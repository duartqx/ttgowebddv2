import React, { useContext, useEffect, useState } from "react";
import {
    Task,
    TaskCreate,
    TaskFilter,
    TasksGroupedBySprint,
} from "../domains/Task";
import TaskService from "../services/TaskService";
import Actions from "../components/elements/Actions";
import TaskCardsGroupedBySprints from "../components/tasks/card/TasksGroupedBySprint";
import { SprintsContext } from "../middleware/SprintsContextProvider";
import LoadingSpinnerIcon from "../icons/LoadingSpinnerIcon";

function TaskList({
    tasksGroupedBySprint,
    updateHandler,
}: {
    tasksGroupedBySprint: TasksGroupedBySprint;
    updateHandler: (task: Task) => () => void;
}) {
    return Object.entries(tasksGroupedBySprint).map(([sprint, tasks]) => {
        return (
            <div
                key={`${sprint}__task__cards`}
                className="
                    flex flex-col rounded-md border-2 border-zinc-950 my-4
                    shadow-md shadow-zinc-950 bg-opacity-5 bg-zinc-950
                "
            >
                <TaskCardsGroupedBySprints
                    sprint={sprint.toString()}
                    updateHandler={updateHandler}
                    tasks={tasks}
                />
            </div>
        );
    });
}

export default function Home() {
    const [taskFilter, setTaskFilter] = useState({} as TaskFilter);
    const [tasks, setTasks] = useState([] as TasksGroupedBySprint);
    const [newTask, setNewTask] = useState({} as TaskCreate);
    const [isLoading, setIsLoading] = useState(true);
    const { setToSprints, pullSprintsTrigger } = useContext(SprintsContext);

    useEffect(() => pullSprintsTrigger(), []);

    useEffect(() => {
        setIsLoading(true);

        TaskService.filter(taskFilter).then((tasksGroupedBySprint) => {
            setTasks(tasksGroupedBySprint);
            setIsLoading(false);
        });

    }, [taskFilter]);

    const setTaskHandler = (task: Task) => {
        const tasksCopy = { ...tasks };
        const group = tasksCopy[Number(task.sprint)] || [];

        const taskId = task.id
            ? group.findIndex((t) => t.id === task.id)
            : null;

        if (taskId !== null && taskId !== -1) {
            group[taskId] = task;
            tasksCopy[Number(task.sprint)] = group;
        } else {
            tasksCopy[Number(task.sprint)] = group.concat([task]);
        }

        setTasks(tasksCopy);
    };

    const updateHandler = (task: Task) => () => {
        const updTask: Task = {
            ...task,
            end_at: new Date(),
            completed: true,
        };
        setTaskHandler(updTask);
        TaskService.update(updTask);
    };

    useEffect(() => {
        try {
            if (newTask.tag && newTask.sprint && newTask.description) {
                TaskService.create(newTask).then((task) => {
                    if (task !== null) {
                        setTaskHandler(task);
                        setToSprints(task?.sprint?.valueOf());
                    }
                    setNewTask({} as TaskCreate);
                });
            }
        } catch (e) {
            console.log(e);
        }
    }, [newTask]);

    return (
        <>
            <Actions
                setTaskFilter={setTaskFilter}
                newTaskHandler={(t: TaskCreate) => setNewTask(t)}
            />
            <div className="w-[80vw] flex flex-col justify-center pt-24">
                {isLoading ? (
                    <div className="
                        h-8 w-8 fixed top-[50%] left-[50%]
                        -translate-x-[1rem] -translate-y-[1rem]
                    ">
                        <LoadingSpinnerIcon />
                    </div>
                ) : (
                    <TaskList
                        tasksGroupedBySprint={tasks}
                        updateHandler={updateHandler}
                    />
                )}
            </div>
        </>
    );
}
