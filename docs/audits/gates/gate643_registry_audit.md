# Gate 643 — CompactSplit ResidualTensor BlockStructure Audit

## Purpose

Gate 642 compressed the compact/split obstruction into the internal projective angle

```text
cos(theta_twist)=13/sqrt(217),
sin(theta_twist)=4*sqrt(3)/sqrt(217).
```

Gate 643 stops treating this as only a scalar angle.  It constructs the projectively normalized residual tensor

```text
B_hat = B_K / ||B_K||_F,
G_hat = g_twist / ||g_twist||_F,
R_hat = [G_hat - <G_hat,B_hat>_F B_hat] / rho_twist,
```

for the three repeated Gate638/Gate639 routes:

```text
omega_1_alt,
omega_2_alt,
omega_B_alt.
```

This is an internal finite-geometry obstruction audit only.  It does not derive split-G2 structure, physical spacetime, boundary stress, scalar/flavor transport, Higgs mass, CKM/PMNS, gauge unification, or a native `7/72` theorem.

## Inherited angle

Gate 643 inherits from Gate 642:

```text
<g_twist,B_K>_F^2 : ||g_twist||_F^2||B_K||_F^2 = 169 : 217,
failure^2         : ||g_twist||_F^2||B_K||_F^2 = 48  : 217.
```

Equivalently:

```text
cos(theta_twist)=13/sqrt(217),
rho_twist=sin(theta_twist)=4*sqrt(3)/sqrt(217).
```

## Residual tensor certificate

For each route, Gate 643 certifies:

```text
<R_hat,B_hat>_F ≈ 0,
||R_hat||_F ≈ 1,
cos(theta) ≈ 13/sqrt(217),
rho_twist ≈ 4*sqrt(3)/sqrt(217).
```

The maximum audited drifts are numerical tolerance effects:

```text
max |<R_hat,B_hat>| ≈ 1.53e-16,
max | ||R_hat||_F - 1 | = 0,
max cosine drift ≈ 7.77e-16,
max rho drift ≈ 1.33e-15.
```

## Hodge-polarity block decomposition

Let `Q_+` and `Q_-` be orthonormal bases for the Gate634 Hodge sectors:

```text
K_7 = K_7^+ ⊕ K_7^-,
dim K_7^+ = 4,
dim K_7^- = 3.
```

Gate 643 computes:

```text
R_++ = Q_+^T R_hat Q_+,
R_-- = Q_-^T R_hat Q_-,
R_+- = Q_+^T R_hat Q_-.
```

Across all three routes, the block profile repeats:

```text
||R_++||_F^2      = 3/7,
||R_--||_F^2      = 4/7,
2||R_+-||_F^2     = 0,
rank(R_++)        = 4,
rank(R_--)        = 3,
rank(R_+-)        = 0.
```

Representative eigenvalue structure:

```text
R_++ eigenvalues ≈ -sqrt(3/28) repeated 4 times,
R_-- eigenvalues ≈ -sqrt(4/21) repeated 3 times.
```

Therefore the residual is **same-sector Hodge-diagonal**, not off-sector.

## Consequence for the Gate642 skeleton

Gate 642 conditionally organized the scalar obstruction as

```text
48 = p^2 q,
p = dim(K_7^+) = 4,
q = dim(K_7^-) = 3.
```

Gate 643 shows that the actual residual tensor does not place the failure in the off-sector block `K_7^+ × K_7^-`.  Instead, the repeated finite matrix result is:

```text
R_hat is diagonal by Hodge sector,
with norm split 3/7 on K_7^+ and 4/7 on K_7^-.
```

Thus `48=p^2q` remains a compressed scalar obstruction skeleton, not a certified off-sector block-source theorem.

## Verdict

```text
PASS_GATE642_PROJECTIVE_ANGLE_INHERITED
PASS_RESIDUAL_TENSOR_DEFINED_ORTHOGONAL_TO_BK
PASS_HODGE_POLARITY_BLOCKS_COMPUTED
PASS_ROUTE_BLOCK_PROFILES_COMPUTED
CONDITIONAL_SUPPORT_RESIDUAL_HAS_TYPED_BLOCK_STRUCTURE
CONDITIONAL_SUPPORT_RESIDUAL_IS_SAME_SECTOR_HODGE_DIAGONAL
FAILED_ROUTE_OFF_SECTOR_BLOCK_DOES_NOT_CARRY_RESIDUAL_TENSOR
FAILED_ROUTE_NO_NATIVE_TRACE_IDENTITY_FOR_PROJECTIVE_ANGLE_YET
FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM
FAILED_ROUTE_PROJECTIVE_ANGLE_IS_NOT_PHYSICAL_ANGLE
FAILED_ROUTE_PROJECTIVE_ANGLE_IS_NOT_PHYSICAL_METRIC_THEOREM
FIREWALL_PRESERVED_GATE643_RESIDUAL_TENSOR_IS_INTERNAL_OBSTRUCTION_ONLY
```

## Final classification

Gate 643 upgrades the scalar angle obstruction into an actual residual-tensor audit.  The residual tensor is clean and typed by the Gate634 Hodge polarity, but it is not the expected off-sector carrier.  Its repeated profile is same-sector and diagonal:

```text
R_hat ≈ -sqrt(3/28) P_{K7+} - sqrt(4/21) P_{K7-}
```

in the Hodge-polarity frame, up to route orientation conventions.

This is a real structural result, but not yet a theorem for the projective angle.  The missing theorem remains a native trace/projector identity deriving the `169:48:217` pair from the actual residual blocks.  All split-G2, boundary-stress, scalar/flavor, physical-geometry, and native `7/72` firewalls remain intact.
