import { useEffect, useState } from "react";
import useWindowDimensions from "../../lib/hooks/useWindowsDimensions";

interface Props {
    callback: (description: string) => void;
}

function Description({ callback }: Props) {
    const { width } = useWindowDimensions();
    const [rows, setRows] = useState<number>(4);
    useEffect(() => {
        if (width >= 768) {
            setRows(7);
        } else {
            setRows(4);
        }
    }, [width]);
    return (
        <div>
            <label
                htmlFor="description"
                className="block mb-2 text-sm md:text-xl font-medium text-gray-900"
            >Description
            </label>
            <textarea
                onChange={(e) => callback(e.target.value)}
                id="description"
                rows={rows}
                minLength={3}
                maxLength={120}
                className="block p-2.5 w-full text-sm md:text-xl
               text-gray-900 bg-gray-50 rounded-lg border border-gray-300
                focus:ring-blue-500 focus:border-blue-500"
                placeholder="Shortly describe the problem here...">
            </textarea>
        </div>
    )
}

export default Description