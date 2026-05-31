<div align="center">
  <h1> QUBEX Sentinel: AI-Driven Post-Quantum DePIN Shield</h1>
  <p><b>The First Decentralized AI-PQC Middleware for EVM Ecosystems & Cross-Chain Bridges</b></p>
  
  [![Status](https://img.shields.io/badge/Status-MVP%20%7C%20Genesis%20Pilots-00f5ff?style=flat-square)](#)
  [![Security](https://img.shields.io/badge/Security-NIST_Level_5-blue?style=flat-square)](#)
  [![L1 Impact](https://img.shields.io/badge/L1_Impact-Zero-success?style=flat-square)](#)
  [![Verification](https://img.shields.io/badge/Verification_Latency-~0.17ms-critical?style=flat-square)](#)
</div>

---

## The "Quantum Collapse" & HNDL Threat
Legacy ECDSA rollups face an existential threat from quantum vectors. State actors and malicious entities are already executing "Harvest Now, Decrypt Later" (HNDL) attacks—archiving encrypted L1/L2 state data to decrypt retroactively once Cryptographically Relevant Quantum Computers (CRQCs) come online. The choice is binary: natively proof your infrastructure, or leave your ecosystem vulnerable to total TVL extraction.

---

## ARCHITECTURE & OPEN-SOURCE DISCLOSURE

What this repository is: This repository contains the Cryptographic Benchmark Engine (Proof-of-Concept) used by QUBEX Sentinel. It demonstrates the sub-millisecond local execution of ML-DSA signatures (Sign & Verify) and verifies public RPC state configurations. It is designed for auditors and developers to independently verify our cryptographic latency claims.

What this repository is NOT: This is NOT the complete Qubex Pre-Batcher / Interceptor middleware. The core interceptor logic, mempool state-handling, and zero-latency sequencer-bypassing architecture are currently Closed-Source / Proprietary Infrastructure. 

*We are open-sourcing the cryptographic benchmarks to prove EVM viability. Full architectural integration will be available selectively to L2 Core Teams under NDA/partnership agreements.*

---

## Core Architecture: The "Zero L1 Impact" Paradigm

QUBEX Sentinel operates as a non-invasive interceptor layer between the Sequencer and the Batcher, securing L2s against "Harvest Now, Decrypt Later" threats without modifying existing consensus logic.

Integrating PQC directly on-chain forces unsustainable gas inflation and state bloat. QUBEX solves this via a decoupled pre-batcher execution:

1. AI Threat Detection (Chaos Engine): The middleware intercepts transactions, passing them through our predictive AI engine to detect anomalous "Harvest Now, Decrypt Later" sequencing patterns in real-time.

2. Decentralized Verification (Sentinel Nodes): Clean traffic is routed to our decentralized node network, where NIST-standard ML-DSA signatures are mathematically validated in ~0.17ms.

3. Compress: Only cryptographically verified, quantum-resistant payloads reach the op-batcher.

Result: Absolute PQC integrity on L2 with zero bytes of overhead added to L1 calldata. We enable post-quantum security at the speed of legacy ECDSA

---

## Live Devnet Omnichain Benchmark (Latest Audit Update)
*Executed via QUBEX Omnichain Broadcaster on commercial-grade hardware. Architecture is fully upgraded to support extreme parameter sets.*

Load Test Parameters:
* Security Standard: NIST Level 5 (ML-DSA-87)
* Operations: 2.2 Million Concurrent PQC Operations (1.1M Sign + 1.1M Verify)
* Concurrency: 11 Tier-1 EVM Ecosystems Simultaneously
* Total Execution Time: ~3m 10s

| Target Ecosystem | Signing Latency | Verification Latency | Status |
| :--- | :--- | :--- | :--- |
| Arbitrum | 476,748 ns (~0.47 ms) | 179,985 ns (~0.17 ms) | SECURED |
| zkSync | 532,161 ns (~0.53 ms) | 186,484 ns (~0.18 ms) | SECURED |
| Metis | 569,466 ns (~0.56 ms) | 175,481 ns (~0.17 ms) | SECURED |
| BNB | 752,569 ns (~0.75 ms) | 217,239 ns (~0.21 ms) | SECURED |
| Base | 1,023,278 ns (~1.02 ms) | 118,730 ns (~0.11 ms) | SECURED |
| Optimism | 1,038,430 ns (~1.03 ms) | 120,701 ns (~0.12 ms) | SECURED |
| Blast | 1,057,001 ns (~1.05 ms) | 134,870 ns (~0.13 ms) | SECURED |
| Polygon | 1,401,724 ns (~1.40 ms) | 74,529 ns (~0.07 ms) | SECURED |
| Mantle | 1,395,950 ns (~1.39 ms) | 78,393 ns (~0.07 ms) | SECURED |
| Linea | 1,709,490 ns (~1.70 ms) | 70,578 ns (~0.07 ms) | SECURED |
| Scroll | 1,813,690 ns (~1.81 ms) | 58,989 ns (~0.05 ms) | SECURED |

> Hardware Scalability Note: The above logs reflect extreme stress-testing under *consumer-grade CPU bottlenecks*. The QUBEX decoupled algorithm inherently solves PQC overhead, engineered to scale at < 100,000 ns latency per operation when deployed on Enterprise Bare Metal architectures (e.g., AWS Graviton, 64+ core instances).

---

## The QUBEX Trust Index
Integrating QUBEX Sentinel immediately enrolls your network into the QUBEX Trust Index—the emerging industry standard for quantum-resilient liquidity. Chains with high Trust Index scores provide institutional capital the cryptographic guarantee required for the next decade of Web3 operations.

---

## Omnichain Reproduction & Local Verification
*"Don't trust, verify."* QUBEX Sentinel operates with absolute cryptographic transparency. Enterprise architects and protocol core developers can independently reproduce the 2.2M operations stress test locally.

Infrastructure Prerequisites:
* Go Environment: Version 1.20 or higher.
* Hardware: Multi-core CPU. *(Warning: The benchmark is designed to force 100% CPU thread exhaustion to simulate peak enterprise sequencer loads. Proceed with caution on commercial laptops).*

Execution Protocol:
Run the following commands in your terminal to initialize the omnichain broadcaster:

### 1. Clone the infrastructure repository

```go
git clone https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git
```

```go
cd Qubex-PQC-Benchmarks
```

### 2. Synchronize cryptographic dependencies (Cloudflare CIRCL ML-DSA)

```go
go mod tidy
```

### 3. Ignite the Devnet Broadcaster

```go
go run main.go
```

## Integration & Genesis Pilots

QUBEX Sentinel is currently in the MVP Phase, actively selecting high-TVL Cross-Chain Bridges, Oracles, and Tier-1 L2 Rollups for exclusive Genesis Pilot integrations

Networks participating in the Genesis Pilot phase are eligible for Strategic Integration Bounties (subsidized from our Ecosystem Grants allocation) to eliminate the financial friction of transitioning to a quantum-safe architecture.

### Initiate Integration:

To request a technical handoff, review the Trust Index framework, and secure a Genesis Pilot slot, submit your protocol's architecture for review by our core team.

Institutional Contact: spyridongagr@qubexsentinel.com

X (Twitter): @QUBEX_SENTINEL

[Submit Genesis Pilot Request Here](https://forms.gle/hmUdBiQz3PT2x8TT7)
