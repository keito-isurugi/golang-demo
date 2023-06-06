FROM golang:1.18-alpine

RUN apk update &&  apk add git
RUN go install github.com/cosmtrek/air@v1.29.0
WORKDIR /app

# air -c [tomlファイル名] // 設定ファイルを指定してair実行(WORKDIRに.air.tomlを配置しておくこと)
CMD ["air", "-c", ".air.toml"]