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
    const [isLoading, setIsLoading] = useState(false);

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();

        if (!email || !password) {
            return false;
        }

        setIsLoading(true);

        try {
            if (await login({ email, password })) {
                navigate(from);
            }
        } catch (e) {
            console.log(e);
        }
        setIsLoading(false);
    };

    return (
        <div className="flex flex-grow mx-auto my-auto h-[90vh] w-screen">
            <form
                onSubmit={handleLogin}
                autoComplete="off"
                className="
                    container m-auto h-[300px] w-[350px] p-6
                    flex flex-col justify-center z-10
                    rounded-md bg-zinc-900 shadow-lg shadow-zinc-950
                "
            >
                <Input
                    label="Email"
                    inputType="email"
                    placeholder="email@email.com"
                    isDisabled={isLoading}
                    onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setEmail(e.target.value)
                    }
                />
                <Input
                    label="Password"
                    inputType="password"
                    placeholder="******"
                    isDisabled={isLoading}
                    onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setPassword(e.target.value)
                    }
                />
                <DarkButton isLoading={isLoading} label="Login" />
            </form>
            <div
                className="
            bg-[radial-gradient(ellipse_at_center,_var(--tw-gradient-stops))]
            from-zinc-800 to-zinc-900 m-auto fixed h-screen w-screen
            "
            ></div>
        </div>
    );
}
