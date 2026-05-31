# Gate 1336 — VacuumBoundary and MajoranaPhase Selector Audit

## Purpose

Separate boundary phase and vacuum quantities from mass spectra.

## Result

Spectra and ordinary oscillation data do not determine:

$$
\Lambda_{cosmo},\qquad \delta_{PMNS},\qquad \alpha_2,\qquad \alpha_3.
$$

The cosmological constant is an independent vacuum-density coefficient. Majorana phases cancel from oscillation probabilities but enter lepton-number-violating observables such as:

$$
m_{\beta\beta}=\left|\sum_i m_iU_{ei}^2\right|.
$$

For rank-2 normal ordering:

$$
m_{\beta\beta}=\left|m_2s_{12}^2c_{13}^2e^{i\alpha}+m_3s_{13}^2e^{-2i\delta}\right|.
$$

## Verdict

```text
PASS_BOUNDARY_SPECTRUM_INDEPENDENCE_THEOREM
PASS_VACUUM_ZERO_INDEPENDENCE_THEOREM
REMAINING_WOUND_BOUNDARY_PHASE_SELECTOR
```
