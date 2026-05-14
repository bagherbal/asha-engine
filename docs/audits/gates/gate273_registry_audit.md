# Gate 273 Registry Audit — Weak/Quaternionic Sub-Bimodule Selector / Finite Inner-Product Normalization Audit

## Gate identity

- **Gate:** 273
- **Package:** `pkg/bridge/weakquaternionicnormalization`
- **Theorem:** `WeakQuaternionicSubBimoduleSelectorFiniteInnerProductNormalizationAuditTheorem`
- **Registry status:** `theorem.BridgeRequired`
- **Audit ID:** `GATE273-WEAK-QUATERNIONIC-SUB-BIMODULE-SELECTOR-FINITE-INNER-PRODUCT-NORMALIZATION-AUDIT`

## Boundary inherited from Gate 272

Gate 272 repaired the representation category for the finite spectral-action branch:

```text
A_F = C ⊕ M3(C)
H_ij = V_i ⊗ V_j*,  i,j ∈ {C,Q}
```

with universal Morita summands:

```text
H_CC dim 1
H_CQ dim 3
H_QC dim 3
H_QQ dim 9
```

The order-one edge rule inherited from Gate 272 is:

```text
H_ij ↔ H_kl is order-one compatible if i=k or j=l.
It is non-vacuous for left one-forms only if i≠k.
```

Therefore the non-vacuous order-one edges are:

```text
H_CC ↔ H_QC   right C shared   amplitude label m_C
H_CQ ↔ H_QQ   right Q shared   amplitude label m_Q
```

Gate 272 did not lock `m_C:m_Q`; Gate 273 audits whether weak/chiral/quaternionic sub-bimodule selection and inner-product normalization can lock it.

## Result summary

Gate 273 computes a legitimate finite Morita trace-multiplicity ledger:

```text
κ_C : κ_Q = 1 : 3
```

This is the expected `1⊕3` trace weighting of the lepton-like and quark-like right sectors. However, the gate proves that this is only a multiplicity weighting. It does **not** determine the independent finite Dirac amplitudes `x:y`, nor the edge-map norms `||T_C||` and `||T_Q||`.

Final status:

```text
CONDITIONAL_SUPPORT_GATE272_MORITA_BIMODULE_LEDGER_INHERITED
CONDITIONAL_SUPPORT_WEAK_CHIRAL_SUB_BIMODULE_SIEVE_AUDITED
CONDITIONAL_SUPPORT_ORDER_ONE_NONVACUOUS_EDGES_RECOVERED
CONDITIONAL_SUPPORT_FINITE_INNER_PRODUCT_NORMALIZATION_LEDGER_BUILT
CONDITIONAL_SUPPORT_LEPTON_QUARK_TRACE_MULTIPLICITIES_COMPUTED
CONDITIONAL_SUPPORT_NORMALIZED_TRACE_MOMENTS_REEVALUATED
FAILED_ROUTE_WEAK_QUATERNIONIC_SELECTOR_NOT_NATIVE_TO_C_PLUS_M3C
FAILED_ROUTE_PHYSICAL_SM_SUB_BIMODULE_NOT_DERIVED
FAILED_ROUTE_EDGE_MAP_NORMS_REMAIN_UNSELECTED
FAILED_ROUTE_INNER_PRODUCT_NORMALIZATION_DOES_NOT_LOCK_XY_RATIO
FAILED_ROUTE_CANONICAL_DF_AMPLITUDES_NOT_LOCKED_VIA_NORMALIZATION
FAILED_ROUTE_A2_A4_HIGGS_RATIO_STILL_NOT_DERIVED
FAILED_ROUTE_EMPIRICAL_YUKAWA_SEAL_REMAINS_ACTIVE
```

## Detailed findings

| Category | Finding | Status | Rigor assessment |
| --- | --- | --- | --- |
| Gate 272 inheritance | The first-quantized Morita arena and two non-vacuous order-one edges are inherited. | `CONDITIONAL_SUPPORT` | The gate starts from the lawful spectral-triple carrier, not full second-quantized `S_C`. |
| Weak/chiral sieve | The native Morita chiral edge sieve retains `H_CC↔H_QC` and `H_CQ↔H_QQ`. | `CONDITIONAL_SUPPORT` | Correctly preserves the Gate 272 order-one structure. |
| Quaternionic selector | A true weak/quaternionic selector is not native to the active algebra `C⊕M3(C)`. | `FAILED_ROUTE` | Prevents importing Connes' Standard Model algebra or a physical Hilbert space by hand. |
| Inner product | The Morita Hilbert inner product gives canonical trace multiplicities `κ_C=1`, `κ_Q=3`. | `CONDITIONAL_SUPPORT` | Computes a real geometric weighting ledger. |
| Amplitude locking | Multiplicities do not fix `x:y`; edge-map norms remain independent. | `FAILED_ROUTE` | A trace count is not a mass/amplitude theorem. |
| Spectral ratio | Normalized raw trace moments still vary with `x:y`. | `FAILED_ROUTE` | `a₂/a₄` and the Higgs ratio remain blocked. |

## Finite inner-product normalization ledger

The finite Morita inner product is audited as:

```text
⟨u⊗φ, v⊗ψ⟩ = ⟨u,v⟩_{V_i} · ⟨ψ,φ⟩_{V_j*}
```

For minimal normalized order-one edge maps, the trace contribution is:

```text
Tr_edge(D²) ∝ dim(V_j) |m_j|² ||T_j||²
```

Thus:

```text
right C edge: H_CC ↔ H_QC, dim(V_C)=1, κ_C=1
right Q edge: H_CQ ↔ H_QQ, dim(V_Q)=3, κ_Q=3
```

The gate explicitly separates the geometric multiplicity from the still-missing amplitude theorem:

```text
κ_C:κ_Q = 1:3      derived multiplicity ledger
x:y                not derived
||T_C||:||T_Q||    not derived
```

## Normalized trace moment stress test

Using the normalized multiplicities but leaving `x,y` free:

```text
Tr(D_F²) proxy = κ_C |x|² + κ_Q |y|²
Tr(D_F⁴) proxy = κ_C |x|⁴ + κ_Q |y|⁴
```

with `κ_C=1`, `κ_Q=3`, representative legal choices give:

| Representative | `Tr(D_F²)` | `Tr(D_F⁴)` | Ratio |
| --- | ---: | ---: | ---: |
| `x=1, y=1` | 4 | 4 | 1 |
| `x=2, y=1` | 7 | 19 | 0.368421052631579 |
| `x=1, y=2` | 13 | 49 | 0.26530612244898 |

The ratio remains amplitude-dependent, so no invariant Seeley-de Witt ratio is derived.

## Firewall ledger

Gate 273 preserves the following firewalls:

- No observed fermion mass inserted.
- No CKM/PMNS data inserted.
- No VEV inserted.
- No cutoff scale inserted.
- No Standard Model quaternionic algebra imported as a theorem.
- No Higgs mass prediction claimed.
- No multiplicity ledger promoted into an amplitude theorem.
- `EmpiricalYukawaSeal` remains active.

## Future theorem criteria

To reopen the `a₂/a₄` route lawfully, a future gate must derive at least the following:

1. A native weak/quaternionic algebra or equivalent finite selector inside the ASHA finite geometry.
2. A physical finite Hilbert sub-bimodule, not just the universal Morita ledger.
3. A physical anti-linear `J` with particle/antiparticle semantics.
4. Edge-map norms `||T_C||` and `||T_Q||`, or a finite action theorem that replaces them.
5. A canonical rule locking `x:y` before any comparison to Higgs data.
6. A heat-kernel / Seeley-de Witt projection and subtraction scheme.

Recommended next gate:

```text
Gate 274 — Native Weak Quaternionic Algebra / Physical Finite Hilbert Space Reconstruction Audit
```

## Validation commands

Focused tests only:

```bash
go test -p=1 ./pkg/bridge/weakquaternionicnormalization -count=1 -timeout=120s -v

go test -p=1 ./pkg/bridge/weakquaternionicnormalization ./pkg/bridge/moritabimodulesearch ./pkg/bridge/fullscrepresentationsearch -count=1 -timeout=120s -v

go list ./internal/app

go list ./cmd/asha
```

No full internal tests, full package tests, or `go test ./...` were run.
