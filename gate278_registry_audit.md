# Gate 278 Registry Audit — Quartic Root-to-Yukawa Sector Bijection / Contact Projector Semantics Audit

## Gate boundary

Gate 277 selected the **sector-level** pairing

```text
{u,d}|{e,ν}
```

using the combined semantic/topological tags:

```text
τ_eta  -> binds the weak {u,d} pair
B_gap  -> tags the neutrino / Majorana side of the lepton pair
```

Gate 277 deliberately did **not** select a contact resolvent root, because the four quartic contact roots were still unlabeled. Gate 278 audits whether the internal arithmetic of those roots, together with the Gate-273 Morita multiplicities and the B-gap scale tag, supplies a native bijection

```text
{q1,q2,q3,q4} ↔ {u,d,e,ν}.
```

## Retrieved contact data

The quartic contact polynomial is carried forward as

```text
q4(x) = 3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271.
```

The isolated roots are audited as finite contact eigenvalues:

| root | interval | approx | closest simple rational | verdict |
|---|---:|---:|---:|---|
| q1 | [2839/10000, 2840/10000] | 0.2839121926 | 1/4 | finite O(1), not zero/suppressed |
| q2 | [4411/10000, 4412/10000] | 0.4411227573 | 1/2 | finite O(1), not zero/suppressed |
| q3 | [7440/10000, 7441/10000] | 0.7440966380 | 3/4 | finite O(1), not zero/suppressed |
| q4 | [8975/10000, 8976/10000] | 0.8975350788 | 1 | finite O(1), not zero/suppressed |

Gate 278 records that real ordering and numerical magnitude are useful diagnostics, but they are **not** invariant contact-projector semantics. The gate therefore refuses to assign “smallest root = neutrino” or any similar ordering rule.

## Constraint audit

| Constraint | Source | Native reach | Root-label reach | Verdict |
|---|---|---|---|---|
| Morita multiplicity `κ_C:κ_Q = 1:3` | Gate 273 | finite Hilbert trace sectors | no | Counts lepton/quark multiplicity, not individual quartic roots. |
| B-gap Majorana tag | Gates 229–231 / NeutrinoTextureSeal | neutrino-sector semantics once labels exist | no | Tags ν after a sector map exists; does not map a contact root to a suppressed scale. |
| `τ_eta` weak-doublet tag | Gates 242, 259, 277 | weak-sector `{u,d}` labels | no | Binds `{u,d}` semantically, but lacks action on contact-root projectors. |

## Projector semantics obstruction

Gate 278 records the algebraic reason the bijection remains blocked:

```text
quartic irreducible over Q:          true
resolvent cubic irreducible over Q:  true
individual root projectors over Q:   false
2+2 pair projectors over Q:          false
```

Individual root idempotents require the quartic splitting field. A 2+2 pair idempotent requires choosing or adjoining a resolvent root. Neither operation has been derived as a native finite-core projector action.

This is the precise meaning of the firewall:

```text
labels are not roots;
sector tags are not contact projectors;
numerical ordering is not a sector theorem.
```

## Resolvent pairing audit

The three formal contact pairings remain:

| branch | root pairing | compatibility with `{u,d}|{e,ν}` | selected? | verdict |
|---|---|---:|---:|---|
| R12_34 | `{q1,q2}|{q3,q4}` | compatible after arbitrary bijection | no | no intrinsic selector |
| R13_24 | `{q1,q3}|{q2,q4}` | compatible after arbitrary bijection | no | no intrinsic selector |
| R14_23 | `{q1,q4}|{q2,q3}` | compatible after arbitrary bijection | no | no intrinsic selector |

All three root pairings can represent the sector split if one is allowed to assign root labels externally. Gate 278 therefore logs that no contact resolvent root has been selected.

## Root-sector bijection count

Gate 278 audits the degeneracy rather than hiding it:

```text
total root-sector bijections:        24
after sector-level {u,d}|{e,ν}:      12
after B_gap neutrino semantic tag:    6
after τ_eta weak-pair semantic tag:   6
unique bijection:                     false
```

The semantic tags reduce the interpretation space, but do not produce a unique root assignment.

## Gate 275 branch projection

Gate 278 inherits the Gate-275 scalar-Morita branches:

```text
r_+ = (3591 + 136√123)/3099 ≈ 1.645470463011191
|y/x|_+ ≈ 1.282758926303454

r_- = (3591 - 136√123)/3099 ≈ 0.672051318208557
|y/x|_- ≈ 0.819787361581378
```

Because no contact resolvent root is selected, and because no theorem maps a selected contact pairing to `r_+` or `r_-`, the amplitude branch remains unlocked.

## Result status

```text
CONDITIONAL_SUPPORT_CONTACT_QUARTIC_ROOTS_RETRIEVED
CONDITIONAL_SUPPORT_CONTACT_ROOT_MAGNITUDE_SIEVE_COMPLETED
CONDITIONAL_SUPPORT_MORITA_AND_B_GAP_CONSTRAINTS_APPLIED_AS_AUDIT_TESTS
CONDITIONAL_SUPPORT_RESOLVENT_PAIRINGS_INHERITED
CONDITIONAL_SUPPORT_ROOT_BIJECTION_FIREWALLS_PRESERVED
FAILED_ROUTE_NO_CONTACT_ROOT_IS_NATIVE_NULL_OR_MAJORANA_SUPPRESSED
FAILED_ROUTE_1_PLUS_3_MULTIPLICITY_DOES_NOT_LABEL_INDIVIDUAL_ROOTS
FAILED_ROUTE_CONTACT_ROOT_PROJECTOR_SEMANTICS_NOT_DERIVED
FAILED_ROUTE_ROOT_TO_YUKAWA_SECTOR_BIJECTION_MISSING
FAILED_ROUTE_CONTACT_RESOLVENT_ROOT_NOT_SELECTED
FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED
FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED
```

## Firewall ledger

Gate 278 uses no observed masses, no CKM/PMNS data, and no empirical Yukawa matrices. It does not promote root ordering, magnitude, or aesthetic pairing into a finite theorem. It preserves Gate 277's sector-level support while keeping the contact-root and amplitude-branch results blocked.

## Theorem statement

Gate 278 confirms the root-semantics obstruction. The four contact roots are finite O(1) members of one irreducible quartic orbit. Morita `1+3` multiplicity counts trace sectors. B-gap tags the neutrino sector only after labels exist. `τ_eta` binds `{u,d}` only at sector level. None of these data provides a rational contact-root projector or a native `q_i ↔ {u,d,e,ν}` bijection.

Therefore the resolvent root, Gate-275 `r` branch, and Higgs ratio remain unclaimed.

## Next lawful target

The next gate should search for a true **contact projector action** or **quartic companion module semantics**:

```text
Gate 279 — Contact Projector Action / Quartic Companion Module Semantics Audit
```

Required future objects:

1. root or pair idempotents with physical sector semantics;
2. a finite operator action linking contact-root projectors to the Morita sector ledger;
3. a theorem mapping resolvent branches to Gate-275 `r_±` branches;
4. physical `J`/hypercharge completion;
5. heat-kernel field-normalization projection before any `a₂/a₄` or Higgs-ratio claim.
