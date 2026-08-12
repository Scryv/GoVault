# GoVault 

[![License: GPL v2](https://img.shields.io/badge/License-GPL_v2-blue.svg)](https://github.com/Scryv/GoVault/blob/main/LICENSE)

**GoVault** is a SelfHosted password manager offering both a terminal TUI and a lightweight web interface designed for privacy and for being lightweight

> *Your passwords stay between you and Gopher.*

<img src="https://sdgscryv.xyz/img/GoVault.png" width="300px" alt="GoVault Logo">

---

## Features and Structure

### Tech Stack
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) 
![Rust](https://img.shields.io/badge/rust-%23000000.svg?style=for-the-badge&logo=rust&logoColor=white) 
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white) 
![HTMX](https://img.shields.io/badge/HTMX-%233366CC.svg?style=for-the-badge&logo=htmx&logoColor=white)

* **Privacy Design:** Your primary secret stays hidden on your own device and is never exposed as readable text over the network.
* **Rust Cryptography Core:** authentication and encryption stuff is written in Rust for safety and speed.
* **Go Backend and Web Server:** Built on Golang standard `net/http` lib paired with **HTMX** templates for a light web UI.
* **Terminal UI:** Terminal interface built for getting quick access to your passwords offline without having to go to the web UI
* **Secure Storage:** Uses embedded **SQLite** for local vault management (stored under `~/.govault` or `/home/user/`), employing AES encryption for vault entries and salted hashing (migrating from SHA-512 to Argon2).

---

## Project Structure

```text
GoVault/
├── cmd/
│   └── main.go           • main files
├── internal/
│   ├── api/
│   │   ├── web/          • HTMX templates and static assets(pictures and css probs)
│   │   └── api.go        • HTTP handlers routes and most of the logic
│   ├── auth/             • Cryptographic operations and token auth
│   ├── database/         • SQLite connection and migrations
│   ├── user/             • User management and vault actions
│   └── utils/            • Shared helper functions
├── .gitignore
└── LICENSE
