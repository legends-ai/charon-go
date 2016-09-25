FROM golang:1.7.1-wheezy

# Copy files. we don't need glide etc with this
COPY . /go/src/github.com/asunaio/charon
WORKDIR /go/src/github.com/asunaio/charon

# Build binary
RUN go build .
