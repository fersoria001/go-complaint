/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import { Enterprise } from "../../../lib/types";
import UploadFileIcon from "../../icons/UploadFileIcon";
import { useRouter } from "@tanstack/react-router";
import { Folder, uploadFile } from "../../../lib/upload_files";

interface Props {
    enterprise: Enterprise
}
const UpdateBannerImage: React.FC<Props> = ({ enterprise }: Props) => {
    const [file, setFile] = useState<File | null>(null)
    const router = useRouter()
    function handleClick(e: React.MouseEvent<HTMLDivElement>) {
        e.preventDefault();
        const input = document.getElementById("banner-input");
        if (input) {
            input.click();
        }
    }
    useEffect(() => {
        const wrap = async (file : File) => {
            const ok = await uploadFile(file, Folder.Banner,enterprise.name)
            if (ok) {
                router.invalidate()
            }
        }
        if (file) {
            wrap(file)
        }
    }, [enterprise.name, file, router])
    useEffect(() => {
        const input = document.getElementById("banner-input")
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
            <div className="flex flex-col items-center  self-center">
                <img className="bg-cover md:row-start-3 md:col-span-3 md:w-[1200px] md:h-[400px] h-52 w-full rounded-md mb-4"  src={enterprise.bannerIMG} />
                <div
                    onClick={handleClick}
                    className="inline-flex cursor-pointer border p-2 rounded-md">
                    <button type="button" className="text-gray-700 mr-2">Banner Image</button>
                    <UploadFileIcon />
                </div>
                <input id="banner-input" type="file" className="hidden" />
                <p className="text-xs ml-16 mt-1 cursor-default ">Max size 4mb.</p>
            </div>
        </div>
    )
}

export default UpdateBannerImage;