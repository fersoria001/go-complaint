interface props {
    text: string;
    variant?: string;
}
function ThinButton({ text, variant }: props) {
    let style = `text-white bg-gradient-to-r from-cyan-400 via-cyan-500
    to-cyan-600 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-cyan-300
     font-medium rounded-lg 
     text-sm px-5  text-center me-2 mb-2`
    switch (variant) {
        case "red":
            style = `text-white bg-gradient-to-r from-red-400 via-red-500
            to-red-600 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-red-300
             font-medium rounded-lg 
             text-sm px-5  text-center me-2 mb-2`
            break;
        case "gray":
            style = `text-white bg-gradient-to-r from-gray-400 via-gray-500
        to-gray-600 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-gray-300
         font-medium rounded-lg 
         text-sm px-5  text-center me-2 mb-2`
            break;
        default:
            break;
    }
    return (
        <button type="button" className={style}>{text}</button>
    )
}

export default ThinButton