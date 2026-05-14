# Gate 245 Registry Audit — Lie Algebra Isomorphism / Scalar-to-Spatial Carrier Projection

## Verdict

```text
CONDITIONAL_SUPPORT_EW_OPERATOR_DECOMPOSITION_TRACED
FAILED_ROUTE_TAU_ETA_SLOTS_NOT_SU2_BASIS
CONDITIONAL_SUPPORT_SU2_BIVECTOR_CAPACITY_PREFLIGHT
FAILED_ROUTE_NATIVE_SU2_TO_SPATIAL_AXIS_ISOMORPHISM
FAILED_ROUTE_SCALAR_TO_SPATIAL_CARRIER_PROJECTION
FAILED_ROUTE_LIE_PULLBACK_EXTERIOR_FORM_REPRESENTATIVE
FAILED_ROUTE_LIE_PULLBACK_WEAK_PLANE_SELECTION
FAILED_ROUTE_LIE_PULLBACK_GENERATION_TEXTURE
FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED
```

Gate 245 traces the proposed projection chain and rejects the final carrier projection. The scalar trace values are structured, but they do not land on the spatial Fock axes through the current finite algebra.

## Inherited obstruction

Gate 244 established:

```text
tau_eta = (2,-2,1)

tau_eta(Q^TQ)        =  2
tau_eta(Z^TZ)        = -2
tau_eta(T3L^T Y_phi) =  1
```

and also established that these are scalar-bundle curvature observables on `H_Phi`, not spatial Fock-mode projectors.

## 1. Operator decomposition audit

Gate 245 decomposes the source observables:

| tau slot | Source expression | Electroweak decomposition | Result |
|---:|---|---|---|
| `0` | `tau_eta(Q^TQ)=2` | `Q = T3L + Y_phi` | neutral scalar quadratic observable |
| `1` | `tau_eta(Z^TZ)=-2` | `Z = T3L - Y_phi` | neutral scalar quadratic observable |
| `2` | `tau_eta(T3L^T Y_phi)=1` | mixed `T3L/Y_phi` pairing | neutral scalar bilinear observable |

This is a strict obstruction to the proposed chain. The three tau slots are not:

```text
T1, T2, T3
```

They are scalar records built from:

```text
span{T3L, Y_phi}
```

So the source triple has three entries, but it does not represent the three nonabelian `su(2)` basis generators.

## 2. Derivation-to-blade isomorphism audit

The finite core still supports a useful preflight:

```text
spatial bivectors = {e1∧e2, e2∧e3, e3∧e1}
```

These have the right abstract shape to carry an `su(2)`-like rotation algebra. However, the engine has not derived:

```text
contact-preserving su(2) generator matrices
ordered map {T1,T2,T3} -> {e1∧e2,e2∧e3,e3∧e1}
ordered map from bivectors to Fock spatial axes {a†_1,a†_2,a†_3}
```

So the second link of the projection chain is also missing.

## 3. Carrier projection theorem

The tempting projection remains rejected:

```text
tau_eta(Q^TQ, Z^TZ, T3L^T Y_phi) ?-> (2e_1, -2e_2, e_3)
```

Failure reason:

```text
first link: tau_eta slots are neutral scalar observables, not su(2) basis elements
second link: contact-su(2) to ordered spatial-axis map is not derived
```

Therefore Gate 245 does not construct:

```text
omega_tau = 2e_1 - 2e_2 + e_3
```

## 4. Weak-plane and generation consequences

The conditional roadmap remains visible:

```text
if omega_tau were lawfully derived:
|tau_eta| = (2,2,1) would tag a†_3
weak plane would be U={a†_1,a†_2}
```

But this remains conditional only.

The signed spectrum also remains visible:

```text
tau_eta = (2,-2,1)
```

which has the right `1+1+1` capacity for generation breaking, but Gate 245 does not derive the separate `tau_eta -> triality generation carrier` map.

## Firewall audit

Gate 245 does not:

```text
force Q,Z,T3Y slots onto axes
force su(2) generators onto spatial bivectors
construct omega_tau by hand
import the weak plane
import Connes' algebra
claim physical chirality
claim global H
claim CKM/PMNS
claim fermion masses
```

## Final theorem distinction

```text
operator decomposition known: yes
scalar slots are su(2) basis: no
spatial bivector su(2) capacity: yes
canonical su(2)->axis map: no
carrier projection theorem: no
```

Gate 245 narrows the missing theorem to an explicit `H_Phi -> W` representation functor or a future seal acknowledging scalar-to-spatial projection as phenomenological boundary data.
