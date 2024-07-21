interface props {
    height: number;
    width: number;
    fill: string;
    className: string;
    onClick: () => void 
}
export default interface IconsProps extends Partial<props> { }