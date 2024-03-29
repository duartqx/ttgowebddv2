import { User } from "./User"

export type Task = {
	id: Number,
	tag: String,
	sprint: String,
	description: String,
	completed: Boolean,
	start_at: Date,
	end_at: Date,
	user: User
}

export enum completedStatus {
	incompleted = 0,
	completed = 1,
	ignored = 2,
}

export type TaskFilter = {
    tag?: String,
	completed: completedStatus,
	sprints: Number[],
	start_at?: Date,
	end_at?: Date,
}