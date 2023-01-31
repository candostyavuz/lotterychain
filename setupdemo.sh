export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

make remove
make reset
make init
make config
make keys
make parameter_token_denomination
make allocate_genesis_accounts
make sign_genesis_transaction
make collect_genesis_tx
make validate_genesis
if [ "$1" = "start" ]; then
   lotterychaind start --log_level info
fi