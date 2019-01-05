const { UnauthorizedError } = require('../lib/errors');

const resolvers = {
  Query: {
    announcements: async (root, args, context, info) => {
      const { db } = context;

      const announcements = await db
        .select('*')
        .from('announcements')
        .where(db.raw('DATE(`timestamp`) = CURDATE()'))
        .orderBy('timestamp', 'desc');

      return announcements.map(announcement => ({
        timestampISO8601: announcement.timestamp.toISOString(),
        message: announcement.message
      }));
    }
  },
  Mutation: {
    makeAnnouncement: async (root, args, context, info) => {
      if (!context.user) {
        throw new UnauthorizedError();
      }

      const timestamp = new Date();
      const { message } = args;

      await context.db('announcements').insert({ timestamp, message });

      const announcement = {
        timestampISO8601: timestamp.toISOString(),
        message
      };

      context.pubsub.publish('announcement', { newAnnouncement: announcement });

      await context.firebase
        .messaging()
        .sendToTopic(context.config.fcm.topicName, {
          notification: {
            body: message,
            icon: context.config.fcm.notificationIconUrl,
            clickAction: context.config.fcm.notificationClickedTargetUrl
          }
        });

      return announcement;
    },

    subscribeToNotifications: async (root, args, context, info) => {
      const { fcmToken } = args;

      await context.firebase
        .messaging()
        .subscribeToTopic([fcmToken], context.config.fcm.topicName);

      return true;
    }
  },
  Subscription: {
    newAnnouncement: {
      subscribe: (root, args, context, info) =>
        context.pubsub.asyncIterator('announcement')
    }
  }
};

module.exports = { resolvers };
