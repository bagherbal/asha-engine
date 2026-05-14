# Gate 190 Registry Audit — Eta-odd scalar-orientation source / matter-pullback search audit

## Package

`pkg/bridge/scalarorientationsource`

## Theorem

`BRIDGE-ETA-ODD-SCALAR-ORIENTATION-SOURCE-MATTER-PULLBACK-SEARCH-AUDIT`

## Status

`FAILED_ROUTE` as a constructive eta-source search, with a positive obstruction theorem.

## Purpose

Gate 189 proved that the abstract branchwise projector pair `{P_A,P_B}` and the physical `H_Phi` high/low projector pair are dimensionally compatible, but no canonical `eta -> high/low` assignment or unique scalar-bundle trivialization was derived.

Gate 190 audits whether any already-derived finite source can break the `eta -> -eta` symmetry lawfully:

```text
P_A <-> P_B
P_high <-> P_low
eta -> -eta
```

The gate tests weak isospin, scalar hypercharge, SU(2) plane-swap/Weyl action, charge conjugation, B-L/Fock charge, scalar complex structure, contact signed diagnostics, and broken-sector/gauge-eating diagnostics.

## Main theorem result

| Claim | Result | Meaning |
|---|---:|---|
| Gate 189 compatibility inherited | Pass | branchwise maps still exist after a choice |
| Weak isospin / scalar hypercharge audited | Pass | `T3L` and `Y_phi` preserve high/low planes; `T1/T2` mix planes as gauge action |
| SU(2) Weyl-style plane swap audited | Pass | exchanges high/low planes; exchange is not selection |
| Charge conjugation audited | Pass | charge conjugation is an involution that exchanges orientations, not a branch selector |
| B-L/Fock pullback audited | Obstructed | `1+3` matter polarization, not `2+2` scalar eta source |
| Contact signed diagnostics audited | Obstructed | signed/contact-current diagnostics are not physical eta-odd scalar pullbacks |
| Broken-sector diagnostics audited | Obstructed | W/Z/photon and Goldstone signatures use diagnostic vacuum convention; cannot select eta retroactively |
| Gauge-invariant eta-odd source | Not found | no finite observable selects `eta -> high` over `eta -> low` |
| Physical scalar bundle | Not derived | requires explicit spontaneous orientation and gauge-frame seal |

## Weak isospin / hypercharge audit

The gate verifies on the physical scalar frame:

```text
[T3L, P_high] = 0
[T3L, P_low]  = 0
[Y_phi, P_high] = 0
[Y_phi, P_low]  = 0
```

while the non-diagonal weak generators mix the two planes:

```text
[T1, P_high] != 0
[T2, P_high] != 0
```

The SU(2)-style Weyl/plane-swap operation exchanges:

```text
P_high <-> P_low
```

but does not select either orientation. Therefore weak gauge structure explains why the orientation is gauge/spontaneous data; it does not provide a gauge-invariant eta-odd selector.

## Charge conjugation audit

Existing contact sign-source gates already proved that charge conjugation acts as an involution exchanging orientations with zero selected branches. Gate 190 imports that result and compares it with the eta swap.

Result:

```text
charge conjugation mirrors eta -> -eta
but selects no eta sign
```

The Higgs-conjugate quotient correction is also inherited: actual Gate-25 Yukawa support does not collapse into Higgs-conjugate pairs, so conjugation does not solve scalar high/low orientation.

## Source search table

| Candidate | Available | Eta-odd | Selects eta? | Verdict |
|---|---:|---:|---:|---|
| B-L / Fock charge | yes | no | no | wrong tensor factor and wrong shape: `1+3`, not scalar `2+2` |
| Weak isospin `T3L` | yes | no | no | preserves high/low projectors; no branch pullback |
| Scalar hypercharge `Y_phi` | yes | no | no | pair-preserving representation datum |
| SU(2) Weyl / plane swap | yes | no | no | exchanges planes as gauge symmetry; exchange is not selection |
| Charge conjugation | yes | no | no | proves Z2 degeneracy; no C-breaking eta pullback |
| Scalar complex structure | yes | no | no | pair-compatible but noncanonical orientation |
| Centered contact signed diagnostic | yes | no | no | diagnostic signed current, not physical scalar eta source |
| Broken generator / covariant derivative diagnostic | yes | no | no | depends on diagnostic vacuum orientation |
| Observed high/low assignment | forbidden | yes | no | blocked by finite-derivation firewall |

## Firewall

```text
observed physical input: no
numeric root approximation: no
individual quartic root diagonalization: no
arbitrary eta-to-high/low assignment: no
Gate 189 compatibility inherited: yes
weak/hypercharge source audited: yes
charge conjugation source audited: yes
broken-sector source audited: yes
contact signed source audited: yes
eta-odd finite source found: no
gauge-invariant eta-odd source found: no
canonical eta orientation derived: no
eta orientation classified spontaneous/gauge: yes
physical scalar bundle derived: no
Chern-Weil carrier: not derived
heat-kernel matching: not derived
threshold beta rows: not derived
absolute coupling promotion: not derived
physical constants: not derived
strict nullity: 3 -> 3
conditional eta-source-search nullity: 1 -> 0
```

## Validation

Focused theorem tests were run in two batches because one long combined command hit the outer tool timeout after several packages had already passed.

Primary focused package:

```bash
go test -p=1 ./pkg/bridge/scalarorientationsource -count=1 -timeout=300s
```

Passed.

Focused dependency batch 1:

```bash
go test -p=1 \
  ./pkg/bridge/scalarorientationsource \
  ./pkg/bridge/scalarbundlemap \
  ./pkg/bridge/branchprojector \
  ./pkg/bridge/resolventvacuum \
  ./pkg/bridge/scalarcovariant \
  ./pkg/bridge/scalarcomplex \
  ./pkg/bridge/scalarsu2 \
  ./pkg/matter/hypercharge \
  ./pkg/matter/su2lgauge \
  ./pkg/bridge/contactsignsource \
  -count=1 -timeout=300s
```

Passed before the outer command timeout.

Focused dependency batch 2:

```bash
go test -p=1 \
  ./pkg/bridge/contactcoddsource \
  ./pkg/bridge/higgsconjugatequotient \
  ./pkg/bridge/gaugeeating \
  ./pkg/bridge/brokenmetric \
  -count=1 -timeout=300s
```

Passed.

Compile smoke:

```bash
go test -p=1 ./internal/app ./cmd/asha -run '^$' -count=1 -timeout=300s
```

Passed.

Full theorem ladder smoke:

```bash
timeout 120s go run ./cmd/asha
```

Completed and printed Gate 190 successfully.

No full historical `go test ./...` suite was run.

## Verdict

Gate 190 closes the eta-source search. No finite weak, hypercharge, matter, contact, charge-conjugation, scalar-complex, or broken-sector datum can lawfully select `eta -> high` rather than `eta -> low`.

Therefore the eta orientation is now precisely localized as spontaneous/gauge-frame data, not as a missing internal finite observable.

## Recommended next gate

Gate 191 — spontaneous scalar-orientation seal / gauge-fixed `H_Phi` trivialization axiom audit.

This gate should explicitly record the spontaneous orientation insertion, construct the conditional gauge-fixed scalar-bundle trivialization, and keep all physical constants, Chern-Weil, heat-kernel, and threshold promotions sealed until that conditional map is cleanly established.
