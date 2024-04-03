import React, { useEffect, useState } from "react";
import { Task, TaskCreate, TaskFilter } from "../domains/Task";
import TaskService from "../services/TaskService";
import Actions from "../components/elements/Actions";

export default function Home() {
    const [taskFilter, setTaskFilter] = useState({} as TaskFilter);
    const [tasks, setTasks] = useState([] as Task[]);
    const [newTask, setNewTask] = useState({} as TaskCreate);

    useEffect(() => {
        TaskService.filter(taskFilter).then((tks) => setTasks(tks));
    }, [taskFilter]);

    useEffect(() => {
        try {
            if (newTask.tag && newTask.sprint && newTask.description) {
                TaskService.create(newTask).then((task) => {
                    if (task !== null) {
                        setTasks(tasks.concat(task));
                    }
                });
            }
        } catch (e) {
            console.log(e);
        }
    }, [newTask]);

    return (
        <>
            <div>
                <Actions
                    setTaskFilter={setTaskFilter}
                    newTaskHandler={(t: TaskCreate) => setNewTask(t)}
                />
                <div className="w-[80vw] mx-auto flex flex-wrap pt-24">
                    {tasks.map((t) => (
                        <div
                            key={`${t.id}__${t.sprint}`}
                            className="
                                rounded-lg border-2 border-zinc-800 
                                hover:border-indigo-500 bg-zinc-800 m-2 
                                w-[240px] h-[200px] shadow-md shadow-zinc-950
                                mx-auto break-all p-3
                            "
                        >
                            <div className="w-full flex justify-between p-2 items-center">
                                <div
                                    className="
                                        px-4 rounded-md border-zinc-200
                                        bg-zinc-100 text-indigo-400 flex
                                        justify-center items-center cursor-pointer
                                    "
                                    title={t.tag.valueOf()}
                                >
                                    {t.tag.length > 8
                                        ? t.tag.slice(0, 6) + "..."
                                        : t.tag}
                                </div>
                                <div
                                    title={`Sprint: ${t.sprint}`}
                                    className="
                                        px-2 rounded-md border-indigo-400
                                        bg-indigo-200 text-indigo-500 flex
                                        text-sm justify-center items-center
                                        cursor-pointer
                                    "
                                >
                                    {t.sprint}
                                </div>
                            </div>
                            <div
                                className="
                                    text-wrap font-light p-4 border-b-[0.5px]
                                    border-b-indigo-200 h-[80px]"
                                title={t.description.valueOf()}
                            >
                                {t.description.length > 22
                                    ? t.description.slice(0, 22) + "..."
                                    : t.description}
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </>
    );
}
