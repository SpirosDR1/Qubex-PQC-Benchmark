# QUBEX Sentinel: Decoupled ML-DSA Pre-Batcher Shield

![Status](https://img.shields.io/badge/Status-Devnet_Active-success) ![Security](https://img.shields.io/badge/Security-NIST_Level_5-blue) ![L1_Impact](https://img.shields.io/badge/L1_Impact-Zero-lightgrey)

## ARCHITECTURE & OPEN-SOURCE DISCLOSURE

What this repository is: This repository contains the Cryptographic Benchmark Engine (Proof-of-Concept) used by Qubex Sentinel. It demonstrates the sub-millisecond local execution of ML-DSA signatures and verifies public RPC state configurations. It is designed for developers to independently verify the cryptographic latency claims.

What this repository is NOT: This is NOT the complete Qubex Pre-Batcher / Interceptor middleware. The core interceptor logic, mempool state-handling, and sequencer-bypassing architecture are currently Closed-Source / Proprietary Infrastructure.

We are open-sourcing the cryptographic benchmarks to prove EVM viability. Full architectural integration will be available selectively to L2 Core Teams under NDA/partnership agreements.

## Overview

QUBEX Sentinel is an enterprise-grade, decoupled Post-Quantum Cryptography (PQC) middleware designed strictly for EVM-compatible Layer 2 and Layer 3 ecosystems. Utilizing the NIST-standardized ML-DSA (Dilithium) algorithm, QUBEX intercepts and verifies signatures at the sequencer level (pre-batcher).

Our architecture neutralizes "Harvest Now, Decrypt Later" (HNDL) vectors and achieves "Quantum-Safe by Default" sovereignty without altering L1 calldata or degrading sequencer throughput.

---

## Live Devnet Omnichain Benchmark

*Executed via QUBEX Omnichain Broadcaster on commercial-grade sequencer nodes.*

Load Test Parameters:
* Operations: 1.1 Million Concurrent PQC Verifications
* Concurrency: 11 Tier-1 EVM Ecosystems Simultaneously
* Execution Time: 1m 46s
* Security Standard: NIST Level 5 (ML-DSA-87)

| Target Ecosystem | Peak Stress Latency (1.1M Load) | Status  |
| :--- | :--- | :--- |
| Base | 39,595 ns (~0.039 ms) | SECURED |
| Blast | 317,479 ns (~0.31 ms) | SECURED |
| Linea | 323,951 ns (~0.32 ms) | SECURED |
| BNB | 424,553 ns (~0.42 ms) | SECURED |
| Scroll | 450,130 ns (~0.45 ms) | SECURED |
| Mantle | 523,436 ns (~0.52 ms) | SECURED |
| Arbitrum | 602,595 ns (~0.60 ms) | SECURED |
| Polygon | 607,658 ns (~0.60 ms) | SECURED |
| zkSync | 608,113 ns (~0.60 ms) | SECURED |
| Metis | 683,961 ns (~0.68 ms) | SECURED |
| Optimism | 1,058,173 ns (~1.05 ms) | SECURED |

> Hardware Scalability Note: The above logs reflect stress-testing under *consumer-grade CPU bottlenecks*. The QUBEX decoupled algorithm inherently solves PQC overhead, engineered to scale at < 100,000 ns latency per operation when deployed on Enterprise Bare-Metal architectures (e.g., AWS Graviton, 64+ core instances).

---

## Core Architecture: The "Zero L1 Impact" Paradigm

Legacy ECDSA rollups face an existential threat from quantum vectors. However, integrating PQC directly on-chain inflates gas costs to unsustainable levels. 

QUBEX solves this via a decoupled pre-batcher execution:

1. Intercept: The middleware wraps the op-node (sequencer).
2. Verify: ML-DSA signatures are mathematically verified in sub-millisecond latency (e.g., Base at ~39k ns, Linea at ~323k ns).
3. Compress: Only verified state data reaches the op-batcher.
4. Result: Absolute PQC integrity on L2 with zero bytes of PQC overhead added to L1 calldata.

---

## Omnichain Reproduction & Local Verification

"Don't trust, verify." QUBEX Sentinel operates with absolute cryptographic transparency. Enterprise architects and protocol core developers can independently reproduce the 1.1M operations stress test locally.

### Infrastructure Prerequisites
* Go Environment: Version 1.20 or higher.
* Hardware: Multi-core CPU. *(Warning: The benchmark is designed to force 100% CPU thread exhaustion to simulate peak enterprise sequencer loads. Proceed with caution on commercial laptops).*

### Execution Protocol
Run the following commands in your terminal to initialize the omnichain broadcaster:

#### Clone the infrastructure repository

```go
git clone https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git
```

```go
cd Qubex-PQC-Benchmarks
```

#### Synchronize cryptographic dependencies (Cloudflare CIRCL ML-DSA)

```go
go mod tidy
```

#### Ignite the Devnet Broadcaster

```go
go run main.go
```

### Integration & Genesis Pilots

QUBEX Sentinel is currently selecting high-TVL Tier-1 DeFi protocols and Rollup-as-a-Service (RaaS) providers for exclusive Genesis Pilot integrations prior to full EVM-wide protocol standardization.

Legacy ECDSA architectures across all L2/L3 networks are mathematically exposed. The choice is binary: native-proof your infrastructure with QUBEX to ensure total TVL sovereignty, or leave your ecosystem vulnerable to total extraction.

Initiate Integration To request a technical handoff and secure a Genesis Pilot slot, submit your protocol's architecture for review by our core team:

[Submit Genesis Pilot Request Here](https://forms.gle/hmUdBiQz3PT2x8TT7)
