# CollabBoard — Progress Tracker

> A Distributed, Real-Time Collaborative Workspace

---

## Section 0: Git Workflow
- [x] `.gitignore` in place
- [x] Feature branch strategy documented
- [x] Conventional Commits format adopted
- [x] Tags planned per milestone (v0.1.0–v0.4.0)

## Section 1: Real-Time Synchronization (task/board state only)
- [x] WebSocket connection handler per client
- [x] Redis Pub/Sub subscription per `board:{id}`
- [x] Postgres write → Pub/Sub publish pipeline
- [x] Sequence ID per board for gap detection
- [x] Client reconnect with last-known sequence_id replay
- [x] Idempotency key for duplicate writes
- [x] Event batching for high-frequency events (cursor moves)
- [x] Backend crash → client auto-reconnect with exponential backoff + jitter
- [x] Out-of-order detection and resync
- [ ] Measured p50/p99 latency *(requires load testing infra)*
- [ ] Acceptance: kill instance under load → clients catch up with zero missed/duplicated updates *(requires live cluster)*

## Section 2: Authentication & Authorization
- [x] RS256 JWT access tokens (5–15 min TTL)
- [x] Auth service with private key; JWKS endpoint for verification
- [x] Opaque refresh tokens, stored hashed in Postgres, rotated on use
- [x] Refresh token reuse detection → revoke session family
- [x] Redis-backed jti denylist with TTL
- [x] Role-based authorization (Owner/Admin/Member/Viewer) per workspace
- [x] Token-bucket rate limiter in Redis (per-IP + per-account)
- [x] Append-only audit log for all mutating actions
- [x] Clock skew tolerance (~30s leeway on exp/iat)
- [x] Per-session revocation
- [x] Signing key rotation via `kid` header
- [x] Role downgrade → re-check from DB on sensitive actions
- [ ] Acceptance: revoked token rejected; downgraded role loses privileges; reused refresh token kills session chain *(requires live cluster)*

## Section 3: Optimistic Concurrency Control
- [x] `version` integer on every mutable row
- [x] `UPDATE ... SET version = version + 1 WHERE id = $1 AND version = $2`
- [x] 409 Conflict → return current server state
- [ ] Frontend conflict resolution UI ("keep mine" / "take theirs" / manual) *(frontend work)*
- [x] Row-level vs field-level versioning decision documented (row-level chosen)
- [x] Idempotency key prevents double version bump
- [x] Stale client version → 409 with full current state
- [x] Bulk operation transaction boundary (all-or-nothing vs per-item — per-item partial success chosen)
- [x] Acceptance: two concurrent updates → exactly one succeeds, other gets 409, no silent data loss

## Section 4: Search & Caching

### 4a. Search
- [x] Postgres full-text search (tsvector/tsquery)
- [x] Transactionally consistent index updates
- [x] Permission-filtered search results (via memberships join)

### 4b. Caching
- [x] Cache-aside: Redis → miss → Postgres → populate with TTL
- [x] Write-through cache invalidation (not TTL-dependent)
- [x] Cache invalidation events over Redis Pub/Sub
- [x] Single-flight pattern for thundering herd prevention
- [x] Short TTL safety net
- [x] Permission-scoped cache keys
- [ ] Acceptance: measured hit ratio; writes reflected in subsequent reads *(requires load testing infra)*

## Section 5: Data Model
- [x] `workspaces` table
- [x] `memberships` table
- [x] `boards` table
- [x] `tasks` table
- [x] `documents` table
- [x] `audit_log` table
- [x] `refresh_tokens` table
- [x] Migration files (001 + 002 for FTS index)

## Section 6: Observability
- [x] Structured JSON logging with request IDs
- [x] WebSocket connection count metric
- [ ] Fan-out latency metric *(requires instrumentation)*
- [ ] Cache hit ratio metric *(requires instrumentation)*
- [x] 409-conflict rate metric
- [x] Auth failure rate metric
- [ ] Tracing across auth-service ↔ API-service boundary *(requires OpenTelemetry)*

## Section 7: Deployment / Scaling
- [x] Stateless reconnect (any instance, not sticky sessions)
- [x] Graceful drain on rolling deploy (SIGTERM → Shutdown)
- [x] Client reconnect signal on instance shutdown

---

## Definition of Done
- [ ] Live demo survives: killing a pod mid-session, two tabs editing same task, revoked token rejected, permission-scoped search under cache load *(requires live deployment)*
- [x] Git log shows feature-by-feature history (no giant commits)
