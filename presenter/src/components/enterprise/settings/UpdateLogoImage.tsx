/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import { Enterprise } from "../../../lib/types";
import UploadFileIcon from "../../icons/UploadFileIcon";
import { useRouter } from "@tanstack/react-router";
import { Folder, uploadFile } from "../../../lib/upload_files";

interface Props {
    enterprise: Enterprise
}
const UpdateLogoImage: React.FC<Props> = ({ enterprise }: Props) => {
    const [file, setFile] = useState<File | null>(null)
    const router = useRouter()
    function handleClick(e: React.MouseEvent<HTMLDivElement>) {
        e.preventDefault();
        const input = document.getElementById("file-input");
        if (input) {
            input.click();
        }
    }
    useEffect(() => {
        const wrap = async (file : File) => {
            const ok = await uploadFile(file, Folder.Enterprise,enterprise.name)
            if (ok) {
                router.invalidate()
            }
        }
        if (file) {
            wrap(file)
        }
    }, [enterprise.name, file, router])
    useEffect(() => {
        const input = document.getElementById("file-input")
        if (input) {
            input.addEventListener("change", function (e: any) {
                if (e.target.files) {
                    setFile(e.target.files[0]);
                }
            })
        }
        return () => {
            if (input) {
                input.removeEventListener("change", function (e: any) {
                    if (e.target.files) {
                        setFile(e.target.files[0]);
                    }
                })
            }
        }
    }, [])
    return (
        <div className="flex flex-col">
            <div className="flex flex-col items-center w-36 self-center">
                <img className="h-24 w-24 rounded-full my-4" src={enterprise.logoIMG} />
                <div
                    onClick={handleClick}
                    className="inline-flex cursor-pointer border p-2 rounded-md">
                    <button type="button" className="text-gray-700 mr-2">Logo Image</button>
                    <UploadFileIcon />
                </div>
                <input id="file-input" type="file" className="hidden" />
                <p className="text-xs self-end mr-2 mt-1 cursor-default">Max size 4mb.</p>
            </div>
        </div>
    )
}

export default UpdateLogoImage;