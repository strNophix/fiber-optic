# fiber-optic
Personal gofiber template with goodies like a Dockerfile, Swagger and Just.

## Features
### Dockerfile
Every created repository using this template includes a Dockerfile. This can be used to build production-ready container images. The image leverages multi-stage build to minimize the size of the image and [dumb-init](https://github.com/Yelp/dumb-init) to reduce the complexity of the container.

### Just
Every created repository using this template includes a Justfile. This file defines all commands that can be ran within the repository. Just has a syntax inspired by make but has a lot of additional features like OS-specific variants of the same recipe ie. `just build` and loading of `.env`-files. The Justfile is used in this case to store procedure's for things like building and containerizing.

### Swagger
Every created repository using this template has Swagger configured. Swagger can be used as a means of documenting your API and exploring endpoints.

### Much more...
Some features are very subtle details in among other things the directory structure and the way things like routing are handled.

## Dependencies
- [cookiecutter](https://github.com/cookiecutter/cookiecutter)
- [just](https://github.com/casey/just)
- [podman](https://github.com/containers/podman)
- [go 1.20](https://github.com/golang/go)

## Usage
```sh
cookiecutter https://git.cesium.pw/niku/fiber-optic.git
```

Which generates the following structure where among other things demo-app will be substituted with the name/data you provide:
```
.
├── cmd
│   └── demo-app
│       └── main.go
├── Dockerfile
├── .dockerignore
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── .gitignore
├── go.mod
├── go.sum
├── internal
│   └── app
│       ├── api
│       │   ├── router.go
│       │   └── v1
│       │       └── router.go
│       └── main.go
└── justfile
```