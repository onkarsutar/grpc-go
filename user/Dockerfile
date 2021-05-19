# Base Image.
FROM golang:1.16.2-alpine3.13

WORKDIR /go/src/gitlab.com/onkarsutar/grpc-go/user/user-server
COPY go.mod ./ go.sum ./

RUN go mod download
COPY ./ ./
# RUN go build 

# CMD ["./user_server"]

RUN go install /go/src/github.com/matzhouse/go-grpc/user/user_server
ENTRYPOINT ["/go/bin/user_server"]
EXPOSE 50001