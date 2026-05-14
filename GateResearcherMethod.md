# GateResearcherMethod

- Start from the latest registry/audit files and identify the exact gate boundary: proven facts, sealed no-gos, conditional bridges, and quarantined empirical inputs.
- Read only the minimum source chain needed for the next gate first: registry entry, theorem package, tests, app wiring, and the immediately preceding gate audit.
- State the mathematical object before coding: inputs, invariants, unknowns, equations, admissible outputs, and firewall conditions.
- Implement the gate as a small isolated package when possible; keep phenomenology, finite algebra, and interpretation in separate structs/functions.
- Prefer exact/rational or auditable floating calculations; log residuals, tolerances, degeneracies, and underdetermined dimensions explicitly.
- Treat failure as data: if the system is underdetermined, overconstrained, irrational, noncanonical, or representation-free, record a no-go instead of forcing a fit.
- Reuse previous-gate audit snapshots for lightweight checks when importing a deep historical package would only retrieve already-audited constants.
- Avoid temporary `go run` probes that import broad theorem chains; use source/audit inspection, small package-local tests, or standalone arithmetic instead.
- Add focused tests for the new theorem package first; use compile-only checks for app/CLI wiring before broader runs.
- Avoid slow first-pass commands such as full `go test ./...` or full CLI execution until the focused gate path is stable.
- Use faster validation loops: `go test -p=1 ./pkg/path -count=1`, then selected dependent packages, then `go test -p=1 ./internal/app -run '^$'` and `go test -p=1 ./cmd/asha -run '^$'`.
- If a command times out, narrow scope rather than retrying blindly: identify package boundaries, warm the build cache, reduce verbose output, or replace deep imports with audited snapshots when scientifically legitimate.
- Before packaging, clean generated binaries and transient artifacts with a root-level artifact check, especially `*.test`, logs, temporary folders, and accidental build outputs.
- End every gate with a registry audit that separates theorem status into permanent, sealed/no-go, conditional, empirical/quarantined, and next-gate obligations.
- For inverse/phenomenology gates, classify source mechanisms before fitting numbers: exact row lattice, finite carrier permission, regulator/measure route, then firewall outcome.
- For representation-lattice gates, separate three ledgers: exact row grammar, lattice membership, and physical activation. Passing the first two must never imply the third.
- When a mathematical representation class is infinite, declare the finite alphabet explicitly and treat it as an audited grammar, not a universal enumeration.
- For carrier/threshold gates, require all three activation pillars before any beta-row promotion: charge semantics, spin-statistics semantics, and mass-activation/decoupling semantics.
- Never infer a representation from multiplicity alone; a seven-mode carrier is not automatically an adjoint, color sector, ghost ledger, or threshold multiplet.
- When a gate introduces a seal, split the result into two ledgers: native-search obstruction and sealed conditional consequence. Never let the sealed consequence rewrite the native theorem status.
- For sealed phenomenology gates, require anomaly/consistency checks before emitting numbers, and label every number with the exact seal and inherited external assumptions.
- Distinguish an activated conditional carrier from a finite-derived particle: the former can be used for stress tests; the latter requires charge, spin-statistics, mass, and decoupling theorems.
- For stress-test gates, split the output into branch verdicts. A collider branch can pass while a UV-completion branch fails; do not collapse mixed evidence into a false global success.
- External experimental limits belong in a quarantined ledger. Use them for stress-testing sealed phenomenology, never for finite-core derivation.
- When a boundary scale is near proton-decay danger, first audit mediator/operator support. Do not compute a lifetime unless the engine derives or explicitly seals the relevant `B/L`-violating operator basis.
- For high-scale beta completions, test one-loop pole/asymptotic-safety behavior before declaring numerical predictions viable. A sub-Planck pole is a failed route unless a new UV completion or matching theorem is supplied.
- For proton-decay gates, do not use `B-L` as a blanket firewall: the standard `QQQL` and `UUD E` classes preserve `B-L`, so the real question is operator construction and mediator activation.
- Distinguish current-connection stability from absolute baryon conservation. If quark-lepton current slots exist but lack an action/propagator/coefficient, record an operator-construction obstruction, not an all-future conservation theorem.
- For dormant-current gates, separate inventory from dynamics. A current slot becomes dangerous only after curvature/action/local-field/propagator/mass/coefficient semantics are derived or explicitly sealed.
- When sealing a dormant threat surface, state what the seal forbids operationally and what future theorem could lift it. The seal must not rewrite the native failed route.
- For rational RG-lattice gates, check exact determinant/field-structure obstructions before brute-force search; use bounded enumeration only as supporting evidence, never as the sole proof of an infinite semigroup no-go.
- For multi-threshold RG gates, separate algebraic solvability from physical viability. Once the number of independent threshold rows equals the coupling-space dimension, closure is a linear solve; the theorem lives in ordering, positivity, pole, anomaly, and seal filters.
- Use the correct coupling convention before solving: `α⁻¹` formulas carry `2π`, while `u=1/g²` formulas carry `8π²`. Do not mix them across gates.
- When a prompt gives an explicit logarithmic cutoff, treat that as the operative filter even if a named physical scale would imply a slightly different logarithm.
- For viable multi-threshold witnesses, record reverse orderings carefully: ordered threshold labels may be mathematically distinct even when the unordered physical row set is the same.
- Record asymptotic freedom and Landau safety separately. Positive high-scale beta coefficients can still pass a bounded no-pole audit over a finite interval, but they are not asymptotically free.
- For degenerate viable witness sets, separate ordered solver labels from unordered physical spectra before doing minimality or uniqueness analysis.
- Never treat a ranking function as a finite theorem. Ranking can prioritize witnesses for inspection, but uniqueness requires a derived selector, symmetry, spectrum rule, or explicit seal.
- For parentage audits, external branching patterns are only hints unless the engine derives the parent gauge connection, branching map, missing partners, and splitting rule.
- For threshold-splitting arguments, numerical closeness of logs is not a mass theorem. Require an activation/splitting mechanism or log a `ThresholdSpectrumSeal` obligation.
- For degenerate viable spectra, introduce an explicit spectrum seal before inspecting a ranked witness. The seal must quarantine the selection and prevent ranked-best from becoming finite-derived uniqueness.
- Matching corrections require more than finite traces: demand a subtraction scheme, counterterm functional, heat-kernel/spectral-action map, and threshold matching convention before deriving `δ_i^match`.
- Two-loop coefficients may be computed as standard-QFT preflight data, but their provenance must stay separate from finite-core theorems unless the engine derives the action, field normalization, and scheme.
- For two-loop preflights, report correction-size diagnostics before claiming stability. If the two-loop/one-loop derivative ratio is not small, keep one-loop scales as reference values only and require full integration plus matching envelopes.
- For two-loop integration gates, solve in the same coupling coordinate used by the previous gate. In the ASHA RG branch this means `u=1/g²`, so the two-loop equation is `du_i/dlnμ = -b_i/(8π²) - Σ_j B_ij/u_j /(128π⁴)`.
- Treat two-loop numerical solves as phenomenological unless the engine has derived the action, normalization, and subtraction scheme. Corrected scales are not automatically finite-core predictions.
- When matching corrections are not derived, use an explicit uncertainty envelope rather than hiding scheme dependence. Label the envelope as a proxy and record its size as theory uncertainty.
- Allow higher-order integration to change threshold ordering if the spectrum is sealed and the scales are continuous fit parameters. Record the ordering change, but do not reinterpret it as a finite theorem.

## Gate 215 method refinement — global sealed scan before interpretation

- When a previous gate hints that a multi-scale result is nearly degenerate, do not collapse the scales by assumption. First run a forced-degenerate scan and measure the residual.
- Compare residuals to an explicit uncertainty envelope already declared by earlier gates; do not silently widen the envelope to save a preferred spectrum.
- Scan all previously viable classes, not only the favorite witness, before claiming a spectrum is structurally preferred.
- Rank by required correction size, but keep ranking separate from derivation. A small required correction is a target for the next theorem, not proof that the correction exists.
- Keep the validation focused: new package tests, direct predecessor tests, `go list` wiring checks, and no timeout-prone full registry execution unless the gate specifically requires it.

## Gate 216 method refinement — residual targets are not coefficients

- When a previous gate produces a small required residual, treat it as a target vector first, not as a discovered constant.
- Compare sign pattern, relative magnitude, and absolute magnitude separately. A sign-only resonance is useful, but it is not a matching theorem.
- Reject loop-factor near-misses unless the normalization is canonical and branch-free. Do not insert fitted coefficients to turn spectral scalars into threshold corrections.
- Heat-kernel language requires concrete structure: finite Dirac operator, spectral triple, gauge-curvature projection, cutoff moments, and subtraction scheme.
- Existing finite traces may be action-level diagnostics while still lacking permission to become `δ_i^match` rows.
- Preserve the difference between a future matching target and a derived matching correction in every theorem, audit, and README summary.

## Gate 217 method refinement — spectral-action machinery must be complete before matching constants exist

- Treat a required matching vector as a target, not as evidence that a finite coefficient has been derived.
- Split spectral-action audits into separate gates: finite Hilbert/`D_F`, gauge projection, and cutoff/subtraction. A failure in any one piece blocks `δ_i^match`.
- Reject vacuous or hand-built `D_F` candidates: zero operators, identity mass matrices, and phenomenological mass insertions are diagnostics only.
- Keep representation-trace facts separate from heat-kernel matching facts. A beta row or Dynkin index is not a finite threshold constant.
- Never import `MSbar`, dimensional regularization, or a cutoff function as a finite theorem. They must be explicit seals or external phenomenological inputs.
- For speed, test the new bridge package and its direct predecessor only; use `go list` for registry wiring and avoid the known full-registry compile path unless intentionally validating the whole historical ladder.

## Gate 218 method refinement — empirical SM running must stay sealed

- When a missing matching correction has been proven non-derived, introduce an explicit matching seal before using it in downstream numerical audits.
- Treat top mass, Higgs mass, `y_t`, and `λ` as empirical ledger inputs unless the finite engine has derived their texture and amplitudes.
- Add SM Yukawa/scalar running as a controlled phenomenology upgrade, not as a finite-core theorem.
- Separate three claims: the forced single-scale fit converges, the required matching residual is inside the envelope, and the residual is derived. Gate 218 can prove the first two only.
- Keep heavy-sector Yukawa couplings absent unless a new seal or finite local-field theorem introduces them.
- When adding full-SM effects, compare against the previous no-Yukawa target and report the shift instead of overwriting the earlier result.

## Gate 219 method refinement — input uncertainty is not tuning freedom

- When a numerical gate depends on empirical low-energy inputs, treat every uncertainty as a bounded audit variable, not as a parameter to tune the target residual to zero.
- Add missing small SM couplings only when they have a clear role in the running equations; keep lighter negligible terms explicitly listed as omissions rather than silently forgotten.
- Report central values, induced scale ranges, residual-envelope ranges, and the dominant sensitivity driver separately.
- If a parameter changes a running variable that does not feed back at the current loop order, say so directly instead of inventing an effect.
- Preserve the difference between experimental uncertainty, theoretical matching uncertainty, and finite-core derivation. They are three separate ledgers.
- For speed, avoid full-grid covariance scans until one-at-a-time `±1σ` scans establish whether the envelope is threatened.

## Gate 220 workflow refinement

- Separate indirect-observable safety from cosmological safety: decoupling can make EW/Higgs effects tiny while stable relics remain dangerous.
- Do not invent decay operators, heavy Yukawa couplings, or mass splittings to make a PeV spectrum phenomenologically safe.
- Use dimensionless suppression proxies first (`v/M`, `v²/M²`) before attempting model-specific observable calculations.
- If decay/splitting data is absent, log a warning and push it into a dedicated future gate rather than computing relic abundance from missing dynamics.

## Gate 221 workflow refinement — relic safety requires dynamics, not decoupling

- Treat precision safety and cosmological safety as separate ledgers. A PeV state can decouple from EW/Higgs observables while still being fatal as a stable relic.
- Search decay portals by required semantics: gauge invariance, Lorentz/local-field map, coupling coefficient, suppression scale, seal compatibility, and width calculation. Missing any one piece blocks a lifetime theorem.
- Do not use toy dimension-5 or dimension-6 widths to rescue a spectrum unless the operator coefficient and scale are derived or explicitly sealed.
- Use BBN only as a safety threshold when dynamics are absent. If no width exists, classify the lifetime as unbounded for safety purposes and log a failed route.
- Keep leptoquark-mediated decays blocked while the `LeptoquarkDynamicsSeal` is active; dormant current slots are not propagators.
- A `RelicDecaySeal` can be required without being granted. Grant it only after a concrete decay/splitting sector supplies BBN-safe lifetimes.

## Gate 222 workflow refinement — relic seals require every carrier to decay

- Correct representation identities before building portals. A field with `(8,2,Y=1/2)` is not SM `Q=(3,2,Y=1/6)`, so simple mass mixing is illegal.
- Audit each heavy carrier separately. A BBN-safe decay for the neutral/colorless carrier does not rescue a stable colored carrier.
- A partial portal can be useful, but full relic safety requires neutral, charged, and colored components to have legal decay/cascade paths.
- Keep EFT portals quarantined: gauge-invariant operators, Wilson coefficients, and suppression scales are phenomenological unless derived from the finite core.
- Use BBN bounds to constrain sealed couplings/scales, not to invent operators. If the operator is not certified, its width formula is diagnostic only.
- Keep leptoquark-mediated decays blocked while `LeptoquarkDynamicsSeal` is active. Do not use dormant current slots to rescue relic safety.

## Gate 223 workflow refinement — relic rescue must be a tensor-product theorem

- Do not trust intuitive identity claims for representations. Recompute color, weak, hypercharge, Lorentz parity, and dimension before accepting a decay portal.
- For a heavy fermion `Ψ`, search for a pure-SM spinor composite `O_SM` such that `bar(Ψ)O_SM` is a singlet. Keep the heavy-field dimension in the total operator-dimension bound.
- Keep baryon/proton firewalls active while searching relic portals. A gauge-invariant operator can still be rejected if it would reopen a sealed baryon-violation route.
- Distinguish operator existence from coefficient existence. Finding a dimension-six portal permits a `RelicDecaySeal`, but the Wilson coefficient and suppression scale remain phenomenological unless derived.
- If a portal is found, compute BBN safety as a bound on `Λ` or `c`, not as a relic abundance. Full abundance and flavor constraints require a later gate.
- If no portal exists within the bounded field alphabet and dimension limit, falsify the sealed spectrum before moving to lower-ranked RG candidates.

## Gate 224 method note — flavor and relic closure

- After a decay seal is granted, immediately audit flavor tensors; gauge invariance does not imply flavor safety.
- Distinguish three levels: operator existence, BBN lifetime safety, and flavor/rare-decay safety.
- Introduce a flavor seal only when exact rare-decay machinery is missing; never claim FCNC bounds are passed without Wilson matrices, basis choice, and matrix elements.
- If a sector is forced to decay before BBN, log its present-day dark-matter contribution as zero for that sector and defer dark matter to a separate inventory audit.

## Gate 225 method note — dark matter requires symmetry, scale, and history

- After a sector is proven to decay before BBN, treat its present-day dark-matter abundance as zero and move the dark-matter question to a separate finite inventory audit.
- Do not identify a dimensionless spectral scalar with an axion mass, decay constant, or relic abundance without a derived dimensionful scale.
- For an ALP claim, require a shift symmetry, periodic coordinate, anomaly/Pontryagin coupling, and production mechanism before computing Ωh².
- For a contact-mode dark-sector claim, require gauge-singlet semantics, a stability symmetry, a dark action, and mass/production laws; non-promotion to SM carriers is not enough.
- Near loop-factor numerical values should be logged as diagnostics only. Never insert a coefficient to turn a finite scalar into a matching, mass, or relic-density fit.
- Use the observed dark-matter abundance as a target for future validation, not as an input for deriving finite scales.

## Gate 226 method note — sealed cosmology is not native cosmology

- After a finite-anchor dark-matter route fails, a phenomenological seal may be used to ask scale questions, but only after explicitly preserving the failed native theorem.
- For ALP audits, keep four ledgers separate: shift symmetry, anomaly coupling, decay constant, and relic production. Sealing one does not derive the others.
- Use observed `Ω_DM h²` only as a target for a conditional parameter extraction, not as input that rewrites the finite core.
- Compare any extracted dark scale against existing hierarchy scales in log space; report non-resonance as clearly as resonance.
- Diagnostic variants such as `θ_i = B_gap` are useful only if marked noncanonical unless a prior gate derived that identification.
- Do not claim dark matter is solved when the result depends on an un-derived axion seal. The native task remains deriving the symmetry, `f_a`, and anomaly map.

## Gate 227 method note — resonance is not derivation

- When two sealed phenomenological scales appear near each other, first test the null hypothesis quantitatively in log space before inventing a new mechanism.
- Use inherited sealed values only when testing hierarchy relations; do not refit `M_B`, `M_*`, `f_a`, or `Λ_EFT` to improve a resonance.
- A geometric mean relation can be a structural clue, but it is not a finite theorem until an order parameter, breaking potential, or exact operator derives the intermediate scale.
- Keep “bracketing” and “matching” distinct: `Λ_EFT < M_int < f_a` is stronger than both scales merely being within one decade, but still conditional.
- Do not reopen sealed leptoquark or Pati-Salam dynamics just because the intermediate scale is suggestive. Existing seals remain active until a gate derives the missing curvature/action/propagator data.
- Treat resonance gates as target-selection gates: they identify where to search next, not what has already been proven.

## Gate 228 method note — test lethal baryon channels before building intermediate symmetry

- When an intermediate scale is identified, run proton-decay and baryon-violation kill-switches before constructing elegant common-origin models.
- Temporarily unseal a dangerous mediator only for a bounded falsification estimate; do not treat that as restored dynamics or a finite theorem.
- Keep `B-L` separate from proton safety. Dimension-six proton-decay operators can preserve `B-L`, so the relevant question is mediator/action/operator existence and mass scale.
- For non-perturbative hierarchy searches, distinguish three levels: functional shape, target coefficient, and finite-derived coefficient. Only the third can grant a native mechanism.
- Near-resonances such as `c≈4/π` should be logged as diagnostics unless the finite algebra derives the coefficient through an action, trace, volume, or instanton normalization.
- A hidden-sector route can be favored by falsifying baryon-unsafe routes, but that still does not derive the hidden order parameter or breaking potential.
