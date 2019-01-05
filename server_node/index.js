const { createError } = require('apollo-errors');

const MyError = createError('MyError', { message: 'My awesome error' });

const myError = new MyError();

console.log(myError instanceof MyError);
