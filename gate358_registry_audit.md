# Gate 358 Registry Audit — Exponential tau_eta Texture / B-Gap Mixing Hierarchy Audit

## Gate identity

- **Gate:** 358
- **Package:** `pkg/bridge/exponentialtauetatexture`
- **Theorem:** `ExponentialTauEtaTextureBGapMixingHierarchyAuditTheorem`
- **Audit ID:** `GATE358-EXPONENTIAL-TAU-ETA-TEXTURE-BGAP-MIXING-HIERARCHY-AUDIT`
- **Layer:** Bridge / Phase-III Flavor Geometry
- **Purpose:** audit whether the nonlinear exponential texture
  `Y = y0 exp(B_gap C) diag(2,-2,1)` can turn the signed triality seed into a kinetic-safe fermion hierarchy and CKM/PMNS-like mixing structure.

---

## 1. Exponential map formalization

Gate 358 tests the nonlinear escape route left by Gates 356–357.

The audited texture is:

```text
Y = y0 · exp(B_gap C) · diag(2,-2,1)
```

with:

```text
B_gap = 0.102464921191
τ_eta = (2,-2,1)
```

The representative Hermitian triality-mixing generators tested were:

```text
C12 = E12 + E21
C13 = E13 + E31
C23 = E23 + E32
```

and one non-native amplified witness:

```text
C12(c=5) = 5(E12 + E21)
```

The exponential map is rank-preserving and positive-metric-safe in the sense that `exp(B C)` is invertible for finite Hermitian/symmetric `C`.

**Status:** `CONDITIONAL_SUPPORT_EXPONENTIAL_MAP_FORMALIZED`

---

## 2. Exponential sign-interference mechanism

For the 1–2 block:

```text
D = diag(2,-2)
C = c σ_x
```

The texture gives:

```text
Y†Y = D · exp(2 B_gap C) · D
```

and its two singular values split as:

```text
σ_+ = 2 exp(+B_gap c)
σ_- = 2 exp(-B_gap c)
```

Therefore:

```text
σ_+ / σ_- = exp(2 B_gap c)
```

This proves the important mathematical point: **the exponential map does activate the sign interference of `+2` and `-2` without destroying rank.**

**Status:** `CONDITIONAL_SUPPORT_SIGN_INTERFERENCE_VERIFIED_IN_EXPONENTIAL_TEXTURE`

---

## 3. Numerical hierarchy sieve

### Canonical C12 generator

```text
c = 1
singular values = (2.215796874941, 1.805219623350, 1.000000000000)
first/second ratio = 1.227438947750
high/low ratio     = 2.215796874941
```

The canonical generator breaks the degeneracy, but only mildly.

### Amplified witness c=5

```text
c = 5
singular values = (3.338333687415, 1.198202568868, 1.000000000000)
first/second ratio = 2.786117952132
high/low ratio     = 3.338333687415
```

Even the user-suggested `c≈5` witness is still far below the observed charged-fermion hierarchy scale.

### C13 and C23 canonical generators

```text
singular values = (2.034434499502, 2.000000000000, 0.983074166551)
high/low ratio ≈ 2.069462
```

Again, the texture is mathematically legal, but the hierarchy is mild.

---

## 4. Required generator magnitude

For the clean 1–2 exponential block:

```text
ratio = exp(2 B_gap c)
```

The required `c` values are:

```text
ratio 17   -> c = 13.825284
ratio 44   -> c = 18.465781
ratio 136  -> c = 23.972374
ratio 207  -> c = 26.022168
```

So the mechanism is real, but observed charged-fermion mass ratios require a large generator norm / repeated-flow amplifier not currently derived from the finite core.

**Status:** `CONDITIONAL_TENSION_OBSERVED_HIERARCHY_REQUIRES_LARGE_GENERATOR_NORM`

---

## 5. CKM / PMNS shadow evaluation

Gate 358 also audits sector misalignment by comparing canonical generator eigenbases, for example:

```text
up-sector witness:   C12
 down-sector witness: C23
```

This produces nontrivial basis misalignment, but the resulting angles are large / democratic in character and no native Morita-charge rule selects which triality generator belongs to which fermion sector.

The gate therefore does **not** derive CKM or PMNS textures.

**Status:** `CONDITIONAL_SUPPORT_CKM_EIGENVECTOR_SHADOW_AUDITED`
**Status:** `FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED`

---

## 6. Parameter census

```text
starting minimal vacuum coordinates = 15
texture reduction                   = 0
CKM/PMNS reduction                  = 0
total reduction                     = 0
remaining vacuum coordinates        = 15
seven-seal target reached           = false
```

The exponential texture is a valid nonlinear mechanism, but without a native generator-amplification theorem or sector-assignment theorem, it does not reduce the empirical vacuum-coordinate ledger.

---

## Final status ledger

```text
CONDITIONAL_SUPPORT_EXPONENTIAL_MAP_FORMALIZED
CONDITIONAL_SUPPORT_CANONICAL_TRIALITY_GENERATOR_AUDITED
CONDITIONAL_SUPPORT_EXPONENTIAL_HIERARCHY_SIEVE_EXECUTED
CONDITIONAL_SUPPORT_SIGN_INTERFERENCE_VERIFIED_IN_EXPONENTIAL_TEXTURE
CONDITIONAL_SUPPORT_CKM_EIGENVECTOR_SHADOW_AUDITED
CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED

CONDITIONAL_TENSION_CANONICAL_GENERATOR_SPLITS_ONLY_MILDLY
CONDITIONAL_TENSION_OBSERVED_HIERARCHY_REQUIRES_LARGE_GENERATOR_NORM
CONDITIONAL_TENSION_SECTOR_TRIALITY_OPERATOR_ASSIGNMENT_NOT_CANONICAL
CONDITIONAL_TENSION_EXPONENTIAL_TEXTURE_DOES_NOT_DERIVE_CKM_PMNS

FAILED_ROUTE_EXPONENTIAL_TEXTURE_NOT_DERIVED_AS_OBSERVED_HIERARCHY
FAILED_ROUTE_REQUIRED_TRIALITY_GENERATOR_MAGNITUDE_NOT_DERIVED
FAILED_ROUTE_HIERARCHY_DEGENERACY_NOT_BROKEN_TO_OBSERVED_SCALE
FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED
FAILED_ROUTE_NO_ADDITIONAL_PARAMETER_REDUCTION_PROVED
FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED
```

---

## Verdict

Gate 358 confirms the central insight: the signed triality seed must be passed through a nonlinear exponential map for its signs to become dynamically visible. The mechanism is mathematically sound, rank-safe, and capable in principle of exponential splitting.

However, with canonical normalized triality generators and the derived `B_gap`, the splitting is far too mild. The observed fermion hierarchy would require a large generator coefficient, repeated exponential flow, or a new native amplification theorem. CKM/PMNS mixing also remains unselected.

Therefore Gate 358 is a structural advance, but not a parameter-reduction theorem.
