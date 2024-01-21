# 基本イメージ
FROM golang:1.21.5-alpine3.18
#FROM golang:1.21.5-alpine3.18 as builder

# セキュリティアップデートとツールのインストール
RUN apk update && apk upgrade && \
    apk add --no-cache git

# ワーキングディレクトリの設定
WORKDIR /data
COPY . .

# 依存関係のキャッシュ
WORKDIR /data/app
RUN go mod download

# デバッグと開発ツールのインストール
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install go.uber.org/mock/mockgen@v0.3.0
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.23.0
RUN go install github.com/cosmtrek/air@latest

# ビルド (本番用イメージの場合、ここでビルドしておく)
# RUN go build -o myapp .

# 本番用イメージの場合、マルチステージビルドを使用する
# FROM alpine:latest  
# WORKDIR /root/
# COPY --from=builder /app/myapp .
# CMD ["./myapp"]

# 開発用イメージの場合のコマンド
CMD ["air", "-c", ".air.toml"]

# 非ルートユーザーでの実行 (本番環境での実行時に有効にする)
# USER nobody
