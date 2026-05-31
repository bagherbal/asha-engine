# Gate 1315 — MajoranaMatrix Structure Audit

## Purpose

Source-type the neutrino seesaw, rank-2 normal ordering, and Majorana orientation wound.

## Result

The seesaw bridge is:

$$
M_\nu=-\frac{v^2}{2}Y_\nu^TM_R^{-1}Y_\nu.
$$

Takagi decomposition is required:

$$
U_\nu^TM_\nu U_\nu=\operatorname{diag}(m_1,m_2,m_3).
$$

Rank-2 normal lane:

$$
m_1\approx0,\qquad m_2=(4L+10S)m_3.
$$

Heavy scale:

$$
M_{R3}=(\sqrt{2\pi}+49S+90S^2)\sqrt{vM_P}.
$$

Second heavy scale:

$$
M_{R2}=M_{R3}\frac{e^{-4\pi/3}}{4L+10S}.
$$

## Verdict

```text
PASS_MAJORANA_TAKAGI_DECOMPOSITION_THEOREM
PASS_RANK2_MINIMAL_SEESAW_THEOREM
PASS_MAJORANA_MATRIX_CONTROLS_PMNS_ORIENTATION
REMAINING_WOUND_MAJORANA_SELECTOR_AND_PHASES
```
