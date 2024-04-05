const infoColors = {
    green: [
        "bg-green-100 text-green-400 hover:shadow-green-400",
        "hover:shadow-green-400 hover:ring-green-500",
    ],
    yellow: [
        "bg-yellow-100 text-yellow-400 hover:shadow-yellow-400",
        "hover:shadow-yellow-400 hover:ring-yellow-500",
    ],
    indigo: [
        "bg-indigo-100 text-indigo-400 hover:shadow-indigo-400",
        "hover:shadow-indigo-400 hover:ring-indigo-500",
    ],
};

type InfoButtonProps = {
    color: "green" | "yellow" | "indigo";
    title: string;
    children: React.ReactNode;
    handlers?: { [key: string]: React.MouseEventHandler };
};

export default function InfoButton({
    color,
    title,
    children,
    handlers,
}: InfoButtonProps) {
    return (
        <div
            className={`
                h-11 w-11 p-2 ${infoColors[color].join(" ")}
                rounded-full flex items-center justify-around
                cursor-pointer transform-all duration-300 ease-in-out
                hover:shadow-md  hover:ring-2
            `}
            title={title}
            {...handlers}
        >
            {children}
        </div>
    );
}
