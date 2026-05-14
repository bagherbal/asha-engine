# ASHA Engine

A typed Go research engine for the finite Boolean--Octonionic source program.

The engine is intentionally theorem-gated:

1. exact algebra first;
2. variational finite dynamics second;
3. gauge and topology tests third;
4. bridge-layer physics last.

No observed physical constants are hard-coded in the finite core.

## Current gates

- `ALG-EXT-GRADES-R8` — exterior grade structure of `R^8`.
- `ALG-CLIFFORD-CL-1-7` — Clifford algebra `Cℓ(1,7)` bookkeeping.
- `PHY-PHASE-COVARIANT-4D` — covariant phase-space split `x_mu + p_mu`.
- `GEO-BOOL-INCIDENCE-R8-L3-U4` — Boolean incidence `Λ^3R^8 -> Λ^4R^8`.
- `GEO-G2-CALIBRATION-L4-R8` — octonionic `G2` rank-14 calibration sector.
- `GEO-CONTACT-K-BG-L4-R8` — contact space `K = Im(P_B) ∩ Im(P_G)`.
- `DYN-BSECTOR-CONTACT-VACUUM` — finite B-sector action whose zero modes are `K`.
- `GAUGE-G2-R-CENTRALIZER` — tangent-level centralizer `g2^R ≅ su(2) ⊕ u(1)`.
- `GAUGE-BOOLEAN-LIFT-COMPRESSION` — harsh Boolean compression diagnostic.
- `GAUGE-BOUNDARY-FIXED-CLOSURE` — minimal boundary-fixed Lie-closure growth diagnostic.

## Current truth from the gauge branch

The contact-preserving tangent algebra exists:

```text
g2^R ≅ su(2) ⊕ u(1)
```

But the naive Boolean lift/compression does **not** produce a closed finite
`su(2) ⊕ u(1)` algebra. After imposing the contact boundary, the generated
finite Lie algebra continues growing beyond the four seed generators. This is
not hidden or patched: the engine reports it as an open theorem failure and a
new structural constraint.

## Run

```bash
go test ./...
go run ./cmd/asha
```

## v0.10 — Projected connection truth: finite Higgs/vacuum-mixing sector

Gate 11 adds `pkg/gauge/higgs`. It keeps the off-diagonal blocks discarded by strict Boolean compression:

```text
Φ_i = P_C A_i P_K + P_K A_i P_C
```

The gate verifies that `Φ_i` is purely off-diagonal and skew, builds positive mixing operators on the contact vacuum and its complement, and reports the derived finite spectra. This gate does not claim the physical Higgs mass. It identifies the finite object that the bridge layer must use for electroweak symmetry breaking, Yukawa texture, and vacuum mixing.

Current finite result:

- Higgs/vacuum-mixing span rank: 2
- Contact-vacuum mixing rank: 4 inside `dim(K)=7`
- Unmixed contact directions: 3
- Trace balance: `Tr(M_K)=Tr(M_C)=1.1333333333`


## v0.12 — Fock / Matter Bridge

Added the first typed matter bridge:

- `pkg/spinor`: four-mode Witt/Fock bookkeeping with `2^4 = 16` states.
- `pkg/matter`: bridge analysis between the Fock generation seed and the contact/Higgs finite data.
- New theorem gates:
  - `SPINOR-WITT-FOCK-16`
  - `MATTER-FOCK-CONTACT-BRIDGE`

The bridge verifies a conservative kinematic resonance:

- four Fock creation modes;
- sixteen occupation states;
- one B−L-neutral sterile vacuum seed;
- one temporal/lepton seed plus three spatial/color seeds;
- four active Higgs/contact directions;
- three protected unmixed contact directions.

It intentionally does **not** claim Standard Model hypercharge, Yukawa masses, a Higgs mass, or dark matter stability. Those remain future theorem gates.


## v0.14 — Canonical Fock/contact embedding search

Gate 15 tests whether the finite Higgs/contact active eigenstructure is already sufficient to assign the four active directions to the Fock `1+3` split. The result is intentionally honest: the active sector is four-dimensional and the Fock sector has a `1 temporal + 3 spatial` bookkeeping split, but the current Higgs/contact spectrum is pair-degenerate `2+2`. Therefore the eigenvectors are not canonically aligned with lepton/color modes yet.

The new missing object is a finite charge-polarizing operator compatible with the block connection on `K ⊕ K⊥`, such as a `B−L`, hypercharge, or central `U(1)` action.

## v0.15 — Charge polarization bridge

Gate 16 adds the first explicit charge-polarizing bridge. Gate 15 showed that the finite Higgs/contact active spectrum is pair-degenerate (`2+2`) and cannot by itself determine the required Fock `1+3` split. Gate 16 resolves the missing object on the Fock side with the standard B−L number operator:

```text
Q_B-L = (1/3)(N1+N2+N3) - N0
```

This supplies a one-particle charge spectrum `[-1, 1/3, 1/3, 1/3]`, exactly the `1+3` polarization. The gate also rejects a false shortcut: the active Higgs/contact scalar eigenvalues must not be identified directly with the three color modes. Charge polarization and scalar/vacuum mixing are separate structures until a true tensor-factor or representation-action bridge is constructed.

## v0.16 — Tensor-factor matter/Higgs bridge

Gate 17 adds `pkg/matter/tensor`. Gate 16 supplied the `B−L` charge polarization on the Fock/matter side, but it also showed that the pair-degenerate finite Higgs/contact scalar spectrum must not be collapsed into color charges. Gate 17 resolves this by introducing the standard tensor product bridge:

```text
H_total = H_Fock ⊗ H_Φ
Q_total = Q_B-L ⊗ I_Φ
S_total = I_Fock ⊗ S_Φ
```

The gate verifies that matter charge and scalar/contact response commute because they live on different tensor factors. It also verifies the trace identities, the neutral vacuum scalar fiber `|Ω⟩⊗H_Φ`, and the charge-sector scalar fibers. The result is a cleaner architecture for the next missing object: a gauge-compatible Yukawa/intertwiner map between matter states and the scalar factor. No Yukawa texture or physical mass matrix is claimed yet.


## v0.17 — Yukawa/intertwiner selection rule

Gate 18 adds `pkg/matter/yukawa`. The tensor-factor bridge made the correct domain explicit:

```text
H_total = H_Fock ⊗ H_Φ
Q_total = Q_B-L ⊗ I_Φ
```

The new gate formulates the first honest Yukawa/intertwiner selection rule:

```text
[Q_total, Y] = 0
```

With the scalar factor currently neutral under `B−L`, the allowed neutral intertwiner space is block-diagonal in charge sectors. The engine computes:

- `dim End(H_total) = 4096`;
- `dim{Y : [Q,Y]=0} = 672`;
- `3424` charge-changing entries are forbidden unless the scalar/contact factor receives its own charge/hypercharge bridge;
- the one-particle neutral selection space has dimension `160`;
- parity-preserving and parity-flipping dimensions are balanced, but true chirality remains an open theorem.

This gate does not derive fermion masses. It turns `U-07-YUKAWA` into a precise mathematical problem: construct a non-arbitrary hypercharge/chirality-compatible intertwiner rather than a fitted mass matrix.


## Gate 20 — Scalar Hypercharge / T3 Bridge Search

The engine now extracts the finite part genuinely supported by the active Higgs/contact sector: a 2+2 real scalar doublet with a canonical trace-zero scalar charge `T_Φ = diag(+1/2,+1/2,-1/2,-1/2)`. It also proves the harder truth: this scalar-side charge is not yet full Standard Model hypercharge. The matter-side `T3_R` and physical chirality operator remain open bridge theorems.

## v0.20 — Gate 21: Matter-side T3_R / chirality search

Adds `pkg/matter/t3r`, which searches for the matter-side operator missing after the scalar hypercharge bridge. The gate tests the canonical temporal occupation polarization `T0 = 1/2 - N0`, verifies that it is compatible with `B-L`, rejects the vectorlike version as physical chirality, and shows that chiral restrictions of `T0` unlock gauge-compatible grading-flipping channels. The result remains bridge-layer: the even/odd physical orientation is still open, and no Yukawa texture or mass spectrum is claimed.

## v0.22 — SU(2)_L doublet charge audit

Gate 23 adds `pkg/matter/su2l`, a charge-level audit of the missing left-doublet bridge. It uses the Gate 22 odd right-singlet/conjugate hypercharge table and the Gate 20 finite scalar doublet charge to test the Yukawa charge-balance equation:

```text
Y_L = Y_R - Y_Φ
```

Result:

- standard orientation: `u_R(2/3)-(+1/2)=1/6`, `d_R(-1/3)-(-1/2)=1/6`, giving `Q_L: Y=1/6 × 6`;
- standard lepton orientation: `ν_R(0)-(+1/2)=-1/2`, `e_R(-1)-(-1/2)=-1/2`, giving `L_L: Y=-1/2 × 2`;
- conjugate mirror orientation is also present with `Q_L^c: Y=-1/6 × 6` and `L_L^c: Y=+1/2 × 2`.

The gate deliberately does not claim the full nonabelian `SU(2)_L` theorem. It proves the doublet hypercharges at charge-selection level and leaves the actual raising/lowering generators, conjugation convention, and explicit Yukawa intertwiners open.

## v0.23 — SU(2)_L finite generator audit

Gate 24 adds `pkg/matter/su2lgauge`. Gate 23 proved the left-doublet hypercharges at the charge-selection level; Gate 24 now builds the explicit finite `SU(2)_L` ladder representation on the derived left-doublet space:

```text
H_L = Q_L(3 colors × 2 weak states) ⊕ L_L(2 weak states)
T3 = diag(+1/2,-1/2) on each doublet
T+ |down⟩ = |up⟩
T- |up⟩ = |down⟩
```

The gate verifies the ladder algebra:

```text
[T3,T+] = T+
[T3,T-] = -T-
[T+,T-] = 2T3
```

and also verifies:

```text
[Y,T±] = 0
Q = T3 + Y
```

This upgrades `U-13A-SU2L-GENERATORS` from missing to solved at the audited left-doublet representation level. The remaining deeper bridge is to derive the same generators directly from the finite Boolean/contact block connection rather than from the hypercharge table.

## v0.24 — Gate 25: Gauge-Compatible Yukawa/Intertwiner Audit

Gate 25 adds the first explicit one-generation Yukawa channel audit. It uses the finite left-doublet representation from Gate 24 and the scalar branch charges from Gate 20, then enforces the selection rule:

```text
Y_R = Y_L + Y_Φ
```

The engine derives eight minimal channels without fitting masses:

```text
u_L^c ⊗ Φ_+ → u_R^c  for c=1,2,3
d_L^c ⊗ Φ_- → d_R^c  for c=1,2,3
ν_L   ⊗ Φ_+ → ν_R
e_L   ⊗ Φ_- → e_R
```

Each scalar branch has the finite pair multiplicity inherited from the 2+2 Higgs/contact spectrum, so the eight minimal channels correspond to sixteen scalar-fiber entries. This gate is intentionally a selection-rule theorem only: coupling constants, masses, flavor mixing, and generation structure remain open bridge problems.

### v0.27 — Gate 28: Generation-Breaking Texture Search

Added `pkg/matter/texture`, a no-go/search gate for Yukawa texture selection. The gate proves that exact triality can replicate the one-generation Yukawa channel pattern into three generations, but cannot by itself select numerical `3×3` Yukawa matrices or three distinct generation eigenvalues. A triality-invariant symmetric texture has a `1+2` eigenvalue pattern, so a real hierarchy requires a new finite generation-breaking operator compatible with the contact/BF/Higgs geometry.

## v0.29 — Gate 30: Curvature on Generation Carrier Search

Gate 30 tests the contact-side mirror curvature

```text
R^K_AB = P_K A P_C B P_K - P_K B P_C A P_K
```

on the three protected contact directions exposed by the Higgs/contact sector.
The result is an important no-go: the protected 3D generation carrier is
curvature-flat under this second-fundamental source, while the active 4D
Higgs/contact carrier receives nonzero curvature.

This means the diagonal generation-breaking spurion from Gate 29 is still the
best bridge-level generation-splitting object, but curvature-induced generation
mixing is not yet selected. The next mathematical target is a finite BF
curvature operator or a new projection from active curvature into generation
texture space.


## v0.30 — Active curvature to generation projection bridge

Gate 31 adds `pkg/matter/bfbridge`. Gate 30 found that second-fundamental curvature is real on the active Higgs/contact carrier but flat on the protected 3D generation carrier. Gate 31 tests the tempting bridge honestly: can the existing finite block connection project active curvature into the protected generation carrier?

It computes the natural cross maps

```text
B_i = G^T A_i H
```

where `G` is the protected 3D carrier, `H` is the active 4D carrier, and `A_i` are the Boolean-compressed connection generators. Then it tests induced generation operators

```text
B_i F_active B_j^T
B_i F_active^T F_active B_j^T
```

The result is a no-go for the current implementation: active curvature exists, but the existing connection gives no nonzero active-to-protected bridge. Therefore the engine refuses to claim CKM/PMNS or generation mixing from active Higgs curvature alone. The next missing theorem is a genuine finite BF/Maurer-Cartan curvature operator or another canonical active-generation projection principle.


## v0.31 — Gate 32: Finite BF / Maurer-Cartan curvature

Added `pkg/matter/bfcurvature`.

This gate implements the first genuine finite Maurer-Cartan residual of the Boolean-compressed block connection:

```text
F_ij = [A_i, A_j] - Π_seed([A_i, A_j])
```

It then restricts that residual to:

- the protected 3D generation carrier `G`,
- the active 4D Higgs/contact carrier `H`,
- the cross bridge `Gᵀ F H`.

The result is a useful no-go: the finite curvature is real on the full Boolean support and active Higgs/contact sector, but it does not yet project into protected generation mixing. Therefore CKM/PMNS and non-diagonal Yukawa textures remain open.

## v0.33 — Gate 34: Source tensor selection / active-generation map search

Gate 34 adds `pkg/matter/sourcemap`. It searches for a canonical source tensor

```text
M : H_active → H_generation
```

that could turn active Higgs/contact curvature into a genuine `3×3` generation texture. The gate tests the existing compressed-connection cross maps `GᵀAH`, the finite Maurer-Cartan curvature cross maps `GᵀFH`, and the BF action mixed source contraction. All natural canonical maps remain zero at the current stage. The abstract map space `Hom(H_active,H_generation)` has dimension `3×4=12`, but choosing a tensor there would be fitting rather than derivation.

The gate therefore preserves the hard truth: diagonal generation splitting exists as a bridge-level spurion, but no canonical non-diagonal source tensor, CKM structure, PMNS structure, or physical Yukawa-strength bridge has been derived yet.

## v0.34 — Source Tensor Action / Variational Selection

Gate 35 adds the variational source-tensor action over `M : H_active → H_generation`.

The minimal stable action is:

```text
S[M] = 1/2 ||M||_F^2 - <J, M>
```

The stationary equation is `M = J`. Since the finite connection, BF curvature,
and BF source contractions currently select `J ≈ 0`, the unique stable stationary
source tensor is `M = 0`. This is a useful no-go theorem: the abstract 12D tensor
space `Hom(R^4,R^3)` exists, but choosing a nonzero tensor would still be fitting
unless a new finite interaction derives a nonzero source, constraint, or
symmetry-breaking action.

New package:

```text
pkg/matter/sourceaction
```

The engine now exposes the remaining unknown as:

```text
U-17E-SOURCE-TENSOR-ACTION
```


## v0.36 — Scalar-sector effective potential normal form

Gate 37 adds `pkg/dynamics/scalarpotential`, which derives the finite scalar/Higgs normal-form invariants from the active contact sector:

- active scalar dimension: `4` real directions = `2` complex doublet components;
- protected contact resonance: `3` unmixed directions;
- finite radius: `r0² = Tr(M_K)`;
- quartic shape: `Tr(M_K²)/Tr(M_K)²`;
- bridge-level shifted normal form: `V(r)=λ_shape(r²-r0²)²`.

This gate intentionally does **not** claim the electroweak vev, observed Higgs mass, or a proven gauge-eating theorem. Those remain bridge-layer unknowns.

## v0.43 — Threshold spectrum / matching audit

Gate 44 adds `pkg/bridge/threshold`, which audits finite spectral threshold anchors from the B-sector gap, contact partial-overlap modes, scalar/contact active spectrum, leakage invariant, and scalar radial curvature. It deliberately keeps them dimensionless and refuses physical threshold masses until a non-fitted mass unit, activation rule, and finite-to-continuum matching map are derived.


## v0.45 — Threshold Activation / Decoupling Audit

Gate 46 adds `pkg/bridge/thresholdactivation`.

It classifies finite threshold anchors by activation status:

- scalar/contact active sector: continuum-field candidate at sector level, but not a heavy threshold;
- scalar radial response: scalar-sector bridge object, not a separate threshold;
- contact leakage: vacuum-frustration-only;
- B-sector gap: threshold-open, no representation/activation/scale;
- contact partial-overlap modes: threshold-open, no physical-field/regulator/frustration decision.

The gate explicitly rejects threshold-corrected beta coefficients until a physical unit, activation rule, representation assignment, and decoupling prescription are derived.

## v0.58 — Gate 59: Current Action on Scalar LR Projector

Adds `pkg/bridge/currentprojection`, which lifts the finite `u(4)`-shaped current inventory to the left×scalar domain and right-singlet image of the scalar LR projector. The gate computes finite representation-overlap diagnostics for central, color, B−L, and leptoquark current sectors while keeping Fierz signs, current kinetic normalization, propagator rules, and up/down splitting explicitly open.


## v0.62 — Gate 63: Finite Exchange Action / Propagator Normalization Search

Adds `pkg/bridge/exchangeaction`, which audits whether the signed Fierz result and conditional exchange kernel can be promoted to a derived finite NJL kernel. The gate computes unit, inverse-kinetic, and kinetic-weighted exchange diagnostics, but keeps the exchange action sign, propagator denominators, relative current couplings, up/down splitting, and NJL criticality open.
