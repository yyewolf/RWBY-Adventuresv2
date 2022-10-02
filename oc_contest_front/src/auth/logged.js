import { backend } from "@/plugins/axios";

async function loggedIn() {
    let response = await backend.get("/auth/status");
    let logged = await response.data.logged
    return logged;
}

export { loggedIn };