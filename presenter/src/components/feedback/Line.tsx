import { useEffect, useState } from "react";
import useWindowDimensions from "../../lib/hooks/useWindowsDimensions";
import { LinePoints } from "./feedback";

interface Props {
    stroke: string;
    points: LinePoints;
}
function Line({ stroke, points }: Props) {
    const key = Math.random().toString(36).substring(7);
    const [viewBox, setViewBox] = useState<string>("0 0 1280 768");
    const { height, width } = useWindowDimensions();
    useEffect(() => {
        setViewBox(`0 0 ${width} ${height}`)
    }, [height, width])
    return (
        <svg
            key={key}
            style={{ height: "100%", position: "absolute" }}
            viewBox={viewBox}
            xmlns="http://www.w3.org/2000/svg">
            <line x1={points.x1} x2={points.x2} y1={points.y1} y2={points.y2} stroke={stroke} />
        </svg>
    )
}

export default Line;
