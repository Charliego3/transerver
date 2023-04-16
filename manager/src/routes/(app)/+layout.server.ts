import type { LayoutServerLoad } from './$types';
import type { Menus } from "$lib/types/nav";

export const load: LayoutServerLoad = (async (): Promise<Menus> => {
    return {
        menus: [
            {
                url: '/',
                name: 'Dashboard',
                icon: 'https://file.dd.net/statics/img/v4/app_icon/icon_default/new_icon_usdt_64.png',
            },
            {
                url: '/settings',
                name: 'Settings',
                icon: 'https://file.dd.net/statics/img/v4/app_icon/icon_default/new_icon_usdt_64.png',
            },
            {
                url: '/menu1',
                name: 'Menu1',
                icon: 'https://file.dd.net/statics/img/v4/app_icon/icon_default/new_icon_usdt_64.png',
            }
        ]
    }
});
