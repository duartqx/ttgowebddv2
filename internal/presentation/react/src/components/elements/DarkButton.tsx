import LoadingSpinner from "../../icons/LoadingSpinner";

type ButtonProps = {
    label: string;
    type?: "submit" | "reset" | "button" | undefined;
    isLoading?: Boolean;
};

export default function DarkButton({ label, type, isLoading }: ButtonProps) {
    const backgroundColors = isLoading
        ? "hover:border-neutral-900 bg-neutral-900"
        : `
            hover:border-indigo-800 bg-gradient-to-br
            from-neutral-950 to-black hover:shadow-indigo-900
        `;

    return (
        <div className="flex justify-between items-center p-3 w-full">
            <button
                disabled={Boolean(isLoading)}
                type={type || "submit"}
                className={`
                    w-full transition-all duration-500 ease-in-out
                    flex items-center justify-center font-bold hover:shadow-sm
                    ${backgroundColors}
                `}
            >
                {isLoading ? <LoadingSpinner /> : label}
            </button>
        </div>
    );
}
