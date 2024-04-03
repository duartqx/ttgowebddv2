export default function ActionsFormContainer({
    isHidden,
    children,
}: {
    isHidden: Boolean;
    children: React.ReactNode;
}) {
    return (
        <div
            className={`
                bg-zinc-800 rounded-md flex-grow transform-all
                ${isHidden && "hidden"}
            `}
        >
            {children}
        </div>
    );
}
