import { ColorSquare, FeedbackErrorType } from "./feedback";



interface Props {
    isDone: boolean;
    colorSquare: ColorSquare;
    callback: (c: string) => void;
    errors: FeedbackErrorType
}

function ColorSelectorItem({ isDone, colorSquare, callback }: Props) {
    
    let iconClass = "";
    if (isDone) {
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
        callback(colorSquare.color)
    }
    return (
        <div>
            <svg
                onClick={wrap}
                fill={colorSquare.color}
                className={iconClass}
                xmlns="http://www.w3.org/2000/svg"
                height={32}
                viewBox="0 -960 960 960"
                width={32}>
                <path d="M200-160q-33 0-56.5-23.5T120-240v-480q0-33 23.5-56.5T200-800h560q33 0 56.5 23.5T840-720v480q0 33-23.5 56.5T760-160H200Z" />
            </svg>
        </div>
    )
}

export default ColorSelectorItem;