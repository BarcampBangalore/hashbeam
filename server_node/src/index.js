const { GraphQLServer } = require('graphql-yoga');
const { formatError } = require('./lib/errors');
const { authMiddleware } = require('./auth-middleware');
const { resolvers } = require('./resolvers');
const config = require('../config.json');
const knex = require('knex');

const db = knex({
  client: 'mysql2',
  connection: config.mySql
});

const server = new GraphQLServer({
  typeDefs: 'src/schema.graphql',
  resolvers,
  context: req => ({ ...req, db }),
  middlewares: [authMiddleware]
});

server.start({ formatError, port: config.app.port || 3000 }, () =>
  console.log(`Server is running on port ${config.app.port || 3000}`)
);
