# ASHA Artifact Index

This is the canonical navigation surface for generated artifacts, curated summaries, publication files, visuals, and reproducibility documents.

## Root policy

The repository root is intentionally small. Generated audits and report outputs belong under `docs/`, not beside `README.md`.

## Artifact map

| Path | Kind | Owner | Purpose | Policy | Validation |
|---|---|---|---|---|---|
| `README.md` | `root-doc` | human overview | high-level project status, claim boundaries, and citation template | bounded edits only; avoid huge rewrites | manual review + doc markers |
| `QUICK_START.md` | `reproducibility` | operators | fast setup, targeted tests, and navigation commands | keep concise and command-oriented | go list / targeted go test examples |
| `GateResearcherMethod.md` | `root-doc` | method | theorem-gated research method | stable method reference | manual review |
| `docs/INDEX.md` | `index` | documentation | top-level documentation map | update when major doc folders change | link/path review |
| `docs/ARTIFACT_INDEX.md` | `index` | documentation | canonical index of generated and curated artifacts | primary artifact navigation surface | Gate 424 export |
| `docs/REPRODUCIBILITY_CHECKLIST.md` | `reproducibility` | reviewers/operators | targeted validation commands and no-full-suite policy | must avoid timeout-prone default commands | targeted go tests |
| `docs/architecture.md` | `root-doc` | architecture | large architecture ledger | append bounded addenda only | manual review + markers |
| `docs/audits/README.md` | `index` | audits | audit folder policy | all generated audits live under docs/audits | path review |
| `docs/audits/gates/INDEX.md` | `audit` | gates | gate audit index | new gate audits use gateNNN_registry_audit.md | count and missing-number check |
| `docs/audits/gates/gateNNN_registry_audit.md` | `audit` | gates | individual generated gate audit | never store in root | package theorem + rendered audit |
| `docs/audits/phenomenology/` | `audit` | phenomenology | empirical-quarantine phenomenology reports | do not promote empirical inputs | manual review |
| `docs/audits/final/` | `publication-support` | legacy/final | aggregate/final result snapshots | historical report area | manual review |
| `docs/summaries/` | `summary` | summaries | gate summaries and logical/ontological maps | curated prose, not generated gate audits | manual review |
| `docs/paper/` | `paper-workspace` | paper | paper drafts, final manuscript, references | publication workspace; no claim drift | review against claim audit |
| `docs/visuals/` | `visual-workspace` | visuals | source and exported figures/diagrams | keep source and exported files separate | figure checklist |
| `cmd/asha/` | `code` | CLI | command entrypoint | go list before broad execution | go list ./cmd/asha |
| `internal/app/` | `code` | registry | theorem registry wiring | avoid internal/app tests when timeout risk matters | go list ./internal/app |
| `pkg/bridge/` | `code` | bridge gates | bridge and publication-support theorem packages | targeted go test only | go test -p=1 ./pkg/bridge/<package> -count=1 |
| `pkg/matter/` | `code` | matter gates | matter/electroweak/yukawa packages | run selected package groups only | targeted matter go tests |

## Audit coverage

- Gate audits indexed: 228
- Gate range: G187--G425
- Known missing audit numbers: 191, 192, 198, 324, 329, 360, 388, 389, 390, 391, 392
- Phenomenology reports: 1
- Final aggregate reports: 1
- Gate audit index: `docs/audits/gates/INDEX.md`

## Firewalled boundaries

- Native charged flavor moduli remain `13`.
- Conditional K/X/Y family axiom ledger remains `9` symbolic charged coefficients.
- No Yukawa values, CKM angles, CP phase, PMNS parameters, cosmology coordinates, or quarantined axioms are promoted by this index.


## Gate 425 publication bundle additions

- `docs/paper/PUBLICATION_BUNDLE_PREFLIGHT.md` — paper-facing bundle readiness report.
- `docs/paper/BUNDLE_MANIFEST.md` — manifest of required paper/support artifacts.
- `docs/paper/SECTION_SOURCE_MAP.md` — manuscript section-to-gate source map.
- `docs/paper/FIGURE_SLOT_LEDGER.md` — reserved visual slots and source expectations.
- `docs/paper/CLAIM_FIREWALL_CHECKLIST.md` — publication claim/firewall checklist.
- `docs/paper/ASSEMBLY_CHECKLIST.md` — paper assembly workflow and targeted validation commands.

## Final paper artifacts

- `docs/paper/final/asha_paper_final_manuscript.docx` — finalized editable manuscript with source-code link, acknowledgments, final unified formula, and completed bibliography.
- `docs/paper/final/asha_paper_final_manuscript.pdf` — PDF export generated from the DOCX.
- `docs/audits/final/final_result.md` — formula-first final ASHA result ledger updated after Gates 0–425.

## Final Runtime Artifacts

| Artifact | Path | Purpose |
|---|---|---|
| Runtime package | `pkg/asha/` | Standalone final calculation/report API. |
| Runtime CLI | `cmd/asha/` | CI-safe command using `pkg/asha`. |
| Runtime docs | `docs/runtime/README.md` | Usage, scenarios, formats, and epistemology. |
| Runtime CI Markdown | `docs/runtime/reports/asha_runtime_ci.md` | Deterministic human-readable runtime report. |
| Runtime CI JSON | `docs/runtime/reports/asha_runtime_ci.json` | Machine-readable runtime report. |
| Runtime consolidation note | `docs/audits/final/runtime_step_gate425_20260514T0035.md` | Timestamped Gate-425 final runtime step. |

## Runtime verification update — 2026-05-14

The standalone runtime board was rerun against the latest embedded Gate-425 data. Reports are available at:

```text
docs/runtime/reports/asha_runtime_verification_20260514.md
docs/runtime/reports/asha_runtime_all_20260514.md
docs/runtime/reports/asha_runtime_ci_20260514.json
docs/audits/final/runtime_verification_gate425_20260514.md
```

The verification found no mismatch requiring manuscript updates.


## Runtime environment scenario update

- `docs/audits/final/runtime_environment_scenarios_gate425_20260514.md` — conditional dark/cosmology/vacuum-fate numeric scenario verification.
- `docs/runtime/reports/asha_runtime_environment_latest.md` — latest environment scenario runtime report.
