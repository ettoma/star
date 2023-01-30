const token = document.cookie.split("=")[1]

export async function handleKudos(recipient: string, message: string, sender: string) {
    console.log(recipient, message, sender)
}

export async function handleGetKudosPerUser(recipient: string) {
    const res = await fetch(`http://localhost:8000/kudos/${recipient}`, {
        method: 'GET',
        headers: new Headers({
            Authorization: `Bearer ${token}`
        })
    })
    return res
}