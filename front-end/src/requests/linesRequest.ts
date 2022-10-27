export class LinesRequest {
    public url:string

    constructor(url:string) {
        this.url = url;
    }

    async request() {
        try {
            const request = await fetch(`${this.url}lines`);
            const response = await request.json();
            if(import.meta.env.DEV) console.log(response);

            return response;
        } catch (exception) {
            if(exception instanceof Error) {
                if(import.meta.env.DEV) throw new Error(exception.message);
            }
        }
    }
}