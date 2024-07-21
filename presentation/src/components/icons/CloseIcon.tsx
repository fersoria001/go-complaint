import IconsProps from "./IconsProps";

//dont put the onClick on it, it's used inside some RSC
const CloseIcon: React.FC<IconsProps> = ({ height = 0, width = 0, fill = "#5f6368", className = "" }: IconsProps) => {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            height={height}
            viewBox="0 -960 960 960"
            width={width}
            className={className}
            fill={fill}>
            <path d="m256-200-56-56 224-224-224-224 56-56 224 224 224-224 56 56-224 224 224 224-56 56-224-224-224 224Z" />
        </svg>
    )
}
export default CloseIcon;