# Gate 511 Registry Audit — Gravitational a4 Curvature-Squared and Topological Counterterm Audit

## Verdict

```text
CONDITIONAL_SUPPORT_GATE510_A2_CURVATURE_COEFFICIENT_INHERITED
CONDITIONAL_SUPPORT_PRODUCT_SPECTRAL_A4_LEDGER_INHERITED
CONDITIONAL_SUPPORT_A4_CURVATURE_SQUARED_SOCKET_DEFINED
CONDITIONAL_SUPPORT_FOUR_DIMENSIONAL_CURVATURE_BASIS_CLASSIFIED
CONDITIONAL_SUPPORT_GAUSS_BONNET_TOPOLOGICAL_COUNTERTERM_IDENTIFIED
CONDITIONAL_SUPPORT_WEYL_SQUARED_DYNAMICAL_SOCKET_PRESENT
CONDITIONAL_SUPPORT_A4_DIMENSIONLESS_F0_CHANNEL_ISOLATED
CONDITIONAL_SUPPORT_A4_CHANNEL_DOES_NOT_USE_F2_LAMBDA_SQUARED
FAILED_ROUTE_NEWTON_CONSTANT_STILL_NOT_DERIVED_BY_A4
FAILED_ROUTE_CUTOFF_SCALE_LAMBDA_STILL_NOT_SELECTED_BY_A4
FAILED_ROUTE_EINSTEIN_HILBERT_NORMALIZATION_STILL_OPEN_AFTER_A4
FAILED_ROUTE_COSMOLOGICAL_F4_VACUUM_CHANNEL_STILL_UNSOLVED
FAILED_ROUTE_A4_CURVATURE_SQUARED_COEFFICIENTS_NOT_UNIQUE_PHYSICAL_GRAVITY_DYNAMICS
FAILED_ROUTE_A4_BOUNDARY_TERMS_AND_RENORMALIZATION_SCHEME_NOT_CLOSED
FAILED_ROUTE_HIGHER_DERIVATIVE_METRIC_EQUATIONS_NOT_NATIVE_DERIVED
FIREWALL_PRESERVED_NO_NEWTON_PLANCK_COSMOLOGY_EW_OR_FLAVOR_DATA_IMPORTED
FIREWALL_BLOCKED_A4_CURVATURE_SQUARED_PHYSICAL_DYNAMICS_WRITE
```

## Inherited boundary

Gate510 supplies the a2 curvature coefficient provenance and its normalization firewall; Gate511 inherits only the product spectral-action a4 curvature² channel and the dimensionless f0 moment.

```text
Gate510 inherited=true; a2 weight native=true; Newton blocked=true; f4 excluded=true; product valid=true; a4 declared=true; f0 available=true; all coeffs closed=false; hard ToE=false
```

## a4 curvature-squared basis audit

the four-dimensional curvature² vector space is classified into a topological Euler counterterm, a Weyl² dynamical socket, and scheme-dependent scalar/boundary curvature² pieces; classification is native, metric dynamics selection is not.

```text
raw basis=Riem²,Ric²,R²; E4=E₄ = Riem² - 4 Ric² + R²; C2=C² = Riem² - 2 Ric² + R²/3; rank=3; topological=true; dynamical=true; unique dynamics=false
∫√g E₄ is Gauss-Bonnet/Euler topological data in four dimensions, up to boundary conventions
C² is the conformal/dynamical curvature-squared socket of the four-dimensional spectral action
```

## a4 coefficient channel

unlike the a2 Einstein-Hilbert channel, the a4 curvature² channel is dimensionless and controlled by f0 times universal heat-kernel curvature polynomials; this does not derive Newton's constant or a unique low-energy gravity action.

```text
S_a4,grav = f0·(4π)^(-2)·∫√g Tr_F[universal curvature² polynomial]/360; TrF=96; f0=7; prefactor/f0=0.00168868639404; f0-weighted=0.0118208047583; dimensionless=true; uses f2Λ²=false; uses f4Λ4=false; physical=false
```

## Topological counterterm audit

the Euler/Gauss-Bonnet density is the topological curvature² counterterm socket; ASHA may classify its presence, but no manifold topology, boundary condition, or physical coefficient is selected here.

```text
E₄ = Riem² - 4 Ric² + R²; integral topological=true; variation boundary-only=true; numeric Euler characteristic selected=false; native socket=true; physical coefficient=false
```

## Dynamical curvature firewall

Weyl² and related curvature² terms are legitimate spectral-action sockets, but their physical role depends on renormalization, boundary data, metric-sign conventions, and the already-unclosed Einstein-Hilbert normalization.

```text
Weyl²=true; all raw sockets=true; higher derivative=true; scheme selected=false; boundary selected=false; low-energy Einstein limit=false; metric equations native=false; physical dynamics closed=false
```

## Firewall result

Gate511 classifies dimensionless a4 curvature² sockets only. It imports no G, M_P, Λ value, cosmological constant, electroweak scale, Yukawa, CKM, or PMNS data, and writes no physical a4 dynamics or gravity normalization.

```text
G imported=false; G derived=false; Planck imported=false; Λ selected=false; EH closed=false; cosmological imported=false; cosmological derived=false; f4 subtraction=false; EW imported=false; flavor imported=false; a4 dynamics write=false; native gravity write=false
```

## Registry update

### Native entries

- The product spectral-action a4 gravitational channel contains a four-dimensional curvature² socket.
- The curvature² basis decomposes into Euler/Gauss-Bonnet topological data, Weyl² conformal/dynamical data, and scalar/boundary curvature² data.
- The a4 channel is dimensionless and does not require the f₂Λ² Einstein-Hilbert normalization product.

### Bridge entries

- The symbolic curvature² action has the form f0·(4π)^(-2)·Tr_F(universal curvature² polynomial)/360.
- The Weyl²/dynamical curvature² socket is present, but physical metric equations and low-energy gravity interpretation remain bridge-level.

### Environmental entries

- Renormalization scheme, boundary conditions, manifold topology, Newton normalization, cutoff selection, and cosmological vacuum subtraction remain quarantined.

### Failed routes

- FAILED_ROUTE_NEWTON_CONSTANT_STILL_NOT_DERIVED_BY_A4
- FAILED_ROUTE_CUTOFF_SCALE_LAMBDA_STILL_NOT_SELECTED_BY_A4
- FAILED_ROUTE_EINSTEIN_HILBERT_NORMALIZATION_STILL_OPEN_AFTER_A4
- FAILED_ROUTE_COSMOLOGICAL_F4_VACUUM_CHANNEL_STILL_UNSOLVED
- FAILED_ROUTE_A4_CURVATURE_SQUARED_COEFFICIENTS_NOT_UNIQUE_PHYSICAL_GRAVITY_DYNAMICS
- FAILED_ROUTE_A4_BOUNDARY_TERMS_AND_RENORMALIZATION_SCHEME_NOT_CLOSED
- FAILED_ROUTE_HIGHER_DERIVATIVE_METRIC_EQUATIONS_NOT_NATIVE_DERIVED

### Open theorems

- Audit the f4Λ4 cosmological/vacuum-energy channel and vacuum-subtraction obligation separately.
- Prove or reject a native renormalization/boundary condition selector for curvature² dynamics.
- Prove or reject a native manifold-topology selector for the Euler characteristic contribution.

## Next step

Gate512 should be:

```text
Gate 512 — Cosmological f4 Vacuum Energy and Subtraction Airlock Audit
```

Primary task:

```text
separate the native a0 volume prefactor from the physical cosmological constant, test whether any finite trace cancels vacuum energy, and formally quarantine Λ_cosmo/f4/subtraction data if no theorem appears
```

## Truth statement

Gate 511 proves that the product spectral action has a native, scale-independent a4 curvature-squared socket: the four-dimensional curvature basis decomposes into Gauss-Bonnet topological data and Weyl/scalar curvature-squared channels. This is a genuine spectral-geometry ledger entry, not a mass, flavor, or electroweak claim. But the gate does not derive Newton's constant, the cutoff, cosmological vacuum subtraction, a physical renormalization scheme, boundary conditions, or complete metric dynamics. The a4 channel is present; physical higher-derivative gravity remains quarantined.
