#Source: https://medium.com/easyread/today-i-learned-golang-live-reload-for-development-using-docker-compose-air-ecc688ee076
FROM golang:1.15.7-buster

RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y git \
    make openssh-client

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air -c /app/cmd/server/.air.toml