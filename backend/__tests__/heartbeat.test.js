const request = require('supertest');
const app = require('../src/index');

describe('GET /v1/heartbeat', () => {
  it('should return status 200', async () => {
    const res = await request(app).get('/v1/heartbeat');
    expect(res.status).toBe(200);
  });

  it('should return application/json content type', async () => {
    const res = await request(app).get('/v1/heartbeat');
    expect(res.headers['content-type']).toMatch(/application\/json/);
  });

  it('should return status ok', async () => {
    const res = await request(app).get('/v1/heartbeat');
    expect(res.body.status).toBe('ok');
  });

  it('should return valid uptimeSeconds as non-negative number', async () => {
    const res = await request(app).get('/v1/heartbeat');
    expect(typeof res.body.uptimeSeconds).toBe('number');
    expect(res.body.uptimeSeconds).toBeGreaterThanOrEqual(0);
  });

  it('should return valid ISO-8601 timestamp', async () => {
    const res = await request(app).get('/v1/heartbeat');
    const timestamp = new Date(res.body.timestamp);
    expect(timestamp.toString()).not.toBe('Invalid Date');
    // Verify it's a valid ISO-8601 string
    expect(res.body.timestamp).toMatch(/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}/);
  });

  it('should return all required fields', async () => {
    const res = await request(app).get('/v1/heartbeat');
    expect(res.body).toHaveProperty('status');
    expect(res.body).toHaveProperty('uptimeSeconds');
    expect(res.body).toHaveProperty('timestamp');
  });
});
