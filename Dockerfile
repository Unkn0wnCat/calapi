# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN --mount=type=cache,target=~/.cache/go-build,type=cache,target=~/go/pkg/mod go mod download

COPY . .

RUN --mount=type=cache,target=~/.cache/go-build,type=cache,target=~/go/pkg/mod CGO_LDFLAGS="-L./lib -Wl,-rpath,\$ORIGIN/lib" go build -v

##
## Package
##
FROM alpine:3.15

RUN apk add gcompat libgcc libstdc++

EXPOSE 8080
WORKDIR /app
VOLUME /app/data
COPY ./lib ./lib

ENTRYPOINT ["/app/calapi"]
CMD ["serve"]

COPY --from=build /app/calapi /app/calapi