# Gate 380 Registry Audit — Self-Consistent CCM + Pfaffian Coefficient Closure & f0 Sieve

## Gate identity

- **Gate:** 380
- **Package:** `pkg/bridge/ccmpfaffianf0closure`
- **Theorem:** `SelfConsistentCCMPfaffianCoefficientClosureF0SieveTheorem`
- **Audit ID:** `GATE380-SELF-CONSISTENT-CCM-PFAFFIAN-F0-CLOSURE`
- **Layer:** Bridge / product spectral-action coefficient closure

## Purpose

Gate 379 installed the direct Chamseddine--Connes--Marcolli coefficient ledger for the almost-commutative product `M × F`. Gate 380 asks whether the CCM Higgs-sector normalization, combined with the ASHA Pfaffian hierarchy VEV, closes the Higgs mass through a native value of the zeroth spectral moment `f0`.

The target is not to fit the Higgs mass silently. The target is to extract the required `f0_eff`, compare it with native ASHA finite data, and determine whether the native geometry already contains a lawful origin for the required value.

## Inherited facts

| Source | Inherited fact | Status |
|---|---|---|
| Gate 379 | CCM Einstein-Hilbert coefficient supersedes generic Gate-378 arithmetic | installed |
| Gate 379 | `e/a² = 1197/4624` is a finite trace ratio, not automatically the canonical quartic | installed |
| Gate 379 | Canonical CCM Higgs quartic requires `f0` and field normalization | installed |
| Pfaffian hierarchy ledger | `v/M_P = 2^(3/2) exp(-4π²)` | inherited |
| Finite Dirac edge ledger | four charged/Dirac Yukawa edges plus neutral Majorana edge | inherited |
| Real structure ledger | `J` doubles edge slots by particle/antiparticle conjugation | inherited |

## CCM + Pfaffian formulas

Gate 380 uses the CCM canonical quartic read-off:

```text
λ_H(f0) = π²(e/a²)/(2f0)
```

with

```text
e/a² = 1197/4624 = 0.2588667820069204
```

and the Higgs relation

```text
m_H = v sqrt(2λ_H)
```

Therefore the effective zeroth moment required by a Higgs boundary is

```text
f0_eff = π²(e/a²)(v/m_H)²
```

## Numerical extraction

Using `m_H = 125.10 GeV`:

| VEV convention | VEV | Extracted `f0_eff` | Distance from 10 |
|---|---:|---:|---:|
| Standard EW VEV | `246.22 GeV` | `9.897103339899` | `0.102896660101` |
| Unreduced-Planck Pfaffian VEV | `247.151135557136 GeV` | `9.972101066697` | `0.027898933303` |

The Pfaffian VEV convention is closest to the integer 10.

## Candidate f0 predictions

### Standard VEV, `v = 246.22 GeV`

| Candidate | Interpretation | `λ_H` | Predicted `m_H` | Error vs 125.10 GeV |
|---:|---|---:|---:|---:|
| 7 | contact `ζ(0)` candidate | `0.182493766499` | `148.751838519467 GeV` | `+18.9063%` |
| 10 | J-doubled finite-Dirac edge candidate | `0.127745636550` | `124.454717162690 GeV` | `-0.5158%` |
| 14 | J-doubled contact candidate | `0.091246883250` | `105.183433731081 GeV` | `-15.9205%` |

### Pfaffian VEV, `v = 247.151135557136 GeV`

| Candidate | Interpretation | `λ_H` | Predicted `m_H` | Error vs 125.10 GeV |
|---:|---|---:|---:|---:|
| 7 | contact `ζ(0)` candidate | `0.182493766499` | `149.314376599374 GeV` | `+19.3560%` |
| 10 | J-doubled finite-Dirac edge candidate | `0.127745636550` | `124.925370287551 GeV` | `-0.1396%` |
| 14 | J-doubled contact candidate | `0.091246883250` | `105.581208222059 GeV` | `-15.6026%` |

## Rule-of-10 edge sieve

The finite Dirac edge graph contains five structural edge classes:

```text
1. Q_L ↔ u_R
2. Q_L ↔ d_R
3. L_L ↔ e_R
4. L_L ↔ ν_R
5. ν_R ↔ ν_R^c
```

The real structure `J` contributes the conjugate partner of each edge, producing:

```text
5 × 2 = 10
```

This is a native integer-10 capacity witness.

## Firewall

The edge count is not yet a spectral-action moment theorem.

In the spectral action, `f0` is the zeroth moment/value of the test function. The J-doubled edge count is a finite graph dimension. These are different mathematical objects unless a theorem identifies the test-function zero-moment with the finite-Dirac edge projection.

Therefore Gate 380 does **not** log final Higgs-mass derivation. It logs conditional near-closure.

## Status ledger

### Conditional supports

```text
CONDITIONAL_SUPPORT_CCM_PFAFFIAN_FRAMEWORKS_COMBINED
CONDITIONAL_SUPPORT_CCM_CANONICAL_HIGGS_QUARTIC_FORMULA_INSTALLED
CONDITIONAL_SUPPORT_PFAFFIAN_VEV_HIERARCHY_INSTALLED
CONDITIONAL_SUPPORT_EFFECTIVE_F0_TARGET_EXTRACTED
CONDITIONAL_SUPPORT_F0_TEN_PREDICTS_NEAR_HIGGS_MASS
CONDITIONAL_SUPPORT_RULE_OF_TEN_EDGE_SIEVE_EXECUTED
CONDITIONAL_SUPPORT_J_DOUBLED_FINITE_DIRAC_EDGE_COUNT_EQUALS_TEN
CONDITIONAL_SUPPORT_HIGGS_MASS_CONDITIONAL_NEAR_CLOSURE
CONDITIONAL_SUPPORT_COEFFICIENT_SENSITIVITY_AUDITED
```

### Tensions

```text
CONDITIONAL_TENSION_F0_EFF_NOT_EXACTLY_TEN_WITH_STANDARD_EW_VEV
CONDITIONAL_TENSION_F0_EFF_CLOSE_TO_TEN_WITH_UNREDUCED_PLANCK_PFAFFIAN_VEV
CONDITIONAL_TENSION_SPECTRAL_ACTION_F0_MOMENT_NOT_AUTOMATICALLY_EDGE_COUNT
CONDITIONAL_TENSION_CONTACT_F0_SEVEN_OVERPREDICTS_HIGGS_MASS
CONDITIONAL_TENSION_J_DOUBLED_CONTACT_F0_FOURTEEN_UNDERPREDICTS_HIGGS_MASS
```

### Failed routes / open seals

```text
FAILED_ROUTE_F0_MOMENT_NOT_DERIVED_FROM_EDGE_COUNT
FAILED_ROUTE_HIGGS_MASS_NOT_NATIVELY_CLOSED
FAILED_ROUTE_F0_TEN_ORIGIN_STILL_CONDITIONAL_NOT_A_MOMENT_THEOREM
FAILED_ROUTE_FULL_NUMERICAL_TOE_CLOSURE_STILL_NOT_REACHED
```

## Final truth statement

Gate 380 is a major sharpening gate. It shows that CCM + Pfaffian data reduce the Higgs-sector closure problem to a single effective `f0`, and that this value is very close to the native integer 10 supplied by the J-doubled finite-Dirac edge count. With `f0=10`, the Higgs mass lands near 125 GeV.

However, ASHA has not yet proven that the CCM spectral-action `f0` moment equals the J-doubled finite-Dirac edge count. Therefore the Higgs mass is not natively closed yet. The next required theorem is an `f0` moment-origin theorem, not another broad phenomenological scan.
