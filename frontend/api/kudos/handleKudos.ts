export async function handleKudos(recipient: string, message: string) {
    console.log(recipient, message)
}

export async function handleGetKudosPerUser(recipient: string) {
    const res = await fetch('http://localhost:8000/kudos/users', {
        method: 'POST',
        body: JSON.stringify({
            receiver: recipient
        })
    })
    return res
}