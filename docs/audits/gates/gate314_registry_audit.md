# Gate 314 Registry Audit — Intermediate Threshold Decoupling / Quartic Jump Transport Audit

## Gate identity

- **Gate:** 314
- **Package:** `pkg/bridge/intermediatethresholdjump`
- **Theorem:** `IntermediateThresholdDecouplingQuarticJumpTransportAuditTheorem`
- **Audit ID:** `GATE314-INTERMEDIATE-THRESHOLD-DECOUPLING-QUARTIC-JUMP-TRANSPORT-AUDIT`
- **Layer:** Bridge / Threshold Matching / Higgs IR Transport
- **Purpose:** formalize a two-stage RG transport interrupted by a finite quartic threshold jump and extract the required `Δλ` target that moves the Gate-313 continuous-flow lower envelope to the 125.10 GeV comparison target.

---

## Inherited inputs

Gate 314 inherits the structural and numerical diagnostic results from Gates 308–313:

```text
λ_H(Λ_GUT) = 1197 / 4624 = 0.258866782007
Λ_GUT      = 2.40099519719e15 GeV   conditional closed-triangle lane
M_thr      = 1.46774973718e6 GeV    inherited PeV threshold lane
v          = 246.22 GeV
m_ref      = 125.10 GeV
λ_ref(v)   = (m_ref / v)^2 / 2 = 0.129073762456
```

The gate does **not** claim `125.10 GeV` as a derived ASHA output. It is used only as a diagnostic comparison target to quantify the threshold jump obligation.

---

## Two-stage RG formalization

The transport is explicitly split into two segments:

```text
Segment A: Λ_GUT -> M_threshold
Segment B: M_threshold -> v
```

High-scale PeV/vectorlike beta ledger:

```text
b1 = 41/10 + 7.78628724237
b2 = -19/6 + 9.65295390904
b3 = -7    + 8.98628724237
```

Low-scale SM beta ledger:

```text
b1 = 41/10
b2 = -19/6
b3 = -7
```

The quartic beta function remains the same one-loop Gate-309/Gate-313 diagnostic equation. No two-loop terms, pole-mass corrections, or finite self-energy shifts are inserted.

**Status:** `CONDITIONAL_SUPPORT_TWO_STAGE_RG_TRANSPORT_FORMALIZED`

---

## Threshold matching rule

At the inherited PeV threshold, Gate 314 inserts the formal matching rule:

```text
λ(M_threshold^-) = λ(M_threshold^+) + Δλ
```

Sign convention:

```text
Δλ < 0  => lowers the post-decoupling effective Higgs quartic
```

This is the correct sign for a tree-level heavy-scalar portal threshold of the form:

```text
Δλ_theory = - λ_mix^2 / (4 λ_heavy)
```

Gate 314 does not derive `λ_mix`, `λ_heavy`, or the threshold mass from the finite graph. It extracts the target value that Phase II must derive.

**Status:** `CONDITIONAL_SUPPORT_THRESHOLD_JUMP_INSERTION_FORMALIZED`

---

## Diagnostic lanes

| Lane | Top fraction | `y_t(Λ)` | Role | Gate 314 verdict |
| --- | ---: | ---: | --- | --- |
| `legacy_all_rplus_assigned_to_top` | `1` | `1.282758926` | inherited high-tension top-attractor lane | no moderate jump bracketed |
| `tau_eta_unique_low_top_fraction` | `1/9` | `0.427586309` | most aggressive Gate-313 τη witness | no moderate jump bracketed |
| `gauge_only_zero_top_lower_envelope` | `0` | `0` | preferred continuous-flow floor lane | solved |

The first two lanes remain diagnostic. Gate 314 does not force them into an artificial threshold solution. The preferred extracted obligation is the gauge-only lower-envelope jump because Gate 313 proved this is the hard lower floor of continuous one-loop transport.

---

## Required jump extraction — preferred lane

Preferred lane:

```text
gauge_only_zero_top_lower_envelope
```

Baseline without jump:

```text
λ(M_threshold^+) = 0.226170677080
λ(v)             = 0.206657145096
m_H              = 158.293666 GeV
```

Required PeV-scale jump to land at the 125.10 GeV comparison target:

```text
Δλ_required      = -0.097561578813
λ(M_threshold^-) = 0.128609098267
λ(v)             = 0.129073762455
m_H              = 125.100000 GeV
```

This jump has the required negative sign and moderate quartic magnitude.

**Status:** `CONDITIONAL_SUPPORT_INTERMEDIATE_THRESHOLD_JUMP_EXTRACTED`

---

## Heavy-portal viability sieve

For a tree-level scalar-portal threshold:

```text
Δλ = -λ_mix² / (4 λ_heavy)
```

The preferred extracted target implies:

```text
λ_mix² / λ_heavy = -4 Δλ = 0.390246315254
```

If `λ_heavy = 1` as a normalization witness, then:

```text
λ_mix ≈ 0.624696978746
```

This is a plausible perturbative-size target, but it is **not derived**. It is the quantitative obligation for the next heavy-sector theorem.

**Status:** `CONDITIONAL_SUPPORT_PORTAL_MAGNITUDE_TARGET_FORMALIZED`

---

## Final status ledger

```text
CONDITIONAL_SUPPORT_TWO_STAGE_RG_TRANSPORT_FORMALIZED
CONDITIONAL_SUPPORT_THRESHOLD_JUMP_INSERTION_FORMALIZED
CONDITIONAL_SUPPORT_INTERMEDIATE_THRESHOLD_JUMP_EXTRACTED
CONDITIONAL_SUPPORT_THRESHOLD_JUMP_VIABILITY_SIEVE_FORMALIZED
CONDITIONAL_SUPPORT_PORTAL_MAGNITUDE_TARGET_FORMALIZED
CONDITIONAL_TENSION_DISCONTINUOUS_THRESHOLD_MECHANISM_REQUIRED
CONDITIONAL_SUPPORT_GATE314_FIREWALLS_PRESERVED

FAILED_ROUTE_THRESHOLD_JUMP_VALUE_NOT_DERIVED_FROM_FINITE_GEOMETRY
FAILED_ROUTE_HEAVY_PORTAL_COUPLING_NOT_DERIVED
FAILED_ROUTE_HEAVY_SELF_QUARTIC_NOT_DERIVED
FAILED_ROUTE_THRESHOLD_SCALE_STILL_CONDITIONAL
FAILED_ROUTE_TOP_YUKAWA_FRACTION_STILL_CONDITIONAL
FAILED_ROUTE_TWO_LOOP_RGE_NOT_EXECUTED
FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_COMPUTED
FAILED_ROUTE_FINAL_HIGGS_MASS_NOT_CLAIMED
FAILED_ROUTE_HEAVY_INTERMEDIATE_SECTOR_NOT_CONSTRUCTED
```

---

## Verification

Only the related package test was run:

```text
go test ./pkg/bridge/intermediatethresholdjump
ok  	github.com/bagherbal/asha-engine/pkg/bridge/intermediatethresholdjump	2.763s
```

No full-suite, `go test ./...`, `go test ./cmd/asha`, `go test ./internal/app`, or broad package sweep was run.

---

## Verdict

Gate 314 proves that the Gate-313 continuous-flow floor is not the end of the story. A finite negative PeV-scale threshold jump can mathematically move the one-loop gauge-only lower envelope from approximately `158.294 GeV` to the `125.10 GeV` comparison target.

The extracted requirement is:

```text
Δλ_required ≈ -0.097561578813
```

This is exactly the sign expected from integrating out a heavy scalar portal, and its magnitude is perturbatively plausible. However, Gate 314 does not derive the heavy portal, the heavy self-quartic, or the threshold scale from the finite ASHA geometry. It therefore records the threshold jump as a precise Phase-II obligation, not as a final Higgs-mass prediction.
