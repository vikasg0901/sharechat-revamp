'use strict';

const request = require('supertest');
const app = require('./server');

describe('GET /api/v7/ready', () => {
  it('returns 200 with status ok', async () => {
    const res = await request(app).get('/api/v7/ready');
    expect(res.status).toBe(200);
    expect(res.body).toEqual({ status: 'ok' });
  });

  it('returns application/json content-type', async () => {
    const res = await request(app).get('/api/v7/ready');
    expect(res.headers['content-type']).toMatch(/application\/json/);
  });

  it('returns 404 for an unknown route', async () => {
    const res = await request(app).get('/api/v7/nonexistent');
    expect(res.status).toBe(404);
  });
});
