import type { Actions } from "./$types";
import { redirect } from "@sveltejs/kit";

type User = {
    uname: string;
    password: string;
}

export const actions: Actions = {
    login: async ({ request }) => {
        const data = Object.fromEntries(await request.formData()) as User;
        console.log({data});
        throw redirect(301, '/');
        return { success: true };
    }
};
