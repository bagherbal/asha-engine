# ASHA Final Result — Formula Ledger and Boundary Board

**Status:** final theorem-board ledger after Gates 0–425.  
**Purpose:** compact formula-first record of the current ASHA result, including native coefficients, bridge coefficients, quarantined inputs, and explicit physical firewalls.  
**Scope:** this document is a scientific ledger, not a narrative paper. The manuscript lives in `docs/paper/final/`.

---

## 0. One-line closure

ASHA closes as a theorem-gated finite-geometric **law-space**:

```math
C\ell(1,7)
\Rightarrow K_7
\Rightarrow A_F=\mathbb C\oplus\mathbb H\oplus M_3(\mathbb C)
\Rightarrow D_{M\times F}
\Rightarrow S_{\rm CCM}
\Rightarrow {\rm SM+gravity~law\ space}
```

with a sealed tree-level Higgs proxy

```math
m_H^{\rm tree}\approx124.925\ {\rm GeV},
```

and with explicit environmental frontiers

```math
\dim \mathcal M_{\rm charged}^{\rm native}=13,
\qquad
(\Omega_{\rm DM}h^2,\rho_\Lambda,t_{\rm universe})\notin{\rm native~ASHA~output}.
```

---

## 1. Complete board formula

### 1.1 Native law-space action

The almost-commutative product data are

```math
A=C^\infty(M)\otimes A_F,
\qquad
A_F=\mathbb C\oplus\mathbb H\oplus M_3(\mathbb C),
```

```math
H=L^2(M,S)\otimes H_F,
\qquad
D=D_M\otimes1_F+\gamma_5\otimes D_F,
```

```math
J=J_M\otimes J_F,
\qquad
\gamma=\gamma_M\otimes\gamma_F.
```

The fluctuated Dirac operator is

```math
D_A=D+A+JAJ^{-1},
\qquad
A=\sum_i a_i[D,b_i].
```

The board-level action is

```math
S_{\rm ASHA}
=
{\rm Tr}\,f\!\left(\frac{D_A}{\Lambda}\right)
+
\langle\Psi,D_A\Psi\rangle
+
S_{\rm seals},
```

where `S_seals` is not native dynamics; it records quarantined extension/boundary data.

### 1.2 Fully separated representation

```math
S_{\rm ASHA}
=
S_{\rm native}(C\ell(1,7),K_7,A_F,D_{M\times F})
+
S_{\rm bridge}(f_0,f_2,f_4,\Lambda,v_{\rm Pf})
+
S_{\rm family}^{\rm sealed}(K,X,Y;c_s)
+
S_{\rm env}(\mathcal M_{13},\Omega_{\rm DM},\rho_\Lambda,\ldots).
```

The four coefficient classes are:

| Class | Symbols | Status |
|---|---:|---|
| Native finite geometry | `P_B`, `P_G`, `K_7`, `A_F`, edge counts, charge tables | theorem/audit native |
| Bridge/spectral-action lane | `f_0,f_2,f_4,Λ`, continuum heat-kernel interpretation, RG transport | bridge-required |
| Quarantined family axiom | `K_gen,X_gen,Y_gen`, symbolic `c_s` | conditional capacity only |
| Environmental coordinates | Yukawa values, CKM/PMNS numbers, cosmological history | not predicted |

---

## 2. Native finite measurement ladder — Gates 0–2

### 2.1 Exterior algebra

```math
\dim \Lambda^k\mathbb R^8=\binom{8}{k},
\qquad
\sum_{k=0}^{8}\binom{8}{k}=2^8=256.
```

Grade dimensions:

```math
[1,8,28,56,70,56,28,8,1].
```

The middle chamber is

```math
\dim\Lambda^4\mathbb R^8=70.
```

### 2.2 Clifford algebra

```math
e_ie_j+e_je_i=2\eta_{ij},
\qquad
\eta={\rm diag}(+,-,-,-,-,-,-,-),
```

```math
\dim C\ell(1,7)=2^8=256.
```

### 2.3 Covariant phase-space bookkeeping

```math
(x_0,x_1,x_2,x_3,p_0,p_1,p_2,p_3),
\qquad
\dim=8.
```

**Boundary:** this is the finite measurement ladder, not yet spacetime dynamics and not a flavor theorem.

---

## 3. Boolean/G2 contact vacuum — Gates 3–6

### 3.1 Boolean incidence support

```math
W:\Lambda^3\mathbb R^8\rightarrow\Lambda^4\mathbb R^8,
\qquad
P_B=WW^T,
```

```math
W^TW=I_{56},
\qquad
P_B^2=P_B,
\qquad
{\rm Tr}(P_B)=56.
```

Harmonic complement:

```math
70-56=14.
```

### 3.2 Octonionic `G_2` calibration support

```math
P_G^2=P_G,
\qquad
{\rm Tr}(P_G)=14.
```

### 3.3 Contact vacuum

```math
K={\rm Im}(P_B)\cap{\rm Im}(P_G),
\qquad
\dim K=7.
```

Contact index:

```math
I_{BG}=\frac{\dim K}{7}=1.
```

### 3.4 B-sector quadratic action

```math
O_B=W^T(I-P_G)W,
```

```math
S_B[b]=\|(I-P_G)Wb\|^2=b^TO_Bb,
```

```math
\ker O_B=K,
\qquad
\dim\ker O_B=7,
```

```math
\lambda_{B,+}^{\min}\approx0.1024649212.
```

**Meaning:** `K_7` is a finite zero-energy contact vacuum selected by a positive semidefinite action.

---

## 4. Early gauge/Higgs repair — Gates 7–12

### 4.1 Tangent-level electroweak shape

Contact-preserving centralizer:

```math
\mathfrak g_2^R=\{X\in\mathfrak g_2\mid [X,R]=0\},
```

```math
\dim \mathfrak g_2^R=4,
\qquad
\mathfrak g_2^R\cong\mathfrak{su}(2)\oplus\mathfrak u(1).
```

Naive Boolean-compressed closure fails: tangent shape does not directly become the finite gauge theorem.

### 4.2 Projected-connection curvature identity

```math
A=PAP+PAQ+QAP+QAQ,
```

```math
P[A,B]P-[PAP,PBP]=PAQBP-PBQAP.
```

This identifies off-diagonal curvature/leakage as the correct repair channel.

### 4.3 Finite Higgs/vacuum mixing

```math
\Phi_i=P_CA_iP_K+P_KA_iP_C,
```

```math
M_K=\sum_i\Phi_i^T\Phi_i\big|_K.
```

Audit result:

```math
{\rm rank}(M_K)=4,
\qquad
7-4=3\ {\rm protected~contact~directions}.
```

### 4.4 Scalar spectrum and potential shape

```math
\tau={\rm Tr}(M_K)=1.1333333333,
```

```math
{\rm Tr}(M_K^2)=0.3325,
```

```math
\lambda_{\rm shape}
=\frac{{\rm Tr}(M_K^2)}{{\rm Tr}(M_K)^2}
=0.2588667820.
```

Active scalar response:

```math
4\ {\rm real~directions}=2\ {\rm complex~Higgs-like~components}.
```

Normal form:

```math
V(r)=\lambda_{\rm shape}(r^2-r_0^2)^2,
\qquad
r_0^2=1.1333333333.
```

---

## 5. Matter carrier and charge skeleton — Gates 13–26

### 5.1 Fock/Witt matter carrier

```math
\mathcal F=\Lambda^*(\mathbb C^4),
\qquad
\dim\mathcal F=2^4=16.
```

Baryon-minus-lepton charge:

```math
B-L=-N_0+\frac13(N_1+N_2+N_3).
```

The tensor split is

```math
H=H_{\rm Fock}\otimes H_\Phi,
\qquad
\dim H=16\times4=64.
```

Charge/scalar separation:

```math
Q_{\rm total}=Q_{B-L}\otimes I_\Phi,
\qquad
S_{\rm total}=I_{\rm Fock}\otimes S_\Phi,
```

```math
[Q_{\rm total},S_{\rm total}]=0.
```

### 5.2 Chirality and hypercharge

Finite parity:

```math
\Gamma_F=(-1)^N.
```

Matter-side candidate:

```math
T_R^3=\frac12-N_0.
```

Hypercharge:

```math
Y=T_R^3+\frac12(B-L).
```

Electric charge:

```math
Q=T_L^3+Y.
```

Odd branch audit matches right-singlet/conjugate hypercharges:

```math
Y\in\{-1,-2/3,-1/3,0,1/3,2/3,1\}.
```

Left-doublet hypercharges:

```math
Q_L:Y=\frac16\quad(3\ {\rm colors}\times2),
\qquad
L_L:Y=-\frac12\quad(2\ {\rm leptons}).
```

### 5.3 `SU(2)_L` ladder

```math
[T_3,T_+]=T_+,
\qquad
[T_3,T_-]=-T_-,
\qquad
[T_+,T_-]=2T_3.
```

### 5.4 Gauge-compatible Yukawa channels

Selection rule:

```math
Y_R=Y_L+Y_\Phi.
```

One-generation allowed channels:

```math
u_L\otimes\Phi_+\rightarrow\nu_R,
\qquad
e_L\otimes\Phi_-\rightarrow e_R,
```

```math
u_L^c?\ {\rm not~promoted};
\qquad
u_R^c\ {\rm only~under~separate~Majorana/seesaw~audit},
```

```math
u,d,u,e\ {\rm channel~types}:\quad 3+3+1+1=8\ {\rm minimal~one-generation~channels}.
```

Three-generation lift:

```math
8\times3=24\ {\rm diagonal~generation~channels},
```

```math
8\times3\times3=72\ {\rm charge-compatible~generation-mixing~maps}.
```

**Boundary:** channel permission is not amplitude derivation.

---

## 6. Electroweak normalization — Gates 40–43, 74, 100–102

Hypercharge normalization:

```math
k_Y=\frac{{\rm Tr}(Y^2)}{{\rm Tr}(T_3^2)}=\frac53.
```

Boundary weak angle:

```math
\sin^2\theta_* =\frac{1}{1+k_Y}=\frac38.
```

Formal RG transport ledger:

```math
\frac1{g_i^2(\mu)}
=
\frac1{g_i^2(M_*)}+B_i\log\frac{M_*}{\mu}.
```

Continuum one-loop diagnostic:

```math
(b_1,b_2,b_3)=\left(\frac{41}{10},-\frac{19}{6},-7\right).
```

**Boundary:** `3/8` is a boundary normalization result, not the observed low-energy weak angle without RG/threshold matching.

---

## 7. Finite spectral triple and inner fluctuations — Gates 272–299

Correct finite algebra:

```math
A_F=\mathbb C\oplus\mathbb H\oplus M_3(\mathbb C).
```

Morita bimodule form:

```math
H_{ij}=V_i\otimes V_j^*.
```

Product Dirac:

```math
D_{\rm total}=D_M\otimes1_F+\gamma_5\otimes D_F.
```

Inner fluctuations:

```math
D_A=D_F+A+JAJ^{-1},
\qquad
\Omega_D^1(A_F)={\rm span}\{a[D_F,b]\}.
```

Field inventory:

```math
U(1)_Y\times SU(2)_L\times SU(3)_C,
```

```math
12\ {\rm gauge~directions}+1\ {\rm complex~Higgs~doublet}.
```

**Boundary:** this gives field inventory and representation consistency, not Yukawa amplitudes.

---

## 8. Spectral-action coefficient lane — Gates 376–387

### 8.1 CCM spectral action form

```math
S_{\rm CCM}={\rm Tr}\,f(D/\Lambda).
```

Schematic bosonic expansion:

```math
S_{\rm CCM}=\int_M\sqrt g\,d^4x\,\pi^{-2}
\left[
48f_4\Lambda^4-f_2\Lambda^2c+\frac{f_0}{4}d
+\frac{96f_2\Lambda^2-f_0c}{24}R
+f_0a|D_\mu\phi|^2
-\frac{f_2\Lambda^2a}{2}|\phi|^2
+\frac{f_0e}{2}|\phi|^4
+\cdots
\right].
```

Finite trace symbols:

```math
a={\rm Tr}(Y^\dagger Y),
\qquad
c={\rm Tr}(D_F^2),
\qquad
d={\rm Tr}(D_F^4),
\qquad
e={\rm Tr}((Y^\dagger Y)^2).
```

Einstein normalization:

```math
C_R=\frac{96f_2\Lambda^2-f_0c}{24\pi^2},
\qquad
\frac{C_R}{M_P^2}=\frac12.
```

Leading relation:

```math
f_2\left(\frac{\Lambda}{M_P}\right)^2=\frac{\pi^2}{8}.
```

This corrected the earlier `π/64` route by

```math
\frac{\pi^2/8}{\pi/64}=8\pi.
```

### 8.2 Higgs one-form edge measure

The Higgs is a finite one-form:

```math
A_F=\sum_i a_i[D_F,b_i].
```

Therefore the scalar kinetic measure is edge-supported:

```math
A_F=P_EA_FP_E,
\qquad
{\rm Tr}_{H_F}(A_F^\dagger A_F)={\rm Tr}_E(A_F^\dagger A_F).
```

Node/edge ledger:

```math
N_{\rm node}=7,
\qquad
N_{\rm edge,J}=10,
\qquad
R_{\rm edge}=\frac{7}{10}\frac{1197}{4624}.
```

Raw trace conversion:

```math
a_{\rm edge}=\frac{10}{7}a_{\rm node},
\qquad
e_{\rm edge}=\frac{10}{7}e_{\rm node},
```

```math
\left(\frac{e}{a^2}\right)_{\rm edge}
=\frac{7}{10}\left(\frac{e}{a^2}\right)_{\rm node}
=\frac{7}{10}\frac{1197}{4624}.
```

### 8.3 Pfaffian scale lane

```math
v_{\rm Pf}=M_P\,2^{3/2}e^{-4\pi^2}.
```

Numerical value in the project convention:

```math
v_{\rm Pf}\approx247.153\ {\rm GeV}.
```

### 8.4 Tree-level Higgs proxy

```math
\lambda_H
=
\frac{\pi^2}{2\cdot10}\frac{1197}{4624}
=
\frac{\pi^2(1197/4624)}{20}
\approx0.12774563655.
```

```math
m_H^{\rm tree}=v_{\rm Pf}\sqrt{2\lambda_H}
\approx124.925370288\ {\rm GeV}.
```

**Historical repair:** old node-style NCG normalization produced a Higgs mass near `170 GeV`. ASHA's one-form edge support changes the scalar normalization to the edge measure and yields the `~124.9 GeV` tree proxy without adding a new scalar field.

**Boundary:** this is not a full pole-mass theorem. Required future inputs:

```math
\text{RG running}+\text{threshold matching}+\text{self-energy/pole conversion}.
```

---

## 9. Thresholds, hidden scale, and dark-sector boundary

B-sector finite gap:

```math
B_{\rm gap}=0.1024649212.
```

Sealed PeV/intermediate ledger includes

```math
M_B\approx1.46774973718\times10^6\ {\rm GeV}.
```

Heavy-sector relic seal path:

```math
\Omega_{\rm heavy}h^2=0
```

only after sealed decay/portal assumptions; otherwise stable heavy relics are unsafe.

Simple stable thermal B-gap Majorana relic interpretation is rejected by overclosure:

```math
\Omega_{B{\rm -gap~Majorana}}h^2\gg\Omega_{\rm DM}h^2,
\qquad
\frac{\Omega_{\rm candidate}}{\Omega_{\rm DM}}\sim1.3\times10^{13}.
```

Thus the B-gap sector is not a simple stable thermal dark-matter theorem.

Holographic IR/UV bridge route:

```math
\rho_\Lambda\sim\frac{M_P^2}{L^2},
```

from the Cohen-Kaplan-Nelson style bound, is a possible continuum/cosmology bridge if the cutoff is promoted to a dynamical dilaton-like field:

```math
\Lambda\rightarrow\Lambda(x)
\quad\text{or}\quad
\Lambda=M_Pe^{-\sigma(x)}.
```

**Boundary:** ASHA currently does not predict

```math
\Omega_{\rm DM}h^2,
\qquad
\rho_\Lambda,
\qquad
t_{\rm universe},
\qquad
\eta_B.
```

---

## 10. Contact spectral algebra and `q4` boundary — Gates 139–183, 393–406

Contact overlap spectral polynomial includes rational blocks and a quartic primary.

Quartic primary:

```math
q_4(x)=3240x^4-7668x^3+6426x^2-2235x+271.
```

Internal contact reconstruction:

```math
C_{q4}=\mathbb Q[x]/(q_4),
\qquad
T_{q4}:p(x)\mapsto xp(x).
```

```math
\chi_{T_{q4}}(x)=q_4(x),
\qquad
\mu_{T_{q4}}(x)=q_4(x).
```

Over `Q`:

```math
q_4\ {\rm irreducible},
\qquad
{\rm nontrivial~idempotents}=0.
```

Failed scalar/edge promotions:

```math
q_4\not\Rightarrow H_\Phi\ {m selector},
\qquad
q_4\not\Rightarrow\Omega_D^1(A_F)\ {m edge~weight},
\qquad
q_4\not\Rightarrow\text{Yukawa~amplitudes}.
```

The `q4` invariant is therefore classified as

```math
q_4\in{\rm contact~spectral~sector},
\qquad
q_4\notin{\rm scalar/flavor~selector~sector}.
```

---

## 11. Scalar/Higgs flavor-blindness — Gates 398–408

Native `H_Φ` data:

```math
\dim_{\mathbb R}H_\Phi=4,
\qquad
H_\Phi\cong\mathbb C^2\quad\text{as weak-doublet arena}.
```

Quaternionic unit endomorphism:

```math
J^2=-I,
\qquad
\mu_J(x)=x^2+1,
\qquad
\chi_J(x)=(x^2+1)^2.
```

Thus a single quaternionic action is quadratic, not quartic.

Native scalar response is pair-degenerate:

```math
\chi_{S_\Phi}(x)=(x-a)^2(x-b)^2,
\qquad
\mu_{S_\Phi}(x)=(x-a)(x-b).
```

`H_Φ` full algebra capacity:

```math
\langle S_\Phi,L_i,R_j\rangle_{\rm alg}={\rm End}_{\mathbb R}(H_\Phi),
\qquad
\dim_{\mathbb R}{\rm End}_{\mathbb R}(H_\Phi)=16.
```

But no native variational selector chooses a nondegenerate element:

```math
\delta V_{\rm native}/\delta A=0
\Rightarrow
A\in\{I,S_\Phi,{\rm pair\text{-}degenerate~family}\},
```

not a flavor texture.

Conclusion:

```math
H_\Phi\ {m is~flavor\text{-}blind~under~native~ASHA~selectors}.
```

---

## 12. Native fermionic generation audit — Gates 393–410

Standard family factor:

```math
H_F^{\rm charged}=H_{\rm one\text{-}generation}\otimes\mathbb C^3_{\rm gen}.
```

Native ASHA currently acts as

```math
O_{\rm native}=O_{\rm one\text{-}generation}\otimes I_3.
```

Therefore

```math
[O_{\rm native},I_3]=0,
\qquad
{\rm Comm}(G_{\rm SM})\supset U(3)_{\rm gen}.
```

Rejected generation origins:

```math
16=8_s\oplus8_c\quad\Rightarrow\quad\text{chirality/two-sector split, not generations},
```

```math
(3\ {\rm spatial/Fock~seeds})\quad\Rightarrow\quad\text{color/spatial semantics, not generations},
```

```math
8_v\leftrightarrow8_s\leftrightarrow8_c\quad\Rightarrow\quad\text{triality arena, not generation functor}.
```

Exact triality texture no-go:

```math
Y=aI+b(\mathbf1-I),
```

```math
\lambda_s=a+2b,
\qquad
\lambda_d=a-b\quad({\rm multiplicity}=2).
```

So exact triality gives `1+2` degeneracy, not three distinct masses.

Native charged moduli firewall:

```math
\dim\mathcal M_{\rm charged}=13.
```

Physical decomposition:

```math
13=6\ {\rm quark~masses}+4\ {\rm CKM~parameters}+3\ {\rm charged\text{-}lepton~masses}.
```

---

## 13. Quarantined family-axiom chain — Gates 411–418

Gate 411 ranked extension candidates. The minimal capacity chain is quarantined as an axiom, not native ASHA.

### 13.1 Hierarchy axiom: `K_gen`

```math
K_{\rm gen}={\rm diag}(-1,0,1).
```

```math
{\rm Tr}(K_{\rm gen})=0,
\qquad
{\rm Tr}(K_{\rm gen}^2)=2,
\qquad
\mu_K(x)=x(x^2-1).
```

KMS state:

```math
\rho_\beta=\frac{e^{-\beta K_{\rm gen}}}{Z(\beta)},
\qquad
Z(\beta)={\rm Tr}(e^{-\beta K_{\rm gen}}).
```

For `β=1`:

```math
\rho\approx(0.665240955775,\,0.244728471055,\,0.09003057317),
```

```math
\frac{\rho_{\max}}{\rho_{\min}}=e^2\approx7.389056098931.
```

Boundary:

```math
K_{\rm gen}\Rightarrow\text{hierarchy capacity},
\qquad
K_{\rm gen}\not\Rightarrow\text{CKM/PMNS}.
```

### 13.2 Mixing axiom: shift/real quadrature

Cyclic shift:

```math
S_{\rm gen}e_1=e_2,
\qquad
S_{\rm gen}e_2=e_3,
\qquad
S_{\rm gen}e_3=e_1,
\qquad
S_{\rm gen}^3=I.
```

Hermitian real quadrature:

```math
X_{\rm gen}=S_{\rm gen}+S_{\rm gen}^T.
```

Noncommutation:

```math
[K_{\rm gen},S_{\rm gen}]\ne0,
\qquad
[K_{\rm gen},X_{\rm gen}]\ne0.
```

Audit norms:

```math
\|[K,S]\|_F\approx2.449489742783,
\qquad
\|[K,X]\|_F\approx3.464101615138.
```

Full real/complex texture capacity is conditional:

```math
\langle K,S\rangle_{\rm alg}=M_3(\mathbb C)
\quad\text{only after family-connection axiom}.
```

### 13.3 Minimal real sector-source axiom

For sector `s∈{u,d,e}`:

```math
M_s=a_sK_{\rm gen}+b_sX_{\rm gen}.
```

Real charged parameter count:

```math
3\ {\rm sectors}\times2=6.
```

Mixing criterion:

```math
[M_u,M_d]=(a_ub_d-b_ua_d)[K_{\rm gen},X_{\rm gen}].
```

CKM real mixing capacity exists iff

```math
a_ub_d-b_ua_d\ne0.
```

But no CP phase.

### 13.4 CP-capable phase axiom

```math
Y_{\rm gen}=i(S_{\rm gen}-S_{\rm gen}^T).
```

```math
M_s=a_sK_{\rm gen}+b_sX_{\rm gen}+c_sY_{\rm gen}.
```

Hermitian family basis:

```math
\dim_{\mathbb R}{\rm Herm}(3)=9.
```

Conditional charged coefficient ledger:

```math
3\ {\rm sectors}\times3=9\ {\rm symbolic~charged~coefficients}.
```

CP-capacity witness:

```math
{\rm Im}\,{\rm Tr}([M_u,M_d]^3)\ne0
```

for generic symbolic sector rays.

Gate-417 diagnostic sample:

```math
{\rm Im}\,{\rm Tr}([M_u,M_d]^3)\approx8.397024.
```

### 13.5 Family closure seal

```math
K/X/Y\Rightarrow\text{hierarchy + mixing + CP capacity},
```

```math
K/X/Y\not\Rightarrow\text{coefficient values}.
```

Final family boundary:

```math
\{a_s,b_s,c_s\}_{s=u,d,e}\quad\text{remain environmental boundary coordinates}.
```

Native ASHA does not promote them:

```math
\dim\mathcal M_{\rm charged}^{\rm native}=13,
\qquad
\dim\mathcal C_{KXY}^{\rm charged}=9\quad({\rm conditional~axiom~ledger}).
```

---

## 14. Publication and artifact closure — Gates 419–425

Gate-419 board:

```text
native law-space → bridge/coefficient lanes → quarantined family axioms → environmental frontiers
```

Gate-420 theorem atlas:

```math
N_{\rm theorem~blocks}=23,
\qquad
N_{\rm dependency~edges}=28.
```

Classification:

```math
7\ {\rm native~nodes}+7\ {\rm bridge~nodes}+4\ {\rm quarantined~axiom~nodes}+2\ {\rm environmental~frontiers}+3\ {\rm failed\text{-}route~boundaries}.
```

Gate-421 manuscript skeleton:

```math
13\ {\rm manuscript~sections},
\qquad
26\ {\rm proof~obligations},
\qquad
4\ {\rm appendices}.
```

Gate-423 reviewer matrix:

```math
12\ {\rm objections}=4\ {\rm high}+6\ {\rm medium}+2\ {\rm low}.
```

Gate-424 artifact index:

```math
N_{\rm gate~audits~indexed}=227.
```

Gate-425 publication preflight:

```text
paper manifest + section-source map + figure-slot ledger + claim/firewall checklist + assembly checklist
```

---

## 15. Inputs, coefficients, and status ledger

### 15.1 Native/audited finite inputs

| Quantity | Value/formula | Status |
|---|---:|---|
| `dim Cℓ(1,7)` | `256` | native algebra |
| Exterior grades | `[1,8,28,56,70,56,28,8,1]` | native combinatorics |
| `rank(P_B)` | `56` | Boolean support |
| `rank(P_G)` | `14` | G2 calibration support |
| `dim K` | `7` | contact vacuum |
| `I_BG` | `1` | contact index |
| `λ_B,+^min` | `0.1024649212` | finite B-gap |
| Active scalar directions | `4` real | Higgs/contact seed |
| Protected directions | `3` | contact residual |
| `Tr(M_K)` | `1.1333333333` | scalar response |
| `Tr(M_K²)` | `0.3325` | scalar response |
| `λ_shape` | `0.2588667820` | scalar normal form |
| `A_F` | `C ⊕ H ⊕ M₃(C)` | finite spectral triple |
| Gauge group | `U(1)_Y × SU(2)_L × SU(3)_C` | inner fluctuations |
| Gauge directions | `12` | inner fluctuations |
| Higgs | `1 complex doublet` | finite one-form |
| `k_Y` | `5/3` | electroweak normalization |
| `sin²θ_*` | `3/8` | boundary weak angle |
| `N_edge,J` | `10` | one-form edge support |
| `N_node` | `7` | contact-node count |

### 15.2 Bridge coefficients

| Quantity | Value/formula | Status |
|---|---:|---|
| `f₂(Λ/M_P)²` | `π²/8` | CCM/Einstein bridge |
| correction factor | `8π` | fixes earlier channel |
| `v_Pf/M_P` | `2^(3/2)e^(-4π²)` | Pfaffian scale lane |
| `v_Pf` | `≈247.153 GeV` | scale bridge |
| `(e/a²)_node` | `1197/4624` | finite trace ratio |
| `(e/a²)_edge` | `(7/10)(1197/4624)` | one-form edge measure |
| `λ_H` | `π²(1197/4624)/20 ≈0.12774563655` | tree proxy |
| `m_H^tree` | `≈124.925370288 GeV` | tree proxy |
| `M_B` | `≈1.46774973718×10⁶ GeV` | sealed heavy-scale ledger |
| Majorana overclosure ratio | `~1.3×10¹³` | stable thermal relic rejection |

### 15.3 Quarantined/imported/environmental inputs

| Quantity | Formula/role | Status |
|---|---|---|
| `K_gen` | `diag(-1,0,1)` | explicit family axiom |
| `S_gen` | cyclic shift | explicit family axiom |
| `X_gen` | `S+S^T` | explicit family axiom |
| `Y_gen` | `i(S-S^T)` | explicit family axiom |
| charged source coefficients | `{a_s,b_s,c_s}_{s=u,d,e}` | environmental |
| Yukawa values | observed masses/textures | not derived |
| CKM angles/phase | observed mixing | not derived |
| PMNS parameters | observed neutrino mixing | not derived |
| `Ω_DM h²` | relic abundance | not derived |
| `ρ_Λ` | observed vacuum energy | not derived |
| cosmological history | expansion/thermal/baryogenesis | not derived |

---

## 16. Non-claims and firewalls

ASHA does **not** currently claim:

```math
\text{Yukawa matrices from pure }C\ell(1,7),
```

```math
\text{CKM/PMNS numerical prediction},
```

```math
\text{loop-corrected Higgs pole mass theorem},
```

```math
\text{dark matter relic abundance prediction},
```

```math
\text{observed cosmological constant prediction},
```

```math
\text{consciousness/biology theorem}.
```

Main firewalls:

```math
\dim\mathcal M_{\rm charged}^{\rm native}=13,
```

```math
\mathcal M_{\rm cosmology}=(\Omega_{\rm DM}h^2,\rho_\Lambda,t_{\rm universe},\eta_B,\ldots)\quad{\rm environmental/history~dependent},
```

```math
q_4\in{\rm contact~spectral~sector},
\qquad
q_4\notin H_\Phi\ {\rm selector},
```

```math
H_\Phi\Rightarrow\text{weak doublet + scalar potential},
\qquad
H_\Phi\not\Rightarrow\text{flavor selector}.
```

---

## 17. Gate coverage index

Essential logical tower:

```text
G0-G2        measurement ladder
G3-G6        Boolean/G2 contact vacuum K7
G7-G12       tangent gauge shape, failure, off-diagonal Higgs repair
G13-G26      Fock matter carrier, charges, SU(2)L, Yukawa channels
G27-G36      triality/flavor first frontier and no-go sequence
G37-G49      scalar potential, RG/threshold/Higgs-mechanism bridge scaffolding
G50-G69      condensate/Fierz/current kernel diagnostics
G70-G102     normalization, U(1), hypercharge, weak-angle boundary refinements
G187-G299    contact spectral, Morita, finite spectral triple, inner fluctuations
G300-G387    product geometry, CCM/Pfaffian/Higgs closure, flavor/cosmology firewalls
G393-G418    final flavor search, q4/Hphi boundary, K/X/Y axiom chain, flavor seal
G419-G425    architecture, theorem atlas, manuscript, artifacts, publication preflight
```

Current canonical artifacts:

```text
docs/audits/gates/          per-gate audit ledger
docs/audits/final/          this final result
docs/summaries/             gates summary + ontological tower map
docs/paper/final/           final manuscript DOCX/PDF
docs/ARTIFACT_INDEX.md      artifact map
docs/REPRODUCIBILITY_CHECKLIST.md
```

---

## 18. Final theorem statement

```math
\boxed{
C\ell(1,7)
\Rightarrow
K_7
\Rightarrow
A_F=\mathbb C\oplus\mathbb H\oplus M_3(\mathbb C)
\Rightarrow
\Omega_D^1(A_F)
\Rightarrow
U(1)_Y\times SU(2)_L\times SU(3)_C+H_\Phi
\Rightarrow
D_{M\times F}
\Rightarrow
S_{\rm CCM}
}
```

with

```math
\boxed{
\sin^2\theta_*=\frac38,
\qquad
v_{\rm Pf}=M_P2^{3/2}e^{-4\pi^2},
\qquad
m_H^{\rm tree}\approx124.925\ {\rm GeV}
}
```

and the honest boundary

```math
\boxed{
\dim\mathcal M_{\rm charged}^{\rm native}=13,
\qquad
\{a_s,b_s,c_s\}_{u,d,e}\ {\rm environmental~under~the~K/X/Y~axiom},
\qquad
\Omega_{\rm DM},\rho_\Lambda,{\rm history}\ {\rm not~derived}.
}
```

