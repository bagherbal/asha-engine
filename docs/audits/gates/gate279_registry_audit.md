# Gate 279 Registry Audit

## Gate

**Gate 279 — Contact Projector Action / Quartic Companion Module Semantics Audit**

## Purpose

Gate 278 proved that numerical ordering of the four quartic contact roots is not a lawful algebraic selector. Gate 279 therefore treats the quartic contact polynomial as a companion-module action over the native rational base and asks whether any finite-geometric operator derives a nontrivial idempotent projector that block-diagonalizes the contact root space into the physical `2+2` sector pairing.

## Inputs inherited from prior gates

| Source | Data | Status |
|---|---:|---|
| Gate 169 | Contact quartic / scalar-contact shape tower | finite/contact-derived |
| Gate 275 | `r = |y/x|² = (3591 ± 136√123)/3099` | two-branch finite shape constraint |
| Gate 277 | sector-level `{u,d}|{e,ν}` selected by `τ_eta + B_gap` | conditional support |
| Gate 278 | no root-to-sector bijection; quartic roots form a Galois orbit | strict firewall |

## Companion module construction

The contact quartic is

```text
q4(x) = 3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271.
```

The monic form is

```text
x^4 -(71/30)x^3 +(119/60)x^2 -(149/216)x + 271/3240.
```

The companion module is

```text
Q[x]/(q4), basis {1, x, x^2, x^3}.
```

Using the convention that the companion matrix multiplies by `x`, the matrix is

```text
C_q4 = [[0, 0, 0, -271/3240],
        [1, 0, 0, 149/216],
        [0, 1, 0, -119/60],
        [0, 0, 1, 71/30]]
```

with

```text
Tr(C_q4) = 71/30
Det(C_q4) = -271/3240.
```

## Irreducibility certificates

| Object | Modular witness | Result |
|---|---|---|
| Quartic `q4` | `x^4 + 3x^3 + 2x + 2` over `F_7` | irreducible mod 7, hence irreducible over `Q` |
| Resolvent cubic | `z^3 - 4z^2 - 4z - 2` over `F_11` | irreducible mod 11, hence irreducible over `Q` |

## Centralizer and idempotent ledger

Because `q4` is irreducible, the companion module is a field extension over `Q`:

```text
Q[C_q4] ≅ Q[x]/(q4).
```

Therefore:

```text
centralizer_Q(C_q4) = Q[C_q4]
dim_Q centralizer = 4
idempotents over Q = {0, 1}
```

So the gate derives the strict no-go:

```text
FAILED_ROUTE_NO_NONTRIVIAL_RATIONAL_COMPANION_IDEMPOTENT
FAILED_ROUTE_COMPANION_MODULE_DOES_NOT_BLOCK_DIAGONALIZE_OVER_Q
```

## Native finite action candidates

| Candidate | Source | Companion action verdict |
|---|---|---|
| `τ_eta=(2,-2,1)` | Gates 242/259 | 3-component topological tag; no native 4D companion-module action |
| `diag(1,3,3,3)` Morita multiplicity | Gate 273 | diagnostic only; does not commute with `C_q4` |
| `B_gap` | B-gap / Majorana branch | scalar/identity-like; commutes but cannot distinguish roots or pairs |

The Morita multiplicity diagnostic has positive commutator residual against the companion action, so multiplicity cannot be promoted into a contact projector.

## Resolvent obligation

The three possible `2+2` pairings remain:

```text
(q1,q2)|(q3,q4)
(q1,q3)|(q2,q4)
(q1,q4)|(q2,q3)
```

A projector onto one of these pairings requires adjoining/selecting a resolvent root. Gate 279 does not derive such an adjunction.

## Result statuses

```text
CONDITIONAL_SUPPORT_CONTACT_QUARTIC_COMPANION_MATRIX_CONSTRUCTED
CONDITIONAL_SUPPORT_QUARTIC_AND_RESOLVENT_IRREDUCIBILITY_CERTIFIED
CONDITIONAL_SUPPORT_COMPANION_COMMUTANT_AND_IDEMPOTENT_LEDGER_AUDITED
CONDITIONAL_SUPPORT_NATIVE_FINITE_GEOMETRY_ACTION_LIFTS_TESTED
CONDITIONAL_SUPPORT_RESOLVENT_ADJUNCTION_OBLIGATION_EXPLICIT
FAILED_ROUTE_NO_NATIVE_FINITE_GEOMETRY_ACTION_ON_COMPANION_MODULE
FAILED_ROUTE_NO_NONTRIVIAL_RATIONAL_COMPANION_IDEMPOTENT
FAILED_ROUTE_COMPANION_MODULE_DOES_NOT_BLOCK_DIAGONALIZE_OVER_Q
FAILED_ROUTE_RESOLVENT_ROOT_NOT_ADJOINED_OR_SELECTED
FAILED_ROUTE_ROOT_TO_YUKAWA_SECTOR_BIJECTION_STILL_MISSING
FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED
FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED
```

## Firewall ledger

Gate 279 explicitly avoids:

```text
numerical-ordering promotion
observed-mass insertion
empirical Yukawa insertion
arbitrary resolvent-root choice
aesthetic root pairing
B_gap-to-root-magnitude identification
Higgs-ratio claim
```

## Final verdict

Gate 279 proves that the contact quartic companion module admits no nontrivial rational commuting idempotent. The topological sector split `{u,d}|{e,ν}` remains supported, but the contact root pairing requires a new theorem: a native operator that selects/adjoins a resolvent root and produces a lawful projector on the companion module.

## Next gate obligation

A lawful next gate should not try to select `r_+` or `r_-` by magnitude. It should audit whether a **resolvent-field adjunction** can be treated as a sealed conditional construction, or whether another finite invariant supplies a native resolvent-root selector.
