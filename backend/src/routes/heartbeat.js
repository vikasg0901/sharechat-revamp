/**
 * GET /v1/heartbeat
 * Returns health status and uptime information.
 */
const heartbeat = (req, res) => {
  const payload = {
    status: 'ok',
    uptimeSeconds: Math.floor(process.uptime()),
    timestamp: new Date().toISOString(),
  };
  res.json(payload);
};

module.exports = { heartbeat };
