remove:
	rm -rf ~/.lotterychain

reset:
	lotterychaind tendermint unsafe-reset-all

init:
	lotterychaind init lotterychain --chain-id lotterychain

config:
	lotterychaind config keyring-backend test
	lotterychaind config chain-id lotterychain

keys:
	lotterychaind keys add client1 --keyring-backend test --algo secp256k1
	lotterychaind keys add client2 --keyring-backend test --algo secp256k1
	lotterychaind keys add client3 --keyring-backend test --algo secp256k1
	lotterychaind keys add client4 --keyring-backend test --algo secp256k1
	lotterychaind keys add client5 --keyring-backend test --algo secp256k1
	lotterychaind keys add client6 --keyring-backend test --algo secp256k1
	lotterychaind keys add client7 --keyring-backend test --algo secp256k1
	lotterychaind keys add client8 --keyring-backend test --algo secp256k1
	lotterychaind keys add client9 --keyring-backend test --algo secp256k1
	lotterychaind keys add client10 --keyring-backend test --algo secp256k1
	lotterychaind keys add client11 --keyring-backend test --algo secp256k1
	lotterychaind keys add client12 --keyring-backend test --algo secp256k1
	lotterychaind keys add client13 --keyring-backend test --algo secp256k1
	lotterychaind keys add client14 --keyring-backend test --algo secp256k1
	lotterychaind keys add client15 --keyring-backend test --algo secp256k1
	lotterychaind keys add client16 --keyring-backend test --algo secp256k1
	lotterychaind keys add client17 --keyring-backend test --algo secp256k1
	lotterychaind keys add client18 --keyring-backend test --algo secp256k1
	lotterychaind keys add client19 --keyring-backend test --algo secp256k1
	lotterychaind keys add client20 --keyring-backend test --algo secp256k1
	lotterychaind keys add lotteryvalidator --keyring-backend test --algo secp256k1

parameter_token_denomination:
	cat ~/.lotterychain/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="token"' > ~/.lotterychain/config/tmp_genesis.json && mv ~/.lotterychain/config/tmp_genesis.json ~/.lotterychain/config/genesis.json
	cat ~/.lotterychain/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="token"' > ~/.lotterychain/config/tmp_genesis.json && mv ~/.lotterychain/config/tmp_genesis.json ~/.lotterychain/config/genesis.json
	cat ~/.lotterychain/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="token"' > ~/.lotterychain/config/tmp_genesis.json && mv ~/.lotterychain/config/tmp_genesis.json ~/.lotterychain/config/genesis.json
	cat ~/.lotterychain/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="token"' > ~/.lotterychain/config/tmp_genesis.json && mv ~/.lotterychain/config/tmp_genesis.json ~/.lotterychain/config/genesis.json

allocate_genesis_accounts:
	lotterychaind add-genesis-account client1 500000000token --keyring-backend test
	lotterychaind add-genesis-account client2 500000000token --keyring-backend test
	lotterychaind add-genesis-account client3 500000000token --keyring-backend test
	lotterychaind add-genesis-account client4 500000000token --keyring-backend test
	lotterychaind add-genesis-account client5 500000000token --keyring-backend test
	lotterychaind add-genesis-account client6 500000000token --keyring-backend test
	lotterychaind add-genesis-account client7 500000000token --keyring-backend test
	lotterychaind add-genesis-account client8 500000000token --keyring-backend test
	lotterychaind add-genesis-account client9 500000000token --keyring-backend test
	lotterychaind add-genesis-account client10 500000000token --keyring-backend test
	lotterychaind add-genesis-account client11 500000000token --keyring-backend test
	lotterychaind add-genesis-account client12 500000000token --keyring-backend test
	lotterychaind add-genesis-account client13 500000000token --keyring-backend test
	lotterychaind add-genesis-account client14 500000000token --keyring-backend test
	lotterychaind add-genesis-account client15 500000000token --keyring-backend test
	lotterychaind add-genesis-account client16 500000000token --keyring-backend test
	lotterychaind add-genesis-account client17 500000000token --keyring-backend test
	lotterychaind add-genesis-account client18 500000000token --keyring-backend test
	lotterychaind add-genesis-account client19 500000000token --keyring-backend test
	lotterychaind add-genesis-account client20 500000000token --keyring-backend test
	lotterychaind add-genesis-account lotteryvalidator 10000000000000token --keyring-backend test

sign_genesis_transaction:
	lotterychaind gentx lotteryvalidator 100000000token --keyring-backend test --chain-id lotterychain

collect_genesis_tx:
	lotterychaind collect-gentxs

validate_genesis:
	lotterychaind validate-genesis

ifdef start
	lotterychaind start --log_level info
endif

all: reset init config keys parameter_token_denomination allocate_genesis_accounts sign_genesis_transaction collect_genesis_tx validate_genesis