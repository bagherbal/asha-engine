# Paper Workspace

Publication-facing material belongs here, not in the repository root.

Suggested layout:

```text
docs/paper/
├── drafts/       # working manuscript drafts
├── final/        # submitted or published manuscript versions
└── references/   # BibTeX, citation notes, external reference ledgers
```

Recommended convention:

- Keep the theorem/code repository independent from paper drafts.
- Use `docs/summaries/` and `docs/audits/` as source material.
- Keep any published DOI/arXiv references mirrored in the README citation template.


## Gate 425 paper-facing files

- `PUBLICATION_BUNDLE_PREFLIGHT.md` — readiness report.
- `BUNDLE_MANIFEST.md` — required and optional paper-support artifact manifest.
- `SECTION_SOURCE_MAP.md` — section-to-gate source map.
- `FIGURE_SLOT_LEDGER.md` — figure/diagram slots.
- `CLAIM_FIREWALL_CHECKLIST.md` — allowed and forbidden claim language.
- `ASSEMBLY_CHECKLIST.md` — draft assembly workflow and targeted validation commands.

## Final manuscript artifacts

The current finalized manuscript artifacts are stored in `docs/paper/final/`:

- `asha_paper_final_manuscript.docx` — editable Word manuscript.
- `asha_paper_final_manuscript.pdf` — PDF export generated from the DOCX.

The manuscript source link is documented inside the paper under **Source Code and Reproducibility**. The self-authored repository entry is intentionally not included as a bibliography item; if the repository is archived later, cite the DOI/archive in the References.
