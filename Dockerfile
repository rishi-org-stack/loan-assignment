FROM golang:latest
LABEL maintainer="Rishi"
EXPOSE 8000
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["loan"]