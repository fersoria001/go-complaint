'use client'
import UploadFileIcon from "@/components/icons/UploadFileIcon";
import { Enterprise } from "@/gql/graphql";
import getGraphQLClient from "@/graphql/graphQLClient";
import changeEnterpriseBannerImgMutation from "@/graphql/mutations/changeEnterpriseBannerImgMutation";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useEffect } from "react";
import Image from "next/image";
interface Props {
    enterprise: Enterprise
}
const UpdateBannerImage: React.FC<Props> = ({ enterprise }: Props) => {
    const queryClient = useQueryClient()
    const uploadMutation = useMutation({
        mutationFn: async (file: File) => await getGraphQLClient().request(changeEnterpriseBannerImgMutation, {
            enterpriseId: enterprise.id,
            file: file
        }),
        onSuccess: () => queryClient.refetchQueries({ queryKey: ["enterpriseByName", enterprise.name] })
    })

    function handleClick(e: React.MouseEvent<HTMLDivElement>) {
        e.preventDefault();
        const input = document.getElementById("banner-input");
        if (input) {
            input.click();
        }
    }
    useEffect(() => {
        const input = document.getElementById("banner-input")
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
            <div className="flex flex-col items-center  self-center">
                <div className="relative md:w-[1200px] md:h-[400px] h-52 w-full rounded-md mb-4">
                    <Image
                        src={enterprise.bannerImg}
                        sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                        fill
                        alt="user photo" />
                </div>
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