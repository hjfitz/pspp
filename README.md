# PSPP: Pub/Sub Push Proxy

A lightweight local emulator that transforms plain HTTP requests into [Google Cloud Pub/Sub push request](https://cloud.google.com/pubsub/docs/push) format and forwards them to your upstream endpoint.

Useful for local development when you want to test how your service handles Pub/Sub push deliveries without running the full emulator or connecting to GCP.

## How it works

```
Your request → PSPP (adds Pub/Sub envelope) → Your upstream endpoint
```

PSPP accepts any HTTP POST body, wraps it in the standard Pub/Sub push envelope (base64-encoded `data`, `messageId`, `publishTime`, `subscription`), and proxies the result to your configured upstream URL.

## Installation

### From source

**Requirements:** Go 1.25+

```bash
git clone https://github.com/hjfitz/pspp.git
cd pspp
go build ./cmd/proxy
```

### Go install

```bash
go install github.com/hjfitz/pspp/cmd/proxy@latest
```

## Usage

```bash
./proxy -u http://localhost:3000 -p 8080
```

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--upstream` | `-u` | Upstream URL to forward Pub/Sub push requests to | *required* |
| `--port` | `-p` | Port to listen on | `8080` |

### Example

1. Start your application (e.g. running on port 3000):

   ```bash
   # Your app receives Pub/Sub push format
   pnpm start:dev  # listens on :3000
   ```

2. Start PSPP:

   ```bash
   ./proxy -u http://localhost:3000 -p 8080
   ```

3. Send a raw HTTP POST to PSPP:

   ```bash
   curl -X POST http://localhost:8080/your-push-endpoint \
     -H "Content-Type: application/json" \
     -d '{"event": "order.created", "orderId": "123"}'
   ```

PSPP wraps the body in the Pub/Sub envelope and forwards it to `http://localhost:3000/your-push-endpoint`. Your app receives the same format it would get from GCP Pub/Sub push subscriptions.

## License

MIT
