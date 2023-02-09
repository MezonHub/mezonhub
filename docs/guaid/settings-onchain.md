# Settings on chain

## Setting examples for admin

### set protocol fee receiver

```bash
mezon tx ledger set-protocol-fee-receiver mezon1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --from admin --chain-id local-mezonhub --keyring-backend file

mezon query ledger protocol-fee-receiver 
```

### add new rtoken

```bash
# set rtoken metadata
mezon tx zbank add-denom cosmos cosmosvaloper ./metadata/metadata_ratom.json --chain-id local-mezonhub --from admin --keyring-backend file

mezon query bank denom-metadata

mezon query zbank address-prefix uratom

# set relay fee receiver
mezon tx ledger set-relay-fee-receiver uratom mezon1mgjkpyfm00mxk0nmhvfvwhlr65067d538l6cxa --from admin --chain-id local-mezonhub --keyring-backend file

mezon query ledger relay-fee-receiver uratom

# this will init bonded pool, exchange rate, pipeline
mezon tx ledger init-pool uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 --from admin --chain-id local-mezonhub --keyring-backend file

mezon query ledger bonded-pools uratom

mezon query ledger exchange-rate uratom

mezon query ledger bond-pipeline uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



# add relayers
mezon tx relayers add-relayers ledger uratom mezon1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx:mezon1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --from admin --chain-id local-mezonhub

mezon query relayers relayers ledger uratom

# set threshold
mezon tx relayers set-threshold ledger uratom 1 --from admin --keyring-backend file --chain-id local-mezonhub

mezon query relayers threshold ledger uratom

# set params used by relay
mezon tx ledger set-r-params uratom 0.00001stake 600 0 2 0stake --from admin --keyring-backend file --chain-id local-mezonhub

mezon query ledger r-params uratom

# set pool detail for multisig/ica pool
mezon tx ledger set-pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmos1cad0efr25faywnjp8qp36l8zlqa2sgz0jwn0hl:cosmos13mwxtgrljf9d5r72sc28496ua4lsga0jvmqz8x 1 --from admin --chain-id local-mezonhub --keyring-backend file

mezon query ledger pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



# default 0.1
mezon tx ledger set-staking-reward-commission uratom 0.15 --from admin --chain-id local-mezonhub --keyring-backend file

mezon q ledger staking-reward-commission uratom

# default 0.002
mezon tx ledger set-unbond-commission uratom 0.0025 --from admin --chain-id local-mezonhub --keyring-backend file

mezon q ledger unbond-commission uratom

# default 1000000umez
mezon tx ledger set-unbond-relay-fee uratom 1000005umez --from admin --chain-id local-mezonhub --keyring-backend file

mezon q ledger unbond-relay-fee uratom

```

### register ica pool
```
# register ica pool (need set rtoken metadata before this)
mezon tx ledger register-ica-pool uratom connection-0 --keyring-backend file --from admin --chain-id local-mezonhub --gas 410000

mezon q ledger ica-pool-list uratom

# set withdrawal address
mezon tx ledger set-withdrawal-addr cosmos1gsth46z50w256p4kq36xquh4q90mfjq0t4lm9scln6zucg64epyqudzqzm --keyring-backend file --from admin --chain-id local-mezonhub --gas 410000

```

### zvalidator

```bash
# add relayers
mezon tx relayers add-relayers zvalidator uratom mezon14z467aut40mcrt2ddyxf7e74fq99udul7kaf9g:mezon15lne70yk254s0pm2da6g59r82cjymzjqvvqxz7 --keyring-backend file --from admin --chain-id local-mezonhub

mezon q relayers relayers zvalidator uratom

# set threshold
mezon tx relayers set-threshold zvalidator uratom 1 --from admin --keyring-backend file --chain-id local-mezonhub

mezon q relayers threshold zvalidator uratom

# init zvalidator (should init target validators of pool before rtoken relay start)
mezon tx zvalidator init-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --chain-id local-mezonhub --keyring-backend file  

mezon q zvalidator r-validator-list uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

# add zvalidator
mezon tx zvalidator add-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper1cad0efr25faywnjp8qp36l8zlqa2sgz0h686mv  --chain-id local-mezonhub --keyring-backend file --from admin

mezon q zvalidator r-validator-list uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

# rm zvalidator
mezon tx zvalidator rm-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper1cad0efr25faywnjp8qp36l8zlqa2sgz0h686mv cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --chain-id local-mezonhub --keyring-backend file
```



### bridge

```bash
mezon tx bridge add-chain-id 1 --from admin --keyring-backend file --chain-id local-mezonhub

mezon query bridge chain-ids



mezon tx relayers add-relayers bridge 1 mezon1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-mezonhub

mezon query relayers relayers bridge 1



mezon tx relayers set-threshold bridge 1 1 --from admin --keyring-backend file --chain-id local-mezonhub

mezon query relayers threshold bridge 1



mezon tx bridge set-resourceid-to-denom  000000000000000000000000000000a9e0095b8965c01e6a09c97938f3860901 uratom NATIVE --from admin --keyring-backend file --chain-id local-mezonhub

mezon query bridge resourceid-to-denoms



mezon tx bridge set-relay-fee-receiver mezon1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-mezonhub

mezon query bridge relay-fee-receiver



mezon tx bridge set-relay-fee 1 1000000umez --from admin --keyring-backend file --chain-id local-mezonhub

mezon query bridge  relay-fee 1


mezon tx bridge add-banned-denom 1 uratom --from admin --keyring-backend file --chain-id local-mezonhub

mezon q bridge banned-denom-list
```

### migrate rtoken (after adding new rtoken step)

```bash
mezon tx ledger migrate-init uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100000000 150000000 200000000 300000000 1.23 --from admin --keyring-backend file --chain-id local-mezonhub

mezon query bank  total 

mezon query ledger exchange-rate uratom

mezon query ledger bond-pipeline uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



mezon tx ledger migrate-unbondings uratom --unbondings ./unbondings_example.json --from admin --keyring-backend file --chain-id local-mezonhub

mezon query ledger pool-unbond uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 3



mezon tx bridge set-denom-type uratom  1 --from admin --keyring-backend file --chain-id local-mezonhub

mezon query bridge denom-types
```

### zdex

```bash
mezon tx zdex create-pool 10umez 20uratom --from admin --chain-id local-mezonhub --keyring-backend file

mezon tx zdex add-provider mezon1qzt0qajzr9df3en5sk06xlk26n30003c8uhdkg --from admin --chain-id local-mezonhub --keyring-backend file

mezon tx zdex add-liquidity  100umez 200uratom --from admin --chain-id local-mezonhub --keyring-backend file

mezon tx zdex remove-liquidity 10 5 1uratom 1umez umez --from admin --chain-id local-mezonhub --keyring-backend file

mezon tx zdex swap 2umez 1uratom  --from admin --chain-id local-mezonhub --keyring-backend file
```

### mining

```bash
mezon tx mining add-mining-provider mezon1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx  --from admin --chain-id local-mezonhub --keyring-backend file

mezon tx mining add-reward-token umez 200 --from admin --chain-id local-mezonhub --keyring-backend file


mezon tx mining add-stake-pool umez ./add_stake_pool_example.json  --from relay1 --chain-id local-mezonhub --keyring-backend file

mezon tx mining stake 0 10umez 0 --from my-account --chain-id local-mezonhub --keyring-backend file 

mezon tx mining claim-reward 0 0 --from my-account --chain-id local-mezonhub --keyring-backend file

mezon tx mining add-reward 1 0 300 0 0 --from relay1 --chain-id local-mezonhub --keyring-backend file

mezon tx mining withdraw 1 10umez 0 --from test --chain-id local-mezonhub --keyring-backend file
```



## Operate examples for user

### liquidity bond (gaiad example)

```bash
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1000000stake --note 1:mezon1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --chain-id local-cosmos
```

### recover (gaiad example)

```bash
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1stake --note 2:mezon1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc:9A80F3E6A007E1144BE34F4A0AC35B9288C19641BCAD3464277168000AF5FC66 --keyring-backend file --chain-id local-cosmos
```

### liquidity unbond

```bash
mezon tx ledger liquidity-unbond cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100uratom cosmos1j9dues7ey2a39nes4ewfvyma96d3f5zrdhnfan --keyring-backend file --from user --home /Users/tpkeeper/gowork/mezon/rtoken-relay-core/keys/mezonhub --chain-id local-mezonhub
```

### deposit (transfer token to external chain)

```bash
mezon tx bridge deposit 1 uratom 800 dccf954570063847d73746afa0b0878f2c779d42089c5d9a107f2aca176e985f --from my-account --chain-id local-mezonhub --keyring-backend file
```
