package mempoolspace

import "testing"

func TestClient_GetAddressNewTransactions(t *testing.T) {
	client := NewClient("https://mempool.space", 2, 1)
	addr := "1AJbsFZ64EpEfS5UAjAfcUG8pH8Jn3rn1F"
	txs, err := client.GetAddressNewTransactions(addr, "")
	if err != nil {
		t.Fatalf("GetAddressNewTransactions failed: %v", err)
	}
	t.Log("len(txs)", len(txs))
}

func TestClient_GetAddressNewTransactions1(t *testing.T) {
	client := NewClient("https://mempool.space", 2, 1)
	addr := "1AJbsFZ64EpEfS5UAjAfcUG8pH8Jn3rn1F"
	txs, err := client.GetAddressNewTransactions(addr, "599e47a8114fe098103663029548811d2651991b62397e057f0c863c2bc9f9ea")
	if err != nil {
		t.Fatalf("GetAddressNewTransactions failed: %v", err)
	}
	t.Log("len(txs)", len(txs))
}

func TestClient_GetLatestBlockNumber(t *testing.T) {
	client := NewClient("https://mempool.space", 2, 1)
	height, err := client.GetLatestBlockNumber()
	if err != nil {
		t.Fatalf("GetLatestBlockNumber failed: %v", err)
	}
	t.Log("height", height)
}

func TestClient_GetBlockHash(t *testing.T) {
	client := NewClient("https://mempool.space", 2, 1)
	hash, err := client.GetBlockHash(919140)
	if err != nil {
		t.Fatalf("GetBlockHash failed: %v", err)
	}
	t.Log("hash", hash)
}

func TestClient_GetBlockTxId(t *testing.T) {
	client := NewClient("https://mempool.space", 2, 1)
	txids, err := client.GetBlockTxId("00000000000000000000588844a5157684d920df5e6e9b3ff31888171b18c0af")
	if err != nil {
		t.Fatalf("GetBlockTxId failed: %v", err)
	}
	t.Log("len(txids)", len(txids))
}

func TestClient_GetTx(t *testing.T) {
	client := NewClient("https://mempool.space", 2, 1)
	tx, err := client.GetTx("15e10745f15593a899cef391191bdd3d7c12412cc4696b7bcb669d0feadc8521")
	if err != nil {
		t.Fatalf("GetTx failed: %v", err)
	}
	t.Log("tx", tx)
}
func TestClient_GetTxNotFound(t *testing.T) {
	client := NewClient("https://mempool.space", 2, 1)
	tx, err := client.GetTx("15e10745a15593a899cef391191bdd3d7c12412cc4696b7bcb669d0feadc8521")
	if err != nil {
		t.Fatalf("GetTx failed: %v", err)
	}
	t.Log("tx", tx)
}
