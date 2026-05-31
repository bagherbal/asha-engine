# ASHA Manuscript Figure Toolkit

Reusable figure-generation workspace for the ASHA manuscript and README visual atlas.

The toolkit separates mathematical contracts from layout/rendering. Each figure has:

- a small geometry module in `asha_figures/`,
- a deterministic SVG renderer,
- a build entrypoint in `scripts/`,
- a validation manifest,
- tests that enforce the mathematical contract,
- a Penrose handoff note when direct Penrose rendering is unavailable.

## Implemented figures

### Section 1 — Measurement Ladder `Cℓ(1,7)`

Mathematical contract:

- 9 horizontal tiers.
- Exact Pascal counts: `1, 8, 28, 56, 70, 56, 28, 8, 1`.
- Exact total nodes: `256` basis blades.
- Exact Boolean-lattice cover edges: `1024`.
- Signature cue: blades containing `e0` are warm; pure space-like blades are cyan.
- Boundary: finite measurement bookkeeping, not particle numerology.

### Section 2 — Boolean/G₂ Contact Vacuum `K7`

Mathematical contract:

- Exact ranks: `rank(P_B)=56`, `rank(P_G)=14`, `rank(K7)=7`.
- `K7 = Im(P_B) ∩ Im(P_G)` is the payload.
- Visible areas obey `area = scale × ln(rank)`.
- Boundary: topological/logarithmic projector support visualization, not literal high-dimensional embedding.

### README Figure — Contact Seven and Depth Triple

Mathematical contract:

- `V8 = (x0,p0,x1,p1,x2,p2,x3,p3)`.
- Selecting `x0` as observer-time reference leaves exactly seven directions.
- `V7_contact = R p0 ⊕ Pi1 ⊕ Pi2 ⊕ Pi3`.
- `Pi_i = span(xi,pi)` for `i=1,2,3`.
- `Q_contact^3 ≅ C^3`.
- `N_Q = diag(1/3,1/2,2/3)`.
- `W_Q = exp(-4π N_Q)`.
- Boundary: carrier/depth grammar only, not flavor, CKM/PMNS, or mass eigenvalues.

## Build

```bash
cd asha_figure_toolkit
python scripts/build_section1.py --png
python scripts/build_section2.py --png
python scripts/build_section3.py --png
pytest -q
```

Outputs are written to:

```text
outputs/section1/
outputs/section2/
outputs/section3/
```

## Structure

```text
asha_figures/boolean_lattice.py  # exact Section 1 combinatorics
asha_figures/contact_vacuum.py   # exact Section 2 ranks and log-area scaling
asha_figures/contact_depth.py    # exact README contact-seven/depth-triple contract
asha_figures/style.py            # ASHA visual palette and dimensions
asha_figures/svg.py              # Section 1 SVG renderer
asha_figures/section2_svg.py     # Section 2 SVG renderer
asha_figures/section3_svg.py     # README contact-depth SVG renderer
scripts/build_section1.py        # Section 1 build entrypoint
scripts/build_section2.py        # Section 2 build entrypoint
scripts/build_section3.py        # README contact-depth build entrypoint
config/section*.yaml             # figure contracts
penrose/section*/README.md       # Penrose semantic handoff notes
tests/                           # mathematical validation tests
```

### README Figure — Matter Sockets and Product Depth

Mathematical contract:

- Finite matter algebra visualized as `A_F = C ⊕ H ⊕ M3(C)`.
- Block cell counts: `1, 4, 9`.
- Contact-depth broadcast layers: `N_Q = diag(1/3, 1/2, 2/3)`.
- Depth weights: `exp(-4πN_Q)`.
- Boundary: product-depth grammar only, not particle assignment or flavor theorem.

### README Figure — Locked Constants and Source Alphabet

Mathematical contract:

- `L = 1/(8π)`.
- `S_split = 0.0012924448188162962`.
- Finite source alphabet: `3, 4, 7, 27, 56, 70, 72`.
- Typed coefficient checks: `158=2(72+7)`, `148=2·72+4`, `106=2(56-3)`.
- Boundary: coefficient firewall; no free fitting.

### README Figure — Planck-to-Electroweak Scale Bridge

Mathematical contract:

- `v = M_P exp[-12π + √3/2 + 2S + 148S²]`.
- Uses reduced Planck stiffness seal `M_P = 2.435e18 GeV` for numerical visualization.
- Computed output: `v ≈ 246.219669 GeV`.
- Boundary: physical filling scale bridge, not a native Planck-stiffness theorem.

### README Figure — Higgs Sector: Quartic and Mass Chain

Mathematical contract:

- `lambda_ASHA = 3/8(1+L)(1/3-S)`.
- `m_H = v sqrt(2 lambda_ASHA)`.
- Computed output from the README chain: `m_H ≈ 125.291520 GeV`.
- Boundary: locked physical-filling chain, not a native scalar theorem.

### README Figure — Charged Lepton Anchor and Shape Laws

Mathematical contract:

- Tau action anchor: `A_tau = 4π/3 + 3/10 + 7/72 - S + 1/2(72+27)S²`.
- Shape laws: `K_e = 2/3 - 4(1-2L)S²`, `D_e = sqrt(2π)+2S-4(1-L)S²`.
- Solves the charged-lepton hierarchy as a logarithmic mass ladder.
- Boundary: locked physical-filling formula stack, not a native flavor theorem.

Updated build:

```bash
cd asha_figure_toolkit
python scripts/build_next5.py --png
pytest -q
```


### Section 15 / Master Snapshot — Programmatic Projector Universe

- Figure: `asha_projector_universe_x4_p4_one_trajectory`
- Contract: `V8 = X4 ⊕ P4`, `η_ASHA = diag(+1,-1,-1,-1,-1,-1,-1,-1)`, `Ω = Σ dp_μ ∧ dx^μ`
- Projectors: `Π_X = diag(1,1,1,1,0,0,0,0)`, `Π_P = diag(0,0,0,0,1,1,1,1)`
- Flow: canonical Hamiltonian movement of the seed “One”
- Render: text-free SVG/PNG; formulas and validation are preserved in JSON/manifest.
- Validation: `PASS_PROGRAMMATIC_X4_P4_PROJECTOR_ONE_TRAJECTORY`

## Section 16 — Programmatic Zoom: All Six ASHA S-Terms Visible

- Figure: `asha_projector_universe_zoom6s_all_visible`.
- Text-free render of the zoomed `V8 = X4 ⊕ P4` One-trajectory.
- Makes all six low-energy ASHA action sectors visible in one view:
  `S_grav`, `S_gauge`, `S_Higgs^ASHA`, `S_fermion`, `S_Yukawa^ASHA`, `S_nu^seesaw`.
- Validation: `PASS_ZOOM6S_ALL_SIX_ACTION_SECTORS_VISIBLE`.
