# Asha Engine v0.89 — O(3) Gauge Quotient

The codebase is organized around a theorem ladder:

```text
Cℓ(1,7) covariant phase space
→ exterior algebra / grades
→ Boolean incidence geometry
→ octonionic G₂ calibration
→ contact vacuum K₇
→ finite dynamics / spectral action / BF action
→ contact-preserving gauge algebra
→ matter / Fock space / triality
→ bridge layer for constants
→ theorem engine / falsification gates
→ cinematic reporting
```

## Non-negotiable rules

1. No physical constants are hard-coded in the finite core.
2. Mathematical dimensions are derived from definitions.
3. Every output has a theorem status.
4. Failed routes are recorded as scientific information.
5. Bridge-layer claims are never labeled exact until the bridge theorem exists.

## Status taxonomy

- `EXACT_FINITE`
- `VARIATIONAL`
- `VERIFIED_NUMERIC`
- `OPEN_TEST`
- `BRIDGE_REQUIRED`
- `PHENOMENOLOGY`
- `FAILED_ROUTE`
- `INVALID_COMPARISON`


## Gate 4 — Octonionic / G₂ Calibration

The G₂ gate constructs the octonionic matter-calibration sector inside `Λ⁴R⁸`. The engine uses the standard Fano convention

```text
φ = 123 + 145 + 167 + 246 − 257 − 347 − 356
```

and its Hodge-dual coassociative form `*φ` to build two seven-dimensional calibration copies, `7_t` and `7_s`. Their direct sum is the rank-14 sector `M₁₄ᴳ`.

This gate verifies `P_G² = P_G`, `P_Gᵀ = P_G`, and `Tr(P_G)=14`. It does not yet compare the G₂ support to the Boolean support. The contact theorem belongs to the next gate.

## Gate 27 — finite-core cache / fixture layer

The theorem ladder now treats the heavy finite-core constructions as reusable runtime fixtures. This does not change the mathematics; it makes the engine behave like an instrument rather than a disposable script.

Cached default constructors now protect the repeated construction path for:

- Boolean–Octonionic contact space `K`.
- B-sector vacuum operator `O_B`.
- Boolean lift/compression diagnostics.
- Projected connection analysis.
- Higgs/vacuum-mixing sector.
- Higgs potential candidate.
- Matter/Fock bridge layers through the triality Yukawa extension.

The boundary-fixed Lie closure diagnostic is intentionally runtime-capped. Its scientific purpose is to expose dimension growth under the current projection, not to regenerate a large closure basis at every full run.


## Gate 28 — Generation-Breaking Texture Search

Exact triality is now treated as a replication principle, not as a texture-selection principle. The gate classifies the `3×3` flavor texture spaces exposed by the Yukawa/triality audit and proves the immediate no-go: a symmetric texture invariant under exact triality has only a singlet plus doublet spectrum. Therefore the engine must discover an additional finite generation-breaking operator before making any mass, CKM, or PMNS claim.

## Gate 84 — Finite Scalar Covariant Derivative / Gauge-Boson Mass Matrix

Gate 84 constructs an abstract finite scalar covariant-derivative template on the four-real-dimensional scalar/contact frame. It verifies the W/Z/photon signature at the dimensionless bridge level: two degenerate charged modes, one neutral massive mode, and one electromagnetic null direction. It deliberately does not claim physical W/Z masses because scalar kinetic normalization, gauge couplings, gauge-field Hessian, and canonical vacuum orientation remain open.

## Gate 85 — Scalar kinetic / gauge-eating diagnostic

Gate 85 strengthens the scalar covariant-derivative bridge by auditing the scalar kinetic frame and Goldstone image structure. It verifies that the diagnostic scalar vacuum has one radial direction, three independent broken-generator images from `{T1,T2,Z=T3−YΦ}`, and an electromagnetic null generator `Q=T3+YΦ`. The result is still bridge-level: the Euclidean scalar metric is not yet an action-selected kinetic normalization, and physical W/Z masses remain forbidden until the gauge Hessian, couplings, and vacuum orientation are derived.

## Gate 87 — Protected-Contact / Broken-Generator Intertwiner

Gate 87 tests whether the three protected contact directions can be canonically
identified with the three broken electroweak generator image directions.

The gate verifies the 3-3-3 count-level resonance and the positive broken-image
metric. It deliberately does not declare a gauge-eating theorem: an abstract
O(3) family of isometries exists between 3D frames, but the finite action has
not yet selected one. The next mathematical target is a protected-contact metric
or connection form that can reduce this O(3) freedom canonically.

## Gate 88 — Protected-Contact Metric / Connection Form Search

Gate 88 audits the protected side of the gauge-eating resonance.  Gate 87
showed the count-level match `3 protected contact directions ↔ 3 broken
generator images`, but count equality left an abstract `O(3)` family of
possible intertwiners.  Gate 88 therefore distinguishes an abstract Euclidean
metric on a three-dimensional carrier from a metric or connection actually
derived from finite BF/contact dynamics.

Current result: the protected side admits an abstract `I_3` metric, and the
broken-generator image metric is positive with condition number 4.  Pulling that
metric back to the protected side would be circular until the protected-to-broken
intertwiner is independently derived.  The `O(3)` freedom may be pure gauge, but
that quotient theorem remains open.

## Gate 91 — Gauge-Quotiented Protected-to-Broken Correspondence Audit

Gate 91 follows Gate 90's conclusion that the protected-contact `O(3)` frame freedom behaves as gauge for the currently implemented intrinsic protected diagnostics. It therefore forbids component-wise matching between protected and broken frames.

The gate keeps only quotient-safe data:

- protected carrier dimension = 3;
- broken-generator image rank = 3;
- electromagnetic null direction from the scalar covariant-derivative diagnostic;
- broken-image metric spectrum and condition number.

The result is a strengthened but still bridge-level correspondence. The count/rank resonance survives the quotient, but the broken image metric is anisotropic, while the protected carrier currently has only the abstract Euclidean reference metric. Therefore no quotient-safe protected-to-broken intertwiner is derived yet.

Next target: derive whether the broken-image anisotropy is physical kinetic data, gauge normalization data, or an artifact of diagnostic normalization.


## Gate 92 — Broken-Image Metric / Kinetic Normalization Audit

Gate 92 audits the anisotropic metric on the broken-generator image after the protected `O(3)` quotient. The raw diagnostic metric has charged eigenvalue `0.2833333333` and neutral eigenvalue `1.1333333333`, giving condition number `4`. The gate shows this anisotropy is exactly removable by normalizing the neutral broken generator by `1/2`, so the current anisotropy is not yet a physical W/Z mass prediction. It is unresolved gauge/scalar kinetic normalization data.

The gate therefore preserves the structural Goldstone/gauge-eating signature while rejecting premature physical-mass claims. The next missing object is an action-selected gauge kinetic normalization or normalized broken-generator basis.

## Gate 93 — Normalized Broken-Generator Basis / Gauge-Kinetic Candidate

Gate 93 refines the broken-image metric result.  The raw broken-generator image
metric has neutral-to-charged ratio `4`; the neutral broken generator can be
scaled by `1/2`, making the quotient-safe broken metric isotropic.  In the raw
field-coordinate basis this exposes the diagnostic kinetic candidate
`diag(1,1,4)`.  The engine records this as a canonical diagnostic extracted
from the broken-image metric, but not as a physical gauge kinetic Hessian until
a finite action second variation selects it.

## Gate 96 — Finite Broken Gauge Field Variables / Curvature Term Search

Gate 96 upgrades the broken-sector diagnostic into typed gauge-field variables.  The broken variables are `W1`, `W2`, and `Z_raw=T3-Y_phi`; the unbroken electromagnetic variable is `A_em` along `Q=T3+Y_phi`.

The closure audit shows that the broken variables alone are not a curvature carrier: `[T1,T2]=T3=(Z+Q)/2`, so the electromagnetic direction cannot be discarded when constructing a finite field strength.  The result redirects the next gate toward the full electroweak connection `{T1,T2,Z,Q}` and forbids treating `diag(1,1,4)` as action-selected before a full finite curvature term and second variation are derived.


## Gate 97 — Full Electroweak Connection Curvature / Field-Strength Audit

Gate 97 adds `pkg/bridge/ewcurvature`. Gate 96 showed that the broken-only variables `{T1,T2,Z=T3-Y_phi}` are not closed because `[T1,T2]=T3=(Z+Q)/2`. Gate 97 therefore keeps the full connection `{T1,T2,Z,Q}` and audits its Lie closure and adjoint/Killing diagnostic.

The full connection closes and supports a formal field-strength carrier. However, the adjoint diagnostic is rank three and has the pure abelian direction `Q-Z=2Y_phi` as a null vector. It sees the semisimple neutral direction `T3=(Z+Q)/2`, not a positive physical `U(1)` kinetic Hessian. Consequently `diag(1,1,4)` remains a strong broken-image metric-whitening candidate, not an action-selected physical Hessian.

Truth: full electroweak curvature must include the electromagnetic direction, but the curvature algebra alone still does not derive `g2`, `gY`, `thetaW`, `alpha`, or physical W/Z masses.

## Gate 98 — Full Electroweak Quadratic Action / Abelian Completion

Gate 98 completes the Gate 97 full electroweak curvature carrier by adding the missing abelian quadratic term as a one-parameter family:

```text
K(kappa)=K_SU2+kappa(Q-Z)(Q-Z)^T
```

In the basis `[T1,T2,Z,Q]`, the semisimple curvature sees the `SU(2)` direction but leaves the pure abelian `Q-Z=2Y_phi` direction null.  The abelian completion exposes the correct missing coefficient problem.  In the chosen convention, the earlier raw broken-coordinate candidate `diag(1,1,4)` is reachable at `kappa=6`, but this value is not action-selected.  The gate therefore keeps `g2`, `gY`, `thetaW`, `alpha`, and physical W/Z masses bridge-gated.
