import { KudoPayload } from "../models/kudos"

function getToken(): string {

    const token = document.cookie.split("=")[1]

    if (token == null) {
        return ""
    } else {
        return "Bearer " + token
    }
}


export async function handleSendKudos(payload: KudoPayload) {
    const res = await fetch(`http://127.0.0.1:8000/kudos`, {
        method: "POST",
        headers: new Headers({
            Authorization: `${getToken()}`
        }),
        body: JSON.stringify({
            sender: payload.sender,
            receiver: payload.receiver,
            content: payload.content
        })
    })
    return res
}

export async function handleGetKudosPerUser(recipient: string) {
    const res = await fetch(`http://127.0.0.1:8000/kudos/${recipient}`, {
        method: 'GET',
        headers: new Headers({
            Authorization: `${getToken()}`
        })
    })
    return res
}