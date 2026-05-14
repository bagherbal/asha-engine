# Paper Bundle Manifest

This manifest lists the files and workspaces needed to assemble the final paper without claim drift.

| Path | Kind | Required | Source | Purpose | Claim rule | Readiness |
|---|---|---:|---|---|---|---|
| `docs/paper/PUBLICATION_BUNDLE_PREFLIGHT.md` | `checklist` | `true` | Gate 425 | top-level bundle readiness report | claim-neutral preflight only | ready |
| `docs/paper/BUNDLE_MANIFEST.md` | `manuscript` | `true` | Gate 425 | paper-facing artifact manifest | must link to existing source artifacts | ready |
| `docs/paper/SECTION_SOURCE_MAP.md` | `manuscript` | `true` | Gate 421 + Gate 425 | section-by-section source map | every claim section has a source and boundary | ready |
| `docs/paper/CLAIM_FIREWALL_CHECKLIST.md` | `boundary` | `true` | Gates 418--423 | publication firewall checklist | forbid promotion of sealed coordinates | ready |
| `docs/paper/ASSEMBLY_CHECKLIST.md` | `checklist` | `true` | Gate 425 | step-by-step paper assembly checklist | no claim drift while assembling | ready |
| `docs/ARTIFACT_INDEX.md` | `theorem-atlas` | `true` | Gate 424 | canonical repository artifact index | source of artifact paths | ready |
| `docs/REPRODUCIBILITY_CHECKLIST.md` | `checklist` | `true` | Gate 424 | targeted validation commands | avoid full-suite timeout by default | ready |
| `docs/audits/gates/gate420_registry_audit.md` | `theorem-atlas` | `true` | Gate 420 | theorem atlas and dependency graph | publication theorem source | ready |
| `docs/audits/gates/gate421_registry_audit.md` | `manuscript` | `true` | Gate 421 | manuscript skeleton | section skeleton source | ready |
| `docs/audits/gates/gate422_registry_audit.md` | `front-matter` | `true` | Gate 422 | executive claim audit | front-matter claim language | ready |
| `docs/audits/gates/gate423_registry_audit.md` | `reviewer-support` | `true` | Gate 423 | reviewer objection matrix | rebuttal boundaries | ready |
| `docs/audits/gates/gate424_registry_audit.md` | `checklist` | `true` | Gate 424 | artifact index audit | reproducibility support | ready |
| `docs/audits/gates/gate425_registry_audit.md` | `audit` | `true` | Gate 425 | this preflight audit | publication-support only | ready |
| `docs/summaries/essential_ontological_tower_map.md` | `proof-source` | `true` | curated summary | core logical tower | non-chronological orientation only | ready |
| `docs/summaries/gates_summary.md` | `proof-source` | `true` | curated summary | chronological gate summary | summary, not theorem replacement | ready |
| `docs/paper/drafts/` | `manuscript` | `false` | user-supplied | working manuscript drafts | drafts checked against claim audit | workspace ready |
| `docs/paper/final/` | `manuscript` | `false` | user-supplied | final manuscript outputs | only after claim/firewall review | workspace ready |
| `docs/paper/references/` | `reference` | `false` | user-supplied | bibliography and published-paper metadata | citation template preserved | workspace ready |
| `docs/visuals/source/` | `figure-slot` | `false` | user-supplied | editable figure sources | keep source separate from exported visuals | workspace ready |
| `docs/visuals/exported/` | `figure-slot` | `false` | user-supplied | publication figure exports | check captions against claim audit | workspace ready |
| `docs/visuals/diagrams/` | `figure-slot` | `false` | Gate 420 / user-supplied | dependency graphs and architecture diagrams | must match theorem atlas | workspace ready |
