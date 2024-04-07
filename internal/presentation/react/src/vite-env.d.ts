/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VITE_DEV_BASE_API_URL: string;
    readonly VITE_PROD_BASE_API_URL: string;
    readonly VITE_DEBUG: string;
}

interface ImportMeta {
    readonly env: ImportMetaEnv;
}
