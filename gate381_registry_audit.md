# Gate 381 Registry Audit — Spectral Graph Projection / `f0` Index Theorem Sieve

## Gate identity

- **Gate:** 381
- **Package:** `pkg/bridge/spectralgraphf0index`
- **Theorem:** `SpectralGraphProjectionF0IndexTheoremSieveTheorem`
- **Audit ID:** `GATE381-SPECTRAL-GRAPH-F0-INDEX-THEOREM-SIEVE`
- **Layer:** Bridge / CCM-Pfaffian coefficient closure / spectral-moment firewall

## Purpose

Gate 380 found a sharp near-closure:

```text
f0_eff ≈ 9.8971      using v = 246.22 GeV
f0_eff ≈ 9.9721      using the unreduced-Planck Pfaffian VEV
f0 = 10              predicts a near-125 GeV Higgs mass
```

The ASHA finite Dirac graph also contains exactly ten `J`-doubled edge slots:

```text
5 finite Dirac edge classes × J-doubling = 10.
```

Gate 381 asks whether this equality can be promoted from numerical/topological resonance into a theorem:

```text
f0_CCM ?= Tr_E(P_edge) = 10.
```

## Inherited facts

| Source | Inherited fact | Status |
|---|---|---|
| Gate 379 | CCM coefficient ledger is the correct almost-commutative product-action coefficient source. | Preserved |
| Gate 379 | CCM Higgs quartic read-off depends on the `f0` moment and the trace ratio `e/a²`. | Preserved |
| Gate 380 | `λ_H(f0)=π²(e/a²)/(2f0)` under the active canonical convention. | Preserved |
| Gate 380 | `e/a² = 1197/4624`. | Preserved |
| Gate 380 | `f0=10` predicts a Higgs mass near 125 GeV. | Preserved as conditional near-closure |
| Gate 380 | The finite Dirac graph has five fundamental edge classes and ten `J`-doubled edge slots. | Preserved |

## Edge projection construction

The finite Dirac edge classes are:

```text
Q_L ↔ u_R
Q_L ↔ d_R
L_L ↔ e_R
L_L ↔ ν_R
ν_R ↔ ν_R^c
```

With `J`-doubling:

```text
N_edge,J = 2 × 5 = 10.
```

Define an edge-slot projection on the finite Dirac edge/operator slot space `E_edges`:

```text
P_edge : E_edges → E_edges
Tr_E(P_edge) = dim(E_edges) = 10.
```

This part is native and exact.

## Type audit

The requested formula was:

```text
Tr_{H_F}(P_edge) = 10.
```

Gate 381 does not accept this as written without an additional embedding theorem, because the edge slots are finite Dirac operator entries/modes, not ordinary state vectors in `H_F`.

The lawful statement currently available is:

```text
Tr_E(P_edge) = 10
```

not:

```text
Tr_{H_F}(P_edge) = 10.
```

To make the latter true, the project would need a theorem embedding the edge-slot basis as an orthogonal projection subspace of `H_F` with the correct CCM trace measure.

## CCM `f0` definition audit

In the CCM spectral-action coefficient ledger, `f0` is the zeroth test-function moment/value:

```text
f0 = f(0)
```

For a unit sharp cutoff:

```text
f0 = 1.
```

The edge projection trace is instead a finite multiplicity:

```text
Tr_E(P_edge) = 10.
```

These are not the same mathematical object. A finite multiplicity can multiply a spectral-action coefficient, but it does not automatically redefine the scalar cutoff moment.

## Index-theorem sieve

The Atiyah-Singer intuition is structurally useful but does not yet prove the target identity.

A true index theorem usually counts a signed kernel/Fredholm index, schematically:

```text
Index(D) = dim ker D_+ - dim ker D_-.
```

The ASHA Rule-of-10 count is instead an unsigned count of all allowed finite Dirac interaction edge slots:

```text
#(Yukawa/Majorana edge slots) = 10.
```

Therefore the ordinary index-theorem analogy does not by itself prove:

```text
f0 = 10.
```

## Higgs near-closure inherited from Gate 380

Using the active Gate-380 formula:

```text
λ_H(f0) = π²(e/a²)/(2f0)
e/a² = 1197/4624
```

for `f0=10`:

```text
λ_H ≈ 0.12774563655
m_H ≈ 124.455 GeV    with v = 246.22 GeV
m_H ≈ 124.925 GeV    with the unreduced-Planck Pfaffian VEV
```

This is a strong near-closure. Gate 381 preserves it as conditional evidence, not final proof.

## Status ledger

### Conditional supports

```text
CONDITIONAL_SUPPORT_SPECTRAL_EDGE_PROJECTION_FORMALIZED
CONDITIONAL_SUPPORT_DISCRETE_EDGE_PROJECTION_TRACE_COMPUTED
CONDITIONAL_SUPPORT_J_DOUBLED_EDGE_PROJECTION_TRACE_EQUALS_TEN
CONDITIONAL_SUPPORT_F0_TEN_HIGGS_NEAR_CLOSURE_INHERITED
CONDITIONAL_SUPPORT_CCM_F0_MOMENT_DEFINITION_AUDITED
CONDITIONAL_SUPPORT_INDEX_THEOREM_SIEVE_EXECUTED
CONDITIONAL_SUPPORT_ANALYTIC_TO_DISCRETE_BRIDGE_AUDITED
```

### Tensions

```text
CONDITIONAL_TENSION_EDGE_MODES_ARE_OPERATOR_SLOTS_NOT_HF_STATE_VECTORS
CONDITIONAL_TENSION_CCM_F0_IS_TEST_FUNCTION_MOMENT_NOT_EDGE_TRACE
CONDITIONAL_TENSION_SHARP_CUTOFF_F0_VALUE_WOULD_BE_ONE_NOT_TEN
CONDITIONAL_TENSION_INDEX_THEOREM_COUNTS_KERNEL_OR_CHIRAL_INDEX_NOT_GENERIC_EDGES
CONDITIONAL_TENSION_EDGE_PROJECTION_TRACE_IS_MULTIPLICITY_FACTOR_NOT_MOMENT_VALUE
```

### Failed routes preserved

```text
FAILED_ROUTE_F0_MOMENT_INDEX_NOT_DERIVED
FAILED_ROUTE_EDGE_PROJECTION_TRACE_NOT_CCM_F0
FAILED_ROUTE_HIGGS_MASS_NOT_GEOMETRICALLY_SEALED
FAILED_ROUTE_CUTOFF_FUNCTION_AMPLITUDE_TEN_NOT_NATIVE
FAILED_ROUTE_FULL_NUMERICAL_TOE_CLOSURE_STILL_NOT_REACHED
```

## Exact missing theorem

The project now knows exactly what theorem would seal the route:

```text
SpectralGraphMomentTheorem:
The CCM zeroth moment functional f0 on the finite ASHA product geometry is canonically represented by the trace of the J-doubled finite Dirac edge projection:

f0 = Tr_E(P_edge) = 10.
```

But Gate 381 does not prove this theorem. It only proves the edge-side integer and the near-Higgs numerical resonance.

## Final truth statement

Gate 381 proves the precise state of the `f0=10` idea. The finite ASHA Dirac graph really has five structural edge classes and ten `J`-doubled edge slots, so `Tr_E(P_edge)=10` on the edge-slot space. This exactly matches the `f0` value that gives near-Higgs closure in Gate 380. However, CCM `f0` is the zeroth test-function moment/value `f(0)`, while the edge projection trace is a finite multiplicity over `D_F` operator slots. Edge slots are not automatically `H_F` eigenvectors, and the standard index-theorem analogy counts a signed kernel/Fredholm index rather than the unsigned set of all Yukawa/Majorana edges. Therefore Gate 381 preserves `f0=10` as a powerful capacity witness, but it does not derive `f0=10` as a native spectral-action moment. Higgs mass geometric sealing remains open.
