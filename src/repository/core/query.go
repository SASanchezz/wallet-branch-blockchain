package core

import (
	"context"
	"time"
	"wallet-branch-blockchain/src/logger"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Query struct {
	dbTx     neo4j.ManagedTransaction
	template string
	params   map[string]interface{}
	logPath  string
}

func (q Query) Run() []*neo4j.Record {
	ctx := context.Background()
	var records []*neo4j.Record

	start := time.Now()

	if result, err := q.dbTx.Run(ctx, q.template, q.params); err != nil {
		panic(err)
	} else if records, err = result.Collect(ctx); err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	logger := logger.Logger{Path: q.logPath}
	logger.LogInt64(int64(elapsed / time.Microsecond))

	return records
}
