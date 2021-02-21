# Con base en la imágen golang
FROM golang
ARG app_env
ENV APP_ENV $app_env

# Se copia el directorio del proyecto a la ruta del contenedor /go/.... y lo usa como referencia para más adelante.
COPY ./ /go/src/github.com/goRest
WORKDIR /go/src/github.com/goRest

# Se descargan las dependencias y construye el proyecto
#RUN go get .
RUN go mod download
RUN go mod vendor
RUN go mod verify
RUN go build .

# Si se especifica production, ejecuta la aplicación; sino, ejecuta el código con recarga en caliente.
CMD if [ ${APP_ENV} = production ]; \
    then \
    app; \
    else \
    go get github.com/gravityblast/fresh && \
    fresh; \
    fi

EXPOSE 8080