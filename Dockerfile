FROM golang AS build
WORKDIR /app

# Go modulesがない場合は先に初期化
# RUN go mod init your-app-name

# go.modファイルがある場合のみコピー（エラー回避）
COPY go.* ./
RUN if [ -f go.mod ]; then go mod download; fi

# ソースコードをコピー
COPY . .

# デバッグ用：何があるか確認
RUN ls -la /app

# ビルド（エラーがあれば止まる）
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 実行時イメージ
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=build /app/main .
RUN chmod +x ./main
EXPOSE 8080
CMD ["./main"]