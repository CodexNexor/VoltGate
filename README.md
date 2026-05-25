# VoltGate ⚡

A lightweight, high-performance API gateway for microservices with built-in rate limiting, authentication, and request transformation.

## Features
- **Rate Limiting** — Token bucket algorithm, configurable per route
- **Auth Middleware** — JWT, OAuth2, and API key support
- **Request/Response Transformation** — Modify headers, bodies, and params on the fly
- **Circuit Breaker** — Automatic failover for unhealthy services
- **Metrics Dashboard** — Real-time request monitoring via Prometheus
- **Hot Reload** — Update config without restarting

## Quick Start
```bash
git clone https://github.com/CodexNexor/VoltGate.git
cd VoltGate
pip install -r requirements.txt
python main.py
```

## Configuration
See `config.yaml` for all available options.

## License
MIT
