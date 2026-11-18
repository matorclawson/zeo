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

Additional edges verified or explicitly attributed during relay reconciliation:

- Agenda → Runtime: `cmd/zeoagenda/main.go:31–32` installs RuntimeReminders for an explicitly supplied socket. Without adapters several actions remain unavailable (`internal/app/server.go:133–150`). Live wire/identity binding remains unproven.
- Agenda → Scheduler/Connect/Messenger: owner packet assigns host timers, identity/actions and message delivery to those components. These are owner-declared integration dependencies, pending independent endpoint/contract inspection; Scheduler must be added to the remaining component inventory.
- App → Connect and producer feeds: `server/main.ts:33–44` wires Connect-backed sessions and published feeds. Current production publisher/identity routes remain unbound per packet; configuration schema alone is not an access grant.
- App web backend → Next frontend: loopback API origin and rewrites (`src/shared/api-origin.ts:1–5`; `next.config.ts:18–22`); multi-instance port/origin support is unresolved.

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
| App | Fixed API origin/port, frontend port and numeric service-UID exclusions | `app/TodosWeb/src/shared/api-origin.ts:1–5`; `package.json` start; `server/config.ts:59–63` | Same-host multi-instance port and identity allocation require design and tests; no configuration-only portability claim. |
| App | Root gate hardcodes host/user admission; native plist packaging remains | `Makefile:496–498`; exact tree plist enumeration | Gate venue must remain owner-admitted; source gate success is not a portable installer. |
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
| Repository and revision | `rodriveracom/zeoagenda` candidate `c2aabfe85c3027213a82de1811e41630041053f4`, tree `65f44ab71ef3ac8a663fa29118493fb243e6b7a3`; pinned WIP, not local main. Owner packet at records commit `ca8ebe5f00e70c4bc0730cb63c1b229fcb842604`; packet SHA256 and all 15 source file SHA256/blob IDs independently match local objects (A1). |
| Current owner | zeoagenda-maintainer identifies itself as component owner in the pinned packet. Packet assigns host timers to Scheduler and actions/identity to Runtime/Connect; their current production bindings remain UNKNOWN. Distribution integration: zeo-maintainer. |
| Purpose | Durable temporal/task intent and personal workspace APIs. Source wires task/productivity/personal handlers (`internal/app/server.go:133–170`) and a reminder registration/status pump (`:73–111`). Capacity planning is described by owner as proposed, not implemented or adopted. |
| Implementation | Go 1.27.0; `go.mod:1–3` declares no external modules. Three command entrypoints: `cmd/zeoagenda`, `cmd/agenda`, `cmd/agenda-commitments`, with build commands recorded in packet. `make verify` is the requested gate; no candidate gate executed or passing receipt supplied in this packet. |
| Platform assumptions | Unix-domain sockets and kernel ownership/locking assumptions confirmed (`internal/app/server.go:46,65`; `internal/agenda/store.go:94–123`). No service/timer/plist files in pinned tree. macOS-only runtime dependency not established from inspected paths; Linux behavior and ARM64/AMD64 qualification remain unverified. |
| Execution model | Long-running Go daemon via `ServeWithAdapters`; signal handling and explicit state/credential/runtime-socket arguments (`cmd/zeoagenda/main.go:15–37`). Reminder registration pump is present; do not infer due-time host scheduling from it. Owner assigns host timers to Scheduler. |
| State | Absolute state root, mode 0700, effective-UID ownership, exclusive `writer.lock` mode 0600 and Flock (`internal/agenda/store.go:94–123`); synchronized atomic state replacement (`:254–260`). AES-GCM encrypted component backup implementation (`internal/agenda/backup.go:69–77,185–258`). Actual deployed state ownership, supported migrations and clean-host restore proof UNKNOWN. |
| Interfaces | Unix `api.sock`, task/productivity/personal dispatch (`internal/app/server.go:46,65,152–169`); injected adapters and explicit Runtime socket. Missing Runtime/Budget/Acceptor adapters become Unavailable; default authorization refuses owner data (`:133–150`). Credential registry and CLI token-file interfaces are specified in packet; production verifier/wire proof absent. |
| Deployment | Owner reports no held supported installation/service/adoption receipt. Pinned tree has no service/timer/plist packaging. No daemon started, protected identity bound or live deployment inspected by maintainer. |
| Dependencies | Standard library runtime; compatible Go/Make and race/C toolchain for gate per packet. Runtime reminder gateway, Connect/Runtime verifiers, Messenger and Scheduler are integration dependencies. Versioned zeocontracts feed contracts are reported proposed; not treated as available. |
| Publication blockers | Private module namespace in go.mod. No LICENSE/COPYING files found by exact-tree filename enumeration. Rights/licence choice, complete secret/history review, public artifacts and private-access-free build remain unresolved. |
| Release blockers | Exact candidate production gate and both readers; real Runtime/Connect authorization and reminder wire tests; Linux service/lifecycle packaging; migration and backup/restore consumer proof; versioned contracts and publication clearance. No independent-organization acceptance evidence. |
| Evidence | A1–A4 below. Packet digest, source tree and 15 file pins verified; selected implementation read at exact candidate. Source metadata findings only; no build, test, service or release demonstration run. |

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
| Repository and revision | Owner reports `rodriveracom/zeoapp` main `a3142585e42a73ca0efc2eba397f717301915048`, tree `b73d7a8b9a1fbe75d8c89b897576fd1b832bc140`. Main commit object absent locally. Inspected gate commit `19fc44892ec06cfb1aacd7928a19a19efbf20737` has exactly that tree; source-byte equivalence confirmed, main ancestry/publication not independently confirmed. Packet at records commit `7076b73a9b9c6698d62755bc1a88775ce5475edd` (W1). |
| Current owner | zeoapp retains product/source authority per pinned owner packet; Ops/Runtime retain host installation/supervision and Connect retains identity/delegation. No installation or custody authority transferred by this inventory. |
| Purpose | Operator web workspace with sessions, feeds and workflow/outbox state. `app/TodosWeb/server/main.ts:31–45` constructs Sessions, Feeds and WorkflowStore into the API. Owner reports new planning/project export and external Projects return unadopted; they are not current distribution features. |
| Implementation | TypeScript/Node, Next, Fastify and React. Inspected scripts regenerate validators, build Next and write build stamp (`app/TodosWeb/package.json`). Historical gate receipt records Node 24.22.1/npm 11.19.1. Lock SHA256 `930f77aa8653d42d2f5601bf595a6888a8ec1bf0a74a6bf431670297b168077b` independently matches. Root tree also retains native code; Linux web scope is narrower than root verify. |
| Platform assumptions | POSIX UID/mode checks and built-in `node:sqlite`. Historical production-03 TypeScript gate evidence exists, with limits below; no Linux ARM64 acceptance. Root native code and plist packaging remain Mac-specific per packet and exact-tree enumeration. Separate desktop source not independently inspected. |
| Execution model | Two long-running loopback processes: Fastify API (`server/main.ts:54`) and Next (`package.json` start script, port 4173). API origin is fixed to `http://127.0.0.1:4380` (`src/shared/api-origin.ts:1–5`). Publisher CLI is bounded synchronous input processing (`server/publish-cli.ts:1–17`), not a scheduler. |
| State | `private_store/web-sessions.sqlite3`; custody checks for database/WAL/SHM and shared Sessions/WorkflowStore database (`server/main.ts:16–44`). Sessions enables WAL, FULL sync and busy timeout (`server/sessions.ts:45`). Consistent DB/WAL backup, clean-host restore and coordinated migration proof remain pending; constructor DDL is present. |
| Interfaces | Config schema `zeoapp.typescript-service/v1`; dedicated service UID and ordinary-seat exclusions (`server/config.ts:14–22,59–65`); feed-backed API and loopback rewrites. CLI publisher accepts named profiles, expected digest and stdin JSON limited to 2 MiB (`server/publish-cli.ts:1–17`). Packet identifies Connect browser-delegation-v1 and versioned HTTP/feed contracts; complete live producer and enrollment route not verified. |
| Deployment | Packet reports no App-held original installed TypeScript/live-service receipt and a frozen earlier release chain. No live service or installed source revision independently observed. Source gate receipt explicitly records deployment=false. Supervisor and HTTPS ingress remain external responsibilities. |
| Dependencies | Exact package/lock pins include Next 16.4.6, Fastify 5.12.8, React 19.3.1, TypeScript 7.0.2 and Playwright 1.64.1; Node built-in SQLite. Connect and real upstream feed publishers are external dependencies. Browser/OS prerequisites belong to test qualification; complete runtime artifact dependencies remain to be established. |
| Publication blockers | MIT LICENSE observed at inspected tree; package.json private=true. Remote privacy is owner-reported, not freshly queried. Whole-tree/history secrets, third-party assets/licences and private dependencies still need review. Root App licence observation does not establish separate desktop rights or publication clearance. |
| Release blockers | Clean external install and independent-org onboarding; port/origin/service-UID portability; installed publisher and producer bindings; consistent backup/restore and migration; ARM64 tests; retained native separation; public/private dependency clearance. No broad release acceptance is supplied by the narrow web gate. |
| Evidence | W1–W5 below. Original log SHA256/size/line count match packet, receipt binds gate commit and records exit 0, 158 unit tests, service/SSR PASS, 43 browser tests, 0 skipped. Source-head binding is runner-attested, not a raw-log banner. No rerun here; root/native gate NOT_RUN under recorded scope; no live Connect, attended device, deployment or external-org proof. |

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
| A1 | Records `org-zeroemployeeorg@ca8ebe5f00e70c4bc0730cb63c1b229fcb842604:projects/zeoagenda/evidence/zeoagenda/distribution-inventory-20261117/AGENDA-COMPONENT-INVENTORY-v1.json` | Packet SHA256 `df893e411ee296fad4fefabb1351ca5489814bf13a9ef4811948a6b92eca80e9` matches relay; candidate tree and all 15 SHA256/blob pins match local Git objects. |
| A2 | Agenda candidate `go.mod:1–3`; `cmd/zeoagenda/main.go:15–37`; `internal/app/server.go:46–111,133–170` | Toolchain, daemon, socket, pump, routing and missing-adapter defaults; source only. |
| A3 | Agenda candidate `internal/agenda/store.go:94–123,254–260`; `internal/agenda/backup.go:69–77,185–258` | Private state, fencing and encrypted backup source; no recovery acceptance. |
| A4 | Agenda candidate `git ls-tree -r --name-only` | No LICENSE/COPYING or service/timer/plist filename present. Packet's no-licence/no-service finding confirmed narrowly by names, not a rights audit. |
| W1 | Records `matorclawson/org-zeroemployeeorg@7076b73a9b9c6698d62755bc1a88775ce5475edd:projects/zeoapp/evidence/distribution-inventory-20261117/APP-COMPONENT-INVENTORY.md` | Committed packet read; SHA256 `8dbf8aae5a5f2e76c435f3a3910548eaf10f726b39aebda5b7665e7f892e3b05`. App gate commit tree matches packet main tree; main commit itself unavailable locally. |
| W2 | App gate commit `app/TodosWeb/package.json`, `package-lock.json`; `Makefile:496–500`; `LICENSE` | Scripts, exact dependencies, matching lock digest, narrow gate admission and MIT file. No execution here. |
| W3 | App gate commit `app/TodosWeb/server/main.ts:16–54`; `server/sessions.ts:45`; `server/config.ts:14–22,59–65`; `server/publish-cli.ts:1–17` | State, service UID checks, API and publisher inputs. Paths after the first are relative to app/TodosWeb. |
| W4 | App gate commit `app/TodosWeb/src/shared/api-origin.ts:1–5`; `app/TodosWeb/next.config.ts:1–22` | Fixed origin and loopback rewrites; actual external access not tested. |
| W5 | Owner evidence `pr169-gate-package-20261117/PASS-ORIGINAL.log`, `PASS-RECEIPT.json` | Log SHA256 `bf304602f0eec575f14c36be0312bb2bf67bbc6ef6647e38b80cf9ec8a9942c9`, 24,169 bytes, 315 lines; receipt SHA256 `e753cd6a7b586a62e165dee0638c6830c6fb656b9927ff116556232f058a281e`. Historical narrow gate integrity, not a new gate or deployment. |

## Open work and blockers

- Confirm current component owners and map these findings into permanent tracked implementation work; confirm the ID mechanism with zeo-ledger after mailbox registration.
- Inspect actual installed services and deployment configurations through the authorized route. `zeoconnect` is not on this session's PATH; no authenticated production gate route has been established in this first cut. production-03/04 are the SOW's gate venues, not yet verified platform support.
- Agenda and App packets are reconciled above. Inspect their actual deployments and integration contracts; inspect Architect, Controlling, Connect, credentials, Recorder/artifacts, shared libraries and newly identified Scheduler; expand the complete repository destination map. Separate desktop source at owner-reported `b2a3a2a0684c51b5f0b051d6da41792420497b5c` is unavailable locally and remains an owner-only report, not verified Linux capability.
- Complete hardcoded-reference and publication scans, including history, private modules, licences and assets. Initial source findings are not publication clearance.
- Establish coordinated release versions, measured resources and all 14 release demonstrations. No product requirement is Verified from this inventory.
