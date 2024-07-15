import { graphql } from "@/gql";

const CountryFragment = graphql(`
  fragment CountryFragment on Country {
    id
    name
  }
`);

export default CountryFragment;