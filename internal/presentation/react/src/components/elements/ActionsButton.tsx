import ChevronUp from "../../icons/ChevronUp";

export default function ActionsButton({
    active,
    setActive,
}: {
    active: Boolean;
    setActive: () => void;
}) {
    return (
        <button
            type="button"
            onClick={setActive}
            className="
                fixed rounded-md border-gray-600 flex justify-center
                left-1/2 bottom-4 w-[200px] -translate-x-[100px]
            "
        >
            <div className="font-light mr-2">Actions</div>
            <div
                className={`
                    transform-all duration-200 ease-in-out
                    ${active && "scale-y-[-1]"}
                `}
            >
                <ChevronUp />
            </div>
        </button>
    );
}
