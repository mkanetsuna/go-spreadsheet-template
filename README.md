# Jooto Go WebApp Template

Goでwebアプリケーションを立ち上げる時のテンプレートです。

## Development

ホットリロードを伴うdevモードでアプリケーションをrunする:

```
make install-tools // 最初だけ
make dev
```

 - これは`http://localhost:8080`でサーバーを立ち上げ、リロードをすればコードの変更を即座に反映することができる

## Building

```
make build
```

This will create a binary in the `bin` directory.

## Running

To build and run the application:

```
make run
```

## Endpoints

- `/scheduled-task`: Simulates a scheduled task
- `/webhook1`: Webhook endpoint 1
- `/webhook2`: Webhook endpoint 2
- `/health`: Health check endpoint
- `/`: Development endpoint