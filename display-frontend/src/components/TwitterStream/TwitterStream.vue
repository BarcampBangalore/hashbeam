<template>
  <div :class="$style.grid">
    <twitter-header :trackText="trackText"></twitter-header>
    <template v-if="tweets[0]">
      <tweet-author :user="tweets[0].user"></tweet-author>
      <image-tweet-body v-if="isImageTweet" :tweet="tweets[0]"></image-tweet-body>
      <text-tweet-body v-else :tweet="tweets[0]"></text-tweet-body>
    </template>
  </div>
</template>

<script>
import TwitterHeader from "./TwitterHeader";
import gql from "graphql-tag";
import TweetAuthor from "./TweetAuthor";
import TextTweetBody from "./TextTweetBody";
import ImageTweetBody from "./ImageTweetBody";

export default {
  components: {
    TwitterHeader,
    TweetAuthor,
    TextTweetBody,
    ImageTweetBody
  },

  data() {
    return {
      trackText: "",
      tweets: []
    };
  },

  watch: {
    async tweets(tweets) {
      if (tweets.length <= 1) {
        const response = await this.$apollo.provider.defaultClient.query({
          query: gql`
            query {
              tweetsToDisplay {
                originalTweetJson
              }
            }
          `,
          fetchPolicy: "network-only"
        });

        const tweets = response.data.tweetsToDisplay.map(tweet =>
          JSON.parse(tweet.originalTweetJson)
        );
        if (tweets.length > 0) {
          this.tweets.push(...tweets);
        }
      }
    }
  },

  computed: {
    isImageTweet() {
      return (
        this.tweets[0].extended_entities ||
        (this.tweets[0].quoted_status &&
          this.tweets[0].quoted_status.extended_entities)
      );
    }
  },

  async created() {
    const response = await this.$apollo.provider.defaultClient.query({
      query: gql`
        query {
          tweetsToDisplay {
            originalTweetJson
          }
        }
      `,
      fetchPolicy: "network-only"
    });

    const tweets = response.data.tweetsToDisplay.map(tweet =>
      JSON.parse(tweet.originalTweetJson)
    );
    this.tweets = tweets;
    setInterval(() => this.tweets.splice(0, 1), 10000);
  },

  sockets: {
    tweet(tweet) {
      this.tweets.push(tweet);
    }
  }
};
</script>

<style module>
.grid {
  display: grid;
  height: inherit;
  grid-template-rows: 1fr 1fr 5fr;
}
</style>
