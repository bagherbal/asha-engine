# Gate 338 Registry Audit — On-Shell Renormalization Scheme / Passarino-Veltman Pole Matching Audit

## Gate identity

- **Gate:** 338
- **Package:** `pkg/bridge/higgsonshellrenormalizationscheme`
- **Theorem:** `OnShellRenormalizationSchemePassarinoVeltmanPoleMatchingAuditTheorem`
- **Audit ID:** `GATE338-ON-SHELL-RENORMALIZATION-SCHEME-PASSARINO-VELTMAN-POLE-MATCHING-AUDIT`
- **Layer:** Bridge / Higgs Precision QFT Renormalization
- **Purpose:** formalize the QFT renormalization machinery required to convert the native ASHA Higgs tree/running proxy into a collider pole-mass comparison without modifying the immutable contact scalar shape.

---

## Inherited precision target

Gate 338 inherits the Gate 337 precision route result:

```text
m_native = 125.274157149698972 GeV
M_ref    = 125.10 GeV
```

The pole equation convention is retained:

```text
M_H² - m_run² + ReΠ_HH(M_H²) = 0
```

Therefore the required finite self-energy target is:

```text
ReΠ_required = m_native² - M_ref²
             = 43.604449567481 GeV²
```

Gate 337 also showed that the raw polynomial one-loop capacity kernel is not the pole correction:

```text
Π_poly_raw = -991.567029891610 GeV²
```

so the renormalized finite residue/counterterm ledger must absorb:

```text
Π_finite_remainder = ReΠ_required - Π_poly_raw
                   = 1035.171479459092 GeV²
```

**Status:** `CONDITIONAL_SUPPORT_GATE337_PRECISION_ROUTE_SIEVE_INHERITED`

---

## Passarino-Veltman formalization

Gate 338 installs the finite scalar integral basis used by the one-loop Higgs self-energy.

Finite tadpole block:

```text
A0_fin(m²; μ²) = m²[1 - ln(m²/μ²)]
```

Finite equal-mass two-point block on the real below-threshold branch:

```text
B0_fin(s; m², m²; μ²)
  = -ln(m²/μ²)
    + 2
    - 2 sqrt(4m²/s - 1) atan(1/sqrt(4m²/s - 1))
```

Renormalized self-energy schematic:

```text
ReΠ_HH(p²)
  = Π_top + Π_W + Π_Z + Π_H
    + δM_H²
    + (p² - M_H²)δZ_H
    + ...
```

**Status:** `CONDITIONAL_SUPPORT_PASSARINO_VELTMAN_STRUCTURE_FORMALIZED`

---

## Finite PV block witnesses

Using the quarantined comparison scale:

```text
μ = m_native = 125.274157149699 GeV
s = M_ref² = 125.10² GeV²
```

and conventional comparison masses:

```text
m_t = 172.76 GeV
m_W = 80.379 GeV
m_Z = 91.1876 GeV
```

Gate 338 computes:

| block | A0_fin [GeV²] | B0_fin | 4m²/s branch |
| --- | ---: | ---: | ---: |
| top | 10661.071787177943 | -0.550445463802 | 7.628370231 |
| W | 12194.750375579539 | 1.448101252519 | 1.651317447 |
| Z | 13596.741394774119 | 1.031449328240 | 2.125283854 |
| H-native | 15693.614449567480 | 0.185619588061 | 4.011144900 |

All four are below pair threshold for the chosen comparison point, so the real equal-mass branch is sufficient for this audit.

**Status:** `CONDITIONAL_SUPPORT_FINITE_PV_BLOCKS_COMPUTED`

---

## Renormalization scheme sieve

Gate 338 audits three lanes.

| Scheme lane | What it fixes | Gate 338 verdict |
| --- | --- | --- |
| **On-Shell** | Defines physical masses as real propagator poles; finite counterterms enforce the pole equation. | Correct language for collider comparison, but requires full SM coefficient/counterterm table. |
| **MS-bar** | Defines running parameters by UV-pole subtraction and RG scale dependence. | Natural for RG transport, but not directly the collider pole mass. |
| **Native ASHA contact boundary** | Fixes `λ/g² = 1197/4624` and the native tree proxy. | Does not select the IR finite renormalization prescription. |

**Status:** `CONDITIONAL_SUPPORT_RENORMALIZATION_SCHEME_DEPENDENCY_AUDITED`

---

## Counterterm target mapping

The finite target is mapped into the on-shell ledger:

```text
M_H² - m_run² + ReΠ_HH(M_H²) = 0
```

with:

```text
ReΠ_required       = +43.604449567481 GeV²
Π_poly_raw         = -991.567029891610 GeV²
finite remainder   = +1035.171479459092 GeV²
remainder / target = +23.739
remainder / |raw|  = +1.044
```

This is not treated as a derived counterterm. It is the exact mathematical debt owed by the missing renormalized SM pole-matching scheme.

**Status:** `CONDITIONAL_SUPPORT_COUNTERTERM_TARGET_MAPPED_TO_ON_SHELL_LEDGER`

---

## Geometric alignment audit

Gate 338 preserves the ASHA geometry:

```text
λ_H/g_*² = 1197/4624
α_GUT⁻¹ = 8π
m_native = v sqrt(1197/4624)
```

The contact scalar shape is not modified, and the electroweak VEV is not shifted. The audit concludes that these UV spectral-action boundary data do not themselves choose the IR renormalization scheme.

**Status:** `CONDITIONAL_SUPPORT_GEOMETRIC_ALIGNMENT_AUDITED`

---

## Final status ledger

```text
CONDITIONAL_SUPPORT_GATE337_PRECISION_ROUTE_SIEVE_INHERITED
CONDITIONAL_SUPPORT_PASSARINO_VELTMAN_STRUCTURE_FORMALIZED
CONDITIONAL_SUPPORT_FINITE_PV_BLOCKS_COMPUTED
CONDITIONAL_SUPPORT_RENORMALIZATION_SCHEME_DEPENDENCY_AUDITED
CONDITIONAL_SUPPORT_COUNTERTERM_TARGET_MAPPED_TO_ON_SHELL_LEDGER
CONDITIONAL_SUPPORT_GEOMETRIC_ALIGNMENT_AUDITED
CONDITIONAL_SUPPORT_PRECISION_FIREWALLS_PRESERVED
CONDITIONAL_TENSION_IR_RENORMALIZATION_SCHEME_NOT_SELECTED_BY_FINITE_CORE
CONDITIONAL_TENSION_FINITE_COUNTERTERM_REQUIRED_TO_REACH_POLE_TARGET
CONDITIONAL_TENSION_PV_BASIS_WITHOUT_COEFFICIENTS_DOES_NOT_CLOSE_POLE_MASS
FAILED_ROUTE_FULL_SM_ONE_LOOP_COEFFICIENT_TABLE_NOT_INSTALLED
FAILED_ROUTE_RENORMALIZED_COUNTERTERMS_NOT_DERIVED_FROM_NATIVE_SCHEME
FAILED_ROUTE_GAUGE_INPUT_SCHEME_NOT_DERIVED
FAILED_ROUTE_ASHA_BOUNDARY_DOES_NOT_SELECT_IR_RENORMALIZATION_SCHEME
FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED
```

---

## Verdict

Gate 338 successfully formalizes the Passarino-Veltman integral structure and renormalization scheme ledger required for a serious Higgs pole-mass comparison.

It does **not** claim the exact collider Higgs mass. The missing object is now sharply localized: a full Standard Model one-loop coefficient and counterterm table in a chosen input scheme, contracted against the installed finite PV blocks.
