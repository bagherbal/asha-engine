# Gate 602 — Unsealed Lepton-Wall / PMNS-Row Branch Selector Audit

## Purpose

Gate 601 showed that the environmental balance

```text
B_flav(sigma,i,s_J)
=
1 - 8*pi*epsilon_sigma(H_e)
- (1/4)Tr(P_eP_i^nu)
+ s_J J_CKM
```

selects the third neutrino projector `P_3^nu` and the positive CKM orientation sign, but does not select the charged-lepton ordering when every branch is already measured against the electron-zero wall. Gate 602 removes that preselection by allowing the wall label and PMNS row to vary.

This remains a bridge-layer branch-compatibility audit only. It does not derive Koide, charged-lepton masses, PMNS, CKM, neutrino data, flavor texture, or `B_flav=0` as ASHA-native law.

## Branch-row balance

Gate 602 evaluates:

```text
B_flav(sigma,alpha,i,s_J)
=
1
- 8*pi*epsilon_{sigma,alpha}(H_e)
- (1/4)Tr(P_alpha P_i^nu)
+ s_J J_CKM
```

over:

```text
sigma in six charged-lepton cyclic/permutation branches
alpha in {e,mu,tau}
i in {1,2,3}
s_J in {+1,-1}
```

## ChargedLeptonWallCandidateTable

For each branch, `delta_sigma` is recomputed and the distance to each component-zero wall is evaluated.

| sigma | alpha | delta deg | wall deg | epsilon deg | kappa | positive chamber |
|---|---|---:|---:|---:|---:|---:|
| `(e,mu,tau)` | e | 132.732819967108 | 135 | 2.267180032892 | 0.005503554191562 | true |
| `(e,mu,tau)` | mu | 132.732819967108 | 105 | 27.732819967108 | -11.1649760889774 | true |
| `(e,mu,tau)` | tau | 132.732819967108 | 255 | 122.267180032892 | -52.6323865849517 | true |
| `(e,tau,mu)` | e | 227.267180032892 | 225 | 2.267180032892 | 0.005503554191575 | true |
| `(e,tau,mu)` | mu | 227.267180032892 | 255 | 27.732819967108 | -11.1649760889774 | true |
| `(e,tau,mu)` | tau | 227.267180032892 | 105 | 122.267180032892 | -52.6323865849517 | true |
| `(mu,e,tau)` | e | 107.267180032892 | 105 | 2.267180032892 | 0.005503554191575 | true |
| `(mu,e,tau)` | mu | 107.267180032892 | 135 | 27.732819967108 | -11.1649760889774 | true |
| `(mu,e,tau)` | tau | 107.267180032892 | 345 | 122.267180032892 | -52.6323865849517 | true |

The remaining three permutations give the same wall-distance classes: electron wall near `2.26718°`, muon wall near `27.73282°`, and tau wall near `122.26718°`.

## PMNSRowProjectorOverlapTable

Using the inherited NuFIT 6.0 central values and standard PMNS convention:

```text
L_{alpha i} = (1/4)Tr(P_alpha P_i^nu) = |U_{alpha i}|^2/4
```

| alpha | i | `|U_{alpha i}|^2` | `L_{alpha i}` |
|---|---:|---:|---:|
| e | 1 | 0.6766722 | 0.16916805 |
| e | 2 | 0.3011778 | 0.07529445 |
| e | 3 | 0.02215 | 0.0055375 |
| mu | 1 | 0.112280287257248 | 0.028070071814312 |
| mu | 2 | 0.428130212742752 | 0.107032553185688 |
| mu | 3 | 0.4595895 | 0.114897375 |
| tau | 1 | 0.211047512742752 | 0.052761878185688 |
| tau | 2 | 0.270691987257248 | 0.067672996814312 |
| tau | 3 | 0.5182605 | 0.129565125 |

## CKMSignTable

| sign | value | convention |
|---:|---:|---|
| +1 | `+3.11699352875547e-05` | observed CKM orientation convention |
| -1 | `-3.11699352875547e-05` | reversed CKM orientation convention |

## FullBranchRowBalanceTable — leading rows

Sorted by `|B_flav|`:

| rank class | sigma | alpha | i | `s_J` | `B_flav` | `|B_flav|` |
|---:|---|---|---:|---:|---:|---:|
| 1 | `(e,tau,mu)` | e | 3 | +1 | -2.77587313788870e-06 | 2.77587313788870e-06 |
| 1 | `(mu,e,tau)` | e | 3 | +1 | -2.77587313788870e-06 | 2.77587313788870e-06 |
| 1 | `(e,mu,tau)` | e | 3 | +1 | -2.77587315043422e-06 | 2.77587315043422e-06 |
| 1 | `(mu,tau,e)` | e | 3 | +1 | -2.77587315043422e-06 | 2.77587315043422e-06 |
| 1 | `(tau,mu,e)` | e | 3 | +1 | -2.77587315043422e-06 | 2.77587315043422e-06 |
| 1 | `(tau,e,mu)` | e | 3 | +1 | -2.77587315343183e-06 | 2.77587315343183e-06 |
| next | `(e,tau,mu)` | e | 3 | -1 | -6.51157437129981e-05 | 6.51157437129981e-05 |
| next | `(mu,e,tau)` | e | 3 | -1 | -6.51157437129981e-05 | 6.51157437129981e-05 |

## ObservedTupleRank

```text
observed = (sigma=(e,mu,tau), alpha=e, i=3, s_J=+1)
rank = 1
minimal_class_size = 6
unique = false
```

## GapToNextBestTuple

```text
best |B_flav|          = 2.77587313788870e-06
next distinct |B_flav| = 6.51157437129981e-05
gap                    = 6.23398705751094e-05
```

The gap is after the sixfold charged-lepton sigma degeneracy. It distinguishes PMNS row/projector/sign choices, not the full charged-lepton cyclic ordering.

## RemainingDegeneracyLedger

The minimal class has:

```text
alpha = e only
i = 3 only
s_J = +1 only
sigma = all six charged-lepton cyclic/permutation branches
```

Therefore Gate 602 improves Gate 601: it shows that the balance does select the electron row once the wall label is unsealed. It still does not select the full charged-lepton sigma/cyclic ordering.

## Verdict

```text
PASS_BRANCH_ROW_BALANCE_FUNCTION_DEFINED
PASS_LEPTON_WALLS_PMNS_ROWS_AND_CKM_SIGNS_ENUMERATED
PASS_CHARGED_LEPTON_ZERO_WALL_CANDIDATES_ENUMERATED
PASS_PMNS_ROW_PROJECTOR_OVERLAPS_ENUMERATED
PASS_CKM_ORIENTATION_SIGNS_ENUMERATED
PASS_FULL_BRANCH_ROW_BALANCE_TABLE_COMPUTED
PASS_OBSERVED_TUPLE_IN_MINIMAL_RESIDUAL_CLASS
CONDITIONAL_SUPPORT_BALANCE_SELECTS_ELECTRON_ROW
CONDITIONAL_SUPPORT_BALANCE_SELECTS_P3_NU_AND_POSITIVE_CKM_SIGN
CONDITIONAL_SUPPORT_CHARGED_LEPTON_SIGMA_DEGENERACY_REMAINS
FAILED_ROUTE_FULL_CHARGED_LEPTON_ORDERING_NOT_UNIQUELY_SELECTED
FAILED_ROUTE_NO_NATIVE_BRANCH_SELECTION_THEOREM
FAILED_ROUTE_NO_NATIVE_B_FLAV_ZERO_THEOREM
FAILED_ROUTE_B_FLAV_REMAINS_ENVIRONMENTAL_BRANCH_ROW_COMPATIBILITY_TEST
FIREWALL_PRESERVED_GATE600_BRANCH_CHAMBER_BOUNDARY
FIREWALL_PRESERVED_GATE601_BRANCH_COMPATIBILITY_BOUNDARY
FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_GATE596_FOURTH_ROOT_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_NO_KOIDE_DERIVATION
FIREWALL_PRESERVED_NO_CHARGED_LEPTON_MASS_DERIVATION
FIREWALL_PRESERVED_NO_PMNS_CKM_NEUTRINO_OR_FLAVOR_DERIVATION
FIREWALL_PRESERVED_OBSERVED_BRANCH_LABELS_REMAIN_ENVIRONMENTAL_DATA
FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED
FIREWALL_PRESERVED_GATE602_UNSEALED_LEPTON_WALL_BOUNDARY
```

## Interpretation

`B_flav` is now a stronger environmental history compatibility filter. When the lepton wall is unsealed, the near-zero balance identifies the electron row, the third neutrino projector, and the positive CKM orientation sign. However, it still leaves the sixfold charged-lepton sigma/cyclic-order degeneracy unresolved. A separate chamber-orientation or discriminant-orientation seal would still be needed for a unique full history branch.
