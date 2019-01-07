const { createError } = require('apollo-errors');
const jwt = require('jsonwebtoken');
const config = require('../../config.json');

const InvalidCredentialsError = createError('BCB/InvalidCredentialsError', {
  message: 'Username or password was incorrect.'
});

const resolvers = {
  Mutation: {
    authenticate: (root, args, context, info) => {
      const { username, password } = args;

      if (
        !config.app.admins[username] ||
        config.app.admins[username] !== password
      ) {
        throw new InvalidCredentialsError();
      }

      const token = jwt.sign({ username }, config.app.jwtSecret, {
        expiresIn: '1d'
      });
      return { token };
    }
  }
};

module.exports = { resolvers };
