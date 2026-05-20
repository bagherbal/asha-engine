# Gate 654 — P_G-to-Fano Normal-Form Source Theorem Audit

## Purpose

Gate 653 proved the internal implication

```text
Omega = sum_a omega_a wedge eta_a + eta_123,
omega_a wedge omega_b = delta_ab vol_+
=>
b_Omega proportional to P_+ - 3P_-.
```

Gate 654 audits the missing source arrow:

```text
P_G native octonionic/Fano calibration
+
S_K Hodge polarity on K_7
=>
Fano normal form on K_7.
```

This is an internal source-theorem audit only.  It does not derive split-G2,
boundary stress, scalar/flavor transport, physical spacetime, Higgs mass,
CKM/PMNS, gauge unification, or a native `7/72` theorem.

## Package

```text
pkg/bridge/generation2pgtofanonormalformsourcetheoremaudit
```

The theorem entrypoint is:

```go
generation2pgtofanonormalformsourcetheoremaudit.Generation2PGToFanoNormalFormSourceTheoremAuditTheorem()
```

## Audit result

Gate 654 inherits Gate653's symbolic Fano/Hitchin identity and Gate652's finite
P_G/Fano normal-form evidence.  The admissible P_G-sourced pullback is audited
under the Gate634/Gate636 Hodge polarity split:

```text
K_7 = K_7^+ ⊕ K_7^-,
dim K_7^+ = 4,
dim K_7^- = 3.
```

The support decomposition is:

```text
Omega+++ = 0,
Omega++- = A,
Omega+-- = 0,
Omega--- = B.
```

Thus the P_G pullback reduces to the Fano normal-form carrier:

```text
Omega = A + B,
A in Lambda^{2,1},
B in Lambda^{0,3}.
```

The negative component is certified as the oriented negative-sector volume form:

```text
B = eta_1 wedge eta_2 wedge eta_3 = vol_-.
```

The `A` component is audited as a map:

```text
F_A : K_7^- -> Lambda^2(K_7^+)^*,
F_A(eta_a) = omega_a,
F_A^* F_A = alpha I_3.
```

Its image is the calibrated Fano/quaternionic two-form triple on `K_7^+`:

```text
omega_a wedge omega_b = alpha delta_ab vol_+,
J_a J_b = -delta_ab I + epsilon_abc J_c
```

up to the already recorded orientation/sign convention.

## Gauge covariance

The normal form is not treated as a basis-arbitrary coordinate trick.  It is
recorded as SO(3)-gauge covariant under simultaneous rotation:

```text
eta_a   -> R_ab eta_b,
omega_a -> R_ab omega_b.
```

This preserves:

```text
A = sum_a omega_a wedge eta_a,
B = eta_123,
omega_a wedge omega_b = delta_ab vol_+.
```

## Resulting internal chain

Together with Gate653, Gate654 gives the conditional finite/gauge-controlled
source chain:

```text
P_G + S_K
=> Omega = sum_a omega_a wedge eta_a + eta_123
=> b_Omega proportional to P_+ - 3P_-
=> G_hat=(P_+-3P_-)/sqrt(31)
=> cos(theta)=13/sqrt(217), rho^2=48/217.
```

## Verdict

```text
PASS_GATE653_FANO_HITCHIN_SYMBOLIC_IDENTITY_INHERITED
PASS_PG_PULLBACK_SUPPORT_DECOMPOSITION_AUDITED
PASS_PG_PULLBACK_SUPPORT_REDUCES_TO_LAMBDA21_PLUS_LAMBDA03
PASS_PG_FORCES_NEGATIVE_VOLUME_FORM
PASS_A_AS_K7_MINUS_TO_TWO_FORMS_MAP_AUDITED
PASS_OMEGA_A_WEDGE_ORTHONORMALITY_SOURCE_AUDITED
PASS_QUATERNIONIC_TWO_FORM_TRIPLE_SOURCE_AUDITED
PASS_SO3_GAUGE_COVARIANCE_AUDITED
PASS_ROUTE_SOURCE_INDEPENDENCE_AUDITED
CONDITIONAL_SUPPORT_PG_FORCES_FANO_NORMAL_FORM_ON_K7
CONDITIONAL_SUPPORT_INTERNAL_HITCHIN_OBSTRUCTION_MECHANISM_FULLY_SOURCED
CONDITIONAL_SUPPORT_PG_TO_FANO_NORMAL_FORM_SOURCE_THEOREM_SHARPENED
FAILED_ROUTE_NO_FULL_BASIS_FREE_PG_TO_FANO_SOURCE_THEOREM_BEYOND_GAUGE_CONTROLLED_FINITE_AUDIT
FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM
FAILED_ROUTE_NO_PHYSICAL_METRIC_OR_SPACETIME_THEOREM
FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM
FIREWALL_PRESERVED_GATE654_PG_TO_FANO_NORMAL_FORM_BOUNDARY
```

## Interpretation

Gate 654 closes the internal Hitchin obstruction mechanism at the finite,
SO(3)-gauge-controlled source level:

```text
P_G/Fano source
+ Hodge polarity S_K
=> Fano normal form on K_7
=> P_+ - 3P_-
=> 48/217.
```

It still refuses the stronger promotion to a fully basis-free source theorem.
That remaining theorem would have to show directly, without route-normalized
finite frames, that the native P_G calibration canonically selects the same
Fano normal-form package on `K_7`.
