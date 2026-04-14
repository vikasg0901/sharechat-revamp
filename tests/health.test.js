'use strict';

const request = require('supertest');
const express = require('express');
const createHealthRouter = require('../src/routes/health');

function buildApp(startTime) {
  const app = express();
  app.use('/api', createHealthRouter(startTime));
  return app;
}

describe('GET /api/health', () => {
  test('returns 200 with application/json Content-Type', async () => {
    const app = buildApp();

    const res = await request(app).get('/api/health');

    expect(res.status).toBe(200);
    expect(res.headers['content-type']).toMatch(/application\/json/);
  });

  test('body has status "ok" and uptime_seconds >= 5 when started 5 s ago', async () => {
    const startTime = Date.now() - 5000;
    const app = buildApp(startTime);

    const res = await request(app).get('/api/health');

    expect(res.body.status).toBe('ok');
    expect(typeof res.body.uptime_seconds).toBe('number');
    expect(Number.isInteger(res.body.uptime_seconds)).toBe(true);
    expect(res.body.uptime_seconds).toBeGreaterThanOrEqual(5);
  });

  test('uptime_seconds grows between successive calls', async () => {
    const startTime = Date.now();
    const app = buildApp(startTime);

    const res1 = await request(app).get('/api/health');
    await new Promise(resolve => setTimeout(resolve, 1100));
    const res2 = await request(app).get('/api/health');

    expect(res2.body.uptime_seconds).toBeGreaterThan(res1.body.uptime_seconds);
  });
});
