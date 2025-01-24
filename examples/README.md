# Examples

This directory contains example implementations of the Common Grants Protocol API.

## Custom API Example

The `custom-api` directory demonstrates how to extend and customize the base Common Grants Protocol API using TypeSpec.

### Getting Started

1. Install dependencies:
```bash
cd custom-api
npm install
```

2. Compile the TypeSpec files:
```bash
tsp compile src/main.tsp
```

This will:
- Generate the OpenAPI specification in `generated/openapi.yaml`
- Generate the Go API code in `server.gen.go`

### Making Changes

To customize the API:

1. Modify the TypeSpec files in `src/`:
   - `models.tsp`: Define custom models and extend base models
   - `routes.tsp`: Define custom routes and override base routes
   - `main.tsp`: Configure the API service

2. Recompile to apply changes:
```bash
tsp compile src/main.tsp
```

### Example: Adding a Custom Field

Here's an example of how to add a new field to the `CustomOpportunity` model:

1. Open `src/models.tsp`
2. Add your field to the appropriate model:
```typescript
model OpportunityMetadata {
    // ... existing fields ...
    
    @doc("Your new field description")
    newField: string;
}
```
3. Run `tsp compile src/main.tsp` to propagate the changes

The changes will be automatically reflected in:
- The OpenAPI specification (`generated/openapi.yaml`)
- The generated Go API code (`server.gen.go`)

### Using the Generated API

The TypeSpec compiler generates two key files:

1. `generated/openapi.yaml`: The OpenAPI specification that documents your API
2. `server.gen.go`: The Go code that defines your API's types and server interface

To implement your API:

1. Create a new Go project
2. Import the generated `server.gen.go` file
3. Implement the `ServerInterface` with your business logic
4. Use the generated types for request/response handling

Example implementation:
```go
package api

type OpportunityServer struct {}

// Implement the ServerInterface
func (s *OpportunityServer) OperationsRead(c *gin.Context, id string) {
    // Your implementation here
}

func (s *OpportunityServer) OperationsUpdate(c *gin.Context, id string) {
    // Your implementation here
}
```

### Project Structure

- `src/`: TypeSpec source files
  - `models.tsp`: Custom model definitions
  - `routes.tsp`: Custom route definitions
  - `main.tsp`: API service configuration
- `generated/`: Generated API files
  - `openapi.yaml`: OpenAPI 3.0 specification
- `server.gen.go`: Generated Go API code
