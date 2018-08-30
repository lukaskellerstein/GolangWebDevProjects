FROM golang
COPY . $GOPATH/src/github.com/GolangWebDevProjects/test



# Cellar websockets - get packages & build
WORKDIR $GOPATH/src/github.com/GolangWebDevProjects/test
RUN go-wrapper download
RUN go-wrapper install
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o test  

# helpers
# RUN echo $GOPATH
RUN pwd
RUN ls -l


# -----------------------
# Stage 2
# -----------------------
#  FROM alpine:latest
#  RUN apk --no-cache add ca-certificates
FROM alpine:3.5
COPY --from=0 /go/src/github.com/GolangWebDevProjects/test/test /app/test
RUN chmod +x /app/test

RUN ls -l  

EXPOSE 10000
ENTRYPOINT /app/test


