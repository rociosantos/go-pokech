# POKECH

## Overview

Bring your best friends alive for battles with this single API

## Requirements

- Docker
- Go
- golangci-lint

## Run the project

Install dependencies

```bash
make deps
```

```bash
make run
```

## Testing and Linting

Run unit tests

```bash
make unit-test
```

Run linters

```bash
make lint
```

## Build docker image and run container locally

Create image

```bash
docker build --tag [your tag] .
```

Create and run container based in the image that you created before.

```bash
docker run --name pokech -p 80:8085 [your tag]
```
