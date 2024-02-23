import { mainStore } from "../store";

export class DispatchSystemStatusesRequest {
    private url:string

    constructor() {
        const store = mainStore();
        const { url } = store;
        this.url = url;
    }

    async req() {
        try {
            const request = await fetch(`${this.url}dispatch_system_statuses`);
            const response = await request.json();
            if(import.meta.env.DEV) console.log(`dispatch system statuse`, response);
            return response;
        } catch (exception) {
            if(import.meta.env.DEV) {
                if(exception instanceof Error) {
                    throw new Error(exception.message)
                }
            }
        }
    }
}