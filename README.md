# Turbo Waffle

This project is a REST API application in the Golang language.
Written for educational purposes.

The main idea was to learn how to write applications using the Pure Architecture approach in building the application structure, as well as to learn how to use the dependency injection technique. At the same time I wanted to learn how to work with the database running in Docker, and to set up the migration of that database.

If you, like me, are in the process of learning Go and do not know how to start writing your first big program, I hope the structure of my project will help you understand how an application in Go should look like. 

[I highly recommend reading it](https://github.com/bxcodec/go-clean-arch)

Here is my project's file structure tree to make it easier for you to understand.
```
.
├── app
│   ├── app.go
│   └── config.go
├── controller
│   └── httpcontroller
│       ├── auth.go
│       ├── controller.go
│       ├── middleware.go
│       ├── response.go
│       ├── router.go
│       ├── shortener.go
│       └── statistics.go
├── domain
│   ├── link.go
│   └── user.go
├── infrastructure
│   └── repository
│       ├── pgrepository
│       │   ├── auth.go
│       │   ├── config.go
│       │   ├── postgres.go
│       │   ├── shortener.go
│       │   └── statistics.go
│       └── repository.go
└── usecase
    ├── auth.go
    ├── shortener.go
    ├── statistics.go
    └── usecase.go
```
