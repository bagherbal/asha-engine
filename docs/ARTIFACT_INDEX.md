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

- Gate audits indexed: 321
- Gate range: G187--G546
- Known missing audit numbers: 191, 192, 198, 324, 329, 360, 388, 389, 390, 391, 392, 426, 427, 428, 429, 430, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 472, 479
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


## Gate 534 additions

- `pkg/bridge/generation2syntheticosreflectionpositivityadapter/` — synthetic OS reflection-positivity adapter package.
- `data/synthetic_os_reflection_positivity_ledger_gate534.json` — source-tagged bridge-only OS kernel fixture.
- `docs/audits/gates/gate534_registry_audit.md` — Gate 534 registry audit.


## Gate 535 additions

- `pkg/bridge/generation2oswickhilbertsectorclosureledger/` — OS/Wick/Hilbert sector closure ledger package.
- `docs/audits/gates/gate535_registry_audit.md` — Gate 535 registry audit and frontier map.

## Gate 536 additions

- `pkg/bridge/generation2physicalschwingerledgerairlock/` — physical Schwinger-function source ledger airlock package.
- `docs/audits/gates/gate536_registry_audit.md` — Gate 536 registry audit and source-ledger schema.


## Gate 537 additions

- `pkg/bridge/generation2syntheticschwingerledgeradapter/` — synthetic Schwinger-function source ledger adapter package.
- `data/synthetic_schwinger_function_ledger_gate537.json` — source-tagged bridge-only 19-row Schwinger fixture.
- `docs/audits/gates/gate537_registry_audit.md` — Gate 537 registry audit and synthetic dry-run result.

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

## Gate 538 additions

- `pkg/bridge/generation2schwingersourceauthenticityairlock/` — Schwinger source-authenticity comparator airlock preflight.
- `docs/audits/gates/gate538_registry_audit.md` — Gate 538 registry audit and authenticity/preflight result.

## Gate 539 additions

- `pkg/bridge/generation2syntheticsourceauthenticityadapter/` — synthetic source-authenticity ledger adapter and rejection dry run.
- `data/synthetic_source_authenticity_ledger_gate539.json` — source-tagged bridge-only 13-row authenticity fixture with canonical checksum.
- `docs/audits/gates/gate539_registry_audit.md` — Gate 539 registry audit and synthetic authenticity result.

## Gate 540 additions

- `pkg/bridge/generation2realschwingerimportswitchairlock/` — Gate 540 real Schwinger source import switch airlock preflight.
- `docs/audits/gates/gate540_registry_audit.md` — Gate 540 registry audit and default-off switch result.

## Gate 541 additions

- `pkg/bridge/generation2reallookingschwingersourcenegativecontroladapter/` — real-looking Schwinger source negative-control adapter package.
- `data/real_looking_schwinger_negative_control_ledger_gate541.json` — intentionally real-looking but untrusted bridge-only negative-control ledger.
- `docs/audits/gates/gate541_registry_audit.md` — Gate 541 registry audit and rejection result.

## Gate 542 additions

- `pkg/bridge/generation2realsourcecomparatorauthorizationairlock/` — real-source comparator authorization manifest airlock preflight.
- `docs/audits/gates/gate542_registry_audit.md` — Gate 542 registry audit and authorization-manifest result.

## Gate 543 additions

- `pkg/bridge/generation2syntheticauthorizationmanifestadapter/` — synthetic comparator authorization manifest adapter dry run.
- `data/synthetic_authorization_manifest_ledger_gate543.json` — source-tagged bridge-only 14-row authorization manifest fixture.
- `docs/audits/gates/gate543_registry_audit.md` — Gate 543 registry audit and synthetic authorization-manifest result.

## Gate 544 additions

- `pkg/bridge/generation2realsourcecomparatorharnessairlock/` — real-source comparator execution harness airlock preflight.
- `docs/audits/gates/gate544_registry_audit.md` — Gate 544 registry audit and comparator-harness contract result.

## Gate 545 additions

- `pkg/bridge/generation2syntheticcomparatorharnessadapter/` — synthetic comparator-harness result adapter dry run.
- `data/synthetic_comparator_harness_result_bundle_gate545.json` — checksum-protected bridge-only fake comparator output bundle.
- `docs/audits/gates/gate545_registry_audit.md` — Gate 545 registry audit and synthetic comparator-output quarantine result.

## Gate 546 additions

- `pkg/bridge/generation2comparatoroutputreleaseairlock/` — comparator-output release airlock preflight package.
- `docs/audits/gates/gate546_registry_audit.md` — Gate 546 registry audit and release-review schema.

## Gate 547 additions

- `pkg/bridge/generation2syntheticreleasereviewmanifestadapter/` — synthetic release-review manifest adapter dry-run package.
- `data/synthetic_release_review_manifest_gate547.json` — checksum-protected bridge-only 15-row synthetic release-review manifest.
- `docs/audits/gates/gate547_registry_audit.md` — Gate 547 registry audit and synthetic release-review rejection result.

## Gate 548 additions

- `pkg/bridge/generation2physicalcorrelationreleaseclosureledger/` — physical-correlation import/release sector closure ledger package.
- `docs/audits/gates/gate548_registry_audit.md` — Gate 548 registry audit.

## Gate 549 additions

- `pkg/bridge/generation2physicalcorrelationevidenceboardairlock/` — physical-correlation evidence-board airlock package.
- `docs/audits/gates/gate549_registry_audit.md` — Gate 549 registry audit and evidence-board schema result.

## Gate 550 additions

- `pkg/bridge/generation2syntheticevidenceboardadapter/` — synthetic physical-correlation evidence-board adapter dry-run package.
- `data/synthetic_evidence_board_manifest_gate550.json` — checksum-protected bridge-only 17-row synthetic evidence-board manifest.
- `docs/audits/gates/gate550_registry_audit.md` — Gate 550 registry audit and synthetic evidence-board rejection result.

## Gate 551 additions

- `pkg/bridge/generation2physicalcorrelationevidenceclosureledger/` — physical-correlation evidence-board sector closure ledger package.
- `docs/audits/gates/gate551_registry_audit.md` — Gate 551 registry audit and evidence-board sector closure result.
