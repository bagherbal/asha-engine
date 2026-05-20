# Gate 560 Registry Audit — Pauli-Hopf Scalar Moment Map Audit

## Verdict

`CONDITIONAL_SUPPORT_SEALED_PAULI_HOPF_SCALAR_MOMENT_MAP_FOUND`

Gate 560 changes the route. Gates 558 and 559 proved that the eta-record algebra

```text
A_eta_rec = span{I_HPhi, eta} ≅ R ⊕ R
```

splits the sealed scalar carrier as `4=2+2`, but cannot transfer linearly and trace-preservingly to a 3-dimensional spatial or generation carrier. Gate 560 therefore audits a different, lawful object already present in the sealed scalar lane: the Pauli/Hopf quadratic moment map on

```text
H_phi = R^4 ≅ C^2,    phi=(z1,z2),    z1=a+ib, z2=c+id.
```

It finds a sealed scalar-sector `4 -> 1+3` decomposition into scalar radius plus Pauli moment triplet, and for nonzero moment vector a scalar-sector `3 -> 1+2` orbit/stabilizer split. This is **not** a weak-plane, generation, gauge-boson, Yukawa, CKM/PMNS, or observed-flavor theorem.

## 1. Scalar Complex Structure

The carrier is certified under the existing sealed scalar-bundle orientation:

```text
H_phi = R^4 with basis (Re z1, Im z1, Re z2, Im z2)
      ≅ C^2 with phi=(z1,z2)
I_HPhi = I_4
```

Status:

```text
CONDITIONAL_SUPPORT_SEALED_HPHI_COMPLEX_STRUCTURE_CERTIFIED
```

## 2. Real Pauli Matrices

Gate 560 constructs the real symmetric matrices

```text
Sigma_1 = [[0,0,1,0],
           [0,0,0,1],
           [1,0,0,0],
           [0,1,0,0]]

Sigma_2 = [[0,0,0,1],
           [0,0,-1,0],
           [0,-1,0,0],
           [1,0,0,0]]

Sigma_3 = eta = diag(1,1,-1,-1)
```

and verifies

```text
Sigma_a^2 = I
Sigma_a Sigma_b + Sigma_b Sigma_a = 2 delta_ab I
```

Status:

```text
CONDITIONAL_SUPPORT_SEALED_PAULI_CL30_TRIPLET_CONSTRUCTED_ON_HPHI
PASS_PAULI_CL30_RELATIONS_VERIFIED
```

The Pauli triple is constructible under the sealed `H_phi=C^2` scalar carrier. It is not promoted to an unsealed physical gauge or weak-isospin action.

## 3. Moment-Map Coordinates

For

```text
x=(a,b,c,d)^T,
```

define

```text
r^2 = x^T x = a^2+b^2+c^2+d^2 = |z1|^2+|z2|^2
mu_a = x^T Sigma_a x.
```

Gate 560 verifies

```text
mu_1 = 2(ac+bd)
mu_2 = 2(ad-bc)
mu_3 = a^2+b^2-c^2-d^2
```

Status:

```text
PASS_PAULI_MOMENT_COORDINATES_VERIFIED
```

## 4. Hopf Identity

The scalar moment coordinates satisfy

```text
mu_1^2 + mu_2^2 + mu_3^2 = (r^2)^2.
```

Equivalently, in complex notation,

```text
phi phi^dagger = 1/2(r^2 I_2 + mu_a sigma_a).
```

Status:

```text
PASS_HOPF_MOMENT_IDENTITY_VERIFIED
```

## 5. Scalar 4 -> 1+3 Audit

The map

```text
phi -> (r^2, mu_1, mu_2, mu_3)
```

exhibits a sealed scalar-sector decomposition

```text
H_phi scalar coordinates -> scalar radius + Pauli moment triplet
4 -> 1 + 3
```

Status:

```text
CONDITIONAL_SUPPORT_SCALAR_SECTOR_4_TO_1PLUS3_RADIUS_AND_MOMENT_TRIPLET
```

This is scalar-sector record geometry only. It does not identify the Pauli moment triplet with gauge bosons, `W_spatial`, weak isospin, weak-plane candidates, flavor, Yukawa texture, or CKM/PMNS data.

## 6. Moment-Vector 3 -> 1+2 Audit

For nonzero scalar moment vector `mu`, the Pauli record space has the canonical stabilizer/orbit split

```text
R^3_sigma = R mu ⊕ mu^perp.
```

Status:

```text
CONDITIONAL_SUPPORT_SCALAR_MOMENT_VECTOR_3_TO_1PLUS2_ORBIT_STABILIZER_SPLIT
```

This is a scalar-sector `3=1+2` split. It is not a selected Fock weak plane `U_12`, `U_13`, or `U_23`.

## 7. Relation to Gate 558 Eta Records

Gate 560 proves that Gate 558 lives on the `Sigma_3` axis of the larger Pauli triplet:

```text
eta = Sigma_3
O1 = (I+Sigma_3)/2
O2 = (I-Sigma_3)/2
O3 = Sigma_3/4
```

Thus

```text
tau_eta = (2,-2,1)
```

is the `Sigma_3`-axis eta-trace shadow of the larger scalar Pauli moment structure, not an operator spectrum and not a spatial/generation selector.

Status:

```text
PASS_ETA_RECORDS_IDENTIFIED_AS_SIGMA3_AXIS_SHADOW
CONDITIONAL_SUPPORT_TAU_ETA_IS_SIGMA3_AXIS_TRACE_SHADOW_OF_PAULI_TRIPLET
```

## 8. Transfer Firewall

Gate 560 finds no current native project data connecting the scalar Pauli moment triplet

```text
R^3_sigma
```

to

```text
W_spatial
weak-plane candidates U_12, U_13, U_23
C^3_gen
```

Status:

```text
FAILED_ROUTE_NO_PAULI_MOMENT_TO_FOCK_OR_GENERATION_FUNCTOR
FAILED_ROUTE_PAULI_MOMENT_TRIPLET_DOES_NOT_SELECT_W_SPATIAL_WEAK_PLANE_OR_GENERATION
FIREWALL_PRESERVED_GATE560_PAULI_HOPF_SCALAR_MOMENT_BOUNDARY
```

## Final Answers

```text
A. Does H_phi contain a sealed Pauli/Cl(3,0) triplet?
   Yes. Under the sealed H_phi=C^2 scalar carrier, Sigma_1,Sigma_2,Sigma_3 satisfy the Cl(3,0) Pauli relations.

B. Does the Hopf moment identity hold?
   Yes. |mu|^2=(r^2)^2.

C. Does this produce scalar-sector 4=1+3?
   Yes. The scalar coordinates map to radius plus Pauli moment triplet.

D. Does nonzero mu produce scalar-sector 3=1+2?
   Yes. R^3_sigma=R mu plus mu^perp for nonzero mu.

E. Is eta only the Sigma_3 axis of this larger structure?
   Yes. eta=Sigma_3, and the Gate 558 records O1,O2,O3 live only on that axis.

F. Is there any lawful transfer to W_spatial, weak-plane candidates, or generation carrier?
   No. A separate native functor/intertwiner is still missing.
```

## Required Next Theorem

The next nontrivial route is not another trace/rank transfer. It is a functor/intertwiner problem:

```text
F : R^3_sigma -> W_spatial or weak-plane incidence data
```

or

```text
F : R^3_sigma -> C^3_gen
```

with basis-independent target labels, B-L compatibility, grading compatibility, J compatibility, D compatibility, first-order compatibility, and an explicit proof that the image is not a hand-chosen weak plane or generation hierarchy.
