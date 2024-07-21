import IconsProps from "./IconsProps";

const SearchIcon: React.FC<IconsProps> = ({ onClick = () => { }, height = 0, width = 0, fill = "none", className = "w-4 h-4" }: IconsProps) => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg"
            onClick={onClick}
            height={height}
            viewBox="0 0 20 20"
            width={width}
            className={className}
            fill={fill}>
            <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z" />
        </svg>
    )
}
export default SearchIcon;