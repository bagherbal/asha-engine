# Gate 680 — Global Augmented Trace versus Kernel-Conditional Trace Audit

## Purpose

Gate 679 showed that the natural split projection

```text
pi_split : H_72 -> Q_boundary
```

has kernel

```text
ker(pi_split)=Lambda^4 R^8 ⊕ L_anti,
dim ker(pi_split)=71.
```

Therefore `K_7` is not the full projection kernel. Gate 680 audits the next denominator question: why the active response uses the global full-extension density `7/72` instead of the kernel-conditional density `7/71` or the finite-only density `7/70`.

## Implemented package

```text
pkg/bridge/generation2globalaugmentedtracekernelconditionalaudit
```

Registered theorem:

```text
generation2globalaugmentedtracekernelconditionalaudit.Generation2GlobalAugmentedTraceVersusKernelConditionalTraceAuditTheorem()
```

## Short exact projection sequence

The natural split projection gives the short exact projection ledger:

```text
0 -> ker(pi_split) -> H_72 -> Q_boundary -> 0
```

with:

```text
dim ker(pi_split)=71,
dim Q_boundary=1,
dim H_72=72.
```

The defect inclusion is:

```text
K_7 ⊕ 0_boundary ⊂ ker(pi_split) ⊂ H_72.
```

## Trace normalization alternatives

Gate 680 audits the typed densities:

```text
tau_global = 7/72   full augmented extension H_72
tau_kernel = 7/71   split-projection kernel only
tau_finite = 7/70   finite Lambda^4 R^8 chamber only
tau_half   = 7/144  half-boundary-coordinate clue
```

The active response remains:

```text
D_base ≈ tau_global S_split,
D_base - (7/72)S_split ≈ 8.5258e-10.
```

The kernel-only `7/71` and finite-only `7/70` alternatives are typed but weaker response normalizations.

## Interpretation

The response operator acts on:

```text
Q_boundary = H_72 / ker(pi_split).
```

Therefore the quotient line is part of the response system. Normalizing only over the kernel erases the quotient input; normalizing only over `Lambda^4 R^8` erases the boundary extension. Gate 680 conditionally supports `7/72` as the full-extension defect density of `K_7` in the augmented response chamber.

## Verdict

```text
PASS_GATE679_RELATIVE_TRACE_RESPONSE_INHERITED
PASS_SHORT_EXACT_PROJECTION_SEQUENCE_DEFINED
PASS_K7_DEFECT_INSIDE_KERNEL_CLASSIFIED
PASS_GLOBAL_KERNEL_FINITE_TRACE_NORMALIZATIONS_AUDITED
PASS_RESPONSE_COMPATIBILITY_WITH_QUOTIENT_LINE_AUDITED
CONDITIONAL_SUPPORT_GLOBAL_H72_TRACE_IS_TYPE_CORRECT_FOR_QUOTIENT_RESPONSE
CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_IS_FULL_EXTENSION_DEFECT_DENSITY
FAILED_ROUTE_KERNEL_CONDITIONAL_TRACE_7_OVER_71_NOT_ACTIVE_RESPONSE_NORMALIZATION
FAILED_ROUTE_FINITE_ONLY_TRACE_7_OVER_70_OMITS_BOUNDARY_QUOTIENT_INPUT
FAILED_ROUTE_NO_NATIVE_GLOBAL_TRACE_RESPONSE_PRINCIPLE
FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM
FAILED_ROUTE_NO_NATIVE_TRACE_TO_BOUNDARY_QUOTIENT_RESPONSE_THEOREM
FIREWALL_PRESERVED_GATE680_GLOBAL_TRACE_NORMALIZATION_BOUNDARY
```
