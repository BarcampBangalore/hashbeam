const merge = require('lodash.merge');
const { resolvers: authResolvers } = require('./auth');

const resolvers = merge(authResolvers);

module.exports = { resolvers };
