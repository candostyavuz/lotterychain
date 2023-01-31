# lotterychain
**lotterychain** is a blockchain built using Cosmos SDK (v0.46.7) and Tendermint and created with [Ignite CLI (v0.26.1)](https://ignite.com/cli).

🥷🏻 [Can Dost Yavuz](https://www.linkedin.com/in/candosty/)
---
## Start Lottery Chain

Start the chain locally with

```
ignite chain serve --reset-once
```

`serve` command installs dependencies, builds, initializes, and starts the blockchain in development.
The blockchain (devnet) is configured with `config.yml`. Currently it has 20 client test accounts and a single validator account.
---

## Modules

**Lottery** module includes handler and keeper implementation of `EnterLottery` transactions. The module also includes storage objects for `Lottery` object and `Participant` object. 

`./x/lottery/abci.go` includes `EndBlocker` routine for determining the lottery winners and for distributing the reward payouts if the conditions are met. 

**Cosmos SDK Modules**: The chain utilizes methods from the native Cosmos SDK modules, `x/bank` and `x/staking`, for asset transactions and validator set retrieval to prevent block proposers from participating in the lottery.

## EnterLottery Transaction

**CLI Format**
`lotterychaind tx lottery enter-lottery [fee] [bet] [flags]`

**Transaction Rules**
* Block proposers are not allowed to join the lottery as participants. Active validator set is constantly being checked by the module.
* `fee` must be `5000000token` (10^6 decimal fields included)
* `bet` must be between `1_000_000token` and `100_000_000token` (10^6 decimal fields included)
* Each user can only participate once in the lottery session. 
* **If multiple transactions are sent from the same address:**
    - Only last transaction counts and last sent `bet` amount will be recorded in the `Participant` object.
    - Previous `bet` that user paid will be refunded back to its wallet.
    - Any `fee` paid to contract is NOT refunded and stored in the lottery prize pool.
    - If previous `bet` of the user is max / min bet of the current lottery session, the new min/max bet is automatically adjusted based on the current bets in the pool and last bet amount in the participant transaction.
* After total transaction count in the lottery hits 10 (or greater), `EndBlocker` is triggered and winner is detected according to modulus of hash of all appended txData with total number of transactions.
* Payment is done according to rules below:
    - If winner has the highest bet, entire pool including all the `fee` and `bet` is paid to winner.
    - If winner has the lowest bet, nothing is paid to winner and lottery prize pool is carried over to next session.
    - Any other winner bet case, all the `bet` amount (excluding `fee`) in the prize pool is paid to winner.
    - All funds are stored in the `lottery` module.

**Example transactions**
Note: export GOPATH if not already:
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

```
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
```

## EnterLottery Query Methods

1. `lotterychaind q lottery list-participant` 
Queries the information of all participants of the current lottery session

```
Participant:
- address: cosmos1cqc23mayjs6fd4fwmk0jc58ryuag8ztaftpudw
  bet:
    amount: "1000000"
    denom: token
  id: "0"
  txData: creator:"cosmos1cqc23mayjs6fd4fwmk0jc58ryuag8ztaftpudw" fee:<denom:"token"
    amount:"5000000" > bet:<denom:"token" amount:"1000000" > Tue Jan 31 07:31:33 UTC
    2023
pagination:
  next_key: null
  total: "0"
  ```

2. `lotterychaind q lottery show-lottery`
Queries the information about the current lottery session.
```
Lottery:
  currentMaxBet:
    amount: "2000000"
    denom: token
  currentMinBet:
    amount: "1000000"
    denom: token
  lastWinner: cosmos1rf8kzu97tjr0379nz89uf3sta73svhpyjfvary
  lastWinnerIdx: "6"
  totalBets:
    amount: "3000000"
    denom: token
  totalFees:
    amount: "115000000"
    denom: token
  txCounter: "2"
  txDataAll: creator:"cosmos1cqc23mayjs6fd4fwmk0jc58ryuag8ztaftpudw" fee:<denom:"token"
    amount:"5000000" > bet:<denom:"token" amount:"1000000" > Tue Jan 31 07:34:09 UTC
    2023creator:"cosmos1gl73n529fhxw55yv6mtsf26f9n47602qem4a8q" fee:<denom:"token"
    amount:"5000000" > bet:<denom:"token" amount:"2000000" > Tue Jan 31 07:34:10 UTC
    2023
```

3. `lotterychaind q bank balances cosmos1helefwcscjl8k3rlqe0zrvcps5acf9jtsnfelu`
Queries total balances stored in `Lottery`module's pool

4. `lotterychaind q bank balances $(lotterychaind keys show client1 -a)`
Queries total balances of a specific client
