import { usePodcasts } from "@/hooks/usePodcasts";
import { CircularProgress, Grid, Pagination, Typography } from "@mui/material";
import { PodcastCard } from "./PodcastCard";

type PodcastListProps = {
  onPodcastClick: (id: string) => void;
};

export const PodcastList = (props: PodcastListProps) => {
  const { onPodcastClick } = props;
  const { podcasts, pageInfo, onPageChange, listPodcasts } = usePodcasts();
  const { isFetching } = listPodcasts || {};
  return (
    <Grid
      container
      spacing={4}
      pt={3}
      sx={{ display: "flex", justifyContent: "center" }}
    >
      {isFetching && (
        <Grid
          item
          xs={12}
          sx={{ p: 3, display: "flex", justifyContent: "center" }}
        >
          <CircularProgress />
        </Grid>
      )}
      {podcasts.length > 0 &&
        podcasts.map((podcast) => (
          <Grid
            item
            key={podcast?.id}
            xs={12}
            md={4}
            lg={3}
            sx={{ display: "flex", justifyContent: "center" }}
          >
            {podcast ? (
              <PodcastCard podcast={podcast} onPodcastClick={onPodcastClick} />
            ) : null}
          </Grid>
        ))}
      {podcasts.length === 0 && (
        <Grid
          item
          xs={12}
          sx={{ p: 3, display: "flex", justifyContent: "center" }}
        >
          <Typography variant="body1">No podcasts found</Typography>
        </Grid>
      )}
      <Grid item xs={12} sx={{ display: "flex", justifyContent: "center" }}>
        <Pagination
          count={pageInfo?.totalPages || 1}
          page={pageInfo?.currentPage || 1}
          onChange={(_, page) => onPageChange(page)}
        />
      </Grid>
    </Grid>
  );
};
