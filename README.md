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


## v0.63 — Gate 64: Finite Propagator Spectrum Search

Adds `pkg/bridge/propagatorspectrum`. This gate searches whether existing finite spectra can provide current-sector propagator denominators for the native NJL/four-fermion kernel. It exposes diagnostic denominator families from the B-sector gap, scalar/contact active mean, contact partial-overlap mean, scalar radial curvature, and contact leakage. It deliberately rejects all of them as derived propagators because no theorem maps those spectral anchors to the central/color/B−L/leptoquark current sectors yet.

Result: finite spectral anchors exist, but current-sector propagator weights remain unassigned; no condensation, Higgs VEV, or mass scale is claimed.

## v0.65 — Gate 66: Current-Sector Operator Construction Search

Gate 66 adds `pkg/bridge/sectoroperators`. Gate 65 showed that raw finite spectral lists could not be assigned to current sectors by multiplicity alone. Gate 66 fixes the direction: it constructs the actual finite current-sector Casimir operators

```text
C_A = Σ_a T_aᵀ T_a
```

on the one-generation `1 lepton + 3 color` flavor space for the sectors central, color-su3, B−L, and leptoquark. This produces real representation-level sector data:

- central: uniform flavor carrier;
- color-su3: zero on lepton and equal Casimir on the three color seeds;
- B−L: lepton/color-polarizing charge-square carrier;
- leptoquark: off-diagonal lepton-color carrier.

This resolves the operator-construction part of the current-sector problem, but it still does **not** derive propagator denominators, exchange signs, up/down splitting, NJL criticality, top condensation, the Higgs VEV, or fermion masses. The next task is to decide whether these sector Casimirs, or another finite kinetic/action operator, control current exchange.


## v0.67 — Gate 68: Finite Exchange-Action Selection Principle

Gate 68 exposes the action-selection problem for current exchange.  The engine now audits four finite candidate propagator rules:

- direct Casimir kernel `K_A = C_A`
- inverse nonzero Casimir kernel `K_A = C_A^+`
- trace-normalized Casimir kernel `K_A = C_A / Tr(C_A)`
- unit-sector rule `K_A = I`

All are positive finite diagnostics, but none is promoted to a physical propagator.  Direct and inverse diagnostics disagree on the dominant sector, so the engine rejects choosing either by convenience.  The next missing object is a finite current-field Hessian / action second variation that selects the exchange kernel.

## v0.69 — Gate 70: Current Field Embedding into Finite BF/Contact Action

Gate 70 adds `pkg/bridge/currentembedding`. Gate 69 exposed the missing current Hessian `K_current`; this gate turns that absence into a typed finite-action architecture. It defines current-sector fields for central, color-su3, B−L, and leptoquark sectors, embeds them as formal action variables, and exposes the minimal template

```text
S[B,A,j] = S_B[B] + S_block[A;K⊕K⊥] + 1/2 jᵀK_current j - <j,J_source[B,A]>
```

The gate is intentionally not a success claim: the map from finite `u(4)` Fock currents into Boolean/contact block-connection operators is still open, so `J_source`, the second variation `δ²S/δjδj`, the propagator rule, NJL attraction, top condensation, and fermion masses remain unclaimed.

## v0.70 — Gate 71: Current-to-contact embedding map search

Gate 71 adds `pkg/bridge/currentcontact`. Gate 70 typed the current-sector action slots but left the key map open:

```text
E_current_to_block : u(4) Fock currents -> Boolean/contact block operators.
```

The new gate searches for this embedding and exposes the representation obstruction. The available Boolean/contact block target is the four-generator contact-preserving `su(2)+u(1)`-shaped seed, while the Fock current inventory is the sixteen-generator `u(4)=central+color-su3+B-L+leptoquark` inventory. An arbitrary linear map `R^16 -> R^4` exists, but it would have a large kernel and is not selected by finite data.

Current result:

- abelian slot exists but cannot separate central `u(1)` from `B-L`;
- no `SU(3)c` adjoint 8D carrier is present in the contact block target;
- no 6D leptoquark carrier is derived in the contact block target;
- therefore `E_current_to_block`, the current Hessian, the exchange kernel, and NJL attraction remain open.

The next direction is a dual-carrier architecture: do not force all `u(4)` currents into the contact `su(2)+u(1)` block. Instead, test whether color/Pati-Salam currents and electroweak/contact currents live on separate coupled finite carriers.

## v0.71 — Gate 72: Dual-Carrier Gauge Architecture Split

Gate 72 adds `pkg/bridge/dualcarrier`. Gate 71 showed that forcing the full sixteen-dimensional `u(4)` Fock/Pati-Salam current inventory into the four-dimensional Boolean/contact `su(2)+u(1)` block is structurally wrong. Gate 72 therefore splits the theory into two typed finite carriers:

```text
Carrier A: Pati-Salam / Fock current carrier
  dim = 16 = central 1 + color-su3 8 + B-L 1 + leptoquark 6

Carrier B: Boolean/contact electroweak block carrier
  dim = 4 = contact-su2 3 + contact-u1 1
```

This preserves the color and leptoquark sectors on their native `u(4)` carrier instead of crushing them into the contact block. The missing object is now a coupling action/tensor between the two carriers, not a direct embedding:

```text
S_total = S_PS[j] + S_contact[A,Φ] + S_coupling[j,A,Φ]
```

The formal current-contact coupling tensor has dimension `16×4 = 64`, but no finite theorem selects its entries yet. Therefore the current Hessian, propagator rule, NJL attraction, top condensation, and fermion masses remain open.


## v0.76 — Gate 77: Non-Factorized Abelian Action / Kinetic-Mixing Search

Adds `pkg/bridge/u1nonfactor`, which tests the first non-factorized abelian source candidate: the Yukawa-incidence correlation between matter-side `B-L` and scalar/contact `T_phi`. The gate finds nonzero local correlation on the gauge-compatible Yukawa support, but the signed moment cancels exactly between up/down quark branches and neutrino/electron lepton branches. Therefore no net `B-L`/contact-`u1` kinetic source or physical `U(1)_Y` coupling is derived yet.

### v0.86 — Gate 87: Protected-Contact / Broken-Generator Intertwiner

Adds a bridge audit for the would-be gauge-eating map. The engine now verifies
that protected contact directions, scalar angular directions, and broken gauge
images all have dimension three, while refusing to identify them by hand. The
remaining obstruction is the absence of a canonical protected-contact metric or
connection selecting a unique O(3) intertwiner.


## v0.90 — Gauge-Quotiented Correspondence

Adds Gate 91, which compares protected contact directions to broken gauge-generator images only after quotienting arbitrary protected `O(3)` frame choices. It rejects component-wise frame matching and keeps only quotient-safe invariants: dimension/rank correspondence and broken-image metric spectrum. The gauge-eating bridge remains open because no quotient-safe intertwiner is derived.


## v0.92 — Normalized Broken-Generator Basis

Adds Gate 93, which normalizes the neutral broken generator by 1/2, isotropizes the quotient-safe broken-image metric, and exposes diag(1,1,4) as a gauge-kinetic candidate rather than an action-selected physical Hessian.


## v0.94 — Gate 95 Broken-Sector Action Second Variation

Adds `pkg/bridge/brokenaction`, which audits whether the raw-coordinate kinetic candidate `diag(1,1,4)` is selected by a finite scalar/gauge action second variation. The gate preserves the Gate 94 result that `diag(1,1,4)` exactly whitens the broken-generator image metric, but it refuses to promote that diagnostic to a physical gauge kinetic Hessian until finite gauge-field variables, scalar kinetic action, curvature/field-strength term, and `δ²S` are derived.

## v0.95 — Finite broken gauge-field variables / curvature search

Gate 96 adds `pkg/bridge/brokengaugefields`. Gate 95 identified `diag(1,1,4)` as the raw-coordinate Hessian candidate that whitens the broken-generator image metric, but it could not be selected by an action because broken gauge-field variables and a field-strength term were absent.

Gate 96 types the broken gauge variables `{W1, W2, Z_raw}` and audits their closure. The crucial result is that the broken sector alone is not closed: `[T1,T2]=T3=(Z+Q)/2`, so the electromagnetic direction `Q=T3+Y_phi` is required in any finite curvature construction. The next action must therefore be a full electroweak connection curvature term, not a broken-only curvature term.

Current result: typed finite gauge variables exist, but the finite field-strength action and second variation remain open. No physical coupling, weak angle, alpha, or W/Z mass is claimed.


## Gate 97 — Full Electroweak Connection Curvature / Field-Strength Audit

Gate 97 adds `pkg/bridge/ewcurvature`. Gate 96 showed that the broken-only variables `{T1,T2,Z=T3-Y_phi}` are not closed because `[T1,T2]=T3=(Z+Q)/2`. Gate 97 therefore keeps the full connection `{T1,T2,Z,Q}` and audits its Lie closure and adjoint/Killing diagnostic.

The full connection closes and supports a formal field-strength carrier. However, the adjoint diagnostic is rank three and has the pure abelian direction `Q-Z=2Y_phi` as a null vector. It sees the semisimple neutral direction `T3=(Z+Q)/2`, not a positive physical `U(1)` kinetic Hessian. Consequently `diag(1,1,4)` remains a strong broken-image metric-whitening candidate, not an action-selected physical Hessian.

Truth: full electroweak curvature must include the electromagnetic direction, but the curvature algebra alone still does not derive `g2`, `gY`, `thetaW`, `alpha`, or physical W/Z masses.

## v1.58 — Gate 160: Quartic External Selector Firewall

Adds `pkg/bridge/quarticexternalselector`.

Gate 159 proved that the quartic contact block cannot be ghost-graded internally: the four quartic roots form one transitive Galois orbit, so any nontrivial parity split is a branch choice. Gate 160 tests the logically prior escape hatch before moving to spectral-action methods: whether any already-derived external physical source can break the quartic orbit canonically.

Audited selector sources:

- scalar vacuum orientation;
- broken gauge generator images `{T1,T2,Z}`;
- matter-side `B−L` charge pullback;
- canonical action second variation;
- rational/quartic spectral cross-coupling.

Current result:

- scalar and broken-gauge sources do not yet have canonical maps into the contact quartic block;
- matter-side `B−L` still lacks a canonical Fock-to-contact pullback;
- the action second variation can restrict to the quartic primary block only as a Galois-safe isotropic scalar action, with spectrum `[1,1,1,1]`;
- `P_rational Ω P_quartic = 0` exactly by spectral orthogonality, so rational roots do not select quartic branches;
- no nondegenerate spectrum, no 2+2 split, no ghost grading, no BRST cancellation, no contact beta row, and no physical constant is derived.

The mode-by-mode quartic route is now blocked both internally and by all currently available external finite sources. The next gate should treat the quartic data collectively through Galois-invariant spectral functionals.

## v1.59 — Gate 161: Collective Quartic Spectral Functional

Adds `pkg/bridge/quarticspectralfunctional`.

Gate 160 closed the last mode-by-mode escape hatch for the quartic contact block. Gate 161 therefore changes strategy: it treats the four quartic roots only collectively, through exact Galois-invariant symmetric spectral data.

Exact quartic factor:

```text
3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271
```

The gate computes the branch-free Newton/zeta ledger:

- quartic trace `p1 = 71/30`;
- quartic quadratic power sum `p2 = 1471/900`;
- quartic cubic power sum `p3 = 33581/27000`;
- quartic quartic power sum `p4 = 809891/810000`;
- quartic inverse trace `ζ_q(1) = 2235/271`;
- full seven-root contact trace `58/15`;
- full seven-root inverse trace `7993/542`.

Current result:

- the quartic block now has a clean collective action-level spectral ledger;
- every listed moment is exact over `Q`, Galois-invariant, and branch-free;
- none of the audited collective scalar functionals reproduces or constrains `κ_U1 = 6`, the embedded `5/3` normalization, `sin²_* = 3/8`, or the generator-basis `1/4` diagnostic;
- no gauge representation, local field, mass activation, decoupling rule, Dynkin index, threshold beta row, or physical constant is derived.

Gate 161 is therefore a positive spectral-data theorem but a negative physics-permission theorem. The next gate should build the full seven-root finite contact zeta/action functional and test whether that stronger object supplies constraints beyond scalar moment matching.

## v1.60 — Gate 162: Finite Contact Spectral Zeta Regularization

Adds `pkg/bridge/contactzeta`.

Gate 162 extends the Gate 161 collective spectral ledger from positive moments to the full seven-root finite zeta function:

```text
ζ_contact(s) = Σ_i λ_i^(-s)
```

The audited exact zeta values are:

```text
ζ(0) = 7
ζ(1) = 7993/542
ζ(2) = 10529233/293764
ζ(3) = 15529024549/159220088
ζ(4) = 24783201328945/86297287696
```

The quartic contribution is computed by reciprocal Newton identities, without choosing quartic branches:

```text
ζ_q(0) = 4
ζ_q(1) = 2235/271
ζ_q(2) = 1512333/73441
ζ_q(3) = 1177369209/19902511
ζ_q(4) = 998467775217/5393580481
```

Current result:

- the full contact zeta ledger is exact over `Q`, Galois-invariant, branch-free, and pole-free;
- no analytic continuation is needed because the audited spectrum is finite and nonzero;
- zeta/action scalar candidates such as `ζ(0)`, `ζ(1)`, `ζ(2)`, `ζ(1)/7`, `ζ(2)/ζ(1)^2`, `Tr(Ω)ζ(1)/49`, and the full determinant do not reproduce or constrain `κ_U1 = 6`, the embedded `5/3` normalization, `sin²_* = 3/8`, or the generator-basis `1/4` diagnostic;
- the zeta ledger does not define a finite spectral triple, real structure, grading, finite Dirac operator, canonical cutoff function, or gauge-kinetic representation map;
- no threshold beta row, physical coupling, mass, scale, CKM, or PMNS datum is derived.

Gate 162 is therefore a positive finite-zeta theorem but a negative spectral-action permission theorem. The next gate should attempt the logically stronger object: a finite spectral action principle, including the spectral triple data required to interpret zeta values as gauge/action coefficients.

## v1.61 — Gate 163: Finite Spectral Action Principle / Spectral Triple Construction Audit

Adds `pkg/bridge/spectralaction`.

Gate 163 tests the logically stronger move after Gate 162: whether the exact seven-root contact zeta ledger can be promoted into a finite spectral action principle rather than remaining a collection of spectral scalars.

The gate distinguishes exact spectral pre-data from missing spectral-triple structure.

Available pre-data:

```text
contact rows:                   7
zeta values:                    5
exact rational overlap:         true
exact characteristic certified: true
exact root isolation certified: true
positive nonzero spectrum rows: 7
finite zeta poles:              0
analytic continuation needed:   false
rational singleton rows:        3
quartic orbit rows:             4
quartic collective blocks:      1
```

Audited spectral-triple ingredients:

```text
ingredients audited:            11
available ingredients:          4
canonical ingredients:          3
missing required canonical:     8
spectral triple complete:       false
finite Dirac selected:          false
real structure selected:        false
grading selected:               false
order-one calculus verified:    false
canonical cutoff selected:      false
gauge fluctuation map derived:  false
```

Gate 163 also audits five finite Dirac-like candidates:

```text
D = Omega_contact
D = Omega_contact^{-1}
D = Omega_contact - Tr(Omega)/7 I
D = 7 Omega_contact / Tr(Omega)
D_q = spectral scalar on the quartic primary block
```

All five are exact, finite-spectrum, Galois-invariant, and branch-free diagnostics. None is promoted to a spectral-triple Dirac operator because each still requires a canonical algebra representation, real structure `J`, grading `gamma`, order-one calculus, and gauge fluctuation map.

Audited spectral-action ansatzes:

```text
Tr f(D/Lambda) = sum_k f_k Tr(D^k)
sum_s c_s zeta_contact(s)
sum_s c_s zeta_q(s)
c log det(D^2)
Tr f((D+A+JAJ^{-1})/Lambda)
```

All remain formal. None supplies canonical coefficients, gauge kinetic rows, boundary constraints, threshold beta rows, masses, scales, or physical constants.

Result:

- Gate 163 is a positive finite spectral pre-data theorem.
- Gate 163 is a negative finite spectral-action permission theorem.
- The spectral-action route is not rejected; it is postponed until a lawful finite spectral triple is constructed.
- The beta-permission firewall remains closed.
- Residual nullity remains `3 -> 3`.

Next gate: Gate 164 — finite Dirac candidate construction / order-one axiom obstruction audit.
