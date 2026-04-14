'use strict';

const express = require('express');

const app = express();

app.use(express.json());

// GET /api/v7/ready
// Readiness probe — returns 200 when the server is ready to accept traffic.
// Does not require authentication per health-check convention.
app.get('/api/v7/ready', (_req, res) => {
  res.status(200).json({ status: 'ok' });
});

// Only start listening when this module is run directly, not when imported in tests.
if (require.main === module) {
  const PORT = process.env.PORT || 3000;
  const server = app.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}`);
  });

  process.on('SIGTERM', () => {
    server.close(() => {
      process.exit(0);
    });
  });
}

module.exports = app;
