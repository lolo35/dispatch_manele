export class ChangeDispatchTypeRequest {
    private url:string;
    private dispatchnumbers:string;
    private dispatchtype:number;
    private isClosed:boolean;

    constructor(url:string, dispatchnumbers:string, dispatchtype:number, isClosed:boolean) {
        this.url = url;
        this.dispatchnumbers = dispatchnumbers;
        this.dispatchtype = dispatchtype;
        this.isClosed = isClosed;
    }

    async send() {
        try {
            const headers = new Headers();
            headers.append(`Content-Type`, `application/x-www-form-urlencoded`);
            const body = new URLSearchParams();
            let isClosed = this.isClosed ? 1 : 0;
            body.append(`dispatchnumbers`, this.dispatchnumbers);
            body.append(`dispatchtype`, this.dispatchtype.toString());
            body.append(`is_closed`, `${isClosed}`);

            const options = {
                method: `POST`,
                headers: headers,
                body: body,
            }

            const request = await fetch(`${this.url}change_dispatch_type`, options);
            const response = await request.json();
            if(import.meta.env.DEV) console.log(`change dispatch type response`, response);
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