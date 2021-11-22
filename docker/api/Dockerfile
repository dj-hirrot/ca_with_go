FROM golang:1.17.3

# go moduleを使用
ENV GO111MODULE=on

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8080

# freshコマンドでサーバーを起動
CMD ["fresh"]
