export type Kudo = {
    sender: string,
    receiver: string,
    content: string,
    id: number,
    timestamp: number
}

export type KudoPayload = {
    sender: string,
    receiver: string,
    content: string,
}