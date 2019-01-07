const {
  createError,
  formatError: apolloFormatError
} = require('apollo-errors');

const InternalServerError = createError('BCB/InternalServerError', {
  message: 'Something went wrong and Adithya was too stupid to handle it.'
});

const UnauthorizedError = createError('BCB/UnauthorizedError', {
  message: "You aren't allowed to do this."
});

const formatError = (...args) => {
  const [error] = args;
  if (error.originalError.name.startsWith('BCB/')) {
    return apolloFormatError(...args);
  }

  console.error(error);
  return apolloFormatError(new InternalServerError());
};

module.exports = {
  formatError,
  UnauthorizedError
};
