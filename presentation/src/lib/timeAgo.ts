const timeAgo = (date: string): string => {
    const obj = new Date(parseInt(date))
    const now = new Date()
    const diff = now.getTime() - obj.getTime()
    const seconds = Math.floor(diff / 1000)
    let result = 0
    if (seconds < 3600) {
        result = Math.floor(seconds / 60)
        return `${result}m ago`
    }
    result = Math.floor(seconds / 3600)
    if (result > 24) {
        return `${Math.floor(result / 24)}d ago`
    }
    return `${result}h ago`
}

export default timeAgo;