import { User } from "./User";

export enum Completed {
    INCOMPLETED,
    COMPLETED,
    IGNORED,
}

export type Task = {
    id: Number;
    tag: String;
    sprint: String;
    description: String;
    completed: Boolean;
    start_at: Date;
    end_at: Date;
    user: User;
};

export type TaskCreate = {
    tag: String;
    sprint: Number;
    description: String;
    completed: Boolean;
};

export type TaskFilter = {
    tag?: String;
    completed?: Completed;
    sprints?: Number[];
    start_at?: Date;
    end_at?: Date;
};
