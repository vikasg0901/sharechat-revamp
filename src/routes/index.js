const { Router } = require('express');

const router = Router();

// GET /api/v7/ready — readiness probe; returns 200 once the server is accepting requests
router.get('/v7/ready', (req, res) => {
  res.status(200).json({ ready: true });
});

module.exports = router;
