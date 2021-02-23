# Image based golang
FROM golang:1.15.7-buster AS build
ARG app_env
ENV APP_ENV $app_env

# First the copy of go .mod/.sum are made.
WORKDIR /src
COPY go.mod .
COPY go.sum .

# Download dependencies only if repository is empty
RUN go mod download
RUN go mod vendor
RUN go mod verify

# Copy all project inside
COPY . .

# Build the executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/gorest.bin -p 4 -v ./cmd/server

# Another stage. From new container, copy binary. Now isn't needed golang development container.
FROM scratch AS bin
COPY --from=build /out/gorest.bin /bin/gorest.bin
ENTRYPOINT ["/bin/gorest.bin"]

# If production is set, then run. Else, use hot reload.
CMD if [ ${APP_ENV} = production ]; \
    then \
    app; \
    else \
    go get github.com/gravityblast/fresh && \
    fresh; \
    fi

EXPOSE 8080