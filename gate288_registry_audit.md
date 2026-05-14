# Gate 288 Registry Audit — Contact-Spectral Cutoff Identification / S_top Branch Selector

## Gate ID

`GATE288-CONTACT-SPECTRAL-CUTOFF-IDENTIFICATION-S-TOP-BRANCH-SELECTOR-AUDIT`

## Purpose

Gate 287 proved that the topological action boundary

```text
S_top = 8π²
```

is meaningful but underdetermined when the spectral-action cutoff moments are treated as free. Gate 288 audits the proposed identification:

```text
cutoff moments = contact spectral zeta/moment values.
```

The test is whether this identification turns the Gate-287 boundary into a branch selector for the Gate-275 amplitude branches.

---

## Inputs Retrieved

### Contact cutoff moments from the contact spectrum

```text
f0 = ζ_contact(0) = 7
f2 = Tr(Ω²)       = 61/25
f4 = Tr(Ω⁴)       = 257629/202500
```

These are treated as exact contact-spectrum data for the audit. The gate does **not** promote the identification `cutoff function = contact spectrum` into a completed heat-kernel theorem.

### Scalar-Morita Dirac moment proxy

From Gates 273 and 275:

```text
κ_C : κ_Q = 1 : 3
X = |x|²
r = |y/x|²
Tr(D_F²) = X(1 + 3r)
Tr(D_F⁴) = X²(1 + 3r²)
```

The scalar shape constraint is:

```text
(1 + 3r²)/(1 + 3r)² = 1197/4624
```

with branches:

```text
r_+ = (3591 + 136√123)/3099 ≈ 1.645470463011191
r_- = (3591 - 136√123)/3099 ≈ 0.672051318208557
```

For the reduced proxy, the identity trace is:

```text
a0 = κ_C + κ_Q = 4
```

This `a0` is a reduced scalar-Morita proxy, not the final physical heat-kernel `a0` theorem.

---

## Quadratic Scale Constraint

Substituting the contact cutoff moments into the topological boundary gives:

```text
7·X²(1+3r²) + (61/25)·X(1+3r) + (257629/202500)·a0 = 8π²
```

For `a0 = 4`, this is a quadratic equation in `X` for each branch.

---

## Branch Sieve Results

The physical positivity condition is:

```text
X = |x|² > 0
```

### Upper branch

```text
r_+ ≈ 1.645470463011191
|y/x|_+ ≈ 1.282758926303454
X_+ ≈ 0.9680658202595966
```

### Lower branch

```text
r_- ≈ 0.672051318208557
|y/x|_- ≈ 0.819787361581378
X_- ≈ 1.905352660102002
```

Both branches admit positive real `X`.

---

## Important Structural Finding

Although the branch-dependent `X` values differ, both branches produce the same total reduced spectral moments:

```text
Tr(D_F²) ≈ 5.746836960723197
Tr(D_F⁴) ≈ 8.549369303330813
Tr(D_F⁴)/Tr(D_F²)² = 1197/4624 ≈ 0.2588667820069204
Tr(D_F⁴)/Tr(D_F²) ≈ 1.4876651907408451
```

Therefore the contact-spectral cutoff identification locks the **total reduced trace moments**, but it does not lock the distribution of the trace between the lepton and quark edge amplitudes.

---

## Status Ledger

```text
CONDITIONAL_SUPPORT_GATE287_S_TOP_UNDERDETERMINATION_INHERITED
CONDITIONAL_SUPPORT_CONTACT_SPECTRAL_CUTOFF_MOMENTS_RETRIEVED
CONDITIONAL_SUPPORT_CONTACT_SPECTRAL_CUTOFF_IDENTIFICATION_AUDITED
CONDITIONAL_SUPPORT_QUADRATIC_SCALE_CONSTRAINT_CONSTRUCTED
CONDITIONAL_SUPPORT_R_BRANCH_POSITIVITY_SIEVE_COMPLETED
CONDITIONAL_SUPPORT_TOTAL_TRACE_MOMENTS_LOCKED_IN_REDUCED_PROXY
CONDITIONAL_SUPPORT_CONTACT_SPECTRAL_CUTOFF_FIREWALLS_PRESERVED
FAILED_ROUTE_BOTH_R_BRANCHES_ADMIT_POSITIVE_REAL_X
FAILED_ROUTE_CONTACT_SPECTRAL_CUTOFF_DOES_NOT_SELECT_R_BRANCH
FAILED_ROUTE_CUTOFF_FUNCTION_EQUALS_CONTACT_SPECTRUM_NOT_DERIVED_AS_HEAT_KERNEL_THEOREM
FAILED_ROUTE_A0_IDENTITY_TRACE_NORMALIZATION_STILL_PROXY_LEVEL
FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED
```

---

## Theorem Verdict

Gate 288 is a meaningful reduction of Gate-287 underdetermination:

```text
free cutoff moments → exact contact cutoff moments
```

But it is **not** a branch selector:

```text
r_+ survives
r_- survives
```

The contact spectrum fixes the total reduced spectral action size, while the two branches remain physically indistinguishable at the level of `Tr(D_F²)` and `Tr(D_F⁴)`.

---

## Firewall Statement

Gate 288 does not claim:

```text
physical heat-kernel normalization
physical a0 coefficient
unique vacuum branch
Higgs mass prediction
Seeley-de Witt a2/a4 theorem
```

The Higgs prediction remains firewalled pending a branch-sensitive observable, full heat-kernel projection, physical `J`, and chiral/hypercharge representation.
