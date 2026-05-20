# ASHA Runtime Package

This directory documents the standalone final runtime layer updated through Gate 570.

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

The runtime marker and latest `all`/`ci` reports were refreshed against the current embedded Gate-562 data. The runtime layer remains a reporting board; Gate 551 adds no new physical constants; it validates the synthetic evidence-board manifest parser, checksum, governance metadata, and zero-native-delta path. It boards no real evidence, imports no real correlation source, and promotes no physical dynamics natively.

See:

```text
docs/runtime/reports/asha_runtime_verification_20260515.md
```

Gate 556 adds a selector-origin audit. It updates the runtime marker only; it does not modify prior ASHA numerical calculations.

Gate 557 adds an eta-trace representative and record-algebra audit. It updates the runtime marker only; it does not modify prior ASHA numerical calculations.

Gate 558 adds an eta-record `End(H_phi)` matrix certificate and product-closure audit. It updates the runtime marker only; it does not modify prior ASHA numerical calculations.

Gate 559 adds an eta-record transfer rank/trace obstruction audit. It updates the runtime marker only; it does not modify prior ASHA numerical calculations.
Gate 560 adds a Pauli-Hopf scalar moment-map audit. It updates the runtime marker only; it does not modify prior ASHA numerical calculations.

Gate 561 adds a Pauli moment to weak-plane incidence obstruction audit. It updates the runtime marker only; it does not modify prior ASHA numerical calculations.

Gate 563 adds a Pauli-Hopf to quaternionic weak-socket audit. It recognizes a structural scalar/quaternionic moment-map bridge to `Im(H)` and updates the runtime marker only; it does not modify prior ASHA numerical calculations or derive physical electroweak dynamics.

- Gate 563 runtime marker: `gate564-symbolic-electroweak-hessian-bridge-audit-20260516`.

Latest runtime marker: Gate 565 — boundary gauge-normalization to electroweak Hessian alignment; no observed low-energy W/Z/photon or flavor prediction is promoted.

Latest runtime marker: Gate 566 — contact/Reeb law-space clock and product-time airlock; K_7 is certified but no contact form, Reeb vector, physical time, RG scale, OS/Wick/Hilbert dynamics, or arrow of time is promoted.


Gate 567 updates the runtime marker after adding the contact form certificate and distinguished covector obstruction audit.

Gate 568 updates the runtime marker after adding the finite contact differential source search audit. It does not modify prior numerical calculations; it blocks Boolean incidence, G2 calibration, projector relative-position data, q4 spectral data, and exterior-language notation from being promoted to a native finite `d` on `K_7`.

Gate 569 updates the runtime marker after adding the finite contact cochain-complex and `d²=0` certificate audit. It proves the current unsigned Boolean incidence fails the differential `d²=0` test and still does not provide a signed finite `d` on `K_7`.


Gate 570 updates the runtime marker after adding the Witt/Fock Hopf `S^7` contact form and Reeb phase audit. It certifies Hopf contact/Reeb phase on the Fock unit sphere, but does not identify it with `K_7`, physical time, RG scale, OS/Wick/Hilbert dynamics, or observed history.

Gate 571 updates the runtime marker after adding the Hopf `S^7` to Boolean-octonionic `K_7` functor and product-time airlock obstruction audit. It preserves the separation between Witt/Fock Hopf phase, Boolean-octonionic `K_7`, physical time, RG scale, OS/Wick/Hilbert dynamics, and observed history.
