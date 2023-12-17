import { graphql } from "@/gql";
import { PageInfo, Podcast } from "@/gql/graphql";
import { useQuery } from "@tanstack/react-query";
import request from "graphql-request";
import React, { useCallback, useContext, useEffect, useState } from "react";

const listPodcastsQuery = graphql(/* GraphQL */ `
  query listPodcasts(
    $search: String
    $category: String
    $title: String
    $limit: Int
    $page: Int
  ) {
    podcasts(
      search: $search
      title: $title
      category: $category
      limit: $limit
      page: $page
    ) {
      items {
        id
        title
        description
        images {
          thumbnail
        }
        category
      }
      pageInfo {
        currentPage
        totalPages
      }
    }
  }
`);

export interface PodcastsFilters {
  search?: string;
  category?: string;
  title?: string;
  limit?: number;
  page?: number;
}

const useListPodcasts = (filters: PodcastsFilters) => {
  console.log("received filters", filters);
  return useQuery({
    queryKey: ["podcasts", filters],
    queryFn: async () =>
      request(
        process.env.NEXT_PUBLIC_GRAPHQL_URL || "http://localhost:3001/graphql",
        listPodcastsQuery,
        filters
      ),
  });
};

type PodcastsContextType = {
  podcasts: Array<Podcast | null>;
  pageInfo: PageInfo | null;
  onFilterUpdate: (filters: PodcastsFilters) => void;
  onPageChange: (page: number) => void;
  listPodcasts: ReturnType<typeof useListPodcasts> | null;
};

const PodcastsContext = React.createContext<PodcastsContextType>({
  podcasts: [],
  pageInfo: null,
  onFilterUpdate: () => {},
  onPageChange: () => {},
  listPodcasts: null,
});

export const PodcastsProvider = (props: { children: React.ReactNode }) => {
  const [filters, setFilters] = useState<PodcastsFilters>({
    search: "",
    page: 1,
    limit: 10,
  });
  const listPodcasts = useListPodcasts(filters);

  // Update filters and keep current page
  const onFilterUpdate = useCallback((value: PodcastsFilters) => {
    setFilters((prev) => {
      // If limit is updated then reset page to 1
      if (value.limit !== prev.limit) {
        return {
          ...value,
          page: 1,
        };
      }
      return { ...value, page: prev.page };
    });
  }, []);

  // Update page and keep current filters
  const onPageChange = useCallback((page: number) => {
    setFilters((prev) => ({ ...prev, page }));
  }, []);

  return (
    <PodcastsContext.Provider
      value={{
        podcasts: listPodcasts.data?.podcasts?.items || [],
        pageInfo: listPodcasts.data?.podcasts?.pageInfo || null,
        onFilterUpdate,
        onPageChange,
        listPodcasts,
      }}
    >
      {props.children}
    </PodcastsContext.Provider>
  );
};

export const usePodcasts = (): PodcastsContextType => {
  const context = useContext(PodcastsContext);
  return context;
};
