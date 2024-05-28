import KeyboardReturnIcon from "../icons/KeyboardReturnIcon";

function ChatButton() {
    return (
        <div>
            <button
                className="relative inline-flex items-center justify-center p-0.5
             mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900
              rounded-lg group bg-gradient-to-br
               from-cyan-500 to-blue-500 group-hover:from-cyan-500
                group-hover:to-blue-500 hover:text-white 
                 focus:ring-4 focus:outline-none focus:ring-cyan-200 ">
                <span className="relative px-5 py-1 md:py-2.5 md:px-12
                 transition-all ease-in duration-75
                  bg-white rounded-md group-hover:bg-opacity-0">
                    <KeyboardReturnIcon fill={"#4B5563"} />
                </span>
            </button>
        </div>
    );
}

export default ChatButton;