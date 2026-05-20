# Gate 650 — Hitchin Sector-Degree Top-Form Selection Rule Audit

## Purpose

Gate 649 showed that the admissible `S_K`-twisted Hitchin contraction has the finite channel algebra

```text
A = Omega++-
B = Omega---

AAA -> +P_+
AAB + ABA + BAA -> -3P_-.
```

Gate 650 asks whether this channel selection is forced by sector-degree saturation in the native Hodge split

```text
K_7 = K_7^+ ⊕ K_7^-,

dim K_7^+ = 4,
dim K_7^- = 3.
```

This is an internal finite tensor-degree audit only.  It does not derive split-G2, boundary stress, scalar/flavor transport, physical spacetime, Higgs mass, CKM/PMNS, gauge unification, or a native `7/72` theorem.

## Inherited data

Gate 649 supplies the finite channel ledger and firewalls:

```text
A = Omega++-
B = Omega---

g_twist ∝ P_+ - 3P_-.
```

The slot count and the negative Hodge-sector dimension are both three in ASHA,

```text
d = 3,
q = dim K_7^- = 3,
```

but Gate 649 keeps this as an ASHA carrier coincidence rather than a general dimension theorem.

## Sector-degree ledger

Gate 650 records the sector degrees

```text
A = Omega++- has degree (2,1),
B = Omega--- has degree (0,3).
```

For interior contractions:

```text
x ∈ K_7^+:
  i_x A has degree (1,1),
  i_x B = 0.

x ∈ K_7^-:
  i_x A has degree (2,0),
  i_x B has degree (0,2).
```

Because the Hitchin cubic produces a top form on the `4|3` carrier, any nonzero contribution must saturate sector degree

```text
(4,3).
```

## Positive block degree audit

For `x,y ∈ K_7^+`, the only ordered channel reaching the top degree is

```text
i_xA ∧ i_yA ∧ A:
(1,1) + (1,1) + (2,1) = (4,3).
```

Thus the positive block has the degree selection rule

```text
AAA only.
```

All other positive-block channels either are killed by a positive interior contraction into `B` or fail the `(4,3)` top-degree requirement.

## Negative block degree audit

For `x,y ∈ K_7^-`, the degree-saturating channels are exactly

```text
i_xA ∧ i_yA ∧ B:
(2,0) + (2,0) + (0,3) = (4,3),

i_xA ∧ i_yB ∧ A:
(2,0) + (0,2) + (2,1) = (4,3),

i_xB ∧ i_yA ∧ A:
(0,2) + (2,0) + (2,1) = (4,3).
```

Therefore the negative block has exactly three degree-allowed ordered placements:

```text
AAB,
ABA,
BAA.
```

This is the degree-level source of the cubic slot count.

## Mixed block degree audit

For mixed inputs `K_7^+ × K_7^-` or `K_7^- × K_7^+`, no ordered channel reaches sector degree `(4,3)`.  Gate 650 therefore explains the mixed block as

```text
g_+- = 0
```

by sector-degree impossibility.

## Sign and normalization audit

Sector-degree saturation explains which channels survive.  It does **not** by itself certify the negative sign or the equal unit weights.  Those are inherited from Gate 649's finite ledger, but a basis-free proof still requires the native octonionic calibration, orientation, and antisymmetrization identity.

Therefore Gate 650 preserves

```text
FAILED_ROUTE_SIGN_AND_EQUAL_UNIT_WEIGHT_STILL_REQUIRE_CALIBRATION_IDENTITY.
```

## Resulting slot formula

With the separate equal-unit calibration identity, the degree selection rule gives

```text
G_slot = (P_+ - 3P_-)/sqrt(31).
```

Equivalently, with `p=4`, `q=3`, and slot multiplicity `d=3`,

```text
G_slot = (P_+ - dP_-)/sqrt(p+d^2 q),

cos(theta) = (p+dq)/sqrt((p+q)(p+d^2q)),
rho^2     = pq(d-1)^2 / [(p+q)(p+d^2q)].
```

For ASHA this recovers

```text
cos(theta) = 13/sqrt(217),
rho^2 = 48/217.
```

## Carrier resonance

Gate 650 records separately:

```text
d = 3 from Hitchin cubic ordered placements,
q = 3 from dim K_7^-.
```

The degree rule sources the three placements from cubic order.  ASHA also has `dim K_7^- = 3`, so

```text
d = q
```

is a carrier resonance, not an independent general dimension theorem.

## Final verdict

```text
PASS_GATE649_CHANNEL_ALGEBRA_INHERITED
PASS_SECTOR_DEGREE_LEDGER_DEFINED
PASS_POSITIVE_BLOCK_AAA_ONLY_BY_TOP_FORM_DEGREE
PASS_NEGATIVE_BLOCK_AAB_ABA_BAA_BY_TOP_FORM_DEGREE
PASS_MIXED_BLOCK_ZERO_BY_TOP_FORM_DEGREE
CONDITIONAL_SUPPORT_CHANNEL_SELECTION_RULE_FROM_4_BY_3_SECTOR_DEGREE_SATURATION
CONDITIONAL_SUPPORT_MINUS_THREE_FROM_THREE_DEGREE_ALLOWED_CUBIC_PLACEMENTS
CONDITIONAL_SUPPORT_D_EQUALS_Q_AS_ASHA_CARRIER_RESONANCE
FAILED_ROUTE_SIGN_AND_EQUAL_UNIT_WEIGHT_STILL_REQUIRE_CALIBRATION_IDENTITY
FAILED_ROUTE_NO_FULL_SYMBOLIC_DEGREE_SELECTION_THEOREM
FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM
FAILED_ROUTE_HITCHIN_DEGREE_SELECTION_IS_NOT_PHYSICAL_METRIC
FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM
FIREWALL_PRESERVED_GATE650_HITCHIN_DEGREE_SELECTION_BOUNDARY
```
