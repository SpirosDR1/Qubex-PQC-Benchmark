# QUBEX SENTINEL | PQC Performance Benchmarks (Chaos Engine)

A high-performance audit tool designed to measure the latency impact of Post-Quantum Cryptography (PQC) on EVM-compatible networks. This utility implements NIST ML-DSA-44 (FIPS 204) signing and utilizes on-chain verification to ensure transparency in performance metrics.

### Live Portals & Genesis Program
* Official Portal: [qubexsentinel.com](https://qubexsentinel.com)
* Apply for Genesis Partnership (Zero-Fee Integration): [Access the Application Form Here](https://forms.gle/hmUdBiQz3PT2x8TT7)

---

## Why QUBEX SENTINEL?
As blockchain infrastructure transitions toward quantum resistance to combat HNDL (Harvest Now, Decrypt Later) attacks, understanding cryptographic overhead is critical. The Chaos Engine provides a deterministic way to benchmark signature latency, ensuring that PQC integration introduces virtually zero drag to high-frequency sequencer performance.

## Multi-Chain Performance Matrix (NIST ML-DSA)
*Verified metrics interacting directly with respective network RPCs (May 2026).*

| Network | Type | Avg Signing Latency (ns) | Verification | Status |
| :--- | :--- | :--- | :--- | :--- |
| Mantle | Optimistic Rollup | 39,032 ns | On-chain | Ready |
| Metis | Optimistic Rollup | 39,499 ns | On-chain | Ready |
| Polygon (Amoy)| zkEVM / AggLayer | 39,616 ns | On-chain | Ready |
| Optimism | Optimistic Rollup | 39,712 ns | On-chain | Ready |
| Base | Optimistic Rollup | 68,244 ns | On-chain | Ready |
| BNB Smart Chain| L1 / Sidechain | 68,630 ns | On-chain | Ready |
| Linea | zkEVM | 93,278 ns | On-chain | Ready |
| Arbitrum | Optimistic Rollup | 122,314 ns | On-chain | Ready |
| Scroll | zkEVM | 122,538 ns | On-chain | Ready |
| zkSync | ZK Rollup | 149,699 ns | On-chain | Ready |
| Blast | Optimistic Rollup | 230,952 ns | On-chain | Ready |

## Quick Start (Verify it yourself)
Prerequisites: Ensure you have Go installed on your machine.

1. Clone the repository:
```bash
git clone [https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git](https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git)
cd Qubex-PQC-Benchmarks
go mod tidy
```
2. Run the Chaos Engine Stress Test:

Supported networks: base, polygon, arbitrum, optimism, bnb, mantle, blast, zksync, linea, metis, scroll
go run main.go polygon

Expected Terminal Output & Audit Trail
The tool outputs real-time signing latency and verifies connectivity to the specified network.

--- QUBEX SENTINEL | MULTI-CHAIN PQC BENCHMARK ---
[SYSTEM] Targeting Network: polygon
[SYSTEM] Starting PQC Stress Test: 100000 iterations...

[AUDIT] Infrastructure: QUBEX SENTINEL ENGINE
[AUDIT] Target Network: polygon
[AUDIT] Avg PQC Signing Latency: 39616 ns
[AUDIT] Signature Integrity: true (NIST ML-DSA Verified)
[AUDIT] Verified Block Height: 38887690
[AUDIT] Status: ALL CHECKS PASSED
--------------------------------------------------
QUBEX SENTINEL: polygon Ecosystem is Ready for PQC Deployment.

## Secure Audit Logging

Each run automatically appends verifiable data to a network-specific log file (e.g., qubex_polygon_audit.log). The audit trail provides an immutable-like record including:

Execution Timestamp (RFC3339)
Targeted Network
Average PQC Latency (ns)
Verified Blockchain Height
Signature Verification Status
Ready for Production?
We are actively selecting Tier-1 RaaS providers and L2/L3 ecosystems for the Genesis Partnership Program (6-month zero-fee integration).
Initiate Integration via the Genesis Form
Developed for QUBEX SENTINEL Research Initiative | License: MIT
