# Gate 601 — Flavor Branch-Compatibility Selector Audit

## Purpose

Gate 600 showed that `epsilon(H_e)` is branch-algebraic over the charged-lepton trace ring only after choosing a cubic splitting branch, positive fourth roots, a charged-lepton chamber ordering, Fourier cyclic orientation, and an electron-zero wall. Gate 601 asks whether the environmental balance

```text
B_flav(sigma,i,s_J)
=
1 - 8*pi*epsilon_sigma(H_e)
- (1/4) Tr(P_e P_i^nu)
+ s_J J_CKM
```

acts as a branch-compatibility selector over charged-lepton chamber permutations, neutrino projector choices, and CKM orientation signs.

This is bridge-layer only. It does not derive Koide, charged-lepton masses, PMNS, CKM, neutrino data, or `B_flav=0`.

## Inherited balance

```text
epsilon_obs = 0.039569756309433 rad
kappa_e     = 1 - 8*pi*epsilon_obs = 0.00550355419157456
J_CKM       = 3.1169935287554706e-05
```

PMNS reactor data inherited from the NuFIT 6.0 Normal Ordering ledger:

```text
sin^2(theta12) = 0.308
sin^2(theta13) = 0.02215
```

## ChargedLeptonBranchTable

For each charged-lepton permutation, the Fourier phase is recomputed and `epsilon_sigma` is measured to the electron-zero wall in that chamber.

| sigma | delta_sigma_deg | electron wall deg | epsilon_sigma_deg | kappa_sigma | positive Koide chamber |
|---|---:|---:|---:|---:|---:|
| `(e,mu,tau)` | 132.732819967108 | 135 | 2.267180032892 | 0.005503554191575 | true |
| `(e,tau,mu)` | 227.267180032892 | 225 | 2.267180032892 | 0.005503554191562 | true |
| `(mu,e,tau)` | 107.267180032892 | 105 | 2.267180032892 | 0.005503554191550 | true |
| `(mu,tau,e)` | 252.732819967108 | 255 | 2.267180032892 | 0.005503554191562 | true |
| `(tau,e,mu)` | 12.732819967108 | 15 | 2.267180032892 | 0.005503554191563 | true |
| `(tau,mu,e)` | 347.267180032892 | 345 | 2.267180032892 | 0.005503554191562 | true |

The branch balance sees the same electron-wall offset for all six charged-lepton permutations. Therefore it cannot uniquely select the charged-lepton ordering.

## PMNSProjectorOverlapTable

```text
L_i = (1/4) Tr(P_e P_i^nu) = |U_ei|^2/4
```

| i | projector | `|U_ei|^2` | `L_i` |
|---:|---|---:|---:|
| 1 | `P_1^nu` | 0.6766722 | 0.16916805 |
| 2 | `P_2^nu` | 0.3011778 | 0.07529445 |
| 3 | `P_3^nu` | 0.02215 | 0.0055375 |

## CKMSignTable

| sign | value | convention |
|---:|---:|---|
| `+1` | `+3.11699352875547e-05` | observed CKM orientation convention |
| `-1` | `-3.11699352875547e-05` | reversed CKM orientation convention |

## FullBranchBalanceTable — leading rows

Sorted by `|B_flav|`:

| rank class | sigma | i | `s_J` | `B_flav` | `|B_flav|` |
|---:|---|---:|---:|---:|---:|
| 1 | `(e,mu,tau)` | 3 | +1 | -2.77587313788957e-06 | 2.77587313788957e-06 |
| 1 | `(tau,e,mu)` | 3 | +1 | -2.77587314954691e-06 | 2.77587314954691e-06 |
| 1 | `(e,tau,mu)` | 3 | +1 | -2.77587315043509e-06 | 2.77587315043509e-06 |
| 1 | `(mu,tau,e)` | 3 | +1 | -2.77587315043509e-06 | 2.77587315043509e-06 |
| 1 | `(tau,mu,e)` | 3 | +1 | -2.77587315043509e-06 | 2.77587315043509e-06 |
| 1 | `(mu,e,tau)` | 3 | +1 | -2.77587316286959e-06 | 2.77587316286959e-06 |
| next | `(e,mu,tau)` | 3 | -1 | -6.51157437129990e-05 | 6.51157437129990e-05 |

The observed branch belongs to the minimal residual class, but the class has a sixfold charged-lepton permutation degeneracy.

## ObservedBranchRank

```text
observed = (sigma=(e,mu,tau), i=3, s_J=+1)
rank = 1
minimal_class_size = 6
unique = false
```

## GapToNextBestBranch

```text
best |B_flav|          = 2.77587313788957e-06
next distinct |B_flav| = 6.51157437129990e-05
gap                    = 6.23398705751094e-05
```

This gap is large compared with the residual itself, but it is the gap after the charged-lepton permutation tie. The balance selects `P_3^nu` and positive CKM orientation, not the full charged-lepton branch.

## Verdict

```text
PASS_BRANCH_BALANCE_FUNCTION_DEFINED
PASS_BRANCH_SPACE_ENUMERATED
PASS_CHARGED_LEPTON_BRANCHES_ENUMERATED
PASS_PMNS_PROJECTOR_CHOICES_ENUMERATED
PASS_CKM_ORIENTATION_SIGNS_ENUMERATED
PASS_FULL_BRANCH_BALANCE_TABLE_COMPUTED
PASS_OBSERVED_BRANCH_IN_MINIMAL_RESIDUAL_CLASS
CONDITIONAL_SUPPORT_BALANCE_SELECTS_THIRD_NEUTRINO_PROJECTOR_AND_POSITIVE_CKM_SIGN
CONDITIONAL_SUPPORT_CHARGED_LEPTON_PERMUTATION_DEGENERACY_EXPOSED
FAILED_ROUTE_BRANCH_SELECTOR_NOT_UNIQUE
FAILED_ROUTE_BALANCE_DOES_NOT_UNIQUELY_SELECT_CHARGED_LEPTON_ORDERING
FAILED_ROUTE_NO_NATIVE_BRANCH_SELECTION_THEOREM
FAILED_ROUTE_NO_NATIVE_B_FLAV_ZERO_THEOREM
FAILED_ROUTE_B_FLAV_REMAINS_ENVIRONMENTAL_BRANCH_COMPATIBILITY_TEST
FIREWALL_PRESERVED_GATE600_BRANCH_CHAMBER_BOUNDARY
FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_GATE596_FOURTH_ROOT_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_NO_KOIDE_DERIVATION
FIREWALL_PRESERVED_NO_CHARGED_LEPTON_MASS_DERIVATION
FIREWALL_PRESERVED_NO_PMNS_CKM_NEUTRINO_OR_FLAVOR_DERIVATION
FIREWALL_PRESERVED_OBSERVED_BRANCH_LABELS_REMAIN_ENVIRONMENTAL_DATA
FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED
FIREWALL_PRESERVED_GATE601_BRANCH_COMPATIBILITY_SELECTOR_BOUNDARY
```

## Interpretation

`B_flav` is a strong environmental branch-compatibility filter: it strongly favors the third neutrino projector and the observed positive CKM orientation sign. However, once `epsilon_sigma` is measured relative to the physical electron-zero wall, all six charged-lepton permutations share the same wall offset. Therefore Gate 601 does not produce a unique history branch selector. A further chamber-ordering principle would still be required.
