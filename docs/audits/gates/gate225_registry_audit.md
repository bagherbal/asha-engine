# Gate 225 Registry Audit — Finite anchor Dark Matter viability / ALP and Dark Sector audit

## Gate identity

```text
Gate: 225
Package: pkg/bridge/finiteanchordm
Theorem: FiniteAnchorDarkMatterViabilityAuditTheorem
Registry ID: BRIDGE-FINITE-ANCHOR-DARK-MATTER-VIABILITY-AUDIT
Primary status: FAILED_ROUTE
```

## Executive result

Gate 225 inherits the Gate-224 result that the sealed PeV threshold carriers decay before BBN and therefore cannot be present-day dark matter:

```text
Ω_heavy h² = 0
```

It then audits whether the remaining finite anchors can supply the missing dark matter:

```text
B-sector first spectral gap = 0.1024649212
seven contact partial-overlap modes
```

The result is a strict obstruction:

```text
FAILED_ROUTE_FINITE_ANCHOR_DARK_MATTER_DERIVATION
FAILED_ROUTE_ALP_SHIFT_ANOMALY_SCALE_OBSTRUCTION
FAILED_ROUTE_CONTACT_DARK_SECTOR_STABILITY_OBSTRUCTION
HEAVY_SECTOR_DM_ABSENCE_REMAINS_BINDING
```

The finite anchors remain valuable inventory, but they are not yet ALPs, not yet a stable dark sector, and not yet a relic-density theorem.

## Gate-224 inheritance

| Inherited object | Status |
|---|---|
| `RelicDecaySeal` | active |
| `FlavorAlignmentSeal` | active |
| PeV carriers decay before BBN | yes, conditionally |
| Heavy-sector dark matter | absent |
| `Ω_heavy h²` | `0` |

This inheritance is essential: Gate 225 is not trying to make the PeV unification carriers into dark matter. Gate 224 already closed that route.

## Finite anchor inventory

| Anchor | Value / count | Current semantic status |
|---|---:|---|
| B-sector first spectral gap | `0.1024649212` | dimensionless scalar spectral anchor |
| Loop-scaled B-gap diagnostic | `0.000648866694` | diagnostic only, not a mass or coupling |
| Contact partial-overlap modes | `7` | positive finite overlap anchors, class-open |

## ALP route audit

A valid ALP or QCD-axion route requires all of the following:

| Required ALP structure | Derived? | Comment |
|---|---:|---|
| Continuous shift symmetry `a → a + c` | no | no global `U(1)_PQ`-like theorem |
| Compact periodic coordinate | no | no finite periodic axion coordinate |
| Axion decay constant `f_a` | no | no dimensionful finite scale |
| Instanton potential / mass law | no | no axion mass theorem |
| Pontryagin coupling `a F∧F` | no | no gauge-anomaly projection row |
| QCD theta relaxation | no | no QCD-axion mechanism |

Verdict:

```text
B-gap is a dimensionless spectral scalar only; no continuous shift symmetry,
periodic axion coordinate, f_a, instanton potential, or F∧F anomaly coupling
is derived.
```

## Contact dark-sector audit

A stable contact dark sector would require:

| Required dark-sector structure | Derived? |
|---|---:|
| Gauge-singlet theorem | no |
| Stability symmetry | no |
| Local dark-field action | no |
| Mass scale | no |
| Self-interaction law | no |
| Thermal history / production law | no |
| Relic abundance calculation | no |

The seven contact modes are sequestered only in the weak sense that previous gates refused to promote them to SM carriers. That is not a proof of a stable gauge-singlet dark sector.

## Misalignment preflight

The usual ALP misalignment route schematically requires:

```text
Ω_a h² ∝ θ_i² f_a² m_a^(1/2)
```

Gate 225 cannot evaluate this because the engine lacks:

```text
m_a
f_a
θ_i
cosmological history
```

The B-sector gap is dimensionless and cannot dimensionalize itself.

## Relic accounting

| Sector | Relic result |
|---|---|
| PeV heavy threshold sector | `Ω_heavy h² = 0` |
| B-gap ALP candidate | not computed |
| Contact-mode dark sector | not computed |
| Total model dark matter | still open |

Dark matter is deferred to a future finite or sealed route:

```text
derive a finite shift symmetry and anomaly map for the B-sector gap
derive a stable gauge-singlet dark action for the seven contact partial-overlap modes
derive a dimensionful scale f_dark or f_a without using observed Ω_DM as input
derive a production mechanism and cosmological history before computing Ω h²
```

## Firewall ledger

Gate 225 does **not** claim:

```text
finite-derived dark matter
B-gap physical mass
B-gap axion decay constant
invented shift symmetry
invented a F∧F coupling
contact-mode particle promotion
contact-mode singlet theorem
misalignment relic abundance
observed Ω_DM-derived finite scale
```

## Registry theorem checks

The theorem verifies:

1. Gate 224 heavy-sector dark-matter absence is inherited.
2. Finite anchor inventory is present.
3. ALP route is obstructed without shift symmetry, `f_a`, and `F∧F` map.
4. Contact dark-sector route remains compatible future inventory only.
5. Misalignment relic density cannot be computed.
6. Dark matter remains open outside the heavy threshold sector.
7. Firewalls remain closed.

## Final theorem statement

Gate 225 preserves the Gate-224 heavy-sector dark-matter absence theorem and audits finite anchors for a replacement dark sector. The B-sector gap and seven contact partial-overlap modes are real finite data, but no shift symmetry, Pontryagin/anomaly map, axion decay constant, stable contact singlet action, or relic-production law is derived. Therefore finite-anchor dark matter remains an open route, not a theorem.
