'use strict';

const express = require('express');
const { v4: uuidv4 } = require('uuid');
const routes = require('./routes/index');

const app = express();

// Attach a request ID to every response for traceability
app.use((_req, res, next) => {
  res.setHeader('X-Request-Id', uuidv4());
  next();
});

app.use(express.json());

// Mount all routes
app.use('/', routes);

// 404 catch-all
app.use((_req, res) => {
  res.status(404).json({
    error: { code: 'NOT_FOUND', message: 'The requested resource does not exist.' },
  });
});

// Central error handler
// eslint-disable-next-line no-unused-vars
app.use((err, _req, res, _next) => {
  console.error(err);
  res.status(500).json({
    error: { code: 'INTERNAL_ERROR', message: 'An unexpected error occurred.' },
  });
});

const PORT = process.env.PORT || 3000;

if (require.main === module) {
  const server = app.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}`);
  });

  // Graceful shutdown on SIGTERM
  process.on('SIGTERM', () => {
    server.close(() => {
      console.log('Server shut down gracefully.');
      process.exit(0);
    });
  });
}

module.exports = app;
