// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"time"

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
	resp = &types.GetBtcAddressIsTrialResp{
		Address: req.Address,
	}
	item, err := l.collectTrialItems(req.Address)
	if err != nil {
		return nil, err
	}
	groupedItem := lo.GroupBy(item, func(btcTran types.Task) string {
		return btcTran.Status
	})
	if len(groupedItem[model.BtcTranStatusApprovedInEvm]) > 0 {
		resp.TrialComplete = true
		resp.TrialInfo = &groupedItem[model.BtcTranStatusApprovedInEvm][0]
		return
	}
	if len(groupedItem[model.BtcTranStatusBinded]) > 0 {
		resp.TrialInfo = &groupedItem[model.BtcTranStatusBinded][0]
		return
	}
	if len(groupedItem[model.BtcTranStatusInit]) > 0 {
		latest := groupedItem[model.BtcTranStatusInit][len(groupedItem[model.BtcTranStatusInit])-1]
		if time.Now().AddDate(0, 0, -3).Unix() <= int64(latest.BlockTime) {
			resp.TrialInfo = &latest
			return
		}
	}
	return
}

func (l *GetBtcAddressIsTrialLogic) collectTrialItems(address string) ([]types.Task, error) {
	var btcTasks []model.BtcTran
	var bindEvmSigns []model.BindEvmSign
	var evmHashInfos []model.EvmHashInfo

	sql := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BtcTran{})
	// sql.Where("JSON_EXTRACT(input_utxo, '$[0]') = ?", address)
	sql.Where("input0 = ?", address)
	if err := sql.Where("CAST(amount_satoshi AS UNSIGNED) + CAST(fee_satoshi AS UNSIGNED) = ?", l.svcCtx.Config.TinyTry).
		Where("trial_skip = ?", false).
		Order("block_time").Order("id desc").Find(&btcTasks).Error; err != nil {
		return nil, err
	}

	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BindEvmSign{}).
		Where("btc_tran_hash IN ?", lo.Map(btcTasks, func(item model.BtcTran, index int) string {
			return item.TransactionHash
		})).Find(&bindEvmSigns).Error; err != nil {
		return nil, err
	}

	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&model.EvmHashInfo{}).
		Where("transaction_hash IN ?", lo.Compact(lo.FlatMap(btcTasks, func(item model.BtcTran, index int) []string {
			return []string{item.RecievedEvmTxHash, item.AcceptedEvmTxHash, item.RejectedEvmTxHash}
		}))).Find(&evmHashInfos).Error; err != nil {
		return nil, err
	}
	return ItemsToTask(btcTasks, bindEvmSigns, evmHashInfos), nil
}

func (l *GetBtcAddressIsTrialLogic) v1GetBtcAddressIsTrial(req *types.GetBtcAddressIsTrialReq) (resp *types.GetBtcAddressIsTrialResp, err error) {
	var /*memBtcTask,*/ btcTasks []model.BtcTran
	// _ = l.svcCtx.MemDB.WithContext(l.ctx).Model(&model.BtcTran{}).Find(&memBtcTask).Error
	var bindEvmSigns []model.BindEvmSign
	var evmHashInfos []model.EvmHashInfo
	resp = &types.GetBtcAddressIsTrialResp{
		Address: req.Address,
	}

	sql := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BtcTran{})
	// sql.Where("JSON_EXTRACT(input_utxo, '$[0]') = ?", req.Address)
	sql.Where("input0 = ?", req.Address)
	if err := sql.Where("CAST(amount_satoshi AS UNSIGNED) + CAST(fee_satoshi AS UNSIGNED) = ?", l.svcCtx.Config.TinyTry).
		Where("trial_skip = ?", false).
		Order("block_time").Order("id desc").Limit(1).
		Find(&btcTasks).Error; err != nil {
		return nil, err
	}

	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BindEvmSign{}).
		Where("btc_tran_hash IN ?", lo.Map(btcTasks, func(item model.BtcTran, index int) string {
			return item.TransactionHash
		})).Find(&bindEvmSigns).Error; err != nil {
		return nil, err
	}

	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&model.EvmHashInfo{}).
		Where("transaction_hash IN ?", lo.Compact(lo.FlatMap(btcTasks, func(item model.BtcTran, index int) []string {
			return []string{item.RecievedEvmTxHash, item.AcceptedEvmTxHash, item.RejectedEvmTxHash}
		}))).Find(&evmHashInfos).Error; err != nil {
		return nil, err
	}
	item := ItemsToTask(btcTasks, bindEvmSigns, evmHashInfos)
	if len(item) > 0 && item[0].Status == model.BtcTranStatusApprovedInEvm {
		resp.TrialComplete = true
	}
	if len(item) > 0 {
		resp.TrialInfo = &item[0]
	}
	return
}
