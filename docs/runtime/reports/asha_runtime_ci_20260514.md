--- scenario ci text ---
ASHA runtime board — gate425-runtime-20260514 (Gate 425)
Scenario: ci
Source: github.com/bagherbal/asha-engine

== Native finite law-space ==
Finite measurement ladder, Boolean/G₂ contact vacuum, charge skeleton, and inner-fluctuation field inventory.
SYMBOL       VALUE                     STATUS          FORMULA/NOTE
dim Cℓ(1,7)  256                       native/audited  2^8 
grades       [1 8 28 56 70 56 28 8 1]  native/audited  dim Λ^k R^8 = C(8,k) 
rank(P_B)    56                        native/audited   
rank(P_G)    14                        native/audited   
dim K        7                         native/audited  Im(P_B)∩Im(P_G) 
I_BG         1                         native/audited  dim K / 7 
k_Y          5/3                       native/audited  Tr(Y²)/Tr(T₃²) 
sin²θ*       3/8                       native/audited  1/(1+k_Y) 

== Higgs and coefficient bridge ==
One-form edge measure and Pfaffian scale lane yield the tree-level Higgs proxy; pole-mass and RG thresholds remain bridge work.
SYMBOL            VALUE              STATUS           FORMULA/NOTE
f₂(Λ/M_P)²        π²/8               bridge-required  f₂(Λ/M_P)² = π²/8 CCM/Einstein normalization bridge
8π                25.1327412287      bridge-required  (π²/8)/(π/64)=8π correction of earlier coefficient route
v_Pf              247.151135557 GeV  bridge-required  v_Pf = M_P 2^(3/2) exp(-4π²) 
(e/a²)_node       1197/4624          native/audited   1197/4624 
(e/a²)_edge       0.181206747405     bridge-required  (7/10)(1197/4624) 
λ_H               0.12774563655      bridge-required  π²(1197/4624)/20 
m_H^tree          124.925370288 GeV  bridge-required  v_Pf sqrt(2λ_H) tree-level proxy, not pole-mass theorem
M_B               1467749.73718 GeV  bridge-required  sealed B-gap Majorana ledger 
Ω_candidate/Ω_DM  1.3e+13            failed-route     stable thermal B-gap Majorana relic overclosure 
BOUNDARY: Higgs pole mass | m_H^tree + RG + thresholds + self-energy | runtime reports tree proxy only

== Family/flavor frontier ==
Native flavor remains 13-dimensional; quarantined K/X/Y axioms activate hierarchy, mixing, and CP capacity with 9 symbolic charged coefficients.
SYMBOL                VALUE                                              STATUS             FORMULA/NOTE
dim M_charged^native  13                                                 environmental      6 quark masses + 4 CKM + 3 charged-lepton masses native firewall
K_gen                 diag(-1,0,1)                                       quarantined-axiom  K_gen = diag(-1,0,1) hierarchy capacity only
ρ_β                   [0.665240955775, 0.244728471055, 0.0900305731704]  quarantined-axiom  exp(-βK)/Tr exp(-βK) 
ρ_max/ρ_min           7.38905609893                                      quarantined-axiom  exp(2β) 
X_gen                 S+S^T                                              quarantined-axiom  real shift quadrature real mixing capacity
Y_gen                 i(S-S^T)                                           quarantined-axiom  imaginary shift quadrature CP capacity
||[K,S]||_F           2.44948974278                                      quarantined-axiom  sqrt(6) 
||[K,X]||_F           3.46410161514                                      quarantined-axiom  sqrt(12) 
Im Tr([M_u,M_d]^3)    8.397024                                           quarantined-axiom  sample nonzero CP-capacity witness 
dim C_KXY^charged     9                                                  quarantined-axiom  3 charged sectors × 3 symbolic coefficients 
BOUNDARY: charged flavor | dim M_charged^native = 13 | Yukawa values and CKM coordinates are not native ASHA outputs.
BOUNDARY: K/X/Y coefficients | {a_s,b_s,c_s}_{s=u,d,e} | conditional family source coefficients remain boundary data.

== Dark-sector scenarios ==
Heavy finite sectors are scenario-classified. Stable thermal B-gap Majorana relics are rejected by overclosure; decay/nonthermal routes require extra history.
SYMBOL            VALUE              STATUS           FORMULA/NOTE
M_B               1467749.73718 GeV  bridge-required   
Ω_candidate/Ω_DM  1.3e+13            failed-route      
BOUNDARY: B-gap Majorana stable thermal relic | Ω_candidate/Ω_DM ~ 1.3×10¹³ | simple stable thermal interpretation is rejected by overclosure.
BOUNDARY: decaying/portal heavy sector | Ω_heavy h² = 0 after decay assumptions | allowed only with sealed decay/portal dynamics.
BOUNDARY: nonthermal/axion-like routes | requires production-history axiom | not native ASHA output.

== Cosmology boundary ==
The spectral-action cosmological term is present, but observed cosmology requires continuum subtraction, history, or holographic/dilaton bridge data.
BOUNDARY: spectral-action cosmological term | 48 f₄ Λ⁴ + subtraction/renormalization | bare term exists; observed ρΛ needs continuum/history rule.
BOUNDARY: holographic/dilaton bridge | ρΛ ~ M_P²/L², Λ → Λ(x) | possible pathway, not native prediction.
BOUNDARY: cosmological coordinates | (Ω_DM h², ρΛ, t_universe, η_B) | history/state dependent and not predicted by current law-space.

Checks:
- PASS Clifford dimension: dim Cℓ(1,7)=2^8
- PASS Exterior grade dimensions: [1,8,28,56,70,56,28,8,1]
- PASS Boolean/G2 contact vacuum: rank(P_B)=56 rank(P_G)=14 dim K=7
- PASS Scalar shape: Tr(M_K^2)/Tr(M_K)^2
- PASS Hypercharge normalization: k_Y=5/3
- PASS Boundary weak angle: sin²θ*=3/8
- PASS Gauge/Higgs inventory: U(1)_Y × SU(2)_L × SU(3)_C + one complex Higgs doublet
- PASS Pfaffian scale positive: v_Pf computed from Planck mass bridge
- PASS Higgs tree proxy: m_H^tree ≈ 124.925 GeV under project Planck convention
- PASS Majorana stable thermal relic rejected: overcloses by ~1.3e13
- PASS Native charged flavor firewall: dim M_charged^native=13
- PASS KMS family hierarchy capacity: ρβ nontracial for β≠0
- PASS Noncommuting capacity: K does not commute with shift/quadrature
- PASS CP capacity not CP prediction: phase coefficients remain free
- PASS Latest gate marker: gate=425

Verdict: PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data.
