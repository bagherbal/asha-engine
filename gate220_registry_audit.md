# Gate 220 Registry Audit — PeV-threshold indirect-signature / experimental observability audit

## Gate identity

```text
Gate 220 — PeV-threshold indirect-signature / experimental observability audit
Package: pkg/bridge/pevobservabilityaudit
Theorem: BRIDGE-PEV-THRESHOLD-INDIRECT-SIGNATURE-OBSERVABILITY
Status: CONDITIONAL_PHENOMENOLOGY_WITH_STABLE_RELIC_WARNING
```

Gate 220 inherits the Gate-219 sealed precision result. It does not derive a physical mass scale from the finite core. It audits whether the already sealed PeV threshold spectrum has immediate indirect phenomenological failures.

## Inherited sealed spectrum

```text
Dirac electroweak triplet:        (1,3,Y=1)
Dirac color-octet weak doublet:   (8,2,Y=1/2)
M_B central:                      2.56895727e6 GeV
M_B 1σ envelope:                  [2.46868509e6, 2.67089887e6] GeV
M_* central:                      1.72179441e17 GeV
M_* 1σ envelope:                  [1.66008302e17, 1.78344443e17] GeV
Worst Gate-219 residual/epsilon:  0.411919
```

Active seals:

```text
ThresholdSpectrumSeal:       active
MatchingCorrectionSeal:      active
EmpiricalCarrierSeal:        inherited
LeptoquarkDynamicsSeal:      inherited
```

## 1. Direct reach audit

Gate 220 uses a conservative 100 TeV proxy for future direct production reach.

```text
M_B / 100 TeV = 25.6896
```

Verdict:

```text
DIRECT_PRODUCTION_PARAMETRICALLY_OUT_OF_REACH
```

This is not a detector-simulation claim. It only records that the sealed PeV mass is far above direct production at current or 100 TeV-class machines.

## 2. Electroweak precision / oblique audit

The audit uses the parametric VEV-suppressed splitting proxy

```text
ΔM ≈ v² / M_B
```

and the small-splitting oblique proxy

```text
T_proxy ≈ (ΔM)² / (12π s_W² c_W² M_Z²)
```

Numerical result:

```text
v/M_B      = 9.58441982e-5
(v/M_B)²   = 9.18611034e-9
ΔM_proxy   = 0.0235987249 GeV
T_proxy    = 9.99417417e-9
```

Verdict:

```text
EWPO_PARAMETRICALLY_SAFE_UNDER_CURRENT_SEALS
```

Firewall: no tree-level violation, exact mass splitting, or heavy Yukawa coupling is derived. The result is a decoupling estimate only.

## 3. Higgs-loop imprint audit

The colored octet doublet can in principle contribute to gluon fusion, and the charged triplet/doublet states can in principle contribute to `h -> gamma gamma`. Under the current seals, however, the heavy mass is not derived from the Higgs VEV and no heavy Higgs Yukawa coupling is derived.

The audit therefore applies the standard decoupling proxy:

```text
v²/M_B² = 9.18611034e-9
```

With simple representation-multiplicity proxies:

```text
h -> gamma gamma amplitude proxy = 3.21513862e-7
gg -> h amplitude proxy          = 1.46977765e-7
```

Verdict:

```text
HIGGS_LOOP_IMPRINTS_DECOUPLE_UNDER_CURRENT_SEALS
```

Firewall: no non-decoupling Higgs coupling is invented.

## 4. Cosmological safety / dark matter audit

The `(1,3,Y=1)` Dirac triplet has charge components

```text
Q = T3 + Y = {0, 1, 2}
```

so it contains a neutral component. The `(8,2,Y=1/2)` sector is colored.

Missing semantics:

```text
decay operator:       not derived
mass splitting:       not derived
relic abundance:      not computed
DM candidate claim:   false
```

Warnings:

```text
stable neutral relic warning:  true
stable charged relic warning:  true
stable colored relic warning:  true
```

Verdict:

```text
COSMOLOGY_WARNING_DECAY_OR_RELIC_SEAL_REQUIRED
```

This is the binding Gate-220 warning. The PeV scale suppresses precision observables, but it does not by itself guarantee cosmological safety. The next gate must derive or seal heavy-carrier decay operators and charged/neutral splittings before relic abundance can be computed.

## Final theorem statement

Gate 220 finds:

```text
precision/direct observability: parametrically safe
Higgs-loop observability:       parametrically decoupled
cosmology/relic safety:          unresolved warning
```

Final status:

```text
CONDITIONAL_PHENOMENOLOGY_WITH_STABLE_RELIC_WARNING
```

## Firewalls preserved

Gate 220 does not claim:

```text
finite-derived PeV mass
observed collider signal
derived heavy decay operator
derived heavy Higgs Yukawa coupling
derived mass splitting
dark matter candidate
relic abundance
proton lifetime
contact/B-sector particle promotion
```

## Validation

Passed:

```bash
go test -p=1 ./pkg/bridge/pevobservabilityaudit -count=1 -timeout=300s
go test -p=1 ./pkg/bridge/inputsensitivityaudit -count=1 -timeout=300s
go list ./pkg/bridge/pevobservabilityaudit ./internal/app ./cmd/asha
```

A combined two-package test command timed out in this environment even though both packages passed when run separately; it was not retried blindly.
