package faker

import (
	"github.com/wundergraph/graphql-go-tools/execution/graphql"
)

type GraphqlFaker struct {
}

func (g *GraphqlFaker) Fake() string {
	var gqlRequest graphql.Request
}
