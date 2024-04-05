import { Task } from "../../../domains/Task";
import CompletedButton from "./CompletedButton";
import InfoButton from "./InfoButton";
import PresentationChartIcon from "../../../icons/PresentationChartIcon";

type TaskCardProps = {
    task: Task;
    updateHandler: () => void;
};

export default function TaskCard({ task, updateHandler }: TaskCardProps) {
    return (
        <div
            className="
                rounded-lg border-2 border-zinc-800  flex justify-between
                hover:border-indigo-500 bg-zinc-800 items-center
                shadow-md shadow-zinc-950 break-all p-3 h-[60px]
            "
        >
            <div className="flex justify-between gap-2">
                <div
                    className="
                        px-4 py-1 rounded-md border-zinc-200
                        bg-zinc-100 text-indigo-400 flex w-[120px]
                        justify-center items-center cursor-default
                    "
                    title={task.tag.valueOf()}
                >
                    {task.tag.length > 8
                        ? task.tag.slice(0, 5) + "..."
                        : task.tag}
                </div>
            </div>
            <div
                className="text-wrap font-light p-4 cursor-pointer"
                title={task.description.valueOf()}
            >
                {task.description.length > 48
                    ? task.description.slice(0, 48) + "..."
                    : task.description}
            </div>

            <div className="flex items-center justify-between gap-2">
                <InfoButton
                    color="indigo"
                    title={`Start at: ${task.start_at}`}
                    children={<PresentationChartIcon />}
                />
                <CompletedButton task={task} updateHandler={updateHandler} />
            </div>
        </div>
    );
}
