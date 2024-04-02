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
        <>
            <form
                onSubmit={handleLogin}
                autoComplete="off"
                className="
                    container m-auto h-[400px] w-[400px] p-6
                    flex flex-col justify-center z-10
                    rounded-md bg-neutral-900 shadow-lg shadow-neutral-950
                "
            >
                <Input
                    label="Email"
                    inputId="login__email"
                    inputType="email"
                    placeholder="email@email.com"
                    isDisabled={isLoading}
                    onChangeHandler={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setEmail(e.target.value)
                    }
                />
                <Input
                    label="Password"
                    inputId="login__password"
                    inputType="password"
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
            from-transparent to-neutral-900 m-auto fixed h-screen w-screen
            "
            ></div>
        </>
    );
}
