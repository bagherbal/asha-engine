# Gate 277 Registry Audit — Resolvent Cubic Selector / B-Gap and Tau-Eta Symmetry Breaking

## Gate ID

`GATE277-RESOLVENT-CUBIC-SELECTOR-BGAP-TAUETA-SYMMETRY-BREAKING-AUDIT`

## Purpose

Gate 277 tests whether the topological tags available from the finite core can resolve the shared branch obstruction behind:

```text
Gate 186/187: quartic contact 2+2 resolvent ambiguity
Gate 275/276: scalar-Morita r_+ vs r_- amplitude ambiguity
```

The audit applies two native/sealed topological tags:

```text
τ_eta = (2,-2,1)    → weak-pair / generation selector
B_gap               → neutrino / Majorana intermediate-scale diagnostic
```

## Retrieved contact-resolvent data

The quartic contact factor is inherited as:

```text
q4(x) = 3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271
```

with isolated roots:

```text
q1 ≈ 0.2839121926
q2 ≈ 0.4411227573
q3 ≈ 0.7440966380
q4 ≈ 0.8975350788
```

The resolvent cubic is:

```text
5832000z^3 - 11566800z^2 + 7569900z - 1637467
```

It encodes the three unordered `2+2` pairings:

```text
{q1,q2}|{q3,q4}
{q1,q3}|{q2,q4}
{q1,q4}|{q2,q3}
```

## Topological tag audit

The tags are applied only at sector-label level:

```text
τ_eta binds the weak pair {u,d}
B_gap tags the neutrino/Majorana sector ν
```

This gives a unique surviving sector pairing:

```text
{u,d}|{e,ν}
```

and rejects:

```text
{u,e}|{d,ν}
{u,ν}|{d,e}
```

## Critical firewall

The selected sector pairing is not yet a selected contact resolvent root.

The missing theorem is a native bijection:

```text
{q1,q2,q3,q4} ↔ {u,d,e,ν}
```

without using empirical masses, CKM/PMNS data, or arbitrary root ordering.

Because this bijection is missing, Gate 277 cannot lawfully say which contact branch is selected. It also cannot map the sector pairing to the Gate-275 amplitude branches:

```text
r_+ = (3591 + 136√123)/3099
r_- = (3591 - 136√123)/3099
```

## Final status ledger

```text
CONDITIONAL_SUPPORT_RESOLVENT_CUBIC_RETRIEVED
CONDITIONAL_SUPPORT_TAU_ETA_AND_B_GAP_TAGS_APPLIED
CONDITIONAL_SUPPORT_SECTOR_LEVEL_UD_ENU_PAIRING_SELECTED
CONDITIONAL_SUPPORT_GALOIS_ORBIT_SIEVE_COMPLETED
CONDITIONAL_SUPPORT_GATE275_AMPLITUDE_BRANCHES_INHERITED
CONDITIONAL_SUPPORT_RESOLVENT_SELECTOR_FIREWALLS_PRESERVED
FAILED_ROUTE_QUARTIC_ROOT_TO_YUKAWA_SECTOR_BIJECTION_MISSING
FAILED_ROUTE_CONTACT_RESOLVENT_ROOT_NOT_SELECTED
FAILED_ROUTE_RESOLVENT_TO_RPLUS_RMINUS_BRANCH_MAP_MISSING
FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED
FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED
```

## Interpretation

Gate 277 is a real advance, but not the final branch selector.

It proves:

```text
τ_eta + B_gap select the sector-level physical 2+2 pairing {u,d}|{e,ν}.
```

It does not prove:

```text
which quartic contact roots represent u,d,e,ν;
which resolvent cubic root is physically selected;
which Gate-275 branch r_+ or r_- is selected;
any Seeley-de Witt a₂/a₄ or Higgs mass ratio.
```

## Future theorem obligations

A future gate must derive:

1. A contact-root to Yukawa-sector bijection.
2. Branchwise contact projectors/idempotents for the selected sector pairing.
3. A map from the selected resolvent root to `r_+` or `r_-`.
4. Physical `J` and chiral/hypercharge completion.
5. Heat-kernel projection and field normalization before any Higgs-ratio claim.

Recommended next gate:

```text
Gate 278 — Quartic Root-to-Yukawa Sector Bijection / Contact Projector Semantics Audit
```

## Tests

Focused tests passed:

```bash
go test -p=1 ./pkg/bridge/resolventcubictagselector -count=1 -timeout=120s -v

go test -p=1 ./pkg/bridge/resolventcubictagselector ./pkg/bridge/scalarmoritaspectralbridge -count=1 -timeout=120s -v

go list ./internal/app

go list ./cmd/asha
```

No full internal tests, full package tests, or `go test ./...` were run.
