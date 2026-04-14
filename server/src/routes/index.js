const { Router } = require('express');
const pingRouter = require('./ping');

const router = Router();

router.use('/ping', pingRouter);

module.exports = router;
