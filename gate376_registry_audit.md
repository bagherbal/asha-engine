# Gate 376 Registry Audit — Almost-Commutative Product Geometry / Full SM+Gravity Spectral Action Assembly

## Gate identity

| Field | Value |
|---|---|
| Gate | 376 |
| Package | `pkg/bridge/almostcommutativeproduct` |
| Theorem | `AlmostCommutativeProductGeometryFullSMGravitySpectralActionAssemblyTheorem()` |
| Audit ID | `GATE376-ALMOST-COMMUTATIVE-PRODUCT-GEOMETRY-FULL-SM-GRAVITY-SPECTRAL-ACTION-ASSEMBLY` |
| Layer | Bridge / Continuum Product Assembly |

## Executive truth statement

Gate 376 corrects the direction of the continuum bridge.

The finite ASHA spectral triple `F` is **not** required to derive spacetime `M`. The correct noncommutative-geometric bridge is the almost-commutative product:

```text
Total geometry = M x F
D_total = D_M ⊗ 1_F + gamma5 ⊗ D_F
```

The spectral action on `M x F` expands into continuum geometric invariants on `M` multiplied by finite ASHA spectral invariants. This assembles the Standard Model plus Einstein-gravity Lagrangian skeleton while preserving the open firewalls: the cosmological constant, cosmological initial data, universe lifetime, dark-matter relic density, and the 13 charged flavor moduli are not derived by this assembly alone.

## Inherited ledger

| Source | Inherited result | Gate-376 use |
|---|---|---|
| Gates 297–298 | completed finite spectral-triple skeleton and inner fluctuation field content | finite factor `F` supplies `A_F`, `H_F`, `D_F`, `J_F`, `gamma_F`, gauge/Higgs content |
| Gate 317 | canonical doubled three-generation finite carrier has dimension `96` | used as `Tr_F(1)` in product-action multiplicity ledger |
| Gate 343 | gravitational spectral-action product `f2(Lambda/M_P)^2 = pi/64` | Einstein-Hilbert coefficient channel |
| Gate 344 | cosmological `f4 Lambda^4` channel remains firewalled | cosmological constant is present structurally but not predicted |
| Gates 362–373 | modular/Fock/holographic attempts do not select flavor texture | 13 charged moduli remain after product assembly |
| Gate 374 | scoped finite-kinematic closure with `13 + theta_QCD + scale` ledger | flavor vacuum remains outside pure finite geometry |
| Gate 375 | hard cosmological observables require continuum model inputs | Gate 376 supplies product Lagrangian interface, not observable numbers |

## Product triple formalization

The gate formalizes:

```text
A = C∞(M) ⊗ A_F,       A_F = C ⊕ H ⊕ M3(C)
H = L²(M,S) ⊗ H_F
D = D_M ⊗ 1_F + gamma5 ⊗ D_F
J = J_M ⊗ J_F
gamma = gamma_M ⊗ gamma_F
```

Status:

```text
CONDITIONAL_SUPPORT_ALMOST_COMMUTATIVE_PRODUCT_TRIPLE_FORMALIZED
CONDITIONAL_SUPPORT_ASHA_FINITE_FACTOR_ACCEPTED_AS_INTERNAL_GEOMETRY
CONDITIONAL_SUPPORT_CONTINUUM_SPIN_MANIFOLD_SUPPLIED_AS_PRODUCT_FACTOR
CONDITIONAL_SUPPORT_TOTAL_DIRAC_OPERATOR_ASSEMBLED
CONDITIONAL_TENSION_SPACETIME_M_NOT_DERIVED_FROM_FINITE_ALGEBRA
FAILED_ROUTE_DISCRETE_TO_CONTINUUM_DERIVATION_REJECTED
```

The last failed route is not negative for the theory. It rejects the wrong direction. The continuum bridge is product geometry, not derivation of `M` from `F`.

## Seeley-deWitt product expansion

The product spectral action is formalized as:

```text
Tr f(D_total²/Lambda²) ~ Σ_n f_{4-n} Lambda^{4-n} a_n(D_total²)
```

On `M x F`, the terms factor into continuum invariants on `M` and finite traces over `H_F`.

| Heat-kernel order | Continuum invariant | Finite invariant | Physical term | Gate-376 verdict |
|---|---|---|---|---|
| `Lambda^4 a0` | `int_M sqrt(g)` | `Tr_F(1)`, vacuum multiplicity/counterterm | bare cosmological/vacuum-energy channel | present, not predicted |
| `Lambda^2 a2` | `int_M sqrt(g) R` | `Tr_F(1)`, `f2 Lambda^2` | Einstein-Hilbert gravity | assembled |
| `a4` gauge | `int_M sqrt(g) F_{mu nu}F^{mu nu}` | representation trace / inner fluctuations | SM gauge kinetic terms | assembled |
| `a4` scalar kinetic | `int_M sqrt(g) |nabla H|^2` | finite one-form Higgs doublet | Higgs kinetic term | assembled |
| `a2/a4` scalar potential | `int_M sqrt(g)(lambda |H|^4 - mu^2|H|^2)` | `Tr(D_F²)`, `Tr(D_F⁴)` | Higgs potential | assembled with ASHA boundary ratio |
| fermionic action | `int_M sqrt(g) psibar(D_M + gamma5 D_F)psi` | finite Dirac edge graph and moduli | Yukawa masses and mixing | structurally assembled; texture remains free |
| `a4` gravity squared | `int_M sqrt(g)(C², R*R, ...)` | `Tr_F(1)` | higher-curvature gravity | structurally present |

Status:

```text
CONDITIONAL_SUPPORT_SEELEY_DEWITT_HEAT_KERNEL_PRODUCT_EXPANSION_FORMALIZED
CONDITIONAL_TENSION_HEAT_KERNEL_NORMALIZATION_CONVENTIONS_MUST_BE_TRACKED
```

## ASHA finite spectral invariants substituted

| Invariant | Value | Enters | Caveat |
|---|---:|---|---|
| `Tr_F(1)` | `96` | multiplicity channels | raw dimension is not the weighted trace capacity |
| `f0` | `7` | gauge/Higgs `a4` channels | heat-kernel convention must be tracked |
| `f2(Lambda/M_P)^2` | `pi/64` | Einstein-Hilbert coefficient | separates into `f2` and `Lambda` only after cutoff choice |
| `sin²(theta_W)` | `3/8` | gauge kinetic normalization | IR comparison requires RG transport |
| `alpha_branch^-1` | `8pi` | gauge branch ledger | absolute empirical unification remains threshold/convention sensitive |
| `lambda_H/g_*²` | `1197/4624` | Higgs quartic boundary | IR pole mass requires RG and matching |
| `dim M_charged(D_F)` | `13` | Yukawa sector | flat coordinates, not fixed constants |

Status:

```text
CONDITIONAL_SUPPORT_ASHA_FINITE_SPECTRAL_INVARIANTS_SUBSTITUTED
```

## Full Lagrangian skeleton

Gate 376 assembles the product-action Lagrangian skeleton:

```text
S[M x F] = int_M sqrt(g) {
    Einstein-Hilbert gravity
  + bare cosmological/vacuum term
  + SU(3)xSU(2)xU(1) gauge kinetic terms
  + Higgs kinetic term
  + Higgs potential
  + fermionic/Yukawa action
  + higher-curvature gravity terms
}
```

| Sector | Identified | Fully predicted by current ledger? | Remaining freedom |
|---|---:|---:|---|
| Einstein-Hilbert gravity | yes | yes, as product coefficient | cutoff convention if splitting `f2` from `Lambda` |
| Cosmological/vacuum term | yes | no | renormalized `f4 Lambda^4` counterterm/subtraction |
| Gauge kinetic | yes | structurally yes | IR RG and threshold matching |
| Higgs kinetic | yes | structurally yes | field normalization convention |
| Higgs potential | yes | boundary ratio yes | continuum RG/pole matching |
| Yukawa/fermion masses | yes | no | 13 charged flavor moduli |
| Higher-curvature gravity | yes | structurally yes | continuum gravity regime |

Status:

```text
CONDITIONAL_SUPPORT_FULL_SM_GRAVITY_LAGRANGIAN_SKELETON_ASSEMBLED
CONDITIONAL_SUPPORT_PRODUCT_GEOMETRY_BRIDGE_DERIVED
CONDITIONAL_SUPPORT_EINSTEIN_HILBERT_TERM_IDENTIFIED
CONDITIONAL_SUPPORT_SM_GAUGE_KINETIC_TERMS_IDENTIFIED
CONDITIONAL_SUPPORT_HIGGS_KINETIC_AND_POTENTIAL_TERMS_IDENTIFIED
CONDITIONAL_SUPPORT_YUKAWA_SECTOR_IDENTIFIED_WITH_13_MODULI
CONDITIONAL_TENSION_COSMOLOGICAL_CONSTANT_F4_COUNTERTERM_CHANNEL_OPEN
CONDITIONAL_TENSION_YUKAWA_TEXTURE_REMAINS_13_MODULI
```

## Continuum computation interface

Gate 376 explains why Gate 375 could not compute cosmological observables directly. The finite ledger alone was missing the product-action interface. After Gate 376, the following computations become well-posed **once continuum data are supplied**:

| Computation | Enabled by product Lagrangian? | Still required |
|---|---:|---|
| RG running | yes | renormalization scheme, threshold ledger, moduli values |
| Boltzmann evolution | yes | stable dark candidate, rates, reheating/initial conditions |
| Vacuum bounce / lifetime | yes | full `lambda(mu)`, bounce scale, prefactor, gravity corrections |
| Classical cosmology | yes | metric ansatz, topology, matter content, initial data |
| Dark energy / `Lambda_cosmo` | structurally present | renormalized vacuum counterterm or saturation theorem |

Status:

```text
CONDITIONAL_SUPPORT_CONTINUUM_COMPUTATION_INTERFACE_OPENED
CONDITIONAL_TENSION_CONTINUUM_OBSERVABLES_REQUIRE_SPACETIME_AND_INITIAL_DATA
FAILED_ROUTE_DARK_MATTER_RELIC_DENSITY_STILL_NOT_PREDICTED_BY_PRODUCT_ASSEMBLY
FAILED_ROUTE_VACUUM_LIFETIME_STILL_NOT_PREDICTED_BY_PRODUCT_ASSEMBLY
FAILED_ROUTE_COSMOLOGICAL_CONSTANT_STILL_NOT_PREDICTED_BY_PRODUCT_ASSEMBLY
```

## Firewall audit

| Firewall | Preserved? |
|---|---:|
| Does not derive spacetime `M` from finite `F` | yes |
| Does not predict observed cosmological constant | yes |
| Does not predict dark matter relic density | yes |
| Does not predict universe lifetime | yes |
| Does not select Yukawa/CKM texture | yes |
| Does not erase the 13 charged finite-Dirac moduli | yes |
| Does not hide heat-kernel conventions | yes |
| Does not claim full-suite cosmology | yes |

## Registry status names

Successful / conditional support:

```text
CONDITIONAL_SUPPORT_GATE375_COSMOLOGICAL_OBSERVABLE_FIREWALL_INHERITED
CONDITIONAL_SUPPORT_ALMOST_COMMUTATIVE_PRODUCT_TRIPLE_FORMALIZED
CONDITIONAL_SUPPORT_ASHA_FINITE_FACTOR_ACCEPTED_AS_INTERNAL_GEOMETRY
CONDITIONAL_SUPPORT_CONTINUUM_SPIN_MANIFOLD_SUPPLIED_AS_PRODUCT_FACTOR
CONDITIONAL_SUPPORT_TOTAL_DIRAC_OPERATOR_ASSEMBLED
CONDITIONAL_SUPPORT_SEELEY_DEWITT_HEAT_KERNEL_PRODUCT_EXPANSION_FORMALIZED
CONDITIONAL_SUPPORT_ASHA_FINITE_SPECTRAL_INVARIANTS_SUBSTITUTED
CONDITIONAL_SUPPORT_FULL_SM_GRAVITY_LAGRANGIAN_SKELETON_ASSEMBLED
CONDITIONAL_SUPPORT_EINSTEIN_HILBERT_TERM_IDENTIFIED
CONDITIONAL_SUPPORT_SM_GAUGE_KINETIC_TERMS_IDENTIFIED
CONDITIONAL_SUPPORT_HIGGS_KINETIC_AND_POTENTIAL_TERMS_IDENTIFIED
CONDITIONAL_SUPPORT_YUKAWA_SECTOR_IDENTIFIED_WITH_13_MODULI
CONDITIONAL_SUPPORT_CONTINUUM_COMPUTATION_INTERFACE_OPENED
CONDITIONAL_SUPPORT_PRODUCT_GEOMETRY_BRIDGE_DERIVED
```

Tensions / preserved firewalls:

```text
CONDITIONAL_TENSION_SPACETIME_M_NOT_DERIVED_FROM_FINITE_ALGEBRA
CONDITIONAL_TENSION_COSMOLOGICAL_CONSTANT_F4_COUNTERTERM_CHANNEL_OPEN
CONDITIONAL_TENSION_YUKAWA_TEXTURE_REMAINS_13_MODULI
CONDITIONAL_TENSION_CONTINUUM_OBSERVABLES_REQUIRE_SPACETIME_AND_INITIAL_DATA
CONDITIONAL_TENSION_ABSOLUTE_GAUGE_NORMALIZATION_REMAINS_CONVENTION_DEPENDENT
CONDITIONAL_TENSION_HEAT_KERNEL_NORMALIZATION_CONVENTIONS_MUST_BE_TRACKED
```

Failed routes deliberately preserved:

```text
FAILED_ROUTE_DISCRETE_TO_CONTINUUM_DERIVATION_REJECTED
FAILED_ROUTE_COSMOLOGICAL_CONSTANT_STILL_NOT_PREDICTED_BY_PRODUCT_ASSEMBLY
FAILED_ROUTE_DARK_MATTER_RELIC_DENSITY_STILL_NOT_PREDICTED_BY_PRODUCT_ASSEMBLY
FAILED_ROUTE_VACUUM_LIFETIME_STILL_NOT_PREDICTED_BY_PRODUCT_ASSEMBLY
FAILED_ROUTE_FLAVOR_VACUUM_STILL_NOT_SELECTED_BY_PRODUCT_ASSEMBLY
```

## Final truth statement

Gate 376 is the marriage theorem.

The finite ASHA geometry supplies the internal factor `F`; spacetime supplies the continuum factor `M`; the spectral action on `M x F` assembles the Standard Model plus Einstein gravity with ASHA finite coefficients. This is the correct bridge from the 375-gate finite ledger to continuum physics.

The gate does not claim that the finite algebra derives spacetime, the cosmological constant, dark-matter relic abundance, universe lifetime, or the 13 charged flavor parameters. It instead installs the exact product-action framework in which those continuum computations can now be performed lawfully.
