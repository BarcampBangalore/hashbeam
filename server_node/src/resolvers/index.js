const merge = require('lodash.merge');
const { resolvers: authResolvers } = require('./auth');
const { resolvers: announcementsResolvers } = require('./announcements');
const { resolvers: tweetsResolvers } = require('./tweets');

const resolvers = merge(authResolvers, announcementsResolvers, tweetsResolvers);

module.exports = { resolvers };
