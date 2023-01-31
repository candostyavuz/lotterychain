export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

make show_module_balance
make show_client_balances

make enter_lottery
make show_lottery

make show_module_balance
make show_client_balances