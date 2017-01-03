FROM golang:1.7-onbuild
RUN go get github.com/mattes/migrate
 
CMD ["./scripts/run-docker.sh"]
