import { defineStore } from 'pinia';

export const mainStore = defineStore({
    id: "mainStore",
    state: ():Store => {
        return {
            url: "",
            dispatchType: null,
            line: [],
            resourseid: "",
            description: "",
            tradecode: "",
            randStart: 1,
            randEnd: 1,
            descriptionIsRandom: false,
        }
    },
    actions: {
        setUrl(value: string) {
            this.$state.url = value;
        },
        setDispatchType(value: DispatchTypeInterface) {
            this.$state.dispatchType = value;
        },
        setLine(value: LineInterface) {
            this.$state.line.push(value);
        },
        removeLine(index: number) {
            for(let i = 0; i < this.$state.line.length; i++) {
                if(this.$state.line[i].id === index) {
                    this.$state.line.splice(i);
                    return;
                }
            }
        },
        resetLines() {
            this.$state.line = [];
        },
        selectAllLines(value: LineInterface[]) {
            this.$state.line = value;
        },
        setResourseid(value: string) {
            this.$state.resourseid = value;
        },
        setDescription(value: string) {
            this.$state.description = value;
        },
        setTradecode(value: string) {
            this.$state.tradecode = value;
        },
        setRandStart(value: number) {
            this.$state.randStart = value;
        },
        setRandEnd(value:number) {
            this.$state.randEnd = value;
        },
        setDescriptionIsRandom(value: boolean) {
            this.$state.descriptionIsRandom = value;
        }
    }
});

interface Store {
    url: string,
    dispatchType: DispatchTypeInterface | null,
    line: LineInterface[],
    resourseid: string,
    description: string,
    tradecode: string,
    randStart: number,
    randEnd: number,
    descriptionIsRandom: boolean,
}

export interface DispatchTypeInterface {
    id: number,
    site: number,
    code: string,
    description: string,
}

export interface LineInterface {
    id: number,
    code: string,
    areacode: string,
    area: number,
    description: string,
    abbreviation: string,
    defaultmachine: number
}