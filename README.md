# Claimclam challenge

## Workdone

- [x] Golang api-gaetway service for podcasts microservice with basic error handling and input validation functions
- [x] Nextjs webapp that render list of podcasts and allow to search for podcasts
   - Handle UI states for:
    - [x] When the podcasts are being fetched.
    - [x] When the response returns some podasts.
    - [x] When the response returns no podcasts that match the search value.

#### Bonus

- [x] Added some styling to the search results.
   - Added MUI Library for components.
- [x] Added some pagination. The api supports "page" and "limit" params.
  - Note: The api don't return a total count so I mocked this in the podcast client
- [x] Added a debounce of 500ms to avoid fetching on every key press.
- [x] Use graphQL as a protocol between UI and API gateway
  - I used graphql codegen to generate typescript types for the graphql api.
- [x] Add basic rate limiting 
- [x] Add health basic health endpoint for gateway service
  
   
## Doc

This repo is a golang monorepo. Here's the repo breakdown:
- `internal`: Contains common libraries for go services.
- `services/*`: Contains all the services.
- `webapp`: Contains the nextjs webapp.
- `deployment`: Meant to contain deployment code (Docker compose, Terraform, K8s...) The application is dockerized and can be run using docker compose.
- `graphql`: Contains the graphql sechemas for the api gateway. 
    - `gen`: Contains the generated code from gqlgen

The graphql api is exposed through the api gateway.
The approcah is a Schema first approach — Types and query are define using the GraphQL and then code is generated using gqlgen.

### Gateway service

This services is a simple api gateway that exposes the graphql api. It is delivered as a CLI and meant to be ship in light docker images. It's build with uber-go/fx for depencies injection and zap for logging.

### Webapp

This is a Nextjs14 webapp that uses the AppRouter. It's build with typescript and MUI. Graphql types are generated using graphql codegen to enhance type safety. Queries are done using @tanstack/react-query which brings amazing error handling, caching etc. I used the built-in React Context API as solution for managing state given the simplicity of the app. However for a more complex state management I would suggest using recoil or redux.

## How to run

Simmply run `make dev`. It will start the docker-compose stack. To clean up the stack run `make dev.clean` and to rebuild containers run `make dev.build`

### Run unit tests

Run `make tests.unit`

### Generate graphql code for gateway
Run `make gql`

### Generate graphql code for webapp
Run `yarn codegen` inside webapp directory

