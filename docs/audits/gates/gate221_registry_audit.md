# Gate 221 Registry Audit — Heavy-carrier decay and mass-splitting / cosmological-relic safety audit

## Gate identity

- **Gate:** 221
- **Package:** `pkg/bridge/heavycarrierdecayaudit`
- **Theorem:** `HeavyCarrierDecayRelicSafetyAuditTheorem`
- **Status:** `FAILED_ROUTE_COSMOLOGICAL_PATHOLOGY`
- **Seal outcome:** `RELIC_DECAY_SEAL_REQUIRED_NOT_GRANTED`

Gate 221 inherits the Gate-220 PeV-threshold observability audit. Gate 220 established that the sealed PeV spectrum is safe against direct-reach, EW precision, and Higgs-loop probes by decoupling, but it logged a stable-relic warning. Gate 221 asks whether the engine can legally make the heavy carriers decay or split their charged/neutral components without inventing new couplings.

## Inherited sealed spectrum

```text
Dirac electroweak triplet      (1,3,Y=1)
Dirac color-octet weak doublet (8,2,Y=1/2)
M_B ≈ 2.56895727e6 GeV
M_B 1σ range ≈ [2.46868509e6, 2.67089887e6] GeV
```

Active inherited seals:

```text
EmpiricalCarrierSeal
ThresholdSpectrumSeal
MatchingCorrectionSeal
LeptoquarkDynamicsSeal
```

These seals allow the spectrum to be audited phenomenologically, but they do not derive decay operators, heavy-light couplings, mass splittings, relic abundance, or dark-matter stability.

## Carrier inventory

| Carrier | Representation | Electric charges | Threat |
|---|---|---:|---|
| Dirac electroweak triplet | `(1,3,Y=1)` | `{0,1,2}` | neutral plus charged relic risk |
| Dirac color-octet weak doublet | `(8,2,Y=1/2)` | `{0,1}` | colored, neutral, and charged relic risk |

The triplet contains a neutral component, but Gate 221 does **not** declare dark matter. A neutral PeV relic without a decay or annihilation/relic calculation is a warning, not a candidate.

## Operator basis audit

Gate 221 audits five representative portal classes:

| Candidate | Dimension | Verdict |
|---|---:|---|
| Renormalizable triplet-lepton-Higgs portal | 4 | hypercharge/field semantics fail for sealed `Y=1` fermion triplet |
| Dimension-five triplet Higgs-lepton portal | 5 | no canonical finite operator, Lorentz contraction, coefficient, or suppression scale |
| Octet-doublet quark-Higgs portal | 4 | color-octet fermion cannot mix with SM triplet quarks without an additional colored operator not derived by ASHA |
| Dimension-six neutral-current decay portal | 6 | external EFT coefficient only; not finite-derived |
| Leptoquark-mediated colored-carrier decay | 6 | blocked by `LeptoquarkDynamicsSeal` |

Summary:

```text
candidates audited:       5
gauge-invariant templates: 1
finite-supported portals: 0
derived decay operators:   0
computable widths:         0
```

The one gauge-invariant-looking class is the leptoquark-mediated channel, but it is explicitly forbidden by the `LeptoquarkDynamicsSeal`; dormant `u(4)` slots cannot be used as propagators or operator coefficients.

## Mass-splitting audit

Gate 221 keeps the Gate-220 splitting proxy only as a diagnostic:

```text
ΔM_proxy = v² / M_B ≈ 0.0235987 GeV
m_π± ≈ 0.13957039 GeV
```

This proxy is below the charged-pion threshold and, more importantly, it is **not** a derived splitting theorem. The engine still lacks:

```text
charged-neutral splitting operator
electroweak loop splitting theorem for the sealed heavy sector
VEV-coupling splitting theorem
charged-to-neutral cascade rule
colored-state hadronization/decay rule
```

Therefore charged and colored relic risks remain active.

## BBN lifetime filter

The BBN safety threshold is treated only as a filter:

```text
τ < 1 second
Γ_required > ℏ / 1s ≈ 6.582119569e-25 GeV
```

Because no decay operator is derived, no decay width is legal. The lifetime is treated as unbounded/infinite for safety classification, so the route fails cosmological safety by operator absence.

Gate 221 does **not** use toy dimension-5 or dimension-6 widths to rescue the spectrum. Such formulas would require a coupling, Lorentz contraction, suppression scale, and operator coefficient that the engine has not supplied.

## RelicDecaySeal status

Gate 221 defines the required future seal:

```text
RelicDecaySeal
```

but does **not** grant it.

To grant this seal, a future gate must supply at least:

```text
gauge-invariant heavy-to-SM operator basis
Lorentz/local-field contraction
finite or sealed coupling coefficient
suppression scale
charged/neutral mass splitting
colored-state hadronization/decay channel
width Γ with τ < 1 second
```

Current operational status:

```text
RELIC_DECAY_SEAL_REQUIRED_NOT_GRANTED
```

## Final theorem

Gate 221 records:

```text
FAILED_ROUTE_COSMOLOGICAL_PATHOLOGY
```

This is not a failure of the RG bridge. It is a phenomenological obstruction: the PeV spectrum is precision-safe but not cosmologically safe until a decay/splitting sector is derived or explicitly sealed.

## Firewalls preserved

Gate 221 does **not** claim or compute:

```text
finite-derived decay operators
finite-derived mass splittings
heavy-light coupling constants
relic abundance
dark matter candidate status
BBN-safe lifetime from absent dynamics
proton lifetime
PeV mass as finite-core derivation
```

## Validation

Passed:

```bash
go test -p=1 ./pkg/bridge/heavycarrierdecayaudit -count=1 -timeout=300s
go test -p=1 ./pkg/bridge/pevobservabilityaudit -count=1 -timeout=300s
go list ./pkg/bridge/heavycarrierdecayaudit ./internal/app ./cmd/asha
```

A chained multi-package validation timed out in this environment after the predecessor package completed, so the checks were rerun separately instead of retrying the same timeout-prone path blindly.
