FROM golang:1.16
COPY *.go go.* /go/src/
WORKDIR /go/src
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
COPY --from=0 /go/src/app /
CMD ["./app"]
