export type User = {
    name: string,
    username: string,
    createdAt: string,
    id: number,
}

export type RegisterPayload = {
    name: string,
    username: string,
    password: string,
}

export type LoginPayload = {
    username: string,
    password: string,
    token: string
}
