# استفاده از تصویر پایه مناسب برای اجرای Go
FROM golang:1.23 AS builder

# تنظیم متغیر کاری
WORKDIR /app

# کپی کردن فایل‌های مورد نیاز
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ساختن فایل اجرایی
RUN go build -o server ./cmd/main.go

# مرحله دوم: ایجاد تصویر نهایی
FROM debian:bullseye-slim

# نصب ابزارهای مورد نیاز
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# تنظیم متغیر کاری
WORKDIR /app

# کپی فایل اجرایی از مرحله اول
COPY --from=builder /app/server /app/server

# کپی فایل کانفیگ (در صورت نیاز)
COPY ./sample-config.json /app

# متغیر پیش‌فرض برای آدرس کانفیگ
ENV CONFIG_FILE /app/sample-config.json

# فرمان اجرای سرور
ENTRYPOINT ["/app/server"]
CMD ["--config=/app/config.json"]
