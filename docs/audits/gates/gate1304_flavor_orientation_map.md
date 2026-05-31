# Gate 1304 — FlavorOrientationMap Construction Audit

## Purpose

Separate mass eigenvalues from mixing orientation.

## Result

For Hermitian Yukawa-square operators:

$$
H_u=U_uD_uU_u^\dagger,\qquad H_d=U_dD_dU_d^\dagger,
$$

masses come from $D_u,D_d$, while CKM is:

$$
V_{CKM}=U_u^\dagger U_d.
$$

Only the relative orientation is physical because common basis changes cancel.

## Locked CKM boundary

$$
\theta_{12}^0=1/4-18S+158S^2,
$$

$$
\theta_{23}^0=L+5S/3-(8-2L)S^2,
$$

$$
\theta_{13}^0=72LS-3S^2/2,
$$

$$
\delta_{CKM}^0=\pi/3+71S+93S^2/4.
$$

## Verdict

```text
PASS_RELATIVE_FLAVOR_ORIENTATION_THEOREM
CKM_RELATIVE_ORIENTATION_LOCKED_AS_BRIDGE_FORMULA
REMAINING_WOUND_MAJORANA_NEUTRINO_ORIENTATION_SOURCE
```
