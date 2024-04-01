FROM golang:latest
WORKDIR /go/tours/tours-service
RUN go mod init tours-service
COPY tours .
RUN go build -o main .
EXPOSE 81
CMD ["./main"]