# One-line installation contract

**Created:** 2026-11-19 · **Status:** Specification; no conforming integrated ZEO installer is qualified.

Every ZEO product must offer one documented installation entry point. The application owner implements its installer; distribution maintenance owns the common contract and integration with host services. A command example is a product promise only after its acceptance has passed on the declared Linux profile.

## Entry point and result

A released product provides a published bootstrap command or one installation target. The entry point obtains and verifies a pinned release for the actual OS and architecture, checks prerequisites, configures its dedicated service identity and private state, installs its supervisor, and runs its declared health check. Authentication comes through the supported user flow or an admitted credential reference; credentials and private configuration never appear in command output.

The installation must survive the invoking terminal closing and a host restart. It must not depend on a builder's checkout, running development session, private network, or unqualified replacement for a component's supported installer. Distribution installation composes component installers and their version contracts through the same entry point.

Before effects, the installer reports its intended service, state, ingress and replacement actions. Missing prerequisites and unknown state stop with a specific cause. Repeated invocation reconciles a completed installation or refuses an uncertain prior operation; it never repeats a possibly committed mutation blindly. Exit zero requires the declared installed health and version checks, rather than successful file copying alone.

## Replacement and retained tasks

The operator may choose a fresh replacement that retains tasks only. The plan names the observed legacy task authority and the exact tables, relations and files needed to retain those tasks. The application owner defines the export/import schema; the host installer does not guess it from a filename or copy only the main file of a live WAL database.

Before removing the legacy application:

1. Fence the actual legacy task writers and create a coherent export. Record the schema revision, per-table and attachment counts, stable task IDs, byte counts and content hashes without printing task contents or credentials.
2. Copy the export to both approved backup destinations. Verify each full copy and record its receipt. Use the SQLite snapshot and bounded-memory rules in requirements section 25; a mutable mirror or tar archive is not a replacement service-state authority.
3. Restore in an isolated location through the proposed importer. Verify the selected task content, identities, relationships, statuses and attachment hashes against the export. A zero count requires an observed and reviewed explanation.
4. Require both reader clearances for the concrete go-live script. Its removal phase may affect only the observed legacy application resources covered by the operator's authorization. Preserve the verified task backups and recovery receipts.
5. Install the new application, import the retained tasks, and verify its actual health and intended task interface. Preserve newly accepted writes during any recovery; never restore a stale database over them.

An export or restore failure prevents removal. A partially completed replacement preserves evidence and reports which service is active, which tasks were imported, and whether reconciliation is required. A restored origin and a publicly visible recovered application are separate checks.

## Release acceptance

Qualification records the immutable product and installer revisions, release hashes, Linux OS/architecture, supervisor and service identity, installation exit, installed health/version, restart survival, and measured resource limits. A task-retaining replacement additionally records both backup receipts, isolated restore results, import counts and a real task readback. Installation tests use an admitted Linux venue; source checks and cross-compilation alone do not establish installation acceptance.

Document the actual command only when these results exist. Until then, identify the missing implementation and evidence and label proposed commands as unqualified. No operator-use date is established by a source estimate or a passing component gate alone.

Tracked by ZEO-02-06 and ZEO-27-01 through ZEO-27-10 in [REQUIREMENTS.md](../REQUIREMENTS.md).
