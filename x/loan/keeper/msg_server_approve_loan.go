package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"loan/x/loan/types"
)

// ApproveLoan 处理批准贷款的请求消息
func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	// 将上下文从 goCtx 转换为 SDK 的上下文
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 根据贷款 ID 获取对应的贷款信息
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}

	// 检查贷款状态是否为 "requested"
	if loan.State != "requested" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}

	// 获取放款人和借款人的地址
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)     // 放款人地址
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower) // 借款人地址

	// 解析贷款金额为 Coins 对象
	amount, err := sdk.ParseCoinsNormalized(loan.Amount)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrWrongLoanState, "Cannot parse coins in loan amount")
	}

	// 将贷款金额从放款人账户转移到借款人账户
	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, err
	}

	// 更新贷款信息：设置放款人和状态
	loan.Lender = msg.Creator // 设置放款人为消息的创建者
	loan.State = "approved"   // 更新贷款状态为 "approved"
	k.SetLoan(ctx, loan)      // 保存更新后的贷款信息

	return &types.MsgApproveLoanResponse{}, nil
}
