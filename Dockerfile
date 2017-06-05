FROM bluele/go-mecab AS build-env
ENV GOPATH /go
ADD . /go/src/github.com/bluele/mecab-docker
WORKDIR /go/src/github.com/bluele/mecab-docker
RUN CGO_LDFLAGS="`mecab-config --libs`" CGO_FLAGS="`mecab-config --inc-dir`" go build -o hello app.go

FROM bluele/mecab
COPY --from=build-env /go/src/github.com/bluele/mecab-docker/hello /usr/local/bin/hello
ENTRYPOINT ["/usr/local/bin/hello"]
