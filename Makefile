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

enter_lottery:
	lotterychaind tx lottery enter-lottery 5000000token 1000000token --from client1 -y
	lotterychaind tx lottery enter-lottery 5000000token 2000000token --from client2 -y
	lotterychaind tx lottery enter-lottery 5000000token 3000000token --from client3 -y
	lotterychaind tx lottery enter-lottery 5000000token 4000000token --from client4 -y
	lotterychaind tx lottery enter-lottery 5000000token 5000000token --from client5 -y
	lotterychaind tx lottery enter-lottery 5000000token 6000000token --from client6 -y
	lotterychaind tx lottery enter-lottery 5000000token 7000000token --from client7 -y
	lotterychaind tx lottery enter-lottery 5000000token 8000000token --from client8 -y
	lotterychaind tx lottery enter-lottery 5000000token 9000000token --from client9 -y
	lotterychaind tx lottery enter-lottery 5000000token 10000000token --from client10 -y
	lotterychaind tx lottery enter-lottery 5000000token 11000000token --from client11 -y
	lotterychaind tx lottery enter-lottery 5000000token 12000000token --from client12 -y
	lotterychaind tx lottery enter-lottery 5000000token 13000000token --from client13 -y
	lotterychaind tx lottery enter-lottery 5000000token 14000000token --from client14 -y
	lotterychaind tx lottery enter-lottery 5000000token 15000000token --from client15 -y
	lotterychaind tx lottery enter-lottery 5000000token 16000000token --from client16 -y
	lotterychaind tx lottery enter-lottery 5000000token 17000000token --from client17 -y
	lotterychaind tx lottery enter-lottery 5000000token 18000000token --from client18 -y
	lotterychaind tx lottery enter-lottery 5000000token 19000000token --from client19 -y
	lotterychaind tx lottery enter-lottery 5000000token 20000000token --from client20 -y

show_lottery:
	lotterychaind q lottery show-lottery

show_module_balance:
	lotterychaind q bank balances cosmos1helefwcscjl8k3rlqe0zrvcps5acf9jtsnfelu   

show_client_balances:
	lotterychaind q bank balances $$(lotterychaind keys show client1 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client2 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client3 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client4 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client5 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client6 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client7 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client8 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client9 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client10 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client11 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client12 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client13 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client14 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client15 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client16 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client17 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client18 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client19 -a)
	lotterychaind q bank balances $$(lotterychaind keys show client20 -a)

all: reset init config keys parameter_token_denomination allocate_genesis_accounts sign_genesis_transaction collect_genesis_tx validate_genesis