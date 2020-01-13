/** Make call to /availablePodcasts to get all podcasts
    then iterate over all, fetch the xml feed and use the data needed to display the available podcasts nicely
    infos to display:
        channel.Title
        channel.image.url
        URL to subscribe to podcast

    this file has to be transformed using browserify to be abtle to use `require` on the client
 */

const axios = require("axios").default;

function parsePodcastInfofromXML(data, podcast) {
  parser = new DOMParser();
  xmlDoc = parser.parseFromString(data, "text/xml");
  podcastInfo = {};

  podcastInfo.title = xmlDoc.getElementsByTagName(
    "title"
  )[0].childNodes[0].nodeValue;
  podcastInfo.imageURL = xmlDoc.getElementsByTagName(
    "url"
  )[0].childNodes[0].nodeValue;
  podcastInfo.subscriptionURL =
    "http://" + window.location.host + "/podcasts" + podcast;

  return podcastInfo;
}

function getPodcastXMLFeed(podcast) {
  axios
    .get("/podcasts/" + podcast)
    .then(function(response) {
      podcastInfo = parsePodcastInfofromXML(response.data, podcast);
      displayPodcastsInHtml(podcastInfo);
    })
    .catch(function(error) {
      console.log(error);
    });
}

function displayPodcastsInHtml(podcastInfo) {
  document.getElementById("podcastList").innerHTML += `<li class="media">
  <img src=${podcastInfo.imageURL} class="align-self-center mr-3 img-fluid" alt="...">
  <div class="media-body">
    <h5 class="mt-0">${podcastInfo.title}</h5>
    <h7>Subscribe at:</h7>
    <p>${podcastInfo.subscriptionURL}</p>
  </div>
</li>`;
}

function showAvailablePodcasts() {
  axios
    .get("/availablePodcasts")
    .then(function(response) {
      // handle success
      response.data.forEach(podcast => {
        getPodcastXMLFeed(podcast);
      });
    })
    .catch(function(error) {
      // handle error
      console.log(error);
    });
}

showAvailablePodcasts();
