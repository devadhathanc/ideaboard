# CollabBoard — Progress Tracker

> A Distributed, Real-Time Collaborative Workspace

---

## Section 0: Git Workflow
- [ ] `.gitignore` in place
- [ ] Feature branch strategy documented
- [ ] Conventional Commits format adopted
- [ ] Tags planned per milestone

## Section 1: Real-Time Synchronization (task/board state only)
- [ ] WebSocket connection handler per client
- [ ] Redis Pub/Sub subscription per `board:{id}`
- [ ] Postgres write → Pub/Sub publish pipeline
- [ ] Sequence ID per board for gap detection
- [ ] Client reconnect with last-known sequence_id replay
- [ ] Idempotency key for duplicate writes
- [ ] Event batching for high-frequency events (cursor moves)
- [ ] Backend crash → client auto-reconnect with exponential backoff + jitter
- [ ] Out-of-order detection and resync
- [ ] Measured p50/p99 latency
- [ ] Acceptance: kill instance under load → clients catch up with zero missed/duplicated updates

## Section 2: Authentication & Authorization
- [ ] RS256 JWT access tokens (5–15 min TTL)
- [ ] Auth service with private key; JWKS endpoint for verification
- [ ] Opaque refresh tokens, stored hashed in Postgres, rotated on use
- [ ] Refresh token reuse detection → revoke session family
- [ ] Redis-backed jti denylist with TTL
- [ ] Role-based authorization (Owner/Admin/Member/Viewer) per workspace
- [ ] Token-bucket rate limiter in Redis (per-IP + per-account)
- [ ] Append-only audit log for all mutating actions
- [ ] Clock skew tolerance (~30s leeway on exp/iat)
- [ ] Per-session revocation
- [ ] Signing key rotation via `kid` header
- [ ] Role downgrade → re-check from DB on sensitive actions
- [ ] Acceptance: revoked token rejected; downgraded role loses privileges; reused refresh token kills session chain

## Section 3: Optimistic Concurrency Control
- [ ] `version` integer on every mutable row
- [ ] `UPDATE ... SET version = version + 1 WHERE id = $1 AND version = $2`
- [ ] 409 Conflict → return current server state
- [ ] Frontend conflict resolution UI ("keep mine" / "take theirs" / manual)
- [ ] Row-level vs field-level versioning decision documented
- [ ] Idempotency key prevents double version bump
- [ ] Stale client version → 409 with full current state
- [ ] Bulk operation transaction boundary (all-or-nothing vs per-item)
- [ ] Acceptance: two concurrent updates → exactly one succeeds, other gets 409, no silent data loss

## Section 4: Search & Caching

### 4a. Search
- [ ] Postgres full-text search (tsvector/tsquery)
- [ ] Transactionally consistent index updates
- [ ] Permission-filtered search results

### 4b. Caching
- [ ] Cache-aside: Redis → miss → Postgres → populate with TTL
- [ ] Write-through cache invalidation (not TTL-dependent)
- [ ] Cache invalidation events over Redis Pub/Sub
- [ ] Single-flight pattern for thundering herd prevention
- [ ] Short TTL safety net
- [ ] Permission-scoped cache keys
- [ ] Acceptance: measured hit ratio; writes reflected in subsequent reads

## Section 5: Data Model
- [ ] `workspaces` table
- [ ] `memberships` table
- [ ] `boards` table
- [ ] `tasks` table
- [ ] `documents` table
- [ ] `audit_log` table
- [ ] `refresh_tokens` table
- [ ] Migration files

## Section 6: Observability
- [ ] Structured JSON logging with request IDs
- [ ] WebSocket connection count metric
- [ ] Fan-out latency metric
- [ ] Cache hit ratio metric
- [ ] 409-conflict rate metric
- [ ] Auth failure rate metric
- [ ] Tracing across auth-service ↔ API-service boundary

## Section 7: Deployment / Scaling
- [ ] Stateless reconnect (any instance, not sticky sessions)
- [ ] Graceful drain on rolling deploy
- [ ] Client reconnect signal on instance shutdown

---

## Definition of Done
- [ ] Live demo survives: killing a pod mid-session, two tabs editing same task, revoked token rejected, permission-scoped search under cache load
- [ ] Git log shows feature-by-feature history (no giant commits)
