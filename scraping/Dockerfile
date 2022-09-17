# デプロイ用のコンテナ
FROM golang:1.18.2-bullseye as deploy

WORKDIR /app

# 開発環境用
FROM golang:1.18.2 as dev

WORKDIR /app/scraping
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
