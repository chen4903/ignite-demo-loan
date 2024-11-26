package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"loan/x/loan/types"
)

// RepayLoan 处理还款的请求消息
func (k msgServer) RepayLoan(goCtx context.Context, msg *types.MsgRepayLoan) (*types.MsgRepayLoanResponse, error) {
	// 将上下文从 goCtx 转换为 SDK 的上下文
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 根据贷款 ID 获取对应的贷款信息
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}

	// 检查贷款状态是否为 "approved"
	if loan.State != "approved" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}

	// 获取放款人和借款人的地址
	lender, _ := sdk.AccAddressFromBech32(loan.Lender)     // 放款人地址
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower) // 借款人地址

	// 验证还款请求是否由借款人发起
	if msg.Creator != loan.Borrower {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot repay: not the borrower")
	}

	// 解析贷款金额、手续费和抵押品为 Coins 对象
	amount, _ := sdk.ParseCoinsNormalized(loan.Amount)         // 贷款金额
	fee, _ := sdk.ParseCoinsNormalized(loan.Fee)               // 手续费
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral) // 抵押品

	// 借款人向放款人还款
	err := k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, err
	}

	// 借款人支付手续费给放款人
	err = k.bankKeeper.SendCoins(ctx, borrower, lender, fee)
	if err != nil {
		return nil, err
	}

	// 将抵押品从模块账户退还给借款人
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}

	// 更新贷款状态为 "repayed"
	loan.State = "repayed"
	k.SetLoan(ctx, loan) // 保存更新后的贷款信息

	return &types.MsgRepayLoanResponse{}, nil
}
