const request = require('supertest');
const { createApp } = require('../src/app');

describe('Express app', () => {
  let app;

  beforeAll(() => {
    app = createApp();
  });

  it('serves the static index.html at /', async () => {
    const res = await request(app).get('/');
    expect(res.status).toBe(200);
    expect(res.headers['content-type']).toMatch(/html/);
  });
});
