import TaskCard from "./TaskCard";
import { Task } from "../../../domains/Task";
import { useState } from "react";

type TaskCardsGroupedBySprintsProps = {
    sprint: string;
    tasks: Task[];
    updateHandler: (t: Task) => () => void;
};

export default function TaskCardsGroupedBySprints({
    sprint,
    tasks,
    updateHandler,
}: TaskCardsGroupedBySprintsProps) {
    return (
        <>
            <div
                className={`
                    bg-indigo-300 text-center cursor-pointer
                    text-indigo-500 p-3 text-xl
                    flex gap-2 px-12 items-center rounded-t-md
                `}
                title={`Sprint ${sprint} with ${tasks.length} tasks.`}
            >
                <div className="font-bold">{sprint}</div>
                <div className="text-sm">({tasks.length})</div>
            </div>
            <div className="p-4 flex flex-col gap-4">
                {tasks.map((t) => (
                    <TaskCard
                        task={t}
                        updateHandler={updateHandler(t)}
                        key={`${t.id}__${t.sprint}`}
                    />
                ))}
            </div>
        </>
    );
}
