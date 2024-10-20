# aculei-be

![go](https://img.shields.io/badge/Go-00ADD8.svg?style=plain&logo=Go&logoColor=white)
[![task](https://img.shields.io/badge/Task-29BEB0.svg?style=plain&logo=Task&logoColor=white)](https://taskfile.dev/installation/)
![swagger](https://img.shields.io/badge/Swagger-85EA2D.svg?style=plain&logo=Swagger&logoColor=black)

Server for [`aculei.xyz`](https://aculei.xyz)

## Install

Clone the repository

```console
git clone https://github.com/micheledinelli/aculei-be.git
```

> [!NOTE]
> Docker image coming soon

## Run

### From source

At `root` of the project run

```console
go run main.go
```

or if you have [task](https://taskfile.dev/installation/) installed

```console
task run
```

### Built binary

Install `task` (installation guide [here](https://taskfile.dev/installation/)). Then at `root` of the project run

```console
task build
```

Then under `/bin` you'll see a binary file, just open it and you'll have the backend up and running

![screenshot](docs/bin-screenshot.png)

## Documentation

Swagger available at `host`/swagger/index.html
