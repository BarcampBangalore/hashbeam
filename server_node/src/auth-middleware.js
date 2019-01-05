const { createError } = require('apollo-errors');
const jwt = require('jsonwebtoken');
const config = require('../config.json');

const InvalidTokenError = createError('BCB/InvalidTokenError', {
  message: 'Token in payload was not valid.'
});

const authMiddleware = (resolve, parent, args, context, info) => {
  const token = context.request.get('Authorization');

  if (!token) {
    context = { ...context, user: null };
    return resolve(parent, args, context, info);
  }

  try {
    const user = jwt.verify(token, config.jwtSecret);
    context = { ...context, user };
    return resolve(parent, args, context, info);
  } catch (err) {
    return new InvalidTokenError();
  }
};

module.exports = { authMiddleware };
