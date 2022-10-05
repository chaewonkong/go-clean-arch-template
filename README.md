![Go Version 1.19](https://img.shields.io/badge/Go%20Version-1.19-blue)

# Go Clean Architecture Template

Go clean architecture template inspired by [golang-standars/project-layout](https://github.com/golang-standards/project-layout)

Uses Echo, Gorm, uber-go/fx.

- [echo](https://echo.labstack.com/): High performance web framework
- [Gorm](https://gorm.io/): ORM library
- [uber-go/fx](https://github.com/uber-go/fx): Dependency Injection framework maintained by Uber.

## Project Layout

Inspired by [golang-standars/project-layout](https://github.com/golang-standards/project-layout).

Checked [go-clean-arch](https://github.com/bxcodec/go-clean-arch/tree/master/article) for an option, yet it seemed a little bit unfarmiliar to Java/Spring developers.

Created template based on project-layout, however implemented popular clean architecture layout(Popular in Korea).
Intended to provide more comfortable and farmiliar environment for other developers.

Overall architectures are as follows(On Update):

- cmd: where main application is located
- internal: contains applications that are not intended to be public. Impelmentation of main application should be here.
- pkg: public packages

```
.
├── README.md
├── cmd
│   └── root.go
├── go.mod
├── go.sum
├── internal
│   └── _app
│       ├── controllers
│       │   ├── modules.go
│       │   ├── root_controller.go
│       │   └── routes.go
│       ├── entities
│       ├── modules.go
│       ├── repositories
│       ├── server.go
│       └── services
├── main.go
├── makefile
└── pkg
```
