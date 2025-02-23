import { defineStore } from 'pinia';
import {bigUploadStates} from "/frontend/src/stores/interface";

export const bigUpload = defineStore('bigUpload', {
    state:():bigUploadStates=>({
        panelShow:false
    }),
    actions: {
        async setPanelShow(bool: boolean) {
            this.panelShow = bool;
        }
    },
})
