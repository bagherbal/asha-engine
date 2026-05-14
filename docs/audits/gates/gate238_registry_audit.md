# Gate 238 Registry Audit — Chiral Alignment (`γ`) and Weak Plane Selector Audit

## Verdict

```text
CONDITIONAL_SUPPORT_GAMMA_PARITY_PREFLIGHT
CONDITIONAL_SUPPORT_TEMPORAL_SPATIAL_CLASS_SIEVE
FAILED_ROUTE_UNIFORM_CHIRAL_DOUBLET_ALIGNMENT
FAILED_ROUTE_CHIRAL_WEAK_PLANE_SELECTION
FAILED_ROUTE_LEFT_HANDED_WEAK_ACTION_DERIVATION
FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED
```

Gate 238 audits whether the native occupation-parity grading `γ=(-1)^N`, introduced in the finite Dirac scaffold, can select one of the six candidate two-mode weak planes found in Gate 237.

It cannot.

This does **not** invalidate Gate 237's local quaternionic support. It proves that raw Fock occupation parity is not yet the physical Standard Model chirality selector.

---

## Inherited context

Gate 237 established that for every two-mode plane `U⊂W`, with

```text
S_C = Λ*(W),     W = U ⊕ V,     dim_C U = 2,
```

the exterior representation decomposes as:

```text
Λ*(U) = 1 ⊕ 2 ⊕ 1
Λ*(W) = (1 ⊕ 2 ⊕ 1) ⊗ Λ*(V)
```

Thus every selected plane gives:

```text
8 complex doublet-state dimensions
8 complex singlet-state dimensions
```

and the fundamental doublet is pseudo-real, giving local quaternionic support.

Gate 237 failed only because no finite theorem selected the physical weak plane.

---

## Gamma grading audit

The native grading is:

```text
γ = (-1)^N
```

where `N` is total occupation number in the four-mode Fock carrier.

It splits the complexified spinor into:

| Sector | Complex dimension |
|---|---:|
| even occupation | `8` |
| odd occupation | `8` |

This is a valid finite grading, but Gate 238 does **not** equate it to physical Standard Model chirality.

---

## Six-plane degeneracy sieve

For a candidate plane `U={i,j}`, the doublet sector is the set of states with exactly one `U`-mode occupied. The complement `V` has degrees `0,1,2`, so the doublet sector inevitably contains both parities.

For every one of the six planes:

| Plane type | Doublet even | Doublet odd | Singlet even | Singlet odd | Chiral uniform? |
|---|---:|---:|---:|---:|---|
| any `U_{ij}` | `4` | `4` | `4` | `4` | no |

Therefore:

```text
uniform doublet planes = 0
uniform singlet planes = 0
γ-selected planes      = 0
```

The exterior `su(2)` lift preserves total occupation parity; it commutes with `γ`. It does not act only on one parity sector.

---

## Temporal/spatial split audit

The native `1⊕3` split distinguishes two classes of planes:

| Class | Count |
|---|---:|
| temporal-spatial `{e0,ei}` | `3` |
| pure spatial `{ei,ej}` | `3` |

This is a genuine class-level distinction, but it still leaves a threefold degeneracy in each class. It does not select one electroweak plane.

---

## Theorem distinction

Gate 238 records:

```text
local H support: inherited from Gate 237
raw γ parity selector: failed
physical left-handed weak action: not derived
global H summand: still unselected
```

The correct interpretation is:

```text
occupation parity is a grading,
but it is not yet the SM chirality operator.
```

---

## Firewalls preserved

Gate 238 does **not** claim:

```text
physical left-handed weak assignment
unique electroweak plane
global quaternionic H summand
faithful C⊕H⊕M₃(C) representation
opposite algebra action
order-one calculus
PMNS / Yukawa / mass texture derivation
```

It does not import Pauli matrices, Connes' algebra, or Standard Model chirality as proof.

---

## Next obstruction

The next selector cannot be raw occupation parity. A stronger finite theorem is required, likely involving one of:

```text
contact-vacuum orientation
η-source / scalar orientation seal pullback
physical chirality map distinct from Fock parity
intertwiner between contact su(2) and a selected plane U⊂W
opposite-algebra/order-one compatibility
```

Until such a selector exists, the weak quaternionic summand remains locally supported but globally underived.
