import React, { useContext } from "react";
import { Navigate, useLocation } from "react-router-dom";
import { AuthContext } from "./AuthContextProvider";
import { Paths } from "../paths";

type PrivateRouterProps = {
    children: React.ReactNode
}

export default function PrivateRouter({ children }: PrivateRouterProps) {
    const { isLoggedIn } = useContext(AuthContext);
    const location = useLocation();

    return isLoggedIn() ? (
        children
    ) : (
        <Navigate
            replace
            to={Paths.login}
            state={{ from: `${location.pathname}${location.search}` }}
        />
    );
}
