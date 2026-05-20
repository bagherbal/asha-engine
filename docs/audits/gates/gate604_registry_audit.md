# Gate 604 — Minimal Flavor History Branch Seal Closure Audit

## Purpose

Gate 604 closes the flavor branch analysis after Gate 603.  Gate 602/603 showed that the environmental balance

```text
B_flav = 1 - 8*pi*epsilon_e - (1/4)Tr(P_e P_3^nu) + J_CKM
```

selects the electron wall, the third neutrino projector, and the positive CKM orientation sign, while the remaining charged-lepton `sigma`/cyclic Fourier presentation is invisible to the balance.  Gate 604 identifies the minimal flavor-history branch seal required for the ASHA history-transport ledger without promoting the balance to native law.

## Branch stack

```text
native trace ring
  R_e = Q[Tr(H_e), Tr(H_e^2), Tr(H_e^3)]
  chi_e(lambda)

algebraic extension
  cubic splitting field K_e
  positive fourth-root sheets x_i=lambda_i^(1/4)

environmental branch seal
  electron wall / epsilon_e
  PMNS electron-to-third-neutrino projector overlap
  positive CKM orientation +J_CKM

gauge/convention layer
  sixfold sigma/cyclic Fourier presentation
  PMNS and CKM phase conventions
```

## Minimality result

Required for `B_flav`:

```text
positive fourth-root charged-lepton branch
electron-wall coordinate epsilon_e
Tr(P_e P_3^nu)=|U_e3|^2
positive CKM orientation +J_CKM
OrientationBalanceSeal B_flav≈0
```

Not required for `B_flav`:

```text
full charged-lepton cyclic sigma
signed Vandermonde orientation
exact Fourier presentation choice
```

The optional full-order seal is:

```text
ChargedLeptonDiscriminantOrientationSeal:
  sign(V_e)=sign(prod_{i<j}(lambda_j-lambda_i))
  or sign(V_x)=sign(prod_{i<j}(x_j-x_i))
  full cyclic order of (e,mu,tau)
```

This optional seal is needed only if full ordered-history reconstruction is demanded.  It is not needed by the current `B_flav` balance.

## Updated flavor transport form

```text
E_flavor(M_Z)
=
T_flavor[
  native trace ring R_e,
  MinimalFlavorHistoryBranchSeal,
  OrientationBalanceSeal,
  remaining raw Yukawa/PMNS/CKM inputs
]
```

with:

```text
Y_core:
  R_e and chi_e(lambda)
  positive fourth-root charged-lepton branch
  epsilon_e
  Tr(P_eP_3^nu)
  J(H_u,H_d)
  remaining Yukawa singular values as observed ledgers

Omega_core:
  electron wall alpha=e
  third neutrino projector P_3^nu
  positive CKM orientation sign
  PMNS/CKM convention labels
  optional signed Vandermonde orientation only for full order
```

## Verdict

```text
PASS_FLAVOR_HISTORY_BRANCH_STACK_CONSTRUCTED
PASS_MINIMAL_FLAVOR_HISTORY_BRANCH_SEAL_DEFINED
PASS_SIGMA_CLASSIFIED_AS_GAUGE_LIKE_FOR_B_FLAV
CONDITIONAL_SUPPORT_OPTIONAL_DISCRIMINANT_ORIENTATION_SEAL_FOR_FULL_ORDER
CONDITIONAL_SUPPORT_B_FLAV_ACTS_AS_ENVIRONMENTAL_BRANCH_COMPATIBILITY_FILTER
FAILED_ROUTE_NO_NATIVE_BRANCH_SELECTION_THEOREM
FAILED_ROUTE_NO_NATIVE_B_FLAV_ZERO_THEOREM
FAILED_ROUTE_NO_NATIVE_FOURTH_ROOT_THEOREM
FIREWALL_PRESERVED_GATE604_MINIMAL_FLAVOR_HISTORY_BRANCH_SEAL_BOUNDARY
```

## Firewalls

Gate 604 does not derive Koide, charged-lepton masses, Yukawa eigenvalues, PMNS, CKM, neutrino data, flavor texture, or `B_flav=0`.  It does not add a carrier or selector.  It preserves the Gate 352 root-trace obstruction, the Gate 596 fourth-root obstruction, the Gate 599 trace-ring extension boundary, and the Gate 603 sigma gauge/orientation boundary.
