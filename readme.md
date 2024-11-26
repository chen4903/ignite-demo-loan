```bash
~/code/???/loan/cmd/loand main* ❯ go run . tx loan request-loan 1000token 100token 1000foocoin 500 --from alice --chain-id loan
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /loan.loan.MsgRequestLoan
    amount: 1000token
    collateral: 1000foocoin
    creator: cosmos1jleaq5w4xt0e5xwfdv4j5gk3yrfmv5x8ya0wma
    deadline: "500"
    fee: 100token
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: ""
timestamp: ""
tx: null
txhash: 9A8888150A0DAFB838BE660D6EA56E51F9AAB098AA1636040B0DB086BAE028BF

~/code/???/loan/cmd/loand main* ❯ go run . tx loan approve-loan 0 --from bob --chain-id loan
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /loan.loan.MsgApproveLoan
    creator: cosmos1puyud4ta9zqaphf8737s2cy4nwj2n09tjddfw3
    id: "0"
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: ""
timestamp: ""
tx: null
txhash: 52335149C0996F1E852227C884D9898543A0E0407C61D63CA93E5C1170A6F8A0

~/code/???/loan/cmd/loand main* ❯ go run . tx loan repay-loan 0 --from alice --chain-id loan
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /loan.loan.MsgRepayLoan
    creator: cosmos1jleaq5w4xt0e5xwfdv4j5gk3yrfmv5x8ya0wma
    id: "0"
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: ""
timestamp: ""
tx: null
txhash: 7F499AFD04A41FAC663FA23C03049F7D23EAE1E0B710FA55B2E7DF8218CBE424

~/code/???/loan/cmd/loand main* ❯ go run . tx loan request-loan 1000token 100token 1000foocoin 20 --from alice --chain-id loan -y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: ""
timestamp: ""
tx: null
txhash: 57615953F18234BA5674627B7F5023F5A3650C4D802631A39982255B4F40C0DF

~/code/???/loan/cmd/loand main* ❯ go run . tx loan approve-loan 1 --from bob --chain-id loan -y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: ""
timestamp: ""
tx: null
txhash: 2A3263D24396C002E989E86A3EC3FB63118DD8E29B158814BFA14694573D3BD9

~/code/???/loan/cmd/loand main* ❯ go run . tx loan liquidate-loan 1 --from bob --chain-id loan -y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: ""
timestamp: ""
tx: null
txhash: 81C4A10E014798DCD3E04A7E5C204DDDA32C94251EB511CA8E70348DAF702514

~/code/???/loan/cmd/loand main* ❯ go run . tx loan liquidate-loan 1 --from bob --chain-id loan -y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: ""
timestamp: ""
tx: null
txhash: 62296C6EF014CEF41939520EC4CA6276D444D3302A7F1BD5E6C9B753C39A3E70

~/code/???/loan/cmd/loand main* ❯ go run . q loan list-loan
Loan:
- amount: 1000token
  borrower: cosmos1jleaq5w4xt0e5xwfdv4j5gk3yrfmv5x8ya0wma
  collateral: 1000foocoin
  deadline: "500"
  fee: 100token
  lender: cosmos1puyud4ta9zqaphf8737s2cy4nwj2n09tjddfw3
  state: repayed
- amount: 1000token
  borrower: cosmos1jleaq5w4xt0e5xwfdv4j5gk3yrfmv5x8ya0wma
  collateral: 1000foocoin
  deadline: "20"
  fee: 100token
  id: "1"
  lender: cosmos1puyud4ta9zqaphf8737s2cy4nwj2n09tjddfw3
  state: liquidated
pagination:
  total: "2"
  
~/code/???/loan/cmd/loand main* ❯ go run . q bank balances $(go run . keys show alice -a)
balances:
- amount: "9000"
  denom: foocoin
- amount: "100000000"
  denom: stake
- amount: "20900"
  denom: token
pagination:
  total: "3"
```