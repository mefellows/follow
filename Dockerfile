FROM golang:1.7-onbuild

RUN go get github.com/codegangsta/gin
RUN go get github.com/mattes/migrate
