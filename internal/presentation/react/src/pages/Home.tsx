import React, { useContext, useEffect, useState } from "react";
import { AuthContext } from "../middleware/AuthContextProvider";
import { Task, TaskFilter } from "../domains/Task";
import TaskService from "../services/TaskService";
import FilterTasks from "../components/FilterTasks";

export default function Home() {
    const { logout } = useContext(AuthContext);

    const [taskFilter, setTaskFilter] = useState({} as TaskFilter);
    const [tasks, setTasks] = useState([] as Task[]);

    useEffect(() => {
        TaskService.filter(taskFilter).then((tks) => setTasks(tks));
    }, [taskFilter]);

    return (
        <>
            <div className="container mx-auto lg:w-2/5 md:w-3/5">
                <FilterTasks setTaskFilter={setTaskFilter} />
                <div>
                    <table>
                        <thead>
                            <tr>
                                <td>Tag</td>
                                <td>Sprint</td>
                                <td>Description</td>
                                <td>Completed</td>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                tasks.map((t) => (
                                    <tr>
                                        <td>{t.tag}</td>
                                        <td>{t.sprint}</td>
                                        <td>{t.description}</td>
                                        <td>{`${t.completed}`}</td>
                                    </tr>
                                ))
                            }
                        </tbody>
                    </table>
                </div>
                <button className="w-full p-4" onClick={() => logout()}>
                    Logout
                </button>
            </div>
        </>
    );
}
