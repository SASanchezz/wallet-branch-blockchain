package listener

import (
	"context"
	"fmt"
	"log"
	"os"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/transaction"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Listen() {
	ts := transaction.New()
	defer ts.Close()

	url, _ := os.LookupEnv("BLOCKCHAIN_URL")

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, _ := client.SubscribeNewHead(context.Background(), headers)
	defer sub.Unsubscribe()

	for header := range headers {
		block, err := client.BlockByHash(context.Background(), header.Hash())
		if err != nil {
			log.Println(err)
			continue
		}

		for _, tx := range block.Transactions() {
			processTransaction(ts, tx)
		}
		fmt.Println("Processed Block Number: ", block.Number().Uint64())
	}
}

func processTransaction(ts *transaction.TransactionService, tx *types.Transaction) {
	if tx.To() == nil {
		return
	}

	from := common.GetFromAddress(tx)
	gas, nonce := tx.Gas(), tx.Nonce()

	transactionArg := &common.Transaction{
		From:                 from,
		To:                   common.BytesToAddress([]byte(tx.To().Hex())),
		Gas:                  &gas,
		GasPrice:             tx.GasPrice(),
		MaxFeePerGas:         tx.GasFeeCap(),
		MaxPriorityFeePerGas: tx.GasTipCap(),
		Value:                tx.Value(),
		Nonce:                &nonce,
	}

	fmt.Printf("Got a transaction! From: %s, To: %s\n", from, tx.To())
	ts.SaveTransaction(ts.GenerateTransaction(transactionArg))
}
