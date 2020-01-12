feed = `[{"Channel":"Q-Dance-Youtube","ChannelURL":"https://www.youtube.com/user/Qdancedotnl","ChannelImageURL":"https://yt3.ggpht.com/a-/AN66SAzyW12uAQRayPY4MS_Fo_Wlj6PFjyNfx3X7CQ=s288-mo-c-c0xffffffff-rj-k-no","PlaylistToDownloadURL":"https://www.youtube.com/user/Qdancedotnl/videos","Items":null},{"Channel":"Q-Dance-Mixcloud","ChannelURL":"https://mixcloud.com/Q-Dance","ChannelImageURL":"https://thumbnailer.mixcloud.com/unsafe/320x320/profile/d/a/8/8/7f54-2819-4c09-9d74-475a2771a834","PlaylistToDownloadURL":"https://mixcloud.com/Q-Dance","Items":null},{"Channel":"B2S-Youtube","ChannelURL":"https://www.youtube.com/user/officialb2s","ChannelImageURL":"https://yt3.ggpht.com/a-/AN66SAzOHiotaZ20EWS6f9SHfzDPHQeAR_gyn-ng9w=s288-mo-c-c0xffffffff-rj-k-no","PlaylistToDownloadURL":"https://www.youtube.com/user/officialb2s/videos","Items":null},{"Channel":"Art-of-Dance-Youtube","ChannelURL":"https://www.youtube.com/user/officialartofdance","ChannelImageURL":"https://yt3.ggpht.com/a/AGF-l79pcH8ynFLsFDpJ7pn1RWs6wMAXWfQY5OLlXQ=s288-c-k-c0xffffffff-no-rj-mo","PlaylistToDownloadURL":"https://www.youtube.com/user/officialartofdance/videos","Items":null}]`;

// podcastConfig = [
//   {
//     channel,
//     channelURL,
//     channelImageURL,
//     playListURL
//   }
// ];

js = JSON.parse(feed);

for (i in js) {
  console.log(js[i]);
}

