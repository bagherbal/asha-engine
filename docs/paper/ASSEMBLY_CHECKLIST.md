# Paper Assembly Checklist

1. Start from the section source map in `docs/paper/SECTION_SOURCE_MAP.md`.
2. Draft or place manuscript files under `docs/paper/drafts/`.
3. For every section, cite the corresponding gate audit or summary source.
4. Add figures only through the slots in `docs/paper/FIGURE_SLOT_LEDGER.md`.
5. Check every claim against `docs/paper/CLAIM_FIREWALL_CHECKLIST.md`.
6. Preserve the README published-paper citation template until publication metadata is known.
7. Run targeted publication-support tests before finalizing:

```bash
go test -p=1 ./pkg/bridge/publicationbundlepreflight -count=1
go test -p=1 ./pkg/bridge/publicationbundlepreflight ./pkg/bridge/artifactindexexport ./pkg/bridge/reviewerobjectionmatrix ./pkg/bridge/executiveabstractclaimaudit ./pkg/bridge/manuscriptskeletonexport ./pkg/bridge/publicationtheorematlas -count=1
go list ./internal/app
```

Do not run `go test ./...` by default when timeout risk matters.
