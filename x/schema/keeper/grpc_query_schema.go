package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (k Keeper) Schema(goCtx context.Context, req *st.QuerySchemaRequest) (*st.QuerySchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ref, found := k.GetWhatIsFromCreator(ctx, req.Creator)
	if !found || len(ref) < 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Schema was not found")
	}

	var what_is *st.WhatIs
	for _, item := range ref {
		if item.Did == req.Did {
			what_is = item
			break
		}
	}

	if what_is == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "Schema was not found for id: %s", req.Did)
	}

	return &st.QuerySchemaResponse{
		Schema: what_is.Schema,
	}, nil
}
