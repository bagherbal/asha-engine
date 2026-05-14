# Gate 289 Registry Audit — Chiral/J-Structure Anomaly Sieve / Asymmetric Trace Audit

## Gate status

Gate 289 audits whether asymmetric operators can break the branch degeneracy left by Gate 288.

Gate 288 proved that the contact-spectrum cutoff identification fixes the reduced total moments

```text
Tr(D_F²) ≈ 5.746836960723197
Tr(D_F⁴) ≈ 8.549369303330813
Tr(D_F⁴)/Tr(D_F²)² = 1197/4624
```

while allowing both scalar-Morita branches:

```text
r_+ ≈ 1.645470463011191,  X_+ ≈ 0.9680658202595966
r_- ≈ 0.672051318208557,  X_- ≈ 1.905352660102002
```

Gate 289 asks whether chiral grading `γ`, real structure `J`, anomaly equations, or projected traces can distinguish these two surviving branches.

## Inputs inherited

| Input | Value | Provenance |
|---|---:|---|
| `κ_C:κ_Q` | `1:3` | Morita trace multiplicity ledger |
| `λ_contact` | `1197/4624` | scalar/contact shape |
| `f0` | `7` | contact spectral cutoff snapshot |
| `f2` | `61/25` | contact spectral cutoff snapshot |
| `f4` | `257629/202500` | contact spectral cutoff snapshot |
| `r_+` | `1.645470463011191` | Gate 275/288 branch |
| `r_-` | `0.672051318208557` | Gate 275/288 branch |

## Chiral grading audit

On the reduced odd-Dirac edge ledger,

```text
γ = +1 on the left edge copy, -1 on the right edge copy
D_F = [[0,M],[M†,0]]
D_F² = diag(MM†, M†M)
```

Therefore, for paired left/right singular values,

```text
Tr(γD_F²) = 0
Tr(γD_F⁴) = 0
```

for both amplitude branches.

| Branch | `Tr(γD_F²)` | `Tr(γD_F⁴)` | Branch-sensitive? |
|---|---:|---:|---|
| `r_+` | `0` | `0` | no |
| `r_-` | `0` | `0` | no |

Conclusion: `γ` alone is branch-blind in the reduced odd-Dirac proxy. It sees left-minus-right, while the branch ambiguity lives in the lepton-vs-quark edge distribution.

## Sector-projected trace diagnostic

Lepton/quark projectors do distinguish the two branches:

| Branch | `Tr(P_C D_F²)` | `Tr(P_Q D_F²)` | `Tr(P_QD²)/Tr(P_CD²)` |
|---|---:|---:|---:|
| `r_+` | `0.968065820260` | `4.778771140464` | `4.936411389034` |
| `r_-` | `1.905352660102` | `3.841484300621` | `2.016153954626` |

and at fourth order:

| Branch | `Tr(P_C D_F⁴)` | `Tr(P_Q D_F⁴)` |
|---|---:|---:|
| `r_+` | `0.937151432355` | `7.612217870976` |
| `r_-` | `3.630368759358` | `4.919000543973` |

However, this is only a diagnostic. No native functional has been derived that says which sector-projected distribution is physical. Promoting this numerical difference into branch selection would reintroduce an unsealed orientation choice.

## `J` and anomaly audit

A candidate anti-linear charge-conjugation idea remains available at the architectural level, but Gate 289 does not derive the physical finite `J` satisfying all KO and representation axioms over the completed `C ⊕ H ⊕ M3(C)` finite Hilbert space.

The missing physical checks remain:

```text
J² = ±1
Jγ = ±γJ
JD_F = D_FJ
ρ°(a)=Jρ(a*)J⁻¹
order-one compatibility on the completed physical H_F
```

The anomaly route also does not select the branch. Standard anomaly cancellation is a charge/chirality constraint. In the current reduced scalar-Morita ledger, `r=|y/x|²` is an edge-amplitude distribution, not a hypercharge assignment. No derived anomaly polynomial depends on `r_+` versus `r_-`.

## Status ledger

```text
CONDITIONAL_SUPPORT_GATE288_BRANCH_MASKING_INHERITED
CONDITIONAL_SUPPORT_GAMMA_PROXY_FORMALIZED_ON_REDUCED_EDGE_LEDGER
CONDITIONAL_SUPPORT_CHIRAL_TRACES_COMPUTED
CONDITIONAL_SUPPORT_SECTOR_PROJECTED_BRANCH_SENSITIVITY_EXPOSED
CONDITIONAL_SUPPORT_J_AND_ANOMALY_REQUIREMENTS_AUDITED
CONDITIONAL_SUPPORT_CHIRAL_J_FIREWALLS_PRESERVED
FAILED_ROUTE_PHYSICAL_J_NOT_DERIVED
FAILED_ROUTE_FULL_CHIRAL_HYPERCHARGE_REPRESENTATION_MISSING
FAILED_ROUTE_GAMMA_TRACES_BRANCH_BLIND
FAILED_ROUTE_SECTOR_PROJECTED_TRACES_LACK_SELECTION_PRINCIPLE
FAILED_ROUTE_ANOMALY_CONDITIONS_DO_NOT_DEPEND_ON_R_BRANCH
FAILED_ROUTE_BRANCH_NOT_SELECTED_BY_CHIRAL_ASYMMETRY
FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED
```

## Verdict

Gate 289 exposes the exact nature of the remaining branch ambiguity:

```text
Global traces are branch-blind.
Chiral traces are also branch-blind in the reduced odd-Dirac proxy.
Sector-projected traces are branch-sensitive, but lack a native selection functional.
Physical J, full chiral/hypercharge representation, and anomaly polynomials remain underived.
```

So the branch remains unselected.

The next theorem must derive a completed physical finite Hilbert space with `J`, `γ`, hypercharge, and a branch-sensitive invariant; or it must introduce a properly quarantined seal. Gate 289 does not claim a Higgs mass ratio.
