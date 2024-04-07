import axios, { AxiosInstance, HttpStatusCode } from "axios";

class HttpClientFacade {
    private _baseUrl = this.getBaseURL();
    private _client: AxiosInstance;
    private _token?: String;

    constructor() {
        this._client = axios.create({
            baseURL: this._baseUrl,
        });

        this._client.interceptors.response.use(
            (res) => res,
            (error) => {
                if (
                    error.response &&
                    error.response.status === HttpStatusCode.Unauthorized
                ) {
                    throw new Error("Unauthorized");
                }
                return Promise.reject(error);
            }
        );
    }

    private getBaseURL() {
        return import.meta.env.VITE_DEBUG == "1"
            ? import.meta.env.VITE_DEV_BASE_API_URL
            : import.meta.env.VITE_PROD_BASE_API_URL;
    }

    setToken() {
        const authData = JSON.parse(localStorage.getItem("auth") || "{}");
        this._token = authData?.token;
        return this;
    }

    client(): AxiosInstance {
        this._client.interceptors.request.use(
            (config) => {
                config.headers.Authorization = `Bearer ${this._token}`;
                return config;
            },
            (error) => {
                return Promise.reject(error);
            }
        );

        return this._client;
    }
}

const facade = new HttpClientFacade();

export default function HttpClient(): AxiosInstance {
    return facade.setToken().client();
}
