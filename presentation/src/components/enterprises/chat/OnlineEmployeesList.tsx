import CircleFillIcon from "@/components/icons/CircleFillIcon"
import { Employee } from "@/gql/graphql"

interface Props {
    items: Employee[]
    openTab: (e: Employee) => void
}
const OnlineEmployeesList: React.FC<Props> = ({ items, openTab }: Props) => {
    return (
        <ul className="h-[200px] md:h-full px-5 pt-2.5">
            {
                items.map((item) => {
                    return (
                        <li
                            //there's no enterprise owner in this list
                            key={item.id}
                            onMouseUp={() => openTab(item)} //
                            className="flex cursor-pointer items-center">
                            <CircleFillIcon fill={item.user.status === "ONLINE" ? '#93c5fd' : '#fca5a5'} className="w-6 h-6" />
                            <p className="ms-10 text-gray-700 text-sm md:text-md">
                                {item.user.person.firstName} {' '} {item.user.person.lastName}
                            </p>
                            {/* 
                            this will need to be part of the aggregate 
                            {user.msgs > 0 ?
                                <div className="ms-auto mr-2 h-6 w-6  bg-blue-300 flex flex-col justify-center  rounded-full">
                                    <p className="text-gray-500 text-center text-sm md:text-xl">
                                        {user.msgs}
                                    </p>
                                </div> : null
                            } */}
                        </li>
                    )
                })
            }
        </ul>
    )
}
export default OnlineEmployeesList