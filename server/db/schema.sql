CREATE TABLE IF NOT EXISTS `tweets`
(
  `tweet_id`            varchar(45) NOT NULL,
  `datetime`            datetime    NOT NULL,
  `original_tweet_json` json        NOT NULL,
  `author_id`           varchar(45) NOT NULL,
  `review_required`     boolean     NOT NULL,
  PRIMARY KEY (`tweet_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `announcements`
(
  `id`       int(11) unsigned NOT NULL AUTO_INCREMENT,
  `datetime` datetime         NOT NULL,
  `message`  text             NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `whitelisted_tweeters`
(
  `twitter_user_id` varchar(45) NOT NULL,
  PRIMARY KEY (`twitter_user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;