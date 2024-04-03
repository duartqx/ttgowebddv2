import React, { useState } from "react";
import AuthService from "../services/AuthService";
import { User } from "../domains/User";

export const AuthContext = React.createContext({
    login: async (user: User): Promise<Boolean> => false,

    logout: () => {},

    register: async (user: User): Promise<Boolean> => {
        return false;
    },

    getUser: (): User => {
        return {} as User;
    },

    isLoggedIn: (): Boolean => false,
});

type AuthProviderProps = {
    children: React.ReactNode;
};

export default function AuthProvider({ children }: AuthProviderProps) {
    const emptyUser: User = {} as User;

    const [user, setUser] = useState(emptyUser);

    const login = async (user: User): Promise<Boolean> => {
        if (!user.email || !user.password) {
            return false;
        }

        const loggedIn = Boolean((await AuthService.login(user)).isLoggedIn());
        getUser();
        return loggedIn;
    };

    const logout = () => {
        setUser(emptyUser);
        AuthService.logout().then(() => {
            window.location.reload();
        });
    };

    const register = async (user: User): Promise<Boolean> => {
        return Boolean((await AuthService.register(user))?.email);
    };

    const getUser = (): User => {
        if (AuthService.getAuth().expired()) {
            logout();
            setUser(emptyUser);
            return emptyUser;
        }

        if (!user?.email) {
            AuthService.getUser().then((u) => {
                if (u && u.email) {
                    setUser(u);
                } else {
                    logout();
                }
            });
        }

        return user;
    };

    const isLoggedIn = () => {
        return AuthService.getAuth().isLoggedIn();
    };

    return (
        <AuthContext.Provider
            value={{ login, logout, register, getUser, isLoggedIn }}
        >
            {children}
        </AuthContext.Provider>
    );
}
