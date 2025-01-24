import { createTypeSpecLibrary } from "@typespec/compiler";

export const $lib = createTypeSpecLibrary({
  name: "typespec-go",
  diagnostics: {},
});

export const internalLib = {
  name: "typespec-go",
  diagnostics: {},
} as const;

export const { reportDiagnostic, createDiagnostic } = $lib;
