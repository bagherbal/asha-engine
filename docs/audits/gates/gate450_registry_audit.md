# Gate 450 Registry Audit — Structural Zero Mass-Mixing Invariants / Ratio Sieve

## Scope

Gate 450 cancels the publication-only lane and tests the strongest post-444 flavor temptation: whether the forced Generation-2 structural zero plus the closed-triangle bridge topology already implies a GST/Fritzsch-style mass-angle ratio. It uses symbolic matrix algebra and counterexample sieves only; no observed lepton, quark, CKM, PMNS, or Yukawa data is imported.

## Inheritance

G444K=true G444Zero=true G445Triangle=true G446PhaseSealed=true G447CoeffSealed=true G449Board=true nativeDim=13 KXY=9 noEmpirical=true verdict=CONDITIONAL_SUPPORT_GATE449_STRUCTURAL_FAMILY_BOARD_INHERITED

## Symbolic matrix arena

K=diag(-1,0,1); X=S+S^T=[[0,1,1],[1,0,1],[1,1,0]]; Y=i(S-S^T)=[[0,i,-i],[-i,0,i],[i,-i,0]]; M(a,b,c)=aK+bX+cY = [[-a,b+ic,b-ic],[b-ic,0,b+ic],[b+ic,b-ic,a]] hermitian=true trace0=true M22Zero=true triangle=true endpointBalanced=true symbolicOnly=true noCollider=true verdict=CONDITIONAL_SUPPORT_SYMBOLIC_TEXTURE_ZERO_MATRIX_CONSTRUCTED reason=The most general Gate-445/Gate-446 K/X/Y texture-zero test matrix is Hermitian, trace-neutral, closed-triangular, and has M22=0 without importing masses or mixing data.

```text
M(a,b,c)=aK+bX+cY = [[-a,b+ic,b-ic],[b-ic,0,b+ic],[b+ic,b-ic,a]]
```

## Symbolic eigen-analysis

chi=chi(lambda)=lambda^3-(a^2+3(b^2+c^2))lambda-2(b^3-3bc^2) P=P=a^2+3r^2, r^2=b^2+c^2 D=D=2(b^3-3bc^2)=2r^3 cos(3phi), phi=atan2(c,b) cardano=lambda_k=2 sqrt(P/3) cos((1/3) arccos((3 sqrt(3) D)/(2 P^(3/2)))-2 pi k/3), k=0,1,2 eigenvector=for z=b+ic and lambda root, v(lambda) proportional to (z^2+lambda conjugate(z), conjugate(z)^2+(a+lambda)z, lambda(a+lambda)-|z|^2) trace=sum_i lambda_i=0 det=prod_i lambda_i=D=2(b^3-3bc^2) convention=physical positive masses would be singular/eigenvalue magnitudes after sector conventions; Gate 450 uses only signed symbolic eigenlevels freeScaleRatio=true freePhase=true coeffRequired=true verdicts=[CONDITIONAL_SUPPORT_CHARACTERISTIC_POLYNOMIAL_DERIVED,CONDITIONAL_SUPPORT_EIGENVECTOR_FORMULA_DERIVED] reason=The symbolic spectrum is computable, but it depends on two independent dimensionless coordinates after removing overall scale: a/r and phi.

| Object | Formula | Status |
|---|---|---|
| Characteristic polynomial | `chi(lambda)=lambda^3-(a^2+3(b^2+c^2))lambda-2(b^3-3bc^2)` | `CONDITIONAL_SUPPORT_CHARACTERISTIC_POLYNOMIAL_DERIVED` |
| Quadratic invariant | `P=a^2+3r^2, r^2=b^2+c^2` | symbolic |
| Cubic/determinant invariant | `D=2(b^3-3bc^2)=2r^3 cos(3phi), phi=atan2(c,b)` | symbolic |
| Eigenvalues | `lambda_k=2 sqrt(P/3) cos((1/3) arccos((3 sqrt(3) D)/(2 P^(3/2)))-2 pi k/3), k=0,1,2` | Cardano form |
| Eigenvectors | `for z=b+ic and lambda root, v(lambda) proportional to (z^2+lambda conjugate(z), conjugate(z)^2+(a+lambda)z, lambda(a+lambda)-\|z\|^2)` | `CONDITIONAL_SUPPORT_EIGENVECTOR_FORMULA_DERIVED` |

## Texture-zero identity

zero=M_22=0 sumRule=0=M_22=sum_i lambda_i |U_{2i}|^2 localAngles=tan(2 theta_12)=tan(2 theta_23)=2r/a for the endpoint-balanced two-level subblocks, r=sqrt(b^2+c^2) GST=sin(theta_ij) ?= sqrt(|m_i/m_j|) exact=true specificRatio=false needsU=true needsCoeff=true verdict=CONDITIONAL_SUPPORT_TEXTURE_ZERO_SUM_RULE_DERIVED reason=The structural zero gives an exact spectral sum rule, but the sum rule contains the eigenvector weights themselves and does not collapse to a pairwise GST ratio without extra texture assumptions.

The structural zero does prove the exact spectral sum rule

```text
0=M_22=sum_i lambda_i |U_{2i}|^2
```

but this is not the same thing as a pairwise relation such as `sin(theta_ij)=sqrt(m_i/m_j)`. The eigenvector weights remain part of the equation.

## Ratio sieve counterexamples

examples=4 sameAngleDifferentMass=true sameMassDifferentAngle=true zeroSumRule=true uniqueInvariant=false coeffRequired=true phaseRequired=true scaleIrrelevant=true verdict=FAILED_ROUTE_RATIOS_REQUIRE_EXACT_AMPLITUDES reason=The topology fixes the zero and support, but the dimensionless coefficients a/r and phi remain free. The counterexamples prove that neither mixing angles nor normalized mass ratios determine the other.

| Witness | a | b | c | phi | q=D/P^(3/2) | theta_local | normalized eigenvalues | abs mass ratios | Demonstrates |
|---|---:|---:|---:|---:|---:|---:|---|---|---|
| same-angle/A: phi=0 | 2 | 1 | 0 | 0 | 0.10799 | 0.392699 | `[-0.940863, -0.109295, 1.05016]` | `[0.116165, 0.895925]` | same theta_12/theta_23 as same-angle/B but different mass-shape q |
| same-angle/B: phi=pi/5 | 2 | 0.809017 | 0.587785 | 0.628319 | -0.0333707 | 0.392699 | `[-1.01629, 0.033408, 0.982877]` | `[0.03399, 0.967127]` | same theta_12/theta_23 as same-angle/A but different mass-shape q |
| same-mass/A: alpha=1, phi=0 | 1 | 1 | 0 | 0 | 0.25 | 0.553574 | `[-0.837565, -0.269594, 1.10716]` | `[0.321879, 0.756499]` | same normalized spectrum as same-mass/B but different local mixing angle |
| same-mass/B: alpha=0, tuned phi | 0 | 0.958829 | 0.283985 | 0.287948 | 0.25 | 0.785398 | `[-0.837565, -0.269594, 1.10716]` | `[0.321879, 0.756499]` | same normalized spectrum as same-mass/A but different local mixing angle |

Two independent obstructions are visible: one pair keeps the same local mixing angle while changing the mass-shape invariant, and another pair keeps the same normalized mass spectrum while changing the local mixing angle. This kills a universal topology-only mass-angle ratio.

## GST / Fritzsch test

historical=GST/Fritzsch relations arise in stricter nearest-neighbor or additional-zero textures, often after hierarchy and phase assumptions. extras=select a coefficient ray such as c=0 or fixed phi; suppress or constrain the 1-3 edge, which Gate 445 did not allow for mass lift; choose sector-dependent amplitude hierarchy; define a physical mass-ordering and rephasing convention fullTriangle=true phaseFree=true coeffRayFree=true forcedGST=false universalApprox=false specialLater=true verdict=FAILED_ROUTE_GST_FRITZSCH_RELATION_NOT_FORCED reason=The ASHA structural topology is texture-zero-like, but it is a full closed triangle with a quarantined cycle phase and coefficient ray. GST can be studied as a special branch, not promoted as a native invariant at Gate 450.

Extra assumptions that would be needed before a GST-like branch could be audited:

- select a coefficient ray such as c=0 or fixed phi
- suppress or constrain the 1-3 edge, which Gate 445 did not allow for mass lift
- choose sector-dependent amplitude hierarchy
- define a physical mass-ordering and rephasing convention

## Result statuses

- `CONDITIONAL_SUPPORT_GATE449_STRUCTURAL_FAMILY_BOARD_INHERITED`
- `CONDITIONAL_SUPPORT_SYMBOLIC_TEXTURE_ZERO_MATRIX_CONSTRUCTED`
- `CONDITIONAL_SUPPORT_CHARACTERISTIC_POLYNOMIAL_DERIVED`
- `CONDITIONAL_SUPPORT_EIGENVECTOR_FORMULA_DERIVED`
- `CONDITIONAL_SUPPORT_TEXTURE_ZERO_SUM_RULE_DERIVED`
- `CONDITIONAL_SUPPORT_INVARIANT_RATIO_SIEVE_EXECUTED`
- `CONDITIONAL_SUPPORT_GST_FRITZSCH_TEST_EXECUTED`
- `FAILED_ROUTE_RATIOS_REQUIRE_EXACT_AMPLITUDES`
- `FAILED_ROUTE_GST_FRITZSCH_RELATION_NOT_FORCED`
- `FAILED_ROUTE_MASS_EIGENVALUES_DO_NOT_DETERMINE_MIXING`
- `FAILED_ROUTE_MIXING_ANGLES_DO_NOT_DETERMINE_MASS_RATIOS`
- `FAILED_ROUTE_PHASE_CONTINUUM_PRESERVED`
- `FAILED_ROUTE_NO_MUON_CHARM_PHYSICAL_MASS_PREDICTION`
- `FAILED_ROUTE_NO_CKM_PMNS_ANGLE_OR_PHASE_PREDICTION`
- `CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED`
- `FIREWALL_REFINED_TEXTURE_ZERO_IDENTITIES_BUT_RATIO_VALUES_SEALED`

## Firewall

noMuon=true noCharm=true noYukawa=true noCKM=true noPMNS=true noFit=true K=true X=true YSealed=true coeffSealed=true ratioSealed=true nativeDim=13 KXY=9 verdict=CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED reason=Gate 450 derives symbolic texture-zero identities but refuses to convert them into observed mass or mixing predictions without a native coefficient/phase selector.

## Next gate

Gate 451 — Texture-Zero Special-Branch Selector / Necessary Boundary Audit: Gate 450 proves the full ASHA triangle does not force GST by itself; the next honest step is to classify which extra native boundary would be required to recover a Fritzsch/GST branch. Task=Audit whether any existing ASHA law suppresses a specific edge, fixes phi, or selects a coefficient ray without importing empirical masses or CKM/PMNS data.

## Truth statement

Gate 450 computes the symbolic eigen-data of the forced structural-zero triangle and derives a real texture-zero identity, 0=sum_i lambda_i |U_2i|^2. It does not derive a GST/Fritzsch mass-angle prediction. The characteristic polynomial depends on P=a^2+3(b^2+c^2) and D=2(b^3-3bc^2), while the mixing data still depends on the independent coefficient ratio a/r and the cycle phase phi. Explicit counterexamples show same mixing with different mass ratios and same normalized mass spectrum with different mixing. Therefore the correct log is FAILED_ROUTE_RATIOS_REQUIRE_EXACT_AMPLITUDES, with the flavor firewall preserved.
