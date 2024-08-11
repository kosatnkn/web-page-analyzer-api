# Web Page Analyzer API

## Overview
This is the backend component of the `Web Page Analyzer`. It is tasked with reading tags of a given web page and generate a summary. The service exposes a REST endpoint that accepts a URL and an array of HTML tag names.

The `OpenAPI` document of the RESTful interface can be found at [docs/api/openapi.yaml](https://github.com/kosatnkn/web-page-analyzer-api/blob/main/docs/api/openapi.yaml).

> **IMPORTANT**
>
> There are two more pieces to the solution.
> - Web Page Analyzer Client, the frontend component ([https://github.com/kosatnkn/web-page-analyzer-client](https://github.com/kosatnkn/web-page-analyzer-client))
> - Deployer, the IoC (infrastructure as code) component that enables single command deployment ([https://github.com/kosatnkn/web-page-analyzer-deploy](https://github.com/kosatnkn/web-page-analyzer-deploy))

## Running
There are two ways to run the service (**from source** or **from a Docker container**). But before that the service needs to be configured.

### Configuring
A set of `.yaml` files are used to configure different aspects of the service. These files are read from the `configs` directory placed at the same level as the compiled binary.

There are sample config files already in the [configs/](https://github.com/kosatnkn/web-page-analyzer-api/tree/main/configs) directory.

Simply make a copy of each `*.yaml.example` file and rename to `*.yaml`. Configurations in these files are minimal and self explanatory. The default values are sufficient to get things started.

### Running From Source
Once you clone the repository and setup configurations use the following make command to run the service.
```bash
make run
```

### Run as a Docker Container
You can also run the service as a Docker container.

> **NOTE:** Before starting make sure that all configurations are in place ([configs/](https://github.com/kosatnkn/web-page-analyzer-api/tree/main/configs))

Use the following commands to build and spin-up the container.
```bash
make docker_run
```
> **NOTE:** This will build and tag the Docker image with `kosatnkn/web-page-analyzer-api:latest`(by default) and then run it. This behavior can be changed from the Makefile.

## Testing
Tests can be run using the following command
```bash
make test
```

## Key Components
This service uses the [Catalyst](https://github.com/kosatnkn/catalyst) microservice base project as its base and follows the **Clean Architecture** paradigm.

The service handles the **Request, Response Cycle** as follows.
```text
                               + ------- +           + -------- +
                               | REQUEST |           | RESPONSE |
                               + ------- +           + -------- +
                                   ||                     /\
                                   \/                     ||
                            + ------------ +              ||
   (Intercepts all requests |  Middleware  |              ||
    coming in)              + ------------ +              ||
                                   ||                     ||
                                   \/                     ||
                            + ------------ +              ||
        (Route request to a |    Router    |              ||
         `Handler` func)    + ------------ +              ||
                                       ||                 ||
                                       ||                 ||
                                       ||   + --------------------------- +
                                       ||   | Transformer | Error Handler | (Transform
                                       ||   + --------------------------- +  responses)
    (convert incoming data             ||    /\
     into domain objects)              \/    ||
    + -------------------- +  =>  + -------------- +
    | Unpacker | Validator |      |   Controller   | (Contain `Handler` functions,
    + -------------------- +  <=  + -------------- +  controls the req, resp cycle)
                                      ||       /\
                                      \/       ||
                                  + -------------- +
                                  |    Use Case    | (`Domain Logic` resides here)
                                  + -------------- +
                                      ||       /\
                                      \/       ||
                          _____________________________________
                              + ---------- +    + ------- +
                              | Repository |    | Service | (Interfacing logic needed to
                              + ---------- +    + ------- +  communicate with external
                                ||    /\          ||  /\     resources are here)
                                \/    ||          \/  ||
                              + ---------- +    + ------- +
                              |  Database  |    |   APIs  | (External resources)
                              + ---------- +    + ------- +
```

Code for each of these components are organized in the file tree of the project as follows.
```text
web-page-analyzer-api/
├── app
│   ├── adapters ===> (Interfaces of commonly used adapters)
│   │   ├── LogAdapterInterface.go
│   │   └── ...
│   ├── config
│   │   ├── Config.go ===> (Configuration object for the service)
│   │   └── ...
│   ├── container
│   │   ├── Container.go ===> (Dependency injection container)
│   │   └── ...
│   ├── metrics
│   │   └── ... ===> (Additional metric registering logic)
│   └── splash
│       └── ... ===> (A splash message when starting the service)
├── configs                  _
│   ├── app.yaml              |
│   ├── app.yaml.example       > (YAML configuration files)
│   ├── logger.yaml           |
│   └── logger.yaml.example   |
├── Dockerfile
├── docs
│   └── api
│       └── openapi.yaml ===> (The OpenAPI REST interface of the service)
├── domain
│   ├── boundary ===> (Interfaces defining the domain boundary)
│   │   └── services
│   │       └── WebPageServiceInterface.go
│   ├── entities ===> (Data structures used to communicate with the domain)
│   │   ├── Component.go
│   │   ├── Page.go
│   │   └── ...
│   ├── errors
│   │   └── DomainError.go ===> (Error objects used in the domain)
│   └── usecases ===> (Domain logic)
│       └── webpage
│           ├── WebPage.go
│           └── ...
├── errors
│   └── BaseError.go ===> (Base error object all other error objects are based on)
├── externals ===> (Logic needed to communicate with third party libs, services, etc.)
│   └── services
│       ├── errors
│       │   └── ...
│       ├── WebPageService.go
│       └── ...
├── main.go
├── Makefile
└── transport
    ├── http ===> (RESTful endpoints handling package)
    └── metrics ===> (Metrics endpoint handling package)
```

## Future Work
The **WebPageService** at `web-page-analyzer-api/externals/services/WebPageService.go` is the main component that does all the heavy lifting. Everything else is basically plumbing. In order to improve the overall service, this piece needs to improve.
- Set up benchmarking so that the service performance can be monitored properly.
- Use parallel processing when it comes to HTTP token parsing.

---
Powered By [Catalyst](https://github.com/kosatnkn/catalyst)
