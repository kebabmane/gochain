package pow

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/kebabmane/gochain/types"
)

var difficulty int

type Chain interface {
	AddBlock(block *types.Block) error
	CurrentBlock() (*types.Block, error)
}

type TxPool interface {
	GetTransactions() []*types.Transaction
}

type PowServer struct {
	mu    sync.Mutex
	chain Chain
	pool  TxPool
	start bool
	quit  chan interface{}
}

func NewPowServer(chain Chain, pool TxPool) *PowServer {
	return &PowServer{
		chain: chain,
		pool:  pool,
		start: false,
		quit:  make(chan interface{}, 1),
	}
}

func (this *PowServer) Start() {
	if this.start {
		return
	}

	go this.Mining()
	this.start = true
}

func (this *PowServer) Stop() {
	if !this.start {
		return
	}
	this.quit <- true
	this.start = false
}

func (this *PowServer) Mining() {
out:
	for {
		select {
		case <-this.quit:
			break out
		default:
		}
		difficulty = 1
		parentBlock, _ := this.chain.CurrentBlock()
		b := types.NewBlock(parentBlock, nil)
		b.Difficulty = difficulty
		this.solveDifficulty(b, difficulty)
		// Validation at this point checks for the correct height and parent hash matches what is stored
		if b.Validation(parentBlock) != nil {
			fmt.Println("it be broke, validation failed")
		} else {
			fmt.Println("validation was successful, storing block")
			this.chain.AddBlock(b)
		}

		fmt.Println("Current Height: ", b.GetHeight())
		//fmt.Println("The Block that has been stored: ", b)
		time.Sleep(5 * time.Second) //TODO

	}
}

func validateBlock(bhash *types.Block, difficulty int) bool {
	fmt.Println("validateBlock in PoW")
	hash, _ := bhash.Hash()
	prefix := strings.Repeat("0", difficulty)
	//fmt.Println("difficulty: ", difficulty)
	//fmt.Println("prefix: ", prefix)
	fmt.Println("hash: ", hash.ToHexString())
	return strings.HasPrefix(hash.ToHexString(), prefix)
}

func (this *PowServer) solveDifficulty(b *types.Block, difficulty int) *types.Block {
	for i := uint64(0); ; i++ {
		b.ChangeNonce(i)
		if !validateBlock(b, difficulty) {
			fmt.Println("hash no good, do more work!")
			continue
		} else {
			fmt.Println("hash good continue")
			break
		}
	}

	return b
}
