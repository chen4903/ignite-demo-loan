package keeper

import (
    "context"

    errorsmod "cosmossdk.io/errors"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

    "loan/x/loan/types"
)

// CancelLoan 处理取消贷款的请求消息
func (k msgServer) CancelLoan(goCtx context.Context, msg *types.MsgCancelLoan) (*types.MsgCancelLoanResponse, error) {
    // 将上下文从 goCtx 转换为 SDK 的上下文
    ctx := sdk.UnwrapSDKContext(goCtx)

    // 根据贷款 ID 获取对应的贷款信息
    loan, found := k.GetLoan(ctx, msg.Id)
    if !found {
        return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
    }

    // 验证操作的发起者是否为贷款的借款人
    if loan.Borrower != msg.Creator {
        return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot cancel: not the borrower")
    }

    // 检查贷款状态是否为 "requested"
    if loan.State != "requested" {
        return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
    }

    // 将贷款的抵押品从模块账户退回到借款人的账户
    borrower, _ := sdk.AccAddressFromBech32(loan.Borrower) // 获取借款人地址
    collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral) // 解析抵押品为 Coins 对象
    err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
    if err != nil {
        return nil, err
    }

    // 更新贷款状态为 "cancelled"
    loan.State = "cancelled"
    k.SetLoan(ctx, loan) // 将更新后的贷款信息保存到存储中

    return &types.MsgCancelLoanResponse{}, nil
}
