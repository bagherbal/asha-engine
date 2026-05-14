# Gate 280 Registry Audit

**Gate 280 — Resolvent Field Adjunction / Contact Projector Construction Audit**

## Purpose

Gate 279 proved the rational field barrier: the contact quartic companion module is irreducible over `Q`, so its rational commutant is the field `Q[C_q4]` and admits no non-trivial idempotents. Gate 280 audits the next mathematically legal move: activate a sealed resolvent-field adjunction and construct the conditional `2+2` projectors that exist only after adjoining a root of the resolvent cubic.

This gate is deliberately conditional. It shows what the adjunction *permits*; it does not pretend the finite core has selected a resolvent root.

## Input Ledger

Inherited from Gate 279:

```text
q4(x) = 3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271
```

Normalized companion module:

```text
Q[x]/(q4), basis {1, x, x^2, x^3}
```

Companion matrix:

```text
C_q4 = [[0, 0, 0, -271/3240],
        [1, 0, 0, 149/216],
        [0, 1, 0, -119/60],
        [0, 0, 1, 71/30]]
```

Resolvent cubic:

```text
R(z) = 5832000z^3 - 11566800z^2 + 7569900z - 1637467
```

Gate 279 certified:

```text
q4 irreducible over Q
R(z) irreducible over Q
no nontrivial rational idempotents
```

## Resolvent Adjunction Seal

Gate 280 formally activates:

```text
ResolventAdjunctionSeal
```

Interpretation:

```text
Q  ->  Q(z_res)
```

where `z_res` is one root of the irreducible resolvent cubic.

The seal grants permission to analyze conditional branches. It does **not** grant a native finite-core theorem selecting which branch is physical.

Status:

```text
CONDITIONAL_SUPPORT_RESOLVENT_ADJUNCTION_SEAL_ACTIVATED
```

## Conditional Branches

Gate 280 constructs the three possible `2+2` branches:

| Branch | Pairing | Resolvent root `z` | Meaning |
|---|---:|---:|---|
| `z_high_pairing_q1q2_q3q4` | `(q1,q2)|(q3,q4)` | `0.793092963834819` | Conditional high resolvent branch |
| `z_mid_pairing_q1q3_q2q4` | `(q1,q3)|(q2,q4)` | `0.607181256713348` | Conditional middle resolvent branch |
| `z_low_pairing_q1q4_q2q3` | `(q1,q4)|(q2,q3)` | `0.583059112785166` | Conditional low resolvent branch |

For each branch, the quartic factors into two quadratics of the form:

```text
(x^2 - s_A x + p_A)(x^2 - s_B x + p_B)
```

where:

```text
z = p_A + p_B
p_A p_B = 271/3240
```

Status:

```text
CONDITIONAL_SUPPORT_RESOLVENT_FIELD_BRANCHES_CONSTRUCTED
CONDITIONAL_SUPPORT_CONDITIONAL_QUADRATIC_FACTORIZATIONS_CONSTRUCTED
```

## Conditional Projector Construction

For each branch, Gate 280 constructs two polynomial projectors:

```text
P_A = p_A(C_q4)
P_B = p_B(C_q4)
```

such that, branch-by-branch:

```text
P_A^2 = P_A
P_B^2 = P_B
[P_A, C_q4] = 0
[P_B, C_q4] = 0
P_A + P_B = I
P_A P_B = 0
```

The residuals are numerically audited and below tolerance for all three branches.

Status:

```text
CONDITIONAL_SUPPORT_CONDITIONAL_CONTACT_PROJECTORS_CONSTRUCTED
CONDITIONAL_SUPPORT_PROJECTOR_ORTHOGONALITY_VERIFIED_PER_BRANCH
```

## Sector Bijection Firewall

Gate 277 selected the sector-level pairing:

```text
{u,d}|{e,ν}
```

Gate 280 constructs three possible algebraic contact pairings. But it does **not** derive a theorem mapping one of these projectors to `{u,d}` and the other to `{e,ν}`.

The remaining possible conditional maps are:

```text
3 resolvent branches × 2 projector orientations = 6 maps
```

Thus the projectors exist after adjunction, but their physical interpretation remains unassigned.

Failed route:

```text
FAILED_ROUTE_PROJECTORS_NOT_MAPPED_TO_PHYSICAL_SECTORS
```

## Gate 275 Branch Firewall

Gate 275 derived:

```text
r_+ = (3591 + 136√123)/3099 ≈ 1.645470463011191
r_- = (3591 - 136√123)/3099 ≈ 0.672051318208557
```

Gate 280 does not derive a map:

```text
resolvent branch -> r_+ or r_-
```

Therefore the amplitude branch remains unlocked.

Failed routes:

```text
FAILED_ROUTE_RESOLVENT_TO_RPLUS_RMINUS_BRANCH_MAP_MISSING
FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED
FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED
```

## Firewall Ledger

Gate 280 explicitly avoids:

- promoting an arbitrary resolvent root to a finite theorem;
- using numerical root ordering as semantics;
- inserting observed masses;
- inserting empirical Yukawa data;
- claiming a Higgs mass ratio;
- identifying projectors with physical sectors without a theorem.

Status:

```text
CONDITIONAL_SUPPORT_RESOLVENT_ADJUNCTION_FIREWALLS_PRESERVED
```

## Final Verdict

Gate 280 proves that the field-extension route is mathematically viable:

```text
Resolvent root adjunction -> conditional quadratic factorization -> conditional 2+2 contact projectors
```

But it does not prove that the finite core selects one branch.

Final status:

```text
CONDITIONAL_SUPPORT_RESOLVENT_ADJUNCTION_SEAL_ACTIVATED
CONDITIONAL_SUPPORT_CONDITIONAL_CONTACT_PROJECTORS_CONSTRUCTED
FAILED_ROUTE_NO_NATIVE_RESOLVENT_ROOT_SELECTOR_DERIVED
FAILED_ROUTE_PROJECTORS_NOT_MAPPED_TO_PHYSICAL_SECTORS
FAILED_ROUTE_RESOLVENT_TO_RPLUS_RMINUS_BRANCH_MAP_MISSING
FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED
FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED
```

## Next Gate Obligation

The next theorem must supply one of the following:

1. a native finite-core selector for one resolvent root;
2. a lawful projector-to-sector semantic map;
3. an explicit sealed branch value that quarantines the selection;
4. a derived map from the selected resolvent branch to `r_+` or `r_-`.

Recommended next gate:

```text
Gate 281 — Resolvent Branch Semantics / Projector-to-Sector Orientation Seal Audit
```
