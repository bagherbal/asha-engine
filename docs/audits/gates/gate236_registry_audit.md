# Gate 236 Registry Audit — Native Finite Algebra Derivation / Contact-Preserving Subalgebra Search

## Verdict

```text
CONDITIONAL_SUPPORT_NATIVE_1PLUS3_SPLIT_PREFLIGHT
CONDITIONAL_SUPPORT_MODE_COMMUTANT_C_PLUS_M3C_PREFLIGHT
CONDITIONAL_SUPPORT_U1_COMPLEX_SUMMAND_PREFLIGHT
FAILED_ROUTE_NATIVE_QUATERNIONIC_H_DERIVATION
FAILED_ROUTE_EXACT_CONNES_ALGEBRA_DERIVATION
FAILED_ROUTE_NATIVE_ALGEBRA_ORDER_ONE_READINESS
FAILED_ROUTE_FULL_NATIVE_FINITE_ALGEBRA_DERIVATION
```

Gate 236 asks whether the finite algebra of the spectral triple can be derived natively from the complexified `Cℓ(1,7)` carrier, the temporal/spatial `1⊕3` split, and the contact-preserving `su(2)⊕u(1)` structure.

It finds a real partial result: the native generator split supports a mode-level commutant

```text
End(C) ⊕ End(C³) = C ⊕ M₃(C)
```

This is the correct algebraic shape for singlet plus color-triplet bookkeeping. However, this is not yet the full Standard Model finite algebra. The quaternionic `H` summand, faithful representation on `S_C`, opposite algebra action, and order-one calculus remain un-derived.

---

## 1. Inherited Gate-235 Carrier

Gate 236 inherits the complexified carrier:

```text
S_C = S ⊗_R C

dim_R(S)   = 16
dim_C(S_C) = 16
dim_R(S_C) = 32
```

The candidate anti-linear structure remains:

```text
J ψ = ψ*
J² = +1
```

Gate 236 does not reinterpret this as a complete physical charge-conjugation theorem. It uses it only as the doubled carrier stage for the native algebra search.

---

## 2. Native `1⊕3` Split Audit

The Fock generator carrier has four covariant modes:

```text
a†_0 : temporal / lepton-like seed
a†_1 : spatial / color-like seed
a†_2 : spatial / color-like seed
a†_3 : spatial / color-like seed
```

Thus the native mode carrier is:

```text
W = C·e0 ⊕ C³_spatial
```

with projectors:

```text
P_lepton(e0) = e0
P_color(ei)  = ei, i=1,2,3
```

Result:

```text
1⊕3 mode split: supported
full particle-species projection: not derived
```

This distinction matters. The `1⊕3` split is native bookkeeping on the generator carrier, not yet a full Standard Model particle representation theorem.

---

## 3. Commutant / Subalgebra Search

The block-preserving associative algebra on the generator carrier is:

```text
End(C) ⊕ End(C³)
```

therefore:

```text
C ⊕ M₃(C)
```

Dimension check:

```text
dim_C(C)      = 1
dim_C(M₃(C))  = 9
total dim_C   = 10
```

This gives conditional support for the color matrix algebra preflight.

However, Gate 236 does not claim that this is the physical color gauge algebra on the full spinor because the following are missing:

```text
explicit lifted matrices on Λ*W = S_C
faithful doubled-space representation
opposite algebra action
non-vacuous order-one calculus
curvature/gauge action data
```

So the result is:

```text
CONDITIONAL_SUPPORT_MODE_COMMUTANT_C_PLUS_M3C_PREFLIGHT
```

not a full derivation of `SU(3)_C` dynamics.

---

## 4. `su(2)⊕u(1)` Integration Check

Gate 236 inherits the contact-preserving Lie algebra:

```text
su(2) ⊕ u(1)
```

The complexification makes a `C`-summand plausible for the `u(1)` phase preflight.

But the quaternionic summand requires more than the existence of a Lie algebra `su(2)`. The engine must derive:

```text
explicit su(2) matrices on S_C
left quaternionic module structure
doublet projection
associative closure equivalent to H
compatibility with J and γ
```

Current status:

```text
u(1) → C preflight: supported
su(2) → H: not derived
```

Therefore:

```text
FAILED_ROUTE_NATIVE_QUATERNIONIC_H_DERIVATION
```

---

## 5. Connes Algebra Import Block

Gate 236 explicitly does not import:

```text
C ⊕ H ⊕ M₃(C)
```

What it has natively is weaker:

```text
C ⊕ M₃(C) mode-level commutant preflight
u(1) complex summand preflight
su(2) Lie algebra availability
```

What it still lacks:

```text
H summand
faithful representation on S_C
opposite algebra action
order-one calculus readiness
Majorana sieve readiness
```

So the exact Standard Model finite algebra remains obstructed:

```text
FAILED_ROUTE_EXACT_CONNES_ALGEBRA_DERIVATION
FAILED_ROUTE_FULL_NATIVE_FINITE_ALGEBRA_DERIVATION
```

---

## 6. Firewalls

Gate 236 does not insert:

```text
Connes algebra by hand
SM gauge group by hand
explicit imported gauge matrices
Yukawa matrices
continuum masses
B-gap Majorana mass
```

It also does not claim:

```text
finite-derived Standard Model algebra
finite-derived quaternionic H module
finite-derived order-one condition
finite-derived B-gap neutrino placement
```

---

## 7. Architectural Status

Gate 235 ended with:

```text
complexified carrier available
native finite algebra representation missing
```

Gate 236 upgrades this to:

```text
native 1⊕3 mode commutant supports C⊕M₃(C)
quaternionic H and full Connes algebra still missing
```

The next hard target is narrow and precise:

```text
derive explicit contact-preserving su(2) matrices on S_C and test whether their associative closure forces a left quaternionic module.
```

Only after that can the order-one condition become non-vacuous and the B-gap Majorana sieve be revisited.
