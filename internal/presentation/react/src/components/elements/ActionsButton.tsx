import ChevronUpIcon from "../../icons/ChevronUpIcon";

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
                w-[140px] flex justify-center ml-[50%]
                -translate-x-[70px] bg-transparent focus:outline-none
            "
        >
            <div className="mr-2">Actions</div>
            <div
                className={`
                    transform-all duration-200 ease-in-out
                    ${!active && "scale-y-[-1]"}
                `}
            >
                <ChevronUpIcon />
            </div>
        </button>
    );
}
