import axios, { AxiosInstance } from "axios";

const baseUrl = "http://127.0.0.1:8080";

const client = axios.create({
    baseURL: baseUrl,
});

client.interceptors.response.use(
    (res) => res,
    (error) => {
        if (error.response && error.response.status === 401) {
            localStorage.clear();
        }
        return Promise.reject(error);
    }
);

function getToken(): String {
    const authData = JSON.parse(localStorage.getItem("auth") || "{}");
    return authData?.token || "";
}

export default function HttpClient(): AxiosInstance {
    client.interceptors.request.use(
        (config) => {
            config.headers.Authorization = `Bearer ${getToken()}`;
            return config;
        },
        (error) => {
            return Promise.reject(error);
        }
    );

    return client;
}
