# ASHA Runtime Verification — 2026-05-14

This report reruns the standalone `pkg/asha` runtime board against the current Gate-425 project data and compares the outputs with the final result ledger and final manuscript.

## Commands executed

```bash
go run ./cmd/asha --scenario all --format text --strict
go run ./cmd/asha --scenario all --format json --strict
go run ./cmd/asha --scenario native --format text --strict
go run ./cmd/asha --scenario native --format json --strict
go run ./cmd/asha --scenario higgs --format text --strict
go run ./cmd/asha --scenario higgs --format json --strict
go run ./cmd/asha --scenario family --format text --strict
go run ./cmd/asha --scenario family --format json --strict
go run ./cmd/asha --scenario dark-stable-thermal --format text --strict
go run ./cmd/asha --scenario dark-stable-thermal --format json --strict
go run ./cmd/asha --scenario cosmology --format text --strict
go run ./cmd/asha --scenario cosmology --format json --strict
go run ./cmd/asha --scenario ci --format text --strict
go run ./cmd/asha --scenario ci --format json --strict
go test -p=1 ./pkg/asha ./cmd/asha -count=1
```

## Scenario results

- `all`: `PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data.`
- `native`: `PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data.`
- `higgs`: `PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data.`
- `family`: `PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data.`
- `dark-stable-thermal`: `PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data.`
- `cosmology`: `PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data.`
- `ci`: `PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data.`

## Key comparison table

| Check | Runtime value | Ledger/paper status | Result |
|---|---:|---|---|
| Clifford dimension | `256` | final_result includes 2^8 / 256 | PASS |
| Exterior grade vector | `[1 8 28 56 70 56 28 8 1]` | final_result includes [1,8,28,56,70,56,28,8,1] | PASS |
| Contact vacuum dimension | `7` | final_result includes dim K=7 | PASS |
| Hypercharge normalization | `5/3` | final_result includes k_Y=5/3 | PASS |
| Weak mixing boundary | `3/8` | final_result includes sin²=3/8 | PASS |
| Pfaffian scale | `247.151135557 GeV` | paper/final_result contain Pfaffian scale lane; exact numeric in runtime | PASS |
| Higgs tree proxy | `124.925370288 GeV` | paper/final_result include 124.925 proxy | PASS |
| Majorana B-gap | `1467749.73718 GeV` | final_result includes exact B-gap ledger; paper states the rounded scale near 1.46×10^6 GeV | PASS |
| Stable thermal overclosure | `1.3e+13` | final_result/runtime/paper include overclosure rejection | PASS |
| Native charged flavor firewall | `13` | paper/final_result include 13 charged moduli firewall | PASS |
| Family axiom symbolic ledger | `9` | paper/final_result include 9 symbolic charged coefficients | PASS |

## Conclusion

No runtime value changed relative to the final result ledger or the final manuscript framing. The paper does not require numerical updates from this verification pass.

The runtime board confirms the same final boundary:

```text
native ASHA law-space derives the audited geometric/gauge/Higgs scaffold
K/X/Y family axioms remain quarantined capacity, not native theorem
dim M_charged^native = 13
cosmology and dark-sector coordinates remain environmental/history-dependent
```