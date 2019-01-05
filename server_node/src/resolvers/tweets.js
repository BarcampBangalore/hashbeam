const { UnauthorizedError } = require('../lib/errors');

const formatTweetForGraphQL = tweet => ({
  tweetId: tweet.tweet_id,
  authorId: tweet.author_id,
  timestampISO8601: tweet.timestamp.toISOString(),
  originalTweetJson: JSON.stringify(tweet.original_tweet_json),
  reviewRequired: tweet.review_required
});

const resolvers = {
  Query: {
    tweetsToDisplay: async (root, args, context, info) => {
      const tweets = await context.db
        .select('*')
        .from('tweets')
        .where({ review_required: false })
        .orderBy('timestamp', 'desc')
        .limit(10);

      return tweets.map(formatTweetForGraphQL);
    },
    tweetToReview: async (root, args, context, info) => {
      if (!context.user) {
        throw new UnauthorizedError();
      }

      const [tweet] = await context.db
        .select('*')
        .from('tweets')
        .where({ review_required: true })
        .orderBy('timestamp', 'asc')
        .limit(1);

      if (tweet) {
        return formatTweetForGraphQL(tweet);
      }

      return null;
    }
  },
  Mutation: {
    whitelistTweeter: async (root, args, context, info) => {
      if (!context.user) {
        throw new UnauthorizedError();
      }

      const { twitterAuthorID } = args;

      await context.db.transaction(async tx => {
        await tx('tweets')
          .update({ review_required: false })
          .where({ author_id: twitterAuthorID });

        await tx('whitelisted_tweeters').insert({
          twitter_user_id: twitterAuthorID
        });

        await tx.commit();
      });

      await context.twitter.updateWhitelistedUserIds();
      return true;
    },
    muteTweeter: async (root, args, context, info) => {
      if (!context.user) {
        throw new UnauthorizedError();
      }

      const { twitterAuthorID } = args;

      await context.db.transaction(async tx => {
        await tx('tweets')
          .delete()
          .where({ author_id: twitterAuthorID });

        await tx('muted_tweeters').insert({ twitter_user_id: twitterAuthorID });
        await tx.commit();
      });

      await context.twitter.updateMutedUserIds();
      return true;
    }
  }
};

module.exports = { resolvers };
