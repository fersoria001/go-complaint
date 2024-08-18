'use client'

import IconsProps from "./IconsProps";

/**
 * client component
 */
const MinimizeIcon: React.FC<IconsProps> = ({
    onClick,
    height = 32,
    width = 32,
    fill = "#5f6368",
    className = "",
}) => {
    return (
        <svg
            onClick={onClick}
            xmlns="http://www.w3.org/2000/svg"
            height={`${height}px`}
            className={className}
            viewBox="0 -960 960 960"
            width={`${width}px`}
            fill={fill}>
            <path d="M240-120v-80h480v80H240Z" />
        </svg>
    )
}
export default MinimizeIcon;