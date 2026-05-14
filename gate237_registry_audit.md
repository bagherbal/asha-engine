# Gate 237 Registry Audit — Explicit `su(2)` Spinor Lift / Quaternionic (`H`) Closure Audit

## Status

```text
CONDITIONAL_SUPPORT_CANDIDATE_WEDGE_SU2_LIFTS
CONDITIONAL_SUPPORT_DOUBLET_DIMENSION_MATCH_PREFLIGHT
CONDITIONAL_SUPPORT_PSEUDOREAL_DOUBLETS_LOCAL_H_PREFLIGHT
FAILED_ROUTE_CONTACT_SU2_TO_SC_NATIVE_LIFT_DERIVATION
FAILED_ROUTE_CANONICAL_WEAK_PLANE_SELECTION
FAILED_ROUTE_NATIVE_GLOBAL_QUATERNIONIC_H_SUMMAND_DERIVATION
FAILED_ROUTE_COMPLETED_CONNES_ALGEBRA_DERIVATION
```

Gate 237 continues the native finite-algebra derivation after Gate 236. Gate 236 found native support for

```text
C ⊕ M₃(C)
```

from the generator split

```text
W = C·e0 ⊕ C³_spatial.
```

The remaining missing Standard-Model-like summand is the weak quaternionic algebra `H`. Gate 237 audits whether the already-derived contact-preserving `su(2)` can be lifted onto the complexified spinor carrier

```text
S_C = Λ*(W),   dim_C(S_C)=16,   dim_R(S_C)=32.
```

## Inherited theorem state

From Gate 235:

```text
S_C = S ⊗_R C
Jψ = ψ*
Majorana bilinear capacity exists kinematically
```

From Gate 236:

```text
1⊕3 split available
C⊕M₃(C) preflight available
u(1) complex summand plausible
quaternionic H summand not derived
Connes algebra import blocked
```

## Exterior `su(2)` lift audit

For any two-mode complex plane

```text
U ⊂ W,   dim_C U = 2,
W = U ⊕ V,
dim_C V = 2,
```

there is a standard exterior lift of the fundamental `su(2)` action:

```text
Λ*(W) = Λ*(U) ⊗ Λ*(V)
Λ*(U) = Λ⁰(U) ⊕ Λ¹(U) ⊕ Λ²(U)
       = 1 ⊕ 2 ⊕ 1.
```

Since `dim_C Λ*(V)=4`, every selected two-mode plane gives:

```text
4 copies of the fundamental doublet
8 complex doublet-state dimensions
8 complex singlet-state dimensions
```

There are six possible two-mode planes in the four-mode carrier. Gate 237 audits all six; none is selected by hand.

## Doublet dimensional resonance

The eight-complex-dimensional doublet sector matches the one-generation Standard Model left-doublet dimension:

```text
Q_L: 3 colors × 2 weak states = 6
L_L: 1 lepton × 2 weak states = 2
Q_L ⊕ L_L = 8 complex states
```

Gate 237 records this as real structural support:

```text
CONDITIONAL_SUPPORT_DOUBLET_DIMENSION_MATCH_PREFLIGHT
```

This is not yet a physical weak-doublet theorem. Still missing:

```text
canonical weak-plane selector
hypercharge attachment
color/lepton assignment to doublet copies
physical chirality map
opposite algebra action
order-one calculus
```

## Quaternionic closure audit

The fundamental `su(2)` doublet is pseudo-real. Therefore, on each selected doublet factor, the representation supports a local quaternionic module structure.

The associative image for a selected plane has the schematic form:

```text
C_singlet ⊕ M₂(C)_doublet-image
```

with the real pseudo-real doublet factor carrying local `H` behavior.

However, Gate 237 does not derive a global `H` summand because:

1. the finite core does not identify the contact-preserving `su(2)` with one canonical two-plane `U`,
2. the doublet projection is not tied to hypercharge/color assignments,
3. the opposite algebra action is not derived,
4. the order-one condition is still unavailable.

Thus the correct result is:

```text
local pseudo-real / quaternionic support: yes
global native H summand: no
exact C⊕H⊕M₃(C): no
```

## Firewall ledger

Gate 237 does **not** import or claim:

```text
Pauli matrices as an answer
Connes algebra C⊕H⊕M₃(C)
Standard Model weak-doublet assignment
manual weak-plane selection
hypercharge forcing
B-gap mass promotion
order-one calculus
Majorana mass theorem
PMNS/Yukawa data
```

## Theorem statement

Gate 237 proves that the complexified exterior spinor carrier has enough native representation capacity for weak doublets: every two-mode plane supplies pseudo-real `su(2)` doublets of the correct dimension. This is the first genuine quaternionic preflight.

But the gate also proves that the full weak algebra summand is still not derived. The finite geometry must next supply a canonical selector/intertwiner identifying the contact-preserving `su(2)` with one specific two-mode plane in `S_C`. Until then, the full finite algebra

```text
C ⊕ H ⊕ M₃(C)
```

remains supported in pieces but not derived as a single faithful spectral-triple algebra.
