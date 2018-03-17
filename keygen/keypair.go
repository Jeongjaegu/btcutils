package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	
)

const (
	ALPHABET = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

var (
	bigRadix = big.NewInt(58)
	bigZero  = big.NewInt(0)
)

func main() {
	

	pub, priv, err := NewKey()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("pubk:", hex.EncodeToString(pub))
	fmt.Println("priv:", hex.EncodeToString(priv))
	fmt.Println("pubk len:", len(pub))
	fmt.Println("priv len:", len(priv))
}

func NewKey() ([]byte, []byte, error) {
	priv, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	pubkey := btcec.PublicKey(priv.PublicKey)
	pubkeyaddr := &pubkey
	return pubkeyaddr.SerializeCompressed(), priv.D.Bytes(), nil
}
