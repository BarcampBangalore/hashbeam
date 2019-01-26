<template>
  <div :class="$style.container">
    <!--If this is a retweet or quoting another tweet-->
    <template v-if="retweetedUser">
      <h2 :class="$style.retweeted">Retweeted @{{ retweetedUser }}</h2>
    </template>

    <div
      :class="$style.image"
      :style="{ background: `${textGradient}, url(${imageUrl}), no-repeat`, backgroundPosition: 'center', backgroundSize: 'cover'}"
    >
      <h1 :class="$style.text">{{ tweetText }}</h1>
    </div>
  </div>
</template>

<script>
import unescapeHTML from "../../helpers/unescapeHTML";

export default {
  props: ["tweet"],

  data() {
    return {
      textGradient:
        "linear-gradient(to bottom, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0) 60%, rgba(0, 0, 0, 0.9) 100%)"
    };
  },

  computed: {
    retweetedUser() {
      if (this.tweet.quoted_status) {
        return this.tweet.quoted_status.user.screen_name;
      } else if (this.tweet.retweeted_status) {
        return this.tweet.retweeted_status.user.screen_name;
      }

      return null;
    },

    tweetText() {
      let tweetText = "";
      if (this.tweet.quoted_status) {
        tweetText = this.tweet.text || this.tweet.quoted_status.text;
      } else if (this.tweet.retweeted_status) {
        tweetText = this.tweet.retweeted_status.text;
      } else {
        tweetText = this.tweet.text || "";
      }

      return unescapeHTML(tweetText).slice(0, 279);
    },

    imageUrl() {
      if (this.tweet.quoted_status) {
        return this.tweet.quoted_status.extended_entities.media[0]
          .media_url_https;
      } else if (this.tweet.retweeted_status) {
        return this.tweet.retweeted_status.extended_entities.media[0]
          .media_url_https;
      }
      return this.tweet.extended_entities.media[0].media_url_https;
    }
  }
};
</script>

<style module>
.container {
  display: flex;
  flex-direction: column;
  margin: 2% 5% 5% 5%;
}

.retweeted {
  color: rgba(255, 255, 255, 0.7);
  font-size: calc((2.5vh + 2.5vw) / 2);
  margin-bottom: 1em;
}

.image {
  display: flex;
  border-radius: 20px;
  height: 100%;
}

.text {
  align-self: flex-end;
  margin: 0 2% 2% 2%;
  color: white;
  font-size: calc((2vh + 2vw) / 2);
  line-height: 1.2em;
}
</style>
