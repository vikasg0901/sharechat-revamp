'use strict';

const request = require('supertest');
const app = require('../app');

describe('GET /api/health', () => {
  it('returns status 200 with ok status and numeric uptime', async () => {
    const res = await request(app).get('/api/health');

    expect(res.status).toBe(200);
    expect(res.body.status).toBe('ok');
    expect(typeof res.body.uptime).toBe('number');
    expect(res.body.uptime).toBeGreaterThanOrEqual(0);
  });
});
