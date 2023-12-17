import * as React from "react";
import Button from "@mui/material/Button";
import Dialog, { DialogProps } from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import { Podcast } from "@/gql/graphql";
import { Box, Chip, Typography } from "@mui/material";
import { usePodcasts } from "@/hooks/usePodcasts";

type PodcastDialogProps = DialogProps & {
  podcastId?: string;
};

export default function PodcastDialog(props: PodcastDialogProps) {
  const { podcastId, ...dialogProps } = props;
  const { podcasts } = usePodcasts();

  const podcast = podcasts.find((p) => p?.id === podcastId);
  console.log(podcast);

  return (
    <Dialog {...dialogProps} sx={{ p: 0 }}>
      <Box
        component="img"
        src={
          podcast?.images?.wide ||
          "https://source.unsplash.com/random?wallpapers"
        }
        sx={{ width: "100%" }}
      />

      <DialogTitle>
        {podcast?.title}
        <Typography variant="body2" sx={{ color: "text.secondary" }}>
          by {podcast?.publisherName}
        </Typography>
      </DialogTitle>
      <DialogContent>
        <DialogContentText>{podcast?.description}</DialogContentText>
        <Box sx={{ py: 2 }}>
          {podcast?.categoryName ? (
            <Chip label={podcast?.categoryName} color="primary" />
          ) : null}
        </Box>
      </DialogContent>
      <DialogActions sx={{ justifyContent: "flex-end" }}>
        <Button
          onClick={(e) =>
            dialogProps.onClose && dialogProps.onClose(e, "escapeKeyDown")
          }
          autoFocus
        >
          Back
        </Button>
      </DialogActions>
    </Dialog>
  );
}
