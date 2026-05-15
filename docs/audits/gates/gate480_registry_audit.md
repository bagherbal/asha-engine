# Gate 480 Registry Audit — Algebraic Null-Cone Bridge & I_K Selection

## Verdict

`CONDITIONAL_SUPPORT_I_K_DERIVED_FROM_NATIVE_NULL_BOUNDARY`

Gate 480 confirms that the native Clifford null-cone calculation can select a **bare vacuum bridge baseline**:

```text
q(a,b,c)=a^2-b^2-c^2=a^2-r^2
q=0, a>0, r>0  =>  a=r
alpha_vac=a/r=1
I_K=alpha/sqrt(alpha^2+3)=1/2
```

## Native null-cone map

| object | result |
|---|---|
| Clifford signature | `Cℓ(1,7)` |
| hierarchy leg | `K_gen hierarchy leg` |
| bridge plane | `X_triangle cosine bridge, Y_phase sine bridge` |
| radial bridge amplitude | `r=sqrt(b^2+c^2)` |
| quadratic form | `q(a,b,c)=a^2-b^2-c^2=a^2-r^2` |
| boundary | `q=0 on the bare family bridge` |

The null cone is native to `Cℓ(1,7)`. The specific statement that the **bare family bridge** lies on that null cone is recorded as the Gate 480 boundary ledger. It was not derived by the earlier raw-mass, electroweak, or PMNS preflight gates.

## Equipartition sieve

| case | a | b | c | r | q=a²-r² | alpha | I_K | accepted |
|---|---:|---:|---:|---:|---:|---:|---:|---|
| null future bridge | 1 | 1 | 0 | 1 | 0 | 1 | 0.5 | true |
| null rotated bridge | 1 | 0.5 | 0.866025403784 | 1 | 0 | 1 | 0.5 | true |
| timelike hierarchy-heavy bridge | 2 | 1 | 0 | 1 | 3 | 2 | 0.755928946018 | false |
| spacelike mixing-heavy bridge | 1 | 2 | 0 | 2 | -3 | 0.5 | 0.277350098113 | false |

Modulo positive scale and rotation inside the X/Y bridge plane, the null branch selects `alpha_vac=1` and therefore `I_K=0.5`.

## Firewall boundary

Gate 480 does **not** compute physical sector coordinates. It does not compute `d_ud`, `d_eν`, CKM entries, PMNS entries, masses, Yukawas, or branch sheets. It exports only a vacuum baseline.

Rejected promotions:

```text
FAILED_ROUTE_NULL_BOUNDARY_NOT_FORCED_BY_PREVIOUS_GATES
FAILED_ROUTE_K_TIMELIKE_XY_SPACELIKE_ASSIGNMENT_REQUIRES_EXPLICIT_BOUNDARY_LEDGER
FAILED_ROUTE_NULL_BASELINE_DOES_NOT_SOLVE_SECTOR_COORDINATES
FAILED_ROUTE_NULL_CONE_AS_CKM_PMNS_PREDICTION_REJECTED
FAILED_ROUTE_NULL_BASELINE_AS_PHYSICAL_SECTOR_I_K_REJECTED
```

## Numerical output

```text
alpha_vac = 1.000000000000
I_K,vac   = 0.500000000000
d_ud      = undefined
d_eν      = undefined
CKM/PMNS  = not constructed
```

## Next step

Gate 481 — Null-baseline perturbation ledger: define bridge-only perturbation variables around alpha_vac=1 and test whether quark/lepton residual ledgers can be expressed as deviations from the null baseline without native-promotion
