import { LineInterface } from '../store/index';

export class AddDispatchRequest {
    public url:string
    public dispatchtypecode:string
    public description:string
    public tradecode:string
    public lines:Array<LineInterface>
    public resourse:string
    public randStart:number
    public randEnd:number
    public descriptionIsRandom:boolean

    constructor(url:string, dispatchtypecode:string, description:string, tradecode:string, lines:Array<LineInterface>, resourse:string, randStart:number, randEnd:number, descriptionIsRandom:boolean) {
        this.url = url;
        this.dispatchtypecode = dispatchtypecode;
        this.description = description;
        this.tradecode = tradecode;
        this.lines = lines;
        this.resourse = resourse;
        this.randStart = randStart
        this.randEnd = randEnd
        this.descriptionIsRandom = descriptionIsRandom;
    }

    async request() {
        try {
            let headers = new Headers();
            headers.append("Content-Type", "application/x-www-form-urlencoded");

            let urlencoded = new URLSearchParams();
            urlencoded.append('dispatchtypecode', this.dispatchtypecode);
            if(this.descriptionIsRandom) {
                urlencoded.append('description', `test`);
            } else {
                urlencoded.append('description', this.description);
            }
            urlencoded.append('tradecode', this.tradecode);
            urlencoded.append('lines', JSON.stringify(this.lines));
            urlencoded.append('resourse', this.resourse);
            urlencoded.append('randstart', this.randStart.toString());
            urlencoded.append('randend', this.randEnd.toString());
            urlencoded.append('descriptionIsRandom', this.descriptionIsRandom.toString());

            const options = {
                method: "POST",
                headers: headers,
                body: urlencoded
            };

            const request = await fetch(`${this.url}addDispatch`, options);
            const response = await request.json();
            if(import.meta.env.DEV) console.log(response);

            return response();

        } catch (exception){
            if(exception instanceof Error) {
                if(import.meta.env.DEV) throw new Error(exception.message);
            }
        }
    }
}