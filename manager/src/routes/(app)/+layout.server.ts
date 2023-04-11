import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = (async () => {
    return {
        menus: [
            {
                url: '/',
                name: 'Dashboard',
                icon: '',
            },
            {
                url: '/login',
                name: 'Menu1',
                icon: '',
            },
            {
                url: '/register',
                name: 'Menu2',
                icon: '',
            }
        ]
    };
});
