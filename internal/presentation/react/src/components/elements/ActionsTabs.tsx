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
                flex flex-around w-full bg-gradient-to-r rounded-b-md mt-4
                ${
                    which
                        ? "from-neutral-900 to-neutral-800"
                        : "from-neutral-800 to-neutral-900"
                }
            `}
        >
            {children}
        </div>
    );
}
