'use client'

import { useEffect, useState } from "react";
import ArrowUpIcon from "../icons/ArrowUpIcon";
import clsx from "clsx";

const ScrollUpButton: React.FC = () => {
    const [show, setShow] = useState<boolean>(false)
    const [scrollTop, setScrollTop] = useState(0);
    useEffect(() => {
        const onScroll = (e: any) => {
            setScrollTop(e.target.documentElement.scrollTop);
        };
        window.addEventListener("scroll", onScroll);
        if (scrollTop > 700) {
            setShow(true)
        } else {
            setShow(false)
        }
        return () => window.removeEventListener("scroll", onScroll);
    }, [scrollTop]);
    const scrollToTop = () => {
        const topElement = document.getElementById("scroll-top")
        topElement?.scrollIntoView({
            behavior: "smooth",
            block: "end"
        })
    }
    return (
        <button
            type="button"
            onClick={scrollToTop}
            className={clsx('fixed bg-blue-500 rounded-full h-12 w-12 z-30 bottom-44 lg:bottom-28 right-4 lg:right-10 cursor-pointer border 2xl:h-24 2xl:w-24 ',
                {
                    'hidden': !show,
                    '': show
                }
            )}>
            <ArrowUpIcon
                className="fixed lg:bottom-28 lg:mb-2 bottom-44 mb-2 right-6 lg:right-12 z-50 2xl:w-20 2xl:h-20"
                fill="#ffffff" />
        </button>
    )
}
export default ScrollUpButton;