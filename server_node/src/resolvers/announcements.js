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
      return announcement;
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
