import { useState } from "react";
import { Task } from "../../../domains/Task";
import CheckCircle from "../../../icons/CheckCircle";
import ExclamationCircleIcon from "../../../icons/ExclamationCircleIcon";
import RedoIcon from "../../../icons/RedoIcon";
import InfoButton from "./InfoButton";

type CompletedButtonProps = {
    task: Task;
    updateHandler: () => void;
};

export default function CompletedButton({
    task,
    updateHandler,
}: CompletedButtonProps) {
    const [isHovered, setIsHovered] = useState(false);
    const handlers = {
        onClick: updateHandler,
        onMouseEnter: () => setIsHovered(true),
        onMouseLeave: () => setIsHovered(false),
    };

    return (
        <div className="flex items-center py-4">
            {task.completed ? (
                <InfoButton
                    color="green"
                    title={`Update completion date from: ${
                        task.end_at && task.end_at
                    } to now`}
                    children={isHovered ? <RedoIcon /> : <CheckCircle />}
                    handlers={handlers}
                />
            ) : (
                <InfoButton
                    color="yellow"
                    title={`Update to Completed`}
                    children={
                        isHovered ? <RedoIcon /> : <ExclamationCircleIcon />
                    }
                    handlers={handlers}
                />
            )}
        </div>
    );
}
