# Documentation Index

Canonical documentation and generated artifacts are organized as follows.

```text
README.md                         # project overview, claim boundaries, citation template
QUICK_START.md                    # fast setup, targeted tests, navigation guide
GateResearcherMethod.md           # theorem-gated research method

docs/
├── INDEX.md                      # this documentation map
├── ARTIFACT_INDEX.md             # canonical artifact/navigation index
├── REPRODUCIBILITY_CHECKLIST.md  # targeted validation commands and run policy
├── ARTIFACT_MAINTENANCE_CHECKLIST.md
├── PUBLICATION_WORKSPACE.md
├── paper/PUBLICATION_BUNDLE_PREFLIGHT.md # paper-facing preflight report
├── architecture.md               # detailed architecture ledger (Needs cleanup)
├── audits/                       # generated audit artifacts
│   ├── gates/                    # gate registry audits and index
│   ├── phenomenology/            # empirical-quarantine phenomenology reports
│   └── final/                    # legacy/final aggregate result reports
├── summaries/                    # gate summaries and ontological tower maps
├── paper/                        # manuscript drafts, final paper, references
└── visuals/                      # diagrams, figures, and visual source files
```

## Start here

- [`QUICK_START.md`](../QUICK_START.md) — how to run targeted tests and avoid timeout-prone commands.
- [`ARTIFACT_INDEX.md`](ARTIFACT_INDEX.md) — canonical map of generated artifacts, summaries, paper workspace, visuals, and code locations.
- [`paper/PUBLICATION_BUNDLE_PREFLIGHT.md`](paper/PUBLICATION_BUNDLE_PREFLIGHT.md) — final paper assembly preflight and bundle readiness report.
- [`REPRODUCIBILITY_CHECKLIST.md`](REPRODUCIBILITY_CHECKLIST.md) — exact command policy for targeted validation.
- [`audits/gates/INDEX.md`](audits/gates/INDEX.md) — gate audit index.

## Artifact policy

Generated gate audits and report outputs should not be committed to the repository root. Put them under `docs/audits/`.

High-level conceptual summaries belong under `docs/summaries/`.

Publication drafts and final paper files belong under `docs/paper/`.

Visual source files and exported figures belong under `docs/visuals/`.

## Runtime Endpoint

- `docs/runtime/README.md` — standalone runtime package and CLI documentation.
- `docs/runtime/reports/asha_runtime_ci.md` — deterministic CI report in Markdown.
- `docs/runtime/reports/asha_runtime_ci.json` — deterministic CI report in JSON.
- `docs/audits/final/runtime_step_gate425_20260514T0035.md` — final runtime consolidation note.

