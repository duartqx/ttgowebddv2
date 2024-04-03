import LoadingSpinnerIcon from "../../icons/LoadingSpinnerIcon";

type ButtonProps = {
    label: string;
    type?: "submit" | "reset" | "button" | undefined;
    isLoading?: Boolean;
};

export default function DarkButton({ label, type, isLoading }: ButtonProps) {
    return (
        <div className="flex justify-between items-center p-3 w-full">
            <button
                disabled={Boolean(isLoading)}
                type={type || "submit"}
                className={`
                    ${
                        isLoading
                            ? "hover:border-zinc-900 bg-zinc-900"
                            : `hover:border-indigo-700 bg-gradient-to-br
                               from-zinc-900 to-black hover:shadow-indigo-800
                               hover:text-indigo-400
                    `
                    }
                    w-full transition-all duration-500 ease-in-out
                    flex items-center justify-center font-bold hover:shadow-md
                `}
            >
                {isLoading ? <LoadingSpinnerIcon /> : label}
            </button>
        </div>
    );
}
