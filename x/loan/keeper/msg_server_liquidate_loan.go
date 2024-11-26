package keeper

import (
    "context"
    "strconv"

    errorsmod "cosmossdk.io/errors"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

    "loan/x/loan/types"
)

// LiquidateLoan 处理清算贷款的请求消息
func (k msgServer) LiquidateLoan(goCtx context.Context, msg *types.MsgLiquidateLoan) (*types.MsgLiquidateLoanResponse, error) {
    // 将上下文从 goCtx 转换为 SDK 的上下文
    ctx := sdk.UnwrapSDKContext(goCtx)

    // 根据贷款 ID 获取对应的贷款信息
    loan, found := k.GetLoan(ctx, msg.Id)
    if !found {
        return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
    }

    // 验证清算请求是否由贷款的放款人发起
    if loan.Lender != msg.Creator {
        return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot liquidate: not the lender")
    }

    // 检查贷款状态是否为 "approved"
    if loan.State != "approved" {
        return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
    }

    // 获取放款人地址和抵押品
    lender, _ := sdk.AccAddressFromBech32(loan.Lender) // 放款人地址
    collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral) // 抵押品

    // 解析贷款的截止区块高度
    deadline, err := strconv.ParseInt(loan.Deadline, 10, 64)
    if err != nil {
        panic(err)
    }

    // 检查当前区块高度是否已超过截止日期
    if ctx.BlockHeight() < deadline {
        return nil, errorsmod.Wrap(types.ErrDeadline, "Cannot liquidate before deadline")
    }

    // 将抵押品从模块账户转移到放款人账户
    err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, collateral)
    if err != nil {
        return nil, err
    }

    // 更新贷款状态为 "liquidated"
    loan.State = "liquidated"
    k.SetLoan(ctx, loan) // 保存更新后的贷款信息

    return &types.MsgLiquidateLoanResponse{}, nil
}
