FROM golang:1.14

WORKDIR /hello
COPY . .

RUN go mod download -x rsc.io/quote
RUN go version
RUN go env
