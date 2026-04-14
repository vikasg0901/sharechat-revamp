'use strict';

const { Router } = require('express');

/**
 * createHealthRouter returns an Express Router with a GET /health handler.
 *
 * startTime is injected rather than read from a package-level variable so
 * tests can supply an arbitrary start time without mutating shared state.
 *
 * @param {number} [startTime=Date.now()] - epoch ms when the process started
 * @returns {import('express').Router}
 */
function createHealthRouter(startTime = Date.now()) {
  const router = Router();

  router.get('/health', (req, res) => {
    res.status(200).json({
      status: 'ok',
      uptime_seconds: Math.floor((Date.now() - startTime) / 1000),
    });
  });

  return router;
}

module.exports = createHealthRouter;
