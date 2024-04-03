export default function ActionsTabs({
    which,
    children,
}: {
    which: Boolean;
    children: React.ReactNode;
}) {
    return (
        <div
            className={`
                flex flex-around w-full bg-gradient-to-r rounded-t-md
                ${
                    which
                        ? "from-zinc-800 to-zinc-700"
                        : "from-zinc-700 to-zinc-800"
                }
            `}
        >
            {children}
        </div>
    );
}
