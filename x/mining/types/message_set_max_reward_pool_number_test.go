package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mezonhub/mezonhub/testutil/sample"
	"github.com/mezonhub/mezonhub/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestMsgSetMaxRewardPoolNumber_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgSetMaxRewardPoolNumber
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgSetMaxRewardPoolNumber{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgSetMaxRewardPoolNumber{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
