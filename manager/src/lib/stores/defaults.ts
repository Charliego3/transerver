import { browser } from '$app/environment';
import { writable } from 'svelte/store';

type Defaults = {
    sidebarWidth: number,
}

interface UpdateCallback {
    (obj: Defaults): void;
}

function createDefaults() {
    let value;
    if (browser) {
        value = window.localStorage.getItem('defaults');
    }

    let dfs: Defaults = {sidebarWidth: 400};
    if (value) {
        dfs = JSON.parse(value);
    }

    const defaults = writable<Defaults>(dfs);
    defaults.subscribe((obj: Defaults): void => {
        if (browser) {
            window.localStorage.setItem('defaults', JSON.stringify(obj));
        }
    });
    return defaults;
}

export function update(callback: UpdateCallback) {
    defaults.update((v: Defaults) => {
        callback(v);
        return v;
    })
}

export const defaults = createDefaults();
