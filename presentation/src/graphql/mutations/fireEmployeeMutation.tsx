import { graphql } from "@/gql";

const fireEmployeeMutation = graphql(`
    mutation fireEmployeeMutation($input: FireEmployee!){
        fireEmployee(input: $input){
            id
        }
    }`)

export default fireEmployeeMutation;