'use client'
import { useEffect, useState } from "react";
import { LineChart, Line, CartesianGrid, XAxis, YAxis, Label, ResponsiveContainer } from 'recharts';
import SelectIcon from "../icons/SelectIcon";
import { TimeDataBias, SelectorsData, TimeData, ChartBuilder } from "./ChartBuilder";
import clsx from "clsx";

interface Props {
    chartLabel: string;
    yLabel: string;
    data: TimeData[]
}



const TimeDataChart: React.FC<Props> = ({ chartLabel, yLabel, data: timeData }: Props) => {
    const [bias, setBias] = useState<TimeDataBias>(TimeDataBias.Month)
    const [selectorsData, setSelectorsData] = useState<SelectorsData | null>(null)
    const [data, setData] = useState<TimeData[]>([])
    const [selectorValue, setSelectorValue] = useState<number | undefined>(undefined)
    useEffect(() => {
        const chartBuilder = new ChartBuilder(bias, {
            day: selectorValue,
            month: selectorValue
        })
        setData(chartBuilder.filter(timeData))
        setSelectorsData(chartBuilder.selectorsData)
    }, [bias, timeData, selectorValue])
    useEffect(() => {
        setSelectorValue(selectorsData?.defaultValue)
    }, [selectorsData?.selectorLabel])
    return (
        <div className="w-full shrink-0 flex flex-col items-center">
            <h1 className="text-center text-gray-700 font-bold p-1.5">{chartLabel}</h1>
            <ResponsiveContainer width={"100%"} height={300}>
                <LineChart data={data} margin={{ top: 5, right: 15, left: 15, bottom: 15 }}>
                    <Line type="monotone" dataKey="qtty" stroke="#3b82f6" />
                    <CartesianGrid stroke="#ccc" />
                    <XAxis dataKey="occurredOn">
                        <Label value={selectorsData?.xLabel} offset={0} position="bottom" />
                    </XAxis>
                    <YAxis label={{ value: yLabel, angle: -90, position: 'insideLeft' }} />
                </LineChart>
            </ResponsiveContainer>
            <div className="flex flex-col md:flex-row  items-center self-center">
                <div className="flex gap-2 shrink-0 pl-2">
                    <button
                        className={clsx("bg-blue-500 px-7 py-3 font-bold text-white rounded-md", {
                            'bg-blue-600': bias === TimeDataBias.Day,
                            'hover:bg-blue-600': bias !== TimeDataBias.Day,
                        })}
                        onClick={() => setBias(TimeDataBias.Day)}>
                        By Day
                    </button>
                    <button
                        className={clsx("bg-blue-500 px-7 py-3 font-bold text-white rounded-md", {
                            'bg-blue-600': bias === TimeDataBias.Month,
                            'hover:bg-blue-600': bias !== TimeDataBias.Month,
                        })}
                        onClick={() => setBias(TimeDataBias.Month)}>
                        By month
                    </button>
                </div>
                {
                    selectorsData &&
                    <div className="p-3 shrink-0">
                        <label
                            className="block text-gray-700 text-sm lg:text-md font-bold"
                            htmlFor={"selector"}>Select a {" "}{selectorsData.selectorLabel}</label>
                        <div className="relative">
                            <select
                                id="selector"
                                className="block appearance-none w-full bg-gray-200 border
                        text-md lg:text-lg
                     border-gray-200 text-gray-700  px-4 pr-8 rounded leading-tight
                      focus:outline-none focus:bg-white focus:border-gray-500"
                                name={selectorsData.selectorLabel}
                                value={selectorValue}
                                onChange={(e) => setSelectorValue(parseInt(e.currentTarget.value))}>
                                {
                                    selectorsData.selectors.map((v) => <option key={v}>{v}</option>)
                                }
                            </select>
                            <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                                <SelectIcon />
                            </div>
                        </div>
                    </div>
                }
            </div>
        </div>
    )
}

export default TimeDataChart