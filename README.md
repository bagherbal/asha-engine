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
