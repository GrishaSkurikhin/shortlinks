FROM golang:1.18.4

RUN mkdir -p /app
ADD . /app/
WORKDIR /app

RUN go mod download && go mod verify
RUN go build -o bin ./cmd/web

EXPOSE 4000

CMD ["/app/bin"]