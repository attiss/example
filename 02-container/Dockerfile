FROM golang:1.18

WORKDIR /build

COPY . /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o lorem-ipsum .

EXPOSE 80

ENTRYPOINT ["/build/lorem-ipsum"]
