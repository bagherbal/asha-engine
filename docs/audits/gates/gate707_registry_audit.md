# Gate 707 — Central Baseline Gauge and Scalar-Wall Reference Selection Audit

## Purpose

Gate 706 isolated the positive-distance observable as:

```text
W_boundary = |lambda|I_H72 + S_split P_K7.
```

Gate 707 audits whether the scalar baseline `|lambda|I_H72` is uniquely typed, or
whether it is one representative inside a central baseline gauge family.

The general central decomposition is:

```text
W_boundary = c I_H72 + (R-c)P_K7 + (|lambda|-c)P_perp.
```

This is a bridge-layer central-baseline gauge audit only. It does not derive
boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
CKM/PMNS, a native response theorem, a native state-selection theorem, or a
native `7/72` theorem.

## Implementation

- Package: `pkg/bridge/generation2centralbaselinegaugeandscalarwallreferenceselectionaudit`
- Registered theorem: `generation2centralbaselinegaugeandscalarwallreferenceselectionaudit.Generation2CentralBaselineGaugeAndScalarWallReferenceSelectionAuditTheorem()`

## Central baseline gauge family

For arbitrary scalar `c`:

```text
W_boundary
= c I_H72 + (R-c)P_K7 + (|lambda|-c)P_perp.
```

Under the full augmented no-bias state:

```text
Tr(rho_72 W_boundary)
= c + (7/72)(R-c) + (65/72)(|lambda|-c)
= (7/72)R + (65/72)|lambda|.
```

Therefore the total wound expectation is invariant under central baseline gauge
choice.

## Active scalar baseline choice

Set:

```text
c = |lambda|.
```

Then:

```text
W_boundary
= |lambda|I_H72
+ (R-|lambda|)P_K7
+ 0 P_perp
= |lambda|I_H72 + S_split P_K7.
```

This is the unique central baseline in the gauge family that makes the complement
uplift vanish and leaves a support-local K7 uplift.

## Alternative baseline gauges

```text
c = 0:
  W = R P_K7 + |lambda|P_perp.
  Raw two-payoff form; no central baseline isolated.

c = |lambda|:
  W = |lambda|I_H72 + S_split P_K7.
  Active; complement-zero K7-local uplift.

c = R:
  W = R I_H72 - S_split P_perp.
  Algebraically valid, but makes the complement the active correction sector.

c = xi_boundary:
  W = xi I_H72 + signed corrections on both sectors.
  Less minimal; not K7-support-local.

c = (R+|lambda|)/2:
  Midpoint baseline; two-sided corrections, not active K7-local uplift.
```

## Support-locality selection

Gate 696 selected the active Bernoulli response by K7 support-locality. Gate 707
shows that the positive-distance observable has the same support-local structure
only when:

```text
c = |lambda|.
```

Thus:

```text
central baseline selection + K7 support-locality
=> W_boundary = |lambda|I_H72 + S_split P_K7.
```

## Scalar-wall airlock compatibility

Gate 703 identified `lambda(Lambda_12)` as the shared scalar-wall airlock. The
choice `c=|lambda|` uses the scalar zero-wall depth as the universal reference
level. The gauge-baseline choice `c=R` is algebraically possible, but it reverses
the active interpretation by assigning the correction to the complement sector.

## Verdict

```text
PASS_GATE706_CENTRAL_BASELINE_UPLIFT_INHERITED
PASS_CENTRAL_BASELINE_GAUGE_FAMILY_DEFINED
PASS_TOTAL_EXPECTATION_BASELINE_GAUGE_INVARIANT
PASS_ACTIVE_SCALAR_BASELINE_CHOICE_COMPUTED
PASS_COMPLEMENT_ZERO_UPLIFT_FOR_C_EQUALS_ABS_LAMBDA
PASS_TYPED_BASELINE_ALTERNATIVES_AUDITED
PASS_SUPPORT_LOCALITY_SELECTS_SCALAR_BASELINE_WITH_K7_UPLIFT
PASS_SCALAR_WALL_AIRLOCK_COMPATIBILITY_AUDITED
CONDITIONAL_SUPPORT_ABS_LAMBDA_IS_UNIQUE_BASELINE_FOR_K7_LOCAL_UPLIFT
CONDITIONAL_SUPPORT_SCALAR_BASELINE_SELECTION_IS_SUPPORT_LOCAL_REFERENCE_GAUGE
CONDITIONAL_SUPPORT_K7_UPLIFT_FORM_IS_SHARPER_THAN_RAW_TWO_PAYOFF_FORM
FAILED_ROUTE_BASELINE_CHOICE_NOT_NATIVE_YET
FAILED_ROUTE_NO_NATIVE_SCALAR_BASELINE_REFERENCE_SELECTION_THEOREM
FAILED_ROUTE_NO_NATIVE_REASON_K7_RATHER_THAN_COMPLEMENT_CARRIES_UPLIFT
FAILED_ROUTE_NO_NATIVE_BOUNDARY_WOUND_UPLIFT_THEOREM
FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM
FIREWALL_PRESERVED_GATE707_CENTRAL_BASELINE_GAUGE_BOUNDARY
```
