# Reproducibility Checklist

Use targeted validation by default. Avoid broad test commands unless a full validation pass is intentional.

## Commands

| Name | Run by default | Risk | Command | Purpose |
|---|---:|---|---|---|
| CLI wiring | `true` | low | `go list ./cmd/asha` | confirm command entrypoint resolves |
| Registry wiring | `true` | low | `go list ./internal/app` | confirm app registry imports resolve without running timeout-prone tests |
| New gate package | `true` | low | `go test -p=1 ./pkg/bridge/publicationbundlepreflight -count=1` | validate Gate 424 package only |
| Publication support bridge group | `true` | medium | `go test -p=1 ./pkg/bridge/artifactindexexport ./pkg/bridge/reviewerobjectionmatrix ./pkg/bridge/executiveabstractclaimaudit ./pkg/bridge/manuscriptskeletonexport ./pkg/bridge/publicationtheorematlas -count=1` | validate recent publication-support chain |
| Selected matter guardrail | `true` | medium | `go test -p=1 ./pkg/matter/yukawaintertwiner ./pkg/matter/trialityyukawa ./pkg/matter/texture ./pkg/matter/generationbreak ./pkg/matter/hypercharge ./pkg/matter/su2l -count=1` | preserve flavor/matter boundary packages |
| Full suite | `false` | high / timeout-prone | `go test ./...` | expensive full validation only when intentionally needed |
| internal/app tests | `false` | high / timeout-prone | `go test ./internal/app` | avoid when timeout risk matters; use go list instead |

## Policy

- Prefer targeted package tests over full-suite runs.
- Use go list ./internal/app to check registry wiring when internal/app tests are timeout-prone.
- Gate audits belong in docs/audits/gates, never in the repository root.
- Generated publication-support gates must not introduce new physics claims.
- Flavor/cosmology firewalls must remain explicit in every export-oriented artifact.

## Minimal validation for this export

```bash
go test -p=1 ./pkg/bridge/publicationbundlepreflight -count=1
go list ./internal/app
```


## Gate 425 publication bundle validation

```bash
go test -p=1 ./pkg/bridge/publicationbundlepreflight -count=1
go test -p=1 ./pkg/bridge/publicationbundlepreflight ./pkg/bridge/artifactindexexport ./pkg/bridge/reviewerobjectionmatrix ./pkg/bridge/executiveabstractclaimaudit ./pkg/bridge/manuscriptskeletonexport ./pkg/bridge/publicationtheorematlas -count=1
go list ./internal/app
```
