export default class Service {
    static endpoint: string = "";

    protected static getResourceEndpoint(id: Number): string {
        return this.endpoint.concat(
            this.endpoint.endsWith("/") ? "" : "/",
            `${id}/`
        );
    }
}
