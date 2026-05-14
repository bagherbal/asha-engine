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

## Gate 229 method note — geometric constants require finite action maps

- When a fitted or near-fitted coefficient becomes recognizable as a geometric constant, decompose it explicitly into prior sealed data and standard mathematical factors before treating it as meaningful.
- Keep three ledgers separate: exact mathematical identity, finite-engine derivation, and phenomenological resonance. `4/π = S_top/(π Vol(S^3))` is exact; a contact-vacuum Hopf-action map is not yet derived.
- For exponential hierarchies, always compute sensitivity. A small coefficient residual can create a large scale displacement, and the audit must quantify this before discussing resonance.
- Be precise about relative versus absolute sensitivity. A derivative of about `53` per unit `B_gap` means a 1% relative `B_gap` shift is about `0.054` decades, while a 10% shift is about `0.54` decades.
- Compare residuals against already sealed uncertainty envelopes, but do not convert “plausibly covered” into “derived.” Matching corrections and higher-loop effects remain separate firewalled mechanisms.
- Do not grant an intermediate-breaking seal from a near-resonance alone. Require a finite order parameter, action normalization, and breaking potential.


- For resonance-to-mechanism gates, split the audit into three layers: inherited numerical/geometric resonance, required dynamic machinery, and seal status. Do not grant a seal merely because the resonance is tight; require an explicit operator, action, field equation, or order parameter.

## Gate 231 method note — check units before celebrating seesaw resonances

- When testing a seesaw scale, compute units explicitly: `v²/M_R` is in GeV and must be multiplied by `1e9` to obtain eV.
- Do not assume an intermediate scale near `10¹² GeV` gives `0.05 eV` neutrinos for order-one Yukawa. With `v≈246 GeV`, it gives tens to hundreds of eV unless the Dirac Yukawa is small.
- Keep three ledgers separate: Majorana scale, Dirac Yukawa amplitude, and full flavor/mixing matrix. A good scale relation does not derive PMNS data.
- Activating a seal can permit phenomenological testing, but it does not retroactively derive the instanton, order parameter, or field content that previous gates obstructed.
- If a target mass can be recovered only by choosing a Yukawa value, report the required Yukawa as a sealed conditional parameter, not as a finite prediction.
- Use neutrino mass observations as comparison bounds, not as inputs to tune the finite core.

- For flavor gates, separate three layers explicitly: inherited scale, texture proxy, and mixing matrix. A ratio-level texture resonance is not a PMNS derivation and must remain behind a seal unless the finite algebra supplies the full matrix structure.

## Gate 233 method note — initialize matrix arenas before claiming spectral actions

- When returning from phenomenology to the finite core, separate matrix-family existence from physical-operator derivation.
- A balanced Fock grading can define a legal odd self-adjoint `D_F` ansatz, but physical chirality requires a separate theorem.
- For finite Dirac work, always audit three ledgers independently: carrier Hilbert space, real/grading structures, and canonical block selection.
- A scalar spectral gap may be inserted into a diagnostic matrix only as an ansatz unless a finite theorem maps it to a bilinear amplitude.
- Computing `Tr(D²)` and `Tr(D⁴)` is not enough. Spectral-action physics also requires gauge projection, heat-kernel/cutoff data, and a subtraction scheme.
- Never insert `v`, `M_B`, `M_*`, or observed fermion masses into a finite `D_F` initialization gate.
- If the legal matrix family exists but the canonical selector is missing, record conditional support for the search space and a failed route for the physical operator.

## Gate 234 method note — real-structure candidates are sieves, not spectral triples

- A candidate `J` can be valuable even if it is not yet physical. Record which matrix parameters it removes, but do not call it charge conjugation without antiunitary structure and particle/antiparticle semantics.
- Compute KO signs as preflight data only unless the project has fixed a KO convention, physical chirality, and the `JD=±DJ` condition for the actual selected `D_F`.
- When applying order-one logic, first ask whether the finite algebra representation is faithful and non-vacuous. A diagonal bookkeeping algebra may be useful diagnostically but must not be mistaken for the spectral triple algebra.
- Distinguish three reductions: grading oddness, `J`-reality, and order-one calculus. Passing the first two does not imply the third.
- Do not force a B-gap Majorana term into the neutral state unless a doubled Hilbert carrier, right-handed-neutrino identification, and Majorana bilinear theorem exist.
- If a gate halves the search space but still leaves many parameters, report it as conditional finite support plus a failed route for physical operator selection.

## Gate 235 method note — derive doubling by complexification, not by adding states

- When a gate requires particle/antiparticle doubling, first ask whether the existing real carrier already contains the doubled real dimension after complexification. Do not append states externally if `S_C = S ⊗_R C` supplies them.
- Treat complex conjugation as an anti-linear candidate `J`, not automatically as physical charge conjugation. Physical charge conjugation still needs representation data and an opposite-algebra action.
- Keep three notions separate: real/imaginary bookkeeping halves, particle/conjugate representations, and physical particle/antiparticle states.
- Never import `C ⊕ H ⊕ M₃(C)` merely because the calculation resembles standard NCG. Search for the native associative algebra generated by the project’s already-derived gauge/contact data.
- A doubled carrier can permit Majorana bilinears kinematically, but it does not identify a right-handed neutrino slot or authorize a B-gap Majorana mass without the algebra representation and order-one calculus.
- Record progress as “carrier available” and obstruction as “algebra missing” when complexification works but the finite algebra is not yet derived.

## Gate 236 method note — derive algebra by commutants, not by recognition

- When searching for the finite algebra, begin with native projections and commutants. Do not import `C ⊕ H ⊕ M₃(C)` because it is familiar from Connes’ model.
- Separate the generator carrier from the full spinor carrier. A `1⊕3` split on the four Fock generators can support `C⊕M₃(C)` as a mode-level preflight without proving a full particle representation on `S_C`.
- Treat `su(2)` as a Lie algebra until an explicit associative action is derived. A quaternionic `H` summand requires a left module or equivalent closure, not just three generators with `su(2)` brackets.
- Keep `u(1) → C`, `spatial 3 → M₃(C)`, and `su(2) → H` as three independent ledgers. Passing two does not derive the third.
- The order-one condition remains blocked until the faithful algebra representation and opposite action are available. Do not use order-one language to force mass-matrix zeros before that.
- Report partial algebraic shapes as preflights. Only claim the Standard Model finite algebra when the exact associative algebra, representation, opposite action, and order-one readiness are all derived.

## Gate 237 method note — local quaternionic support is not global H derivation

- When deriving the weak algebra, distinguish an `su(2)` Lie action, an exterior spinor lift, a local pseudo-real doublet, and a global associative `H` summand. These are different achievements.
- Test all candidate two-mode planes rather than selecting the most Standard-Model-looking one. A plane selector must be derived by finite geometry, not chosen by recognition.
- Dimension matches such as an eight-complex-dimensional doublet sector are strong structural diagnostics, but they do not attach hypercharge, color, or physical chirality by themselves.
- A pseudo-real fundamental doublet supports a quaternionic module locally. Claim only local `H` support unless the action is globally selected and represented faithfully on the full carrier.
- Never import Pauli matrices, Connes’ algebra, or the Standard Model weak-doublet assignment as the proof. The gate may use representation identities only as preflight tests.
- Keep the order-one calculus blocked until a faithful finite algebra representation, opposite action, and canonical weak-plane selector are all available.

## Gate 238 method note — chirality must be tested, not recognized

- When a candidate grading has the right dimension split, compute its action on the actual candidate representation sectors. An `8⊕8` split alone does not make it Standard Model chirality.
- For weak-plane selection, audit all candidate planes and compare parity distributions. Do not pick the plane that looks physically familiar.
- Distinguish three layers: finite grading, physical chirality, and left-handed weak action. Passing the first does not imply the second or third.
- If a Lie action preserves a grading, it is not automatically chiral with respect to that grading. A chiral weak action must be isolated to the appropriate sector by a derived projection.
- Temporal/spatial class distinctions are meaningful, but class-level reduction is not unique plane selection.
- Preserve local quaternionic support from prior gates while logging failed global `H` derivation when the selector is missing.
