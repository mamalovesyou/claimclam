package resolvers

import "github.com/mamalovesyou/getclaim/internal/podcasts"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	podcastsClient *podcasts.Client
}

func NewResolver(podcastsClt *podcasts.Client) *Resolver {
	return &Resolver{
		podcastsClient: podcastsClt,
	}
}
