const { Router } = require('express');
const { heartbeat } = require('./heartbeat');

const v1Router = Router();

// Health / heartbeat endpoint
v1Router.get('/heartbeat', heartbeat);

module.exports = v1Router;
