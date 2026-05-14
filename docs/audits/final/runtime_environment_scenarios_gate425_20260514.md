# ASHA Runtime Environment Scenario Verification — Gate 425

## Purpose

This verification pass updates the standalone runtime so that environmental/dark/cosmology scenarios compute conditional numeric paths rather than only reporting the firewall.

The numbers remain bridge/environmental diagnostics. They are not promoted to native ASHA predictions.

## New runtime surface

```bash
go run ./cmd/asha --scenario dark-stable-thermal --format markdown --strict
go run ./cmd/asha --scenario cosmology --format markdown --strict
go run ./cmd/asha --scenario environment --format markdown --strict
```

## Dark-sector conditional arithmetic

Stable thermal B-gap Majorana stress test:

```math
M_B = 1.46774973718\times10^6\ {\rm GeV}
```

```math
Y_{\rm required}=2.97981078152\times10^{-16},
\qquad
Y_{\rm thermal}=3.90149836766\times10^{-3}
```

```math
\Omega_{\rm thermal}h^2=1.57117293159\times10^{12},
\qquad
\Omega_{\rm thermal}/\Omega_{\rm DM}=1.30931077633\times10^{13}
```

```math
Y_{\rm required}/Y_{\rm thermal}=7.63760612134\times10^{-14}
```

Verdict: stable thermal B-gap Majorana dark matter is rejected by overclosure. Suppressed/nonthermal/dilution/decay routes remain conditional history bridges.

## Cosmology conditional arithmetic

Bare CCM diagnostic convention:

```math
\rho_{\rm bare}/M_P^4=48/\pi^2=4.863416814832
```

Diagnostic target:

```math
\rho_\Lambda/M_P^4=10^{-120}
```

Cancellation severity:

```math
\rho_{\rm bare}/\rho_\Lambda=4.863416814832\times10^{120},
\qquad
\log_{10}(\rho_{\rm bare}/\rho_\Lambda)=120.686941492
```

Holographic/dilaton target scale:

```math
L M_P\sim10^{60}
```

Gate-344 target convention:

```math
L M_P\sim10^{61}
```

Electroweak-vacuum scale tension:

```math
(v_{\rm Pf}/M_P)^4=1.67936189445\times10^{-67}
```

```math
(v_{\rm Pf}/M_P)^4/10^{-120}=1.67936189445\times10^{53}
```

```math
(v_{\rm Pf}/M_P)^4/10^{-122}=1.67936189445\times10^{55}
```

Verdict: holographic/dilaton route is numerically viable as a bridge target, but no native saturation/subtraction theorem is derived.

## Vacuum-fate conditional stress test

ASHA threshold jump:

```math
\Delta\lambda_B=-0.097846792207
```

| Seed | λ before | λ after | μ_inst [GeV] | λ_min | S_E | log10 τ/yr |
|---|---:|---:|---:|---:|---:|---:|
| tree pole top | -0.006880640805 | -0.104727433012 | 5.76733268667e5 | -0.122793907974 | 214.334289900 | 55.642486182 |
| one-loop-QCD MSbar-like top | 0.032625459630 | -0.065221332577 | 1.46774973718e6 | -0.077446343073 | 339.834574820 | 109.740890393 |

Verdict: conditional phenomenology only; it requires empirical top/Higgs inputs and a chosen continuum RG scheme.

## Comparison to paper/final ledger

The paper and `docs/audits/final/final_result.md` were updated to include these conditional numerical paths with explicit warnings.

## Validation

```bash
go test -p=1 ./pkg/asha ./cmd/asha -count=1

go test -p=1 \
  ./pkg/asha \
  ./cmd/asha \
  ./pkg/phenomenology \
  ./pkg/bridge/publicationbundlepreflight \
  ./pkg/bridge/artifactindexexport \
  ./pkg/bridge/familyaxiomclosureledger \
  ./pkg/matter/yukawaintertwiner \
  ./pkg/matter/hypercharge \
  ./pkg/matter/su2l \
  -count=1

go list ./internal/app ./cmd/asha
```

All targeted checks passed.
