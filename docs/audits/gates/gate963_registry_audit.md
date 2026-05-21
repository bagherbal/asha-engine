# Gate 963 — CanonicalFlavorSelector vs ExternalFlavorOrientationSeal Decision Audit

## Package

```text
pkg/bridge/generation2canonicalflavorselectorvsexternalorientationsealdecisionaudit
```

## Registered theorem

```text
generation2canonicalflavorselectorvsexternalorientationsealdecisionaudit.Generation2CanonicalFlavorSelectorVsExternalOrientationSealDecisionAuditTheorem()
```

## Purpose

Gate 963 follows Gate 962:

```text
R4_EXTERNAL_C3_HAS_U3_FAMILY_ORBIT_NO_FLAVOR_BASIS
```

Gate 962 showed that the sealed generation carrier

```text
G_gen^seal = C^3
```

supplies only a family-orientation orbit

```text
[Phi_flav]_{U(3)}
```

and no canonical flavor basis. Gate 963 decides whether the current sealed R4 certificate contains a lawful `CanonicalFlavorSelector`, or whether downstream flavor-ledger testing requires an explicit `ExternalFlavorOrientationSeal`.

This gate does not derive Yukawa eigenvalues, CKM/PMNS, physical particles, observed masses, or official ledger values.

## Audit result

The current certificate does not contain a canonical selector. The following sources remain insufficient:

```text
A_F^orient
R3 tracebody
socket ledger
Fock 1+3 / B-L
boundary activation
K7 remnants
Boolean-octonionic projectors
```

The obstruction is the surviving family-gauge freedom:

```text
C^3_gen,seal -> U(3) C^3_gen,seal
```

Without an additional selector, no representative of the flavor-orientation orbit is canonical.

## Seal decision

The allowed quarantine object is:

```text
ExternalFlavorOrientationSeal
```

It may select a representative:

```text
Phi_flav^seal in [Phi_flav]_{U(3)}
```

Allowed role:

```text
orientation representative for downstream ledger tests
```

Forbidden roles:

```text
native flavor theorem
Yukawa spectrum theorem
CKM/PMNS theorem
particle assignment theorem
official ledger update
```

## Verdict

```text
NO_CANONICAL_FLAVOR_SELECTOR_FOUND_EXTERNAL_FLAVOR_ORIENTATION_SEAL_REQUIRED_FOR_DOWNSTREAM_FLAVOR_LEDGER_TESTS
```

## Classification

```text
R4_FLAVOR_ORIENTATION_SOURCE_DECISION_EXTERNAL_SEAL_REQUIRED
```

## Short status

```text
R4_REQUIRES_EXTERNAL_FLAVOR_ORIENTATION_SEAL
```

## Conditional supports

```text
CONDITIONAL_SUPPORT_CURRENT_CERTIFICATE_SEARCH_COMPLETED_FOR_CANONICAL_FLAVOR_SELECTOR
CONDITIONAL_SUPPORT_EXTERNAL_C3_RETAINS_U3_FAMILY_GAUGE_FREEDOM
CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_CAN_SELECT_REPRESENTATIVE
CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_REQUIRED_FOR_DOWNSTREAM_FLAVOR_LEDGER_TESTS
CONDITIONAL_SUPPORT_DOWNSTREAM_FLAVOR_LEDGER_TESTS_ALLOWED_ONLY_UNDER_EXTERNAL_ORIENTATION_SEAL
CONDITIONAL_SUPPORT_R3_DUALSEAL_EXTERNAL_C3_SEAL_AND_EXTERNAL_FLAVOR_ORIENTATION_SEAL_MUST_REMAIN_VISIBLE
```

## Preserved firewalls

```text
FAILED_ROUTE_NO_CANONICAL_FLAVOR_SELECTOR_IN_CURRENT_CERTIFICATE
FAILED_ROUTE_U3_FAMILY_GAUGE_NOT_BROKEN_BY_CURRENT_ASHA_DATA
FAILED_ROUTE_NO_CANONICAL_REPRESENTATIVE_OF_FLAVOR_ORIENTATION_ORBIT
FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_NATIVE_FLAVOR_THEOREM
FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_YUKAWA_SPECTRUM_THEOREM
FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_CKM_PMNS_THEOREM
FAILED_ROUTE_A_F_ORIENT_DOES_NOT_SUPPLY_FAMILY_ORIENTATION_SELECTOR
FAILED_ROUTE_SOCKET_ORIENTATION_NOT_FAMILY_ORIENTATION
FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS
FAILED_ROUTE_R3_TRACEBODY_CANNOT_ORIENT_EXTERNAL_C3_BY_ROW_MATCHING
FAILED_ROUTE_FOCK_P3_RANK_THREE_NOT_FLAVOR_ORIENTATION
FAILED_ROUTE_B_MINUS_L_SELECTOR_NOT_FAMILY_ORIENTATION
FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR
FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_ORIENT_EXTERNAL_C3
FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES
FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT
FAILED_ROUTE_NO_CKM_PMNS_THEOREM
FAILED_ROUTE_NO_PMNS_THEOREM
FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM
FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE
FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER
FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED
```

## Strategic result

Gate 963 closes the flavor-orientation decision exactly the way Gate 959 closed generation multiplicity. The current sealed lane becomes:

```text
R3DualSeal
ExternalGenerationCarrierSeal(C3)
ExternalFlavorOrientationSeal required for downstream flavor-ledger tests
```

This is not native flavor. It is an honest quarantine boundary for later diagnostics.
