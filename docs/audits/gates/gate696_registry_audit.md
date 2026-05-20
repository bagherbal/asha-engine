# Gate 696 — Bernoulli Payoff Normalization and Zero-Complement Support Audit

## Registry target

```text
pkg/bridge/generation2bernoullipayoffnormalizationandzerocomplementsupportaudit
```

Registered theorem:

```text
generation2bernoullipayoffnormalizationandzerocomplementsupportaudit.Generation2BernoulliPayoffNormalizationAndZeroComplementSupportAuditTheorem()
```

## Purpose

Gate 695 typed the active bridge as a no-bias `K7` event expectation:

```text
D_base ≈ E_rho72[R_split]
rho_72 = I_H72/72
R_split = S_split P_K7.
```

Gate 696 audits the payoff normalization hidden inside this Bernoulli reading.  The most general two-event observable on the `K7` / complement split is:

```text
R_{a,b}=aP_K7+b(I_H72-P_K7).
```

The gate asks whether the active observable is selected by support-locality: the response is carried only by `K7`, while the complement is zero-response.  Under that condition:

```text
a=S_split,
b=0,
R_{a,b}=S_split P_K7.
```

This is a bridge-layer payoff-normalization audit only. It does not derive boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor, CKM/PMNS, a native state-selection theorem, a native payoff theorem, or a native `7/72` theorem.

## Inherited objects

```text
H_72 = Lambda^4 R^8 ⊕ R^2_boundary
dim(H_72)=72
rho_72 = I_H72/72
P_K7 = Boolean-octonionic support-selected event projector
rank(P_K7)=7
P_perp = I_H72 - P_K7
rank(P_perp)=65
S_split = lambda(Lambda_12)+(R_3-1)
D_base = kappa_lambda+kappa_e+lambda(Lambda_12)
```

Inherited numerical ledger:

```text
S_split ≈ 0.0012924448188162962
E_rho72[S_split P_K7] ≈ 0.0001256543573849177
D_base ≈ 0.0001256552099683575
E_1 ≈ 8.525834398014336e-10
```

## General Bernoulli payoff observable

Define:

```text
R_{a,b}=aP_K7+bP_perp.
```

Under `rho_72`:

```text
E[R_{a,b}]
= Tr(rho_72 R_{a,b})
= (7/72)a + (65/72)b.
```

The active observable is the special case:

```text
a=S_split,
b=0.
```

Then:

```text
E[R_split]=(7/72)S_split.
```

## Affine payoff degeneracy

Expectation alone cannot select the active pair.  The equation

```text
(7/72)a+(65/72)b=(7/72)S_split
```

has many solutions.  For example:

```text
a=0,
b=(7/65)S_split
```

gives the same expectation value but puts the payoff on the complement rather than on `K7`.

Therefore:

```text
FAILED_ROUTE_EXPECTATION_VALUE_ALONE_DOES_NOT_SELECT_PAYOFF_NORMALIZATION.
```

## Support-locality condition

Impose `K7` support-locality:

```text
P_K7 R P_K7 = R,
P_perp R = 0,
R P_perp = 0.
```

For:

```text
R_{a,b}=aP_K7+bP_perp,
```

we have:

```text
P_perp R_{a,b}=bP_perp,
R_{a,b}P_perp=bP_perp,
P_K7 R_{a,b} P_K7=aP_K7.
```

Thus support-locality forces:

```text
b=0.
```

It does not determine `a`.  It only removes the complement payoff.

## Boundary payoff assignment

With `b=0`, the observable becomes:

```text
R_a=aP_K7.
```

The boundary quotient scalar supplies the event payoff:

```text
a=S_split.
```

Therefore:

```text
R_split=S_split P_K7.
```

Gate 696 does not prove why `S_split` is the payoff. It only separates the two roles:

```text
support-locality -> zero complement payoff,
boundary quotient scalar -> K7 event payoff.
```

## Alternative payoff observables

Typed alternatives are audited:

```text
R=S_split I_H72
  expectation=S_split
  rejected: no K7 support.

R=S_split P_perp
  expectation=(65/72)S_split
  rejected: wrong event support.

R=S_split(P_K7-(7/72)I_H72)
  expectation=0
  rejected: fluctuation observable, not active response.

R=S_split(P_+-P_-)
  expectation=(1/72)S_split
  rejected: signed polarity response, inactive.

R=S_split P_K7
  expectation=(7/72)S_split
  active.
```

## Source-type classification

```text
P_K7:
  event support selected by Boolean-octonionic support.

P_perp:
  no-response complement under K7 support-locality.

S_split:
  boundary anti-alignment quotient payoff assigned to the active event.

rho_72:
  full augmented no-bias observer state.

R_split:
  support-local Bernoulli payoff observable.
```

## Firewalls

Gate 696 does not prove:

```text
why history uses support-locality,
why K7 is the event,
why S_split is the payoff,
why expectation equals D_base,
why the residual exists,
or why 7/72 is native.
```

It only proves that once support-locality and `K7` event support are imposed, the zero-complement Bernoulli observable is forced.

## Expected status lines

```text
PASS_GATE695_BERNOULLI_OBSERVABLE_INHERITED
PASS_GENERAL_TWO_PAYOFF_OBSERVABLE_DEFINED
PASS_EXPECTATION_FOR_GENERAL_A_B_COMPUTED
PASS_AFFINE_PAYOFF_DEGENERACY_AUDITED
PASS_SUPPORT_LOCALITY_CONDITION_DEFINED
PASS_SUPPORT_LOCALITY_FORCES_ZERO_COMPLEMENT_PAYOFF
PASS_ACTIVE_PAYOFF_ASSIGNMENT_RECONSTRUCTED
PASS_ALTERNATIVE_PAYOFF_OBSERVABLES_AUDITED
CONDITIONAL_SUPPORT_ACTIVE_RESPONSE_IS_SUPPORT_LOCAL_K7_PAYOFF_OBSERVABLE
CONDITIONAL_SUPPORT_ZERO_COMPLEMENT_PAYOFF_FROM_K7_SUPPORT_LOCALITY
FAILED_ROUTE_EXPECTATION_VALUE_ALONE_DOES_NOT_SELECT_PAYOFF_NORMALIZATION
FAILED_ROUTE_NO_NATIVE_REASON_HISTORY_USES_SUPPORT_LOCALITY
FAILED_ROUTE_NO_NATIVE_REASON_K7_EVENT_RECEIVES_S_SPLIT_PAYOFF
FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM
FIREWALL_PRESERVED_GATE696_BERNOULLI_PAYOFF_SUPPORT_BOUNDARY
```
