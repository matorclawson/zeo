# ZEO

**Updated:** 2026-11-18 · **Status:** Early distribution implementation; release qualification pending.

ZEO is being assembled as an open-source platform for persistent agent organizations, including ML research teams and small research laboratories. It aims to let users create independent organizations, connect models, set budgets and permissions, and inspect the results of coordinated work.

Available today: the [product requirements](REQUIREMENTS.md), [component inventory](INVENTORY.md), and candidate release-preflight source in `cmd/zeo-release`. An integrated installer and qualified distribution release are not available yet. Existing component capabilities do not establish that the complete platform works.

The installation target is one entry point for the complete Linux product, using compatible prebuilt component artifacts. Pomar supplies versioned execution environments; the distribution supplies the application platform inside them. Releases, environment profiles, organization settings and experiment configurations remain separately versioned.

Planned capabilities include onboarding, coordinated service operation, organization and seat isolation, upgrade and recovery, and reproducible research workflows. ARM64 and AMD64 support will be claimed only after release acceptance on each qualified backend. Docker is not a prerequisite for the intended standard path.

Implementation and verification are tracked separately. Public quickstart, contribution, security-reporting, licensing and release policies remain incomplete. Contributions must exclude credentials and private operational state.

Release preflight uses `zeo.release/v1` metadata: a distribution version, Linux platform, separately versioned environment profile, pinned components/artifacts, runtime dependencies, shared contracts, state compatibility, upgrade paths, agent tools and build provenance. It requires coverage of the core product responsibilities without assuming a daemon per component. Runtime integration dependencies may be cyclic; explicit installation ordering must be acyclic.

The tool rejects ambiguous JSON, mismatched targets, missing capabilities, conflicting pins, unsafe artifact locations and invalid local artifact bytes. `--artifacts` verifies local files; it does not download or extract them. Plans retain their platform/profile and a digest of the typed metadata representation; that digest is not a signature or a raw-manifest trust anchor. HTTPS URL checks are syntactic and do not establish public availability. A successful plan always reports `ready_for_install: false`: metadata and digest checks do not establish compatible runtime behavior, trusted build provenance, installation readiness or product acceptance.

Maintainer development requires Go 1.27.0 and the Linux race-test toolchain. `make verify` checks formatting, compiler/vet diagnostics and race-enabled tests using only standard-library dependencies. The recorded Linux source gate at `ddd4666fc895244d87770ab629778856bbb0201e` passed `make verify` with Go 1.27.0: vet and race-enabled tests passed in all five packages. The original output digest is `ad31f3462f88c57080f5ca987d37b5df97ef0bbd6e67b59dc079c9f573aa7331`. This source result does not qualify component artifacts, installed services or the integrated product. Use `zeo-release --help` for the preflight arguments after building it in the development environment; it is an assembly tool, not the user installation entry point.

The next milestone is one clean-host Linux installation with one working seat: supervised services, Runtime/Pomar organization and seat creation, App terminal attachment and resize, same-session reconnect across App/host-agent restarts, one bounded real-model task, and VM-failure visibility with owner notification within sixty seconds. None of these integrated demonstrations has passed yet.

Candidate guest service packaging in `cmd/zeo-service` renders a systemd unit for the existing Pomar session keeper. It addresses missing supervision without replacing the session implementation. It does not install or activate the unit, allocate identities, verify deployed artifacts or establish isolation. See [session service packaging](docs/session-service.md) for the layout, ownership requirements and pending qualification. The Linux source gate covers this renderer and its regressions; systemd unit verification and actual service/session behavior remain untested.

Candidate `cmd/zeo-stage` copies a local release-pinned Pomar executable into an existing trusted artifact root. It checks the copied size and digest, publishes without replacing an existing object, and validates existing content on reruns. It does not fetch a release, establish its trust, install OS dependencies or activate services. Its synthetic-fixture regressions passed in the same Linux source gate; a qualified Pomar artifact and real installation remain pending.

The complete installation profile must include its declared application workloads and seat tooling, consume runtime-owned signed releases, and finish with `make doctor` against the installed instance. That target is planned and is not implemented in the current Makefile. A missing required application or unknown required probe must prevent a green installation result.
