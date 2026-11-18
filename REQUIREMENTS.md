# Requirements register

**Created:** 2026-11-17 · **Last-updated:** 2026-11-17 · **Status:** Initial baseline · **Doctrine:** CLAUDE.md Rev 18 body (header still says Rev 17).

Authority: operator mandate `MANDATE-zeo-distribution-2026-11-17.md`, received 2026-11-17, and `zeo-maintainer-SOW-01-zeo-distribution-charter.md`, revision a. The mandate wins over summaries here. Private charter location: `org-architect/projects/zeo/sow/zeo-maintainer/`.

IDs are permanent `ZEO-<mandate section>-<ordinal>` per SOW section 3. Never renumber or reuse them; append new IDs and retain superseded rows with a mapping. Confirmation with zeo-ledger is pending mailbox registration; any existing identifier mechanism must be mapped, without losing these IDs.

Each paragraph, list item and decision-table row is retained in source order, including constraints and qualifying context. Lead-in rows apply to the following items. This deliberately includes contextual statements to avoid losing requirements. Separate implementation and verification fields prevent source inspection or a merge from becoming a product acceptance claim.

Accountable owner for every row is zeo-maintainer; component delivery owners remain to be confirmed in the inventory. `—` means no implementation revision or verification evidence recorded. UNKNOWN dependencies and limitations are unresolved, not an assertion that none exist. Allowed states: Not started / In progress / Implemented / Verified / Blocked. Verified requires reproducible behavioral evidence. All rows start Not started as required by SOW section 8.

Section 00 assigns IDs to the unnumbered preamble requirements. Sections 01–19 follow the mandate numbering. The closing defining outcome is retained in section 19.

## 0. Unnumbered mandate preamble

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-00-01 | Establish the dedicated zeo-maintainer seat to deliver a coherent publicly distributable open-source product. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-00-02 | Use zeroemployeeorg/zeo as the flagship repository. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-00-03 | Deliver a working, installable, maintainable release; a proposal, links or manually assembled infrastructure is insufficient. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-00-04 | Create the formal SOW using the current authoritative organizational format and preserve the mandate requirements. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-00-05 | Assign permanent work identifiers through the existing mechanism and track implementation and verification separately. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 1. Strategic objective

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-01-01 | ZEO will provide infrastructure for persistent agent organizations, with particular focus on **ML research teams, small frontier research labs, and neolabs**. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-02 | Users must be able to install ZEO, create their own organization, configure model access, and have agents perform meaningful work under explicit budgets, permissions, and human oversight. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-03 | The same platform must support: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-04 | Our own ZEO development organization. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-05 | A separate Prof Rod organization. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-06 | A separate organization for configured Rasa work. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-07 | Independent organizations created by external adopters. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-08 | Disposable organizations used for experiments. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-09 | These organizations must use released software without depending on our private infrastructure or personal machine configuration. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-10 | Our public release must also establish the foundation for: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-11 | **Research conducted using ZEO:** implementing and reproducing ML publications, running experiments and ablations, and producing inspectable evidence. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-12 | **Research conducted on ZEO:** comparing organizational policies, communication structures, planning methods, memory mechanisms, and resource allocation. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-01-13 | Do not reduce this mission to a generic agent chat interface. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 2. Decisions established by the operator

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-02-01 | Treat these as requirements: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-02 | Dedicated seat — `zeo-maintainer` | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-03 | Flagship repository — `zeroemployeeorg/zeo` | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-04 | Repository direction — ZEO repositories are to become public and move back under `zeroemployeeorg` | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-05 | Supported execution platform — Linux | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-06 | Public installation — One entry point that installs and configures the complete required system | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-07 | Docker — Must not be a prerequisite for the supported installation | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-08 | Mac hosts — May provide Linux guests through Apple’s Virtualization framework | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-09 | Existing environment infrastructure — Reuse and extend Pomar where appropriate | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-10 | Product experience — One coherent distribution, with coordinated releases and lifecycle management | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-11 | Isolation — Enforced boundaries between organizations and seats | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-12 | Evidence — Completion requires reproducible verification | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-13 | Linux is the common execution platform regardless of whether it runs on a physical Linux machine, an existing VPS, or a Linux VM hosted on a Mac. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-14 | Existing implementation languages may remain where appropriate. A language change must solve a demonstrated problem and be justified separately. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-02-15 | A monorepo migration is not required. Preserve useful component boundaries while delivering one integrated product. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 3. Mandate of the `zeo` seat

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-03-01 | The `zeo` seat owns the **whole-product release and adoption experience**. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-02 | It is accountable for: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-03 | The flagship repository. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-04 | The supported platform specification. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-05 | Distribution assembly and release manifests. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-06 | Installation and onboarding. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-07 | Cross-component integration. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-08 | Release-level acceptance tests. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-09 | Upgrade, backup, restore, and uninstall workflows. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-10 | Public documentation and contributor onboarding. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-11 | Public release artifacts and release notes. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-12 | Coordination of the component work required to make releases function. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-13 | Existing component seats retain responsibility for their components. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-14 | However, `zeo` must actively drive integration defects to resolution. It may not declare its work complete because each component owner reports passing local tests. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-15 | The Architect arbitrates ownership disputes, missing ownership, and incompatible contracts. Every release-blocking dependency must have an owner and tracked work item. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-03-16 | The `zeo` seat must remain dedicated to this mandate rather than becoming a general pool for unrelated tasks. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 4. Begin with an evidence-backed inventory

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-04-01 | Inspect the actual repositories, installed services, and deployment configurations. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-02 | Produce an inventory covering at least: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-03 | Runtime. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-04 | Architect. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-05 | Messenger. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-06 | Agenda. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-07 | Controlling. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-08 | Connect. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-09 | App and its backend. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-10 | Pomar. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-11 | Connect and credential integrations. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-12 | Recorder, evidence, and artifact facilities. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-13 | Shared contracts and libraries. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-14 | Other components discovered to be required. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-15 | For each component, record: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-16 | Repository and revision — Exact source inspected | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-17 | Current owner — Responsible seat | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-18 | Purpose — Actual implemented responsibility | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-19 | Implementation — Languages, build tools, runtime dependencies | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-20 | Platform assumptions — Linux, macOS, architecture, host-specific integrations | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-21 | Execution model — Library, CLI, service, worker, scheduled job | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-22 | State — Databases, files, ownership, migrations | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-23 | Interfaces — APIs, events, contracts, authentication | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-24 | Deployment — Current installation and service supervision | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-25 | Dependencies — Other components and external services | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-26 | Publication blockers — Private dependencies, embedded data, missing rights or licensing decisions | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-27 | Release blockers — Missing functionality or unverified behavior | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-28 | Evidence — Commands, paths, tests, or observations supporting the entry | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-29 | Explicitly identify hardcoded references to personal usernames, hostnames, paths, accounts, domains, organization IDs, and private services. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-04-30 | Mark unknowns as unknown. Do not infer capabilities from names, old plans, or documentation without implementation evidence. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 5. Define the supported Linux platform

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-05-01 | Select and document: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-02 | The initial supported distribution and version. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-03 | CPU architectures supported at launch. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-04 | Minimum and recommended resources, supported by measurements. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-05 | Required system facilities. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-06 | Network requirements. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-07 | Storage requirements. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-08 | Privileged setup operations. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-09 | Supported model and agent authentication flows. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-10 | Design for Linux ARM64 and AMD64. Publish support claims only for architectures that pass the required release tests. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-11 | The supported product must not depend on macOS Keychain, launchd, AppleScript, macOS filesystem conventions, or other macOS-only facilities inside its Linux execution environment. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-12 | Host-side Mac virtualization support may use platform-specific code. Keep that adapter separate from Linux product logic. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-05-13 | Document any existing workflows that remain macOS-specific and provide a migration or optional integration plan. Do not silently remove necessary functionality. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 6. Define responsibilities between ZEO, Runtime, Pomar, and the OS

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-06-01 | Inspect current contracts before finalizing these boundaries. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-06-02 | The intended separation is: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-06-03 | **ZEO distribution:** assembles compatible components and manages installation and release lifecycle. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-06-04 | **Runtime:** manages agent execution and organizational runtime responsibilities. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-06-05 | **Pomar:** provides and manages execution environments through supported backends. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-06-06 | **Linux service management:** starts and supervises operating-system services. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-06-07 | **Architect:** coordinates organizational work. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-06-08 | Do not create duplicate schedulers, registries, credential stores, or environment managers without demonstrating why existing facilities cannot serve. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-06-09 | A repository does not automatically require a daemon. Choose process boundaries based on lifecycle, isolation, resource use, and failure behavior. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-06-10 | Maintain explicit versioned contracts between components. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 7. Build the complete distribution

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-07-01 | A ZEO release must identify an exact compatible combination of components. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-02 | The release manifest must record: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-03 | Distribution version. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-04 | Component versions and source revisions. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-05 | Platform and architecture. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-06 | Artifact checksums. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-07 | Required runtime dependencies. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-08 | Contract versions. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-09 | Database schema and migration compatibility. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-10 | Build provenance. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-11 | Supported upgrade paths. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-12 | Relevant agent-tool versions and compatibility constraints. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-13 | Normal installation must use built release artifacts. Do not require users to clone and compile all component repositories. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-14 | The flagship repository must contain the installer and distribution source, integration tests, documentation, examples, and the materials necessary to understand and reproduce release assembly. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-07-15 | Do not maintain manually copied component source trees that become divergent shadow repositories. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 8. Deliver one installation entry point

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-08-01 | On a clean supported Linux machine, the user must be able to begin installation through one documented entry point. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-02 | The installation must: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-03 | Detect compatibility and resources. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-04 | Explain required privileged changes. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-05 | Download and verify the selected release. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-06 | Install required dependencies through controlled, documented mechanisms. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-07 | Create service identities and protected state locations. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-08 | Configure component connectivity. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-09 | Install and start service definitions. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-10 | Run health and integration checks. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-11 | Present the onboarding interface or a clear remote-access method. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-12 | Report success only when the required system is usable. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-13 | It must handle partial installation, interruption, and reruns without duplicating organizations, corrupting state, or leaving unexplained changes. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-14 | Do not require the user to manually configure Messenger, Runtime, Architect, Agenda, Controlling, and App individually. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-15 | “One-click” allows necessary credentials and user choices. It does not allow manual infrastructure assembly. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-08-16 | A remote VPS installation must work without a desktop environment. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 9. Provide onboarding that reaches useful work

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-09-01 | Onboarding must guide the user through: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-09-02 | Creating an organization. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-09-03 | Configuring supported model or coding-agent access. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-09-04 | Establishing budget and concurrency limits. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-09-05 | Selecting or creating a workspace. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-09-06 | Starting an initial project. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-09-07 | Inspecting progress and results. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-09-08 | Responding to an agent’s request for input. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-09-09 | Never assume our personal provider accounts, subscription arrangements, credentials, or private services are available. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-09-10 | Document the authentication modes actually tested. Do not claim that API-key support proves subscription-based agent authentication works, or vice versa. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-09-11 | Provide a modest, bounded introductory workflow with a verifiable deliverable. It must exercise genuine coordination and execution rather than merely displaying service health. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 10. Support independent organization instances

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-10-01 | Organization identity must be explicit throughout the system. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-02 | Each instance must have its own: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-03 | Architect identity and continuity. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-04 | Agent identities. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-05 | Configuration. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-06 | Workspaces. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-07 | Messages. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-08 | Projects and commitments. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-09 | Credentials and access grants. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-10 | Budgets and usage records. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-11 | Memory and evidence. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-12 | Persistent state. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-13 | Backup and restore scope. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-14 | Shared binaries are acceptable. Implicitly shared authority or organizational state is not. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-15 | Support independent start, stop, inspection, backup, and removal. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-16 | The release design must explain whether different distribution versions can coexist on one host. If that is initially unsupported, enforce and document the limitation. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-10-17 | Prove isolation with two organizations running simultaneously. Include tests that attempt unauthorized cross-organization access. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 11. Preserve seat isolation inside Linux

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-11-01 | Do not equate seat count with VM count. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-02 | Support multiple seats inside a Linux environment with enforced identities, filesystem permissions, process controls, and resource limits. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-03 | Agent processes must not receive unrestricted host administration privileges by default. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-04 | Specify: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-05 | Seat identity allocation. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-06 | Workspace ownership. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-07 | Credential access. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-08 | Shared-resource access. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-09 | Process and session ownership. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-10 | CPU, memory, and process limits. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-11 | Seat creation, suspension, recovery, and deletion. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-12 | The mechanism for privileged operations. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-13 | Document the limits of shared-kernel isolation. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-11-14 | Use separate VMs where the workload or trust boundary requires stronger isolation. Avoid provisioning one VM per seat without a demonstrated need. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 12. Integrate Pomar and the Mac-hosted Linux path

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-12-01 | Inventory Pomar’s implemented capabilities before committing to a backend design. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-02 | Establish a tested route for: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-03 | A Linux target on an existing VPS. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-04 | A Linux guest hosted on an Apple Silicon Mac. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-05 | Disposable execution environments for experiments. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-06 | For Mac-hosted Linux, verify: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-07 | ARM64 guest compatibility. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-08 | VM creation and provisioning. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-09 | CPU and memory configuration. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-10 | Networking and access. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-11 | Disk persistence. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-12 | Start, stop, and restart. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-13 | Host reboot behavior. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-14 | Logging and failure reporting. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-15 | Backup and recovery. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-16 | Resource contention under representative workloads. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-17 | Do not assume nested virtualization or GPU access. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-18 | A management VM and worker VM may be a useful topology, but validate the choice against actual resource consumption. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-19 | The immediate intended migration is to run Architect and ZEO workloads in Linux on the Mac mini. Preserve a recovery route onto Hetzner. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-12-20 | Do not migrate the entire live organization before a representative pilot passes. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 13. Deliver lifecycle management

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-13-01 | The product must support: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-02 | Status — Show component and organization health with actionable failures | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-03 | Start/stop — Respect dependencies and preserve state | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-04 | Reboot — Recover required services and identify interrupted work | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-05 | Upgrade — Check compatibility, preserve data, and apply coordinated migrations | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-06 | Failed upgrade — Leave a diagnosable and recoverable state | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-07 | Backup — Capture a consistent, documented recovery set | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-08 | Restore — Recover onto a clean supported host | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-09 | Diagnostics — Produce useful reports with secrets excluded | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-10 | Uninstall — Remove installed services cleanly and explicitly handle retained data | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-11 | Do not promise rollback across irreversible schema migrations. Define when recovery requires restoring a backup. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-12 | For cross-component backups, specify how consistency is established. Copying live database and state files without a consistency mechanism is insufficient. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-13-13 | Cross-architecture recovery must be tested or explicitly limited. Do not assume that process checkpoints or tool sessions are portable between ARM64 and AMD64. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 14. Establish the experimentation capability

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-14-01 | Create a bounded implementation plan and deliver an initial end-to-end research workflow. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-02 | The experimentation capability must represent: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-03 | Research question and hypothesis. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-04 | Protocol and baselines. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-05 | Workload and dataset identity. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-06 | Source revisions and environment. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-07 | Model and agent configuration. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-08 | Seeds where applicable. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-09 | Budgets and stopping rules. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-10 | Runs and retries. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-11 | Metrics and evaluation. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-12 | Artifacts and provenance. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-13 | Failures and human interventions. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-14 | Conclusions linked to evidence. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-15 | Expose a researcher-friendly Python interface or integration path. Reuse established research tools where appropriate rather than rebuilding training frameworks or metric trackers. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-16 | Separate: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-17 | Distribution version. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-18 | Organization configuration. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-19 | Experimental treatment. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-20 | Workload and evaluator version. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-21 | For example, comparing GTD with another policy must be possible on the same software release. Do not confound a policy change with unrelated development changes. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-22 | Provide one small ML reproduction or experiment workflow and one organizational-policy comparison workflow. Record limitations honestly, including hosted-model variability and any barriers to exact replay. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-14-23 | Experimental instances must not inherit production credentials or external write permissions by accident. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 15. Move and publish the repositories

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-15-01 | Inventory every repository in the ZEO scope and propose the definitive destination mapping under `zeroemployeeorg`. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-02 | Include dependencies and auxiliary repositories required to build, run, test, or understand the public product. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-03 | Prepare and execute the transfers and public releases through the available authorized mechanisms. Report concrete blockers if platform permissions or ownership prevent execution. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-04 | Before changing visibility: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-05 | Inspect tracked contents and history for secrets and private operational data. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-06 | Remove publication blockers; rotate exposed credentials when required. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-07 | Resolve private dependency links. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-08 | Update module paths, package references, documentation, CI, and release automation. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-09 | Record ownership and licence status of included code and assets. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-10 | Identify the open-source licence decision needed where one is not already established. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-11 | Verify the public system can build and run without private repository access. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-12 | Publishing source code does not include publishing organizational databases, customer or employer data, private messages, personal credentials, or live configuration. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-13 | Do not rewrite shared history casually. Where remediation requires history changes, specify the impact and coordinate recovery for existing clones. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-15-14 | The public flagship must clearly distinguish what is available, experimental, and planned. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 16. Make external contribution practical

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-16-01 | Provide: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-02 | A clear product explanation. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-03 | A working quickstart. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-04 | An architecture overview. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-05 | Supported platform and resource information. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-06 | A contributor setup that works without our organization. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-07 | Component ownership and issue-routing guidance. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-08 | Contract-change and compatibility procedures. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-09 | Release and support policies. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-10 | A security-reporting route. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-11 | Documented limitations. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-12 | Public examples and reproducible acceptance workflows. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-16-13 | An external contributor must be able to start a development instance without obtaining our private state or credentials. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 17. Required verification

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-17-01 | A release candidate is not accepted until the following are demonstrated: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-02 | **Clean installation:** complete installation on a fresh supported Linux host. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-03 | **Useful execution:** a configured organization completes a verifiable project. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-04 | **No private dependency:** installation and operation succeed without access to our private repositories or infrastructure. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-05 | **Reboot recovery:** services restart and interrupted work is represented accurately. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-06 | **Component failure:** failures are surfaced and recovery does not silently duplicate work. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-07 | **Organization isolation:** two instances coexist without unauthorized state or credential access. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-08 | **Seat isolation:** worker identities cannot access another seat’s protected resources. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-09 | **Upgrade:** a supported previous release upgrades with state preserved. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-10 | **Restore:** a consistent backup restores onto another clean host. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-11 | **Mac guest:** the Linux release works inside the supported Mac virtualization path. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-12 | **Representative concurrency:** load testing distinguishes registered seats, idle sessions, and actively executing workloads. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-13 | **Research evidence:** an experiment exports a traceable result package. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-14 | **Uninstall:** removal behaves as documented. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-15 | **Public build:** an external-style checkout can reproduce the documented build and checks. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-16 | Use real model calls for at least the bounded end-to-end acceptance workflow. Deterministic fixtures are appropriate for routine regression tests but must not be presented as proof of live operation. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-17-17 | Record actual measurements. Do not declare that 50 active seats are supported based only on creating 50 identities. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 18. Delivery sequence

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-18-01 | Organize the work into these milestones: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-02 | **A. Inventory and ownership** | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-03 | Produce the repository-backed inventory, dependency map, platform gaps, publication blockers, and assigned work. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-04 | **B. First complete Linux installation** | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-05 | Deliver a release candidate that installs the necessary system and completes a bounded workflow on a clean Linux host. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-06 | **C. Lifecycle and isolation** | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-07 | Prove reboot recovery, independent organizations, seat isolation, upgrade, backup, restore, and uninstall. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-08 | **D. Mac-hosted Linux pilot** | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-09 | Run the same release through the supported Pomar/Mac path. Migrate representative seats before scaling. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-10 | **E. Research workflows** | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-11 | Deliver the initial experimentation and publication-reproduction capabilities with traceable outputs. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-12 | **F. Public release** | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-13 | Complete repository migration and publication, public installation tests, documentation, and release artifacts. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-18-14 | Prepare repository publication work in parallel where dependencies allow. Do not postpone discovering publication blockers until the end. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

## 19. Reporting and accountability

| Work ID | Mandate requirement / context | Owner | State | Implementation revision | Verification evidence | Known limitations | Remaining dependencies |
|---|---|---|---|---|---|---|---|
| ZEO-19-01 | For every requirement, track: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-02 | Owner. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-03 | Work identifier. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-04 | Status. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-05 | Implementation revision. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-06 | Verification evidence. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-07 | Known limitations. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-08 | Remaining dependencies. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-09 | Use explicit states: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-10 | **Not started / In progress / Implemented / Verified / Blocked.** | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-11 | “Implemented” does not mean “verified.” “Merged” does not mean “deployed.” “Services running” does not mean “the organization works.” | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-12 | The final release report must include: | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-13 | The exact distribution and component revisions. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-14 | Installation artifacts and public locations. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-15 | Supported configurations. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-16 | Acceptance-test results. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-17 | Observed resource usage. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-18 | Upgrade and recovery evidence. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-19 | Repository migration status. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-20 | Outstanding limitations. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-21 | Every unfulfilled requirement and its reason. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-22 | Your first response should confirm creation of the `zeo` seat and SOW, identify the actual repositories inspected, and provide the initial work breakdown and blockers. Continue into implementation within this mandate; do not stop after returning a roadmap. | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |
| ZEO-19-23 | **The defining outcome is that someone outside our organization can install ZEO on a supported Linux environment, establish an independent agent organization, perform useful research work, and operate and recover the system without our personal assistance.** | zeo-maintainer | Not started | — | — | UNKNOWN | UNKNOWN |

