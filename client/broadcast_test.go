package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/reapchain/reapchain-core/crypto/tmhash"
	"github.com/reapchain/reapchain-core/mempool"
	"github.com/reapchain/reapchain-core/rpc/client/mock"
	ctypes "github.com/reapchain/reapchain-core/rpc/core/types"
	tmtypes "github.com/reapchain/reapchain-core/types"

	"github.com/reapchain/cosmos-sdk/client/flags"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

type MockClient struct {
	mock.Client
	err error
}

func (c MockClient) BroadcastTxCommit(ctx context.Context, tx tmtypes.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	return nil, c.err
}

func (c MockClient) BroadcastTxAsync(ctx context.Context, tx tmtypes.Tx) (*ctypes.ResultBroadcastTx, error) {
	return nil, c.err
}

func (c MockClient) BroadcastTxSync(ctx context.Context, tx tmtypes.Tx) (*ctypes.ResultBroadcastTx, error) {
	return nil, c.err
}

func CreateContextWithErrorAndMode(err error, mode string) Context {
	return Context{
		Client:        MockClient{err: err},
		BroadcastMode: mode,
	}
}

// Test the correct code is returned when
func TestBroadcastError(t *testing.T) {
	errors := map[error]uint32{
		mempool.ErrTxInCache:       sdkerrors.ErrTxInMempoolCache.ABCICode(),
		mempool.ErrTxTooLarge{}:    sdkerrors.ErrTxTooLarge.ABCICode(),
		mempool.ErrMempoolIsFull{}: sdkerrors.ErrMempoolIsFull.ABCICode(),
	}

	modes := []string{
		flags.BroadcastAsync,
		flags.BroadcastBlock,
		flags.BroadcastSync,
	}

	txBytes := []byte{0xA, 0xB}
	txHash := fmt.Sprintf("%X", tmhash.Sum(txBytes))

	for _, mode := range modes {
		for err, code := range errors {
			ctx := CreateContextWithErrorAndMode(err, mode)
			resp, returnedErr := ctx.BroadcastTx(txBytes)
			require.NoError(t, returnedErr)
			require.Equal(t, code, resp.Code)
			require.NotEmpty(t, resp.Codespace)
			require.Equal(t, txHash, resp.TxHash)
		}
	}

}
