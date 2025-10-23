package mempoolspace

type AddressTransaction struct {
	Txid     string `json:"txid"`
	Version  int64  `json:"version"`
	Locktime int64  `json:"locktime"`
	Size     int64  `json:"size"`
	Weight   int64  `json:"weight"`
	Fee      int64  `json:"fee"`
	Vin      []Vin  `json:"vin"`
	Vout     []Vout `json:"vout"`
	Status   struct {
		Confirmed   bool   `json:"confirmed"`
		BlockHeight uint64 `json:"block_height"`
		BlockHash   string `json:"block_hash"`
		BlockTime   uint64 `json:"block_time"`
	} `json:"status"`
}

type Vout struct {
	Value               int64  `json:"value"`
	Scriptpubkey        string `json:"scriptpubkey"`
	ScriptpubkeyAddress string `json:"scriptpubkey_address"`
	ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
	ScriptpubkeyType    string `json:"scriptpubkey_type"`
}

type Vin struct {
	IsCoinbase bool `json:"is_coinbase"`
	Prevout    struct {
		Value               int64  `json:"value"`
		Scriptpubkey        string `json:"scriptpubkey"`
		ScriptpubkeyAddress string `json:"scriptpubkey_address"`
		ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
		ScriptpubkeyType    string `json:"scriptpubkey_type"`
	} `json:"prevout"`
	Scriptsig             string   `json:"scriptsig"`
	ScriptsigAsm          string   `json:"scriptsig_asm"`
	Sequence              int64    `json:"sequence"`
	Txid                  string   `json:"txid"`
	Vout                  int64    `json:"vout"`
	Witness               []string `json:"witness"`
	InnerRedeemscriptAsm  string   `json:"inner_redeemscript_asm"`
	InnerWitnessscriptAsm string   `json:"inner_witnessscript_asm"`
}
