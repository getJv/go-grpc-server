FROM alpine:latest

RUN mkdir /app

COPY  go-grpc-server /app

EXPOSE 50051

CMD ["/app/go-grpc-server"]