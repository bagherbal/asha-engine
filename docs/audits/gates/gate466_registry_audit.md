# Gate 466 Registry Audit — Quark-Sector Observed Comparator Adapter / CKM Data Firewall

## Verdict

`FAILED_ROUTE_CKM_GEOMETRIC_ALIGNMENT_NOT_COMPUTABLE_FROM_MASS_SPECTRA_ONLY`

Gate 466 opened the empirical airlock and imported observed quark-mass plus Cabibbo comparator rows into the quarantined ledger. The calculation does **not** produce a unique `d_ud`: the observed mass rows do not supply a common-scale trace-zero sector model, the independent `I_K` comparator, or the `{sigma_CP,n_C3}` branch tags required by Gates 454–459.

## Inheritance

executed=true K=true triangle=true spectrum_rank_one=true inverse=true branch_tags=true ckm_null_socket=true airlock=true rejects_native=true native_clean=true verdict=CONDITIONAL_SUPPORT_GATE465_EMPIRICAL_AIRLOCK_INHERITED

## Airlock import

executed=true empirical_import=true accepted=7 rejected=0 quark_mass_rows=6 ckm_rows=1 quarantined=true no_native_write=true native_promotion_probe=true native_registry_probe=true theorem_probe=true verdict=CONDITIONAL_SUPPORT_OBSERVED_QUARK_CKM_ROWS_IMPORTED_TO_QUARANTINED_LEDGER reason=empirical_import=true admits seven observed comparator rows into the quarantined quark-sector ledger, while native-promotion probes fail closed

| Row | Accepted | Verdict | Metadata |
|---|---:|---|---|
| `m_u` | true | `CONDITIONAL_SUPPORT_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_COMPARATOR_LEDGER` | m_u sector=u obs=running quark mass value=2.16 MeV source=PDG Review of Particle Physics observed comparator row version=PDG-style Gate466 bridge constants; values are empirical comparator inputs, not theorem premises scale=2 GeV scheme=MS-bar uncertainty=+0.49/-0.26 MeV bridge_only=true |
| `m_c` | true | `CONDITIONAL_SUPPORT_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_COMPARATOR_LEDGER` | m_c sector=u obs=running quark mass value=1.27 GeV source=PDG Review of Particle Physics observed comparator row version=PDG-style Gate466 bridge constants; values are empirical comparator inputs, not theorem premises scale=mu=m_c scheme=MS-bar uncertainty=±0.02 GeV bridge_only=true |
| `m_t` | true | `CONDITIONAL_SUPPORT_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_COMPARATOR_LEDGER` | m_t sector=u obs=running top quark mass value=162.5 GeV source=PDG Review of Particle Physics observed comparator row version=PDG-style Gate466 bridge constants; values are empirical comparator inputs, not theorem premises scale=mu=m_t scheme=MS-bar uncertainty=declared external uncertainty bridge_only=true |
| `m_d` | true | `CONDITIONAL_SUPPORT_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_COMPARATOR_LEDGER` | m_d sector=d obs=running quark mass value=4.67 MeV source=PDG Review of Particle Physics observed comparator row version=PDG-style Gate466 bridge constants; values are empirical comparator inputs, not theorem premises scale=2 GeV scheme=MS-bar uncertainty=+0.48/-0.17 MeV bridge_only=true |
| `m_s` | true | `CONDITIONAL_SUPPORT_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_COMPARATOR_LEDGER` | m_s sector=d obs=running quark mass value=93.4 MeV source=PDG Review of Particle Physics observed comparator row version=PDG-style Gate466 bridge constants; values are empirical comparator inputs, not theorem premises scale=2 GeV scheme=MS-bar uncertainty=+8.6/-3.4 MeV bridge_only=true |
| `m_b` | true | `CONDITIONAL_SUPPORT_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_COMPARATOR_LEDGER` | m_b sector=d obs=running quark mass value=4.18 GeV source=PDG Review of Particle Physics observed comparator row version=PDG-style Gate466 bridge constants; values are empirical comparator inputs, not theorem premises scale=mu=m_b scheme=MS-bar uncertainty=+0.03/-0.02 GeV bridge_only=true |
| `\|V_us\|` | true | `CONDITIONAL_SUPPORT_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_COMPARATOR_LEDGER` | \|V_us\| sector=u-d obs=Cabibbo/CKM 12 comparator value=0.225 dimensionless source=PDG Review of Particle Physics observed comparator row version=PDG-style Gate466 bridge constants; values are empirical comparator inputs, not theorem premises scale=weak charged-current convention scheme=CKM standard parameterization uncertainty=declared external uncertainty bridge_only=true |

## Coordinate-map attempt

executed=true rows=7 up_masses=3 down_masses=3 cabibbo=true |Vus|=0.225 common_scale_required=true common_scale=false trace_zero_required=true trace_zero_supplied=false ray_dof=2 spectrum_rank=1 min_comparators=2 I_K=false branch_tags=false alpha_u=false phi_u=false alpha_d=false phi_d=false d_ud=undefined |d_ud-Vus|=undefined alignment=false verdict=FAILED_ROUTE_CKM_GEOMETRIC_ALIGNMENT_NOT_COMPUTABLE_FROM_MASS_SPECTRA_ONLY reason=PDG-style mass rows pass the airlock but do not provide a common-scale trace-zero sector spectrum, the independent I_K comparator, or the {sigma_CP,n_C3} branch tags required to define alpha_u, phi_u, alpha_d, phi_d; therefore d_ud is undefined and cannot be compared to |V_us|

The Gate 464 socket requires `alpha_u`, `phi_u`, `alpha_d`, and `phi_d`. Gate 454 already proved that spectrum-only information has rank one while the projective coefficient ray has two degrees of freedom. Gate 459 further requires branch metadata. Therefore the adapter refuses to fabricate coordinates from masses alone.

```text
d_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))
Gate466 result: d_ud = undefined
observed bridge target: |V_us| = 0.225
comparison: not computed
```

## Native firewall proof

executed=true observed_rows=7 quarantined=true empirical_in_native=false native_prediction=false native_law=false theorem_input=false quark_mass_native=false CKM_native=false CKM_constructed=false CKM_entry=false d_ud_native=false alignment_native=false K=true triangle=true Y_sealed=true coeffs_sealed=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_OBSERVED_AIRLOCK_ROWS reason=observed rows can enter the bridge ledger, but the attempted cylinder-coordinate map fails before d_ud exists; no quark mass, CKM value, coefficient ray, or alignment claim enters native law-space

No imported row writes to the native theorem registry. No quark mass, CKM value, coefficient ray, `d_ud`, or alignment claim is exported as `native_prediction` or `native_law`. The 13-moduli firewall and the 9 charged K/X/Y coefficient seals remain intact.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE465_EMPIRICAL_AIRLOCK_INHERITED`
- `CONDITIONAL_SUPPORT_OBSERVED_QUARK_CKM_ROWS_IMPORTED_TO_QUARANTINED_LEDGER`
- `CONDITIONAL_SUPPORT_OBSERVED_COMPARATOR_ADAPTER_ATTEMPTED`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_OBSERVED_AIRLOCK_ROWS`
- `FAILED_ROUTE_COMMON_SCALE_SCHEME_REQUIRED_FOR_SECTOR_COORDINATES`
- `FAILED_ROUTE_OBSERVED_MASS_SPECTRA_DO_NOT_DEFINE_ASHA_RAY`
- `FAILED_ROUTE_OBSERVED_MASS_IMPORT_MISSING_IK_COMPARATOR`
- `FAILED_ROUTE_OBSERVED_MASS_IMPORT_MISSING_BRANCH_TAGS`
- `FAILED_ROUTE_DUD_UNDEFINED_FOR_OBSERVED_MASS_ONLY_IMPORT`
- `FAILED_ROUTE_CKM_GEOMETRIC_ALIGNMENT_NOT_COMPUTABLE_FROM_MASS_SPECTRA_ONLY`
- `FAILED_ROUTE_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_EMPIRICAL_DATA_NATIVE_REGISTRY_WRITE_REJECTED`
- `FAILED_ROUTE_OBSERVED_DATA_AS_THEOREM_INPUT_REJECTED`

## Next gate

Gate 467 — Common-Scale Running Ledger / Coefficient-Ray Comparator Design: Gate466 proves PDG-style mass rows alone cannot define d_ud; a future bridge calculation needs common-scale running inputs plus an explicitly labelled second comparator I_K and branch tags. Primary task: define a bridge-only data product containing common-scale sector spectra, I_K comparators, {sigma_CP,n_C3} branch tags, and uncertainty propagation, without native promotion

## Truth statement

Gate 466 safely imports observed quark-mass and Cabibbo comparator rows, but it does not compute d_ud. Mass spectra alone do not define the ASHA coefficient rays: common-scale sector spectra, an independent I_K comparator, and branch tags are missing. Therefore no CKM geometric-alignment claim is mathematically licensed.
