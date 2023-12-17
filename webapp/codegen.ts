
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "../graphql/podcasts.graphqls",
  documents: "src/**/*.tsx",
  ignoreNoDocuments: true, // for better experience with the watcher
  generates: {
    "src/gql/": {
      preset: "client",
      plugins: []
    },
  }
};

export default config;
