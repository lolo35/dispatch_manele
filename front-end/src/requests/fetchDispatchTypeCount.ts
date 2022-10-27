export class FetchDispatchtypeCount {
    public url:string
    public dispatchtypecode:string

    constructor(url:string, dispatchtypecode:string) {
        this.url = url;
        this.dispatchtypecode = dispatchtypecode;
    }

    async req() {
        try {
            const request = await fetch(`${this.url}description_count?dispatchtypecode=${this.dispatchtypecode}`)
            const response = await request.json();

            if(import.meta.env.DEV) console.log(response)
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