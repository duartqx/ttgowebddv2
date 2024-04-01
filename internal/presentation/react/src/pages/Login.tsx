import React, { useContext, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { AuthContext } from "../middleware/AuthContextProvider";
import Input from "../components/elements/Input";
import DarkButton from "../components/elements/DarkButton";

export default function Login() {
    const navigate = useNavigate();
    const location = useLocation();
    const { from } = location.state || { from: { pathname: "/" } };
    const { login } = useContext(AuthContext);

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();

        if (!email || !password) {
            console.log(
                `Missing email or password: Email: '${email}'; Password: '${password}'`
            );
            return false;
        }

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
                className="
                    container m-auto lg:w-2/5 md:w-3/5
                    rounded-md bg-neutral-900 shadow-md shadow-neutral-950
                "
            >
                <Input
                    label="Email"
                    inputId="login__email"
                    inputType="email"
                    placeholder="email@email.com"
                    onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setEmail(e.target.value)
                    }
                />
                <Input
                    label="Password"
                    inputId="login__password"
                    inputType="password"
                    onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setPassword(e.target.value)
                    }
                />
                <DarkButton label="Login" />
            </form>
        </>
    );
}
