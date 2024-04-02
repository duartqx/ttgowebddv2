import React from "react";

type InputProps = {
    label: string;
    inputType: string;
    inputId: string;
    onChangeHandler?: React.ChangeEventHandler;
    placeholder?: string;
    isDisabled?: Boolean;
};

export default function Input({
    label,
    inputId,
    inputType,
    onChangeHandler,
    placeholder,
    isDisabled,
}: InputProps) {
    return (
        <div className="flex flex-col p-4 font-light">
            <label htmlFor={inputId}>{label}</label>
            <input
                id={inputId}
                type={inputType}
                placeholder={placeholder || ""}
                disabled={Boolean(isDisabled)}
                onChange={onChangeHandler}
                className={`
                    rounded-md bg-zinc-900
                    transition-all duration-500 ease-in-out w-full
                    ${
                        isDisabled
                            ? "text-neutral-600"
                            : `
                            focus:shadow-inner focus:shadow-indigo-800
                            focus:border-indigo-800 hover:border-indigo-800
                            hover:ring-indigo-800
                            `
                    }
                `}
            />
        </div>
    );
}
