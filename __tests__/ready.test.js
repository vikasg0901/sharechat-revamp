const request = require('supertest');
const { createApp } = require('../src/app');

describe('GET /api/v7/ready', () => {
  let app;

  beforeAll(() => {
    app = createApp();
  });

  it('returns 200 with ready:true', async () => {
    const res = await request(app).get('/api/v7/ready');
    expect(res.status).toBe(200);
    expect(res.body).toEqual({ ready: true });
  });

  it('responds with application/json content-type', async () => {
    const res = await request(app).get('/api/v7/ready');
    expect(res.headers['content-type']).toMatch(/application\/json/);
  });

  it('returns 404 for unknown routes', async () => {
    const res = await request(app).get('/api/v7/unknown');
    expect(res.status).toBe(404);
  });
});
