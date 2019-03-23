package models

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

// LEADING_ZEROS is the amount of leading zeros required for the proof of work
var LeadingZeros int64 = 20

// ProofOfWork contains values needed for the proof of work method. It has a
// target, which is a number with a specific amount of leading zeros, a block
// and a counter which is used in the generation of the hash. The block is
// hashed with the increasing counter until the hash is smaller than the target
// i.e. hash at least the specific amount of leading zeros
type ProofOfWork struct {
	target *big.Int
	block  *Block
}

// NewProofOfWork creates a proof of work construct for the block with the
// specified amount of leading zeros
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-LeadingZeros))
	return &ProofOfWork{
		target: target,
		block:  block,
	}
}

// createHash creates a hash for the block using SHA256. The hash consists of
// the block data, the previous hash, the
// unix timestamp and the counter.
func (pow ProofOfWork) createHash() []byte {
	counter := pow.block.ProofOfWorkCounter
	if pow.block.ProofOfWorkCounter == nil {
		counter = big.NewInt(0)
	}
	hashValue := bytes.Join(
		[][]byte{
			pow.block.HashTransactions(),
			pow.block.PreviousHash,
			[]byte(strconv.FormatInt(pow.block.Timestamp.Unix(), 10)),
			[]byte(strconv.FormatInt(LeadingZeros, 10)),
			counter.Bytes(),
		}, []byte{})
	hash := sha256.Sum256(hashValue)
	return hash[:]
}

// Run creates a hash with at least the specified amount of leading zeros. The
// counter is incremented on each iteration
func (pow *ProofOfWork) Run() (hash []byte) {
	var compareInt big.Int
	start := time.Now()
	for {
		hash = pow.createHash()
		if pow.target.Cmp(compareInt.SetBytes(hash)) != 1 {
			pow.block.ProofOfWorkCounter.Add(pow.block.ProofOfWorkCounter,
				big.NewInt(1))
		} else {
			break
		}
	}
	duration := time.Now().Sub(start)
	fmt.Printf("hash creation took %v seconds\n", duration.Seconds())
	return
}

// Validate checks if the hash of the block has at least the specified amount of
// leading zeros
func (pow ProofOfWork) Validate() bool {
	var compareInt big.Int
	return pow.target.Cmp(compareInt.SetBytes(pow.createHash())) == 1
}
