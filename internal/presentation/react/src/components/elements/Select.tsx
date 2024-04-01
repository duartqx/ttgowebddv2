type Option = {
    value: number | string;
    label: string;
};

type SelectProps = {
    id?: string;
    options: Option[];
    onChangeHandler?: React.ChangeEventHandler;
};

export default function Select({ id, options, onChangeHandler }: SelectProps) {
    return (
        <select
            id={id}
            onChange={onChangeHandler}
            className="
                bg-zinc-900 rounded-md w-full
                transition-all
                duration-500
                ease-in-out 
                focus:shadow-inner
                focus:shadow-indigo-800
                focus:border-indigo-800
                hover:border-indigo-800
                hover:ring-indigo-800
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
