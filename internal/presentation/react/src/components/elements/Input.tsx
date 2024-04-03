import React from "react";

type InputProps = {
    label: string;
    inputType?: string;
    placeholder?: string;
    isDisabled?: Boolean;
    onChangeHandler?: React.ChangeEventHandler;
};

export default function Input({
    label,
    inputType,
    placeholder,
    isDisabled,
    onChangeHandler,
}: InputProps) {
    return (
        <div className="flex flex-col p-4 font-light">
            <label>{label}</label>
            <input
                type={inputType || "text"}
                placeholder={placeholder || ""}
                disabled={Boolean(isDisabled)}
                onChange={onChangeHandler}
                className={`
                    rounded-md bg-zinc-900 border-zinc-700
                    transition-all duration-500 ease-in-out w-full
                    ${
                        isDisabled
                            ? "text-zinc-800"
                            : `
                            focus:shadow-inner focus:shadow-indigo-700
                            focus:border-indigo-700 hover:border-indigo-700
                            hover:ring-indigo-700
                            `
                    }
                `}
            />
        </div>
    );
}
