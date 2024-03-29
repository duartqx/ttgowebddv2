import { Task, TaskCreate, TaskFilter } from "../domains/Task";
import HttpClient from "./client";

export default class TaskService {
    public static async create(t: TaskCreate): Promise<Task | null> {
        try {
            const res = await HttpClient().post("/api/tasks/", t);
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
            const res = await HttpClient().patch(`/api/tasks/${t.id}/`, t);
            const task: Task = res.data;
            if (!task.id) {
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
            const res = await HttpClient().patch(`/api/tasks/filter/`, tf);
            const tasks: Task[] = res.data;
            return tasks;
        } catch (e) {
            console.log(e);
            const tasks: Task[] = [];
            return tasks;
        }
    }
}
