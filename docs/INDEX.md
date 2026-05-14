# Documentation Index

Canonical documentation and generated artifacts are organized as follows.

```text
README.md                         # project overview and claim boundaries
QUICK_START.md                    # build, test, and navigation guide
GateResearcherMethod.md           # theorem-gated research method

docs/
├── architecture.md               # detailed architecture ledger
├── audits/                       # generated audit artifacts
│   ├── gates/                    # gate registry audits and index
│   ├── phenomenology/            # empirical-quarantine phenomenology reports
│   └── final/                    # legacy/final aggregate result reports
├── summaries/                    # gate summaries and ontological tower maps
├── paper/                        # manuscript drafts, final paper, references
└── visuals/                      # diagrams, figures, and visual source files
```

## Artifact policy

Generated gate audits and report outputs should not be committed to the repository root. Put them under `docs/audits/`.

High-level conceptual summaries belong under `docs/summaries/`.

Publication drafts and final paper files belong under `docs/paper/`.

Visuals belong under `docs/visuals/`.
