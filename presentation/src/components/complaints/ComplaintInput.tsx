'use client'

import { useState } from "react"

interface Props {
    sendCallback?: (body: string) => void
}
const ComplaintInput: React.FC<Props> = ({ sendCallback = (b: string) => { } }) => {
    const handleClick = () => {
        const chatInput = document.getElementById("chat") as HTMLTextAreaElement;
        const message = chatInput.value.trim();
        if (message === "") return;
        chatInput.value = "";
        sendCallback(message);
    }
    const handlePress = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
        if (e.key === "Enter" && !e.shiftKey) {
            const submitBtn = document.getElementById("submit-btn");
            submitBtn?.animate([
                { backgroundColor: "#dbeafe" },
                { transform: "scale(1)" },
                { transform: "scale(1.1)" },
                { transform: "scale(1)" },
                { backgroundColor: "inherit" },
            ], {
                duration: 200,
                iterations: 1
            });
            handleClick();
        }
    }

    return (
        <form className="relative" >
            <label htmlFor="chat" className="sr-only">Your message</label>
            <div className="flex flex-col md:flex-row items-center  md:px-3 md:py-2 rounded-lg bg-gray-100">
                <div className="flex self-start my-auto ">
                    {/* <button type="button" className="inline-flex justify-center p-2 text-gray-500 rounded-lg cursor-pointer hover:text-gray-900 hover:bg-gray-100">
                <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 18">
                    <path fill="currentColor" d="M13 5.5a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0ZM7.565 7.423 4.5 14h11.518l-2.516-3.71L11 13 7.565 7.423Z" />
                    <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M18 1H2a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Z" />
                    <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M13 5.5a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0ZM7.565 7.423 4.5 14h11.518l-2.516-3.71L11 13 7.565 7.423Z" />
                </svg>
                <span className="sr-only">Upload image</span>
            </button> */}
                    <button
                        onMouseUp={() => { }}
                        type="button" className="p-2 text-gray-700 rounded-lg cursor-pointer hover:text-gray-900 hover:bg-gray-100">
                        <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                            <path
                                stroke="currentColor"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth="2"
                                d="M13.408 7.5h.01m-6.876 0h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0ZM4.6 11a5.5 5.5 0 0 0 10.81 0H4.6Z" />
                        </svg>
                        <span className="sr-only">Add emoji</span>
                    </button>
                    {/* {showEmoji && <div
                    ref={emojiRef}
                    className="absolute -top-[251px] left-0 z-50">
                    <EmojiPicker
                        onEmojiClick={(emojiData, _: MouseEvent) => { setInput(input + emojiData.emoji); return; }}
                        theme={Theme.LIGHT}
                        emojiStyle={EmojiStyle.GOOGLE}
                        searchDisabled={true}
                        skinTonesDisabled={true}
                        previewConfig={{ showPreview: false } as PreviewConfig}
                        width={270}
                        height={250} />
                </div>} */}
                </div>
                <div className="flex w-full">
                    <textarea
                        id="chat"
                        onKeyUpCapture={handlePress}
                        rows={2}
                        maxLength={120}
                        className="block md:mx-4 p-2.5 w-full text-sm md:text-xl text-gray-900 bg-white rounded-lg border border-gray-300
     focus:ring-blue-500 focus:border-blue-500 resize-none" placeholder="Your message...">
                    </textarea>
                    <button onMouseUp={() => { handleClick() }}
                        id="submit-btn"
                        type="button"
                        className="my-auto inline-flex justify-center p-2 text-blue-600 rounded-full cursor-pointer hover:bg-blue-100">
                        <svg className="w-5 h-5 rotate-90 rtl:-rotate-90" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 18 20">
                            <path d="m17.914 18.594-8-18a1 1 0 0 0-1.828 0l-8 18a1 1 0 0 0 1.157 1.376L8 18.281V9a1 1 0 0 1 2 0v9.281l6.758 1.689a1 1 0 0 0 1.156-1.376Z" />
                        </svg>
                        <span className="sr-only">Send message</span>
                    </button>
                </div>
            </div>
        </form>
    )
}
export default ComplaintInput;