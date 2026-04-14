'use strict';

const express = require('express');
const createHealthRouter = require('./routes/health');

const app = express();

app.use(express.json());

// Health route is registered OUTSIDE any auth group so load-balancer probes
// never receive 401. Do not move this inside middleware groups that attach auth.
app.use('/api', createHealthRouter());

module.exports = app;
