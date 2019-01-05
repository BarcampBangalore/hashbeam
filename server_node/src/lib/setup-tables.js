const knex = require('knex');
const fs = require('fs').promises;
const path = require('path');

/**
 * @param {knex} db
 */
const setupTables = async db => {
  const schemaFile = await fs.readFile(path.join(__dirname, 'schema.sql'));
  const sqlString = schemaFile.toString();
  const sqlStatements = sqlString.split(';\n');

  for (const statement of sqlStatements) {
    await db.raw(statement);
  }
};

module.exports = { setupTables };
