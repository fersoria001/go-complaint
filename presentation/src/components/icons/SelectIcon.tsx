import React from "react";
import IconsProps from "./IconsProps";

const SelectIcon: React.FC = ({ height = 0, width = 0, fill = "current", className = "h-4 w-4" }: IconsProps) => {
    return (
        <svg
            height={height}
            width={width}
            fill={fill}
            className={className}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20">
            <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
        </svg>
    )

}

export default SelectIcon;