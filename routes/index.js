'use strict';

const { Router } = require('express');
const { getVersion } = require('../handlers/version');

const router = Router();

// Health check — unauthenticated, used by load balancers and uptime monitors
router.get('/healthz', (_req, res) => {
  res.status(200).json({ status: 'ok' });
});

// v1 routes
router.get('/v1/version', getVersion);

module.exports = router;
