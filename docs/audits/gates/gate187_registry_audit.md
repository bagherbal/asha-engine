# Gate 187 Registry Audit — Resolvent-vacuum order parameter / spontaneous Higgs 2+2 pairing audit

## Package

`pkg/bridge/resolventvacuum`

## Theorem

`BRIDGE-RESOLVENT-VACUUM-SPONTANEOUS-HIGGS-PAIRING-AUDIT`

## Status

`BRIDGE_REQUIRED`

## Purpose

Gate 186 proved that the exact quartic contact module cannot be canonically identified with the Gate-37 pair-degenerate Higgs/scalar carrier without choosing one of three 2+2 partitions. Gate 187 turns that obstruction into the lawful finite object: the exact resolvent-vacuum algebra whose three branches are the three possible scalar pairings.

## Exact input

Quartic contact polynomial:

```text
q4(x) = 3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271
```

Monic form:

```text
x^4 - (71/30)x^3 + (1071/540)x^2 - (149/216)x + 271/3240
```

Gate 186 resolvent cubic:

```text
r3(z) = z^3 - (119/60)z^2 + (8411/6480)z - 1637467/5832000
```

Integer form:

```text
5832000z^3 - 11566800z^2 + 7569900z - 1637467 = 0
```

## Semantic correction preserved by Gate 187

The classical cubic resolvent above has roots:

```text
r1r2 + r3r4
r1r3 + r2r4
r1r4 + r2r3
```

not pair-sum products. Each root still labels one unordered 2+2 partition:

```text
12|34
13|24
14|23
```

This does not change Gate 186's obstruction theorem or coefficients; it corrects the interpretation used for the next construction.

## Derived object

```text
R_pair = Q[z] / r3(z)
```

The engine derives the threefold scalar-vacuum pairing orbit branch-free. It does not select a root.

## Branchwise quadratic factor scheme

For monic

```text
q4(x) = x^4 + a x^3 + b x^2 + c x + d
```

and a resolvent branch `z`, define formal pair data:

```text
u + v = -a
uv = b - z
p + q = z
pq = d
u q + v p = -c
```

Then branchwise:

```text
q4(x) = (x^2 - u x + p)(x^2 - v x + q)
```

This gives the conditional 2+2 scalar shape without diagonalizing individual quartic roots.

## Audit result

| Claim | Status |
|---|---|
| Gate 186 selector obstruction inherited | PASS |
| Exact `R_pair = Q[z]/r3` algebra derived | PASS |
| Threefold degenerate vacuum orbit derived | PASS |
| Unique canonical vacuum selected | NOT CLAIMED |
| Branchwise two-quadratic 2+2 scalar shape opened | PASS |
| Ordered quadratic factors require further adjunction/orientation | PASS |
| Gate-37 Higgs compatibility is conditional | PASS |
| Physical scalar bundle derived | NOT CLAIMED |
| Canonical complex/symplectic structure derived | NOT CLAIMED |
| Chern-Weil / heat-kernel / threshold / constants promotion | BLOCKED |

## Firewall

```text
observed physical input: no
branch diagonalization: no
arbitrary pairing choice: no
canonical 2+2 selector: not derived
spontaneous branch data: quarantined
conditional scalar carrier: opened
physical scalar bundle: not derived
canonical scalar projector: not derived
Chern-Weil carrier: not derived
heat-kernel matching: not derived
threshold beta rows: not derived
absolute coupling promotion: not derived
physical constants: not derived
strict nullity: 3 -> 3
conditional nullity: 2 -> 2
```

## Validation

Focused tests:

```bash
go test -p=1 ./pkg/bridge/resolventvacuum ./pkg/bridge/scalarcontactselector ./pkg/bridge/quarticscalaroperator ./pkg/bridge/cliffordcontactcommutant -count=1 -timeout=90s
```

App compile smoke test:

```bash
go test -p=1 ./internal/app ./cmd/asha -run '^$' -count=1 -timeout=300s
```

Executable theorem ladder smoke test:

```bash
timeout 45s go run ./cmd/asha
```

The theorem ladder reached and printed Gate 187 successfully.

## Next gate

Gate 188 — branchwise quadratic idempotent / scalar-projector construction audit.
