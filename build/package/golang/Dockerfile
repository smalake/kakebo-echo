FROM golang:1.21

RUN apt-get update && apt-get -y install zip unzip

# ENV PATH=${PATH}:${GOPATH}/bin
# メモリが掴みっぱなしにならないように
ENV GODEBUG=madvdontneed=1

WORKDIR /go/src/kakebo-echo

ENV GO111MODULE=on
ENV GOOGLE_APPLICATION_CREDENTIALS=/go/src/kakebo-echo/build/package/golang/google-key.json

# RUN go mod init kakebo-echo && go get -u github.com/cosmtrek/air@latest
RUN go install github.com/cosmtrek/air@v1.41.0
ENTRYPOINT ["air", "-c", ".air.conf"]