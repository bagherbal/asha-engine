# ASHA Phenomenology Package — Empirical-Seal Prediction Audit

## Bottom line

The empirical-seal layer computes conditional consequences. It predicts ASHA+empirical vacuum metastability in the simple one-loop audit, rules out an unsuppressed stable thermal B-gap relic by overclosure, and confirms that the cosmological constant still needs an environmental subtraction theorem.

## Environmental quarantine inputs

| Input | Value |
|---|---:|
| Top mass | 172.69 GeV |
| Higgs mass | 125.25 GeV |
| alpha_s(mZ) | 0.1179 |
| Z mass | 91.1876 GeV |
| Target Omega_c h^2 | 0.120 |
| Target rho_Lambda/M_Pl^4 | 1.0e-120 |

## Vacuum fate: one-loop conditional audit

| Seed mode | y_t start | λ before threshold | λ after threshold | instability scale | λ_min | log10 lifetime/yr |
|---|---:|---:|---:|---:|---:|---:|
| tree-pole-top-seed | 0.991879377 | -0.006880641 | -0.104727433 | 5.767333e+05 GeV | -0.122793908 | 55.642486 |
| one-loop-QCD-MSbar-top-seed | 0.942247405 | 0.032625460 | -0.065221333 | 1.467750e+06 GeV | -0.077446343 | 109.740890 |

Default precision lane: `one-loop-QCD-MSbar-top-seed`. Initial λ(mZ)=0.129383477, g3(mZ)=1.217199694, bounce action proxy=339.834575, log10(age/yr)=10.139879.

**Verdict:** conditional metastability = true. conditional one-loop audit; precision vacuum fate requires full MSbar matching, 2/3-loop beta functions, pole-to-running conversion and threshold convention

## Dark matter: B-gap Majorana constraint

| Quantity | Value |
|---|---:|
| candidate mass | 1.467750e+06 GeV |
| required yield Y=n/s for Omega=0.120 | 2.979811e-16 |
| relativistic thermal yield, g=2, g*=106.75 | 3.901498e-03 |
| stable thermal Omega h^2 | 1.571173e+12 |
| overclosure factor | 1.309311e+13 |
| required fraction of thermal yield | 7.637606e-14 |

**Verdict:** A stable thermal relic at the ASHA B-gap mass grossly overcloses the universe; matching Planck requires an extremely suppressed/nonthermal yield or entropy dilution. Reheating temperature, production channel and decay/stability theorem remain empirical seals.

## Cosmological constant: subtraction severity

| Quantity | Value |
|---|---:|
| convention | diagnostic leading CCM bare vacuum, f4=1 and Lambda=M_Pl |
| bare rho / M_Pl^4 | 4.863417e+00 |
| target rho_Lambda / M_Pl^4 | 1.000000e-120 |
| required counterterm | 4.863417e+00 |
| cancellation ratio | 4.863417e+120 |
| decimal digits of cancellation | 120.687 |

**Verdict:** ASHA supplies a bare spectral-action cosmological term but not a vacuum subtraction theorem; matching the observed scale requires a counterterm cancellation at roughly 121 decimal digits under this convention.

## Status ledger

- `CONDITIONAL_SUPPORT_PHENOMENOLOGY_LAYER_EXECUTED`
- `CONDITIONAL_SUPPORT_EMPIRICAL_QUARANTINE_STATE_LOADED`
- `CONDITIONAL_SUPPORT_ONE_LOOP_VACUUM_RG_EXECUTED`
- `CONDITIONAL_SUPPORT_ASHA_THRESHOLD_JUMP_APPLIED`
- `CONDITIONAL_TENSION_OUTPUTS_DEPEND_ON_EMPIRICAL_SEALS`
- `CONDITIONAL_SUPPORT_CONDITIONAL_VACUUM_METASTABILITY_FOUND`
- `FAILED_ROUTE_UNIVERSE_LIFETIME_NOT_NATIVELY_PREDICTED`
- `CONDITIONAL_SUPPORT_DARK_MATTER_REQUIRED_YIELD_COMPUTED`
- `CONDITIONAL_SUPPORT_THERMAL_STABLE_MAJORANA_OVERCLOSURE_COMPUTED`
- `FAILED_ROUTE_DARK_MATTER_ABUNDANCE_NOT_NATIVELY_PREDICTED`
- `CONDITIONAL_SUPPORT_COSMOLOGICAL_CONSTANT_FINE_TUNING_COMPUTED`
- `FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_ORGANICALLY_SOLVED`
