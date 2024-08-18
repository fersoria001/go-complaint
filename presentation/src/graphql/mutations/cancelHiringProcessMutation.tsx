import { graphql } from "@/gql";

const cancelHiringProcessMutation = graphql(`
    mutation cancelHiringProcessMutation($input: CancelHiringProcess!){
        cancelHiringProcess(input: $input){
            id
        }
    }`)
    
export default cancelHiringProcessMutation