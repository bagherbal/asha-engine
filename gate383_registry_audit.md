# Gate 383 Registry Audit — Spectral Graph Trace / Node-to-Edge Kinetic Normalization Sieve

**Audit ID:** `GATE383-SPECTRAL-GRAPH-TRACE-NODE-TO-EDGE-KINETIC-NORMALIZATION-SIEVE`  
**Package:** `pkg/bridge/spectralgraphtracenormalization`  
**Theorem:** `SpectralGraphTraceNodeToEdgeKineticNormalizationSieveTheorem`  
**Layer:** Bridge / Product spectral-action coefficient normalization  
**Status:** Bridge required; Higgs near-closure strengthened, not sealed.

## 1. Purpose

Gate 382 isolated the last Higgs-sector normalization gap as the ratio

```text
10/7
```

where:

- `7` is the contact-node ledger used in the earlier `f0=7` channel;
- `10` is the five finite Dirac edge classes doubled by the real structure `J`.

Gate 383 audits whether the Higgs kinetic term in the product spectral action is mathematically normalized by the node trace or by the finite Dirac edge trace.

## 2. Inherited facts

| Source | Fact |
|---|---|
| Gate 379 | CCM coefficient ledger must be used directly for the almost-commutative product geometry. |
| Gate 380 | `λ_H(f0)=π²(e/a²)/(2f0)` with `e/a²=1197/4624`; Higgs boundary gives `f0_eff≈10`. |
| Gate 381 | The finite Dirac graph has `Tr_E(P_edge)=10`, but this is not CCM `f0`. |
| Gate 382 | `f0=1` belongs to the continuous test-function moment; factor `10` can only enter through finite trace/canonical normalization. |

## 3. Node versus edge trace domains

| Domain | Count | Slot | Formula | Verdict |
|---|---:|---|---|---|
| Contact node trace | `7` | finite contact/vacuum support measure | `Tr_node(P_contact)=7` | Native finite support count, but not automatically the CCM test-function moment and not the natural support of `D_F` interaction terms. |
| J-doubled Dirac edge trace | `10` | finite interaction / endomorphism channel support | `Tr_edge(P_DF)=2×5=10` | Correct-looking support for inner fluctuations and Higgs kinetic normalization, but not allowed to replace CCM `f0`. |

The native graph bridge is exact:

```text
N_edge,J / N_node = 10/7
```

## 4. Higgs kinetic trace audit

The Higgs field in the finite spectral triple arises as an inner fluctuation / finite one-form along allowed finite Dirac channels. Therefore the kinetic trace is structurally edge-supported:

```text
a |D_μH|²,      a = Tr_F(Y†Y)
```

This is the key conceptual win of Gate 383: the Higgs kinetic term is not a bare trace over contact nodes. It is supported by the interaction edges of `D_F`.

However, CCM already packages this finite support into the trace invariant `a`. Therefore this gate does **not** prove that one may apply a new `7→10` replacement after the stored ratio

```text
e/a² = 1197/4624
```

has already been computed. That ratio may already contain the edge support.

## 5. Higgs coefficient lanes

Using

```text
e/a² = 1197/4624
λ = π²(e/a²)/(2D)
m_H = v sqrt(2λ)
```

Gate 383 computes:

| Lane | Denominator `D` | Meaning | Result |
|---|---:|---|---|
| Contact-node denominator | `7` | old contact ledger | Overpredicts Higgs mass. |
| Edge-trace denominator | `10` | node-to-edge correction | Near-closes Higgs mass around the Gate-380 Pfaffian VEV lane. |
| Unit sharp-cutoff denominator | `1` | no finite kinetic normalization conversion | Strongly overpredicts Higgs mass. |

The strongest lane remains:

```text
λ_edge = π²(1197/4624)/(2×10)
```

which reproduces the near-125 GeV Higgs closure inherited from Gate 380.

## 6. Final obstruction

The exact missing theorem is now:

```text
SpectralGraphKineticNormalizationTheorem:
The canonical Higgs kinetic trace in the ASHA finite spectral triple is edge-normalized
with N_edge,J=10 rather than contact-node-normalized with N_node=7, and this replacement
is not already contained in the stored full finite trace ratio e/a².
```

To prove it, the project must recompute raw `a` and `e` under separate node-measure and edge-measure conventions rather than applying `10/7` after the fact.

## 7. Status ledger

Positive / conditional supports:

```text
CONDITIONAL_SUPPORT_NODE_EDGE_TRACE_DOMAINS_FORMALIZED
CONDITIONAL_SUPPORT_HIGGS_KINETIC_TRACE_SUPPORT_AUDITED
CONDITIONAL_SUPPORT_EDGE_SUPPORTED_KINETIC_TRACE_WITNESS_FOUND
CONDITIONAL_SUPPORT_TEN_OVER_SEVEN_NODE_TO_EDGE_BRIDGE_COMPUTED
CONDITIONAL_SUPPORT_EDGE_DENOMINATOR_REPRODUCES_NEAR_HIGGS_CLOSURE
CONDITIONAL_SUPPORT_PFAFFIAN_HIGGS_MASS_LANE_COMPUTED
CONDITIONAL_SUPPORT_CCM_MOMENT_REMAINS_SEPARATE_FROM_GRAPH_MULTIPLICITY
```

Tensions:

```text
CONDITIONAL_TENSION_KINETIC_TERM_IS_EDGE_SUPPORTED_BUT_CCM_USES_A_TRACE
CONDITIONAL_TENSION_E_OVER_A2_MAY_ALREADY_INCLUDE_EDGE_TRACE_NORMALIZATION
CONDITIONAL_TENSION_TEN_OVER_SEVEN_REQUIRES_RAW_A_AND_E_RECOMPUTATION
CONDITIONAL_TENSION_NODE_TO_EDGE_BRIDGE_NOT_YET_A_CCM_THEOREM
```

Failed routes / firewalls:

```text
FAILED_ROUTE_EDGE_TRACE_NORMALIZATION_NOT_DERIVED
FAILED_ROUTE_TEN_OVER_SEVEN_NOT_DERIVED
FAILED_ROUTE_HIGGS_MASS_NOT_GEOMETRICALLY_SEALED
FAILED_ROUTE_FULL_NUMERICAL_TOE_CLOSURE_STILL_NOT_REACHED
```

## 8. Truth statement

Gate 383 proves the conceptual correction: the Higgs kinetic term is structurally supported on finite Dirac interaction edges, while the old contact ledger counted seven nodes. This gives an exact native graph conversion `10/7`, and the edge-denominator lane reproduces the near-125 GeV Higgs mass. But the gate does not seal the Higgs mass, because CCM canonical normalization already uses `a=Tr_F(Y†Y)`, and the stored ratio `e/a²=1197/4624` may already include the relevant edge support.

The missing final theorem is a raw `a/e` recomputation under node versus edge measures proving that `7` must be replaced by `10` in the kinetic normalization denominator without double-counting.
