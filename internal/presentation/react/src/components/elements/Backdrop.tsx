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
                w-screen fixed left-0 right-0 bottom-0
                ${isOpen ? "top-1/2" : "top-[95vh]"}
                bg-gradient-to-t from-neutral-950 to-transparent
                transform-all duration-300 ease-in-out
            `}
        >
            {children}
        </div>
    );
}
