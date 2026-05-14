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
