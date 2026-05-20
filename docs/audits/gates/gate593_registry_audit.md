# Gate 593 — OrientationBalance Invariant Matrix Form Audit

## Purpose

Gate 593 continues from Gate 592.  Gates 590-592 established the strongest current bridge-layer environmental relation:

```text
1 - 8*pi*epsilon_e ≈ sin²(theta13)/4 - J_CKM.
```

Gate 592 showed that no additional residual fit is justified at current v1 precision and isolated the missing object as a `CrossSectorOrientationIntertwiner`.  Gate 593 does not search for new numbers.  It rewrites the relation in invariant matrix/projector form so the missing operator target becomes precise.

This remains bridge-layer environmental geometry.  It does not derive Koide, PMNS, CKM, Yukawa spectra, neutrino physics, or flavor texture from ASHA-native law.

## Inherited OrientationBalance relation

Observed charged-lepton wall coordinate:

```text
epsilon_e = 0.039569756309433 rad
          = 2.26718003289167°
```

Loop-angle deficit:

```text
kappa_e = 1 - 8*pi*epsilon_e
        = 0.00550355419157456.
```

Gate 590 / 592 candidate:

```text
sin²(theta13)/4 - J_CKM = 0.00550633006471245.
```

Residual:

```text
Delta_590 = 2.77587313788925e-06.
```

The equivalent epsilon prediction is:

```text
epsilon_pred = (1/(8*pi)) [1 - sin²(theta13)/4 + J_CKM]
             = 0.0395696458609502 rad.
```

This residual is already inside the propagated uncertainty band inherited from Gate 591.

## Charged-lepton root-space side

Gate 593 records the left side as a functional of the charged-lepton Yukawa singular values:

```text
Y_e -> spec_sing(Y_e) = (y_e,y_mu,y_tau)
x_e = (sqrt(y_e), sqrt(y_mu), sqrt(y_tau))
```

In the Fourier Koide chamber:

```text
x_j = A[1 + sqrt(2) R cos(delta + 2*pi*j/3)]
```

with canonical chamber ordering:

```text
(e, mu, tau)
```

and chamber-wall coordinate:

```text
epsilon(Y_e) = 135° - delta.
```

Therefore:

```text
kappa(Y_e) = 1 - 8*pi*epsilon(Y_e).
```

This map requires a root-spectrum operation and remains blocked natively by Gate 352:

```text
FAILED_ROUTE_ROOT_SPECTRUM_MAP_NOT_NATIVE_GATE352
FAILED_ROUTE_NO_NATIVE_EPSILON_OF_YE_OPERATOR
```

## PMNS projector-trace side

The PMNS reactor term is rewritten as:

```text
sin²(theta13)
= |U_PMNS[e,3]|²
= Tr(P_e U_PMNS P_3^nu U_PMNS†).
```

Required labels/seals:

```text
P_e      = electron flavor projector
P_3^nu   = third neutrino mass-eigenstate projector
PMNS convention
mass ordering seal
```

Thus:

```text
sin²(theta13)/4
=
(1/4) Tr(P_e U_PMNS P_3^nu U_PMNS†).
```

This is a version-pinned observed projector trace, not a native ASHA PMNS operator.

## CKM Jarlskog side

The CKM orientation term is recorded as the rephasing-invariant area:

```text
J_CKM = Im(V_us V_cb V_ub* V_cs*).
```

Gate 593 also records the basis-invariant commutator form:

```text
det([H_u,H_d])
=
2 i J_CKM
prod_{i<j}(y_{u_i}²-y_{u_j}²)
prod_{i<j}(y_{d_i}²-y_{d_j}²),
```

up to sign and convention, with:

```text
H_u = Y_u Y_u†
H_d = Y_d Y_d†.
```

The magnitude is rephasing invariant, while the signed orientation requires generation-ordering and CKM-orientation conventions.

## Invariant OrientationBalance form

Gate 593 rewrites the current environmental bridge as:

```text
kappa_e
?=
(1/4) Tr(P_e U_PMNS P_3^nu U_PMNS†)
-
J_CKM.
```

Equivalently:

```text
1 - 8*pi*epsilon(Y_e)
?=
(1/4) Tr(P_e U_PMNS P_3^nu U_PMNS†)
-
J(Y_u,Y_d).
```

Or in epsilon form:

```text
epsilon(Y_e)
?=
(1/(8*pi))
[
  1
  - (1/4) Tr(P_e U_PMNS P_3^nu U_PMNS†)
  + J(Y_u,Y_d)
].
```

Numerically this is the same Gate 590 relation:

```text
left kappa_e             = 0.00550355419157456
right projector-minus-J  = 0.00550633006471245
residual                 = 2.77587313788925e-06.
```

## Basis and label dependence

| object | invariant part | required label/seal |
|---|---|---|
| `epsilon(Y_e)` | root-spectrum Koide chamber coordinate | charged-lepton ordering chamber `(e,mu,tau)` and electron-zero wall |
| `P_e` | rank-one flavor projector after flavor basis is selected | electron flavor label |
| `P_3^nu` | rank-one neutrino mass projector | third neutrino mass eigenstate and mass ordering |
| `J_CKM` | rephasing-invariant oriented area magnitude | quark generation ordering and CKM orientation sign |
| `1/(8*pi)` | loop-sized angular unit | bridge normalization convention |

Gate 593 therefore makes the required labels explicit instead of hiding them inside angle notation.

## Current ASHA availability audit

| object | observed ledger? | native operator? | can supply the balance? |
|---|---:|---:|---:|
| charged-lepton root-space `epsilon(Y_e)` | yes | no | no |
| PMNS projector trace `Tr(P_e U P_3 U†)` | yes | no | no |
| CKM Jarlskog / commutator area | yes | no | no |
| finite spectral triple `D_F` and one-form edges | no | yes | no |
| quaternionic weak socket `H` / `Im(H)` | no | yes | no |
| `CrossSectorOrientationIntertwiner` | no | no | no |

The observed ledgers supply the ingredients.  They do not supply the native operator.

## Minimal missing operator target

Gate 593 defines the precise operator target:

```text
CrossSectorOrientationIntertwiner
```

with domain:

```text
charged-lepton root-space chamber functional epsilon(Y_e),
PMNS reactor projector trace Tr(P_e U_PMNS P_3^nu U_PMNS†),
CKM Jarlskog commutator area J(Y_u,Y_d),
loop angular unit 1/(8*pi).
```

and scalar residual:

```text
I_orient
:=
1 - 8*pi*epsilon(Y_e)
- (1/4) Tr(P_e U_PMNS P_3^nu U_PMNS†)
+ J(Y_u,Y_d).
```

The native theorem target would be:

```text
I_orient = 0.
```

Current ASHA does not contain such an operator.

## Verdict

```text
PASS_GATE592_ORIENTATION_BALANCE_SEAL_INHERITED
PASS_CHARGED_LEPTON_ROOT_SPACE_EPSILON_MAP_DEFINED
PASS_PMNS_REACTOR_PROJECTOR_TRACE_FORM_DEFINED
PASS_CKM_JARLSKOG_REPHASING_AND_COMMUTATOR_FORM_RECORDED
PASS_ORIENTATION_BALANCE_INVARIANT_MATRIX_FORM_WRITTEN
PASS_BASIS_AND_LABEL_DEPENDENCE_AUDITED
PASS_CURRENT_ASHA_AVAILABILITY_AUDITED
CONDITIONAL_SUPPORT_INVARIANT_FORM_SHARPENS_OPERATOR_TARGET
CONDITIONAL_SUPPORT_ORIENTATIONBALANCESEAL_REWRITTEN_AS_PROJECTOR_COMMUTATOR_BALANCE
FAILED_ROUTE_ROOT_SPECTRUM_MAP_NOT_NATIVE_GATE352
FAILED_ROUTE_NO_NATIVE_EPSILON_OF_YE_OPERATOR
FAILED_ROUTE_PMNS_PROJECTOR_TRACE_IS_OBSERVED_LEDGER_NOT_NATIVE_OPERATOR
FAILED_ROUTE_NO_NATIVE_FLAVOR_COMMUTATOR_TO_KOIDE_WALL_MAP
FAILED_ROUTE_NO_CROSS_SECTOR_TRACE_ORIENTATION_BALANCE_OPERATOR
FAILED_ROUTE_NO_CROSS_SECTOR_ORIENTATION_INTERTWINER
FAILED_ROUTE_ORIENTATIONBALANCESEAL_REMAINS_ENVIRONMENTAL
FIREWALL_PRESERVED_NO_KOIDE_PMNS_CKM_YUKAWA_NEUTRINO_OR_FLAVOR_DERIVATION
FIREWALL_PRESERVED_PMNS_CKM_AND_YUKAWA_INPUTS_REMAIN_OBSERVED_ENVIRONMENTAL_DATA
FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED
FIREWALL_PRESERVED_GATE593_ORIENTATION_BALANCE_INVARIANT_FORM_BOUNDARY
```

Gate 593 therefore turns the Gate 592 symbolic missing object into an exact mathematical target: a native cross-sector trace/orientation balance operator whose shadow would be the current environmental `OrientationBalanceSeal`.
