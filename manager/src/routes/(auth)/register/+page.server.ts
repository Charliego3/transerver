import type { Actions } from "./$types";

type User = {
    uname: string;
    email: string;
    password: string;
    confirm: string;
    privacy: string;
}

export const actions: Actions = {
    register: async ({ request }) => {
        const data = Object.fromEntries(await request.formData()) as User;
        console.log({data});
        return { success: true };
    }
};
