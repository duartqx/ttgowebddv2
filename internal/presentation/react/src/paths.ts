import { Path } from "react-router-dom";

type PathsType = {
    root: Partial<Path>;
    login: Partial<Path>;
    register: Partial<Path>;
};

export const Paths: PathsType = {
    root: { pathname: "/" },
    login: { pathname: "/login" },
    register: { pathname: "/register" },
};
