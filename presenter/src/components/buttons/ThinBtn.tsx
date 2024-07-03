interface props {
    text: string;
    variant?: string;
    state?: string;
}
function ThinBtn({ text, variant = 'blue', state = 'normal' }: props) {
    let className1 = ''
    let className2 = ''
    switch (variant) {
        case "red":
            className1 = `text-gray-900 bg-gradient-to-r from-red-200 via-red-300 to-yellow-200
             hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-red-100 
              font-medium rounded-lg text-sm px-5 py-0.5 text-center me-2 mb-2 
              text-gray-900 bg-gradient-to-r from-red-200 via-red-300 to-yellow-200 hover:bg-gradient-to-bl focus:ring-4 
            focus:outline-none focus:ring-red-100  font-medium rounded-lg text-sm  text-center me-2 mb-2`
            if (state === 'clicked') {
                className1 = `text-gray-900 bg-gradient-to-r from-red-200 via-red-300 to-yellow-200
             hover:bg-gradient-to-bl ring-4 outline-none ring-red-100 
              font-medium rounded-lg text-sm px-5 py-0.5 text-center me-2 mb-2 
              text-gray-900 bg-gradient-to-r from-red-200 via-red-300 to-yellow-200 ring-4 
            outline-none ring-red-100  font-medium rounded-lg text-sm  text-center me-2 mb-2 `
                className2 = `cursor-not-allowed pointer-events-none`
            } else if (state === 'blocked') {
                className1 = `text-gray-900 bg-gradient-to-r from-red-200 via-red-300 to-yellow-200
             hover:bg-gradient-to-bl ring-4 outline-none ring-red-100 
              font-medium rounded-lg text-sm px-5 py-0.5 text-center me-2 mb-2 
              text-gray-900 bg-gradient-to-r from-red-200 via-red-300 to-yellow-200 ring-4 
            outline-none ring-red-100  font-medium rounded-lg text-sm  text-center me-2 mb-2  opacity-50`
                className2 = `cursor-not-allowed pointer-events-none`
            }
            break;
        case "blue":
            className1 = `text-white bg-gradient-to-r from-cyan-500 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 
           focus:outline-none focus:ring-cyan-300 font-medium rounded-lg text-sm px-5 py-0.5 text-center me-2 mb-2`
            if (state === 'clicked') {
                className1 = `text-white bg-gradient-to-r from-cyan-500 to-blue-500 ring-4 
           outline-none ring-cyan-300 font-medium rounded-lg text-sm px-5 py-0.5 text-center me-2 mb-2`
                className2 = `cursor-not-allowed pointer-events-none`
            } else if (state === 'blocked') {
                className1 = `text-white bg-gradient-to-r from-cyan-500 to-blue-500 ring-4 
           outline-none ring-cyan-300 font-medium rounded-lg text-sm px-5 py-0.5 text-center me-2 mb-2  opacity-50 `
                className2 = `cursor-not-allowed pointer-events-none`
            }
            break;
        default:
            break;
    }
    return (
        <div className={className2}>
            <button type="button" className={className1}>{text}</button>
        </div>
    )
}

export default ThinBtn