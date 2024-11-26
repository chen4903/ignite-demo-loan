package types

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestLoan{}

func NewMsgRequestLoan(creator string, amount string, fee string, collateral string, deadline string) *MsgRequestLoan {
	return &MsgRequestLoan{
		Creator:    creator,
		Amount:     amount,
		Fee:        fee,
		Collateral: collateral,
		Deadline:   deadline,
	}
}

// ValidateBasic 对 MsgRequestLoan 消息进行基本的无状态验证。
// 它确保以下条件：
// 1. creator 地址是有效的 Bech32 地址。
// 2. amount 是有效的且非空的 Coins 对象。
// 3. fee 是有效的 Coins 对象。
// 4. deadline 是一个正整数。
// 5. collateral 是有效的且非空的 Coins 对象。
//
// 返回值：
// - 如果消息通过所有验证，返回 nil。
// - 如果任何验证失败，返回对应的错误。
func (msg *MsgRequestLoan) ValidateBasic() error {
    // 验证 creator 地址是否是有效的 Bech32 地址。
    _, err := sdk.AccAddressFromBech32(msg.Creator)
    if err != nil {
        return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
    }

    // 验证 amount 是否是有效的 Coins 对象。
    amount, _ := sdk.ParseCoinsNormalized(msg.Amount)
    if !amount.IsValid() {
        return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coins object")
    }
    if amount.Empty() {
        return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount is empty")
    }

    // 验证 fee 是否是有效的 Coins 对象。
    fee, _ := sdk.ParseCoinsNormalized(msg.Fee)
    if !fee.IsValid() {
        return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "fee is not a valid Coins object")
    }

    // 验证 deadline 是否是一个正整数。
    deadline, err := strconv.ParseInt(msg.Deadline, 10, 64)
    if err != nil {
        return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "deadline is not an integer")
    }
    if deadline <= 0 {
        return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "deadline should be a positive integer")
    }

    // 验证 collateral 是否是有效的 Coins 对象。
    collateral, _ := sdk.ParseCoinsNormalized(msg.Collateral)
    if !collateral.IsValid() {
        return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collateral is not a valid Coins object")
    }
    if collateral.Empty() {
        return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collateral is empty")
    }

    return nil
}
