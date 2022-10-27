export class ImportDispatchtypeDescriptions {
    public url:string
    public dispatchtypecode:string
    
    constructor(url:string,dispatchtypecode:string) {
        this.url = url;
        this.dispatchtypecode = dispatchtypecode;
    }

    async req() {
        try {
            const headers = new Headers();
            headers.append("Content-Type", "application/x-www-form-urlencoded");

            const urlencoded = new URLSearchParams();
            urlencoded.append("dispatchtypecode", this.dispatchtypecode);

            const options = {
                method: "POST",
                headers: headers,
                body: urlencoded,
            };

            const request = await fetch(`${this.url}save_description`, options)
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