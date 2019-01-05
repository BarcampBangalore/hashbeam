const { GraphQLServer, PubSub } = require('graphql-yoga');
const { formatError } = require('./lib/errors');
const { setupTables } = require('./lib/setup-tables');
const process = require('process');
const { authMiddleware } = require('./lib/auth-middleware');
const { resolvers } = require('./resolvers');
const config = require('../config.json');
const knex = require('knex');

const main = async () => {
  const db = knex({
    client: 'mysql2',
    connection: config.mySql,
    debug: process.env.NODE_ENV !== 'production'
  });

  await setupTables(db);

  const pubsub = new PubSub();

  const server = new GraphQLServer({
    typeDefs: 'src/schema.graphql',
    resolvers,
    context: params => ({ ...params, db, pubsub }),
    middlewares: [authMiddleware]
  });

  server.start({ formatError, port: config.app.port || 3000 }, () => {
    console.log(`Server is running on port ${config.app.port || 3000}`);
  });
};

main();
