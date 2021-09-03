FROM golang:1.17-alpine
WORKDIR /go/src/app
COPY . .
RUN go get
RUN go install
CMD [ "go", "run", "main.go" ]
