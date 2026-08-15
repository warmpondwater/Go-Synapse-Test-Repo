# Go-Synapse-Test-Repo

Official multi-language testbed repository for **[Go-Synapse](https://warmpondwater.com)** — the local-first, interactive 2D code intelligence, AST mapping, and security verification engine.

---

## Overview

This repository is designed for developers, architects, and security auditors to test and evaluate Go-Synapse across **11 core programming languages** with pre-configured patterns for:
* **Polyglot Call Graph Mapping**: Go, Python, TypeScript, JavaScript, Rust, C, C++, C#, Java, PHP, and Ruby.
* **Static Taint Tracing**: Source-to-sink vulnerability paths (`SOURCE` $\rightarrow$ `SANITIZER` $\rightarrow$ `SINK`).
* **Dead-Code Quarantine**: 0 in-degree isolated functions automatically grouped into quarantine boundaries.
* **Concurrency & Loop-Scoped Defers**: Heuristic AST detection for defer-in-loops and channel patterns.
* **Hub Portal Spaghetti Prevention**: High-fanout helper methods that collapse into portal badges (`📞`).

---

## Directory Structure & Language Matrix

| Directory | Language | Demonstration Pattern |
| :--- | :--- | :--- |
| `go_src/` | **Go** | Concurrency channels, dataflow taint tracing, loop-scoped defers, dead code blocks |
| `py_src/` | **Python** | Flask routing, request handling, subprocess execution |
| `ts_src/` | **TypeScript** | Typed API controllers, interfaces, and middleware |
| `js_src/` | **JavaScript** | Asynchronous network client and DOM event handlers |
| `rust_src/` | **Rust** | Pattern matching, memory safety models, file I/O |
| `c_src/` | **C** | Non-local jumps (`setjmp`/`longjmp`), raw pointer operations |
| `cpp_src/` | **C++** | Object hierarchies, template abstractions |
| `cs_src/` | **C#** | LINQ queries, asynchronous tasks, dependency injection |
| `java_src/` | **Java** | Object-oriented configuration loader and factory patterns |
| `php_src/` | **PHP** | Dynamic parameter parsing and SQL statement builders |
| `ruby_src/` | **Ruby** | Metaprogramming handlers and block yield dispatchers |

---

## How to Test with Go-Synapse

### 1. Clone this Repository
```bash
git clone https://github.com/warmpondwater/Go-Synapse-Test-Repo.git
cd Go-Synapse-Test-Repo
```

### 2. Launch Go-Synapse Visualizer
```bash
# Start local visual intelligence server
./Go-Synapse -dir . -port 8080
```
Open `http://127.0.0.1:8080` to interact with the 2D spatial canvas.

### 3. Run Cryptographic Audit Pass
```bash
# Verify 100% AST integrity and sign local audit certificate
./Go-Synapse -audit=verify -dir .
```

### 4. Run Offline Software Composition Analysis (SCA)
```bash
# Scan dependency manifests without sending network packets
./Go-Synapse -audit=sca -offline -dir .
```

---

## License

This test repository is provided free and open for public evaluation under the [MIT License](LICENSE).
For Go-Synapse visualizer licensing and documentation, visit **[warmpondwater.com](https://warmpondwater.com)**.
