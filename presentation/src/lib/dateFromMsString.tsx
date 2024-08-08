export function dateFromMsString(msString: string): Date {
    return new Date(parseInt(msString))
}