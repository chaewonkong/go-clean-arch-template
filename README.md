![Go Version 1.19](https://img.shields.io/badge/Go%20Version-1.19-blue)

# NASS v2 Admin API with Go

NASS v2 Admin의 API 서버입니다.
Go를 이용해 개발합니다.

## Project Layout

[golang-standars/project-layout](https://github.com/golang-standards/project-layout)을 참고해 clean-architecture를 도입합니다.

[go-clean-arch](https://github.com/bxcodec/go-clean-arch/tree/master/article)의 도입도 검토하였으나, 위 project-layout에 비해 전체적인 star 수가 적고 한국에서 일반적으로 사용되는 아키텍처와는 조금 다른 느낌이 있어 project-layout을 활용하되 Java/Spring에서 일반적으로 사용되는 clean-architecture 형태를 도입하여 다른 개발자에게 보다 익숙한 환경을 제공하려고 하였습니다.

전반적인 아키텍처는 아래와 같습니다.

- cmd: 메인 애플리케이션의 실행 CLI가 위치합니다.
- internal: 외부에 공개하지 않을 목적의 내부 코드를 보관합니다. 본 프로젝트의 경우 애플리케이션의 코드가 여기 보관됩니다.
- pkg: 외부 공개용 코드의 위치입니다.

```
.
├── README.md
├── cmd
│   └── root.go
├── go.mod
├── go.sum
├── internal
│   ├── controllers
│   │   └── main.go
│   ├── entities
│   ├── repositories
│   └── services
├── main.go
└── pkg
```
