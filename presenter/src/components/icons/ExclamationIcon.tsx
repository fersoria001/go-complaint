import IconsProps from "./IconsProps";

function ExclamationIcon({ fill }: IconsProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            height="24px"
            viewBox="0 -960 960 960"
            width="24px"
            fill={fill}>
            <path d="M440-400v-360h80v360h-80Zm0 200v-80h80v80h-80Z" />
        </svg>)
}

export default ExclamationIcon;