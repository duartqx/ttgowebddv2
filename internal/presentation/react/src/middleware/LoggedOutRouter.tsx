import React, { useContext } from "react";
import { Navigate, useLocation } from "react-router-dom";
import { AuthContext } from "./AuthContextProvider";
import { Paths } from "../paths";

type PrivateRouterProps = {
    children: React.ReactNode;
};

export default function LoggedOutRouter({ children }: PrivateRouterProps) {
    const { isLoggedIn } = useContext(AuthContext);
    const location = useLocation();

    return isLoggedIn() ? (
        <Navigate
            replace
            to={Paths.login}
            state={{ from: `${location.pathname}${location.search}` }}
        />
    ) : (
        children
    );
}
