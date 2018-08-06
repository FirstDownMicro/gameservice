FROM golang:alpine as builder
RUN mkdir -p $GOPATH/src/github.com/Nanixel/FirstDownMicro/gameservice 
#right now this just works locally, but eventually I will run a go get on this repo
ADD . $GOPATH/src/github.com/Nanixel/FirstDownMicro/gameservice
RUN cd $GOPATH/src/github.com/Nanixel/FirstDownMicro/gameservice && go install

FROM alpine
WORKDIR /app
COPY --from=builder /go/bin /app
ENTRYPOINT ./gameservice

