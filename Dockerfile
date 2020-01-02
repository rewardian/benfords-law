FROM golang:1.12.0-alpine3.9
RUN apk add git

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go get github.com/gorilla/mux && \
go get github.com/jfyne/csvd && \
go get github.com/rewardian/benfords-law/layouts
 
RUN go build -o main .
EXPOSE 9021
CMD ["/app/main"]
