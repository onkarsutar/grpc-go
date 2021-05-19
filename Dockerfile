# Base Image.
FROM golang:1.16.2-alpine3.13

WORKDIR /go/src/gitlab.com/onkarsutar/grpc-go/
COPY go.mod ./  go.sum ./

RUN go mod download

WORKDIR /go/src/gitlab.com/onkarsutar/grpc-go/user/user_server
COPY ./ ./
# RUN go build 

# CMD ["./user_server"]

RUN go install ./

ENTRYPOINT ["/go/bin/grpc-go"]
EXPOSE 50001