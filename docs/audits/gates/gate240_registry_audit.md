# Gate 240 Registry Audit — Spin^c Twisted Chirality / Hypercharge-Weak Plane Sieve Audit

## Theorem ID

```text
BRIDGE-SPINC-TWISTED-CHIRALITY-HYPERCHARGE-WEAK-SIEVE
```

## Package

```text
pkg/bridge/spinctwistedchirality
```

## Inherited state

Gate 238 proved that occupation parity

```text
γ = (-1)^N
```

splits the complexified Fock carrier into `8⊕8` but does not align weak doublets into a single chiral sector. Gate 239 proved that the Clifford-volume orientation candidate is proportional to the same `γ`, and that `τ_η` is a scalar trace functional rather than a spinor endomorphism.

Therefore Gate 240 tests the next strictly native possibility: a Spin^c-like twist using the diagonal finite `u(1)` charge bookkeeping.

## Native u(1) generator

The audit uses the existing Fock charge seed:

```text
Y_native(|n⟩) = Σ_i w_i n_i
w = (-1, 1/3, 1/3, 1/3)
```

This is treated as native diagonal `u(1)` bookkeeping only. It is **not** promoted to full Standard Model hypercharge.

## Twisted chirality diagnostic

The tested diagnostic is:

```text
χ_twist = γ · Y_native
```

Properties:

```text
diagonal on S_C: yes
distinct from γ: yes
commutes with γ: yes
involution: no
physical chirality theorem: no
```

Because `χ_twist` is not an involution and has more than two eigenvalue classes, it cannot be declared to be Standard Model chirality.

## Six-plane sieve

For a candidate two-mode weak plane `U={i,j}`, the exterior `su(2)` preserves the diagonal `u(1)` only if the two mode weights match.

Result:

```text
U={a†_0,a†_1}: rejected, temporal-spatial, weights=(-1,1/3)
U={a†_0,a†_2}: rejected, temporal-spatial, weights=(-1,1/3)
U={a†_0,a†_3}: rejected, temporal-spatial, weights=(-1,1/3)

U={a†_1,a†_2}: survives, pure-spatial, weights=(1/3,1/3)
U={a†_1,a†_3}: survives, pure-spatial, weights=(1/3,1/3)
U={a†_2,a†_3}: survives, pure-spatial, weights=(1/3,1/3)
```

Thus the twist improves the sieve:

```text
six planes → three pure-spatial planes
```

but does not select a unique plane.

## Twisted doublet alignment

For every surviving pure-spatial plane, the weak doublet sector has multiple `χ_twist` eigenvalues. Therefore no plane has uniform twisted chirality.

```text
uniform twisted doublet planes: 0
selected planes: 0
```

## Status

```text
CONDITIONAL_SUPPORT_NATIVE_U1_DIAGONAL_GENERATOR_PREFLIGHT
CONDITIONAL_SUPPORT_SPINC_GAMMA_U1_TWIST_PREFLIGHT
CONDITIONAL_SUPPORT_U1_COMMUTANT_TEMPORAL_SPATIAL_CLASS_SIEVE
FAILED_ROUTE_UNIFORM_TWISTED_CHIRALITY_ALIGNMENT
FAILED_ROUTE_SPINC_WEAK_PLANE_SELECTION
FAILED_ROUTE_SPINC_PHYSICAL_CHIRALITY_DERIVATION
FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED
```

## Firewall statement

Gate 240 does not import Standard Model hypercharge, does not tune the diagonal weights, does not force a weak plane, does not claim a left-handed weak action, and does not complete the global quaternionic `H` summand.

## Next obstruction

The remaining obstruction is the pure-spatial `S_3` degeneracy. A future gate must derive a selector that distinguishes one of the three spatial planes, or prove that an additional sealed symmetry-breaking datum is required.
