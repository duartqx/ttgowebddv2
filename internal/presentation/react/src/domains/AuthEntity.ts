import { LoginResponse } from "./Auth";

export default class AuthEntity {
    private token: String;
    private expiresAt: Date;
    private status: Boolean;

    constructor(data: LoginResponse) {
        this.token = data.token || "";
        this.expiresAt = this.setExpiresAt(data.expires_at);
        this.status = Boolean(data.status);
    }

    private setExpiresAt(exp: Date | string | undefined): Date {
        if (exp) {
            return new Date(exp)
        }
        const d = new Date();
        d.setDate(d.getDate() - 99)
        return d
    }

    public isLoggedIn(): Boolean {
        return this.status && this.expiresAt && this.expiresAt > new Date();
    }

    public expired(): Boolean {
        return !this.expiresAt || this.expiresAt < new Date();
    }
}
