# Gate 541 Registry Audit — Real-Looking Schwinger Source Negative-Control Adapter

## Verdict

```text
CONDITIONAL_SUPPORT_GATE540_REAL_IMPORT_SWITCH_INHERITED
CONDITIONAL_SUPPORT_REAL_LOOKING_SCHWINGER_NEGATIVE_CONTROL_LEDGER_LOADED
CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_SWITCH_ROWS_ACCEPTED
CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_CHECKSUM_VERIFIED
CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_METADATA_PARSED
CONDITIONAL_SUPPORT_REAL_LOOKING_SCHWINGER_NEGATIVE_CONTROL_ADAPTER_EXECUTED
CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_REJECTED_SWITCH_OFF
CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_REJECTED_NO_OPERATOR_INTENT
CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_REJECTED_INSUFFICIENT_PROVENANCE
CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_COMPARATOR_AUTHORIZATION_BLOCKED
CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_QUARANTINE_PRESERVED
CONDITIONAL_SUPPORT_NO_REAL_SOURCE_COMPARATOR_EXECUTED_IN_GATE541
CONDITIONAL_SUPPORT_NO_NATIVE_WRITE_FROM_REAL_LOOKING_SOURCE_GATE541
FAILED_ROUTE_REAL_LOOKING_SOURCE_IMPORT_SWITCH_OFF
FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_NO_EXPLICIT_OPERATOR_INTENT
FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_INSUFFICIENT_PROVENANCE
FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_NO_LICENSE_OR_ACCESS_GRANT
FAILED_ROUTE_REAL_LOOKING_SOURCE_URI_NOT_AUTHENTICATED
FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS
FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY
FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_WICK_ROTATION
FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE
FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN
FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS
FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY
FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_SELECT_ARROW_OF_TIME
FIREWALL_PRESERVED_GATE541_REAL_LOOKING_NEGATIVE_CONTROL_BRIDGE_ONLY
FIREWALL_BLOCKED_GATE541_REAL_LOOKING_SOURCE_NATIVE_WRITE
```

## Inherited boundary

Gate 541 inherits Gate 540's real Schwinger source import switch. Gate 540 defined the required switch rows and kept the switch off by default. Gate 541 therefore tests the rejection path, not a real import path.

```text
CONDITIONAL_SUPPORT_GATE540_REAL_IMPORT_SWITCH_INHERITED: switch_defined=true default_off=true intent_needed=true comparator_blocked=true no_real=true native_blocked=true redirects=true; Gate541 inherits Gate540's default-off real-source switch, missing operator-intent guard, comparator block, and native-write lock.
```

## Negative-control fixture

Gate 541 loads a deliberately real-looking but untrusted source fixture:

```text
data/real_looking_schwinger_negative_control_ledger_gate541.json
```

The fixture claims to be non-synthetic/physical-looking, but it is marked as a negative-control ledger and lacks the required operator intent, trusted source URI, license/access grant, authenticity reference, source proof hash, and comparator authorization.

Checksum of the canonical payload:

```text
sha256:176e104152f81376e1dabb59adc67a5ff1dd174df3377f490c1e61da0dcf9884
```

## Switch-row parsing result

The 12 Gate 540 switch rows are present and parsed:

| Row | Negative-control value |
|---|---|
| `real_source_import_switch` | `off/default-deny` |
| `explicit_operator_intent` | `absent` |
| `non_synthetic_source_uri` | invalid negative-control URI |
| `authenticity_ledger_reference` | absent/unverified |
| `checksum_or_proof_hash_reference` | fixture checksum only; no source proof hash |
| `license_and_access_grant_reference` | absent |
| `source_class_non_synthetic_assertion` | claimed non-synthetic/physical-looking but untrusted |
| `gate536_schema_alignment_reference` | claimed alignment; not comparator-authorized |
| `comparator_execution_plan` | requested by fixture, blocked by switch |
| `quarantine_output_target` | bridge quarantine only |
| `native_write_lock` | locked |
| `rollback_audit_trace` | negative-control trace only |

Parser status:

```text
CONDITIONAL_SUPPORT_REAL_LOOKING_SCHWINGER_NEGATIVE_CONTROL_LEDGER_LOADED: rows=12 accepted=12 rejected=0 bridge=true native_request=false real_looking=true negative=true declared_synthetic=false physical_claim=true nonsynthetic_claim=true switch=false intent=false comparator_request=true comparator_grant=false license=false auth_ref=false checksum_ref=false trusted_uri=false physical_loaded=false observed=false measure=false OS=false Wick=false Hamiltonian=false gate540=true gate539=true gate536=true metadata=true required=true all_bridge=true all_no_theorem=true all_negative=true checksum=true expected=sha256:176e104152f81376e1dabb59adc67a5ff1dd174df3377f490c1e61da0dcf9884 actual=sha256:176e104152f81376e1dabb59adc67a5ff1dd174df3377f490c1e61da0dcf9884 failures=[]; the real-looking negative-control fixture parses and verifies checksum plumbing, but it is intentionally untrusted and cannot pass the real-source switch.
```

## Rejection result

The fixture is rejected before comparator execution.

```text
CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_REJECTED_SWITCH_OFF: parsed=true real_looking=true physical_claim=true negative=true switch_off=true no_intent=true no_license=true uri_untrusted=true auth_missing=true insufficient=true allowed=false performed=false authenticated=false imported=false rejected=true before_comparator=true quarantine=true reasons=[FAILED_ROUTE_REAL_LOOKING_SOURCE_IMPORT_SWITCH_OFF FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_NO_EXPLICIT_OPERATOR_INTENT FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_NO_LICENSE_OR_ACCESS_GRANT FAILED_ROUTE_REAL_LOOKING_SOURCE_URI_NOT_AUTHENTICATED FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_INSUFFICIENT_PROVENANCE]; the fixture is deliberately real-looking, but default-off switch state, absent operator intent, untrusted URI, and missing license/access provenance reject it before comparator execution.
```

## Firewall result

```text
FIREWALL_PRESERVED_GATE541_REAL_LOOKING_NEGATIVE_CONTROL_BRIDGE_ONLY;FIREWALL_BLOCKED_GATE541_REAL_LOOKING_SOURCE_NATIVE_WRITE;FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS;FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY;FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_WICK_ROTATION;FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE;FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN;FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS;FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY;FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_SELECT_ARROW_OF_TIME: comparator=false real=false observed=false measure=false OS_import=false wick_import=false hamiltonian_import=false derived=false OS=false Wick=false Hilbert=false Hamiltonian=false unitary=false global=false arrow=false native_Schwinger=false native_measure=false native_OS=false native_Wick=false native_Hilbert=false native_Hamiltonian=false native_unitary=false native_global=false native_arrow=false reopened_flavor=false reopened_EW=false reopened_gravity=false reopened_topology=false reopened_dimension=false reopened_Krein=false native_registry=false; real-looking negative-control data stays quarantined: no comparator runs and no physical correlation/dynamics object writes into the native registry.
```

## Registry update

### Native

- No new native law is written at Gate 541.
- A real-looking fixture, checksum success, and parser success are not evidence of physical Schwinger functions, OS positivity, Wick rotation, Hilbert reconstruction, Hamiltonian dynamics, global causality, or time orientation.

### Bridge

- Gate 541 adds a real-looking negative-control adapter for the real-source import switch.
- The adapter verifies that default-off switch state, missing explicit operator intent, untrusted URI, absent license/access grant, missing authenticity/proof references, and absent comparator authorization reject the input before any comparator can run.

### Environmental

- Any future real or constructive Schwinger source remains sourced bridge/environmental data and must provide explicit operator intent, authenticated URI, access/license grant, authenticity ledger, proof hash, Gate 536 alignment, comparator plan, quarantine target, rollback trace, and a native-write lock.

### Failed routes

- FAILED_ROUTE_REAL_LOOKING_SOURCE_IMPORT_SWITCH_OFF
- FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_NO_EXPLICIT_OPERATOR_INTENT
- FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_INSUFFICIENT_PROVENANCE
- FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_NO_LICENSE_OR_ACCESS_GRANT
- FAILED_ROUTE_REAL_LOOKING_SOURCE_URI_NOT_AUTHENTICATED
- FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS
- FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY
- FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_WICK_ROTATION
- FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE
- FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN
- FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS
- FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY
- FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_SELECT_ARROW_OF_TIME
- FIREWALL_BLOCKED_GATE541_REAL_LOOKING_SOURCE_NATIVE_WRITE

### Open theorems

- Define a controlled comparator-authorization manifest airlock for a future real source.
- Confirm that authorization can target only bridge quarantine and cannot unlock native registry promotion.

## Next step

Gate 542 — Real Source Comparator Authorization Manifest Airlock. Gate 541 proves that real-looking data is rejected when the switch is off or provenance is insufficient. The next safe boundary is the authorization manifest required before a real-source comparator can run in bridge quarantine.

Primary task: Enumerate the comparator-authorization manifest and confirm that authorization can only target bridge quarantine, never native registry promotion.

## Truth statement

Gate541 proves the default-deny import path: a real-looking Schwinger source fixture can parse and verify checksum plumbing, but the off switch, missing operator intent, untrusted URI, incomplete provenance, and absent access grant reject it before comparator execution or native registry write.
