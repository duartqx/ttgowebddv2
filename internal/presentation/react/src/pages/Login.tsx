import React, { useContext, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { AuthContext } from "../middleware/AuthContextProvider";

export default function Login() {
    const navigate = useNavigate();
    const location = useLocation();
    const { from } = location.state || { from: { pathname: "/" } };
    const { login } = useContext(AuthContext);

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();

        try {
            if (await login({ email, password })) {
                navigate(from);
            }
        } catch (e) {
            console.log(e);
        }
    };

    return (
        <>
            <form
                onSubmit={handleLogin}
                className="container m-auto w-screen"
            >
                <div className="flex justify-between items-center md:w-3/5 lg:w-2/5 p-3">
                    <label className="flex-grow-0 w-24">Email</label>
                    <input
                        type="email"
                        placeholder="email@email.com"
                        onChange={(e) => setEmail(e.target.value)}
                        className="rounded-md flex-grow"
                    />
                </div>
                <div className="flex justify-between items-center md:w-3/5 lg:w-2/5 p-3">
                    <label className="flex-grow-0 w-24">Password</label>
                    <input
                        type="password"
                        placeholder=""
                        onChange={(e) => setPassword(e.target.value)}
                        className="rounded-md flex-grow"
                    />
                </div>
            </form>
        </>
    );
}
