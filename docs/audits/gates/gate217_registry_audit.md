# Gate 217 Registry Audit — Finite Spectral Triple / Heavy-Sector Gauge-Curvature Projection Audit

## Gate

```text
Gate 217 — finite spectral triple / heavy-sector gauge-curvature projection audit
Package: pkg/bridge/finitespectraltriple
Registry theorem: BRIDGE-FINITE-SPECTRAL-TRIPLE-HEAVY-SECTOR-GAUGE-CURVATURE-PROJECTION
```

## Epistemic status

```text
FAILED_ROUTE_FINITE_SPECTRAL_TRIPLE_MATCHING_DERIVATION
```

This is the correct strict result. Gate 217 does not disprove the Gate-215 single-scale spectrum. It proves that the current finite algebra still lacks the full spectral-action machinery required to derive the needed finite matching constants.

## Inherited target

Gate 217 inherits the unique Gate-215 plausible single-scale spectrum through Gate 216:

```text
Dirac (1,3,Y=1) + Dirac (8,2,Y=1/2)
M_B ≈ 2.60752425e6 GeV
M_* ≈ 1.71690311e17 GeV
```

The required matching residual remains:

```text
δ_match_required = (-0.000561193804, +0.000561440698, -0.000560508948)
normalized       = (-0.999560249, +1, -0.998340430)
sign pattern     = - + -
```

Gate 216 already proved that current finite scalars provide only diagnostics and near-misses, not a canonical matching row.

## Heavy-sector representation audit

The sealed test subject consists of two continuum threshold rows:

| Carrier | Representation | Internal dimension | Dirac chiral carrier dimension | Beta row |
|---|---:|---:|---:|---:|
| sealed Dirac weak triplet | `(1,3,Y=1)` | 3 | 6 | `(12/5, 8/3, 0)` |
| sealed Dirac color-octet weak doublet | `(8,2,Y=1/2)` | 16 | 32 | `(16/5, 16/3, 8)` |

These rows are valid under the `ThresholdSpectrumSeal`, but the finite core has not derived their heavy Hilbert carrier, inner product, real structure, grading, local field map, or finite mass.

## Finite Dirac operator audit

Gate 217 audits possible `D_F` candidates without promoting any of them:

| Candidate | Result | Why not promoted |
|---|---|---|
| `D_F = 0` on the heavy sector | Rejected | Vacuous; cannot produce `Tr(D_F^-2)`, decoupling, or threshold matching constants. |
| `D_F = M_B I` identity mass ansatz | Rejected | Imports the phenomenological PeV scale and chooses the operator by hand. |
| Off-diagonal `(1,3,1) ↔ (8,2,1/2)` map | Rejected | No gauge-equivariant finite intertwiner exists between inequivalent color/weak/hypercharge reps. |
| Reuse of earlier top-down Fock spectral-triple support | Rejected | It is a representation-trace certificate for the SM seed, not a heavy-sector Dirac operator. |

Summary:

```text
finite Dirac candidates audited: 4
promotable finite D_F:           0
order-one verified:              0
finite mass scale derived:        0
Clifford/G2 dictated D_F:         0
```

Binding obstruction:

```text
FINITE_DIRAC_OPERATOR_NOT_DERIVED
```

## Heat-kernel / gauge-curvature projection audit

A lawful matching derivation requires a gauge-fluctuated operator `D_A` and a heat-kernel projection onto:

```text
U(1)_Y curvature squared
SU(2)_L curvature squared
SU(3)_C curvature squared
```

Gate 217 finds:

```text
a2/a4 language audited:            true
finite spectral triple complete:   false
gauge fluctuation map derived:      false
representation trace rows known:    2
gauge projection rows derived:      0
a4 gauge coefficients derived:      0
projected δ_i^match rows:           0
```

Binding obstruction:

```text
GAUGE_CURVATURE_PROJECTION_NOT_DERIVED
```

Representation traces and beta rows are known, but they are not physical finite matching constants.

## Cutoff and subtraction scheme audit

Even if a formal heat-kernel trace existed, physical threshold constants require a cutoff function, cutoff moments, and a subtraction prescription.

Gate 217 finds:

```text
canonical cutoff function:      false
cutoff moments:                 false
renormalization scheme:         false
threshold subtraction rule:     false
MSbar imported:                 false
dimensional regularization:     false
finite counterterm functional:  false
physical δ_i^match rows:        0
```

Binding obstruction:

```text
CUTOFF_SUBTRACTION_SCHEME_NOT_DERIVED
```

No `MSbar`, dimensional regularization scheme, or cutoff function is imported as a finite theorem.

## Firewall audit

Gate 217 explicitly preserves:

```text
ThresholdSpectrumSeal: inherited
EmpiricalCarrierSeal: inherited
LeptoquarkDynamicsSeal: inherited
Z-pole ledger: quarantined
D_F fitted by hand: false
cutoff function invented: false
heat-kernel projection fitted: false
matching residual promoted: false
matching corrections derived: false
heavy masses finite-derived: false
physical unification claimed: false
contact modes promoted to particles: false
B-gap promoted to mass: false
proton lifetime computed: false
```

## Theorem statement

Gate 217 proves that the ASHA Engine can name the exact spectral-action machinery required to derive the Gate-215 residual, but cannot yet construct it from the finite core.

Therefore:

```text
δ_i^match remains an external target vector, not a finite-derived heat-kernel coefficient.
```

The next valid step is either:

1. construct a finite heavy-sector `D_F, J, gamma` theorem with an explicit order-one calculus, or
2. formalize a matching-correction seal / precision uncertainty ledger for future phenomenological work.

## Validation

Passed:

```bash
go test -p=1 ./pkg/bridge/finitespectraltriple -count=1 -timeout=300s
go test -p=1 ./pkg/bridge/matchingresidualstructure ./pkg/bridge/finitespectraltriple -count=1 -timeout=300s
go list ./pkg/bridge/finitespectraltriple ./internal/app ./cmd/asha
```

The known full historical registry compile/test path was not used for final validation because previous gates repeatedly showed it timing out through the full theorem ladder.
