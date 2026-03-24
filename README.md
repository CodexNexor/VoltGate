<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=waving&height=220&color=0:030712,45:0B3B82,100:06B6D4&text=VoltGate&fontColor=E6F7FF&fontSize=54&fontAlignY=38&animation=fadeIn&desc=Multi-account%20AI%20gateway%20for%20local%20and%20remote%20use&descAlignY=60&descSize=16" alt="VoltGate banner" />

  <a href="#quick-start">
    <img src="https://readme-typing-svg.demolab.com?font=JetBrains+Mono&weight=700&size=21&duration=2600&pause=900&color=22D3EE&center=true&vCenter=true&width=900&lines=Route+Claude%2C+Gemini%2C+Codex%2C+and+Antigravity+through+one+gateway;Use+local+chat%2C+remote+chat%2C+or+the+OpenAI-compatible+API;Multi-account+failover+with+a+clean+local+control+panel" alt="VoltGate typing animation" />
  </a>

  <p>
    <img src="https://img.shields.io/badge/Go-1.26+-0ea5e9?style=for-the-badge&logo=go&logoColor=white" alt="Go 1.26+" />
    <img src="https://img.shields.io/badge/Next.js-16-111827?style=for-the-badge&logo=nextdotjs&logoColor=white" alt="Next.js 16" />
    <img src="https://img.shields.io/badge/OAuth-Claude%20%7C%20Gemini%20%7C%20Codex-0f172a?style=for-the-badge&logo=icloud&logoColor=67e8f9" alt="OAuth providers" />
    <img src="https://img.shields.io/badge/Remote-Chat%20%2B%20API-0891b2?style=for-the-badge&logo=cloudflare&logoColor=white" alt="Remote chat and API" />
    <img src="https://img.shields.io/badge/License-MIT-0f172a?style=for-the-badge" alt="MIT license" />
  </p>
</div>

## What Is VoltGate?

VoltGate is a local-first AI gateway that lets you connect multiple OAuth accounts, route requests across them, and use the same setup through:

- a clean local control panel
- a local chat workspace
- a remote chat link
- an OpenAI-compatible API endpoint

It is built for people who want one place to manage Claude, Gemini, Codex, and Antigravity access across different tools and machines.

## Why VoltGate

- Multi-account failover: if one account or model cools down, VoltGate can move to another ready account.
- Local + remote access: manage everything locally, then use remote chat or the remote API from another network.
- OpenAI-compatible API: point existing tools and scripts at VoltGate without rebuilding your workflow.
- Real account routing: use your connected Claude and Gemini subscriptions without buying separate API credits for every project.
- Simple operator flow: connect accounts locally, create one client API key, and start using chat or the API.

## Quick Start

### 1. Clone the repo

```bash
git clone https://github.com/CodexNexor/VoltGate
cd VoltGate
```

### 2. Start everything with one command

```bash
python run-stack.py
```

That command will:

- install missing runtime dependencies when possible
- start the Go backend
- build and start the UI
- start Cloudflare quick tunnels
- print your local and remote links

### 3. Open VoltGate locally

Open:

- `http://127.0.0.1:3000`

Then:

1. if you want Gemini or Antigravity login, set the Google OAuth environment variables from `.env.example`
2. connect your OAuth accounts locally
3. create a client API key
4. use local chat, remote chat, or the API

## How To Use It

### Local control panel

Use the local UI to:

- connect Claude / Gemini / Codex / Antigravity accounts
- view runtime account state
- create client API keys
- manage your local routing setup

### Local chat

Use:

- `http://127.0.0.1:3000/chat`

Paste a client API key, sync models, and chat directly through VoltGate.

### Remote chat

After startup, VoltGate prints a **Remote Chat** link.

Use that link from another computer, paste your client API key, and chat through the same local account pool.

### Remote API

After startup, VoltGate prints:

- `Remote API Base`
- `Remote Gemini Base`

Use those URLs from any computer or network with the client API key you created locally.

## Example API Call

```bash
curl -X POST "<REMOTE_API_BASE>/chat/completions" \
  -H "Authorization: Bearer <YOUR_CLIENT_API_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-2.5-flash",
    "messages": [
      { "role": "user", "content": "Reply with exactly: API is working" }
    ]
  }'
```

## Project Structure

```text
cmd/server             Go server entrypoint
internal/              core backend runtime and API handlers
sdk/                   routing, auth orchestration, executor logic
ui/                    VoltGate web UI
run-stack.py           one-command cross-platform launcher
run-local-stack.ps1    Windows launcher
auth/                  local auth storage (kept out of git)
.local-stack/          local runtime logs and process state (kept out of git)
```

## Security Model

- OAuth account connection and removal should be done locally.
- Remote chat and remote API use client API keys.
- Local auth files, logs, and runtime state are excluded from git.
- Remote account-management writes are restricted so the local machine stays the source of truth.

## Minimal Run Commands

Start:

```bash
python run-stack.py
```

Stop:

```bash
python run-stack.py --stop
```

## License

MIT
