import { Task, TaskCreate, TaskFilter } from "../domains/Task";
import HttpClient from "./client";
import Service from "./service";

export default class TaskService extends Service {
    static endpoint: string = "/api/tasks/";
    static filterEndpoint: string = "/api/tasks/filter/";

    public static async create(t: TaskCreate): Promise<Task | null> {
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
            const res = await HttpClient().patch(
                TaskService.filterEndpoint,
                tf
            );
            const tasks: Task[] = res.data;
            return tasks;
        } catch (e) {
            console.log(e);
            const tasks: Task[] = [];
            return tasks;
        }
    }
}
