FROM golang

ADD . /go/src/github.com/plutov/games
RUN curl https://glide.sh/get | sh
RUN cd /go/src/github.com/plutov/games && glide install && go test -v ./pkg/...
RUN go build -o /go/src/github.com/plutov/games/web /go/src/github.com/plutov/games/cmd/web/main.go
ENTRYPOINT cd /go/src/github.com/plutov/games && ./web

EXPOSE 8080
