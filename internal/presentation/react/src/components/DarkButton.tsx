type ButtonProps = {
    label: string;
    type?: "submit" | "reset" | "button" | undefined;
};

export default function DarkButton({ label, type }: ButtonProps) {
    return (
        <div className="flex justify-between items-center p-3 w-full">
            <button
                type={type || "submit"}
                className="
                    w-full transition-all duration-500 ease-in-out
                    bg-gradient-to-br from-zinc-950 to-black
                    hover:shadow-sm hover:shadow-indigo-900 hover:border-indigo-800
                "
            >
                {label}
            </button>
        </div>
    );
}
