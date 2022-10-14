import { defineStore } from 'pinia';

export const mainStore = defineStore({
    id: "mainStore",
    state: ():Store => {
        return {
            url: "",
            dispatchType: null,
        }
    },
    actions: {
        setUrl(value: string) {
            this.$state.url = value;
        }
    }
});

interface Store {
    url: string,
    dispatchType: DispatchTypeInterface | null,
}

export interface DispatchTypeInterface {
    id: number,
    site: number,
    code: string,
    description: string,
}