const infoColors = {
    green: "bg-green-100 text-green-400 hover:shadow-green-400",
    yellow: "bg-yellow-100 text-yellow-400 hover:shadow-yellow-400",
    indigo: "bg-indigo-100 text-indigo-400 hover:shadow-indigo-400",
};

type InfoButtonProps = {
    color: "green" | "yellow" | "indigo";
    title: string;
    children: React.ReactNode;
    onClickHandler?: React.MouseEventHandler;
    onMouseEnterHandler?: React.MouseEventHandler;
    onMouseLeaveHandler?: React.MouseEventHandler;
};

export default function InfoButton({
    color,
    title,
    children,
    onClickHandler,
    onMouseEnterHandler,
    onMouseLeaveHandler,
}: InfoButtonProps) {
    return (
        <div
            className={`
                h-8 w-8  ${infoColors[color]}
                rounded-full flex items-center justify-around
                cursor-pointer transform-all duration-300 ease-in-out
                hover:shadow-md hover:shadow-green-400
            `}
            title={title}
            onClick={onClickHandler}
            onMouseEnter={onMouseEnterHandler}
            onMouseLeave={onMouseLeaveHandler}
        >
            {children}
        </div>
    );
}
