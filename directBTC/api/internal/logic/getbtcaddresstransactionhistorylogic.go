// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"encoding/json"

	"directBTC/api/internal/svc"
	"directBTC/api/internal/types"
	"directBTC/model"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBtcAddressTransactionHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBtcAddressTransactionHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBtcAddressTransactionHistoryLogic {
	return &GetBtcAddressTransactionHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBtcAddressTransactionHistoryLogic) GetBtcAddressTransactionHistory(req *types.GetBtcAddressTransactionHistoryReq) (resp *types.GetBtcAddressTransactionHistoryResp, err error) {
	resp = &types.GetBtcAddressTransactionHistoryResp{
		PageData: types.PageData{
			Limit:  req.Limit,
			Offset: req.Offset,
		},
	}
	var btcTasks []model.BtcTran
	sql := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BtcTran{})
	if req.Address != "" {
		sql.Where("JSON_EXTRACT(input_utxo, '$[0]') = ?", req.Address)
	}
	if err := sql.Count(&resp.Total).Order("block_time desc").Order("id desc").
		Limit(req.Limit).Offset(req.Offset).Find(&btcTasks).Error; err != nil {
		return nil, err
	}
	resp.Data = lo.Map(btcTasks, func(item model.BtcTran, _ int) types.Transaction {
		return types.Transaction{
			ID:   uint64(item.ID),
			Hash: item.TransactionHash,
			TreasuryAddress: func() []string {
				var addrs []string
				_ = json.Unmarshal(item.TreasuryAddress, &addrs)
				return addrs
			}(),
			AmountSatoshi: item.AmountSatoshi,
			FeeSatoshi:    item.FeeSatoshi,
			InputAddress: func() []string {
				var addrs []string
				_ = json.Unmarshal(item.InputUtxo, &addrs)
				return addrs
			}(),
			BlockNumber:      item.BlockNumber,
			BlockTime:        item.BlockTime,
			ConfirmNumber:    item.ConfirmNumber,
			ConfirmThreshold: item.ConfirmThreshold,
		}
	})
	return
}
