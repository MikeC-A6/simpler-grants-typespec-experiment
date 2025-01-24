import { EmitContext, emitFile, resolvePath, navigateProgram, Model } from "@typespec/compiler";
import { generateGoStruct } from "./go-types.js";

const INTERNAL_NAMESPACES = new Set([
  "TypeSpec",
  "Rest",
  "Http",
  "OpenAPI",
  "JsonSchema",
]);

export async function $onEmit(context: EmitContext) {
  if (context.program.compilerOptions.noEmit) {
    return;
  }

  const models: string[] = [];
  
  navigateProgram(context.program, {
    model(m: Model) {
      // Skip internal types
      if (m.namespace && INTERNAL_NAMESPACES.has(m.namespace.name)) {
        return;
      }
      const struct = generateGoStruct(m);
      if (struct) {
        models.push(struct);
      }
    },
  });

  const content = `package api

import (
    "time"
)

${models.join("\n\n")}
`;

  await emitFile(context.program, {
    path: resolvePath(context.emitterOutputDir, "models.go"),
    content,
  });
}
