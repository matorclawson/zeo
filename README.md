# ZEO

**Updated:** 2026-11-18 · **Status:** Early distribution implementation; release qualification pending.

ZEO is being assembled as an open-source platform for persistent agent organizations, including ML research teams and small research laboratories. It aims to let users create independent organizations, connect models, set budgets and permissions, and inspect the results of coordinated work.

Available today: the [product requirements](REQUIREMENTS.md) and [component inventory](INVENTORY.md). An integrated installer and qualified distribution release are not available yet. Existing component capabilities do not establish that the complete platform works.

The installation target is one entry point for the complete Linux product, using compatible prebuilt component artifacts. Pomar supplies versioned execution environments; the distribution supplies the application platform inside them. Releases, environment profiles, organization settings and experiment configurations remain separately versioned.

Planned capabilities include onboarding, coordinated service operation, organization and seat isolation, upgrade and recovery, and reproducible research workflows. ARM64 and AMD64 support will be claimed only after release acceptance on each qualified backend. Docker is not a prerequisite for the intended standard path.

Implementation and verification are tracked separately. Public quickstart, contribution, security-reporting, licensing and release policies remain incomplete. Contributions must exclude credentials and private operational state.
