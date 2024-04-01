import React from "react";

type InputProps = {
    label: string;
    inputType: string;
    inputId: string;
    onChangeHandler?: React.ChangeEventHandler;
    placeholder?: string;
};

export default function Input({
    label,
    inputId,
    inputType,
    onChangeHandler,
    placeholder,
}: InputProps) {
    return (
        <div className="flex flex-col p-4 font-light">
            <label htmlFor={inputId}>{label}</label>
            <input
                id={inputId}
                type={inputType}
                placeholder={placeholder || ""}
                onChange={onChangeHandler}
                className="
                    rounded-md bg-zinc-900
                    focus:shadow-inner focus:shadow-indigo-800 focus:border-indigo-800
                    hover:border-indigo-800 hover:ring-indigo-800
                    transition-all duration-500 ease-in-out w-full
                "
            />
        </div>
    );
}
