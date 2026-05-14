# Gate 372 Registry Audit — Native Moduli Space Dimension / Exact Dirac Parameter Census Sieve

## Executive truth statement

Gate 372 changes the question from an external Standard-Model parameter subtraction to a native finite-Dirac moduli-space census.

The result is category-sharper than the earlier phrase “15 vacuum coordinates remain”:

- The **minimal charged finite-Dirac flavor moduli** have dimension **13**.
- The older external **15** ledger decomposes as:

```text
15 = 13 finite-Dirac charged-flavor moduli + theta_QCD + one absolute unit / VEV scale
```

- If the all-allowed right-neutrino / Majorana sector is included, the finite-Dirac moduli dimension is **31**.
- No hidden cross-sector ASHA constraint reduces the flavor moduli below the minimal external 15-vacuum ledger.
- No physical vacuum point, Yukawa spectrum, CKM texture, PMNS texture, or Majorana orientation is selected.

Therefore Gate 372 does **not** magically solve the vacuum problem. It corrects the epistemology: `dim M(D_F)` and the old “15 inputs” are related, but not identical categories.

## Inherited chain

| Gate range | Meaning inherited by Gate 372 |
|---:|---|
| 297–299 | Physical finite Hilbert space, finite spectral triple, order-one edge graph, inner fluctuation field content. |
| 320–347 | B-gap / heavy-light / Majorana and flavor-sector attempts expose support structure but do not derive flavor texture. |
| 345 | External minimal SM-19 census gives 15 remaining vacuum inputs after four ASHA boundary constraints. |
| 362–371 | Modular, KMS, Lorentzian, Bimodule, eta-graded, support-generation, and vibrational routes find capacity witnesses but no lawful vacuum selector. |

## Parameterization of the finite Dirac operator

Gate 372 parameterizes the finite Dirac operator by the allowed first-order edge graph. The J-real mirrored blocks are not counted twice as independent data.

| Block | Edge | Matrix type | Raw real dimension | Role |
|---|---|---:|---:|---|
| `Y_u` | `Q_L ↔ u_R` | generic complex `3×3` | 18 | up-type quark Yukawa matrix |
| `Y_d` | `Q_L ↔ d_R` | generic complex `3×3` | 18 | down-type quark Yukawa matrix |
| `Y_e` | `L_L ↔ e_R` | generic complex `3×3` | 18 | charged-lepton Yukawa matrix |
| `Y_ν` | `L_L ↔ ν_R` | generic complex `3×3` | 18 | Dirac-neutrino Yukawa matrix, extended ledger |
| `M_R` | `ν_R ↔ ν_R^c` | complex symmetric `3×3` | 12 | right-neutrino Majorana block, extended ledger |

Raw dimensions:

| Ledger | Raw real dimension |
|---|---:|
| Minimal charged finite Dirac: `Y_u,Y_d,Y_e` | 54 |
| Dirac-neutrino extension: `Y_u,Y_d,Y_e,Y_ν` | 72 |
| Majorana/seesaw extension: `Y_u,Y_d,Y_e,Y_ν,M_R` | 84 |

## Axiomatic sieve

| Axiom | Effect |
|---|---|
| J-reality | Mirrors particle/antiparticle blocks; does not create independent duplicate variables. |
| Chirality | Allows odd left-right finite Dirac edges; forbids same-chirality charged Dirac edges. |
| First-order condition | Enforces the Standard-Model edge graph; forbids quark-lepton cross Yukawa edges and color-changing finite Dirac edges. |
| Majorana symmetry | Reduces `M_R` from a generic complex `3×3` matrix to a complex symmetric `3×3` matrix. |
| Generation texture constraints | **0 additional constraints found** in the current finite ledger. |

Important result:

```text
The spectral-triple axioms constrain edge shape, not the numerical generation texture.
```

## Quotient audit

Gate 372 separates two quotient notions:

1. **Algebra gauge quotient** from `U(C ⊕ H ⊕ M_3(C))`, reduced to the Standard Model gauge group by unimodularity.
2. **Generation-basis quotient**, i.e. unphysical kinetic flavor-basis rotations preserving the representation.

The first does **not** remove generation moduli. The second is required to recover physical flavor parameters.

| Sector | Raw dim | Basis group dim | Stabilizer dim | Orbit dim | Physical dim | Physical content |
|---|---:|---:|---:|---:|---:|---|
| Quark sector `Y_u,Y_d` | 36 | 27 | 1 | 26 | **10** | 6 quark masses + 3 CKM angles + 1 CKM phase |
| Charged-lepton-only `Y_e` | 18 | 18 | 3 | 15 | **3** | 3 charged-lepton singular values |
| Minimal charged finite Dirac | 54 | 45 | 4 | 41 | **13** | quark 10 + charged lepton 3 |
| Dirac-neutrino lepton sector `Y_e,Y_ν` | 36 | 27 | 1 | 26 | **10** | 3 charged-lepton masses + 3 Dirac-neutrino masses + PMNS 4 |
| Quark + Dirac-neutrino finite Dirac | 72 | 54 | 2 | 52 | **20** | quark 10 + Dirac lepton 10 |
| Majorana/seesaw lepton sector `Y_e,Y_ν,M_R` | 48 | 27 | 0 | 27 | **21** | charged, light/heavy neutrino, PMNS, high-energy seesaw data |
| Quark + Majorana finite Dirac | 84 | 54 | 1 | 53 | **31** | quark 10 + Majorana/seesaw lepton 21 |

## Native dimension result

The direct answer depends on which edge ledger is declared active:

| Interpretation | `dim M(D_F)` |
|---|---:|
| Minimal charged SM finite Dirac | **13** |
| Dirac-neutrino extension | **20** |
| All-allowed Majorana/seesaw finite Dirac | **31** |

Because the prompt explicitly included Majorana sectors, the canonical all-allowed finite-Dirac answer is:

```text
N_physical = dim M(D_F) = 31
```

But the minimal SM-19 vacuum ledger remains:

```text
15 = 13 + theta_QCD + absolute scale
```

So the external 15 count is not falsified; it is refined.

## Status ledger

### Supports

```text
CONDITIONAL_SUPPORT_GENERAL_FINITE_DIRAC_PARAMETERIZATION_EXECUTED
CONDITIONAL_SUPPORT_J_REALITY_SIEVE_EXECUTED
CONDITIONAL_SUPPORT_CHIRALITY_SIEVE_EXECUTED
CONDITIONAL_SUPPORT_FIRST_ORDER_EDGE_GRAPH_SIEVE_EXECUTED
CONDITIONAL_SUPPORT_RAW_AXIOMATIC_DIRAC_DIMENSION_COMPUTED
CONDITIONAL_SUPPORT_ALGEBRA_GAUGE_QUOTIENT_AUDITED
CONDITIONAL_SUPPORT_FLAVOR_BASIS_QUOTIENT_COMPUTED
CONDITIONAL_SUPPORT_NATIVE_MODULI_SPACE_COMPUTED
CONDITIONAL_SUPPORT_MINIMAL_CHARGED_FINITE_DIRAC_MODULI_DIMENSION_13
CONDITIONAL_SUPPORT_MAJORANA_FINITE_DIRAC_MODULI_DIMENSION_31
CONDITIONAL_SUPPORT_EXTERNAL_15_DECOMPOSED_AS_13_PLUS_THETA_PLUS_SCALE
CONDITIONAL_SUPPORT_EXTERNAL_COUNT_VERIFIED_AS_MINIMAL_VACUUM_LEDGER_NOT_DF_MODULI
CONDITIONAL_SUPPORT_NO_HIDDEN_FLAVOR_REDUCTION_FOUND_BY_NATIVE_DIRAC_CENSUS
```

### Tensions

```text
CONDITIONAL_TENSION_FINITE_DIRAC_MODULI_ARE_NOT_IDENTICAL_TO_EXTERNAL_15_LEDGER
CONDITIONAL_TENSION_UA_GAUGE_QUOTIENT_DOES_NOT_REMOVE_GENERATION_BASIS_ROTATIONS
CONDITIONAL_TENSION_MAJORANA_NEUTRINO_CENSUS_IS_EXTENDED_MODEL_DEPENDENT
CONDITIONAL_TENSION_SPECTRAL_TRIPLE_AXIOMS_ALLOW_GENERIC_GENERATION_MATRICES
```

### Failed routes preserved

```text
FAILED_ROUTE_NATIVE_MODULI_SPACE_REDUCTION_BELOW_EXTERNAL_15_NOT_FOUND
FAILED_ROUTE_HIDDEN_CROSS_SECTOR_FLAVOR_CONSTRAINTS_NOT_FOUND
FAILED_ROUTE_PHYSICAL_VACUUM_POINT_STILL_NOT_SELECTED
FAILED_ROUTE_YUKAWA_COORDINATES_STILL_FREE_AFTER_NATIVE_CENSUS
FAILED_ROUTE_CKM_TEXTURE_STILL_FREE_AFTER_NATIVE_CENSUS
FAILED_ROUTE_PMNS_MAJORANA_TEXTURE_REMAINS_EXTENDED_MODEL_COORDINATE
```

## Firewalls

| Firewall | Status |
|---|---|
| No observed Yukawa values imported | preserved |
| No CKM values imported | preserved |
| No PMNS values imported | preserved |
| No observed mass values imported | preserved |
| No vacuum direction forced | preserved |
| Algebra gauge quotient not conflated with flavor-basis quotient | preserved |
| Minimal charged SM ledger not conflated with Majorana extension | preserved |
| ASHA landscape ratios preserved | preserved |

## Final truth

Gate 372 proves that the native finite-Dirac census does not secretly reduce the unresolved minimal vacuum problem below the old 15-count. Instead, it reveals that the old 15 was not literally `dim M(D_F)`:

```text
minimal charged dim M(D_F) = 13
external minimal vacuum ledger = 13 + theta_QCD + absolute scale = 15
all-allowed Majorana dim M(D_F) = 31
```

This is a strong negative result. The 15-count is not an accidental external bookkeeping artifact, but it must be stated category-correctly. The finite spectral-triple axioms determine the allowed edge architecture; they do not determine the generation texture.
