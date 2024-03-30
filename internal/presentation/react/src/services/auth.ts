import { LoginResponse } from "../domains/Auth";
import { User } from "../domains/User";
import HttpClient from "./client";
import Service from "./service";

export default class AuthService extends Service {
    static endpoint: string = "/api/user/";
    static loginEndpoint: string = "/api/user/login/";
    static logoutEndpoint: string = "/api/user/logout/";

    public static async login(user: User): Promise<LoginResponse> {
        try {
            const res = await HttpClient().post(AuthService.loginEndpoint, {
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
            return {} as LoginResponse;
        }
    }

    public static async logout(): Promise<void> {
        try {
            const res = await HttpClient().delete(AuthService.logoutEndpoint);

            if (res.status >= 200 && res.status <= 299) {
                localStorage.clear();
            }
            console.log("Logout");
        } catch (e) {
            console.log(e);
        }
    }

    public static async register(user: User): Promise<User> {
        try {
            const res = await HttpClient().post(AuthService.endpoint, user);

            const userResponse: User = res.data;

            return userResponse;
        } catch (e) {
            console.log(e);
            return {} as User;
        }
    }

    public static async getUser(): Promise<User> {
        var user: User = JSON.parse(localStorage.getItem("user") || "{}");

        if (!user?.email) {
            try {
                const res = await HttpClient().get(AuthService.endpoint);

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

    public static getAuth(): LoginResponse {
        return JSON.parse(localStorage.getItem("auth") || "{}");
    }
}
