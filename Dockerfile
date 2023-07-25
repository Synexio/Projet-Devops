FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go get main
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs

EXPOSE 8080

CMD ["sleep","3600"]