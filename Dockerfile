# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.13-alpine

# copy the local package files to the container workspace
ADD . /go/src/reformatting_tool

# Setting up working directory
WORKDIR /go/src/reformatting_tool

# Build the users command inside the container.
RUN go install reformatting_tool


#####################################

# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.13-alpine

# Install dependencies
RUN apk update && apk upgrade && apk add --no-cache bash git openssh && \
    go get github.com/Jeffail/gabs github.com/gorilla/mux github.com/auth0/go-jwt-middleware github.com/codegangsta/negroni github.com/dgrijalva/jwt-go github.com/go-redis/redis github.com/basgys/goxml2json

# copy the local package files to the container workspace
ADD . /go/src/feed

# Setting up working directory
WORKDIR /go/src/feed

# Build the users command inside the container.
RUN go install feed

# Run the users microservice when the container starts.
ENTRYPOINT /go/bin/feed