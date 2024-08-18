import IconsProps from "./IconsProps";

const CheerIcon: React.FC<IconsProps> = ({ height = 0, width = 0, fill = "#5f6368", className = "" }: IconsProps) =>{

    return (
        <svg

            xmlns="http://www.w3.org/2000/svg"
            height={`${height}px`}
            viewBox="0 -960 960 960"
            width={`${width}px`}
            className={className}
            fill={fill}>
            <path d="m312-751-40-120 56-18 40 119-56 19Zm138-49v-120h60v120h-60Zm198 49-56-19 40-119 56 19-40 119ZM86-40l-12-79 211-32q11-2 19.5-9.5T317-179l34-106q5-14 0-27t-18-20l-33 104-76-24 88-278q2-6 2-13t-2-13L178-304q-16 29-44.5 46.5T72-240H40v-80h32q11 0 20.5-5.5T107-341l177-334 50 28q37 21 52.5 60.5T389-506l-31 98q44 17 63.5 60t5.5 88l-34 106q-11 32-36.5 54.5T297-72L86-40Zm788 0L663-72q-34-5-59.5-27.5T567-154l-34-106q-14-45 5.5-88t63.5-60l-31-98q-13-41 2.5-80.5T626-647l50-28 177 334q5 10 14.5 15.5T888-320h32v80h-32q-33 0-61.5-17.5T782-304L648-556q-2 6-2 13t2 13l88 278-76 24-33-104q-13 7-18 20t0 27l34 106q4 11 12.5 18.5T675-151l211 32-12 79ZM224-252Zm512 0Zm-76 24-58-180 58 180ZM358-408l-58 180 58-180Z" />
        </svg>
    )
}

export default CheerIcon;