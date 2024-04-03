type SprintButtonProps = {
    sprint: string;
    isSelected: Boolean;
    toggleSelected: () => void;
};
export default function SprintButton({
    sprint,
    isSelected,
    toggleSelected,
}: SprintButtonProps) {
    return (
        <button
            className={`
                ${
                    isSelected
                        ? "bg-zinc-950 border-indigo-500 text-indigo-400"
                        : "bg-zinc-800 border-gray-500"
                }
                m-1 shadow-md shadow-zinc-950
                focus:outline-none hover:shadow-indigo-950
                transition-all duration-500 ease-in-out
            `}
            value={sprint.toString()}
            type="button"
            onClick={toggleSelected}
        >
            {sprint.toString()}
        </button>
    );
}
