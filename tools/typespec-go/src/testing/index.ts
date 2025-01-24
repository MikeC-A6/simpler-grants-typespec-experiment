import { resolvePath } from "@typespec/compiler";
import { createTestLibrary, TypeSpecTestLibrary } from "@typespec/compiler/testing";
import { fileURLToPath } from "url";

export const TypespecGoTestLibrary: TypeSpecTestLibrary = createTestLibrary({
  name: "typespec-go",
  packageRoot: resolvePath(fileURLToPath(import.meta.url), "../../../../"),
});
