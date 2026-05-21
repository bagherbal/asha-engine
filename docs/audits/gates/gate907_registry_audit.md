# Gate 907 — GlobalPhaseZ2 Equivariance and OrientationGauge Audit

## Purpose

Gate 907 follows Gate 906's result:

```text
R3_AIRLOCK_PHASE_MODULE_INDUCED_POSITIVE_GENERATOR_MISSING
```

Gate 906 reduced the remaining master wound to the sign of:

```text
Q_phi = e_lambda - e_barlambda.
```

Gate 907 audits whether the absolute choice of `+Q_phi` is required, or whether the reversal:

```text
lambda <-> bar(lambda)
e_lambda <-> e_barlambda
h_lambda <-> h_barlambda
Q_phi -> -Q_phi
```

is a global phase-orientation gauge. This gate does not derive `alpha_B`, does not assign physical sectors, does not derive individual Yukawa values, does not promote to native R3, and does not update official ledgers.

---

## Implemented package

```text
pkg/bridge/generation2globalphasez2equivarianceorientationgaugeaudit
```

Registered theorem:

```text
generation2globalphasez2equivarianceorientationgaugeaudit.Generation2GlobalPhaseZ2EquivarianceOrientationGaugeAuditTheorem()
```

---

## Global Z2 operation

Gate 907 defines the global phase flip:

```text
tau_phi: lambda <-> bar(lambda)
```

with:

```text
tau_phi(e_lambda)=e_barlambda
tau_phi(e_barlambda)=e_lambda
tau_phi(Q_phi)=-Q_phi.
```

This does not yet certify a native theorem. It defines the candidate orientation-gauge operation.

---

## Audit I — airlock pair equivariance

The two punctures are:

```text
p_lambda     = e_lambda tensor P_1
p_barlambda  = e_barlambda tensor P_1.
```

Under `tau_phi`, they exchange.

The corresponding flags have the same ranks:

```text
rank(F_1/F_0)=3
rank(F_2/F_0)=7.
```

Therefore the alpha rank reconstruction:

```text
alpha_B = (3/10)s + (7/72)s^2
```

is invariant under the global phase flip.

Conditional support:

```text
CONDITIONAL_SUPPORT_GLOBAL_PHASE_Z2_EXCHANGES_AIRLOCK_REPRESENTATIVES
CONDITIONAL_SUPPORT_ALPHA_RANK_RECONSTRUCTION_IS_Z2_INVARIANT
CONDITIONAL_SUPPORT_PUNCTURE_AIRLOCK_EXISTS_AS_ORIENTATION_CLASS
```

---

## Audit II — edge table equivariance

The current oriented edge table and the mirror edge table both have active rank pattern:

```text
3 + 3 + 1 = 7.
```

Both leave one neutral singleton outside the image. Therefore the edge skeleton is mirrored, not destroyed, by the phase flip.

Conditional support:

```text
CONDITIONAL_SUPPORT_EDGE_TABLE_HAS_Z2_MIRROR_REPRESENTATIVE
CONDITIONAL_SUPPORT_EDGE_RANK_AND_KERNEL_COUNT_ARE_Z2_INVARIANT
CONDITIONAL_SUPPORT_ORIENTED_EDGE_TABLE_IS_A_GAUGE_REPRESENTATIVE_IF_Z2_EQUIVARIANCE_CERTIFIED
```

Preserved caution:

```text
FAILED_ROUTE_EDGE_EQUIVARIANCE_NOT_YET_NATIVE_OPERATOR_THEOREM
```

---

## Audit III — trace-magnitude readout equivariance

The current readout is:

```text
Y^dagger Y = 1*Pi_lambda3
           + alpha_B(1-alpha_B)*Pi_barlambda3
           + 3 alpha_B^2*Pi_barlambda1.
```

The mirror readout exchanges labels:

```text
Y_tau^dagger Y_tau = 1*Pi_barlambda3
                   + alpha_B(1-alpha_B)*Pi_lambda3
                   + 3 alpha_B^2*Pi_lambda1.
```

The row multiset is unchanged:

```text
(rank 3, weight 1)
(rank 3, weight alpha_B(1-alpha_B))
(rank 1, weight 3 alpha_B^2)
```

Therefore:

```text
a_total/T     = 3 + 3 alpha_B
b_total/T^2   = 3 + 3 alpha_B^2 - 6 alpha_B^3 + 12 alpha_B^4
N_eff^operator and C_Yukawa^operator are phase-orientation invariant.
```

Conditional support:

```text
CONDITIONAL_SUPPORT_TRACE_MAGNITUDE_ROW_MULTISET_IS_Z2_INVARIANT
CONDITIONAL_SUPPORT_OPERATOR_N_EFF_IS_PHASE_ORIENTATION_INVARIANT
CONDITIONAL_SUPPORT_C_YUKAWA_OPERATOR_IS_PHASE_ORIENTATION_INVARIANT
```

---

## Audit IV — label firewall

The socket labels change under the phase flip:

```text
lambda socket <-> bar(lambda) socket
h_lambda <-> h_barlambda.
```

But no physical sector assignment is certified in the current branch, and the trace ledger does not require physical labels.

Conditional support:

```text
CONDITIONAL_SUPPORT_SOCKET_LABELS_ARE_ORIENTATION_DEPENDENT
CONDITIONAL_SUPPORT_TRACE_LEDGER_DOES_NOT_REQUIRE_PHYSICAL_LABEL_ASSIGNMENT
```

Preserved failures:

```text
FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT
FAILED_ROUTE_NO_GENERATION_CARRIER_MAP
FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP
```

---

## Verdict

Gate 907 concludes:

```text
GLOBAL_PHASE_Z2_EQUIVARIANCE_CAN_REMOVE_ABSOLUTE_SIGN_SELECTION_PRESSURE_IF_CERTIFIED
```

Classification:

```text
R3_AIRLOCK_PHASE_SIGN_RECLASSIFIED_AS_ORIENTATION_GAUGE_CANDIDATE
```

Short status:

```text
R3_PHASE_ORIENTATION_GAUGE_EQUIVARIANCE_CANDIDATE_NOT_NATIVE
```

The wound changes from:

```text
select +Q_phi rather than -Q_phi
```

to:

```text
certify a Z2-equivariant airlock orientation class.
```

---

## Conditional supports

```text
CONDITIONAL_SUPPORT_GLOBAL_PHASE_Z2_EXCHANGES_AIRLOCK_REPRESENTATIVES
CONDITIONAL_SUPPORT_ALPHA_RANK_RECONSTRUCTION_IS_Z2_INVARIANT
CONDITIONAL_SUPPORT_EDGE_TABLE_HAS_Z2_MIRROR_REPRESENTATIVE
CONDITIONAL_SUPPORT_TRACE_MAGNITUDE_ROW_MULTISET_IS_Z2_INVARIANT
CONDITIONAL_SUPPORT_OPERATOR_N_EFF_IS_PHASE_ORIENTATION_INVARIANT
CONDITIONAL_SUPPORT_C_YUKAWA_OPERATOR_IS_PHASE_ORIENTATION_INVARIANT
CONDITIONAL_SUPPORT_ABSOLUTE_Q_PHI_SIGN_MAY_BE_ORIENTATION_GAUGE_NOT_NATIVE_OBSERVABLE
CONDITIONAL_SUPPORT_NATIVE_R3_PRESSURE_REDUCES_TO_Z2_EQUIVARIANT_AIRLOCK_THEOREM
```

## Preserved firewalls

```text
FAILED_ROUTE_NO_NATIVE_GLOBAL_PHASE_Z2_EQUIVARIANCE_THEOREM
FAILED_ROUTE_NO_NATIVE_PHASE_MODULE_TO_ASHA_C_R2_IDENTIFICATION
FAILED_ROUTE_NO_NATIVE_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTOR
FAILED_ROUTE_ALPHA_B_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_BOUNDARY_FUNCTOR
FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_WEAK_SELECTOR
FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED
FAILED_ROUTE_NOT_NATIVE_R3
FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT
FAILED_ROUTE_NO_GENERATION_CARRIER_MAP
FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP
FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES
FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED
FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM
```
