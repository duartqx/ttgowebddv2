import { Task } from "../../../domains/Task";
import CompletedButton from "./CompletedButton";
import InfoButton from "./InfoButton";
import PresentationChartIcon from "../../../icons/PresentationChartIcon";

type TaskCardProps = {
    task: Task;
    updateHandler: () => void;
};

function TaskDescription({
    description,
    ignoreVisibility,
}: {
    description: string;
    ignoreVisibility?: Boolean;
}) {
    return (
        <div
            className={`"text-wrap p-4 cursor-pointer ${
                !ignoreVisibility && "hidden sm:block"
            }`}
            title={description.valueOf()}
        >
            {description.length > 48
                ? description.slice(0, 48) + "..."
                : description}
        </div>
    );
}

export default function TaskCard({ task, updateHandler }: TaskCardProps) {
    return (
        <div
            className="
                rounded-lg border-2 border-zinc-800 flex flex-col
                hover:border-indigo-500 bg-zinc-800
                shadow-md shadow-zinc-950 break-all p-3 sm:h-[60px] h-[160px]
            "
        >
            <div className="flex justify-between sm:h-full h-[50%] items-center">
                <div className="flex justify-between gap-2">
                    <div
                        className="
                        px-4 py-1 rounded-md border-zinc-200
                        bg-zinc-900 text-indigo-400 flex w-[120px]
                        justify-center items-center cursor-default
                    "
                        title={task.tag.valueOf()}
                    >
                        {task.tag.length > 8
                            ? task.tag.slice(0, 5) + "..."
                            : task.tag}
                    </div>
                </div>
                <TaskDescription description={task.description.valueOf()} />

                <div className="flex items-center justify-between gap-2">
                    <InfoButton
                        color="indigo"
                        title={`Start at: ${task.start_at}`}
                        children={<PresentationChartIcon />}
                    />
                    <CompletedButton
                        task={task}
                        updateHandler={updateHandler}
                    />
                </div>
            </div>
            <div className="sm:hidden block h-[50%]">
                <TaskDescription
                    description={task.description.valueOf()}
                    ignoreVisibility={true}
                />
            </div>
        </div>
    );
}
