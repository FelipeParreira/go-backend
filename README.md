# iDev Soluções' Engineering Challenge

## Contents
  - [Description](#description)
  - [Commands](#commands)
  - [Design Choices](#design-choices)
  - [Project Structure](#project-structure)
  - [Next Steps](#next-steps)
  - [Docs](#docs)

## Description

* This is a RESTful API written in Go (1.14);
* Data is persisted through [*JSON* file](./data.json);
* The endpoints are documented using Swagger and can be tested using the [Postman collection and environment attached](./postman);


## [Commands](./Makefile)

Command | Description | 
--- | --- 
`$ make run` | Run app (development mode) *
`$ make shutdown` | Shutdown app (development mode)
`$ make test` | Run tests 
`$ make doc` | Generate swagger doc files **
`$ make lint` | Run linter
`$ make pipeline/qa` | QA pipeline (linter, audit dependencies & tests)

\* runs development image (with environment variables)

\** once the server is running locally, the Swagger docs can be seen [here](http://localhost:8080/swagger/index.html)

## Design Choices

* Gingonic was used as the HTTP web framework, for its [benchmarked performance](https://github.com/gin-gonic/gin#benchmarks) (avoiding reflection and a small memory footprint, and with an efficient router);
* A [MVC architecture](#project-structure) was followed, in which the model defines the entities in the given [domain](./domain); the [controller](./controllers) exposes an interface to deal with the entities, and the view (or [services](./services), as they are called here) provide the business rules for the manipulation and presentation/visualization of the entities.
* This approach keeps the code more modular (with separation of concerns), extensible and decoupled, improving developers' experience.

## Project Structure
Folder/File | Description
--- | --- 
[`main.go`](./main.go) | Entry point for the app; starts application.
[`app/`](./app) | Defines the [HTTP server](./app/application.go); [maps the URLs](./app/url_mapping.go).
[`controllers/`](./controllers) | Defines the functions called for each of URLs; they handle HTTP requests and respond to them.
[`services/`](./services) | Defines and apply business rules.
[`domain/`](./domain) | Defines the data models of the domain, some related utilities and CRUD operations.
[`utils/`](./utils) | Defines utility functions.
[`logger/`](./logger) | Sets up the logger, defined as a global variable.
[`docs/`](./docs) | Documentation files generated by Swagger.
  
## Next Steps

* Add more unit tests;
* Add integration/E2E tests;
* Add performance tests;

## Docs
* [Golang](https://golang.org/doc/)
* [GinGonic](https://gin-gonic.com/docs/)
* [Swaggo](https://github.com/swaggo/swag#how-to-using-security-annotations)
* [Docker](https://docs.docker.com)
