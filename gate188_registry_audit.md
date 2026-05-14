# Gate 188 Registry Audit — Branchwise quadratic idempotent / scalar-projector construction audit

## Package

`pkg/bridge/branchprojector`

## Theorem

`BRIDGE-BRANCHWISE-QUADRATIC-IDEMPOTENT-SCALAR-PROJECTOR-AUDIT`

## Status

`BRIDGE_REQUIRED`

## Purpose

Gate 187 derived the exact threefold resolvent-vacuum orbit

```text
R_pair = Q[z]/(r3)
```

but did not construct scalar projectors. Gate 188 tests whether a spontaneous resolvent branch can produce exact `2+2` scalar projectors on the quartic companion module without diagonalizing the four individual roots.

## Key correction

The naive statement

```text
Q(z0) alone writes the two ordered quadratic factors
```

is too strong.

The exact result is:

```text
Q[z]/(r3) selects the unordered 2+2 partition.
K_pair = Q[z, eta]/(r3(z), eta^2 - (z^2 - 271/810)) labels the two quadratic pair factors.
```

The involution

```text
eta -> -eta
```

exchanges the two quadratic factors and exchanges the two projectors. No individual quartic root or linear factor is adjoined.

## Exact inputs

Quartic contact polynomial:

```text
q4(x) = 3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271
```

Monic form:

```text
x^4 - (71/30)x^3 + (1071/540)x^2 - (149/216)x + 271/3240
```

Resolvent cubic:

```text
r3(z) = z^3 - (119/60)z^2 + (8411/6480)z - 1637467/5832000
```

Quadratic factor-label adjunction:

```text
eta^2 = z^2 - 4d = z^2 - 271/810
```

## Constructed quadratic factors

Over `K_pair`, Gate 188 constructs two monic quadratics:

```text
q_A(x) = x^2 + m x + n
q_B(x) = x^2 + p x + q
```

with exact relations:

```text
n = (z - eta)/2
q = (z + eta)/2
m = (c - a n) / eta
p = a - m
```

where

```text
a = -71/30
b = 1071/540
c = -149/216
d = 271/3240
```

The engine verifies exactly:

```text
q_A(x) q_B(x) = q4(x)
```

## Bézout identity

Because `q_A` and `q_B` are coprime, the engine computes exact coefficients by extended Euclidean algorithm:

```text
A(x) q_A(x) + B(x) q_B(x) = 1
```

This uses exact rational arithmetic in `K_pair[x]`, not numerical root approximation.

## Projectors

The branchwise scalar projectors are constructed in the quotient algebra `K_pair[x]/(q4)`:

```text
P_A = B(x) q_B(x) mod q4
P_B = A(x) q_A(x) mod q4
```

The engine verifies exactly:

```text
P_A^2 = P_A
P_B^2 = P_B
P_A P_B = 0
P_A + P_B = 1
Tr(P_A) = 2
Tr(P_B) = 2
```

Therefore Gate 188 constructs a conditional `2+2` scalar-projector pair on the quartic module.

## What was not derived

```text
canonical unique branch: not derived
canonical scalar projector independent of branch data: not derived
individual quartic root projectors: not derived
linear root factors: not derived
physical H_Phi scalar bundle map: not derived
Chern-Weil carrier: not derived
heat-kernel matching: not derived
threshold beta rows: not derived
absolute coupling promotion: not derived
physical constants: not derived
```

## Firewall

```text
observed physical input: no
numeric root approximation: no
individual root diagonalization: no
arbitrary pairing choice: no
resolvent-vacuum inherited: yes
quadratic adjunction recorded: yes
branchwise quadratic factors derived: yes
branchwise projector pair derived: yes
conditional scalar projector derived: yes
physical scalar bundle derived: no
strict nullity: 3 -> 3
conditional nullity: 2 -> 1
```

## Validation

Focused tests:

```bash
go test -p=1 ./pkg/bridge/branchprojector ./pkg/bridge/resolventvacuum ./pkg/bridge/scalarcontactselector ./pkg/bridge/quarticscalaroperator ./pkg/bridge/cliffordcontactcommutant -count=1 -timeout=180s
```

Passed.

App/cmd compile smoke:

```bash
go test -p=1 ./internal/app ./cmd/asha -run '^$' -count=1 -timeout=300s
```

Passed.

Full theorem ladder smoke:

```bash
timeout 90s go run ./cmd/asha
```

Completed and printed Gate 188 successfully.

No full historical `go test ./...` suite was run.

## Next gate

Gate 189 — scalar-bundle map / `H_Phi` projector identification audit.

The next lawful task is to determine whether the branchwise projector pair can be mapped to the existing Gate-37 Higgs carrier without importing observed data, choosing an arbitrary branch convention, or prematurely reopening Chern-Weil / heat-kernel / threshold / coupling promotion.
