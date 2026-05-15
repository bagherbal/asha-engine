# Gate 445 Registry Audit — Seesaw Bridge Mass-Lift / Structural-Zero Compatibility Audit

## Claim tested

Gate 445 tests whether the Gate-444 Generation-2 structural zero can be lifted by a purely off-diagonal family bridge without inserting empirical Yukawa values, muon/charm masses, CKM data, or PMNS data. The gate selects topology only; it does not select a bridge amplitude.

## Prior boundary inherited

Gate444KForced=true gen2BareZero=true noColliderData=true KXYStillFree=true nativeFlavorDim=13 conditionalKXYDim=9 verdict=CONDITIONAL_SUPPORT_GATE444_STRUCTURAL_ZERO_INHERITED

## Bridge arena

ansatz="B(a,b,c)=[[0,a,c],[a,0,b],[c,b,0]], M(ε)=K_gen+εB" epsilonSymbolic=true hermitian=true zeroDiagonal=true traceless=true familyFiber=true primitiveScan=true noYukawa=true verdict=CONDITIONAL_SUPPORT_OFFDIAGONAL_FAMILY_BRIDGE_ARENA_FORMALIZED

The tested bridge is:

```text
K_gen = diag(-1, 0, 1)
B(a,b,c) = [[0,a,c],[a,0,b],[c,b,0]]
M(ε) = K_gen + ε B(a,b,c)
```

## Boundary stack

| Boundary | Formula | Applied | Passed | Verdict | Reason |
|---|---|---:|---:|---|---|
| structural-zero preservation | `diag(B)=0 and diag(M)_2=0` | true | true | `CONDITIONAL_SUPPORT_OFFDIAGONAL_FAMILY_BRIDGE_ARENA_FORMALIZED` | the bridge may mix through the family fiber but must not insert a direct second-generation bare mass |
| endpoint balance | `|B_12|=|B_23|` | true | true | `CONDITIONAL_SUPPORT_ENDPOINT_BALANCE_BOUNDARY_APPLIED` | the K=-1 and K=+1 endpoints must couple symmetrically to the middle resonance; otherwise the lift is a lopsided source |
| determinant mass-lift | `det(K+εB) not identically zero` | true | true | `CONDITIONAL_SUPPORT_DETERMINANT_MASS_LIFT_BOUNDARY_APPLIED` | a physical mass-lift bridge must remove the exact zero eigenvalue for symbolic nonzero ε |
| primitive topology | `gcd(|a|,|b|,|c|)=1 and minimal support` | true | true | `CONDITIONAL_SUPPORT_TRIANGULAR_BRIDGE_TOPOLOGY_FORCED` | only topology, not amplitude, is allowed to be selected without empirical scale data |

## Sieve enumeration

radius=1 raw=27 primitive=26 balanced=12 balancedLift=8 openFailures=4 unbalancedLift=12 uniqueUnsigned=true signedVariants=8 verdict=CONDITIONAL_SUPPORT_TRIANGULAR_BRIDGE_TOPOLOGY_FORCED

The bounded scan uses primitive integer edge weights in `{−1,0,1}`. This is not an amplitude scan; it is a support-topology sieve.

### Endpoint-balanced mass-lift survivors

| Weights | Support | Det polynomial | Closed triangle | Canonical topology |
|---|---:|---|---:|---:|
| `(a=-1,b=-1,c=-1)` | 3 | `-2 ε^3` | true | true |
| `(a=-1,b=-1,c=1)` | 3 | `2 ε^3` | true | true |
| `(a=-1,b=1,c=-1)` | 3 | `2 ε^3` | true | true |
| `(a=-1,b=1,c=1)` | 3 | `-2 ε^3` | true | true |
| `(a=1,b=-1,c=-1)` | 3 | `2 ε^3` | true | true |
| `(a=1,b=-1,c=1)` | 3 | `-2 ε^3` | true | true |
| `(a=1,b=1,c=-1)` | 3 | `-2 ε^3` | true | true |
| `(a=1,b=1,c=1)` | 3 | `2 ε^3` | true | true |

### Open-chain failures

| Weights | Support | Det polynomial | Failure status |
|---|---:|---|---|
| `(a=-1,b=-1,c=0)` | 2 | `0` | `FAILED_ROUTE_OPEN_CHAIN_BRIDGE_PRESERVES_ZERO_DETERMINANT` |
| `(a=-1,b=1,c=0)` | 2 | `0` | `FAILED_ROUTE_OPEN_CHAIN_BRIDGE_PRESERVES_ZERO_DETERMINANT` |
| `(a=1,b=-1,c=0)` | 2 | `0` | `FAILED_ROUTE_OPEN_CHAIN_BRIDGE_PRESERVES_ZERO_DETERMINANT` |
| `(a=1,b=1,c=0)` | 2 | `0` | `FAILED_ROUTE_OPEN_CHAIN_BRIDGE_PRESERVES_ZERO_DETERMINANT` |

## Analytic determinant collapse

K="diag(-1,0,1)" bridge="B(a,b,c) with a=B12, b=B23, c=B13" det="det(K+εB)=(b²-a²)ε²+2abc ε³" balance="|a|=|b|" reduction="endpoint balance cancels the ε² term" open="c=0 ⇒ det(K+εB)=0, so the structural zero survives" triangle="abc≠0 ⇒ det(K+εB)=2abc ε³, so the zero is lifted at cubic order" middleOrder="λ₂^eff = O(ε³) for the balanced primitive bridge" triangleForced=true xSupport=true fixesAmplitude=false fixesSign=false verdict=CONDITIONAL_SUPPORT_GEN2_SEESAW_MASS_LIFT_COMPATIBLE

For `B(a,b,c)` the exact symbolic identity is:

```text
det(K_gen + ε B) = (b^2 - a^2) ε^2 + 2abc ε^3
```

Endpoint balance gives `|a|=|b|`, so the `ε²` term cancels. An open chain has `c=0`, hence determinant zero and the middle structural zero survives. A balanced nonzero lift therefore requires `a,b,c` all nonzero. Primitive normalization reduces the unsigned support to the three-edge triangle.

## Geometrically forced bridge topology

Generation-2 structural-zero mass-lift bridge topology="primitive closed triangular family bridge / X_gen support" support=3 trace=0 traceSquare=6 detLeading="det(K+εB)=2ε³ for the positive canonical representative" gen2DiagZero=true liftsZero=true forcedTopology=true amplitudeSealed=true signSealed=true yukawaPredicted=false muonCharmPredicted=false verdict=CONDITIONAL_SUPPORT_X_GEN_SUPPORT_SELECTED_AS_MINIMAL_TOPOLOGY

Canonical positive representative:

```text
B_lift = [[0,1,1],[1,0,1],[1,1,0]]
det(K_gen + ε B_lift) = 2 ε^3
```

This is the real `X_gen` support topology. The gate does not fix the sign orientation, complex phase, or sector coefficient multiplying this topology.

## Phenomenology/firewall audit

muonImported=false charmImported=false yukawaImported=false CKM=false PMNS=false topologyConditional=true amplitudeSealed=true signedPhaseSealed=true physicalMassNeedsData=true nativeDim=13→13 KXYFree=9 verdict=FIREWALL_REFINED_BRIDGE_TOPOLOGY_FORCED_BUT_AMPLITUDE_SEALED

The result is a topology theorem, not a mass theorem. It says the second-family bare zero is lift-compatible through a closed triangular mixing bridge. The observed muon/charm mass values still require bridge amplitude and sector-source data outside Gate 445.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE444_STRUCTURAL_ZERO_INHERITED`
- `CONDITIONAL_SUPPORT_OFFDIAGONAL_FAMILY_BRIDGE_ARENA_FORMALIZED`
- `CONDITIONAL_SUPPORT_ENDPOINT_BALANCE_BOUNDARY_APPLIED`
- `CONDITIONAL_SUPPORT_DETERMINANT_MASS_LIFT_BOUNDARY_APPLIED`
- `CONDITIONAL_SUPPORT_TRIANGULAR_BRIDGE_TOPOLOGY_FORCED`
- `CONDITIONAL_SUPPORT_GEN2_SEESAW_MASS_LIFT_COMPATIBLE`
- `CONDITIONAL_SUPPORT_X_GEN_SUPPORT_SELECTED_AS_MINIMAL_TOPOLOGY`
- `CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED`
- `FAILED_ROUTE_OPEN_CHAIN_BRIDGE_PRESERVES_ZERO_DETERMINANT`
- `FAILED_ROUTE_UNBALANCED_BRIDGE_LIFTS_BY_ASYMMETRIC_SOURCE`
- `FAILED_ROUTE_BRIDGE_AMPLITUDE_NOT_PREDICTED`
- `FAILED_ROUTE_SIGNED_CYCLE_ORIENTATION_UNFIXED`
- `FAILED_ROUTE_NO_MUON_CHARM_MASS_VALUE_PREDICTION`
- `FIREWALL_REFINED_BRIDGE_TOPOLOGY_FORCED_BUT_AMPLITUDE_SEALED`

## Next gate

Gate 446 — Signed-Cycle / Complex Phase Orientation Sieve: Gate 445 fixes the unsigned triangle support but leaves signed cycle orientation and complex phase data sealed.

## Truth statement

Gate 445 proves a narrow compatibility theorem: once Gate 444 fixes K_gen=diag(-1,0,1) and the second bare diagonal level is zero, an endpoint-balanced off-diagonal bridge can lift that zero only if the family graph closes into the primitive triangle. Open chains keep det(K+εB)=0. Thus the unsigned X_gen support topology is forced as the minimal seesaw/mixing bridge, while the coefficient ε, signed cycle orientation, sector Yukawa amplitudes, and observed muon/charm masses remain sealed.
