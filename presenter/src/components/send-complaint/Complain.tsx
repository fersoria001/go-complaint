import { useContext, useState, useEffect } from "react";
import PrimaryButton from "../buttons/PrimaryButton";
import Stepper from "./Stepper";
import useWindowDimensions from "../../lib/hooks/useWindowsDimensions";
import { ComplaintContext } from "../../react-context/ComplaintContext";
import Modal from "./Modal";
import { Route } from "../../routes/_profile";
import { ErrorType, Receiver, SendComplaintValidationSchema } from "../../lib/types";
import { FindAuthorByIDQuery, FindAuthorByIDType, Query } from "../../lib/queries";
import { useForm } from "../../lib/hooks/useForm";
import { SendComplaintMutation } from "../../lib/mutations";
import { useRouter } from "@tanstack/react-router";



function Complain() {
    const { descriptor } = Route.useLoaderData();
    const [content, setContent] = useState<string>("")
    const [showModal, setShowModal] = useState<boolean>(false);
    const [errors, setErrors] = useState<ErrorType>({})
    const [formData, setFormData] = useState<FormData | null>(null)
    const { success } = useForm(
        formData,
        SendComplaintValidationSchema,
        SendComplaintMutation,
    )
    const { complaint, setKeyValue } = useContext(ComplaintContext);
    const [receiver, setReceiver] = useState<Receiver>({} as Receiver);
    const { width } = useWindowDimensions();
    const [rows, setRows] = useState<number>(4);
    const router = useRouter();
    useEffect(() => {
        if (success) {
            router.navigate({ to: '/complaint-sent' })
        }
        if (width >= 768) {
            setRows(8);
        } else {
            setRows(6);
        }
        async function getAuthor() {
            const receiver = await Query<Receiver>(
                FindAuthorByIDQuery,
                FindAuthorByIDType,
                [complaint.receiverID],
            )
            setReceiver(receiver)
        }
        getAuthor()
    }, [success, router, width, complaint.receiverID]);
    const handleClick = () => {
        const errors = setKeyValue("content", content)
        if (Object.keys(errors).length > 0) {
            setErrors(errors)
        }
        if (errors.content) {
            return
        }
        return setShowModal(true)
    }
    function handleSubmit() {
        const newFormData = new FormData()
        newFormData.append("content", content)
        newFormData.append("receiverID", complaint.receiverID)
        newFormData.append("receiverFullName", complaint.receiverFullName)
        newFormData.append("receiverProfileIMG", complaint.receiverProfileIMG)
        newFormData.append("title", complaint.title)
        newFormData.append("description", complaint.description)
        newFormData.append("title", complaint.title)
        newFormData.append("authorID", descriptor.email)
        setFormData(newFormData)
    }
    return (
        <div className="flex flex-col relative">
            <div className="w-full mb-2 md:w-2/3 self-center md:mb-10">
                <label
                    htmlFor="complain"
                    className="block mb-2 text-sm font-medium text-gray-900">
                    Complain about it
                </label>
                <textarea
                    id="complaint"
                    rows={rows}
                    minLength={50}
                    maxLength={250}
                    onChange={(e) => setContent(e.target.value)}
                    className="block w-full text-sm md:text-xl p-2.5
               text-gray-900 bg-gray-50 rounded-lg border border-gray-300
                focus:ring-blue-500 focus:border-blue-500"
                    placeholder={`Complain to ${receiver.fullName} about ${complaint.title}`}>
                </textarea>
                {errors.content && <span
                    className="self-center text-red-500 text-xs italic">
                    {errors.content} </span>}
            </div>

            <span
                onClick={handleClick}
                className="self-center md:mb-[42px]">
                <PrimaryButton text="Complain!" />
            </span>
            <div
                className="self-center px-2 md:px-6">
                <Stepper step={3} />
            </div>
            {showModal && (<Modal
                id="confirm-complaint-modal"
                fullName={receiver.fullName}
                show={showModal}
                validatedObject={complaint}
                callbackFn={() => { handleSubmit() }}
                closeFn={() => { setShowModal(false) }}
            />)}

        </div>
    );
}

export default Complain;