const request = require('supertest');
const app = require('../src/app');

describe('GET /v1/ping', () => {
  it('returns 200 with status ok and message pong', async () => {
    const res = await request(app).get('/v1/ping');

    expect(res.status).toBe(200);
    expect(res.headers['content-type']).toMatch(/application\/json/);
    expect(res.body).toEqual({ status: 'ok', message: 'pong' });
  });

  it('returns 404 for an unknown route', async () => {
    const res = await request(app).get('/v1/unknown');

    expect(res.status).toBe(404);
    expect(res.body.error.code).toBe('NOT_FOUND');
  });
});
