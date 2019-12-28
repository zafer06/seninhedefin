FROM golang:1.13.4

RUN go get github.com/beego/bee

ENV GO111MODULE=on
RUN mkdir /go/src/api
WORKDIR /go/src/api

COPY go.mod .
COPY go.sum .

RUN go mod download

EXPOSE 8081

COPY . .

CMD ["bee", "run", "-downdoc=true", "-gendoc=true"]
