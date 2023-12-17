import Card, { CardProps } from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";
import { Podcast } from "@/gql/graphql";

export type PodcastCardProps = CardProps & {
  podcast: Podcast;
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
          // 16:9
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
          p: 2,
        }}
      >
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
    </Card>
  );
};
