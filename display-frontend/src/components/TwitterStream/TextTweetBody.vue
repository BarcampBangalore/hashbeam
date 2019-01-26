<template>
  <div :class="$style.container">

    <!--If this is a retweet-->
    <template v-if="tweet.retweeted_status">
      <h2 :class="$style.retweeted">
        Retweeted @{{ tweet.retweeted_status.user.screen_name }}
      </h2>
      <h1 :class="$style.text">
        {{ retweetedStatusText(tweet) }}
      </h1>
    </template>

    <!--If it's not a retweet-->
    <template v-else :class="$style.text">
      <h1 :class="$style.text">
        {{ tweetText(tweet) }}
      </h1>
    </template>

  </div>
</template>

<script>
import unescapeHTML from '../../helpers/unescapeHTML';

export default {
  props: ['tweet'],

  methods: {
    tweetText(tweet) {
      if (tweet.extended_tweet) {
        return unescapeHTML(tweet.extended_tweet.full_text);
      }

      return unescapeHTML(tweet.text);
    },

    retweetedStatusText(tweet) {
      if (tweet.retweeted_status.extended_tweet) {
        return unescapeHTML(tweet.retweeted_status.extended_tweet.full_text);
      }

      return unescapeHTML(tweet.retweeted_status.text);
    }
  }
};
</script>

<style module>
.container {
  display: flex;
  flex-direction: column;
  margin: 5% 8% 5% 8%;
}

.retweeted {
  color: rgba(255, 255, 255, 0.7);
  font-size: calc((3vh + 3vw) / 2);
  margin-bottom: 1em;
}

.text {
  color: white;
  font-size: calc((3vh + 3vw) / 2);
  line-height: 1.5em;
}
</style>
