import { graphql } from "@/gql";

const hireEmployeeMutation = graphql(`
    mutation hireEmployeeMutation($input: HireEmployee!) {
        hireEmployee(input: $input){
            id
        }
    }`)

export default hireEmployeeMutation;