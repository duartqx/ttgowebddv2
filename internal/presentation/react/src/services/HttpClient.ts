import axios, { AxiosInstance, HttpStatusCode } from "axios";

class HttpClientFacade {
    private _baseUrl = "http://127.0.0.1:8000";
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
