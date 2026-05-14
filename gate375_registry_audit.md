# Gate 375 Registry Audit — Cosmological Observables & Dark Sector Prediction Sieve

## Gate identity

| Field | Value |
|---|---|
| Gate | 375 |
| Package | `pkg/bridge/cosmologicalobservables` |
| Theorem | `CosmologicalObservablesDarkSectorPredictionSieveTheorem()` |
| Audit ID | `GATE375-COSMOLOGICAL-OBSERVABLES-DARK-SECTOR-PREDICTION-SIEVE` |
| Layer | Bridge / Cosmological Observables |

## Purpose

Gate 374 sealed the finite ASHA/`Cℓ(1,7)` program as a scoped kinematic and boundary reconstruction, while preserving the irreducible census

```text
15 = 13 charged finite-Dirac flavor moduli + theta_QCD + one absolute scale
```

Gate 375 deliberately reopens the program in the only admissible direction: **observable cosmology**.  The question is not whether ASHA has suggestive heavy-sector structures.  It clearly does.  The question is whether the current ledger is already sufficient to compute hard macroscopic predictions:

1. dark-matter relic abundance `Ω_DM h²`;
2. vacuum-decay bounce action and universe lifetime;
3. dark-energy density / cosmological constant.

The gate enforces the rule:

> A cosmological observable counts as derived only if every required physical input is fixed by the ASHA ledger.  Missing rates, couplings, initial conditions, RG trajectories, counterterms, or saturation postulates cannot be replaced by observed values.

## Inherited facts

| Inherited item | Value | Status |
|---|---:|---|
| Gate 374 scoped closure | finite kinematics closed; flavor vacuum unselected | inherited |
| Charged finite-Dirac moduli | `13` | inherited from Gate 372/374 |
| External minimal ledger | `15 = 13 + theta_QCD + scale` | inherited |
| `sin²θ_W(Λ)` | `3/8` | native boundary ratio |
| `α_GUT^-1` branch | `8π ≈ 25.1327412287` | native boundary branch |
| `λ_H/g_*²` | `1197/4624 ≈ 0.258866782` | native boundary ratio |
| `v/M_P` | `2^(3/2) exp(-4π²) ≈ 2.02416e-17` | hierarchy scale |
| PeV threshold lane | `1.46774973718e6 GeV` | conditional threshold transport lane |
| Sealed intermediate scale ledger | `6.650726476871e11 GeV` | earlier intermediate/seesaw ledger |
| Required quartic jump witness | `Δλ ≈ -0.0978` | transport/threshold witness, not full heavy-sector theorem |

## Candidate observable lanes

| Lane | Observable | Required model ingredients | Current ASHA status | Verdict |
|---|---|---|---|---|
| A | Heavy Majorana/B-gap dark matter relic density | particle identity, mass, stability, annihilation or decay rates, reheating history, entropy dilution, closed Boltzmann kernel | candidate scales exist, but rates/history/stability are not derived | failed as hard prediction |
| B | Higgs-vacuum lifetime | full `λ(μ)` RG trajectory, top/Yukawa data, threshold matching, `λ_min`, bounce scale, prefactor, gravity corrections | boundary ratio and jump witness exist, but no complete RG/bounce kernel | failed as hard prediction |
| C | Dark-energy density / cosmological constant | renormalized vacuum-energy functional, counterterm, sign, normalization, equality/saturation theorem | hierarchy suppressions exist, but no Λ functional or counterterm | failed as hard prediction |

## Lane A — Dark matter relic density

Formal Boltzmann skeleton:

```text
dY/dx = - s <σv>/(H x) (Y² - Y_eq²) + source/decay terms
Ω h² ∝ m Y_∞
```

The current ledger supplies only candidate scale data.  It does **not** supply the complete kernel.

| Input | Native? | Blocks `Ω_DM h²`? | Note |
|---|---:|---:|---|
| Dark particle identity | no | yes | Heavy Majorana/B-gap sector is a candidate semantic region, not a stable relic theorem. |
| Physical relic mass | no | yes | PeV threshold and `M_int` are scale ledgers, not a cosmological particle mass theorem. |
| Stability or lifetime | no | yes | No symmetry or width theorem proves survival to cosmological time. |
| `<σv>` | no | yes | Requires continuum dark-sector interactions and matrix elements. |
| Decay width `Γ` | no | yes | Requires couplings, final states, and mixing angles. |
| Reheating / production history | no | yes | Thermal freeze-out, freeze-in, or nonthermal production is not selected. |
| Entropy dilution / `g_*` history | no | yes | Requires cosmological thermal history. |

### Lane A truth

```text
FAILED_ROUTE_DARK_MATTER_RELIC_DENSITY_NOT_DERIVED
FAILED_ROUTE_OMEGA_DM_H2_NOT_COMPUTED_FROM_ASHA_LEDGER
FAILED_ROUTE_STABLE_DARK_SECTOR_CANDIDATE_NOT_DERIVED
FAILED_ROUTE_BOLTZMANN_KERNEL_NOT_DERIVED
```

A heavy scale by itself is not a relic-density prediction.

## Lane B — Vacuum metastability and universe lifetime

Formal bounce skeleton:

```text
S_E ≈ 8π² / (3 |λ_min|)
Γ/V ≈ μ_B⁴ exp(-S_E)
```

The ASHA ledger has the UV quartic boundary and a threshold-jump witness.  That is not enough to compute a lifetime.

| Input | Native? | Blocks lifetime? | Note |
|---|---:|---:|---|
| `λ_H/g_*² = 1197/4624` | yes | no by itself | Boundary ratio exists. |
| `Δλ ≈ -0.0978` | conditional | no by itself | Threshold-jump witness/target, not a full matching theorem. |
| Full `βλ, βyt, βg` trajectory | no | yes | Requires continuum RG plus threshold schedule. |
| Top/Yukawa transport | no | yes | Gate 372 leaves charged flavor moduli unselected. |
| `λ_min` and instability scale | no | yes | Bounce action depends on the most negative running quartic. |
| Bounce prefactor and gravitational corrections | no | yes | Needed to convert action into lifetime. |

### Lane B truth

```text
FAILED_ROUTE_UNIVERSE_LIFETIME_NOT_DERIVED
FAILED_ROUTE_EUCLIDEAN_BOUNCE_ACTION_NOT_COMPUTED
```

The gate does not deny metastability physics.  It says ASHA has not yet derived the complete continuum RG/bounce functional required to output a universe lifetime.

## Lane C — Cosmological constant / dark energy

Native hierarchy suppressions are available:

| Expression | Native value | Predicts `Λ_cosmo`? | Reason |
|---|---:|---:|---|
| `(v/M_P)^2` | `≈ 4.097e-34` | no | hierarchy-squared scale, not vacuum-energy theorem |
| `(v/M_P)^4` | `≈ 1.6787e-67` | no | dimensional scaling only; needs counterterm/sign/normalization |
| `exp(-8π²)` | `≈ 5.1225e-35` | no | instanton-like suppression, not uniquely dark energy |

A cosmological constant prediction requires a renormalized vacuum-energy functional and counterterm.  Gate 373 already proved that the holographic/Bekenstein lanes are aggregate inequalities unless a saturation theorem is derived.  Gate 375 preserves that firewall.

### Lane C truth

```text
FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_DERIVED
FAILED_ROUTE_DARK_ENERGY_DENSITY_NOT_COMPUTED
```

## Observable census

| Quantity | Value |
|---|---:|
| Requested cosmological observables | `3` |
| Hard predictions derived | `0` |
| Dark matter predictions | `0` |
| Vacuum lifetime predictions | `0` |
| Cosmological constant predictions | `0` |
| Remaining charged moduli | `13` |

## Firewalls

| Firewall | Preserved? |
|---|---:|
| No observed `Ω_DM h²` fitted | yes |
| No observed dark-energy density fitted | yes |
| No universe-lifetime target inserted | yes |
| No reheating temperature inserted | yes |
| No annihilation cross section fitted | yes |
| No decay width fitted | yes |
| No RG trajectory fitted | yes |
| No `λ_min` inserted | yes |
| No vacuum counterterm inserted | yes |
| No holographic saturation assumed | yes |
| No claim beyond available inputs | yes |

## Status ledger

```text
CONDITIONAL_SUPPORT_GATE374_SCOPED_CLOSURE_INHERITED
CONDITIONAL_SUPPORT_COSMOLOGICAL_OBSERVABLE_SIEVE_OPENED
CONDITIONAL_SUPPORT_ASHA_HEAVY_SCALE_LEDGER_FORMALIZED
CONDITIONAL_SUPPORT_DARK_MATTER_BOLTZMANN_SYSTEM_FORMALIZED
CONDITIONAL_SUPPORT_RELIC_DENSITY_INPUT_AUDIT_EXECUTED
CONDITIONAL_SUPPORT_VACUUM_LIFETIME_BOUNCE_FORMALIZED
CONDITIONAL_SUPPORT_METASTABILITY_INPUT_AUDIT_EXECUTED
CONDITIONAL_SUPPORT_COSMOLOGICAL_CONSTANT_CAPACITY_AUDITED
CONDITIONAL_SUPPORT_COSMOLOGICAL_OBSERVABLE_FIREWALL_PRESERVED
CONDITIONAL_SUPPORT_NO_HARD_COSMOLOGICAL_OBSERVABLE_DERIVED_IN_CURRENT_LEDGER

CONDITIONAL_TENSION_HEAVY_SCALE_IS_NOT_A_RELIC_MODEL
CONDITIONAL_TENSION_BGAP_THRESHOLD_IS_CONDITIONAL_TRANSPORT_LEDGER
CONDITIONAL_TENSION_BOLTZMANN_RELIC_NEEDS_RATES_AND_INITIAL_STATE
CONDITIONAL_TENSION_MAJORANA_DARK_STABILITY_NOT_DERIVED
CONDITIONAL_TENSION_VACUUM_LIFETIME_NEEDS_FULL_RG_TRAJECTORY
CONDITIONAL_TENSION_BOUNCE_ACTION_NEEDS_LAMBDA_MINIMUM_AND_PREFACTOR
CONDITIONAL_TENSION_DARK_ENERGY_NEEDS_RENORMALIZED_COUNTERTERM
CONDITIONAL_TENSION_PFAFFIAN_HIERARCHY_IS_NOT_A_NATIVE_LAMBDA_COSMO_PREDICTION

FAILED_ROUTE_COSMOLOGICAL_OBSERVABLES_NOT_DERIVED
FAILED_ROUTE_DARK_MATTER_RELIC_DENSITY_NOT_DERIVED
FAILED_ROUTE_OMEGA_DM_H2_NOT_COMPUTED_FROM_ASHA_LEDGER
FAILED_ROUTE_STABLE_DARK_SECTOR_CANDIDATE_NOT_DERIVED
FAILED_ROUTE_BOLTZMANN_KERNEL_NOT_DERIVED
FAILED_ROUTE_UNIVERSE_LIFETIME_NOT_DERIVED
FAILED_ROUTE_EUCLIDEAN_BOUNCE_ACTION_NOT_COMPUTED
FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_DERIVED
FAILED_ROUTE_DARK_ENERGY_DENSITY_NOT_COMPUTED
```

## Final truth statement

Gate 375 proves that ASHA does **not yet** derive dark matter abundance, universe lifetime, or dark energy from the finite ledger alone.

The result is not a philosophical retreat.  It is a sharp map of the next required theory layer:

```text
derive dark-sector Lagrangian
derive stability/decay-width theorem
derive Boltzmann production history
derive continuum RG trajectory and threshold matching
derive λ_min and bounce functional
derive renormalized vacuum-energy counterterm or holographic saturation theorem
```

Only after those continuum cosmology ingredients are native can ASHA output hard cosmological predictions.
