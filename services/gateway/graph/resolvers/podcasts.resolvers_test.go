package resolvers

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mamalovesyou/claimclam/graphql/gen/model"
	"github.com/mamalovesyou/claimclam/internal/podcasts"
	"github.com/stretchr/testify/assert"
)

func Setup(t *testing.T) (context.Context, *queryResolver, *podcasts.MockPodcastClient) {
	ctrl := gomock.NewController(t)
	mockPodcastsClient := podcasts.NewMockPodcastClient(ctrl)
	resolver := &queryResolver{&Resolver{podcastsClient: mockPodcastsClient}}
	return context.Background(), resolver, mockPodcastsClient
}

func TestPodcastsResolver_Success(t *testing.T) {
	ctx, resolver, mockPodcastsClient := Setup(t)

	// Mock setup
	mockPodcastsClient.EXPECT().ListPodcasts(gomock.Any(), gomock.Any()).Return(&podcasts.ListPodcastsResponse{
		Podcasts:   []*model.Podcast{{ /* ...populate fields... */ }},
		TotalCount: 10,
	}, nil)

	// Test parameters
	search := "comedy"
	title := "example"
	category := "entertainment"
	page := 1
	limit := 5

	// Call resolver
	result, err := resolver.Podcasts(ctx, &search, &title, &category, &page, &limit)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, *result.PageInfo.TotalPages) // TotalPages should be 2 (10 total / 5 limit)
	// ...additional assertions...
}

func TestPodcastsResolver_Error(t *testing.T) {
	ctx, resolver, mockPodcastsClient := Setup(t)

	// Mock setup
	mockPodcastsClient.EXPECT().ListPodcasts(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

	// Test parameters
	search := "comedy"
	title := "example"
	category := "entertainment"
	page := 1
	limit := 5

	// Call resolver
	result, err := resolver.Podcasts(ctx, &search, &title, &category, &page, &limit)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
}
