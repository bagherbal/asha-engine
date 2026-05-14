# Gate 296 Registry Audit — Hypercharge Ledger Sieve / Canonical Finite Dirac Assembly

## Gate metadata

- **Gate:** 296
- **Package:** `pkg/bridge/hyperchargediracassembly`
- **Theorem:** `HyperchargeLedgerSieveCanonicalFiniteDiracAssemblyAuditTheorem`
- **Registry status:** `BRIDGE_REQUIRED`
- **Purpose:** Continue from Gate 295 by using the true left/right bimodule representation to audit the hypercharge ledger and assemble the first-order-compatible finite Dirac edge graph.

---

## 1. Inputs inherited from Gate 295

Gate 295 established the zero-order true bimodule:

```text
Q_L ≈ C²_weak ⊗ C³_color
Weak left action:  L(q) = q ⊗ I₃
Color right action: R(B) = I₂ ⊗ Bᵀ
[L(q), R(B)] = 0
```

Gate 296 inherits this as a valid zero-order arena. It does **not** reinterpret weak and color as a single all-left representation of the direct-sum algebra.

**Status:** `CONDITIONAL_SUPPORT_GATE295_TRUE_BIMODULE_INHERITED`

---

## 2. Hypercharge ledger sieve

The audit constructs the symbolic hypercharge variables:

```text
q = Y(Q_L)
u = Y(u_R)
d = Y(d_R)
l = Y(L_L)
e = Y(e_R)
n = Y(ν_R)
h = Y(H)
```

and applies the standard algebraic consistency equations:

```text
Yukawa compatibility:
  u = q + h
  d = q - h
  e = l - h
  n = l + h

SU(2)^2 U(1):
  3q + l = 0

SU(3)^2 U(1):
  2q - u - d = 0

Gravitational trace:
  6q - 3u - 3d + 2l - e - n = 0

Cubic U(1)^3 anomaly:
  6q^3 - 3u^3 - 3d^3 + 2l^3 - e^3 - n^3 = 0

ν_R neutral/Yukawa-compatible branch:
  n = 0
```

The resulting one-parameter ray is:

```text
(q, u, d, l, e, n, h) = (q, 4q, -2q, -3q, -6q, 0, 3q)
```

With the conventional normalization `q = 1/6`, this becomes:

```text
Y(Q_L) = +1/6
Y(u_R) = +2/3
Y(d_R) = -1/3
Y(L_L) = -1/2
Y(e_R) = -1
Y(ν_R) = 0
Y(H)   = +1/2
```

All tested anomaly/Yukawa/unimodularity residuals vanish on the ray:

```text
Yukawa residual          = 0
SU(2)^2 U(1) residual   = 0
SU(3)^2 U(1) residual   = 0
Gravitational residual  = 0
U(1)^3 residual         = 0
Unimodularity residual  = 0
```

However, the absolute normalization `q = 1/6` is **not** derived by this gate. The equations recover the Standard Model hypercharge **ray**, not the absolute unit.

**Supported:**

```text
CONDITIONAL_SUPPORT_HYPERCHARGE_ANOMALY_YUKAWA_EQUATIONS_BUILT
CONDITIONAL_SUPPORT_STANDARD_MODEL_HYPERCHARGE_RAY_RECOVERED
CONDITIONAL_SUPPORT_UNIMODULARITY_TRACE_CANCELLATION_VERIFIED_ON_RAY
```

**Firewalled:**

```text
FAILED_ROUTE_HYPERCHARGE_ABSOLUTE_NORMALIZATION_NOT_DERIVED
FAILED_ROUTE_HYPERCHARGE_FRACTIONS_REQUIRE_NORMALIZATION_SEAL_OR_PRIOR_LEDGER
```

---

## 3. Canonical finite Dirac edge graph

The finite Hilbert slots audited by this gate are:

```text
Q_L : left weak doublet, right color module M3, Y=q
u_R : right weak singlet, right color module M3, Y=q+h
d_R : right weak singlet, right color module M3, Y=q-h
L_L : left weak doublet, right lepton module C,  Y=-3q
e_R : right weak singlet, right lepton module C, Y=l-h
ν_R : right weak singlet, right lepton module C, Y=0
```

The assembled structural finite Dirac shape is:

```text
D_F = [[0, M], [M†, 0]]
M = diag_edges(Y_u ⊗ I3, Y_d ⊗ I3, Y_e, Y_ν)
```

Allowed edge classes:

```text
Q_L ↔ u_R   allowed, shared right M3 module, color identity I3
Q_L ↔ d_R   allowed, shared right M3 module, color identity I3
L_L ↔ e_R   allowed, shared right C module
L_L ↔ ν_R   allowed as Dirac neutrino edge, shared right C module
```

Forbidden or sealed edge classes:

```text
Q_L ↔ e_R      forbidden by right-module mismatch M3 -> C
L_L ↔ u_R      forbidden by right-module mismatch C -> M3
color-changing quark edges forbidden unless the color map commutes with all right M3 actions
ν_R Majorana edge remains sealed behind NeutrinoTextureSeal / B-gap Majorana theorem
```

**Status:**

```text
CONDITIONAL_SUPPORT_CANONICAL_DF_EDGE_GRAPH_ASSEMBLED
FAILED_ROUTE_NUMERICAL_YUKAWA_MATRICES_NOT_DERIVED
FAILED_ROUTE_BGAP_MAJORANA_EDGE_NOT_DERIVED
```

---

## 4. First-order preflight sieve

The first-order condition is restated as:

```text
[[D_F, ρ(a)], ρ°(b)] = 0
```

The derived Morita edge rule is:

```text
An edge H_ij -> H_kl is first-order compatible when the right/opposite module is shared (j = l).
It is non-vacuous when the left module changes (i ≠ k).
```

A concrete color-intertwiner test was performed on the quark edge:

```text
legal color map:       I3
illegal color-changing map: E12
sample color generator: diag(2, -1, 1/2)
```

Residuals:

```text
||I3 B - B I3|| = 0
||E12 B - B E12|| = 3
```

Therefore, the first-order preflight correctly forces quark Yukawa edges to be color-preserving intertwiners.

**Supported:**

```text
CONDITIONAL_SUPPORT_FIRST_ORDER_EDGE_PREFLIGHT_VERIFIED
CONDITIONAL_SUPPORT_COLOR_INTERTWINER_SIEVE_VERIFIED
```

**Still firewalled:**

```text
FAILED_ROUTE_FULL_FIRST_ORDER_SPECTRAL_TRIPLE_NOT_VERIFIED
```

The gate is a preflight because it does not yet evaluate the full first-order condition over all generators of the completed physical finite spectral triple.

---

## 5. Firewall ledger

Gate 296 explicitly refuses to overpromote the result.

Still not derived:

```text
absolute U(1) hypercharge normalization
full numerical Yukawa matrices
B-gap Majorana activation
full anti-linear physical J semantics
full first-order theorem over all algebra generators
canonical D_F amplitudes
Higgs mass ratio
B-gap instanton hierarchy
```

Firewall statuses:

```text
FAILED_ROUTE_HYPERCHARGE_ABSOLUTE_NORMALIZATION_NOT_DERIVED
FAILED_ROUTE_NUMERICAL_YUKAWA_MATRICES_NOT_DERIVED
FAILED_ROUTE_BGAP_MAJORANA_EDGE_NOT_DERIVED
FAILED_ROUTE_FULL_FIRST_ORDER_SPECTRAL_TRIPLE_NOT_VERIFIED
FAILED_ROUTE_HIGGS_AND_BGAP_DYNAMICS_STILL_FIREWALLED
```

---

## 6. Final verdict

Gate 296 successfully upgrades the ASHA finite spectral-triple scaffold from a zero-order true bimodule to a first-order-compatible Dirac edge graph:

```text
CONDITIONAL_SUPPORT_STANDARD_MODEL_HYPERCHARGE_RAY_RECOVERED
CONDITIONAL_SUPPORT_CANONICAL_DF_EDGE_GRAPH_ASSEMBLED
CONDITIONAL_SUPPORT_FIRST_ORDER_EDGE_PREFLIGHT_VERIFIED
```

But it does not complete the finite spectral triple:

```text
FAILED_ROUTE_HYPERCHARGE_ABSOLUTE_NORMALIZATION_NOT_DERIVED
FAILED_ROUTE_NUMERICAL_YUKAWA_MATRICES_NOT_DERIVED
FAILED_ROUTE_BGAP_MAJORANA_EDGE_NOT_DERIVED
FAILED_ROUTE_FULL_FIRST_ORDER_SPECTRAL_TRIPLE_NOT_VERIFIED
```

The correct next gate is to complete the physical hypercharge normalization/opposite-action semantics and then run the full first-order condition on the assembled Dirac edge graph.
