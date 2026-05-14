# Gate 222 Registry Audit — EFT decay portal construction / RelicDecaySeal activation audit

## Status

```text
PARTIAL_EFT_PORTAL_SUPPORT_TRIPLET_ONLY
FAILED_ROUTE_COLORED_OCTET_DECAY_PORTAL
RELIC_DECAY_SEAL_NOT_GRANTED_FULL_SPECTRUM
```

Gate 222 inherits Gate 221's cosmological pathology: the sealed PeV spectrum is precision-safe but not relic-safe unless the heavy carriers can decay before BBN.

Inherited spectrum:

```text
Dirac (1,3,Y=1)
Dirac (8,2,Y=1/2)
M_B ≈ 2.56895727e6 GeV
```

BBN safety threshold used only as a filter:

```text
τ < 1 second
Γ_required > 6.582119569e-25 GeV
```

## Critical correction

The colored carrier is **not** identical to the SM quark doublet:

```text
sealed carrier:  (8,2,Y=1/2)
SM Q doublet:    (3,2,Y=1/6)
```

Therefore a simple mass-mixing portal `m Ψ_8 Q` is not gauge invariant. Gate 222 rejects this shortcut instead of granting the full relic seal from a false representation identity.

## Operator audit

| Candidate | Target | Result |
| --- | --- | --- |
| `y_T Ψ_3^a (L σ^a H†)` | `(1,3,Y=1)` | Gauge invariant sealed EFT portal; not finite-derived |
| `m_mix Ψ_8 Q` | `(8,2,Y=1/2)` | Rejected: color and hypercharge mismatch with SM `Q` |
| Dimension≤6 pure-SM octet portal | `(8,2,Y=1/2)` | No certified gauge/Lorentz/operator template in audited basis |
| Leptoquark-assisted decay | `(8,2,Y=1/2)` | Blocked by `LeptoquarkDynamicsSeal` |

## BBN parametric bounds

For the triplet portal, using

```text
Γ ≈ |y_T|² M_B / (8π)
```

BBN safety requires only

```text
|y_T| > 2.53760706e-15
```

This is tiny, so the triplet can be rescued by an explicitly sealed phenomenological Yukawa portal.

For diagnostic comparison only, a unit-Wilson dimension-six four-fermion-type decay width would obey

```text
Γ ≈ M_B^5 / (192π^3 Λ^4)
```

and BBN safety would allow roughly

```text
Λ < 1.29992096e13 GeV
```

but Gate 222 does **not** use this to rescue the octet, because the needed octet decay operator is not certified.

## RelicDecaySeal result

Gate 222 grants only a partial triplet sub-support:

```text
Triplet portal support: true
Colored octet portal support: false
Full RelicDecaySeal granted: false
```

The full PeV spectrum remains cosmologically unsafe until a future theorem or explicit phenomenological seal supplies a legal colored-octet decay channel.

## Firewalls preserved

Gate 222 does not claim:

```text
native finite decay operator
finite-derived Yukawa coupling
finite-derived EFT scale
simple octet-Q mixing
leptoquark-mediated decay while LeptoquarkDynamicsSeal is active
relic abundance computation
full cosmological safety
physical prediction
```

## Next obstruction

The next required step is a colored-octet decay portal alternatives audit. The engine must either:

1. find a legal pure-SM higher-dimensional operator for `(8,2,Y=1/2)`,
2. explicitly seal a new mediator sector without violating the proton/leptoquark firewalls, or
3. reject/replace the sealed PeV spectrum as cosmologically fatal.
