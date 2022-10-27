export class DeleteDispatchRequest {
    public url: string
    public dispatchnumbers: string
    constructor(url:string, dispatchnumbers:string) {
        this.url = url;
        this.dispatchnumbers = dispatchnumbers;
    }

    async request() {
        try {
            let headers = new Headers();
            headers.append("Content-Type", "application/x-www-form-urlencoded");

            let urlencoded = new URLSearchParams();
            urlencoded.append("test", this.dispatchnumbers);

            const options = {
                method: "POST",
                headers: headers,
                body: urlencoded
            }

            const req = await fetch(`${this.url}delete`, options)
            const resp = await req.json();

            if(import.meta.env.DEV) console.log(resp)
        } catch(exception) {
            if(exception instanceof Error) {
                if(import.meta.env.DEV) throw new Error(exception.message);
            }
        }
    }
}