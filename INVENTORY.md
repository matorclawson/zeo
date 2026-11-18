# Component inventory

**Created:** 2026-11-17 · **Last-updated:** 2026-11-17 · **Status:** Initial source-backed cut; deployment inspection pending · **Doctrine:** CLAUDE.md Rev 18 body.

Scope: mandate section 4 and milestone A. UNKNOWN means not inspected or not established by evidence. Source presence does not prove deployed behavior or Linux compatibility. No release acceptance tests have run.

## Dependency map

Observed source edges (not a validated release graph):

- Runtime → Messenger: `go.mod:9` imports the private Messenger module; messaging is a cross-repository dependency.
- Messenger → Runtime: `bin/zmsg:250` requires a pinned `zeo` relay; the Linux service installer also admits a Runtime payload (`deploy/service/install-service.sh:10`). This creates a coordinated-release dependency in both directions.
- Runtime → Pomar: `internal/pomarci/client.go:351,509,659` implements start, signed-result retrieval and run through a control-socket client. This is source integration, not proof that an installed runner is adopted.
- Messenger → Tailscale: `cmd/zmsgd/main.go:151` authenticates callers through node lookup; the Linux unit depends on `tailscaled.service`.
- Pomar manager → signed Swift host → Apple Containerization: `internal/manager/manager.go:96`, `host/Package.swift:17`; host-side adapter work must stay separate from the Linux product.

The complete graph, contract versions, Linux service boundaries, and all other component edges remain UNKNOWN. Do not create replacement schedulers, credential stores or environment managers on this preliminary evidence.

## Hardcoded operational references

Initial findings below refer to tracked source at the revisions in each entry. Private operational values are retained only in the restricted evidence report, not repeated here. This is a targeted first pass, not a complete publication scan.

| Component | Hardcoded reference / coupling | Source evidence | Distribution consequence |
|---|---|---|---|
| Runtime | Keychain secret retrieval through `/usr/bin/security` | `internal/delivery/secret_source.go:27,33` | Requires a supported Linux secret source or an explicit optional Mac host adapter. |
| Runtime | launchctl service cutover and native worker launch | `internal/runtimecutover/transition.go:129,271`; `internal/nativeci/runner.go:381` | These workflows need Linux alternatives; core runtime Linux support remains unverified. |
| Messenger | Fixed Mac volume paths for mailbox, client, logs and receipts; fixed tailnet listen address; Tailscale app path | `cmd/zmsgd/main.go:59–90`; `bin/zmsg:126` | Installer must override all relevant paths and host settings per independent organization. |
| Messenger | launchctl install/upgrade and wake operations | `deploy/install-v5.sh:127`; `bin/zmsg:2311,3484` | Linux service definition alone does not prove wake/upgrade parity. |
| Messenger | Fixed internal service domain and bridge/Architect seat mapping in Linux unit | `deploy/service/zmsgd.service:23` | Must be configurable for external organizations and tested with two instances. |
| Pomar | macOS 26 host package, Apple Containerization dependency and codesign gate | `host/Package.swift:9,17`; `Makefile:18,82–86` | Mac host adapter is implemented in inspected worktree; Linux VPS backend remains UNKNOWN. |

Personal usernames, account IDs, organization IDs and additional domains elsewhere in these trees remain UNKNOWN pending a full tracked-content and history scan. No licence or secret review clearance is implied.

## Runtime

| Field | Finding |
|---|---|
| Repository and revision | `rodriveracom/zeoruntime` (upstream remote), local `zeoruntime/main`, HEAD `71cc3fec87ac6faeb1e2b1cabd4689a02b9c2a74`; tracked tree clean at inspection. |
| Current owner | Expected component seat ZEO-RT per component AGENTS.md; current accountable individual and acceptance of distribution work UNKNOWN. Distribution integration: zeo-maintainer. |
| Purpose | Execution requests and event transitions are persisted by `internal/store/files.go:44–124`; coordinator runs lifecycle endpoints and API servers (`internal/service/coordinator.go:152,311–385,1068`). Full responsibility boundary still under inspection. |
| Implementation | Go 1.27.0 (`go.mod:3`), GNU Make >=3.82 (`Makefile:18–28`); private Messenger module, UUID, TOML, YAML, JSON schema and SQLite modules are declared in go.mod. A Swift app also exists in this tree; not assumed a required Linux runtime dependency. |
| Platform assumptions | Mac-only delivery secrets, launchctl cutover and native build launch are confirmed below. Core Linux and ARM64/AMD64 operation NOT TESTED here. |
| Execution model | `cmd/zeo` CLI built by `Makefile:142–145`; coordinator and Unix-socket API service (`internal/service/api.go:113,401`). Runtime service launchd implementation exists (`internal/service/launchd/launchd.go:275`). Installed service supervision UNKNOWN. |
| State | Protected root mode 0700; immutable execution request JSON mode 0600; per-execution events with allocated sequence and transition validation (`internal/store/files.go:28–124`). Other databases, migrations, ownership and cross-component consistency UNKNOWN. |
| Interfaces | Unix-socket API and execution domain; Messenger module import; Pomar control-socket client with signed result verification (`internal/pomarci/client.go:275–310,351,509`). Complete API authentication and contract catalogue UNKNOWN. |
| Deployment | Installed `zeo --version --json` reports 0.13.0, release SHA `3c94130e040815751ba9b0d261235c905c6d1a90`, unmodified build identity. This differs from inspected source HEAD. No service-state inspection or deployment activation performed. |
| Dependencies | Messenger private module at pinned pseudo-version; Pomar integration source; Go module graph (`go.mod`). Provider, Agenda, Connect and credential integration completeness UNKNOWN. |
| Publication blockers | Private upstream and private Messenger module path confirmed; licence/rights, tracked secrets and history review NOT DONE. No public-build claim. |
| Release blockers | Linux secret/service equivalents; exact compatible Messenger/Runtime release pairing; two-org and seat-isolation behavior; coordinated lifecycle and release tests all unverified. No supported resource measurements. |
| Evidence | R1–R5 below; source presence only. No build or behavioral gate run on this Mac. |

## Architect

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN |
| Current owner | UNKNOWN |
| Purpose | UNKNOWN |
| Implementation | UNKNOWN |
| Platform assumptions | UNKNOWN |
| Execution model | UNKNOWN |
| State | UNKNOWN |
| Interfaces | UNKNOWN |
| Deployment | UNKNOWN |
| Dependencies | UNKNOWN |
| Publication blockers | UNKNOWN |
| Release blockers | UNKNOWN |
| Evidence | UNKNOWN |

## Messenger

| Field | Finding |
|---|---|
| Repository and revision | `rodriveracom/zeomessenger` (upstream remote), local `zeomessenger/main`, HEAD `33958c2ebaca24eb7c2df97dd28920a56ce379b8`; no tracked modifications; untracked `worktrees/` excluded from evidence. |
| Current owner | Expected component seat zeomessenger per AGENTS.md and contributor routing; confirmation of distribution work ownership UNKNOWN. Runtime retains message-log boundary per AGENTS.md. |
| Purpose | Seat mail through a pinned relay CLI; REST/MCP routes are registered in `cmd/zmsgd/main.go:1778–1780`; shell client invokes pinned Runtime relay (`bin/zmsg:250`). Notifications and acknowledgments are separate concepts; live delivery not exercised. |
| Implementation | Go 1.27.0 (`go.mod:3`), Bash CLI, Makefile; modernc SQLite module. Linux installation script additionally uses shell, Perl, hashing, flock and system service tools (`deploy/service/install-service.sh:67–113`). Complete dependency list UNKNOWN. |
| Platform assumptions | Mac paths and Tailscale app default; launchctl install/wake paths. Linux systemd definition and installer exist, with explicit Linux paths. Neither source path establishes cross-platform qualification. |
| Execution model | CLI `bin/zmsg`, remote client, Go API `cmd/zmsgd`, wake/backup/delivery command trees. Systemd service uses separate service user and restart-on-failure (`deploy/service/zmsgd.service:18–26`). Installed processes UNKNOWN. |
| State | CLI root, call log, receipts, attention and state paths are separate flags (`cmd/zmsgd/main.go:59–90`); Linux unit specifies mailbox and shadow SQLite mirror paths. Installer transaction receipts, locks and recovery fences are in `deploy/service/install-service.sh:67–79`. Full backup consistency and migration compatibility unverified. |
| Interfaces | REST `/v1/`, `/mcp` handlers; tailnet node lookup (`cmd/zmsgd/main.go:151–168`) and seat/node binding (`:264`); CLI digest pin. Further v2 and delivery APIs present but not yet traced. |
| Deployment | Linux unit and installer inspected as source only. Installer requires protected independent acceptance for cutover and explicit payload digests; staged installation is not service activation (`deploy/service/install-service.sh:14–59`). Actual host, revision and active unit UNKNOWN. |
| Dependencies | Runtime relay binary pinned by CLI; Tailscale; shell/Perl/Linux installation tools; modernc SQLite. Current private network policy and seats are not public distribution inputs. |
| Publication blockers | Private upstream/module namespace; hardcoded operational paths/domain/seat mapping. History, secrets, rights and licence review NOT DONE. Public installation without private infrastructure unverified. |
| Release blockers | Parameterized independent org instances; tested Linux wake/upgrade/recovery path; dependency cycle release coordination; external authentication/network setup; deployment and all acceptance evidence pending. |
| Evidence | M1–M4 below; no mailbox contents, credentials or live configuration read. No gate run. |

## Agenda

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN |
| Current owner | UNKNOWN |
| Purpose | UNKNOWN |
| Implementation | UNKNOWN |
| Platform assumptions | UNKNOWN |
| Execution model | UNKNOWN |
| State | UNKNOWN |
| Interfaces | UNKNOWN |
| Deployment | UNKNOWN |
| Dependencies | UNKNOWN |
| Publication blockers | UNKNOWN |
| Release blockers | UNKNOWN |
| Evidence | UNKNOWN |

## Controlling

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN |
| Current owner | UNKNOWN |
| Purpose | UNKNOWN |
| Implementation | UNKNOWN |
| Platform assumptions | UNKNOWN |
| Execution model | UNKNOWN |
| State | UNKNOWN |
| Interfaces | UNKNOWN |
| Deployment | UNKNOWN |
| Dependencies | UNKNOWN |
| Publication blockers | UNKNOWN |
| Release blockers | UNKNOWN |
| Evidence | UNKNOWN |

## Connect

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN |
| Current owner | UNKNOWN |
| Purpose | UNKNOWN |
| Implementation | UNKNOWN |
| Platform assumptions | UNKNOWN |
| Execution model | UNKNOWN |
| State | UNKNOWN |
| Interfaces | UNKNOWN |
| Deployment | UNKNOWN |
| Dependencies | UNKNOWN |
| Publication blockers | UNKNOWN |
| Release blockers | UNKNOWN |
| Evidence | UNKNOWN |

## App and backend

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN |
| Current owner | UNKNOWN |
| Purpose | UNKNOWN |
| Implementation | UNKNOWN |
| Platform assumptions | UNKNOWN |
| Execution model | UNKNOWN |
| State | UNKNOWN |
| Interfaces | UNKNOWN |
| Deployment | UNKNOWN |
| Dependencies | UNKNOWN |
| Publication blockers | UNKNOWN |
| Release blockers | UNKNOWN |
| Evidence | UNKNOWN |

## Pomar

| Field | Finding |
|---|---|
| Repository and revision | `zeroemployeeorg/pomar` (upstream remote). Local `pomar/main` HEAD `42a887d3c92c339d04dfb606e780eb6b0aa5fd61` is README/LICENSE only. Separately inspected clean `pomar/worktrees/pomar-recovery-20261117` HEAD `3d79cafdc9197e1a87ded3226a967fee5f66b295`. Do not treat the worktree as released main. |
| Current owner | Current delivery owner UNKNOWN; Pomar component seat confirmation needed, including Go manager versus Swift host ownership. Distribution integration: zeo-maintainer. |
| Purpose | Go manager represents pinned guests and job classes (`internal/manager/manager.go:33–96`); Swift helper options represent per-attempt commands, pins, inputs/outputs and resource caps (`host/Sources/PomarHostCore/Helper.swift:8–107`). Retained agent disk handling is represented separately by `AgentEnvironment.swift:5–55`. Behavior not run. |
| Implementation | Go 1.27.0 (`go.mod:3`); Swift tools 6.2, SwiftPM, Apple Containerization exactly 0.45.0 (`host/Package.swift:1–17`); Makefile full gate includes Go checks, Swift build/test and codesign (`Makefile:18,73–91`). |
| Platform assumptions | Swift host requires macOS 26; macOS virtualization entitlement checked by signing gate. Go manager separates host binary. Linux VPS backend, supported guest distro and architecture qualification UNKNOWN. No nested virtualization or GPU support assumed. |
| Execution model | Go manager with Unix-domain HTTP server (`internal/manager/server.go:40–98`); per-attempt Swift host helper; session service source exists (`internal/sessiond/server.go`). Necessity and deployment of each process remain to be established. |
| State | Manager, attempts, store and tmp paths under a data root (`internal/manager/manager.go:33–38`); retained rootfs and identity/pin/socket fields (`AgentEnvironment.swift:29–55`). Actual data owner, schema migration, host-reboot recovery and consistent backup UNKNOWN. |
| Interfaces | Unix socket permissions set by server (`internal/manager/server.go:85–98`); attempt create/get/stop, reconcile and capacity handlers (`:124–250`). Runtime consumes Pomar via its own signed-result client. Full authentication and version compatibility contract UNKNOWN. |
| Deployment | No installed manager/socket/service inspected. Source worktree is newer than local main; public branch/release availability and installed revision UNKNOWN. No VM provisioned or build run. |
| Dependencies | Apple Containerization and pinned guest kernel/image/source/tool inputs; Go manager imports internal proxies, mirror, capacity, result and signing packages. Full external dependency and Linux backend inventory pending. |
| Publication blockers | Apache-2.0 licence present in local main; per-file/asset rights and history scan NOT DONE. Implementation worktree versus public main divergence must be resolved before claiming available functionality. Private references still need scanning. |
| Release blockers | Tested Linux VPS route; reproducible released Mac guest route; ARM64 qualification; restart, networking, disk, backup and contention evidence; adoption of exact manager/host release. These remain unverified. |
| Evidence | P1–P5 below. CLI documentation is context, not execution proof; no installed socket guessed and no jobs started. |

## Connect and credential integrations

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN |
| Current owner | UNKNOWN |
| Purpose | UNKNOWN |
| Implementation | UNKNOWN |
| Platform assumptions | UNKNOWN |
| Execution model | UNKNOWN |
| State | UNKNOWN |
| Interfaces | UNKNOWN |
| Deployment | UNKNOWN |
| Dependencies | UNKNOWN |
| Publication blockers | UNKNOWN |
| Release blockers | UNKNOWN |
| Evidence | UNKNOWN |

## Recorder, evidence and artifacts

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN |
| Current owner | UNKNOWN |
| Purpose | UNKNOWN |
| Implementation | UNKNOWN |
| Platform assumptions | UNKNOWN |
| Execution model | UNKNOWN |
| State | UNKNOWN |
| Interfaces | UNKNOWN |
| Deployment | UNKNOWN |
| Dependencies | UNKNOWN |
| Publication blockers | UNKNOWN |
| Release blockers | UNKNOWN |
| Evidence | UNKNOWN |

## Shared contracts and libraries

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN |
| Current owner | UNKNOWN |
| Purpose | UNKNOWN |
| Implementation | UNKNOWN |
| Platform assumptions | UNKNOWN |
| Execution model | UNKNOWN |
| State | UNKNOWN |
| Interfaces | UNKNOWN |
| Deployment | UNKNOWN |
| Dependencies | UNKNOWN |
| Publication blockers | UNKNOWN |
| Release blockers | UNKNOWN |
| Evidence | UNKNOWN |

## Other required components

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN |
| Current owner | UNKNOWN |
| Purpose | UNKNOWN |
| Implementation | UNKNOWN |
| Platform assumptions | UNKNOWN |
| Execution model | UNKNOWN |
| State | UNKNOWN |
| Interfaces | UNKNOWN |
| Deployment | UNKNOWN |
| Dependencies | UNKNOWN |
| Publication blockers | UNKNOWN |
| Release blockers | UNKNOWN |
| Evidence | UNKNOWN |

## Evidence catalogue and reproduction

All observations dated 2026-11-17. Revisions above are local source observations, not claims about current GitHub visibility or latest upstream HEAD. Paths below are relative to the named component checkout; use `git show <revision>:<path>` to inspect the exact source, and `git ls-tree -r <revision>` to reproduce file-presence findings. Check current `git status --short` before comparing a working tree. Local component roots were siblings of the flagship under `zeroemployeeorg/`.

| Evidence | Inspected source / command | Proves / limit |
|---|---|---|
| R1 | Runtime `git rev-parse HEAD`, `git status --short`, `git remote get-url upstream`, `go.mod:1–25` | Source identity and declared modules; no remote visibility probe. |
| R2 | Runtime `internal/store/files.go:28–124` | Request/event storage implementation; behavior untested. |
| R3 | Runtime `internal/service/coordinator.go:152,311–385,1068`; `internal/service/api.go:113,401` | Coordinator and Unix API source. |
| R4 | Runtime `internal/delivery/secret_source.go:27–41`; `internal/runtimecutover/transition.go:129,271`; `internal/nativeci/runner.go:381`; `internal/pomarci/client.go:275–310,351,509,659` | Mac couplings and Pomar client source. |
| R5 | Installed `zeo --version --json` | Version 0.13.0 with distinct release SHA above; service operation not established. |
| M1 | Messenger `git rev-parse HEAD`, `git status --short`, `git remote get-url upstream`, `go.mod:1–19` | Source identity and declared dependencies. |
| M2 | Messenger `cmd/zmsgd/main.go:59–99,151–168,264,1778–1780`; `bin/zmsg:126,250,2311,3484` | Defaults, node lookup, routing and relay/launchctl source. |
| M3 | Messenger `deploy/service/zmsgd.service:1–44` | Linux unit specification, fixed mappings and hardening; no installed-state claim. |
| M4 | Messenger `deploy/service/install-service.sh:1–113`; `deploy/install-v5.sh:127` | Controlled Linux cutover and Mac installer source; no install run. |
| P1 | Pomar main `git rev-parse HEAD`, `git status --short`, `git ls-tree -r HEAD`, `README.md`, `LICENSE` | Local main is a stub with Apache-2.0 licence. |
| P2 | Pomar recovery worktree `git rev-parse HEAD`, `git status --short`, `git remote get-url upstream`, `go.mod`, `host/Package.swift:1–33` | Separate clean implementation revision and host dependencies. |
| P3 | Pomar recovery `internal/manager/manager.go:33–100`; `internal/manager/server.go:40–98,124–250` | Paths, config and socket/HTTP source. |
| P4 | Pomar recovery `host/Sources/PomarHostCore/Helper.swift:8–107`; `host/Sources/PomarHostCore/AgentEnvironment.swift:1–55` | Attempt and retained environment input/resource fields; no VM behavior proof. |
| P5 | Pomar recovery `Makefile:18,73–91` | Full gate includes macOS Swift/signing stages; never run here. |

## Open work and blockers

- Confirm current component owners and map these findings into permanent tracked implementation work; confirm the ID mechanism with zeo-ledger after mailbox registration.
- Inspect actual installed services and deployment configurations through the authorized route. `zeoconnect` is not on this session's PATH; no authenticated production gate route has been established in this first cut. production-03/04 are the SOW's gate venues, not yet verified platform support.
- Inspect Architect, Agenda, Controlling, Connect, App/backend, credentials, Recorder/artifacts and shared libraries; expand the complete repository destination map.
- Complete hardcoded-reference and publication scans, including history, private modules, licences and assets. Initial source findings are not publication clearance.
- Establish coordinated release versions, measured resources and all 14 release demonstrations. No product requirement is Verified from this inventory.
