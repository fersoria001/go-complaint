/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react"
import DoneIcon from "../icons/DoneIcon"
import LoadingSpinner from "../icons/LoadingSpinner"
import PrimaryButton from "./PrimaryButton"
import ThinButton from "./ThinBtn"

interface Props {
    variant: 'primary' | 'thin',
    text: string,
    callback: (...args: any[]) => Promise<boolean>
    callbackArgs?: any[],
    status?: string; //'pending' | 'rejected' | 'blocked' | 'accepted',
    cleanUp?: () => void
    reset?: boolean
    blocked?: boolean
}

const AcceptBtn: React.FC<Props> = ({ blocked = false, reset = false, variant, text, callback, callbackArgs = [], cleanUp = () => { }, status = "pending", }: Props) => {
    const [loading, setLoading] = useState(false)
    const [done, setDone] = useState(false)
    const [btnState, setBtnState] = useState('normal')
    const color = 'blue'
    useEffect(() => {
        switch (status) {
            case 'accepted': {
                setBtnState('blocked')
                setDone(true)
                break;
            }
            case 'hired': {
                setBtnState('clicked')
                setDone(true)
                break;
            }
            case 'rejected': {
                setBtnState('blocked')
                break;
            }
            case 'canceled': {
                setBtnState('blocked')
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
    }, [status])
    useEffect(() => {
        if (reset) {
            setBtnState('normal')
            setLoading(false)
            setDone(false)
        }
        if (blocked) {
            setBtnState('blocked')
        }
    }, [reset, blocked])
    const handleClick = async () => {
        if (btnState === 'blocked') return
        if (btnState === 'accepted') return
        if (btnState === 'clicked') return
        setLoading(true)
        const ok = await callback(...callbackArgs)
        if (ok) {
            setLoading(false)
            setBtnState('clicked')
            setDone(true)
        } else {
            setBtnState('blocked')
            setLoading(false)
        }
        cleanUp()
    }


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
            <div>
                {variant === 'primary' ?
                    <span onMouseUp={handleClick} className={btnState === 'blocked' || btnState === 'clicked' ? 'cursor-none pointer-events-none' : ''}>
                        <PrimaryButton text={text} variant={color} state={btnState} />
                    </span> :
                    <div onMouseUp={handleClick} className={btnState === 'blocked' || btnState === 'clicked' ? 'cursor-none pointer-events-none' : ''}>
                        <ThinButton text={text} variant={color} state={btnState} />
                    </div>
                }
            </div>
        </div>
    )
}

export default AcceptBtn