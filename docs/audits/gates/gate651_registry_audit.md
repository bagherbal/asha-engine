# Gate 651 — Hitchin Channel Sign and Equal-Unit Calibration Audit

## Purpose

Gate 650 proved the sector-degree selection rule for the admissible `S_K`-twisted Hitchin contraction on the native `4|3` Hodge split of `K_7`:

```text
A = Omega++- ∈ Lambda^{2,1},
B = Omega--- ∈ Lambda^{0,3}.
```

Top-form saturation explains the support pattern:

```text
positive block: AAA only,
negative block: AAB + ABA + BAA only,
mixed block: zero by degree.
```

Gate 651 audits the remaining finite calibration question: why the surviving degree-allowed channels have the signs and equal unit weights

```text
AAA = +c P_+,
AAB = ABA = BAA = -c P_-.
```

This is an internal finite tensor-calibration audit only.  It does not derive split-G2, boundary stress, scalar/flavor transport, physical spacetime, Higgs mass, CKM/PMNS, gauge unification, or a native `7/72` theorem.

## Inherited data

Gate 650 supplies:

```text
K_7 = K_7^+ ⊕ K_7^-,

dim K_7^+ = 4,
dim K_7^- = 3,

A = Omega++-  has sector degree (2,1),
B = Omega---  has sector degree (0,3),

top degree = (4,3).
```

Gate 650 explains channel support, but explicitly preserves the calibration gap:

```text
FAILED_ROUTE_SIGN_AND_EQUAL_UNIT_WEIGHT_STILL_REQUIRE_CALIBRATION_IDENTITY.
```

## Orientation and volume conventions

Gate 651 fixes the oriented split convention

```text
vol_7 = vol_+ ∧ vol_-
```

with

```text
deg(vol_+) = 4,
deg(vol_-) = 3.
```

The audit records the sign conventions involved in the finite channel maps:

```text
interior contraction ordering,
wedge ordering,
Hitchin 1/6 normalization,
S_K action,
octonionic pullback orientation.
```

The sign is convention-sensitive at this level; a basis-free theorem would still need to tie those conventions to the native octonionic calibration identity.

## Surviving channel bilinear maps

Gate 651 defines the surviving channel bilinear maps as top-form coefficients:

```text
H_AAA(x,y) = coeff_top[(i_x A) ∧ (i_y A) ∧ A],
H_AAB(x,y) = coeff_top[(i_x A) ∧ (i_y A) ∧ B],
H_ABA(x,y) = coeff_top[(i_x A) ∧ (i_y B) ∧ A],
H_BAA(x,y) = coeff_top[(i_x B) ∧ (i_y A) ∧ A].
```

The finite route-normalized calibration ledger is:

```text
H_AAA|_{K_7^+} = +c P_+,
H_AAB|_{K_7^-} = -c P_-,
H_ABA|_{K_7^-} = -c P_-,
H_BAA|_{K_7^-} = -c P_-.
```

In the normalized audit convention:

```text
c = 1.
```

## Positive unit audit

The positive channel satisfies

```text
AAA = +P_+.
```

The positive block is scalar and isotropic:

```text
c_+ = (1/4) Tr(P_+ H_AAA P_+) = +1,
||H_AAA - c_+ P_+||_F = 0.
```

## Negative equal-unit audit

For the three negative channels,

```text
c_AAB = (1/3) Tr(P_- H_AAB P_-) = -1,
c_ABA = (1/3) Tr(P_- H_ABA P_-) = -1,
c_BAA = (1/3) Tr(P_- H_BAA P_-) = -1.
```

Thus, in the finite route-normalized ledger,

```text
c_AAB = c_ABA = c_BAA = -c_+.
```

The combined negative block is therefore

```text
AAB + ABA + BAA = -3P_-.
```

## Sign-source audit

Gate 651 classifies the negative sign source as a typed candidate involving:

```text
S_K negative-sector insertion sign,
K_7^+ ∧ K_7^- orientation,
ordered interior-contraction sign,
antisymmetrization convention,
P_G-sourced octonionic calibration normalization.
```

This supports the finite sign/equal-unit pattern but does not yet supply a basis-free symbolic calibration theorem.

## Route universality audit

The same normalized calibration pattern is recorded for the repeated routes:

```text
omega_1_alt: AAA=+1, AAB=ABA=BAA=-1,
omega_2_alt: AAA=+1, AAB=ABA=BAA=-1,
omega_B_alt: AAA=+1, AAB=ABA=BAA=-1.
```

The result is route-universal after each route is normalized by its common coefficient `c`.

## Reconstruction audit

Using the surviving calibrated channels,

```text
g_twist = H_AAA + H_AAB + H_ABA + H_BAA
        = P_+ - 3P_-.
```

Therefore

```text
G_hat = (P_+ - 3P_-)/sqrt(31).
```

Against

```text
B_hat = (P_+ - P_-)/sqrt(7),
```

Gate 651 recovers

```text
cos(theta) = 13/sqrt(217),
rho^2 = 48/217.
```

## Theorem target

The sharpened theorem target is now:

```text
For the admissible S_K-twisted P_G-sourced tensor Omega=A+B
with A∈Lambda^{2,1} and B∈Lambda^{0,3}, the degree-allowed
Hitchin channel bilinear maps satisfy

AAA = +cP_+,
AAB = ABA = BAA = -cP_-,

therefore g_twist ∝ P_+ - 3P_-.
```

Gate 651 certifies the finite route-normalized calibration pattern, not the full basis-free proof.

## Final verdict

```text
PASS_GATE650_DEGREE_SELECTION_INHERITED
PASS_ORIENTATION_AND_VOLUME_CONVENTIONS_AUDITED
PASS_SURVIVING_CHANNEL_BILINEAR_MAPS_COMPUTED
PASS_AAA_POSITIVE_UNIT_AUDITED
PASS_AAB_ABA_BAA_NEGATIVE_EQUAL_UNIT_AUDITED
CONDITIONAL_SUPPORT_SURVIVING_CHANNELS_HAVE_EQUAL_UNIT_MAGNITUDE
CONDITIONAL_SUPPORT_NEGATIVE_SIGN_SOURCE_CLASSIFIED
PASS_RECONSTRUCTION_OF_P_PLUS_MINUS_THREE_P_MINUS_COMPUTED
CONDITIONAL_SUPPORT_HITCHIN_CHANNEL_CALIBRATION_IDENTITY_SHARPENED
FAILED_ROUTE_NO_FULL_SYMBOLIC_CALIBRATION_THEOREM
FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM
FAILED_ROUTE_HITCHIN_CHANNEL_CALIBRATION_IS_NOT_PHYSICAL_METRIC
FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM
FIREWALL_PRESERVED_GATE651_HITCHIN_CHANNEL_CALIBRATION_BOUNDARY
```
