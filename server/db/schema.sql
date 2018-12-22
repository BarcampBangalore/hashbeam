CREATE TABLE IF NOT EXISTS `announcements`
(
  `id`       int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `datetime` DATETIME         NOT NULL,
  `message`  TEXT             NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `tweets`
(
  `tweet_id`        VARCHAR(45) NOT NULL,
  `datetime`        DATETIME    NOT NULL,
  `content`         JSON        NOT NULL,
  `author_id`       VARCHAR(45) NOT NULL,
  `review_required` BOOLEAN     NOT NULL,
  PRIMARY KEY (`tweet_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `whitelisted_tweeters`
(
  `twitter_user_id` VARCHAR(45) NOT NULL,
  `display_name`    VARCHAR(45) NOT NULL,
  PRIMARY KEY (`twitter_user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;