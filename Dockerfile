# Image based golang
FROM golang:1.15.7-buster AS build

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/app
WORKDIR $GOPATH/src/app

# First, were copying go mod and go sum for cache downloads in newer builds.
COPY go.mod .
COPY go.sum .

RUN go mod download -x
RUN go get github.com/githubnemo/CompileDaemon
RUN go mod verify

# Now, the rest of the project
COPY . $GOPATH/src/app

# Compiling
#CGO_ENABLED=0
#RUN GOOS=linux GOARCH=amd64 go build -o /out/gorest.bin -p 4 -v ./cmd/server

# Another stage. From new container, copy binary. Now isn't needed golang development container.
#FROM debian:buster
#COPY --from=build ./out/gorest.bin /goRest/gorest.bin

#ENTRYPOINT ["/goRest/gorest.bin"]
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o ./bin/goRest.bin -p 4 -v ./cmd/server/" -command="./bin/goRest.bin" --color