export default function Backdrop({
    isOpen,
    children,
}: {
    isOpen: Boolean;
    children: React.ReactNode;
}) {
    return (
        <div
            className={`
                w-screen fixed left-0 right-0 top-0
                ${isOpen ? "bottom-1/4" : "bottom-[90vh]"}
                bg-gradient-to-b from-zinc-950 to-transparent
                transform-all duration-300 ease-in-out
            `}
        >
            {children}
        </div>
    );
}
