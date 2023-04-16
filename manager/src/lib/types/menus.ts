export interface Menus {
    menus: MenuItem[];
}

export interface MenuItem {
    url: string | null;
    name: string;
    icon: string;
    children: MenuItem[] | null;
}
