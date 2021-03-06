const twit = require('twit');

class Twitter {
  constructor(config, db) {
    this.client = new twit(config);
    this.config = config;
    this.db = db;
    this.whitelistedUserIds = new Set();
    this.mutedUserIds = new Set();
  }

  async updateMutedUserIds() {
    const rows = await this.db.select('twitter_user_id').from('muted_tweeters');
    const mutedUserIds = rows.map(row => row.twitter_user_id);
    this.mutedUserIds = new Set(mutedUserIds);
  }

  async updateWhitelistedUserIds() {
    const rows = await this.db
      .select('twitter_user_id')
      .from('whitelisted_tweeters');

    const whitelistedUserIds = rows.map(row => row.twitter_user_id);
    this.whitelistedUserIds = new Set(whitelistedUserIds);
  }

  async startStream() {
    await Promise.all([
      this.updateMutedUserIds(),
      this.updateWhitelistedUserIds()
    ]);

    const stream = this.client.stream('statuses/filter', {
      track: this.config.textToTrack
    });

    stream.on('tweet', async tweet => {
      if (this.mutedUserIds.has(tweet.user.id_str)) {
        return;
      }

      try {
        await this.db('tweets').insert({
          tweet_id: tweet.id_str,
          timestamp: new Date(tweet.created_at),
          original_tweet_json: JSON.stringify(tweet),
          author_id: tweet.user.id_str,
          review_required: !this.whitelistedUserIds.has(tweet.user.id_str)
        });
      } catch (err) {
        console.error('Failed to save tweet to database', err);
      }
    });

    stream.on('error', err => {
      console.error('Twitter stream error', err);
    });
  }
}

module.exports = { Twitter };
