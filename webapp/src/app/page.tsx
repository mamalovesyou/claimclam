"use client";

import { PodcastListFilters } from "@/components/podcasts/ProdcastListFilters";
import { Card, Container } from "@mui/material";
import { PodcastsProvider, usePodcasts } from "@/hooks/usePodcasts";
import { PodcastList } from "@/components/podcasts/PodcastList";
import PodcastDialog from "@/components/podcasts/PodcastDialog";
import { useState } from "react";
import { Podcast } from "@/gql/graphql";

export default function Home() {
  const [selectedPodcastId, setSelectedPodcastId] = useState<
    string | undefined
  >(undefined);
  const dialogOpen = selectedPodcastId !== null;

  const handlePodcastClick = (podcastId: string) => {
    console.log("podcastId", podcastId);
    setSelectedPodcastId(podcastId);
  };

  const handleDialogClose = () => {
    setSelectedPodcastId(undefined);
  };

  return (
    <PodcastsProvider>
      <Container sx={{ py: 8 }} maxWidth="lg" component="main">
        <Card>
          <PodcastListFilters />
        </Card>
        <PodcastList onPodcastClick={handlePodcastClick} />
      </Container>
      <PodcastDialog
        podcastId={selectedPodcastId}
        open={selectedPodcastId !== undefined}
        onClose={handleDialogClose}
      />
    </PodcastsProvider>
  );
}
