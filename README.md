# coffein - Clean Architecture [WIP]

This project is a sample implementation for demonstrating the power of Clean Architecture written in GO.

## Motivation

My purpose in creating this project is to try a clean architecture implementation in Go language. Thanks to Clean
Architecture, I aimed to develop a more manageable application by separating business rules and infrastructure
dependencies.

## Structure

My main starting point when implementing this architecture was to separate business rules and infrastructural
dependencies.
I basically started by creating two different packages, business rules and infra.
In the business rules package, I designed it to contain the content directly related to what the package is associated
with,
without any external dependencies, without any infrastructural dependencies.
In this way, I did not have any external dependencies while developing a business rule.

The business rules package is basically divided into two, the first is the domain package,
I aimed to keep the domain objects used in this package, the second package is the usecases package, in this package,
I designed it to contain command, query services and port definitions where these services will receive the information
they need.
The infra package is basically divided into two parts; the first part is incoming adapters,
in this package I defined the services that trigger the system (for example http services, grpc services or event
handlers).
The second part is outgoing adapters, this package contains the adapters used by the system.
The adapters corresponding to the ports in the usecases package are defined in this package.

```bash
.
.
├── business
│   ├── domain
│   │   └── recipe.go
│   └── usecases
│       ├── command_service.go
│       ├── commands.go
│       ├── commands_test.go
│       ├── ports.go
│       └── query_service.go
├── infra
│   ├── incoming
│   │   └── web
│   │       └── handlers.go
│   └── outgoing
│       ├── recipe
│       │   ├── adapter.go
│       │   ├── model.go
│       │   └── persistence
│       │       └── inmem.go
│       └── user
│           └── adapter.go
├── go.mod
├── go.sum
└── main.go
.
.

```

I paid attention that the model used by the output adapters and the model I used in the domain were separate models. In
this way, the model I used in the domain became independent of the database used and its restrictions. With this
approach, I got ahead of database oriented design and made sure to think about business rules first.

On the other hand, I designed the use-cases to be accessible from outside the package. It was developed to take command
type as a parameter so that it doesn't care about any details about where it is used. In this way, it has been ensured
that no changes are made in the functions according to the place of use, that is, it is not affected by any
infrastructure restrictions.

## Credits

[Alistair Cockburn - Hexagonal Architecture](http://alistair.cockburn.us/Hexagonal+architecture)

[Robert C. Martin - Clean Architecture: A Craftsman’s Guide to Software Structure and Design](https://www.oreilly.com/library/view/clean-architecture-a/9780134494272/)

[Ardanlabs - Service Repo](https://github.com/ardanlabs/service)
