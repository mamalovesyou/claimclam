import Card, { CardProps } from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";
import { Podcast } from "@/gql/graphql";
import { Box, Button, CardActions, CardHeader } from "@mui/material";

export type PodcastCardProps = CardProps & {
  podcast: Podcast;
  onPodcastClick: (podcastId: string) => void;
};

export const PodcastCard = (props: PodcastCardProps) => {
  const { podcast, ...cardProps } = props;
  return (
    <Card
      sx={{
        height: "100%",
        display: "flex",
        flexDirection: "column",
      }}
      {...cardProps}
    >
      <CardMedia
        component="div"
        sx={{
          pt: "56.25%",
        }}
        image={
          podcast?.images?.thumbnail ||
          "https://source.unsplash.com/random?wallpapers"
        }
      />
      <CardContent
        sx={{
          flexGrow: 1,
          textOverflow: "ellipsis",
        }}
      >
        <Box sx={{ pb: 2 }}>
          <Typography
            variant="h5"
            sx={{
              overflow: "hidden",
              textOverflow: "ellipsis",
              display: "-webkit-box",
              WebkitLineClamp: "1",
              WebkitBoxOrient: "vertical",
            }}
          >
            {podcast?.title}
          </Typography>
          <Typography variant="body2" sx={{ color: "text.secondary" }}>
            by {podcast?.publisherName}
          </Typography>
        </Box>
        <Typography
          sx={{
            overflow: "hidden",
            textOverflow: "ellipsis",
            display: "-webkit-box",
            WebkitLineClamp: "4",
            WebkitBoxOrient: "vertical",
          }}
        >
          {podcast?.description}
        </Typography>
      </CardContent>
      <CardActions sx={{ p: 1, justifyContent: "center" }}>
        <Button
          variant="contained"
          size="small"
          sx={{ p: 1 }}
          onClick={() => props.onPodcastClick(podcast?.id || "")}
        >
          View
        </Button>
      </CardActions>
    </Card>
  );
};
