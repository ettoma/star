

export function unixToDate(timeUnix: number): string {
    const timeString = new Date(timeUnix)

    return timeString.toLocaleDateString()
}