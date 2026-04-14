const { Router } = require('express');

const router = Router();

/**
 * GET /ping
 * Health check — no auth required, no downstream checks.
 * Returns 200 with a JSON envelope so clients can confirm the API is reachable.
 */
router.get('/', (req, res) => {
  res.status(200).json({ status: 'ok', message: 'pong' });
});

module.exports = router;
