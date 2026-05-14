# Gate 401 Registry Audit — Derived Edge-Weight Operator / Hypercharge Laplacian Sieve

## Claim tested

Gate 401 tests whether native electroweak, B-L, or T3-like charges can differentiate the ten J-doubled finite-Dirac one-form edges strongly enough to produce the irreducible contact quartic `q4` as a canonical invariant on the four-real scalar carrier `H_phi`.

## Inheritance

```text
executed=true gate400_uniform_central=true gate400_pair_deg=true gate400_no_q4=true gate385_edges=true edge_count=10 gate26_yukawa=true gate41_hypercharge=true charged_moduli=13 no_empirical=true verdict=Gate 401 inherits the Gate-400 q4 obstruction, Gate-385 ten-edge one-form support, the charge-compatible Yukawa edge classes, and the 13-moduli flavor firewall.
```

## q4 target

```text
q4=3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271 degree=4 irreducible_Q=true monic=[1, -2.36666666667, 1.98333333333, -0.689814814815, 0.0836419753086] verdict=The target remains the branch-free irreducible contact primary q4. A differentiated edge Laplacian must match this polynomial directly or by a derived finite theorem, not by an arbitrary affine fit.
```

## Edge-weight arena

```text
formalized=true structural_edges=5 J_doubled=10 Hphi_dim=4 native_Y=true native_B-L=true native_T3=true yukawa_amplitudes_used=false observed_masses_used=false edges=[L_L ↔ e_R/Y_R=-1/B-L=-1/T3=-0.5/branch=Φ_-; L_L ↔ ν_R/Y_R=0/B-L=-1/T3=0.5/branch=Φ_+; Q_L ↔ d_R/Y_R=-0.333333/B-L=0.333333/T3=-0.5/branch=Φ_-; Q_L ↔ u_R/Y_R=0.666667/B-L=0.333333/T3=0.5/branch=Φ_+; ν_R ↔ ν_R^c/Y_R=0/B-L=1/T3=0/branch=singlet/Majorana] verdict=The finite one-form edge graph has five structural edge classes and ten J-doubled slots. Hypercharge, scalar branch charge, T3-like branch sign, and B-L provide native charge weights; no Yukawa amplitudes or observed masses are used.
```

## Weighted candidate table

```text
executed=true native_anisotropic=5 native_quartic_capacity=2 canonical_q4_matches=0 sealed_matches=1 best_native=edge-resolved squared-hypercharge stress test best_residual=1.72482562421 verdict=Native electroweak and B-L weights do break edge uniformity, but canonical scalar-branch compression remains central or pair-degenerate. Edge-resolved stress tests can reach degree four only by a noncanonical edge-to-H_phi component placement, and their polynomials are disjoint from q4.
name=uniform J-doubled edge measure source=Gate385 support measure formula=Δ_E=1 on every one-form edge native=true sealed=false circular=false Hphi_endomorphism=true canonical_Hphi=true edge_resolved=false branch_compressed=true J_real=true gauge=true uses_yukawa=false uses_mass=false eigen=[1, 1, 1, 1] distinct=1 min_degree=1 char=1*x^4 + -4*x^3 + 6*x^2 + -4*x + 1 residual_q4=5.53158448173 pair_deg=false central=true quartic_capacity=false q4_exact=false promotable=false verdict=FAILED_ROUTE reason=Uniform support is the already-audited Gate400 central case.
name=scalar branch T3/hypercharge weight source=Gate21/Gate41 scalar weak charge formula=diag(+1/2,+1/2,-1/2,-1/2) native=true sealed=false circular=false Hphi_endomorphism=true canonical_Hphi=true edge_resolved=false branch_compressed=true J_real=false gauge=true uses_yukawa=false uses_mass=false eigen=[0.5, 0.5, -0.5, -0.5] distinct=2 min_degree=2 char=1*x^4 + -0.5*x^2 + 0.0625 residual_q4=3.49919233787 pair_deg=true central=false quartic_capacity=false q4_exact=false promotable=false verdict=CONDITIONAL_SUPPORT_ANISOTROPIC_WEIGHT_BUT_NOT_Q4 reason=The scalar weak charge is native but pair-degenerate by the Higgs doublet structure.
name=branch-averaged right-hypercharge edge Laplacian source=right hypercharge averaged by scalar branch formula=Φ_+:avg(Y_u,Y_ν)=1/3; Φ_-:avg(Y_d,Y_e)=-2/3 native=true sealed=false circular=false Hphi_endomorphism=true canonical_Hphi=true edge_resolved=false branch_compressed=true J_real=false gauge=true uses_yukawa=false uses_mass=false eigen=[0.333333333333, 0.333333333333, -0.666666666667, -0.666666666667] distinct=2 min_degree=2 char=1*x^4 + 0.666666666668*x^3 + -0.333333333333*x^2 + -0.148148148148*x + 0.0493827160493 residual_q4=3.85520842889 pair_deg=true central=false quartic_capacity=false q4_exact=false promotable=false verdict=CONDITIONAL_SUPPORT_ANISOTROPIC_WEIGHT_BUT_NOT_Q4 reason=Canonical branch compression forgets which real component came from which Yukawa class, so it remains 2+2.
name=branch-averaged B-L edge Laplacian source=B-L averaged by scalar branch formula=Φ_+:avg(1/3,-1)=-1/3; Φ_-:avg(1/3,-1)=-1/3 native=true sealed=false circular=false Hphi_endomorphism=true canonical_Hphi=true edge_resolved=false branch_compressed=true J_real=false gauge=true uses_yukawa=false uses_mass=false eigen=[-0.333333333333, -0.333333333333, -0.333333333333, -0.333333333333] distinct=1 min_degree=1 char=1*x^4 + 1.33333333333*x^3 + 0.666666666665*x^2 + 0.148148148148*x + 0.0123456790123 residual_q4=4.0163262069 pair_deg=false central=true quartic_capacity=false q4_exact=false promotable=false verdict=FAILED_ROUTE reason=B-L is native, but the scalar-branch average collapses to a central value for the Dirac Yukawa edges.
name=edge-resolved right-hypercharge four-channel stress test source=right hypercharge per Yukawa structural class formula=diag(Y_u,Y_d,Y_ν,Y_e)=diag(2/3,-1/3,0,-1) native=true sealed=false circular=true Hphi_endomorphism=true canonical_Hphi=false edge_resolved=true branch_compressed=false J_real=false gauge=true uses_yukawa=false uses_mass=false eigen=[0.666666666667, -0.333333333333, 0, -1] distinct=4 min_degree=4 char=1*x^4 + 0.666666666666*x^3 + -0.555555555556*x^2 + -0.222222222222*x residual_q4=3.9840565651 pair_deg=false central=false quartic_capacity=true q4_exact=false promotable=false verdict=CONDITIONAL_SUPPORT_QUARTIC_CAPACITY_BUT_NOT_Q4 reason=This gives four distinct values, but placing four edge classes onto the four real H_phi components is a noncanonical component assignment; its polynomial is also disjoint from q4.
name=edge-resolved squared-hypercharge stress test source=quadratic hypercharge norm per Yukawa structural class formula=diag(Y_u²,Y_d²,Y_ν²,Y_e²)=diag(4/9,1/9,0,1) native=true sealed=false circular=true Hphi_endomorphism=true canonical_Hphi=false edge_resolved=true branch_compressed=false J_real=false gauge=true uses_yukawa=false uses_mass=false eigen=[0.444444444444, 0.111111111111, 0, 1] distinct=4 min_degree=4 char=1*x^4 + -1.55555555556*x^3 + 0.604938271604*x^2 + -0.0493827160493*x residual_q4=1.72482562421 pair_deg=false central=false quartic_capacity=true q4_exact=false promotable=false verdict=CONDITIONAL_SUPPORT_QUARTIC_CAPACITY_BUT_NOT_Q4 reason=Charge norms are native diagnostics, but the edge-to-real-scalar-component assignment is not canonical and the polynomial is not q4.
name=edge-resolved B-L four-channel stress test source=B-L per Yukawa structural class formula=diag(1/3,1/3,-1,-1) native=true sealed=false circular=true Hphi_endomorphism=true canonical_Hphi=false edge_resolved=true branch_compressed=false J_real=true gauge=true uses_yukawa=false uses_mass=false eigen=[0.333333333333, 0.333333333333, -1, -1] distinct=2 min_degree=2 char=1*x^4 + 1.33333333333*x^3 + -0.222222222221*x^2 + -0.444444444444*x + 0.111111111111 residual_q4=4.31456098354 pair_deg=true central=false quartic_capacity=false q4_exact=false promotable=false verdict=CONDITIONAL_SUPPORT_ANISOTROPIC_WEIGHT_BUT_NOT_Q4 reason=B-L distinguishes quark/lepton sectors, not four scalar components; it remains 2+2 and cannot be q4.
name=sealed q4-weighted edge companion source=manual q4 placement formula=companion(q4) declared as Δ_E(w) on H_phi native=false sealed=true circular=true Hphi_endomorphism=true canonical_Hphi=false edge_resolved=true branch_compressed=false J_real=false gauge=false uses_yukawa=false uses_mass=false eigen=[0.283912192592, 0.441122757284, 0.744096637981, 0.897535078809] distinct=4 min_degree=4 char=1*x^4 + -2.36666666667*x^3 + 1.98333333333*x^2 + -0.689814814814*x + 0.0836419753085 residual_q4=1.44357650717e-12 pair_deg=false central=false quartic_capacity=true q4_exact=true promotable=false verdict=SEALED_STRESS_TEST_ONLY reason=The q4 roots can be imposed as four weights, but no edge charge theorem derives them, so this is quarantined.
```

## Identity / impact audit

```text
Hphi_q4_identified=false scalar_bundle_sealed=false diff_weights=true canonical_weighted_laplacian=false yukawa_reduced=false moduli_start=13 moduli_result=13 flavor_firewall=true higgs_lane_preserved=true verdict=Differentiated charge weights exist, but no canonical q4-weighted scalar Laplacian is derived. Higgs coefficient and flavor moduli ledgers remain unchanged.
```

## Firewall status

```text
executed=true no_masses=true no_ckm=true no_pmns=true no_yukawa_amplitudes=true no_manual_q4=true no_arbitrary_edge_component_map=true no_affine_fit=true no_moduli_reduction=true verdict=No observed masses, CKM/PMNS data, fitted Yukawa amplitudes, arbitrary affine charge fit, or manual q4/H_phi identification is promoted.
```

## Statuses

```text
CONDITIONAL_SUPPORT_GATE400_MIXED_EDGE_OBSTRUCTION_INHERITED
CONDITIONAL_SUPPORT_1_FORM_EDGE_SUPPORT_INHERITED
CONDITIONAL_SUPPORT_ELECTROWEAK_CHARGES_INHERITED
CONDITIONAL_SUPPORT_NATIVE_EDGE_WEIGHTS_AUDITED
CONDITIONAL_SUPPORT_DIFFERENTIATED_EDGE_LAPLACIAN_FORMALIZED
CONDITIONAL_SUPPORT_ANISOTROPIC_EDGE_WEIGHTS_FOUND
CONDITIONAL_SUPPORT_EDGE_RESOLVED_QUARTIC_CAPACITY_FOUND
FAILED_ROUTE_UNIFORM_EDGE_WEIGHT_REMAINS_CENTRAL
FAILED_ROUTE_BRANCH_COMPRESSION_REMAINS_PAIR_DEGENERATE
FAILED_ROUTE_HYPERCHARGE_EDGE_POLYNOMIAL_DISJOINT_FROM_Q4
FAILED_ROUTE_B_MINUS_L_EDGE_POLYNOMIAL_DISJOINT_FROM_Q4
FAILED_ROUTE_T3_EDGE_POLYNOMIAL_NOT_Q4
FAILED_ROUTE_NO_CANONICAL_EDGE_TO_HPHI_COMPONENT_MAP
FAILED_ROUTE_NO_NATIVE_Q4_WEIGHTED_EDGE_LAPLACIAN
FAILED_ROUTE_NO_CANONICAL_HPHI_QUARTIC_IDENTIFICATION
FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION
FIREWALL_PRESERVED_13_MODULI
```

## Conclusion

Gate 401 proves that native electroweak/B-L charges can differentiate the ten J-doubled one-form edges, but they do not solve the q4/H_phi identity problem. Canonical scalar-branch compression is still central or 2+2 pair-degenerate; edge-resolved hypercharge can give degree-four capacity only after a noncanonical assignment of edge classes to real H_phi components, and its characteristic polynomial is disjoint from q4 (best native residual 1.72483 from edge-resolved squared-hypercharge stress test). The 13 charged flavor moduli remain sealed.

## Next gate

```text
Gate 402 — Spectral Graph Edge-Adjacency Operator Search
Reason: Charge weights differentiate edges but do not produce q4. The missing object is not a gauge charge weight; it is likely a native adjacency/incidence operator on the full one-form edge graph whose quotient to H_phi has four nondegenerate eigenvalues.
Primary task: Build the finite one-form edge graph adjacency/incidence Laplacian, including edge-edge incidence through shared source/target bimodule nodes, and test its canonical scalar quotient.
```
