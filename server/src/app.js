const express = require('express');
const routes = require('./routes');

const app = express();

app.use(express.json());

// Mount all versioned API routes
app.use('/v1', routes);

// Catch-all 404 for unknown routes
app.use((req, res) => {
  res.status(404).json({ error: { code: 'NOT_FOUND', message: 'Route not found' } });
});

module.exports = app;
