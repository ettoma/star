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

// const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzQ4Mzg5OTAsInVzZXIiOiJ0ZXN0MiJ9.-2IsCzADysf07Rn-vAfK_1XVYpLpsBYltaFp3UyAwkQ"
const token = document.cookie.split("=")[1]