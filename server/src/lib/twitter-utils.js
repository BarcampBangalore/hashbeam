exports.hasVideo = function hasVideo(tweet) {
  return (
    (tweet.extended_entities &&
      tweet.extended_entities.media &&
      tweet.extended_entities.media.some(media => media.type === 'video')) ||
    (tweet.retweeted_tweet &&
      tweet.retweeted_tweet.extended_tweet &&
      tweet.retweeted_tweet.extended_tweet.extended_entities &&
      tweet.retweeted_tweet.extended_tweet.extended_entities.media &&
      tweet.retweeted_tweet.extended_tweet.extended_entities.media.some(
        media => media.type === 'video'
      ))
  );
};
