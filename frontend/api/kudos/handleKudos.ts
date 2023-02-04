import { KudoPayload } from "../models/kudos"
import { getTokenFromCookies } from "../auth/tokens"


export async function handleSendKudos(payload: KudoPayload) {
    const res = await fetch(`http://127.0.0.1:8000/kudos`, {
        method: "POST",
        headers: new Headers({
            Authorization: `${getTokenFromCookies()}`
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
            Authorization: `${getTokenFromCookies()}`
        })
    })
    return res
}