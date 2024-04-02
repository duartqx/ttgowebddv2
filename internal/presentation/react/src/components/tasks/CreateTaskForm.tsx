import { useState } from "react";
import { Completed, TaskCreate } from "../../domains/Task";
import Input from "../elements/Input";
import Select from "../elements/Select";
import TextAreaInput from "../elements/TextAreaInput";
import DarkButton from "../elements/DarkButton";

export default function CreateTaskForm({
    setNewTaskHandler,
}: {
    setNewTaskHandler: (t: TaskCreate) => void;
}) {
    const [tag, setTag] = useState("");
    const [sprint, setSprint] = useState(null as Number | null);
    const [description, setDescription] = useState("");
    const [completed, setCompleted] = useState(Completed.INCOMPLETED);

    const submitHandler = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (tag && sprint && description) {
            setNewTaskHandler({
                tag,
                sprint,
                description,
                completed: Boolean(completed),
            });
            setTag("");
            setSprint(null);
            setDescription("");
        }

        e.currentTarget.reset();
    };

    const completedOptions = [
        {
            value: Completed.INCOMPLETED,
            label: "Incompleted",
        },
        {
            value: Completed.COMPLETED,
            label: "Completed",
        },
    ];

    return (
        <form onSubmit={submitHandler} autoComplete="off">
            <Input
                label="Tag"
                onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setTag(e.target.value)
                }
            />
            <Input
                label="Sprint"
                inputType="number"
                onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setSprint(Number(e.target.value))
                }
            />
            <TextAreaInput
                label="Description"
                rows={4}
                onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setDescription(e.target.value)
                }
            />
            <div className="p-4 font-light">
                <label>Completed</label>
                <Select
                    onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setCompleted(Number(e.target.value))
                    }
                    options={completedOptions}
                />
            </div>
            <DarkButton label="Submit" />
        </form>
    );
}
