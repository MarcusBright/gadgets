package blockchaininfo

import "testing"

func TestClient_GetAddressTransactions(t *testing.T) {
	client := NewClient("https://blockchain.info", 1, 15)
	addr := "1AJbsFZ64EpEfS5UAjAfcUG8pH8Jn3rn1F"
	txs, err := client.GetAddressNewTransactions(addr, "")
	if err != nil {
		t.Fatalf("GetAddressNewTransactions failed: %v", err)
	}
	t.Log("len(txs)", len(txs))
}

func TestClient_GetAddressTransactions1(t *testing.T) {
	client := NewClient("https://blockchain.info", 1, 15)
	addr := "1AJbsFZ64EpEfS5UAjAfcUG8pH8Jn3rn1F"
	txs, err := client.GetAddressNewTransactions(addr, "599e47a8114fe098103663029548811d2651991b62397e057f0c863c2bc9f9ea")
	if err != nil {
		t.Fatalf("GetAddressNewTransactions failed: %v", err)
	}
	t.Log("len(txs)", len(txs))
}
