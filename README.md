# QUBEX Sentinel: Decoupled ML-DSA Pre-Batcher Shield

![Status: Active](https://img.shields.io/badge/Status-Devnet_Active-success) ![Security: ML-DSA](https://img.shields.io/badge/Security-NIST_Level_5-blue) ![Impact: Zero L1 Overhead](https://img.shields.io/badge/L1_Impact-Zero-brightgreen)

## Overview
QUBEX Sentinel is an enterprise-grade, decoupled Post-Quantum Cryptography (PQC) middleware designed strictly for EVM-compatible Layer 2 and Layer 3 ecosystems. Utilizing the NIST-standardized ML-DSA (Dilithium) algorithm, QUBEX intercepts and verifies signatures at the sequencer level (pre-batcher).

Our architecture neutralizes "Harvest Now, Decrypt Later" (HNDL) vectors and achieves "Quantum-Safe by Default" sovereignty without altering L1 calldata or degrading sequencer throughput.

---

## Live Devnet Omnichain Benchmark
*Executed via QUBEX Omnichain Broadcaster on commercial-grade sequencer nodes.*

Load Test Parameters:
- Operations: 1.1 Million Concurrent PQC Verifications
- Concurrency: 11 Tier-1 EVM Ecosystems Simultaneously
- Execution Time: 1m 46s
- Security Standard: NIST Level 5 (ML-DSA-87)

| Target Ecosystem | Baseline Latency (Enterprise CPU) | Peak Stress Latency (1.1M Load) | Status |
| :--- | :--- | :--- | :--- |
| Base (Sepolia) | ~39k - 65k ns | Sub-Millisecond (<1ms) | SECURED |
| Optimism | ~39k - 65k ns | Sub-Millisecond (<1ms) | SECURED |
| Arbitrum | ~39k - 65k ns | Sub-Millisecond (<1ms) | SECURED |
| zkSync Era | ~39k - 65k ns | Sub-Millisecond (<1ms) | SECURED |
| Polygon Amoy | ~39k - 65k ns | Sub-Millisecond (<1ms) | SECURED |

*(Full 11-network concurrency logs including Linea, Blast, Mantle, Scroll, Metis, BNB are preserved in system audits).*

---

## Core Architecture: The "Zero L1 Impact" Paradigm
Legacy ECDSA rollups face an existential threat from quantum vectors. However, integrating PQC directly on-chain inflates gas costs to unsustainable levels. 

QUBEX solves this via a decoupled pre-batcher execution:
1. Intercept: The middleware wraps the op-node (sequencer).
2. Verify: ML-DSA signatures are mathematically verified in sub-millisecond latency (avg. ~39k ns).
3. Compress: Only verified state data reaches the op-batcher.
4. Result: Absolute PQC integrity on L2 with zero bytes of PQC overhead added to L1 calldata.

---

## Omnichain Reproduction & Local Verification
"Don't trust, verify." QUBEX Sentinel operates with absolute cryptographic transparency. Enterprise architects and protocol core developers can independently reproduce the 1.1M operations stress test locally.

### Infrastructure Prerequisites
- Go Environment: Version 1.20 or higher.
- Hardware: Multi-core CPU. *(Warning: The benchmark is designed to force 100% CPU thread exhaustion to simulate peak enterprise sequencer loads. Proceed with caution on commercial laptops).*

### Execution Protocol
Run the following commands in your terminal to initialize the omnichain broadcaster:


#### 1. Clone the infrastructure repository

```go
git clone https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git
```

```go
cd Qubex-PQC-Benchmarks
```

#### 2. Synchronize cryptographic dependencies (Cloudflare CIRCL ML-DSA)

```go
go mod tidy
```

#### 3. Ignite the Devnet Broadcaster

```go
go run main.go
```
## Integration & Genesis Pilots

QUBEX Sentinel is currently selecting high-TVL Tier-1 DeFi protocols and Rollup-as-a-Service (RaaS) providers for exclusive Genesis Pilot integrations prior to full EVM-wide protocol standardization.

Legacy ECDSA architectures across all L2/L3 networks are mathematically exposed. The choice is binary: native-proof your infrastructure with QUBEX to ensure total TVL sovereignty, or leave your ecosystem vulnerable to total extraction.

### Initiate Integration

To request a technical handoff and secure a Genesis Pilot slot, submit your protocol's architecture for review by our core team:

[Submit Genesis Pilot Request](https://docs.google.com/forms/d/e/1FAIpQLSc6HvUe5tBH6fJ-m0B8j3eBBl4vBaPPHN_XF9h0m9OrxGPRww/viewform)
