package core

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

type queryBuilder struct {
	query Query
}

func NewQueryBuilder(dbTx neo4j.ManagedTransaction) *queryBuilder {
	return &queryBuilder{
		query: Query{
			dbTx: dbTx,
		},
	}
}

func (qb *queryBuilder) WithTemplate(template string) *queryBuilder {
	qb.query.template = template
	return qb
}

func (qb *queryBuilder) WithParams(params map[string]interface{}) *queryBuilder {
	qb.query.params = params
	return qb
}

func (qb *queryBuilder) WithLogPath(logPath string) *queryBuilder {
	qb.query.logPath = logPath
	return qb
}

func (qb *queryBuilder) Build() Query {
	return qb.query
}
