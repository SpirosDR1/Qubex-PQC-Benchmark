QUBEX SENTINEL | PQC Performance Benchmark

A high-performance audit tool designed to measure the latency impact of Post-Quantum Cryptography (PQC) on EVM-compatible networks. This utility implements ML-DSA-44 (FIPS 204) signing and utilizes on-chain verification to ensure transparency in performance metrics.

Why QUBEX SENTINEL?

As blockchain infrastructure transitions toward quantum resistance, understanding the cryptographic overhead is critical. This tool provides a deterministic way to benchmark signature latency, ensuring that PQC integration doesn't bottleneck high-frequency sequencer performance.

Technical Specs

Algorithm: ML-DSA-44 (FIPS 204 standard).

Environment: Live integration with Base Sepolia (Testnet).

Metrics:

Mean signing latency (in nanoseconds).

On-chain state verification (via block header synchronization).

Persistent audit logging for reproducibility.


Quick Start

Prerequisites:

Ensure you have Go installed on your machine.

'''Bash
go mod init qubex-benchmark
go mod tidy

Run the Benchmark:

'''Bash
go run main.go

Results:

The tool will output real-time signing latency and verify connectivity to the Base Sepolia network. Audit logs are automatically appended to qubex_benchmark.log in the local directory.

Audit Trail
The generated qubex_benchmark.log provides an immutable-like audit trail of your performance runs, including:

Execution Timestamp.

Average PQC Latency.

Verified Blockchain Height.

Signature Verification Status.

Developed for QUBEX SENTINEL Research Initiative.
License: MIT
