# ASHA Runtime Package

This directory documents the standalone final runtime layer updated through Gate 551.

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

## Verification run — 2026-05-15

The runtime marker and latest `all`/`ci` reports were refreshed against the current embedded Gate-550 data. The runtime layer remains a reporting board; Gate 551 adds no new physical constants; it validates the synthetic evidence-board manifest parser, checksum, governance metadata, and zero-native-delta path. It boards no real evidence, imports no real correlation source, and promotes no physical dynamics natively.

See:

```text
docs/runtime/reports/asha_runtime_verification_20260515.md
```
