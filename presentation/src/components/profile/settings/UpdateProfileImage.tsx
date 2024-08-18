'use client'
import UploadFileIcon from "@/components/icons/UploadFileIcon"
import { UserDescriptor } from "@/gql/graphql"
import getGraphQLClient from "@/graphql/graphQLClient"
import updateProfileImageMutation from "@/graphql/mutations/updateProfileImageMutation"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import Image from "next/image"
import { useEffect } from "react"
interface Props {
    descriptor: UserDescriptor
}
const UpdateProfileImage: React.FC<Props> = ({ descriptor }: Props) => {
    const queryClient = useQueryClient()
    const uploadMutation = useMutation({
        mutationFn: async (file: File) => await getGraphQLClient().request(updateProfileImageMutation, { id: descriptor.id, file: file }),
        onSuccess: () => queryClient.refetchQueries({ queryKey: ['userDescriptor'] })
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
                        src={descriptor.profileImg}
                        className="rounded-full"
                        sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                        fill
                        alt="user photo" />
                </div>
                <div
                    onClick={handleClick}
                    className="inline-flex cursor-pointer border p-2 rounded-md">
                    <button type="button" className="text-gray-700 mr-2">Profile Image</button>
                    <UploadFileIcon />
                </div>
                <input id="file-input" type="file" className="hidden" />
                <p className="text-xs self-end mr-2 mt-1 cursor-default">Max size 4mb.</p>
            </div>
        </div>
    )
}

export default UpdateProfileImage