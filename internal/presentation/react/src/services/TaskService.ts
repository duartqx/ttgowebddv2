import { Task, TaskCreate, TaskFilter } from "../domains/Task";
import HttpClient from "./HttpClient";
import Service from "./Service";

export default class TaskService extends Service {
    static endpoint: string = "/api/tasks/";
    static filterEndpoint: string = "/api/tasks/filter/";
    static sprintsEndpoint: string = "/api/tasks/sprints/";

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

        try {
            const res = await HttpClient().post(TaskService.endpoint, t);
            const task: Task = res.data;
            if (!task.id) {
                throw new Error("Task not properly created");
            }
            return task;
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
                t
            );
            const task: Task = res.data;
            if (!task?.id) {
                throw new Error("Task not properly updated");
            }
            return task;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    public static async filter(tf: TaskFilter): Promise<Task[]> {
        try {
            const res = await HttpClient().post(
                TaskService.filterEndpoint,
                Object.fromEntries(
                    Object.entries(tf).filter(
                        ([key, value]) => value !== undefined
                    )
                )
            );
            const tasks: Task[] = res.data;
            return tasks;
        } catch (e) {
            console.log(e);
            const tasks: Task[] = [];
            return tasks;
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
