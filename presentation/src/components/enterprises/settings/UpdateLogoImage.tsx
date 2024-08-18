'use client'
import UploadFileIcon from "@/components/icons/UploadFileIcon";
import { Enterprise } from "@/gql/graphql";
import getGraphQLClient from "@/graphql/graphQLClient";
import changeEnterpriseLogoImgMutation from "@/graphql/mutations/changeEnterpriseLogoImgMutation";
import { useQueryClient, useMutation } from "@tanstack/react-query";
import { useEffect } from "react";
import Image from "next/image";
interface Props {
    enterprise: Enterprise
}
const UpdateLogoImage: React.FC<Props> = ({ enterprise }: Props) => {
    const queryClient = useQueryClient()
    const uploadMutation = useMutation({
        mutationFn: async (file: File) => await getGraphQLClient().request(changeEnterpriseLogoImgMutation, {
            enterpriseId: enterprise.id,
            file: file
        }),
        onSuccess: () => queryClient.refetchQueries({ queryKey: ["enterpriseByName", enterprise.name] })
    })

    function handleClick(e: React.MouseEvent<HTMLDivElement>) {
        e.preventDefault();
        const input = document.getElementById("file-input");
        if (input) {
            input.click();
        }
    }
    useEffect(() => {
        const input = document.getElementById("file-input")
        if (input) {
            input.addEventListener("change", function (e: any) {
                if (e.target.files) {
                    uploadMutation.mutate(e.target.files[0])
                }
            })
        }
        return () => {
            if (input) {
                input.removeEventListener("change", function (e: any) {
                    if (e.target.files) {
                        uploadMutation.mutate(e.target.files[0])
                    }
                })
            }
        }
    }, [queryClient, uploadMutation])
    return (
        <div className="flex flex-col">
            <div className="flex flex-col items-center w-36 self-center">
                <div className='relative h-24 w-24 my-4'>
                    <Image
                        src={enterprise.logoImg}
                        className="rounded-full"
                        sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                        fill
                        alt="user photo" />
                </div>
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