import {
    ResponseTask,
    Task,
    TaskCreate,
    TaskFilter,
    TasksGroupedBySprint,
} from "../domains/Task";
import HttpClient from "./HttpClient";
import Service from "./Service";

export default class TaskService extends Service {
    static endpoint: string = "/api/tasks/";
    static filterEndpoint: string = "/api/tasks/filter/";
    static sprintsEndpoint: string = "/api/tasks/sprints/";

    private static buildTaskFromResponse(t: ResponseTask) {
        return {
            ...t,
            start_at: new Date(t.start_at),
            end_at: t.end_at ? new Date(t.end_at) : undefined,
        };
    }

    private static validate(t: Task | TaskCreate, update: Boolean = false) {
        switch (true) {
            case !Boolean(t.tag):
                throw new Error("Tasks must have a Tag!");
            case !Boolean(t.description):
                throw new Error("Tasks must have a Description!");
            case !Boolean(t.sprint):
                throw new Error("Tasks must have a Sprint!");
        }
        if (update && !Boolean((t as Task).id)) {
            throw new Error("Tasks must have an Id when updating!");
        }
    }

    public static async create(t: TaskCreate): Promise<Task | null> {
        this.validate(t);

        console.log(t)

        try {
            const res = await HttpClient().post(TaskService.endpoint, {
                ...t,
                sprint: Number(t.sprint),
            });
            const task: ResponseTask = res.data;
            if (!task.id) {
                throw new Error("Task not properly created");
            }
            return this.buildTaskFromResponse(task);
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    public static async update(t: Task): Promise<Task | null> {
        this.validate(t, true);

        try {
            const res = await HttpClient().patch(
                this.getResourceEndpoint(t.id),
                {
                    ...t,
                    sprint: Number(t.sprint)
                }
            );
            const task: ResponseTask = res.data;
            if (!task?.id) {
                throw new Error("Task not properly updated");
            }
            return this.buildTaskFromResponse(task);
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    public static async filter(tf: TaskFilter): Promise<TasksGroupedBySprint> {
        try {
            const res = await HttpClient().post(
                TaskService.filterEndpoint,
                Object.fromEntries(
                    Object.entries(tf).filter(
                        ([key, value]) => value !== undefined
                    )
                )
            );

            const tasksGroupedBy: TasksGroupedBySprint = {};
            for (let i: number = 0; i < res.data.length; i++) {
                const task: ResponseTask = res.data[i];
                const group: Task[] = tasksGroupedBy[Number(task.sprint)] || [];

                tasksGroupedBy[Number(task.sprint)] = group.concat([
                    this.buildTaskFromResponse(task)
                ]);
            }

            return tasksGroupedBy;
        } catch (e) {
            console.log(e);
            return {} as TasksGroupedBySprint;
        }
    }

    public static async sprints(): Promise<Number[]> {
        try {
            const res = await HttpClient().get(TaskService.sprintsEndpoint);
            return res.data;
        } catch (e) {
            console.log(e);
            return [];
        }
    }
}
