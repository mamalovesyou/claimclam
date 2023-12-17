"use client";

import { PodcastListFilters } from "@/components/podcasts/ProdcastListFilters";
import { Card, Container } from "@mui/material";
import { PodcastsProvider } from "@/hooks/usePodcasts";
import { PodcastList } from "@/components/podcasts/PodcastList";

export default function Home() {
  return (
    <PodcastsProvider>
      <Container sx={{ py: 8 }} maxWidth="lg" component="main">
        <Card>
          <PodcastListFilters />
        </Card>
        <PodcastList />
      </Container>
    </PodcastsProvider>
  );
}
