import { useEffect, useState } from "react";
import useWindowDimensions from "../../lib/hooks/useWindowsDimensions";
import IconsProps from "./IconsProps";
interface StarIconProps extends IconsProps {
    index: number;
    rating: number;
    hover: number;
}
function StarIcon({ index, rating, hover }: StarIconProps) {
    const [widthHeight, setWidthHeight] = useState<number>(0)
    const { width } = useWindowDimensions();
    useEffect(() => {
        if (width < 640) {
            setWidthHeight(24)
        } else {
            setWidthHeight(32)
        }
    }, [width])
    return <svg
        xmlns="http://www.w3.org/2000/svg"
        height={widthHeight}
        viewBox="0 -960 960 960"
        width={widthHeight}
        className={index <= (hover || rating) ? "fill-yellow-500" : "fill-gray-200"}
    //fill={fill}
    >
        <path d="m233-120 65-281L80-590l288-25 112-265 112 265 288 25-218 189 65 281-247-149-247 149Z" />
    </svg>
}

export default StarIcon;