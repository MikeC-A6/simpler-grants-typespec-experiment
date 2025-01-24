import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "../lib/models/index.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;
using CommonGrants.Models;

@service({
  title: "Common Grants Protocol API"
})
@server("http://localhost:8080", "Local development server")
@doc("Common Grants Protocol API")
namespace CommonGrantsAPI;

@route("/api")
@tag("Opportunities")
interface Operations {
  @summary("List all opportunities")
  @get
  @route("/opportunities")
  listOpportunities(): Opportunity[];

  @summary("Get a specific opportunity")
  @get 
  @route("/opportunities/{id}")
  getOpportunity(@path id: string): Opportunity;

  @summary("Create a new opportunity")
  @post
  @route("/opportunities")
  createOpportunity(@body opportunity: Opportunity): Opportunity;

  @summary("Update an existing opportunity")
  @put
  @route("/opportunities/{id}")
  updateOpportunity(@path id: string, @body opportunity: Opportunity): Opportunity;

  @summary("Delete an opportunity")
  @delete
  @route("/opportunities/{id}")
  deleteOpportunity(@path id: string): void;
}