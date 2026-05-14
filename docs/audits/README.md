# Audit Artifacts

This folder stores generated audit and registry artifacts. These are outputs of theorem-gated runs, not source code.

## Layout

```text
docs/audits/
├── gates/          # gateNNN_registry_audit.md files + index
├── phenomenology/  # empirical-quarantine phenomenology reports
└── final/          # legacy/final aggregate result reports
```

## Policy

- Do not place generated audits in the repository root.
- Keep root-level files limited to project entry points such as `README.md`, `QUICK_START.md`, `go.mod`, and source folders.
- If an audit is referenced from papers or summaries, reference it through this folder.
