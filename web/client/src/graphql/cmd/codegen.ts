import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  overwrite: true,
  schema: "http://localhost:4000/graphql",
  documents: "src/**/*.tsx",
  generates: {
    "src/graphql/": {
      preset: "client",
      plugins: [],
      config: {
        scalars: {
          DateTime: String,
        },
      },
    },
  },
};

export default config;
