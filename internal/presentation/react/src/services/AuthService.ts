import { LoginResponse } from "../domains/Auth";
import AuthEntity from "../domains/AuthEntity";
import { User } from "../domains/User";
import HttpClient from "./HttpClient";
import Service from "./Service";

export default class AuthService extends Service {
    static endpoint: string = "/api/users/";
    static loginEndpoint: string = "/api/users/login/";
    static logoutEndpoint: string = "/api/users/logout/";

    public static async login(user: User): Promise<AuthEntity> {
        try {
            const res = await HttpClient().post(AuthService.loginEndpoint, {
                email: user.email,
                password: user.password,
            });

            const loginResponse: LoginResponse = res.data;

            if (loginResponse && loginResponse.token) {
                localStorage.setItem("auth", JSON.stringify(loginResponse));
                await this.getUser();
            }

            return new AuthEntity(loginResponse);
        } catch (e) {
            console.log(e);
            return new AuthEntity({} as LoginResponse);
        }
    }

    public static async logout(): Promise<void> {
        try {
            await HttpClient()
                .delete(AuthService.logoutEndpoint)
                .then(() => localStorage.clear());
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
                if (this.getAuth().expired()) {
                    return {} as User;
                }

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

    public static getAuth(): AuthEntity {
        return new AuthEntity(JSON.parse(localStorage.getItem("auth") || "{}"));
    }
}
