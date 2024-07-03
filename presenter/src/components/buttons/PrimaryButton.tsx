
interface props {
    text: string;
    variant?: string; // 'red' | 'blue';
    state?:  string; // 'normal' | 'clicked' | 'hovered';
}
function PrimaryButton({ text, variant = 'blue', state = 'normal' }: props) {
    let className1
    let className2
    let className3 = ''
    switch (variant) {
        case 'red': {
            className1 = `relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden
             text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-red-200 via-red-300 to-yellow-200 
            group-hover:from-red-200 group-hover:via-red-300 group-hover:to-yellow-200 focus:ring-4 focus:outline-none focus:ring-red-100`
            className2 = "text-sm md:text-xl relative px-6 py-3 transition-all ease-in duration-75 bg-white  rounded-md group-hover:bg-opacity-0"
            if (state === 'clicked') {
                className1 = `relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden
                 text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-red-200 via-red-300 to-yellow-200
                 ring-4 outline-none ring-red-100`
                className2 = "relative px-6 py-3 bg-white rounded-md bg-opacity-0"
                className3 = `cursor-not-allowed pointer-events-none`
            } else if (state === 'blocked') {
                className1 = `relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden
                 text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-red-200 via-red-300 to-yellow-200
                 ring-4 outline-none ring-red-100 opacity-50`
                className2 = "relative px-6 py-3 bg-white rounded-md bg-opacity-0"
                className3 = `cursor-not-allowed pointer-events-none`
            }
            break;
        }
        case 'blue': {
            className1 = `relative inline-flex items-center justify-center p-0.5 mb-2 me-2 
            overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br
             from-cyan-500 to-blue-500 group-hover:from-cyan-500 group-hover:to-blue-500
              hover:text-white  focus:ring-4 focus:outline-none focus:ring-cyan-200 `
            className2 = `text-sm md:text-xl relative px-6 py-3
                 transition-all ease-in duration-75 
                 bg-white rounded-md group-hover:bg-opacity-0`
            if (state === 'clicked') {
                className1 = `relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden
                 text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-cyan-500 to-blue-500
                 ring-4 outline-none ring-cyan-200`
                className2 = "relative px-6 py-3 bg-white rounded-md bg-opacity-0"
                className3 = `cursor-not-allowed pointer-events-none`
            } else if (state === 'blocked') {
                className1 = `relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden
                 text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-cyan-500 to-blue-500
                 ring-4 outline-none ring-cyan-200 opacity-50`
                className2 = "relative px-6 py-3 bg-white rounded-md bg-opacity-0"
                className3 = `cursor-not-allowed pointer-events-none`
            }
        }
    }
    return (
        <div className={className3}>
            <button  type="button" className={className1}>
                <span className={className2}>
                    <p>{text}</p>
                </span>
            </button>
        </div>
    )
}

export default PrimaryButton