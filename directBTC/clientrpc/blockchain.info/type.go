package blockchaininfo

type AddressTransactions struct {
	Hash160       string        `json:"hash160"`
	Address       string        `json:"address"`
	NTx           int64         `json:"n_tx"`
	NUnredeemed   int64         `json:"n_unredeemed"`
	TotalReceived int64         `json:"total_received"`
	TotalSent     int64         `json:"total_sent"`
	FinalBalance  int64         `json:"final_balance"`
	Txs           []Transaction `json:"txs"`
}

type Transaction struct {
	Hash        string   `json:"hash"`
	Ver         int64    `json:"ver"`
	VinSz       int64    `json:"vin_sz"`
	VoutSz      int64    `json:"vout_sz"`
	Size        int64    `json:"size"`
	Weight      int64    `json:"weight"`
	Fee         int64    `json:"fee"`
	RelayedBy   string   `json:"relayed_by"`
	LockTime    int64    `json:"lock_time"`
	TxIndex     int64    `json:"tx_index"`
	DoubleSpend bool     `json:"double_spend"`
	Time        int64    `json:"time"`
	BlockIndex  int64    `json:"block_index"`
	BlockHeight int64    `json:"block_height"`
	Inputs      []Input  `json:"inputs"`
	Out         []Output `json:"out"`
	Result      int64    `json:"result"`
	Balance     int64    `json:"balance"`
}

type Output struct {
	Type  int64 `json:"type"`
	Spent bool  `json:"spent"`
	Value int64 `json:"value"`
	// SpendingOutpoints []interface{} `json:"spending_outpoints"`
	N       int64  `json:"n"`
	TxIndex int64  `json:"tx_index"`
	Script  string `json:"script"`
	Addr    string `json:"addr"`
}

type Input struct {
	Sequence int64  `json:"sequence"`
	Witness  string `json:"witness"`
	Script   string `json:"script"`
	Index    int64  `json:"index"`
	PrevOut  struct {
		Type              int64 `json:"type"`
		Spent             bool  `json:"spent"`
		Value             int64 `json:"value"`
		SpendingOutpoints []struct {
			TxIndex int64 `json:"tx_index"`
			N       int64 `json:"n"`
		} `json:"spending_outpoints"`
		N       int64  `json:"n"`
		TxIndex int64  `json:"tx_index"`
		Script  string `json:"script"`
		Addr    string `json:"addr"`
	} `json:"prev_out"`
}
