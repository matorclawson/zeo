# Guest session service packaging

Updated: 2026-11-18. Status: candidate source; not compiled, tested or activated.

The one-seat installation needs the session keeper to run inside the seat's Linux environment, independently of App and the host agent. The existing Pomar interface is `pomar session serve -root DIR` in the foreground. `zeo-service` renders its systemd system unit. It does not implement a second session manager or installer.

Inputs are organization and seat names, Runtime-allocated user/group names, the qualified Pomar executable's SHA256, and explicit memory/task/CPU limits from the environment profile. Names must be lowercase safe identifiers of at most 32 characters. Resource numbers are requested caps, not measured minimums or support claims. One CPU is 100 percent.

The candidate layout is:

| Item | Path or policy |
|---|---|
| Pinned executable | `/opt/zeo/artifacts/<sha256>/pomar` |
| Seat home | `/var/lib/zeo-seats/<organization>/<seat>/home` |
| Workspace | `/var/lib/zeo-seats/<organization>/<seat>/workspace` |
| Session root | `/var/lib/zeo-seats/<organization>/<seat>/sessions` |
| Socket | `sessions.sock` inside the session root |
| Supervisor | System systemd instance inside the guest |
| Ownership | Explicit non-root user/group; private state directories and umask |
| Lifecycle | Boot target; restart on failure; no App/host-agent dependency |

This layout is a candidate for owner adoption, not a completed environment-profile contract. Runtime must allocate distinct UIDs and GIDs for distinct organizations/seats; the renderer cannot detect reuse across calls. A name is not an isolation boundary. App/host-agent attachment must enter the approved OS identity and session socket boundary through Pomar's authenticated interface, not weaken its same-UID check or expose the socket indiscriminately.

The future installer must verify the executable against the trusted release pin, target and qualified source. The digest in a pathname does not check its contents. The executable and its ancestors must be administrator-owned and not writable by the seat. Verify no preexisting symlinks or conflicting custody in any state path before activation: systemd's StateDirectory can recursively change ownership of existing directories. Provisioning must refuse conflicts rather than take over arbitrary state. Service names must identify the organization/seat pair without collisions. No secret values belong in the generated unit.

Type=exec detects failed process execution; it does not establish that the session socket is healthy. After startup the installer must use the owner's readiness interface under the allocated identity and fail installation if readiness is absent. The unit sets HOME and POMAR_SESSION_ROOT; every client must address that same root explicitly or through its approved environment. Profile tool paths, model credential delivery and their writable storage need owner qualification. The read-only filesystem policy may expose missing tool cache/state paths; test the real bounded task before adoption.

App and host-agent restarts must leave the guest keeper running and reconnect to its existing immutable session. Restarting THIS unit, losing the guest or rebooting it destroys live PTYs; retained metadata/logs do not restore those processes. Restart=on-failure provides service availability, not session continuity across keeper failure. Never turn an automatic restart or a fresh session with the same name into a continuity claim.

On an admitted Linux gate, `make verify` must first compile and run the packaging regressions at the exact committed head. Then qualify the rendered unit with `systemd-analyze verify` on the selected target using the actual installed Pomar executable and identities. Unit parsing alone is insufficient: on an approved disposable environment, activate it, confirm socket ownership/readiness, create a real session through Pomar, attach and resize through App, detach/reconnect, restart App and host agent separately, and verify the same session and task remain. Record the exact unit, versions, identity mapping, limits, timestamps and original receipts. Guest stop is a separate failure-visibility test and must leave the platform observer/App/notification path running.

The complete one-entry-point platform installer, compatible component set, identity allocation, integration wiring and seven-point pilot acceptance remain open work. A rendered unit does not report installation success.
