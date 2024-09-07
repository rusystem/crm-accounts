FROM golang:1.22

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update

# build go app
RUN go mod download
RUN go build -o crm-accounts ./cmd/main.go

RUN chmod +x crm-accounts

CMD ["./crm-accounts"]