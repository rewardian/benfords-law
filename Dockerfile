FROM golang:1.12.0-alpine3.9 AS build
RUN apk add git

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go get github.com/gorilla/mux && \
go get github.com/jfyne/csvd
 
RUN go build -o main

FROM alpine
WORKDIR /app
COPY --from=build /app/main /app/main
COPY --from=build /app/layouts /app/layouts
EXPOSE 9021
ENTRYPOINT ["/app/main"]
