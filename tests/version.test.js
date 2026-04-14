'use strict';

const request = require('supertest');
const app = require('../server');
const pkg = require('../package.json');

describe('GET /v1/version', () => {
  test('returns 200 with name and version from package.json', async () => {
    const res = await request(app).get('/v1/version');

    expect(res.status).toBe(200);
    expect(res.headers['content-type']).toMatch(/application\/json/);
    expect(res.body).toEqual({
      name: pkg.name,
      version: pkg.version,
    });
  });

  test('includes X-Request-Id header', async () => {
    const res = await request(app).get('/v1/version');

    expect(res.headers['x-request-id']).toBeDefined();
    expect(typeof res.headers['x-request-id']).toBe('string');
    expect(res.headers['x-request-id'].length).toBeGreaterThan(0);
  });
});

describe('GET /healthz', () => {
  test('returns 200 with status ok', async () => {
    const res = await request(app).get('/healthz');

    expect(res.status).toBe(200);
    expect(res.body).toEqual({ status: 'ok' });
  });
});

describe('unknown route', () => {
  test('returns 404 with error envelope', async () => {
    const res = await request(app).get('/does-not-exist');

    expect(res.status).toBe(404);
    expect(res.body.error).toBeDefined();
    expect(res.body.error.code).toBe('NOT_FOUND');
    expect(typeof res.body.error.message).toBe('string');
  });
});
