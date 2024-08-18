'use client'

import getGraphQLClient from "@/graphql/graphQLClient";
import countriesQuery from "@/graphql/queries/countriesQuery";
import enterpriseByNameQuery from "@/graphql/queries/enterpriseByNameQuery";
import { useSuspenseQueries } from "@tanstack/react-query";
import { useParams } from "next/navigation";
import UpdateBannerImage from "./UpdateBannerImage";
import UpdateLogoImage from "./UpdateLogoImage";
import UpdateWebsite from "./UpdateWebsite";
import UpdateEmail from "./UpdateEmail";
import UpdatePhone from "./UpdatePhone";
import UpdateAddress from "./UpdateAddress";

const Settings: React.FC = () => {
    const { enterpriseId } = useParams()
    const enterpriseName = decodeURIComponent(enterpriseId as string)
    const gqlClient = getGraphQLClient()
    const [{ data: { enterpriseByName: enterprise } }, { data: { countries } }] = useSuspenseQueries({
        queries: [
            {
                queryKey: ["enterpriseByName", enterpriseName],
                queryFn: async () => await gqlClient.request(enterpriseByNameQuery, { name: enterpriseName })
            },
            {
                queryKey: ['countries'],
                queryFn: async () => getGraphQLClient().request(countriesQuery),
            }
        ]
    })
    return (
        <div className="bg-white">
            <UpdateBannerImage enterprise={enterprise} />
            <UpdateLogoImage enterprise={enterprise} />
            <UpdateWebsite enterprise={enterprise} />
            <UpdateEmail enterprise={enterprise} />
            <UpdatePhone enterprise={enterprise} />
            <UpdateAddress enterprise={enterprise} countries={countries} />
        </div>
    );
}

export default Settings;