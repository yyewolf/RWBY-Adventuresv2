import { backend } from "@/plugins/axios";

async function loggedIn() {
    let response = await backend.get("/auth/status");
    let data = await response.data
    return data;
}

export { loggedIn };