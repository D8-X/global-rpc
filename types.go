package globalrpc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type RpcConfig struct {
	ChainId int      `json:"chainId"`
	Wss     []string `json:"wss"`
	Https   []string `json:"https"`
}

type RPCKind int

const (
	TypeHTTPS RPCKind = iota // Starts at 0
	TypeWSS                  // 1
)

func (t RPCKind) String() string {
	switch t {
	case TypeHTTPS:
		return "HTTPS"
	case TypeWSS:
		return "WSS"
	default:
		return "Unknown"
	}
}

func loadRPCConfig(chainId int, filename string) (RpcConfig, error) {
	var config []RpcConfig
	jsonFile, err := os.Open(filename)
	if err != nil {
		return RpcConfig{}, err
	}
	defer jsonFile.Close()

	// Read the file's contents into a byte slice
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	// Unmarshal the JSON data
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return RpcConfig{}, fmt.Errorf("unmarshalling JSON failed: %v", err)
	}
	idx := -1
	for j := range config {
		if config[j].ChainId == chainId {
			idx = j
			break
		}
	}
	if idx == -1 {
		return RpcConfig{}, fmt.Errorf("no rpc config for chain %d", chainId)
	}

	return config[idx], nil
}
