# Component inventory

**Updated:** 2026-11-18 · **Status:** Source inventory in progress; integrated installation and deployment qualification pending.

This public summary omits private operational records, accounts and host configuration. Detailed inspection receipts are retained by maintainers. Until source and evidence publication is complete, some findings cannot be reproduced from public artifacts alone. UNKNOWN means unestablished; a source implementation, test receipt, installed service and working organization are separate claims.

## Integration edges to resolve

- Runtime and Messenger depend on one another's interfaces and pinned artifacts; release assembly must qualify a compatible pair.
- Runtime consumes Pomar execution capabilities. Pomar supplies environments while Runtime retains placement authority.
- Agenda supplies intent; Scheduler provides timing; Runtime/Connect supply action authority. Missing adapters cannot be replaced by assumed grants.
- The application consumes Connect-backed identity and producer feeds. Its API/frontend origin and service identities require independent-instance configuration.
- Organizational coordination consumes journal, source-hosting and messaging interfaces. Existing host-specific job/signature implementations need supported Linux equivalents.
- The complete graph, public source map, contract versions and lifecycle boundaries remain incomplete. Existing registries, schedulers and credential facilities must be reused where appropriate.

## Portability and publication findings

Host-specific credential access, service supervision, signing, filesystem defaults and fixed operational identifiers have been found in several components. Linux packaging and API schemas do not by themselves prove portable workflows. Full tracked-content/history, dependency, licence and asset review remains required. Optional desktop/host adapters must remain distinct from the Linux product.

## Runtime

| Field | Finding |
|---|---|
| Repository and revision | Inspected source revision `71cc3fec87ac6faeb1e2b1cabd4689a02b9c2a74`. A publicly reproducible component source location remains to be established. |
| Current owner | Runtime maintainers; distribution maintainers own cross-component release integration. |
| Purpose | Agent execution requests, durable execution events, lifecycle coordination and local service APIs. |
| Implementation | Go 1.27.0 and GNU Make; declared dependencies include SQLite, configuration formats and the messaging component. |
| Platform assumptions | OS credential access, service cutover and native worker launch contain host-specific implementations. Supported Linux execution and architecture coverage need qualification. |
| Execution model | CLI and coordinator with Unix-domain service APIs. Installed supervision is not established by the source inventory. |
| State | Execution request and event files with explicit directory/file permissions and transition checks. Complete migration and coordinated recovery scope remain unknown. |
| Interfaces | Execution/domain interfaces, local service API, messaging dependency and a Pomar control client with signed-result checks. |
| Deployment | Owner metadata now identifies a supervised non-Linux daemon and a separate Linux CLI artifact. That CLI is a tool build, not a qualified Runtime daemon or Linux installer. Installed health, migration and recovery remain unqualified. |
| Dependencies | Messenger integration, environment execution facilities, provider adapters and the declared module graph. Complete service contracts remain under inspection. |
| Publication blockers | Inaccessible source/dependencies and unresolved publication review. Complete rights and history checks are pending. |
| Release blockers | Linux service/credential paths, compatible component release pins, real integration, independent-organization isolation and lifecycle acceptance. |
| Evidence | Source: go.mod; internal/store/files.go; internal/service/coordinator.go; internal/service/api.go; internal/pomarci/client.go. Observations establish source behavior candidates, not passing tests. |

## Architect

| Field | Finding |
|---|---|
| Repository and revision | Inspected source revision `b4cd7973d58f3a0f3fc81e6468b8533ba109e2f5`; public source publication remains pending. |
| Current owner | Coordination component maintainers; the runtime coordinator retains its operational authority. |
| Purpose | Organizational review/queue/ruling journal, merge monitoring and notifications. Complete conversation and native-harness boundaries are still being inspected. |
| Implementation | Go 1.27 with SQLite. The full declared gate includes formatting, vet, race tests and mutation checks. |
| Platform assumptions | Current job installation uses launchd; artifact verification uses the host signing facility; recall has host-specific filesystem defaults. A Linux lifecycle path needs explicit qualification. |
| Execution model | CLI with scheduled or persistent host-supervised jobs. Source layout does not establish a required additional Linux daemon. |
| State | Append-only SQLite events with unique source keys, WAL and FULL synchronous mode; separate derived recall index. Actual ownership, migrations and consistent recovery remain unverified. |
| Interfaces | Versioned CLI JSON, PR-reader and message-notifier interfaces, journal events and wrapper integrations. Authentication and independent-organization contracts need further tracing. |
| Deployment | Host-specific job and verified-release installation source exists. No current installed artifact or activated-service receipt is held by this inventory. |
| Dependencies | SQLite, source-hosting read interface, Messenger and current host service/signature adapters. Additional harness dependencies remain unknown. |
| Publication blockers | No licence filename found in the inspected tracked tree; source/history, asset and private-dependency review remain open. |
| Release blockers | Linux supervision and signature verification, organization-scoped continuity, installed-version evidence, migration and recovery tests. |
| Evidence | Source: internal/journal/journal.go; internal/monitor/monitor.go; internal/agent/agent.go; internal/release/release.go; cmd/zeoarchitect/recall.go. No private continuity state was inspected. |

## Messenger

| Field | Finding |
|---|---|
| Repository and revision | Inspected source revision `33958c2ebaca24eb7c2df97dd28920a56ce379b8`; publicly accessible source/artifact publication is unresolved. |
| Current owner | Messaging maintainers; underlying runtime log and identity boundaries remain with their respective owners. |
| Purpose | Seat messaging through a pinned relay, REST/MCP APIs and separate notification, acknowledgement and delivery concepts. |
| Implementation | Go 1.27.0, shell clients, Make, SQLite and installer support utilities. The complete runtime package set still needs qualification. |
| Platform assumptions | Both host-specific defaults and Linux service packaging exist. Linux packaging alone does not qualify notification, upgrade or recovery parity. |
| Execution model | Local/remote CLI, API service, wake, backup and delivery commands. The Linux unit declares a dedicated service identity and restart behavior. |
| State | Mail root, logs, receipts, attention/cache state, mirror database and deployment transaction fences. Coordinated backup and schema compatibility remain unverified. |
| Interfaces | REST/MCP routes, node/seat binding and relay digest pinning. Additional delivery and change-feed contracts still need end-to-end tracing. |
| Deployment | A controlled Linux installer requires exact payload bindings and an independently accepted activation operation. Staged files do not establish service activation. |
| Dependencies | Pinned Runtime relay, supported network identity interface, shell utilities and SQLite. Distribution identity setup must work without the maintainers network. |
| Publication blockers | Private module/source coupling, embedded operational assumptions and pending licence/history/asset review. |
| Release blockers | Independent instance parameterization, tested Linux wake/upgrade/recovery, public identity setup and coordinated Runtime compatibility. |
| Evidence | Source: cmd/zmsgd/main.go; bin/zmsg; deploy/service/zmsgd.service; deploy/service/install-service.sh. No mail or live configuration was read. |

## Agenda

| Field | Finding |
|---|---|
| Repository and revision | Inspected candidate `c2aabfe85c3027213a82de1811e41630041053f4`, tree `65f44ab71ef3ac8a663fa29118493fb243e6b7a3`. Candidate qualification is pending. |
| Current owner | Agenda maintainers; Scheduler owns host timing, and Runtime/Connect retain action/identity authority. |
| Purpose | Durable task and temporal intent, personal workspace APIs and reminder registration. Proposed capacity planning is not treated as adopted functionality. |
| Implementation | Go 1.27.0 with standard-library dependencies; daemon, client and commitments entry points. Current candidate full-gate evidence is missing. |
| Platform assumptions | Unix sockets, kernel ownership checks and file locking. No service/timer packaging was found in the inspected tree; Linux/architecture behavior remains unqualified. |
| Execution model | Long-running service with explicit state/credential-reference/runtime-socket arguments and a registration/status pump. Registration is distinct from due-time dispatch. |
| State | Owned private state root, exclusive writer fence and synchronized state replacement; encrypted component-backup implementation. Deployed custody, migrations and restore proof remain missing. |
| Interfaces | Task/productivity/personal routes and injected authority/reminder adapters. Missing adapters remain unavailable; a private socket alone does not grant owner authority. |
| Deployment | Component and runtime owners do not hold an adopted installation tuple. One daemon and two one-shot tools are source-defined; the direct journal CLI needs an exclusive writer window even for export. No source unit or archive establishes adoption. |
| Dependencies | Runtime reminder/verification interfaces, Connect, Messenger, Scheduler and the release build toolchain. Proposed shared-feed contracts are not assumed implemented. |
| Publication blockers | No licence/copying filename found; source/publication scan, public artifacts and private-access-free build remain unresolved. |
| Release blockers | Current candidate gate, real authority/wire integration, maintained Linux lifecycle packaging, compatible contracts and recovery acceptance. |
| Evidence | Owner packet digest, candidate tree and all 15 supplied source-file SHA256/blob pins match inspected objects. Source: cmd/zeoagenda/main.go; internal/app/server.go; internal/agenda/store.go; internal/agenda/backup.go. |

## App and backend

| Field | Finding |
|---|---|
| Repository and revision | Latest owner-pinned candidate `e24013c61b1eb2fc3a5f6a8a62d13e29d1faf8bc`, tree `a727ecebcd556e3c827c2ed00d936391074137d0`. All 15 supplied file size/SHA256/Git-blob pins independently match source objects and the previously inspected selected files at `19fc44892ec06cfb1aacd7928a19a19efbf20737`. Main integration and public release are not inferred. |
| Current owner | Application maintainers own source/product behavior; runtime/operations own installed custody and ingress, Connect owns delegated identity, and upstream components own feeds. |
| Purpose | Operator web workspace, sessions, organizational feeds and durable workflow/outbox state. Proposed project-planning integrations remain unadopted. |
| Implementation | TypeScript, Node, Next, Fastify and React. Retained Linux web-gate receipt records Node 24.22.1/npm 11.19.1. Lock digest `930f77aa8653d42d2f5601bf595a6888a8ec1bf0a74a6bf431670297b168077b` matches inspected bytes. |
| Platform assumptions | POSIX custody checks and built-in Node SQLite. Linux web evidence is narrower than the native/root tree; ARM64, native separation and external-organization behavior remain unqualified. |
| Execution model | Loopback API and web frontend, plus a bounded synchronous feed-publisher CLI. Fixed local origins/ports need multi-instance configuration and tests. |
| State | Shared session/workflow SQLite database with WAL/FULL synchronization and file-custody checks. Consistent backup/restore and a maintained migration contract remain open. |
| Interfaces | Versioned configuration, Connect-backed sessions, feed/HTTP contracts and compare-and-swap publisher inputs. Schema presence is not an enrollment grant or live producer route. |
| Deployment | Component and runtime owners hold no complete current installed-service/publisher tuple; this does not prove global absence. No adopted browser-to-Pomar terminal gateway or qualified install bundle is held. Component web tests do not qualify real terminal attachment. |
| Dependencies | Exact application packages and lockfile; Node SQLite, Connect, producer feeds and owner-managed supervision/ingress. Browser packages are gate requirements, not automatically runtime payload. |
| Publication blockers | MIT licence file observed; package remains private-marked. That observation does not clear history, third-party rights, assets, separate desktop code or publication. |
| Release blockers | External install/onboarding, independent identities, portable ports/service UIDs, producer bindings, migration/restore and ARM64 qualification. |
| Evidence | All 15 latest owner source pins match inspected Git objects. The owner reports a Linux web gate at `e24013c6` with log digest `11191f16635449f0c86afa12172df03de7395a421a9a9f623eba16596737260c`: 158 unit tests, service/SSR and 43 browser tests. The previously retained digest `bf304602f0eec575f14c36be0312bb2bf67bbc6ef6647e38b80cf9ec8a9942c9` belongs to the earlier candidate. No App rerun, installed receipt or complete-product acceptance is claimed here. |

## Pomar

| Field | Finding |
|---|---|
| Repository and revision | Inspected implemented source at `3d79cafdc9197e1a87ded3226a967fee5f66b295`; an older bootstrap checkout had no implementation. Release/adoption identity remains distinct from source presence. |
| Current owner | Pomar maintainers own environment backends; distribution maintainers own the product installed into those environments. |
| Purpose | Pinned guest/job environments, controlled inputs/outputs, resource classes and retained seat-environment support. Actual supported backend capabilities require their own conformance evidence. |
| Implementation | Go 1.27.0, Swift 6.2 and pinned Apple Containerization 0.45.0; the declared full gate includes Go and host-adapter checks. |
| Platform assumptions | Current Apple host adapter targets macOS 26. Linux guest support is distinct from a direct Linux-host backend; supported guest baselines and architectures remain evidence-gated. |
| Execution model | Environment manager over a protected local socket, per-attempt host helpers and optional persistent-session service. Deployment and process boundaries are owner-managed. |
| State | Manager/attempt/object state, pinned environment inputs and retained seat disks. Actual custody, reboot recovery, migration and consistent restore remain to be qualified. |
| Interfaces | Environment attempt, stop, reconcile and capacity APIs; protected socket access and signed-result consumption. Common profile/backend contracts remain under coordination. |
| Deployment | No new guest or manager was started by this inventory. A published release and qualified installed manager must precede adoption claims. |
| Dependencies | Host adapter, pinned kernel/images/tool inputs, source/proxy/cache and result-signing facilities. Backend-specific capability matrices are still needed. |
| Publication blockers | Apache-2.0 licence observed in the bootstrap tree; full source/asset/history review and public artifact qualification remain pending. |
| Release blockers | Qualified published release, common profile, supported host/guest routes, real isolation, resource measurements, recovery and architecture-specific conformance. |
| Evidence | Source: go.mod; host/Package.swift; internal/manager/manager.go; internal/manager/server.go; host helper and retained-environment implementations. No VM behavior or passing current gate is inferred. |

## Scheduler

| Field | Finding |
|---|---|
| Repository and revision | Full maintained source boundary remains unknown. Runtime candidate source was inspected at `71cc3fec87ac6faeb1e2b1cabd4689a02b9c2a74`; candidate code is not assumed to be the entire Scheduler. |
| Current owner | Scheduler maintainer confirmed; precise path, contract and deployment mapping requested. |
| Purpose | Host timing is assigned to this component. Inspected Runtime code supplies bounded Agenda schedule ensure/cancel and intent projections; adopted dispatch behavior still needs proof. |
| Implementation | Candidate implementation is Go within Runtime. Complete binaries, scripts and toolchain requirements remain unknown. |
| Platform assumptions | Candidate API uses Unix peer identity. Host timer backend and architecture support remain unqualified. |
| Execution model | Runtime controller seam observed; an independent Scheduler service is neither required nor inferred from that observation. |
| State | Candidate methods commit to the existing runtime store. Complete timing-state custody, migrations and recovery scope remain unknown. |
| Interfaces | Bounded schedule frames, request IDs and owner-authority fences in candidate code. Actual installed endpoint and accepted contract versions remain unresolved. |
| Deployment | No live Scheduler/timer installation receipt held. |
| Dependencies | Agenda intent and Runtime/Connect authority per component ownership; other dependencies await the owner map. |
| Publication blockers | Source boundary, rights and public dependency review remain incomplete. |
| Release blockers | Owner-confirmed maintained paths, real dispatch/adoption, integration, reboot/recovery and isolation tests. |
| Evidence | Source: internal/service/agenda_schedule.go and internal/autonomy/agenda_schedule.go. Ownership confirmation is separate from implementation and deployment evidence. |

## Controlling

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN — inspection pending. |
| Current owner | UNKNOWN — inspection pending. |
| Purpose | UNKNOWN — inspection pending. |
| Implementation | UNKNOWN — inspection pending. |
| Platform assumptions | UNKNOWN — inspection pending. |
| Execution model | UNKNOWN — inspection pending. |
| State | UNKNOWN — inspection pending. |
| Interfaces | UNKNOWN — inspection pending. |
| Deployment | UNKNOWN — inspection pending. |
| Dependencies | UNKNOWN — inspection pending. |
| Publication blockers | UNKNOWN — inspection pending. |
| Release blockers | UNKNOWN — inspection pending. |
| Evidence | UNKNOWN — inspection pending. |

## Connect

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN — inspection pending. |
| Current owner | UNKNOWN — inspection pending. |
| Purpose | UNKNOWN — inspection pending. |
| Implementation | UNKNOWN — inspection pending. |
| Platform assumptions | UNKNOWN — inspection pending. |
| Execution model | UNKNOWN — inspection pending. |
| State | UNKNOWN — inspection pending. |
| Interfaces | UNKNOWN — inspection pending. |
| Deployment | UNKNOWN — inspection pending. |
| Dependencies | UNKNOWN — inspection pending. |
| Publication blockers | UNKNOWN — inspection pending. |
| Release blockers | UNKNOWN — inspection pending. |
| Evidence | UNKNOWN — inspection pending. |

## Connect and credential integrations

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN — inspection pending. |
| Current owner | UNKNOWN — inspection pending. |
| Purpose | UNKNOWN — inspection pending. |
| Implementation | UNKNOWN — inspection pending. |
| Platform assumptions | UNKNOWN — inspection pending. |
| Execution model | UNKNOWN — inspection pending. |
| State | UNKNOWN — inspection pending. |
| Interfaces | UNKNOWN — inspection pending. |
| Deployment | UNKNOWN — inspection pending. |
| Dependencies | UNKNOWN — inspection pending. |
| Publication blockers | UNKNOWN — inspection pending. |
| Release blockers | UNKNOWN — inspection pending. |
| Evidence | UNKNOWN — inspection pending. |

## Recorder, evidence and artifacts

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN — inspection pending. |
| Current owner | UNKNOWN — inspection pending. |
| Purpose | UNKNOWN — inspection pending. |
| Implementation | UNKNOWN — inspection pending. |
| Platform assumptions | UNKNOWN — inspection pending. |
| Execution model | UNKNOWN — inspection pending. |
| State | UNKNOWN — inspection pending. |
| Interfaces | UNKNOWN — inspection pending. |
| Deployment | UNKNOWN — inspection pending. |
| Dependencies | UNKNOWN — inspection pending. |
| Publication blockers | UNKNOWN — inspection pending. |
| Release blockers | UNKNOWN — inspection pending. |
| Evidence | UNKNOWN — inspection pending. |

## Shared contracts and libraries

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN — inspection pending. |
| Current owner | UNKNOWN — inspection pending. |
| Purpose | UNKNOWN — inspection pending. |
| Implementation | UNKNOWN — inspection pending. |
| Platform assumptions | UNKNOWN — inspection pending. |
| Execution model | UNKNOWN — inspection pending. |
| State | UNKNOWN — inspection pending. |
| Interfaces | UNKNOWN — inspection pending. |
| Deployment | UNKNOWN — inspection pending. |
| Dependencies | UNKNOWN — inspection pending. |
| Publication blockers | UNKNOWN — inspection pending. |
| Release blockers | UNKNOWN — inspection pending. |
| Evidence | UNKNOWN — inspection pending. |

## Other required components

| Field | Finding |
|---|---|
| Repository and revision | UNKNOWN — inspection pending. |
| Current owner | UNKNOWN — inspection pending. |
| Purpose | UNKNOWN — inspection pending. |
| Implementation | UNKNOWN — inspection pending. |
| Platform assumptions | UNKNOWN — inspection pending. |
| Execution model | UNKNOWN — inspection pending. |
| State | UNKNOWN — inspection pending. |
| Interfaces | UNKNOWN — inspection pending. |
| Deployment | UNKNOWN — inspection pending. |
| Dependencies | UNKNOWN — inspection pending. |
| Publication blockers | UNKNOWN — inspection pending. |
| Release blockers | UNKNOWN — inspection pending. |
| Evidence | UNKNOWN — inspection pending. |

## Outstanding release work

Complete remaining source and installed-state inspection; publish the component/revision map and compatible artifacts; establish qualified Linux platforms and measured resources; implement and verify integrated installation, onboarding, isolation and recovery. Existing environment capabilities should be released and adopted through their owner contracts, rather than reimplemented here.

No release acceptance demonstration is marked complete by this inventory. A small component test suite, source merge or running process does not establish a usable distribution. Current and proposed Linux environment-profile work is tracked in REQUIREMENTS.md section 20.
