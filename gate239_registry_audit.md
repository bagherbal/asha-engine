# Gate 239 Registry Audit — Orientation Operator (`χ`) / True Chirality Derivation Audit

## Gate

**Gate 239 — Orientation Operator (`χ`) / True Chirality Derivation Audit**

Package:

```text
pkg/bridge/orientationtruechirality
```

Theorem:

```text
BRIDGE-ORIENTATION-TRUE-CHIRALITY-DERIVATION-AUDIT
```

## Verdict

```text
CONDITIONAL_SUPPORT_CLIFFORD_VOLUME_ORIENTATION_PREFLIGHT
CONDITIONAL_SUPPORT_TAU_ETA_ORIENTATION_FUNCTIONAL_INHERITED
FAILED_ROUTE_DISTINCT_ORIENTATION_CHI_DERIVATION
FAILED_ROUTE_TAU_ETA_TO_SC_OPERATOR_PULLBACK
FAILED_ROUTE_TRUE_CHIRALITY_PLANE_SELECTION
FAILED_ROUTE_LEFT_HANDED_WEAK_ACTION_DERIVATION
FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED
```

Gate 239 tests the proposed route from finite orientation to physical chirality. It does **not** tune signs or import the Standard Model chirality operator.

## Inherited state

Gate 238 proved:

```text
γ = (-1)^N
```

is a valid finite grading on the complexified Fock spinor, but every candidate weak plane has mixed parity:

```text
weak doublets = 4 even + 4 odd
weak singlets = 4 even + 4 odd
```

Therefore occupation parity is not yet Standard Model chirality.

Gate 239 asks whether the finite orientation data supplies a stronger operator `χ`.

## Candidate 1 — Clifford volume orientation

The audited candidate is:

```text
χ_vol = Clifford-volume candidate acting on Λ*(W)
```

In the current exterior/Fock realization:

```text
S_C = Λ*(W),   dim_C(S_C)=16,   dim_R(S_C)=32
```

The finite volume/chirality candidate is proportional to occupation parity:

```text
χ_vol ∝ (-1)^N = γ
```

Audit result:

| Property | Result |
|---|---:|
| Acts on `S_C` | yes |
| `χ_vol` eigenspaces | `8 + 8` |
| Equivalent to `γ` | yes |
| Commutes with `γ` | yes |
| Anticommutes with `γ` | no |
| Distinct eigenspaces from `γ` | no |
| Manual sign adjustment | no |

Thus `χ_vol` is real finite orientation data, but it does not improve the Gate-238 chirality sieve.

## Candidate 2 — scalar fundamental class `τ_η`

Gate 239 inherits the scalar fundamental-class signs:

```text
τ_η = (2, -2, 1)
-τ_η = (-2, 2, -1)
```

These are meaningful signed orientation data, but they are not currently an operator on `S_C`.

| Requirement | Status |
|---|---|
| Scalar-bundle functional | yes |
| Endomorphism of `S_C` | no |
| Canonical pullback to spinor chirality | no |
| Gauge projection map | no |
| Can act as `χ` | no |

So `τ_η` remains inherited orientation-trace data. It is **not** promoted into physical chirality.

## Six-plane weak sieve with `χ_vol`

Because `χ_vol` is equivalent to `γ`, every two-mode plane still has the same mixed chiral split:

| Plane class | Count | Doublet `χ+` | Doublet `χ-` | Singlet `χ+` | Singlet `χ-` |
|---|---:|---:|---:|---:|---:|
| temporal-spatial | 3 | 4 | 4 | 4 | 4 |
| pure-spatial | 3 | 4 | 4 | 4 | 4 |

Summary:

```text
candidate planes: 6
uniform χ-doublet planes: 0
uniform χ-singlet planes: 0
χ-selected planes: []
χ breaks degeneracy: false
```

Therefore the physical weak plane remains unselected.

## Theorem distinction

Gate 239 confirms:

```text
finite orientation preflight: yes
true SM chirality theorem: no
unique weak plane: no
global H summand: no
```

The correct statement is:

```text
The currently available orientation endomorphism on S_C is proportional to γ, and τ_η is not yet pulled back to S_C. Therefore Gate 239 does not derive physical chirality.
```

## Firewall ledger

Gate 239 does **not**:

```text
adjust χ signs to fit the weak plane
import γ₅ from continuum QFT
import Connes chirality
force a weak plane
promote τ_η into a spinor operator
claim left-handed weak action
claim global H
claim order-one calculus
```

## Binding obstruction

The next required theorem is one of:

```text
1. a nontrivial orientation pullback from τ_η/contact η data to S_C,
2. a faithful finite algebra representation whose order-one calculus defines physical chirality,
3. a contact-vacuum intertwiner selecting a weak plane independently of γ.
```

Until such a theorem exists, the weak plane and Standard Model chirality remain blocked.

## Validation

Passed:

```bash
go test -p=1 ./pkg/bridge/orientationtruechirality -count=1 -timeout=300s
go test -p=1 ./pkg/bridge/chiralweakselector ./pkg/bridge/orientationtruechirality -count=1 -timeout=300s
go list ./pkg/bridge/orientationtruechirality ./internal/app ./cmd/asha
```

The known full historical theorem-ladder tests were not run.
