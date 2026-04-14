# Backend Layout & Router Conventions

## Express App Bootstrap
- **File**: `src/index.js`
- **Pattern**: Exports Express app instance for testability
- **Middleware**: JSON parser, v1 router mounting, root health check, error handlers
- **Port**: Configurable via `PORT` env var, defaults to 3000

## /v1 Router Wiring
- **File**: `src/routes/v1.js`
- **Mount**: `app.use('/v1', v1Router)` in bootstrap
- **Convention**: All API routes prefixed with `/v1/` for versioning
- **Route Registration**: Handlers imported from route-specific modules and wired in router file

## Route Handler Structure
- **Location**: `src/routes/[handler-name].js`
- **Export Style**: Named exports (e.g., `export const heartbeat = (req, res) => {...}`)
- **Module System**: ESM (ES6 import/export)
- **Handler Characteristics**: Pure functions, minimal side effects, testable in isolation

## Jest Configuration
- **Config File**: `jest.config.js` (ESM format)
- **Test Location**: `__tests__/[feature-name].test.js` colocated with handlers
- **Test Environment**: Node.js
- **Coverage**: Collected from `src/**/*.js`
- **Transform**: None (native ES module support)

## Conventions Summary
✓ Layered structure: bootstrap → router → handler  
✓ ESM module system throughout  
✓ Colocated tests in `__tests__/` directory  
✓ Named exports for handler functions  
✓ Versioned API routes (/v1/)  
✓ Pure, testable handlers  
