# Gate 840 — RightSocket Character Split and Punctured LeptoColor Compression Audit

## Package

```text
pkg/bridge/generation2rightsocketcharactersplitpuncturedleptocolorcompressionaudit
```

## Registered theorem

```text
generation2rightsocketcharactersplitpuncturedleptocolorcompressionaudit.Generation2RightSocketCharacterSplitPuncturedLeptoColorCompressionAuditTheorem()
```

## Purpose

Gate 840 follows Gate 839's socket-compression obstruction. Gate 839 found the
candidate finite-body location

```text
E tensor W -> (e_t tensor P_3) plus (e_r tensor W)
```

with ranks `3+4=7`, but it did not certify the rank-one socket projectors
`e_t,e_r`.

Gate 840 audits the sharper proposal that the right socket pair

```text
C_R^2
```

is split by two represented `C`-characters, schematically

```text
rho_R(lambda) = diag(lambda, conjugate(lambda)).
```

If this character split is admitted as a representation seal, then the unordered
rank-one socket pair

```text
e_+ plus e_-
```

is source-typed by the represented right-socket action rather than being an
arbitrary basis choice.

This gate does **not** certify the full native `rho_F` action, the dominant/rest
orientation, the socket-compression theorem, `alpha_B`, trace magnitudes, R3, or
R4.

---

## Inherited carrier

Gate 838/839 carrier:

```text
E = C_R^2 plus C_L^2
W = C_lepton plus C_color^3
H_part = E tensor W
```

with:

```text
dim(E)=4
dim(W)=4
dim(H_part)=16
dim(H_F)=32
```

Gate 840 restricts attention to the right lepto-color rectangle:

```text
C_R^2 tensor W
```

with rank:

```text
2 * 4 = 8.
```

---

## Right-socket character split

Candidate sealed character action:

```text
rho_R(lambda)=diag(lambda, conjugate(lambda)).
```

This supports the unordered socket pair:

```text
C_R^2 = e_+ plus e_-
```

with:

```text
rank(e_+) = 1
rank(e_-) = 1
e_+ e_- = 0
e_+ + e_- = I_{C_R^2}.
```

Gate 840 classifies this as:

```text
CONDITIONAL_SUPPORT_RIGHT_SOCKET_CHARACTER_SPLIT_SEAL_CANDIDATE
CONDITIONAL_SUPPORT_E_PLUS_E_MINUS_AS_CHARACTER_PROJECTORS_IF_RHO_R_HAS_LAMBDA_BARLAMBDA_SPLIT
CONDITIONAL_SUPPORT_RANK_ONE_SOCKET_PROJECTORS_NOT_ARBITRARY_IF_CHARACTER_SPLIT_SEALED
```

but preserves:

```text
FAILED_ROUTE_RIGHT_SOCKET_CHARACTER_SPLIT_IS_SEAL_NOT_NATIVE_DERIVATION
FAILED_ROUTE_NO_EXPLICIT_RHO_R_LAMBDA_BARLAMBDA_MATRIX_PROOF_CERTIFIED
FAILED_ROUTE_NO_FULL_RHO_F_ACTION_LEDGER_CERTIFIED
```

---

## Punctured lepto-color support

With an orientation written only as a candidate:

```text
Pi_top  = e_+ tensor P_3
Pi_rest = e_- tensor W
```

Gate 840 audits:

```text
Pi_7 = (e_+ tensor P_3) plus (e_- tensor W)
```

with ranks:

```text
rank(e_+ tensor P_3) = 3
rank(e_- tensor W)   = 4
rank(Pi_7)           = 7.
```

The complement inside the right rectangle is:

```text
e_+ tensor P_1
```

with rank:

```text
1.
```

Therefore:

```text
rank(Pi_7) + rank(e_+ tensor P_1) = 7 + 1 = 8
```

which reconstructs the full right lepto-color rectangle.

Gate 840 classifies this as a support anatomy, not a theorem:

```text
CONDITIONAL_SUPPORT_PUNCTURED_RIGHT_LEPTOCOLOR_RECTANGLE_CANDIDATE
CONDITIONAL_SUPPORT_SELECTED_SUPPORT_HAS_RANK_SEVEN_FROM_THREE_PLUS_FOUR
CONDITIONAL_SUPPORT_EXCLUDED_SINGLETON_HAS_RANK_ONE
FAILED_ROUTE_PUNCTURED_RECTANGLE_IS_SUPPORT_ANATOMY_NOT_COMPRESSION_THEOREM
FAILED_ROUTE_NO_TYPED_SOCKET_PAIR_COMPRESSION_MAP_CERTIFIED
```

---

## B-L puncture conservation

On `W`, inherited:

```text
B-L = -P_1 + (1/3)P_3
```

and:

```text
Tr_W(B-L)= -1 + 3*(1/3)=0.
```

Gate 840 computes the puncture balance:

```text
Tr_{e_+ tensor P_3}(B-L) = 3*(1/3) = 1
Tr_{e_- tensor W}(B-L)   = -1 + 3*(1/3) = 0
Tr_{Pi_7}(B-L)           = 1
Tr_{e_+ tensor P_1}(B-L) = -1
```

Thus:

```text
Tr_{Pi_7}(B-L) + Tr_{e_+ tensor P_1}(B-L) = 1 + (-1) = 0.
```

The full right rectangle remains neutral, but the selected rank-seven support
carries `+1` and the excluded singleton carries `-1`.

Certified support:

```text
CONDITIONAL_SUPPORT_SELECTED_PUNCTURE_HAS_B_MINUS_L_TRACE_PLUS_ONE
CONDITIONAL_SUPPORT_EXCLUDED_SINGLETON_HAS_B_MINUS_L_TRACE_MINUS_ONE
CONDITIONAL_SUPPORT_FULL_RIGHT_RECTANGLE_B_MINUS_L_TRACE_ZERO
```

This is a structural conservation pattern only. It is not a physical particle
assignment.

---

## Orientation firewall

Even if the unordered pair `e_+,e_-` is source-typed by the character split,
Gate 840 does not certify:

```text
e_+ = top socket
e_- = rest socket
```

The ordered dominant/rest assignment remains missing:

```text
FAILED_ROUTE_DOMINANT_REST_ORIENTATION_NOT_CERTIFIED
FAILED_ROUTE_E_PLUS_NOT_IDENTIFIED_WITH_TOP_SOCKET
FAILED_ROUTE_E_MINUS_NOT_IDENTIFIED_WITH_REST_SOCKET
FAILED_ROUTE_NO_DOMINANT_COLOR_SOCKET_SELECTOR
FAILED_ROUTE_NO_REST_LEPTOCOLOR_SOCKET_SELECTOR
```

Possible future selector sources remain outside this gate:

```text
D_F symbolic edge skeleton
finite one-form / Higgs edge dominance
boundary/rest-pressure split
top-dominant trace atom seal
```

---

## Magnitude and sector firewalls

Gate 840 does not derive:

```text
alpha_B
N_eff
C_Yukawa
C_Higgs
observed Yukawa values
masses
CKM
PMNS
three generations
```

The aggregate operator remains, at most, an oriented compression-shadow
candidate:

```text
I_{e_+ tensor P_3}
plus
[alpha_B P_3 - 3 alpha_B^2(B-L)] on e_- tensor W.
```

But no compression map, no trace-magnitude readout, no R3 sector ledger, and no
native Yukawa theorem are certified.

---

## Final verdict

Gate 840 partially resolves Gate 839's missing fine-socket problem by identifying
a sealed right-character source for the unordered rank-one socket pair:

```text
C_R^2 = e_+ plus e_-.
```

It then sharpens the rank-seven aggregate candidate to the punctured right
lepto-color rectangle:

```text
(e_+ tensor P_3) plus (e_- tensor W)
```

with excluded singleton:

```text
e_+ tensor P_1.
```

The B-L balance is exact:

```text
selected support: +1
excluded singleton: -1
full right rectangle: 0.
```

However, this remains a seal-level support anatomy. The ordered dominant/rest
orientation, typed compression map, alpha source, trace-magnitude readout, R3/R4
promotion, and official ledger updates remain blocked.
