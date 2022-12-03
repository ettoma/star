import User from "../api/models/user"
type RegisterPayload = {
    name: string,
    username: string,
    password: string,
}

async function handleRegister(payload: RegisterPayload) {
    const res = await fetch("http://127.0.0.1:8000/users", {
        method: "POST",
        body: JSON.stringify({
            name: payload.name,
            username: payload.username,
            password: payload.password
        })
    })

    return res



}

export default handleRegister