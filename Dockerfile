# Image based golang
FROM golang:1.15.7-buster AS build
ARG app_env
ENV APP_ENV $app_env

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/app
ADD . $GOPATH/src/app

WORKDIR $GOPATH/src/app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/gorest.bin -p 4 -v ./cmd/server

# Another stage. From new container, copy binary. Now isn't needed golang development container.
FROM scratch AS bin
COPY --from=build ./out/gorest.bin /goRest/gorest.bin
ENTRYPOINT ["/goRest/gorest.bin"]

# If production is set, then run. Else, use hot reload.
CMD if [ ${APP_ENV} = production ]; \
    then \
    app; \
    else \
    go get github.com/gravityblast/fresh && \
    fresh; \
    fi

EXPOSE 8080