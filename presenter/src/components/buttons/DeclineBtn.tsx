/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react"
import DoneIcon from "../icons/DoneIcon"
import LoadingSpinner from "../icons/LoadingSpinner"
import PrimaryButton from "./PrimaryButton"
import ThinButton from "./ThinBtn"
import CancelIcon from "../icons/CancelIcon"

interface Props {
    variant: 'primary' | 'thin',
    text: string,
    callback: (...args: any[]) => Promise<boolean>
    callbackArgs?: any[],
    status?: string; //'pending' | 'rejected' | 'blocked' | 'accepted',
    cleanUp?: () => void
    animate?: boolean
    blocked?: boolean
}

function DeclineBtn({ blocked = false, variant, text, callback, callbackArgs = [], cleanUp = () => { }, status = "pending", animate = true }: Props) {
    const [loading, setLoading] = useState(false)
    const [done, setDone] = useState(false)
    const [cancel, setCancel] = useState(false)
    const color = 'red'
    const [btnState, setBtnState] = useState('normal')
    const handleClickWithAnimation = () => {
        setLoading(true)
        callback(...callbackArgs).then((res) => {
            if (res) {
                setLoading(false)
                setDone(true)
            }
        })
        cleanUp()
    }
    const handleClick = () => {
        if (btnState === 'blocked') return
        if (btnState === 'accepted') return
        if (btnState === 'clicked') return
        callback(...callbackArgs).then(() => {

        })
        cleanUp()
    }

    useEffect(() => {
        if (blocked) {
            setBtnState('blocked')
        }
        switch (status) {
            case 'accepted': {
                setBtnState('blocked')
                break;
            }
            case 'hired': {
                setBtnState('blocked')
                break;
            }
            case 'rejected': {
                setBtnState('clicked')
                setCancel(true)
                break;
            }
            case 'canceled': {
                setBtnState('blocked')
                setCancel(true)
                break;
            }
            case 'fired': {
                setBtnState('blocked')
                break;
            }
            case 'leaved': {
                setBtnState('blocked')
                break;
            }
        }
    }, [status, blocked])

    return (
        <div className="flex">
            {
                loading ? <span id="loading" className={`mx-2 ${variant === 'primary' ? 'self-center' : 'self-start'} `}>
                    <LoadingSpinner />
                </span> : done ?
                    <span id="done" className={`mx-2 ${variant === 'primary' ? 'self-center' : 'self-start'} `}>
                        <DoneIcon fill="#3b82f6" />
                    </span> : null
            }

            <div className="">
                {variant === 'primary' ?
                    <span onMouseUp={animate ? handleClickWithAnimation : handleClick} className={btnState === 'blocked' || btnState === 'clicked' ? 'cursor-none pointer-events-none' : ''}>
                        <PrimaryButton text={text} variant={color} state={btnState} />
                    </span> :
                    <span onMouseUp={animate ? handleClickWithAnimation : handleClick} className={btnState === 'blocked' || btnState === 'clicked' ? 'cursor-none pointer-events-none' : ''}>
                        <ThinButton text={text} variant={color} state={btnState} />
                    </span>
                }
            </div>
            {
                cancel &&
                <span id="done" className={`mx-2 ${variant === 'primary' ? 'self-center' : 'self-start'} `}>
                    <CancelIcon fill="#EF4444" />
                </span>
            }
        </div>
    )
}

export default DeclineBtn