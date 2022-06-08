// Deprecated: 測試失敗，沒屌用的學習文章
package createConn

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client defines typed wrappers for the Ethereum RPC API.
// Deprecated: 測試失敗，沒屌用的學習文章
type Client struct {
	rpcClient *rpc.Client
	EthClient *ethclient.Client
}

// Connect creates a client that uses the given host.
// Deprecated: 測試失敗，沒屌用的學習文章
func Connect(host string) (*Client, error) {
	rpcClient, err := rpc.Dial(host)
	if err != nil {
		return nil, err
	}
	ethClient := ethclient.NewClient(rpcClient)
	return &Client{rpcClient, ethClient}, nil
}

// GetBlockNumber returns the block number.
// Deprecated: 測試失敗，沒屌用的學習文章
func (ec *Client) GetBlockNumber(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := ec.rpcClient.CallContext(ctx, &result, "eth_blockNumber")
	return (*big.Int)(&result), err
}
