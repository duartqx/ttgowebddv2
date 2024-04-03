type Option = {
    value: number | string;
    label: string;
};

type SelectProps = {
    options: Option[];
    onChangeHandler?: React.ChangeEventHandler;
};

export default function Select({ options, onChangeHandler }: SelectProps) {
    return (
        <select
            onChange={onChangeHandler}
            className="
                bg-zinc-900 rounded-md w-full
                transition-all
                duration-500
                ease-in-out 
                focus:shadow-inner
                border-zinc-700
                focus:shadow-indigo-700
                focus:border-indigo-700
                hover:border-indigo-700
                hover:ring-indigo-700
            "
        >
            {options.map((o) => (
                <option
                    value={o.value}
                    key={`${o.label.replaceAll(" ", "")}__${o.value}`}
                >
                    {o.label}
                </option>
            ))}
        </select>
    );
}
