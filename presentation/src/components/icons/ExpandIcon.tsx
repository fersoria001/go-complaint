'use client'
import IconsProps from "./IconsProps";

/**
 * client component
 */
function ExpandIcon({ onClick, height = 0, width = 0, fill = "#5f6368", className = "" }: IconsProps) {
    return (
        <svg
            onClick={onClick}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 -960 960 960"
            height={`${height}px`}
            width={`${width}px`}
            className={className}
            fill={fill}>
            <path d="M200-120q-33 0-56.5-23.5T120-200v-560q0-33 23.5-56.5T200-840h560q33 0 56.5 23.5T840-760v560q0 33-23.5 56.5T760-120H200Zm0-80h560v-560H200v560Z" />
        </svg>)
}

export default ExpandIcon;