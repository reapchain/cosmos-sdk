package tmservice

import (
	"context"

	ctypes "github.com/reapchain/reapchain-core/rpc/core/types"

	"github.com/reapchain/cosmos-sdk/client"
)

func getBlock(ctx context.Context, clientCtx client.Context, height *int64) (*ctypes.ResultBlock, error) {
	// get the node
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	return node.Block(ctx, height)
}
