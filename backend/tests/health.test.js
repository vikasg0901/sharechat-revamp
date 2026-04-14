'use strict';

const request = require('supertest');
const app = require('../src/app');

describe('GET /health', () => {
  it('returns 200 with status ok', async () => {
    const res = await request(app).get('/health');

    expect(res.status).toBe(200);
    expect(res.headers['content-type']).toMatch(/application\/json/);
    expect(res.body.status).toBe('ok');
  });

  it('returns a valid ISO 8601 timestamp', async () => {
    const before = Date.now();
    const res = await request(app).get('/health');
    const after = Date.now();

    const ts = new Date(res.body.timestamp).getTime();
    expect(ts).toBeGreaterThanOrEqual(before);
    expect(ts).toBeLessThanOrEqual(after);
  });

  it('returns no unexpected fields', async () => {
    const res = await request(app).get('/health');

    expect(Object.keys(res.body).sort()).toEqual(['status', 'timestamp']);
  });
});
