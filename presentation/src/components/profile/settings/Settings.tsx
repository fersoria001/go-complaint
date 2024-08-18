'use client'
import getGraphQLClient from "@/graphql/graphQLClient";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";
import { useSuspenseQueries } from "@tanstack/react-query";
import UpdatePassword from "./UpdatePassword";
import UpdateProfileImage from "./UpdateProfileImage";
import UpdateGenre from "./UpdateGenre";
import UpdateFirstName from "./UpdateFirstName";
import UpdateLastName from "./UpdateLastName";
import UpdatePhone from "./UpdatePhone";
import UpdateAddress from "./UpdateAddress";
import countriesQuery from "@/graphql/queries/countriesQuery";

const Settings: React.FC = () => {
    const gqlClient = getGraphQLClient()
    const [{ data: { userDescriptor: descriptor } }, { data: { countries } }] = useSuspenseQueries({
        queries: [
            {
                queryKey: ['userDescriptor'],
                queryFn: async () => await gqlClient.request(userDescriptorQuery)
            },
            {
                queryKey: ['countries'],
                queryFn: async () => gqlClient.request(countriesQuery),
            }
        ]

    })

    return (
        <div className="bg-white flex flex-col">
            <UpdateProfileImage descriptor={descriptor} />
            <UpdatePassword descriptor={descriptor} />
            <UpdateGenre descriptor={descriptor} />
            <UpdateFirstName descriptor={descriptor} />
            <UpdateLastName descriptor={descriptor} />
            <UpdatePhone descriptor={descriptor} />
            <UpdateAddress descriptor={descriptor} countries={countries} />
        </div>
    )
}

export default Settings;