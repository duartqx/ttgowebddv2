import React from "react";
import { Outlet } from "react-router-dom";

export default function Layout() {
    return (
        <div className="flex flex-col min-h-[100vh]">
            <div
                className="flex justify-center items-center w-screen max-w-[97vw]"
            >
                <Outlet />
            </div>
            <div className="z-10 p-8 font-light text-center cursor-default mt-auto">
                ttgowebddv2 © {new Date().getFullYear()}
            </div>
        </div>
    );
}
