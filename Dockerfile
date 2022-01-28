#Ensures Docker pulls the offical Go image
FROM golang:1.17.6 AS GO_BUILD

ENV CGO_ENABLED 0
#copies content over
COPY . /go-app
#sets the working directory.
WORKDIR /go-app
#builds your app into a binary
RUN go build -o server

#uses alpine as our runner image in which our GO application will run
FROM alpine:3.15

WORKDIR /go-app
#copies the binary built earlier into the image
COPY --from=GO_BUILD /go-app/server /go-app/server
#exposes application over port 8080
EXPOSE 8080
#sets the default command to run when the container is run
CMD ["./server"]