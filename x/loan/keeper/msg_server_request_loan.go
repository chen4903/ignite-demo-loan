package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"loan/x/loan/types"
)

// RequestLoan 处理贷款请求消息
func (k msgServer) RequestLoan(goCtx context.Context, msg *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
	// 将上下文从 goCtx 转换为 SDK 的上下文
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 创建一个新的 Loan 实例，并初始化其字段
	var loan = types.Loan{
		Amount:     msg.Amount,     // 贷款金额
		Fee:        msg.Fee,        // 手续费
		Collateral: msg.Collateral, // 抵押品
		Deadline:   msg.Deadline,   // 截止日期
		State:      "requested",    // 初始状态为 "requested"
		Borrower:   msg.Creator,    // 借款人
	}

	// 验证并转换借款人的地址
	borrower, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	// 解析抵押品为 Coins 对象
	collateral, err := sdk.ParseCoinsNormalized(loan.Collateral)
	if err != nil {
		panic(err)
	}

	// 从借款人的账户中转移抵押品到模块账户
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}

	// 将贷款记录追加到存储中
	k.AppendLoan(ctx, loan)

	return &types.MsgRequestLoanResponse{}, nil
}
