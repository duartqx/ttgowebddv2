import React, { useState } from "react";
import AuthService from "../services/auth";
import { User } from "../domains/User";

export const AuthContext = React.createContext({
    login: async (user: User): Promise<Boolean> => false,

    logout: async () => {},

    register: async (user: User): Promise<Boolean> => {
        return false;
    },

    getUser: async (): Promise<User | null> => {
        return null;
    },

    isLoggedIn: (): Boolean => false,
});

type AuthProviderProps = {
    children: React.ReactNode
}

export default function AuthProvider({ children }: AuthProviderProps) {
    const emptyUser: User = { email: "", password: "", name: "" };

    const [user, setUser] = useState(emptyUser);

    const login = async (user: User): Promise<Boolean> => {
        const loginResponse = await AuthService.login(user);
        return Boolean(loginResponse?.status);
    };

    const logout = async () => {
        await AuthService.logout();
        setUser(emptyUser);
    };

    const register = async (user: User): Promise<Boolean> => {
        const signUpUser = await AuthService.register(user);
        return signUpUser?.email ? true : false;
    };

    const getUser = async (): Promise<User | null> => {
        const exp = AuthService.getAuth()?.expires_at;
        if (!exp) {
            return null;
        }

        if (new Date(exp) < new Date()) {
            logout();
            return null;
        }

        if (!user?.email) {
            const authUser = await AuthService.getUser();

            if (authUser && authUser.email) {
                setUser(authUser);
            } else {
                logout();
            }
        }
        return user;
    };

    const isLoggedIn = () => {
        const exp = AuthService.getAuth()?.expires_at;
        if (!exp) {
            return false;
        }
        return Boolean(new Date(exp) > new Date());
    };

    return (
        <AuthContext.Provider
            value={{ login, logout, register, getUser, isLoggedIn }}
        >
            {children}
        </AuthContext.Provider>
    );
}
