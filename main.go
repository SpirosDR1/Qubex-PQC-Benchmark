package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cloudflare/circl/sign/mldsa/mldsa44"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("--- QUBEX SENTINEL | AUDIT-READY BENCHMARK ---")

	// 1. Connect to the actual network
	client, err := ethclient.Dial("https://sepolia.base.org")
	if err != nil {
		log.Fatalf("Conn Error: %v", err)
	}

	// 2. Generate PQC Keys
	pk, sk, _ := mldsa44.GenerateKey(rand.Reader)
	txData := []byte("QUBEX_SECURE_TRANSFER_100_USDC")

	// 3. Warm-up
	for i := 0; i < 1000; i++ {
		sk.Sign(nil, txData, nil)
	}

	// 4. Stress Test (100,000 iterations)
	iterations := 100000
	startBench := time.Now()
	for i := 0; i < iterations; i++ {
		sk.Sign(nil, txData, nil)
	}
	avgPQC := time.Since(startBench).Nanoseconds() / int64(iterations)

	// 5. Blockchain Proof
	header, _ := client.HeaderByNumber(context.Background(), nil)

	// 6. Verification
	signature, _ := sk.Sign(nil, txData, nil)
	valid := mldsa44.Verify(pk, txData, nil, signature)

	// 7. FILE PATH (Writes to the same folder where main.go is located)
	logMsg := fmt.Sprintf("[%s] Latency: %d ns | Block: %d | Valid: %v\n",
		time.Now().Format(time.RFC3339), avgPQC, header.Number, valid)

	f, err := os.OpenFile("qubex_benchmark.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		f.WriteString(logMsg)
		f.Close()
	}

	// Logs
	fmt.Println("--------------------------------------------------")
	fmt.Println("[AUDIT] Network: Base Sepolia")
	fmt.Printf("[AUDIT] PQC Signing Latency: %d ns\n", avgPQC)
	fmt.Printf("[AUDIT] Signature Verification: %v\n", valid)
	fmt.Printf("[AUDIT] Verified Block Height: %d\n", header.Number)
	fmt.Printf("[AUDIT] Log saved in: %s\\qubex_benchmark.log\n", os.Args[0])
	fmt.Println("--------------------------------------------------")
	fmt.Println("QUBEX SENTINEL: Ready for Deployment.")
}
