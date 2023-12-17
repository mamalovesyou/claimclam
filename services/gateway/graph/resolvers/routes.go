package resolvers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gorilla/mux"
	"github.com/mamalovesyou/claimclam/services/gateway/graph"
	"go.uber.org/zap"

	"github.com/99designs/gqlgen/graphql/playground"
)

func RegisterRoutes(router *mux.Router, resolver *Resolver, logger *zap.Logger) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql")).Methods("GET", "POST", "OPTIONS")
	router.Handle("/graphql", srv).Methods("GET", "POST", "OPTIONS")
}
