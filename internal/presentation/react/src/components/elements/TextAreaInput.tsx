import React from "react";

type InputProps = {
    label: string;
    rows: number;
    inputId: string;
    onChangeHandler?: React.ChangeEventHandler;
};

export default function TextAreaInput({
    label,
    rows,
    inputId,
    onChangeHandler,
}: InputProps) {
    return (
        <div className="flex flex-col p-4 font-light">
            <label htmlFor={inputId}>{label}</label>
            <textarea
                id={inputId}
                rows={rows}
                onChange={onChangeHandler}
                className="
                    rounded-md bg-zinc-900 max-h-[140px]
                    focus:shadow-inner focus:shadow-indigo-800 focus:border-indigo-800
                    hover:border-indigo-800 hover:ring-indigo-800
                    transition-all duration-500 ease-in-out w-full
                "
            />
        </div>
    );
}
