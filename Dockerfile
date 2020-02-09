# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.13-alpine

# copy the local package files to the container workspace
ADD . /go/src/reformatting_tool

# Setting up working directory
WORKDIR /go/src/reformatting_tool

# Build the users command inside the container.
RUN go install reformatting_tool

ENTRYPOINT /go/bin/reformatting_tool
