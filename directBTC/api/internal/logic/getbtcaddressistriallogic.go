// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"directBTC/api/internal/svc"
	"directBTC/api/internal/types"
	"directBTC/model"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBtcAddressIsTrialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBtcAddressIsTrialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBtcAddressIsTrialLogic {
	return &GetBtcAddressIsTrialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBtcAddressIsTrialLogic) GetBtcAddressIsTrial(req *types.GetBtcAddressIsTrialReq) (resp *types.GetBtcAddressIsTrialResp, err error) {
	var /*memBtcTask,*/ btcTasks []model.BtcTran
	// _ = l.svcCtx.MemDB.WithContext(l.ctx).Model(&model.BtcTran{}).Find(&memBtcTask).Error
	var bindEvmSigns []model.BindEvmSign
	resp = &types.GetBtcAddressIsTrialResp{
		Address: req.Address,
	}

	sql := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BtcTran{})
	sql.Where("JSON_EXTRACT(input_utxo, '$[0]') = ?", req.Address)
	if err := sql.Where("CAST(amount_satoshi AS UNSIGNED) + CAST(fee_satoshi AS UNSIGNED) = ?", l.svcCtx.Config.TinyTry).
		Order("block_time desc").
		Order("id desc").Limit(1).
		Find(&btcTasks).Error; err != nil {
		return nil, err
	}

	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BindEvmSign{}).
		Where("btc_tran_hash IN ?", lo.Map(btcTasks, func(item model.BtcTran, index int) string {
			return item.TransactionHash
		})).Find(&bindEvmSigns).Error; err != nil {
		return nil, err
	}
	item := ItemsToTask(btcTasks, bindEvmSigns)
	if len(item) > 0 && item[0].Status == model.BtcTranStatusApprovedInEvm {
		resp.TrialComplete = true
	}
	if len(item) > 0 {
		resp.TrialInfo = &item[0]
	}
	return
}
