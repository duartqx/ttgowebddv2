export default function Tab({
    label,
    isSelected,
    onClickHandler,
}: {
    label: string;
    isSelected: () => Boolean;
    onClickHandler: () => void;
}) {
    return (
        <div
            className={`
                        flex-grow text-center p-2 cursor-pointer
                        transform-all duration-200 ease-in-out
                        hover:text-white hover:font-medium
                        ${
                            isSelected()
                                ? "bg-neutral-900 rounded-b-md"
                                : "font-light text-gray-400"
                        } 
                    `}
            onClick={onClickHandler}
        >
            {label}
        </div>
    );
}
