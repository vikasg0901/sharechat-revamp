'use strict';

const pkg = require('../package.json');

/**
 * GET /v1/version
 * Returns the running application version sourced from package.json.
 * No authentication required.
 */
function getVersion(req, res) {
  res.status(200).json({
    name: pkg.name,
    version: pkg.version,
  });
}

module.exports = { getVersion };
