import { graphql } from "@/gql";

const changeUserGenreMutation = graphql(`
    mutation changeUserGenreMutation($input:ChangeUserGenre!){
        changeUserGenre(input: $input){
            id
        }
    }`)

export default changeUserGenreMutation;