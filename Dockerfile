# Image based golang
FROM golang:1.15.7-buster AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir -p /goRestProject
WORKDIR /goRestProject

# First, were copying go mod and go sum for cache downloads in newer builds.
COPY go.mod .
COPY go.sum .

RUN echo "Download dependencies"
RUN go mod download -x
RUN go mod verify

# Now, the rest of the project
COPY . .

RUN echo "Building server binary"
RUN go build -gcflags "all=-N -l" -o ./bin/goRest.bin -p 4 -v ./cmd/server/
#RUN gebug init
#RUN gebug start