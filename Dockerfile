FROM golang:1.16.15

RUN apt update

COPY $PWD /ChatGPT
WORKDIR /ChatGPT

EXPOSE 8080

ENTRYPOINT ["go","run","main.go"]