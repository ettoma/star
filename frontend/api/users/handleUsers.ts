import { RegisterPayload, LoginPayload } from "../models/user"


export async function handleRegister(payload: RegisterPayload) {
    const res = await fetch("http://127.0.0.1:8000/register", {
        method: "POST",
        body: JSON.stringify({
            name: payload.name,
            username: payload.username,
            password: payload.password
        })
    })

    return res
}

export async function handleLogin(payload: LoginPayload) {
    const res = await fetch("http://127.0.0.1:8000/login", {
        method: "POST",
        body: JSON.stringify({
            username: payload.username,
            password: payload.password
        })
    })

    return res
}

export async function handleGetUsers() {
    const res = await fetch("http://127.0.0.1:8000/users", {
        method: "GET"
    })

    return res
}

export async function handleDeleteUser(id: number) {
    const res = await fetch("http://127.0.0.1:8000/users", {
        method: "DELETE",
        body: JSON.stringify({
            id: id
        })
    })

    return res
}
