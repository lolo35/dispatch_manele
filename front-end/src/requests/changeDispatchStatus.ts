import { mainStore } from "../store";

export class ChangeDispatchStatusRequest {
    private url:string
    public dispatchnumbers:string
    public dispatchstatus:number

    constructor(dispatchnumbers:string, dispatchstatus:number) {
        const store = mainStore();
        const { url } = store;

        this.url = url;
        this.dispatchnumbers = dispatchnumbers;
        this.dispatchstatus = dispatchstatus;
    }

    async req() {
        try {
            const headers = new Headers();
            headers.append(`Content-Type`, `application/x-www-form-urlencoded`);
            const formData = new URLSearchParams();
            formData.append(`dispatchnumbers`, this.dispatchnumbers);
            formData.append(`status`, this.dispatchstatus.toString())

            const options = {
                method: "POST",
                headers: headers,
                body: formData,
            }

            const request = await fetch(`${this.url}change_dispatch_status`, options)
            const response = await request.json();
            if(import.meta.env.DEV) console.log(`close dispatch response`, response);
            return response;
        } catch (exception) {
            if(import.meta.env.DEV) {
                if(exception instanceof Error) {
                    throw new Error(exception.message);
                }
            }
        }
    }
}