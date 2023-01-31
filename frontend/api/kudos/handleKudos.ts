function getToken(): string {

    const token = document.cookie.split("=")[1]

    if (token == null) {
        return ""
    } else {
        return "Bearer " + token
    }
}


export async function handleKudos(recipient: string, message: string, sender: string) {
    console.log(recipient, message, sender)
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