/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/no-unused-vars */
import { useEffect, useState } from "react";
import useWindowDimensions from "../../lib/hooks/useWindowsDimensions";
import { FeedbackErrorType } from "../feedback/feedback";
interface Props {
    callback?: (...args: any[]) => void;
    colorSquare: ColorSquare;
    errors: FeedbackErrorType;
    end?:boolean;
}
export type ColorSquare = {
    selected: boolean;
    picked: boolean;
    fixed: boolean;
    color: string;
}
function SquareFillIcon({end, errors, colorSquare, callback = (..._: any[]) => { } }: Props) {
    const { width } = useWindowDimensions();
    const [size, setSize] = useState<number>(24);
    useEffect(() => {
        if (width < 768) {
            setSize(24);
        } else {
            setSize(32);
        }
    }, [colorSquare.color, errors.square, width]);
    let iconClass = "";
    if (end) {
        iconClass = `translate-y-1 scale-110   cursor-default `
    } else if (colorSquare.fixed) {
        iconClass = colorSquare.selected ? `translate-y-1 scale-110   cursor-default ` : colorSquare.picked ?
            `translate-y-1 scale-110   cursor-default` : `cursor-default`;
    } else {
        iconClass = colorSquare.selected ? `translate-y-1 scale-110   cursor-pointer` : colorSquare.picked ?
            `transition ease-out hover:translate-y-1 hover:scale-110 hover: cursor-pointer` :
            `transition ease-out hover:translate-y-1 hover:scale-110 hover: cursor-pointer`;
    }
    const wrap = () => {
        if (end) return
        callback()
    }
    return (
        <div>
            <svg
                onClick={wrap}
                fill={colorSquare.color}
                className={iconClass}
                xmlns="http://www.w3.org/2000/svg"
                height={size}
                viewBox="0 -960 960 960"
                width={size}>
                <path d="M200-160q-33 0-56.5-23.5T120-240v-480q0-33 23.5-56.5T200-800h560q33 0 56.5 23.5T840-720v480q0 33-23.5 56.5T760-160H200Z" />
            </svg>
            {errors.square && errors.square[colorSquare.color] &&
                <span
                    className="absolute left-0 text-red-500 whitespace-nowrap">{errors.square[colorSquare.color]}
                </span>
            }
        </div>
    )
}

export default SquareFillIcon;