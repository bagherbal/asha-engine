# Gate 446 Registry Audit — Signed-Cycle / Complex Phase Orientation Sieve

## Claim tested

Gate 446 tests whether the Gate-445 triangular bridge support collapses further to one signed real cycle or one complex CP phase. It deliberately does not import muon/charm masses, Yukawa matrices, CKM data, or PMNS data.

## Prior boundary inherited

Gate445K=true XSupport=true amplitudeSealed=true signSealed=true noEmpiricalMasses=true nativeDim=13 KXYFree=9 verdict=CONDITIONAL_SUPPORT_GATE445_TRIANGLE_TOPOLOGY_INHERITED

## Orientation arena

K="diag(-1,0,1)" ansatz="B(z12,z23,z13) Hermitian, zero-diagonal, full triangular support" hermitian=true zeroDiagonal=true triangle=true endpointBalanced=true vertexGauge=true invariant="Phi=arg(z12 z23 conjugate(z13))" empiricalImported=false verdict=CONDITIONAL_SUPPORT_HERMITIAN_TRIANGULAR_CYCLE_ARENA_FORMALIZED

```text
K_gen = diag(-1,0,1)
B = [[0,z12,z13],[conj(z12),0,z23],[conj(z13),conj(z23),0]]
Phi = arg(z12 z23 conjugate(z13))
det(K_gen + eps B) = (|z23|^2-|z12|^2) eps^2 + 2 Re(z12 z23 conjugate(z13)) eps^3
```

## Boundary stack

| Boundary | Formula | Applied | Passed | Verdict | Reason |
|---|---|---:|---:|---|---|
| Gate-445 support inheritance | `support(B)={(1,2),(2,3),(1,3)}` | true | true | `CONDITIONAL_SUPPORT_GATE445_TRIANGLE_TOPOLOGY_INHERITED` | the unsigned closed triangle is already forced by the mass-lift bridge sieve |
| Hermitian/J/Gamma compatibility | `B_ji=conjugate(B_ij), diag(B)=0, [B,Gamma_gen]=0` | true | true | `CONDITIONAL_SUPPORT_J_GAMMA_ETA_TRACE_BOUNDARIES_APPLIED` | the family bridge remains an internal Hermitian source and does not alter chirality/gauge representation slots |
| eta-graded trace neutrality | `Tr(B)=0 and Tr_eta(B)=0 in family-only lift` | true | true | `CONDITIONAL_SUPPORT_J_GAMMA_ETA_TRACE_BOUNDARIES_APPLIED` | all zero-diagonal family-cycle orientations are trace neutral; this boundary cannot distinguish signs or phases |
| vertex rephasing quotient | `B_ij -> exp(i(theta_i-theta_j)) B_ij` | true | true | `CONDITIONAL_SUPPORT_CYCLE_PHASE_INVARIANT_IDENTIFIED` | two edge phases are gauge convention; only the cycle phase Phi survives as a rephasing invariant |
| determinant mass-lift | `Re(z12 z23 conjugate(z13)) != 0` | true | true | `CONDITIONAL_SUPPORT_CP_PHASE_CAPACITY_AUDITED` | mass lift excludes purely imaginary cycle product but still leaves infinitely many real/complex orientations |

## Real signed-cycle sieve

candidates=8 positive=4 negative=4 Z2classes=2 uniqueSigned=false verdict=FAILED_ROUTE_SIGNED_CYCLE_ORIENTATION_NOT_UNIQUE reason=the eight real sign assignments collapse under vertex sign flips to two invariant cycle-product classes; both pass all structural boundaries

| Weights | Product | Gauge class | Cycle phase | Determinant leading term | Representative |
|---|---:|---|---|---|---:|
| `(a=-1,b=-1,c=-1)` | -1 | negative cycle product | `pi` | `-2 eps^3` | false |
| `(a=-1,b=1,c=1)` | -1 | negative cycle product | `pi` | `-2 eps^3` | false |
| `(a=1,b=-1,c=1)` | -1 | negative cycle product | `pi` | `-2 eps^3` | false |
| `(a=1,b=1,c=-1)` | -1 | negative cycle product | `pi` | `-2 eps^3` | true |
| `(a=-1,b=-1,c=1)` | 1 | positive cycle product | `0` | `2 eps^3` | false |
| `(a=-1,b=1,c=-1)` | 1 | positive cycle product | `0` | `2 eps^3` | false |
| `(a=1,b=-1,c=-1)` | 1 | positive cycle product | `0` | `2 eps^3` | false |
| `(a=1,b=1,c=1)` | 1 | positive cycle product | `0` | `2 eps^3` | true |

After quotienting by vertex sign flips, the eight sign assignments do not collapse to one class. They collapse only to two invariant cycle-product classes: `abc=+1` and `abc=-1`. Both are trace neutral, J/Gamma compatible, and mass-lift compatible.

## Complex phase sieve

invariant="Phi=arg(z12 z23 conjugate(z13))" determinant="det(K+eps B)=2 r^3 cos(Phi) eps^3 for |z12|=|z23|=|z13|=r" CPmap="CP: Phi -> -Phi" massCondition="cos(Phi) != 0" witness="CP-odd cycle witness proportional to sin(Phi)" continuum=true CPpairs=true uniquePhase=false CPValuePredicted=false verdict=FAILED_ROUTE_COMPLEX_PHASE_CONTINUUM_UNDERDETERMINED samples=[0 det=2 eps^3 cp=false mass=true | pi det=-2 eps^3 cp=false mass=true | pi/4 det=1.41421 eps^3 cp=true mass=true | -pi/4 det=1.41421 eps^3 cp=true mass=true | pi/2 det=0 eps^3 cp=true mass=false]

Endpoint balance gives the exact reduction:

```text
det(K_gen + eps B) = 2 r^3 cos(Phi) eps^3
CP-odd cycle witness ∝ sin(Phi)
CP maps Phi -> -Phi
```

| Phase sample | det leading | CP witness | Mass lift | CP capable | CP conjugate |
|---|---|---|---:|---:|---|
| `0` | `2 eps^3` | `0` | true | false | `0` |
| `pi` | `-2 eps^3` | `0` | true | false | `pi` |
| `pi/4` | `1.41421 eps^3` | `0.707107` | true | true | `-pi/4` |
| `-pi/4` | `1.41421 eps^3` | `-0.707107` | true | true | `pi/4` |
| `pi/2` | `0 eps^3` | `1` | false | true | `-pi/2` |

The mass-lift boundary removes the purely imaginary cycle products with `cos(Phi)=0`, but it does not select a unique value of `Phi`. CP-capable pairs such as `Phi=±pi/4` survive as conjugate orientations.

## Orientation conclusion

XSupport=true signedForced=false complexForced=false YNative=false phaseCoeffFixed=false CPViolationPredicted=false massLiftCompatible=true verdict=FIREWALL_PRESERVED_PHASE_ORIENTATION_QUARANTINED reason=the boundary intersection preserves Gate-445 mass-lift compatibility but does not collapse the signed or complex orientation to one survivor

Gate 446 therefore does not promote `Y_gen` or any CP phase value to native law-space. It only confirms that Gate-445 `X_gen` support is compatible with both CP-even and CP-odd phase orientations.

## Phenomenology/firewall audit

muonImported=false charmImported=false yukawaImported=false CKM=false PMNS=false XForced=true amplitudeSealed=true signSealed=true phaseSealed=true YQuarantined=true nativeDim=13→13 KXYFree=9 verdict=FIREWALL_PRESERVED_PHASE_ORIENTATION_QUARANTINED

The firewall remains intact: no observed masses, Yukawa coefficients, CKM angles, CKM CP phase, or PMNS data were used. The native charged flavor dimension remains 13, and the conditional K/X/Y coefficient ledger remains 9.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE445_TRIANGLE_TOPOLOGY_INHERITED`
- `CONDITIONAL_SUPPORT_HERMITIAN_TRIANGULAR_CYCLE_ARENA_FORMALIZED`
- `CONDITIONAL_SUPPORT_J_GAMMA_ETA_TRACE_BOUNDARIES_APPLIED`
- `CONDITIONAL_SUPPORT_REAL_SIGNED_CYCLE_SIEVE_COMPLETED`
- `CONDITIONAL_SUPPORT_CYCLE_PHASE_INVARIANT_IDENTIFIED`
- `CONDITIONAL_SUPPORT_CP_PHASE_CAPACITY_AUDITED`
- `CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED`
- `FAILED_ROUTE_SIGNED_CYCLE_ORIENTATION_NOT_UNIQUE`
- `FAILED_ROUTE_COMPLEX_PHASE_CONTINUUM_UNDERDETERMINED`
- `FAILED_ROUTE_CP_PHASE_VALUE_NOT_PREDICTED`
- `FAILED_ROUTE_Y_GEN_PHASE_QUADRATURE_NOT_NATIVE`
- `FAILED_ROUTE_NO_MUON_CHARM_MASS_VALUE_PREDICTION`
- `FIREWALL_PRESERVED_PHASE_ORIENTATION_QUARANTINED`

## Next gate

Gate 447 — Sector-Coefficient Source Ledger / Amplitude Firewall Closure: Gate 446 proves that phase orientation is not selected by the current boundary stack.

## Truth statement

Gate 446 is a negative but important sieve result: the Gate-445 triangle support remains forced, but the signed real cycle and the complex CP phase are not forced by Hermiticity, J/Gamma compatibility, eta-trace neutrality, determinant mass-lift, or vertex rephasing quotient. Real signs reduce only to two Z2 cycle classes, and complex Hermitian bridges retain the gauge-invariant continuum Phi=arg(z12 z23 conjugate(z13)). Therefore Y_gen and the CP phase remain quarantined; no muon/charm mass value, CKM phase, or Yukawa coefficient is predicted.
