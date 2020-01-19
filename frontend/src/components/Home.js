import React, { useState, useEffect } from "react";
import axios from "axios";
import { Box, Image, Badge, StarIcon, Flex } from "@chakra-ui/core";

const Podcast = ({
  Channel,
  ChannelURL,
  ChannelImageURL,
  PlaylistToDownloadURL
}) => (
  <Box
    m={8}
    flex="1 0 200px"
    borderWidth="1px"
    borderStyle="solid"
    borderColor="gray.200"
    rounded="lg"
    overflow="hidden"
  >
    <Image w="100%" src={ChannelImageURL} alt={Channel + " Logo"} />
    <Box p="6">
      <Box d="flex" alignItems="baseline">
        <Badge rounded="full" px="2" variantColor="teal">
          New
        </Badge>
      </Box>

      <Box mt="1" fontWeight="semibold" as="h4" lineHeight="tight" isTruncated>
        {Channel}
      </Box>

      <Box></Box>
    </Box>
  </Box>
);

const Home = () => {
  const [podcasts, setPodcasts] = useState([]);
  const [error, setError] = useState("");

  useEffect(() => {
    axios
      .get("/podcasts")
      .then((resp) => setPodcasts(resp.data))
      .catch((e) => setError(e.message));
  }, []);

  console.log(podcasts);
  return (
    <div>
      <h1>Podcasts</h1>
      <Flex align="center" justify="center" flexWrap="wrap">
        {podcasts.map((podcast) => (
          <Podcast key={podcast.ChannelURL} {...podcast} />
        ))}
      </Flex>
    </div>
  );
};

export default Home;
