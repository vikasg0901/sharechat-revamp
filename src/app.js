const express = require('express');
const path = require('path');
const routes = require('./routes');

function createApp() {
  const app = express();

  app.use(express.json());

  // Serve static frontend files from project root
  app.use(express.static(path.join(__dirname, '..')));

  // API routes
  app.use('/api', routes);

  return app;
}

module.exports = { createApp };
