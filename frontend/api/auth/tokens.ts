export function getTokenFromCookies(): string {

    const token = document.cookie
        .split('; ')
        .find((row) => row.startsWith('token='))
        ?.split('=')[1];

    if (token == null) {
        return ""
    } else {
        return "Bearer " + token
    }
}

export function getRefresherFromCookies(): string {
    const token = document.cookie
        .split('; ')
        .find((row) => row.startsWith('refresh_token='))
        ?.split('=')[1];

    if (token == null) {
        return ""
    } else {
        return "Bearer " + token
    }
}

export async function refreshToken(username: string) {
    const res = await fetch(`http://127.0.0.1:8000/auth-refresh`, {
        method: "POST",
        headers: new Headers({
            Authorization: `${getRefresherFromCookies()}`
        }),
        body: JSON.stringify({
            username: username
        })
    })
        .then(res => res.json())
        .then(data => document.cookie = "refresh_token=" + data["refreshToken"])

}