import React, { useState, useEffect } from "react";
import TaskService from "../../services/TaskService";
import Select from "../elements/Select";
import Input from "../elements/Input";
import DarkButton from "../elements/DarkButton";
import { Completed, TaskFilter } from "../../domains/Task";
import SprintButton from "../elements/SprintButton";

type FilterTaskProps = {
    setTaskFilter: (tf: TaskFilter) => void;
    dismissForm: () => void;
};

type SprintsState = {
    [sprint: string]: Boolean;
};

export default function FilterTasksForm({
    setTaskFilter,
    dismissForm,
}: FilterTaskProps) {
    const [completed, setCompleted] = useState(Completed.IGNORED);
    const [startAt, setStartAt] = useState("");
    const [endAt, setEndAt] = useState("");
    const [sprints, setSprints] = useState({} as SprintsState);

    useEffect(() => {
        TaskService.sprints().then((s) =>
            setSprints(
                s.reduce((acc: SprintsState, curr: Number) => {
                    acc[curr.valueOf()] = false;
                    return acc;
                }, {} as SprintsState)
            )
        );
    }, []);

    const submitHandler = (e: React.FormEvent) => {
        e.preventDefault();

        setTaskFilter({
            completed: completed,
            start_at: startAt ? new Date(startAt) : undefined,
            end_at: endAt ? new Date(endAt) : undefined,
            sprints: Object.entries(sprints)
                .filter(([sprint, selected]) => selected)
                .map(([sprint, selected]) => Number(sprint)),
        });

        dismissForm();
    };

    const completedChangeHandler = (e: React.ChangeEvent<HTMLInputElement>) =>
        setCompleted(Number(e.target.value) as Completed);

    const completedOptions = [
        { value: Completed.IGNORED, label: "Select" },
        {
            value: Completed.COMPLETED,
            label: "Completed",
        },
        {
            value: Completed.INCOMPLETED,
            label: "Incompleted",
        },
    ];

    const toggleSelectedSprint = (sprint: string) => {
        const sprintsCopy = { ...sprints };
        sprintsCopy[sprint] = !Boolean(sprints[sprint]);
        setSprints(sprintsCopy);
    };

    const sprintIsSelected = (sprint: string): Boolean => sprints[sprint];

    return (
        <form onSubmit={submitHandler}>
            <div className="flex-flex-col p-4">
                <label className="font-light">Status</label>
                <Select
                    onChangeHandler={completedChangeHandler}
                    options={completedOptions}
                />
            </div>
            <div className="flex flex-col p-4">
                <label className="font-light">Sprint</label>
                <div
                    className="
                    self-center flex flex-wrap justify-center
                    transform-all duration-500 ease-in-out
                "
                >
                    {Object.keys(sprints).map((s) => (
                        <SprintButton
                            key={`sprint__${s}`}
                            sprint={s}
                            isSelected={sprintIsSelected(s)}
                            toggleSelected={() => toggleSelectedSprint(s)}
                        />
                    ))}
                </div>
            </div>
            <Input
                inputType="date"
                label="Start At"
                onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setStartAt(e.target.value)
                }
            />

            <Input
                inputType="date"
                label="End At"
                onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setEndAt(e.target.value)
                }
            />
            <DarkButton label="Submit" />
        </form>
    );
}
