# Candidate App service profile

These units are source for review, **not a qualified installation or a runnable
switch packet**. No units have been installed, verified by systemd, started or
measured. Keep them disabled until the distribution's actual target binding,
state migration, resource evidence and guarded switch are qualified.

`zeoapp.slice` is the **one aggregate 384M budget** for API, Next, their stateless
memory observer and owned descendants. All three services explicitly join it.
The API and Next individual 384M ceilings
do not grant 768M: the parent remains 384M total (402653184 bytes). The 320M parent
soft threshold is a proposed measurement value, pending measured justification
and App/Messenger review. Node's 256MiB old-space flag is present on both actual
entrypoints; it is not total process RSS or an aggregate heap reservation.

The API invocation follows 41b's real `server/main.ts`: exactly one configuration
path argument, not an invented `--config` flag. Next retains the qualified
`--experimental-next-config-strip-types` start flag and listens on127.0.0.1:4173.
The private API configuration must bind 127.0.0.1:4380 and the dedicated service
UID. Paths/usernames below are proposed installation conventions, not observed
production-02 metadata:

- root-owned immutable app release exposed as `/opt/zeoapp/current`;
- authenticated, matching-native Node 24.22.1 at `/opt/zeoapp/runtime/bin/node`;
- dedicated `zeoapp` OS user/group, no builder-session dependence;
- private `/etc/zeoapp/runtime.json`, whose source custody checks remain required;
- service state under `/var/lib/zeoapp`, 0700, with the sole App SQLite DB.

Do not start these units against existing state merely to create or change
ownership. Systemd `StateDirectory` is not a migration/fence/backup helper. The
current 41b artifact still has production flat feed/pair/receipt state; it does
not satisfy the later SQLite-only rule. App owns its typed single-DB writer
and forward migration; Runtime authorizes narrowly without DB/session grants.

The candidate Go `zeo-app-memory-guard` only observes `/proc/meminfo` and exits
on a low, unavailable or malformed observation. It never spawns/stops processes,
uses a control API or writes service state. API, Next and their common target
use systemd `BindsTo=`/`After=` so an inactive observer stops the bound App,
without root privileges or a separate process supervisor. A startup observation
precedes monitoring. The proposed600MiB floor comes from Messenger's suggestion;
it is not yet adopted or measured. The proposed200ms interval is a polling bound,
not a claim that every instantaneous fluctuation is caught. Compile/package the
authenticated native helper as `/opt/zeoapp/libexec/zeo-app-memory-guard` before
any unit may start. Neither its authored tests nor the systemd stop propagation
have run yet. Include this third process in the same aggregate measurement.

Mandatory execution gaps remain:

- actual target OS/libc/principal/runtime/artifact/config/state/ingress binding;
- matching ARM runtime peakRSS, aggregate cgroup memory.peak and OOM events under
  the exact workload and source/artifact identity;
- finite systemd verification and actual readback that both services and every
  runtime worker are inside the budgeted slice;
- native continuous App-only MemAvailable abort enforcement and its measured threshold;
  OOMScoreAdjust 1000 is an App-first preference, not proof that Messenger cannot
  be killed; the new observer/dependency binding is still unqualified source;
- a reviewed exact state-safe switch, compatible rollback and Messenger review;
- Messenger's lean SQLite release and 24-hour zero-OOM soak before its 768 MB guard
  is adopted; the existing 3000 MiB guard is not silently lowered;
- one-file SQLite online/VACUUM backup and remote copy with a few MiB working
  memory, no tar/compressor/archive, and no backup overlapping the switch.

Starting the target orders processes; it does not prove health, ingress or
publication readiness. The bounded restart policy does not authorize replaying
unknown mutations. No global sshd/Messenger unit, cap, backup or ingress setting
is touched by these files. Newsletter belongs on the Mac under its own owner;
it is not part of this production-02 slice or target.

The candidate follows systemd 255's [hierarchical resource controls](https://www.freedesktop.org/software/systemd/man/255/systemd.resource-control.html),
[service execution semantics](https://www.freedesktop.org/software/systemd/man/255/systemd.service.html)
and [managed directory semantics](https://www.freedesktop.org/software/systemd/man/255/systemd.exec.html).
The observer's lifecycle dependency uses [BindsTo with ordering](https://www.freedesktop.org/software/systemd/man/255/systemd.unit.html#BindsTo=).
The actual target version and runtime behavior still require native verification.
