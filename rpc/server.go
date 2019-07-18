package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kebabmane/gochain/types"
)

type Chain interface {
	AddBlock(block *types.Block) error
	CurrentBlock() (*types.Block, error)
	GetBlockByHeight(height uint64) (*types.Block, error)
}

type Server struct {
	chain Chain
}

func NewServer(chain Chain) *Server {
	return &Server{
		chain: chain,
	}
}

func (this *Server) Start() {
	http.HandleFunc("/", this.GetHandler)
	http.HandleFunc("/block", this.GetBlockByHeight)
	http.ListenAndServe(":8000", nil)
}

func (this *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := this.chain.CurrentBlock()
	hash, _ := b.Hash()
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, _ := json.Marshal(map[string]interface{}{
		"jsonrpc":       "2.0",
		"CurrentHash":   hash.ToHexString(),
		"TxRootHash":    (b.TxRootHash).ToHexString(),
		"CurrentHeight": b.Height,
		"ChainID":       b.ChainID,
		"ParentHash":    (b.ParentHash).ToHexString(),
		"Timestamp":     b.Timestamp,
		"Difficulty":    b.Difficulty,
		"Nonce":         b.Nonce,
	})

	w.Write(data)
}

func (this *Server) GetBlockByHeight(w http.ResponseWriter, r *http.Request) {
	keys, _ := r.URL.Query()["height"]
	fmt.Println("keys: ", keys)
	height, _ := strconv.Atoi(keys[0])
	b, _ := this.chain.GetBlockByHeight(uint64(height))
	hash, _ := b.Hash()
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, _ := json.Marshal(map[string]interface{}{
		"jsonrpc":    "2.0",
		"Hash":       hash.ToHexString(),
		"TxRootHash": (b.TxRootHash).ToHexString(),
		"Height":     b.Height,
		"ChainID":    b.ChainID,
		"ParentHash": (b.ParentHash).ToHexString(),
		"Timestamp":  b.Timestamp,
		"Difficulty": b.Difficulty,
		"Nonce":      b.Nonce,
	})

	w.Write(data)
}
