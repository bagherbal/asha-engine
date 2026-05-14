# Gate 336 Registry Audit — Exact Inverse Higgs Shape Deviation / Full-Precision Diagnostic Audit

## Gate identity

- **Gate:** 336
- **Package:** `pkg/bridge/higgsinverseshapeprecision`
- **Theorem:** `ExactInverseHiggsShapeDeviationPrecisionAuditTheorem`
- **Audit ID:** `GATE336-EXACT-INVERSE-HIGGS-SHAPE-DEVIATION-PRECISION-AUDIT`
- **Layer:** Bridge / Phase-II Precision Diagnostics
- **Purpose:** invert the Higgs proxy relation using exact rational arithmetic to compare the native contact scalar shape against the shape implied by the nominal collider mass and electroweak VEV, without fitting or modifying the geometry.

---

## Inherited exact branch

Gate 336 inherits the Gate 335 closed-form branch:

```text
α_GUT⁻¹ = 8π
g_*² = 1/2
λ_H/g_*² = 1197/4624
λ_H = 1197/9248
m_native = v √(1197/4624)
```

Using the quarantined decimal-rational input:

```text
v = 246.22 GeV = 12311/50 GeV
M_H,ref = 125.10 GeV = 1251/10 GeV
```

Gate 335 gave:

```text
m_native = 125.274157149698971935740602811547201489421906436146511739793682642913 GeV
```

**Status:** `CONDITIONAL_SUPPORT_GATE335_EXACT_NATIVE_PRECISION_INHERITED`

---

## Exact inverse observed shape

The observed tree-level proxy shape is not guessed; it is computed exactly:

```text
R_obs := (M_H / v)^2
       = (125.10 / 246.22)^2
       = 39125025 / 151560721
       ≈ 0.25814752491181403
```

The corresponding proxy quartic is:

```text
λ_obs := R_obs / 2
       = 39125025 / 303121442
```

The native contact shape is:

```text
R_native = 1197 / 4624
         ≈ 0.2588667820069204
```

**Status:** `CONDITIONAL_SUPPORT_EXACT_INVERSE_HIGGS_SHAPE_COMPUTED`

---

## Exact deviation ledger

Exact native-minus-observed shape gap:

```text
ΔR = R_native - R_obs
   = 504067437 / 700816773904
   ≈ 0.0007192570951063575
```

Exact quartic gap:

```text
Δλ = λ_native - λ_obs
   = 504067437 / 1401633547808
   ≈ 0.00035962854755317876
```

Relative shape excess:

```text
ΔR / R_obs ≈ 0.278622502908778%
```

Mass gap:

```text
m_native - M_H,ref
= 0.174157149698971935740602811547201489421906436146511739793682642913 GeV

relative mass excess ≈ 0.139214348280553%
```

**Status:** `CONDITIONAL_SUPPORT_EXACT_CONTACT_SHAPE_DEVIATION_COMPUTED`

---

## Pole self-energy equivalence target

The exact mass-squared difference is:

```text
m_native² - M_H,ref²
= 504067437 / 11560000 GeV²
≈ 43.60444956747405 GeV²
```

Using the pole equation convention:

```text
M_H² - m_run² + ReΠ_HH(M_H²) = 0
```

Gate 336 therefore preserves the Gate 335/332 target:

```text
ReΠ_required = 504067437 / 11560000 GeV²
             ≈ 43.60444956747405 GeV²
```

**Status:** `CONDITIONAL_SUPPORT_SELF_ENERGY_EQUIVALENCE_RECOMPUTED`

---

## Counterfactual VEV diagnostic

Holding the native contact shape fixed, the VEV that would make the proxy mass exactly `125.10 GeV` is:

```text
v_required = M_H / √(1197/4624)
           ≈ 245.877702958259464840997909714225257613934962666015612816425045659317 GeV
```

Relative to `246.22 GeV`:

```text
Δv = -0.342297041740535159002090285774742386065037333984387183574954340683 GeV
```

This is only a diagnostic. Gate 336 does **not** modify the electroweak VEV or tune the native contact shape.

**Status:** `CONDITIONAL_SUPPORT_REQUIRED_VEV_FOR_EXACT_NATIVE_SHAPE_COMPUTED`

---

## Precision and efficiency ledger

The Gate 336 computation uses:

```text
Rational core: yes
float64 native branch: no
sqrt precision: 768 bits
```

Exact rational objects:

```text
R_native = 1197/4624
λ_native = 1197/9248
v = 12311/50
M_H = 1251/10
R_obs = 39125025/151560721
ΔR = 504067437/700816773904
Δλ = 504067437/1401633547808
ReΠ = 504067437/11560000
```

**Status:** `CONDITIONAL_SUPPORT_FULL_PRECISION_INVERSE_LEDGER_COMPILED`

---

## Final status ledger

```text
CONDITIONAL_SUPPORT_GATE335_EXACT_NATIVE_PRECISION_INHERITED
CONDITIONAL_SUPPORT_EXACT_INVERSE_HIGGS_SHAPE_COMPUTED
CONDITIONAL_SUPPORT_EXACT_CONTACT_SHAPE_DEVIATION_COMPUTED
CONDITIONAL_SUPPORT_SELF_ENERGY_EQUIVALENCE_RECOMPUTED
CONDITIONAL_SUPPORT_REQUIRED_VEV_FOR_EXACT_NATIVE_SHAPE_COMPUTED
CONDITIONAL_SUPPORT_FULL_PRECISION_INVERSE_LEDGER_COMPILED

CONDITIONAL_TENSION_NATIVE_CONTACT_SHAPE_EXCEEDS_OBSERVED_PROXY_BY_0_2786_PERCENT
CONDITIONAL_TENSION_NATIVE_MASS_PROXY_ABOVE_125_10_BY_0_174_GEV

FAILED_ROUTE_POLE_CORRECTION_NOT_COMPUTED
FAILED_ROUTE_CONTACT_SHAPE_NOT_MODIFIED_TO_FIT_DATA
FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED
```

---

## Verdict

Gate 336 gives the exact inverse diagnostic of the Gate 335 Higgs branch.

The native shape `1197/4624` is not altered. It exceeds the nominal tree-level observed proxy shape by exactly:

```text
504067437 / 700816773904
```

This corresponds to a `+0.174157149699 GeV` native mass excess and an exact pole/self-energy target of:

```text
ReΠ_required = 504067437 / 11560000 GeV²
```

The result is a full-precision diagnostic ledger, not a fitted correction and not a collider pole-mass derivation.
