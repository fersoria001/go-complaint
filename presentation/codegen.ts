// codegen.ts
require('dotenv').config()
import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  schema: process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT,
  documents: 'src/graphql/**/*.tsx',
  ignoreNoDocuments:true,
  generates: {
    './src/gql/': {
      preset: 'client'
    }
  }
}

export default config