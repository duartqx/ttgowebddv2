export default function ActionsFormContainer({
    label,
    isHidden,
    children,
}: {
    label: string;
    isHidden: Boolean;
    children: React.ReactNode;
}) {
    return (
        <div
            className={`
                bg-neutral-900 rounded-md flex-grow transform-all
                ${isHidden && "hidden"}
            `}
        >
            <div
                className="
            bg-neutral-950 rounded-t-md p-4
            text-center font-semibold
            "
            >
                {label}
            </div>
            {children}
        </div>
    );
}
