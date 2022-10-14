export class DispatchTypeRequest {
    public url:string

    constructor(url:string) {
        this.url = url;
    }

    async request() {
        try {
            const req = await fetch(`${this.url}dispatchtypes`);
            const resp = await req.json();
            if(import.meta.env.DEV) console.log(resp)

            return resp;
        } catch (exception) {
            if(exception instanceof Error) {
                if(import.meta.env.DEV) throw new Error(exception.message);
            }
        }
    }
}