import { writable } from "svelte/store";

export const pageTitle = writable('');

export const appName = writable('');

export const hideControls = writable(false);

//pocketcms settings
export const pocketCMSSettings = writable({});

export function setPocketCmsSettings(model) {
    pocketCMSSettings.set(model || {});
}