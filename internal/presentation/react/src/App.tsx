import React from "react";
import "./App.css";
import Login from "./pages/Login";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import LoggedOutRouter from "./middleware/LoggedOutRouter";
import AuthProvider from "./middleware/AuthContextProvider";
import PrivateRouter from "./middleware/PrivateRouter";
import Layout from "./pages/Layout";
import { Paths } from "./paths";
import Home from "./pages/Home";

function App() {
    const router = createBrowserRouter([
        {
            path: Paths.root.pathname,
            element: (
                <PrivateRouter>
                    <Layout />
                </PrivateRouter>
            ),
            children: [
                {
                    index: true,
                    element: (<Home />),
                },
            ],
        },
        {
            path: Paths.login.pathname,
            element: (
                <LoggedOutRouter>
                    <Layout />
                </LoggedOutRouter>
            ),
            children: [{ path: "", element: <Login /> }],
        },
    ]);
    return (
        <>
            <AuthProvider>
                <RouterProvider router={router} />
            </AuthProvider>
        </>
    );
}

export default App;
