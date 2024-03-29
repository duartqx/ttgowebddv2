import { LoginResponse } from "../domains/Auth";
import { User } from "../domains/User";
import HttpClient from "./client";

export default class AuthService {
    public static async login(user: User): Promise<LoginResponse | null> {
        try {
            const res = await HttpClient().post("/login", {
                email: user.email,
                password: user.password,
            });

            const loginResponse: LoginResponse = res.data;

            if (loginResponse && loginResponse.token) {
                localStorage.setItem("auth", JSON.stringify(loginResponse));
            }

            return loginResponse;
        } catch (e) {
            console.log(e);
            return null;
        }
    }
    public static async logout(): Promise<void> {
        try {
            const res = await HttpClient().delete("/logout");

            if (res.status >= 200 && res.status <= 299) {
                localStorage.clear();
            }
            console.log("Logout");
        } catch (e) {
            console.log(e);
        }
    }
    public static async register(user: User): Promise<User | null> {
        try {
            const res = await HttpClient().post("/api/user/", user);

            const userResponse: User = res.data;

            return userResponse;
        } catch (e) {
            console.log(e);
            return null;
        }
    }
    public static async getUser(): Promise<User | null> {
        var user: User = JSON.parse(localStorage.getItem("user") || "null");

        if (!user?.email) {
            try {
                const res = await HttpClient().get("/api/user/");

                if (res.data?.email) {
                    user = res.data;
                    localStorage.setItem("user", JSON.stringify(user));
                }
            } catch (e) {
                console.log(e);
            }
        }

        return user;
    }
    public static getAuth(): LoginResponse | null {
        return JSON.parse(localStorage.getItem("auth") || "null");
    }
}
