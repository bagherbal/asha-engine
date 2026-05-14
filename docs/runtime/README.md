# ASHA Runtime Package

This directory documents the standalone final runtime layer introduced after Gate 425.

The historical gate packages remain as theorem/audit references. The runtime endpoint is now:

```text
pkg/asha     standalone calculation/report package
cmd/asha     runtime CLI using pkg/asha
```

## Primary command

```bash
go run ./cmd/asha --scenario all --format text --strict
```

## Supported scenarios

```text
all
native
higgs
family
dark-stable-thermal
cosmology
environment
ci
```

## Supported formats

```text
text
markdown
json
```

## Example CI usage

```bash
go test -p=1 ./pkg/asha ./cmd/asha -count=1
go run ./cmd/asha --scenario ci --format json --strict > docs/runtime/reports/asha_runtime_ci.json
go run ./cmd/asha --scenario environment --format markdown --strict > docs/runtime/reports/asha_runtime_environment_latest.md
```

## Runtime epistemology

The runtime package separates:

```text
native/audited       finite law-space and exact gate outputs
bridge-required      continuum/spectral-action coefficient lanes
quarantined-axiom    K/X/Y family-capacity assumptions
environmental        empirical/flavor/cosmology coordinates
failed-route         rejected interpretations and firewalls
```

It does not promote any quarantined axiom to native ASHA theorem.

## Verification run — 2026-05-14

The runtime board was rerun across all scenarios (`all`, `native`, `higgs`, `family`, `dark-stable-thermal`, `cosmology`, `environment`, `ci`) against the current embedded Gate-425 data. The environment scenario now reports conditional numerical paths for dark-sector yield constraints, cosmological subtraction severity, holographic/dilaton scale, and vacuum-fate stress tests. These remain bridge/environmental diagnostics, not native predictions.

See:

```text
docs/runtime/reports/asha_runtime_verification_20260514.md
```
