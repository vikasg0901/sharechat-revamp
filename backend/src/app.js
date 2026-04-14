'use strict';

const express = require('express');

const app = express();

app.use(express.json());

// GET /health — no auth required, checks service is alive
app.get('/health', (req, res) => {
  res.status(200).json({
    status: 'ok',
    timestamp: new Date().toISOString(),
  });
});

module.exports = app;
