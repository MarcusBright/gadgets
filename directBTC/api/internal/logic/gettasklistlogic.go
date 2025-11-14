// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"encoding/json"
	"slices"

	"directBTC/api/internal/svc"
	"directBTC/api/internal/types"
	"directBTC/model"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskListLogic {
	return &GetTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskListLogic) GetTaskList(req *types.TaskListReq) (resp *types.TaskListResp, err error) {
	resp = &types.TaskListResp{
		PageData: types.PageData{
			Limit:  req.Limit,
			Offset: req.Offset,
		},
	}
	var /*memBtcTask,*/ btcTasks []model.BtcTran
	// _ = l.svcCtx.MemDB.WithContext(l.ctx).Model(&model.BtcTran{}).Find(&memBtcTask).Error
	var bindEvmSigns []model.BindEvmSign

	sql := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BtcTran{})
	if req.Address != "" {
		// sql.Where("JSON_EXTRACT(input_utxo, '$[0]') = ?", req.Address)
		sql.Where("input0 = ?", req.Address)
	}
	if len(req.Status) != 0 {
		sql.Where("status IN ?", req.Status)
		if slices.Contains(req.Status, model.BtcTranStatusRecievedInEvm) {
			sql.Order("process_idx asc")
		}
	}
	// sql.Where("confirm_number >= confirm_threshold")
	if req.OrderDir == "desc" {
		sql.Order("block_number desc")
	} else {
		sql.Order("block_number asc")
	}
	if err := sql.Count(&resp.Total).Limit(req.Limit).Offset(req.Offset).Find(&btcTasks).Error; err != nil {
		return nil, err
	}
	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BindEvmSign{}).
		Where("btc_tran_hash IN ?", lo.Map(btcTasks, func(item model.BtcTran, index int) string {
			return item.TransactionHash
		})).Find(&bindEvmSigns).Error; err != nil {
		return nil, err
	}

	resp.Data = ItemsToTask(btcTasks, bindEvmSigns)

	return
}

func ItemsToTask(item []model.BtcTran, sign []model.BindEvmSign) []types.Task {
	signMap := lo.SliceToMap(sign, func(item model.BindEvmSign) (string, model.BindEvmSign) {
		return item.BtcTranHash, item
	})
	return lo.Map(item, func(item model.BtcTran, index int) types.Task {
		return types.Task{
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
			Status:              item.Status,
			BlockNumber:         item.BlockNumber,
			BlockTime:           item.BlockTime,
			ConfirmNumber:       item.ConfirmNumber,
			ConfirmThreshold:    item.ConfirmThreshold,
			BindedEvmAddress:    item.BindedEvmAddress,
			ChainId:             uint64(item.ChainId),
			RecievedEventTxHash: item.RecievedEvmTxHash,
			AcceptedEventTxHash: item.AcceptedEvmTxHash,
			RejectedEventTxHash: item.RejectedEvmTxHash,
			BindSignInfo: func() *types.BindSignInfo {
				if sign, ok := signMap[item.TransactionHash]; ok {
					return &types.BindSignInfo{
						Message:   sign.Message,
						Signature: sign.Signature,
						Signer:    sign.Signer,
						BindTime:  sign.CreatedAt.Unix(),
					}
				}
				return nil
			}(),
		}
	})
}
