# Gate 228 Registry Audit — Pati-Salam falsification / B-sector non-perturbative hierarchy origin search

## Executive verdict

Gate 228 tests the Gate-227 intermediate scale in the most lethal order:

1. Temporarily unseal the dormant `u(4)` leptoquark route **only** as a proton-decay lifetime stress estimate.
2. If that route fails, test whether the B-sector first spectral gap can generate the intermediate scale through a non-perturbative hierarchy function.

Result:

```text
FAILED_ROUTE_PATI_SALAM_INTERMEDIATE_BREAKING
CONDITIONAL_SUPPORT_BSECTOR_NONPERTURBATIVE_HIERARCHY_SHAPE
INTERMEDIATE_BREAKING_SEAL_REQUIRED_NOT_GRANTED
CONDITIONAL_SUPPORT_HIDDEN_SECTOR_ORIGIN_AFTER_PATI_SALAM_FALSIFICATION
```

This is not a finite derivation of the intermediate scale. It is a falsification of the baryon-unsafe Pati-Salam route plus a controlled diagnostic showing that the B-gap exponential hierarchy has the correct shape, but still lacks a derived coefficient and order parameter.

---

## Inherited sealed hierarchy

Gate 228 inherits Gate 227:

```text
M_B   = 2.56895727e6 GeV
M_*   = 1.72179441e17 GeV
M_int = sqrt(M_B M_*) = 6.650726476871e11 GeV
f_a   = 1.00000000e12 GeV
Λ_EFT ≲ 4.99261316e11 GeV
```

The dormant `u(4)` / Pati-Salam slots remain sealed by the `LeptoquarkDynamicsSeal` except for a deliberately isolated proton-lifetime stress calculation.

---

## 1. Pati-Salam proton-decay falsification

Gate 228 temporarily assigns the dormant leptoquark slots a mass

```text
M_LQ = M_int = 6.650726476871e11 GeV
```

and uses the dimension-six stress proxy

```text
Γ_p ~ α² m_p⁵ / M_LQ⁴
τ_p = ℏ / Γ_p
```

with

```text
α = 1/(4π)
m_p = 0.9382720813 GeV
ℏ = 6.582119569e-25 GeV·s
```

Result:

```text
Γ_p ≈ 2.353682580713e-50 GeV
τ_p ≈ 2.796519642e25 seconds
τ_p ≈ 8.861636000285e17 years
```

Compared with the stress bound

```text
τ_p > 1.0e34 years
```

this is short by

```text
16.052486093 orders of magnitude
```

Therefore an intermediate Pati-Salam / active leptoquark breaking at `M_int` is catastrophically excluded in this sealed phenomenological tower.

This is a filter calculation only. Gate 228 does not claim a finite-derived proton lifetime.

---

## 2. B-sector hierarchy search

Gate 228 tests the B-sector first spectral gap:

```text
B_gap = 0.1024649212
```

against the non-perturbative hierarchy family

```text
M_hidden = M_* exp(-c / B_gap)
```

### Canonical `c = 1`

```text
M_* exp(-1/B_gap) = 9.942862120e12 GeV
ratio to M_int    = 14.95003915
log10 gap         = 1.174642330 decades
```

This is outside the one-decade criterion and is not promoted.

### Required coefficient

Solving for the exact coefficient required by the sealed hierarchy gives

```text
c_req = B_gap ln(M_*/M_int) = 1.277138298532
```

and exactly reconstructs

```text
M_* exp(-c_req/B_gap) = M_int
```

This coefficient is order-one, but it is derived from the sealed target scales, not from the finite algebra. It is therefore a target coefficient, not a theorem.

### Diagnostic near-misses

Gate 228 records the strongest simple diagnostic:

```text
c = 4/π = 1.273239544735
M_* exp(-(4/π)/B_gap) = 6.908660279e11 GeV
ratio to M_int = 1.038782801
log10 gap = 0.016524751 decades
```

This is a very strong structural near-resonance, but it is not promoted because no prior ASHA theorem derives `4/π` as the B-sector instanton coefficient or hidden-sector action normalization.

Another diagnostic:

```text
c = 5/4
M_* exp(-(5/4)/B_gap) = 8.667501947e11 GeV
ratio to M_int = 1.303241379
log10 gap = 0.115024861 decades
```

also remains diagnostic only.

---

## 3. IntermediateBreakingSeal status

Gate 228 defines the required future seal:

```text
IntermediateBreakingSeal
```

but does **not** grant it.

Reason:

```text
Pati-Salam route: falsified by proton decay
B-sector exponential route: structurally plausible
B-sector coefficient c: not finite-derived
hidden order parameter: not finite-derived
breaking potential: not finite-derived
EFT mediator origin: not finite-derived
axion shift-breaking mechanism: not finite-derived
```

The correct status is:

```text
SEAL_PREPARED_NOT_GRANTED
```

---

## 4. Baryon-safety theorem

Gate 228 records the following conditional support:

```text
CONDITIONAL_SUPPORT_HIDDEN_SECTOR_ORIGIN_AFTER_PATI_SALAM_FALSIFICATION
```

The intermediate scale must be baryon-safe. A Pati-Salam/leptoquark realization at `M_int` is not baryon-safe. A hidden B-sector origin is compatible with baryon safety because the B-gap currently carries no baryon-number semantics and does not reopen the sealed leptoquark dynamics.

This is not a proof that the B-sector generates the scale. It is a proof that, after Pati-Salam falsification, any viable intermediate origin must be hidden or otherwise baryon-safe.

---

## Firewall ledger

Gate 228 does not claim or derive:

```text
finite intermediate order parameter
finite B-sector breaking potential
finite axion shift-breaking mechanism
finite EFT mediator mass
finite proton lifetime
genuine Pati-Salam dynamics
active u(4) leptoquark curvature
physical B-gap field
canonical c coefficient in exp(-c/B_gap)
```

The proton-decay calculation is a phenomenological kill-switch. The B-gap exponential is a structural search target. The finite core remains unpolluted.

---

## Next gate target

Gate 229 should audit whether the near-coefficient

```text
c_req ≈ 1.277138298532
```

or its close diagnostic

```text
4/π ≈ 1.273239544735
```

can be derived from a finite B-sector action, instanton normalization, contact-volume ratio, or exact trace functional without fitting.
