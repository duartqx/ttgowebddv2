import React from "react";
import { Outlet } from "react-router-dom";

export default function Layout() {
    return (
        <>
            <div
                className="
                flex items-center justify-center w-screen max-w-[98vw]
            "
            >
                <Outlet />
            </div>
        </>
    );
}
