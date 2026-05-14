# ASHA Runtime Verification — Environment Scenario Extension

Runtime version: `gate425-runtime-env-scenarios-20260514`

## Result

The standalone runtime now computes conditional dark-sector, cosmology, and vacuum-fate scenarios with explicit warnings. These numbers agree with the updated final result ledger and the updated manuscript language.

## Commands

```bash
go run ./cmd/asha --scenario all --format markdown --strict
go run ./cmd/asha --scenario dark-stable-thermal --format markdown --strict
go run ./cmd/asha --scenario cosmology --format markdown --strict
go run ./cmd/asha --scenario environment --format markdown --strict
go test -p=1 ./pkg/asha ./cmd/asha -count=1
```

## Key verified values

| Quantity | Runtime value | Status |
|---|---:|---|
| `M_B` | `1.46774973718e6 GeV` | bridge scale |
| `Y_required` | `2.97981078152e-16` | conditional dark yield target |
| `Y_thermal` | `3.90149836766e-3` | stable thermal stress test |
| `Ω_thermal h²` | `1.57117293159e12` | rejected stable thermal scenario |
| `Ω_thermal/Ω_DM` | `1.30931077633e13` | overclosure |
| `Y_required/Y_thermal` | `7.63760612134e-14` | required suppression/dilution |
| `ρ_bare/M_P^4` | `4.863416814832` | diagnostic bare CCM convention |
| `ρ_bare/ρ_target` | `4.863416814832e120` | subtraction severity |
| `digits cancellation` | `120.686941492` | diagnostic severity |
| `L M_P` for `10^-120` | `1e60` | holographic/dilaton target |
| `L M_P` for `10^-122` | `1e61` | Gate-344 target convention |
| `(v_Pf/M_P)^4` | `1.67936189445e-67` | EW-vacuum scaling |
| pole-seed `log10 τ/yr` | `55.642486182` | conditional RG/bounce stress test |
| MSbar-like top seed `log10 τ/yr` | `109.740890393` | conditional RG/bounce stress test |

## Verdict

No native ASHA boundary changed. The runtime now exposes the conditional scenario numbers that were previously summarized verbally.

```text
PASS: conditional numerical paths computed; environmental firewall preserved.
```
