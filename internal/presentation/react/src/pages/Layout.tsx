import React from "react";
import { Outlet } from "react-router-dom";

export default function Layout() {
    return (
        <>
            <div
                className="
                flex items-start justify-center
                px-16 w-screen h-screen bg-neutral-800
            "
            >
                <Outlet />
            </div>
        </>
    );
}
