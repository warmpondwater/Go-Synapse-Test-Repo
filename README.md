# Go-Synapse Demo Target

This is a mock polyglot application designed to demonstrate the AST analysis, security taint tracing, and verification capabilities of **Go-Synapse**.

## Structure

- `go_src/`: A Go service demonstrating:
  - **Taint Tracer Pathing**: A caller-callee sequence from `GetUserInput()` (SOURCE) -> `SanitizeInput()` (SANITIZER) -> `ExecuteDatabaseQuery()` (SINK).
  - **Isolated Quarantine**: Unused/dead code items in `dead_code.go` showing how Go-Synapse visually separates dead code blocks.
  - **Node Density Hazard**: High fan-out helper methods inside `generated_proto.go` illustrating auto-generated code clusters.
- `c_src/`: A C file demonstrating non-local jumps (`setjmp`/`longjmp`) and resource cleanups.
- `js_src/`: An asynchronous JavaScript Network Client.
- `rust_src/`: A Rust file matching file read results.
- `java_src/`: An object-oriented Java file reading system configuration.
