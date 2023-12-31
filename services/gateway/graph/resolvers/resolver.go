package resolvers

import "github.com/mamalovesyou/claimclam/internal/podcasts"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	podcastsClient podcasts.PodcastClient
}

func NewResolver(podcastsClt podcasts.PodcastClient) *Resolver {
	return &Resolver{
		podcastsClient: podcastsClt,
	}
}
