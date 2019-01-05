const { GraphQLServer, PubSub } = require('graphql-yoga');
const { formatError } = require('./lib/errors');
const { setupTables } = require('./lib/setup-tables');
const process = require('process');
const Firebase = require('firebase-admin');
const { authMiddleware } = require('./lib/auth-middleware');
const { resolvers } = require('./resolvers');
const config = require('../config.json');
const firebaseServiceKeyJson = require('../firebase-service-key.json');
const knex = require('knex');

const main = async () => {
  const db = knex({
    client: 'mysql2',
    connection: config.mySql,
    debug: process.env.NODE_ENV !== 'production'
  });

  await setupTables(db);

  const pubsub = new PubSub();

  const firebase = Firebase.initializeApp({
    credential: Firebase.credential.cert(firebaseServiceKeyJson)
  });

  const server = new GraphQLServer({
    typeDefs: 'src/schema.graphql',
    resolvers,
    context: params => ({ ...params, config, db, pubsub, firebase }),
    middlewares: [authMiddleware]
  });

  server.start({ formatError, port: config.app.port || 3000 }, () => {
    console.log(`Server is running on port ${config.app.port || 3000}`);
  });
};

main();
