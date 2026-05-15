# Gate 487 Registry Audit — CKM Rephasing-Invariant Polynomial Constraint Search

## Verdict

```text
CONDITIONAL_SUPPORT_GATE486_CKM_FIREWALL_INHERITED
CONDITIONAL_SUPPORT_NULL_C3_SPECTRUM_ANSATZ_CONSTRUCTED
CONDITIONAL_SUPPORT_COMMUTATOR_SIEVE_EXECUTED_ON_SYNTHETIC_NULL_SPECTRA
FAILED_ROUTE_SHARED_NULL_CONE_DOES_NOT_SUPPRESS_UP_DOWN_COMMUTATOR_RANK
FAILED_ROUTE_JARLSKOG_COMMUTATOR_POLYNOMIAL_NOT_DERIVED
FAILED_ROUTE_TWO_REPHASING_INVARIANT_CKM_CONSTRAINTS_NOT_DERIVED
FAILED_ROUTE_NATIVE_UP_DOWN_CLIFFORD_OPERATORS_STILL_ABSENT
FAILED_ROUTE_NULL_SPECTRUM_SHAPE_DOES_NOT_DETERMINE_EIGENBASIS_MISALIGNMENT
FIREWALL_BLOCKED_CKM_POLYNOMIAL_CONSTRAINT_NATIVE_WRITE
```

Gate 487 rejects the proposed native CKM commutator compression. The null-C3 boundary fixes a spectrum-shape baseline, but it does not fix the up/down eigenbasis mismatch from which physical CKM and Jarlskog invariants arise.

## Inherited boundary

Gate 485 remains the only accepted native shape statement in this lane:

```text
3S² - (3/2)R² = 0  ⇒  R/S = sqrt(2)  ⇒  Q = 2/3
```

Gate 486 then blocked CKM 4→2 because a null-mirror coordinate chart did not supply two rephasing-invariant CKM constraints. Gate 487 inherits that demand and tests the commutator route without CKM, Wolfenstein, quark-mass, or CP-phase inputs.

| inherited object | status |
|---|---:|
| Gate 485 Koide baseline inherited | `true` |
| Gate 486 null mirror bridge-only | `true` |
| required invariant equations from Gate 486 | `2` |
| inherited derived invariant equations | `0` |
| observed CKM imported | `false` |

## Algebraic commutator sieve

The audit constructs only a synthetic theorem probe: two Hermitian operators with the same null-C3 spectrum, then changes their relative eigenbasis by a bridge-selected unitary. This is allowed as a negative sieve, not as a native physical model.

```text
O_u = diag(x_1,x_2,x_3)
O_d = U diag(x_1,x_2,x_3) U†
x_i = S + R cos(theta_i - psi),  R/S = sqrt(2)
C_ud = [O_u,O_d]
```

| null-spectrum datum | value |
|---|---:|
| S | `1` |
| R | `1.41421356237` |
| ψ | `0.2` |
| R/S | `1.41421356237` |
| null residual `3S²-(3/2)R²` | `-4.44e-16` |
| native eigenbasis supplied | `false` |

## Commutator rank result

All cases preserve the same null-C3 spectrum. Only the relative eigenbasis changes. The commutator rank changes anyway.

| case | relative frame | rank `[O_u,O_d]` | Frobenius norm | abs det | result |
|---|---|---:|---:|---:|---|
| aligned null spectra | `I` | `0` | `0` | `0` | same null spectrum can commute exactly; rank is not forced away from zero |
| real 1-2 bridge rotation | `R12(0.4)` | `2` | `1.70934900521` | `0` | same null spectrum can yield a rank-two commutator under a bridge-chosen real eigenbasis tilt |
| complex Fourier bridge frame | `F3` | `3` | `3` | `0.828321574724` | same null spectrum can yield a full-rank commutator under a bridge-chosen complex eigenbasis frame |

Therefore the shared null cone does not suppress the commutator rank. It permits commuting, rank-two, and full-rank synthetic brackets under the same spectrum. That kills the proposed native implication from null baseline to CKM polynomial compression.

## Rephasing-invariant polynomial hunt

A genuine CKM 4→2 theorem must produce two independent relations in the physical rephasing quotient, not merely a coordinate chart or a synthetic commutator sample. Gate 487 derives none.

| invariant requirement | count/status |
|---|---:|
| physical CKM parameter dimension | `4` |
| proposed compressed dimension | `2` |
| required independent polynomial constraints | `2` |
| derived independent polynomial constraints | `0` |
| moduli polynomial relations derived | `0` |
| Jarlskog polynomial relations derived | `0` |
| commutator determinant relation derived | `false` |
| 4→2 theorem passed | `false` |

## Firewall result

```text
FIREWALL_BLOCKED_CKM_POLYNOMIAL_CONSTRAINT_NATIVE_WRITE
FAILED_ROUTE_EMPIRICAL_CKM_WOLFENSTEIN_QUARK_MASS_INPUT_REJECTED
```

No CKM matrix entries, Wolfenstein parameters, quark masses, or physical CP phase were imported. The synthetic commutator probes are bridge-only counterexamples to overclaiming; they are not native predictions.

| firewall item | status |
|---|---:|
| observed CKM imported | `false` |
| Wolfenstein imported | `false` |
| quark masses imported | `false` |
| observed CP phase imported | `false` |
| Jarlskog native prediction | `false` |
| CKM 4→2 native write | `false` |
| polynomial constraints native write | `false` |
| synthetic commutator bridge-only | `true` |
| native flavor dimension | `13` |
| K/X/Y charged coefficient dimension | `9` |

## Registry update

Native

- `no new native CKM, Jarlskog, or 4->2 theorem`
- `Gate485 null-C3 Koide baseline remains a spectrum-shape theorem only`

Bridge

- `CONDITIONAL_SUPPORT_COMMUTATOR_SIEVE_EXECUTED_ON_SYNTHETIC_NULL_SPECTRA`
- `synthetic same-null-spectrum commutator probes may be used to test future native operator candidates`

Environmental

- `observed CKM matrix, Wolfenstein parameters, quark masses, and CP phase remain forbidden theorem inputs`
- `physical Jarlskog value remains external comparator data`

Failed route

- `FAILED_ROUTE_SHARED_NULL_CONE_DOES_NOT_SUPPRESS_UP_DOWN_COMMUTATOR_RANK`
- `FAILED_ROUTE_JARLSKOG_COMMUTATOR_POLYNOMIAL_NOT_DERIVED`
- `FAILED_ROUTE_TWO_REPHASING_INVARIANT_CKM_CONSTRAINTS_NOT_DERIVED`
- `FAILED_ROUTE_NATIVE_UP_DOWN_CLIFFORD_OPERATORS_STILL_ABSENT`
- `FAILED_ROUTE_NULL_SPECTRUM_SHAPE_DOES_NOT_DETERMINE_EIGENBASIS_MISALIGNMENT`
- `FAILED_ROUTE_EMPIRICAL_CKM_WOLFENSTEIN_QUARK_MASS_INPUT_REJECTED`

Open theorem

- `CONDITIONAL_SUPPORT_GATE488_NATIVE_UP_DOWN_OPERATOR_SOURCE_SEARCH_DEFINED`
- `find a native finite operator source that distinguishes up/down sectors before any CKM invariant polynomial can be claimed`

## Next step

Gate 488 — Native Up/Down Operator Source Search. search the finite Clifford/spectral data for a native up/down operator pair O_u,O_d, or prove that all available native structures remain generation-blind before the empirical airlock

## Truth statement

Gate487 result: the shared null-C3 boundary constrains spectrum shape, not relative eigenbasis. Synthetic up/down operators with the same null spectrum realize commutator ranks [0 2 3], so the null cone does not suppress [O_u,O_d] into two physical CKM constraints. The required invariant constraints remain 2, the derived constraints remain 0, and the CKM/Jarlskog native registry write is blocked.
