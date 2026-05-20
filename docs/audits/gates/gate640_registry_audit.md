# Gate 640 — TwistResidual RationalCompression Audit

## Purpose

Gate 639 certified the repeated compact/split obstruction residual on `K_7`:

```text
rho_twist ≈ 0.470317081001772.
```

Gate 640 audits the sharper rational compression exposed by its square:

```text
rho_twist^2 ≈ 48/217.
```

This is an internal finite-geometry obstruction audit only.  It does not derive boundary stress, scalar RG matching, physical spacetime, flavor, Higgs mass, CKM/PMNS, gauge unification, split-G2, or a native `7/72` theorem.

## Inherited Gate639 data

Gate 639 classified `rho_twist` as a repeated projective mismatch between the admissible `S_K`-twisted split Hitchin metrics and `B_K`:

```text
rho_twist = min_c ||g_twist - c B_K||_F / ||B_K||_F.
```

The residual cluster is inherited from:

| Route | Inertia | Relative residual to `B_K` |
|---|---:|---:|
| `omega_1_alt = Alt[Omega_0(S_Kx,y,z)]` | `(4,3,0)` | `0.470317081001771` |
| `omega_2_alt = Alt[Omega_0(S_Kx,S_Ky,z)]` | `(3,4,0)` | `0.470317081001770` |
| `omega_B_alt = Alt[B_K(x ×_{Omega_0} y,z)]` | `(4,3,0)` | `0.470317081001773` |

Gate 639 also certified that the residual survives projective normalization probes and is not removed by switching compact `P_G` pullback sources.

## Rational compression audit

Gate 640 tests:

```text
rho_twist^2 ?= 48/217.
```

Using the inherited matrix residual:

```text
rho_twist        ≈ 0.470317081001772,
rho_twist^2      ≈ 0.221198156682027,
48/217           ≈ 0.221198156682028,
sqrt(48/217)     ≈ 0.470317081001772.
```

The difference is at float64 matrix tolerance, so the rational compression is conditionally supported as a candidate skeleton of the obstruction residual.

## Route-level repetition

The same rational compression is checked route by route.  Each Gate639 cluster route has residual square equal to `48/217` within matrix tolerance:

```text
omega_1_alt: rho^2 - 48/217 ≈ -1.25e-15,
omega_2_alt: rho^2 - 48/217 ≈ -2.19e-15,
omega_B_alt: rho^2 - 48/217 ≈ +6.38e-16.
```

Thus the compression is not attached only to the averaged cluster value.

## Dimensional skeleton audit

Gate 634 certified the Hodge polarity:

```text
dim K_7^+ = 4,
dim K_7^- = 3.
```

The ambient Hodge split has:

```text
dim Lambda^4_+ R^8 = 35,
dim Lambda^4_- R^8 = 35.
```

Gate 640 records the typed candidate:

```text
48 = 4^2 * 3
   = (dim K_7^+)^2 dim K_7^-,

217 = 7*(35-4)
    = dim K_7 * dim(Lambda^4_+ R^8 / K_7^+).
```

So the rational skeleton can be written as:

```text
rho_twist^2 ?=
((dim K_7^+)^2 dim K_7^-)
/
(dim K_7 * (dim Lambda^4_+ R^8 - dim K_7^+))

= 4^2*3 / [7*(35-4)]
= 48/217.
```

## Trace/projector contraction audit

Gate 640 audits candidate source expressions involving the typed dimensions of:

```text
P_+,
P_-,
P_{K_7^+},
P_{K_7^-}.
```

The dimensional expression above matches `48/217`, while simpler typed ratios such as `dim K_7/dim Lambda^4`, the Hodge trace imbalance `(4-3)/7`, or the occupancy `dim K_7^+/dim Lambda^4_+` do not.

However, this is still only a dimensional compression.  No native trace identity or projector-contraction theorem has been derived that forces:

```text
rho_twist^2 = 48/217.
```

## Native ASHA status

Certified:

```text
rho_twist remains the Gate639 compact/split obstruction witness,
rho_twist^2 compresses to 48/217 within matrix tolerance,
48 matches 4^2*3 from the K_7 Hodge polarity,
217 matches 7*(35-4) from K_7 and the self-dual complement to K_7^+.
```

Not certified:

```text
native trace derivation of 48/217,
B_K-compatible split-G2 structure,
boundary-stress assignment,
scalar/flavor transport theorem,
physical metric theorem,
native 7/72 theorem.
```

## Verdict

```text
PASS_GATE639_RHO_TWIST_INVARIANT_INHERITED
PASS_RHO_TWIST_SQUARED_RATIONAL_COMPRESSION_TESTED
CONDITIONAL_SUPPORT_RHO_TWIST_SQUARED_EQUALS_48_OVER_217_CANDIDATE
CONDITIONAL_SUPPORT_DENOMINATOR_217_MATCHES_7_TIMES_SELF_DUAL_COMPLEMENT_31
CONDITIONAL_SUPPORT_NUMERATOR_48_MATCHES_4_SQUARED_TIMES_3_HODGE_POLARITY
PASS_48_OVER_217_COMPRESSION_REPEATED_ACROSS_GATE639_ROUTES
PASS_TRACE_PROJECTOR_CONTRACTION_CANDIDATES_AUDITED
FAILED_ROUTE_NO_NATIVE_TRACE_DERIVATION_OF_48_OVER_217_YET
FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM
FAILED_ROUTE_RHO_COMPRESSION_IS_NOT_PHYSICAL_METRIC_THEOREM
FIREWALL_PRESERVED_GATE640_RATIONAL_COMPRESSION_IS_OBSTRUCTION_ONLY
```
