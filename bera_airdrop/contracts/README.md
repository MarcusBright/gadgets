# **Airdrop Contract**

A **Merkle Tree-based** token airdrop contract system.

---

## **🚀 Core Features**

- ✅ **Merkle Tree Verification** – Validates user eligibility using Merkle proofs  
- ✅ **Duplicate Claim Prevention** – Each address can only claim once  
- ✅ **Event Tracking** – Records all claim events  

---

## **📜 Contract Architecture**

| Contract      | Description           |
| ------------- | --------------------- |
| `Airdrop.sol` | Main airdrop contract |

---

## **🛠 Development Tools**

This project uses the **Foundry** framework:

- **Forge** – Ethereum testing framework  
- **Cast** – Command-line tool for interacting with EVM smart contracts  
- **Anvil** – Local Ethereum node  

---

## **🚀 Getting Started**

### **📌 Install Dependencies**
```sh
forge install
```

### **🔧 Compile Contracts**
```sh
forge build
```

### **🧪 Run Tests**
```sh
forge test
```

---

## **📤 Deployment Process**

1. Prepare **Merkle Tree root**  
2. Deploy **Airdrop contract** with:
   - **Merkle Root**

```sh
forge script script/Airdrop.s.sol:AirdropScript --rpc-url <your_rpc_url> --private-key <your_private_key>
```

### Error Codes from contracts

| Error Code | Description                           |
| ---------- | ------------------------------------- |
| SYS001     | INVALID_ADDRESS                       |
| SYS002     | INVALID_INPUT_PARAMETER               |
| SYS003     | APPROVE_FAILED                        |
| USR001     | CURRENT_EPOCH_IS_STILL_VALID          |
| USR002     | NO_ACTIVE_EPOCH                       |
| USR003     | INVALID_ROOT                          |
| USR004     | NEW_DURATION_WOULD_EXPIRE_IMMEDIATELY |
| USR005     | ALREADY_CLAIMED                       |
| USR006     | DISTRIBUTION_DISABLED                 |
| USR007     | DISTRIBUTION_NOT_ACTIVATED            |
| USR008     | DISTRIBUTION_EXPIRED                  |
| USR009     | INVALID_PROOF                         |

 