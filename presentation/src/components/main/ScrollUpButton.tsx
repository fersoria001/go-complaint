'use client'

import { useEffect, useState } from "react";
import ArrowUpIcon from "../icons/ArrowUpIcon";
import clsx from "clsx";
import useScroll from "@/lib/hooks/useScroll";

const ScrollUpButton: React.FC = () => {
    const [show, setShow] = useState<boolean>(false)
    const { scrollTop } = useScroll()
    useEffect(() => {
        if (scrollTop > 700) {
            setShow(true)
        } else {
            setShow(false)
        }
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
            className={clsx('fixed bg-blue-500 rounded-full h-12 w-12 z-10 bottom-44 lg:bottom-36 right-4 lg:right-6 cursor-pointer border 2xl:h-24 2xl:w-24 ',
                {
                    'hidden': !show,
                    '': show
                }
            )}>
            <ArrowUpIcon
                className="fixed lg:bottom-36 lg:mb-2 bottom-44 mb-2 right-6 lg:right-8 z-50 2xl:w-20 2xl:h-20"
                fill="#ffffff" />
        </button>
    )
}
export default ScrollUpButton;