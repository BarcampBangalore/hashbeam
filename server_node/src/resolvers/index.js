const merge = require('lodash.merge');
const { resolvers: authResolvers } = require('./auth');
const { resolvers: announcementsResolvers } = require('./announcements');

const resolvers = merge(authResolvers, announcementsResolvers);

module.exports = { resolvers };
