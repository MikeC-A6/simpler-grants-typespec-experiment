import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "@common-grants/core";

import "./routes.tsp";
import "./models.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;
using CommonGrants;

@service({
    title: "Custom Grants API",
    version: "1.0.0"
})
@server("http://localhost:8080", "Local development server")
@doc("Custom implementation of the Common Grants Protocol API")
namespace CustomAPI;

@route("/api")
@tag("Opportunities")
interface Operations extends CustomRoutes.CustomOpportunities {}
