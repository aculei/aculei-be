# aculei-be

![go](https://img.shields.io/badge/Go-00ADD8.svg?style=plain&logo=Go&logoColor=white)
![docker](https://img.shields.io/badge/Docker-2496ED.svg?style=plain&logo=Docker&logoColor=white)

Server for [`aculei.xyz`](https://aculei.xyz)

## Install

### ![github](https://img.shields.io/badge/GitHub-181717.svg?style=plain&logo=GitHub&logoColor=white)

Clone the repository

```console
git clone https://github.com/aculei/aculei-be.git
```

### ![docker](https://img.shields.io/badge/Docker-2496ED.svg?style=plain&logo=Docker&logoColor=white)

```console
docker pull ghcr.io/aculei/aculei-be:main
```

## Run

### ![github](https://img.shields.io/badge/GitHub-181717.svg?style=plain&logo=GitHub&logoColor=white)

At `root` of the project run

```console
go run main.go
```

or if you have [task](https://taskfile.dev/installation/) installed

```console
task run
```

### ![docker](https://img.shields.io/badge/Docker-2496ED.svg?style=plain&logo=Docker&logoColor=white)

```console
docker run -d -p 8080:8080 ghcr.io/aculei/aculei-be:main
```

Platform specific

```console
docker run -d -p 8080:8080 ghcr.io/aculei/aculei-be:main --platform linux/amd64
```

[![task](https://img.shields.io/badge/Task-29BEB0.svg?style=plain&logo=Task&logoColor=white)](https://taskfile.dev/installation/)

Install `task` (installation guide [here](https://taskfile.dev/installation/)). Then at `root` of the project run

```console
task build
```

Then under `/bin` you'll see a binary file, just open it and you'll have the backend up and running

<!-- ![screenshot](docs/bin-screenshot.png) -->

## Documentation

Swagger available at `host`/swagger/index.html
