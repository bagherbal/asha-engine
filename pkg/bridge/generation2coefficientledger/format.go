package generation2coefficientledger

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("KForced=%t gen2Zero=%t XSupport=%t ampSealed=%t signSealed=%t phaseSealed=%t YQuarantined=%t nativeDim=%d KXYBefore=%d noEmpirical=%t verdict=%s", x.Gate444KGenForced, x.Gate444Generation2BareZero, x.Gate445XSupportForced, x.Gate445AmplitudeSealed, x.Gate446SignedCycleSealed, x.Gate446ComplexPhaseSealed, x.Gate446YGenQuarantined, x.NativeFlavorDim, x.KXYCoeffDimBefore, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatArena(x Arena) string {
	return fmt.Sprintf("expr=%q sectors=%s basis=%s KForced=%t XForced=%t YNative=%t Hermitian=%t gaugeBlind=%t traceNeutral=%t coeffs=%d verdict=%s", x.TextureExpression, join(x.ChargedSectorNames), join(x.Basis), x.KAxisGeometricallyForced, x.XSupportGeometricallyForced, x.YQuadratureNative, x.HermitianFamilySources, x.GaugeBlindFamilyFiber, x.TraceNeutralBasis, x.TotalSymbolicCoefficients, x.Verdict)
}

func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("%s formula=%q applied=%t passed=%t selectsValues=%t verdict=%s reason=%s", x.Name, x.Formula, x.Applied, x.Passed, x.SelectsCoefficientValues, x.Verdict, x.Reason)
}

func FormatFunctional(x FunctionalAudit) string {
	return fmt.Sprintf("%s functional=%q gauge=%t empiricalIndependent=%t uniqueRay=%t sectorWeights=%t massValues=%t mixingAngles=%t diagnostic=%.6g verdict=%s reason=%s", x.Name, x.Functional, x.GaugeCompatible, x.EmpiricalIndependent, x.SelectsUniqueCoefficientRay, x.SelectsSectorWeights, x.PredictsMassValues, x.PredictsMixingAngles, x.DiagnosticValue, x.Verdict, x.Reason)
}

func FormatCounterLedger(x CounterLedger) string {
	return fmt.Sprintf("%s up=%s down=%s lepton=%s Hermitian=%t trace=%t gauge=%t KMS=%t massLift=%t importsEmpirical=%t distinct=%t verdict=%s", x.Name, FormatTriple(x.UpCoefficients), FormatTriple(x.DownCoefficients), FormatTriple(x.LeptonCoefficients), x.Hermitian, x.TraceNeutral, x.GaugeCompatible, x.KMSCompatible, x.MassLiftCompatible, x.ImportsEmpiricalData, x.DistinctFromOtherLedgers, x.Verdict)
}

func FormatSieve(x CounterLedgerSieve) string {
	return fmt.Sprintf("ledgers=%d survivors=%d distinct=%d unique=%t universalRay=%t KFixed=%t XFixed=%t YFixed=%t verdict=%s reason=%s", len(x.Ledgers), x.SurvivingLedgers, x.DistinctSurvivors, x.UniqueCoefficientLedger, x.ForcesUniversalSectorRay, x.ForcesKCoefficientValues, x.ForcesXCoefficientValues, x.ForcesYCoefficientValues, x.Verdict, x.Reason)
}

func FormatCoefficientLedger(x CoefficientLedger) string {
	return fmt.Sprintf("sectors=%s basis=%s symbols=%s total=%d nativeValues=%d quarantined=%d KForced=%t XForced=%t YQuarantined=%t amplitudesSealed=%t masses=%t CKM=%t PMNS=%t verdict=%s", join(x.Sectors), join(x.Basis), join(x.SymbolNames), x.TotalSymbols, x.NativeCoefficientValues, x.QuarantinedCoefficientDim, x.KAxisForced, x.XSupportForced, x.YQuadratureQuarantined, x.AmplitudeValuesSealed, x.PhysicalMassesPredicted, x.CKMPredicted, x.PMNSPredicted, x.Verdict)
}

func FormatClosure(x Closure) string {
	return fmt.Sprintf("nativeDim=%d→%d KXY=%d→%d nativeReduction=%t coeffReduction=%t KPreserved=%t XPreserved=%t YNative=%t coeffAxiom=%t firewallClosed=%t verdict=%s reason=%s", x.NativeFlavorDimBefore, x.NativeFlavorDimAfter, x.KXYCoeffDimBefore, x.KXYCoeffDimAfter, x.NativeReductionBelow13, x.CoefficientReductionBelow9, x.KGenStructuralAxiomPreserved, x.XSupportStructuralPreserved, x.YGenPromotedNative, x.AnyCoefficientAxiomPromoted, x.AmplitudeFirewallClosed, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("muonImported=%t charmImported=%t yukawaImported=%t CKM=%t PMNS=%t poleFit=%t curveFit=%t KNative=%t XNative=%t YQuarantined=%t coefficientsQuarantined=%t verdict=%s", !x.NoObservedMuonMassImported, !x.NoObservedCharmMassImported, !x.NoObservedYukawaImported, !x.NoCKMImported, !x.NoPMNSImported, !x.NoPoleMassFit, !x.NoCurveFit, x.KGenNative, x.XSupportNative, x.YGenQuarantined, x.CoefficientsQuarantined, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func FormatTriple(x [3]string) string { return fmt.Sprintf("(%s,%s,%s)", x[0], x[1], x[2]) }

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 447 Registry Audit — Sector-Coefficient Source Ledger / Amplitude Firewall Closure\n\n")
	b.WriteString("## Scope\n\n")
	b.WriteString("Gate 447 audits whether the post-Gate-446 boundary stack selects the charged-sector amplitudes multiplying the structural family operators. It is intentionally not a fit: no muon/charm mass, Yukawa matrix, CKM angle, CKM phase, or PMNS datum is allowed into the sieve.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("The inherited structural state is:\n\n")
	b.WriteString("```text\n")
	b.WriteString("K_gen = diag(-1,0,1)                         // Gate 444 structural axis\n")
	b.WriteString("support(X_triangle) = complete 3-cycle        // Gate 445 unsigned mass-lift topology\n")
	b.WriteString("Phi_cycle, Y_gen, amplitudes = quarantined    // Gate 446 orientation firewall\n")
	b.WriteString("```\n\n")

	b.WriteString("## Coefficient arena\n\n")
	b.WriteString(FormatArena(a.Arena) + "\n\n")
	b.WriteString("The charged-sector symbolic texture ledger is:\n\n")
	b.WriteString("```text\n")
	b.WriteString("M_u = kappa_u K_gen + xi_u X_triangle + upsilon_u Y_phase\n")
	b.WriteString("M_d = kappa_d K_gen + xi_d X_triangle + upsilon_d Y_phase\n")
	b.WriteString("M_e = kappa_e K_gen + xi_e X_triangle + upsilon_e Y_phase\n")
	b.WriteString("dim C_KXY^charged = 3 sectors × 3 coefficients = 9\n")
	b.WriteString("```\n\n")

	b.WriteString("## Native boundary stack\n\n")
	b.WriteString("| Boundary | Formula | Applied | Passed | Selects coefficient values | Verdict | Reason |\n")
	b.WriteString("|---|---|---:|---:|---:|---|---|\n")
	for _, x := range a.Boundaries {
		b.WriteString(fmt.Sprintf("| %s | `%s` | %t | %t | %t | `%s` | %s |\n", x.Name, x.Formula, x.Applied, x.Passed, x.SelectsCoefficientValues, x.Verdict, x.Reason))
	}
	b.WriteString("\n")

	b.WriteString("## Functional selector sieve\n\n")
	b.WriteString("| Functional | Native/empirical status | Unique ray | Sector weights | Mass values | Mixing angles | Verdict | Reason |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---|---|\n")
	for _, x := range a.Functionals {
		status := "native-compatible"
		if !x.EmpiricalIndependent {
			status = "source-dependent"
		}
		b.WriteString(fmt.Sprintf("| `%s` | %s | %t | %t | %t | %t | `%s` | %s |\n", x.Functional, status, x.SelectsUniqueCoefficientRay, x.SelectsSectorWeights, x.PredictsMassValues, x.PredictsMixingAngles, x.Verdict, x.Reason))
	}
	b.WriteString("\n")

	b.WriteString("## Counter-ledger witnesses\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Ledger | Up coefficients | Down coefficients | Charged-lepton coefficients | Hermitian | Trace neutral | Gauge compatible | KMS compatible | Mass-lift compatible | Imports data |\n")
	b.WriteString("|---|---|---|---|---:|---:|---:|---:|---:|---:|\n")
	for _, x := range a.Sieve.Ledgers {
		b.WriteString(fmt.Sprintf("| %s | `%s` | `%s` | `%s` | %t | %t | %t | %t | %t | %t |\n", x.Name, FormatTriple(x.UpCoefficients), FormatTriple(x.DownCoefficients), FormatTriple(x.LeptonCoefficients), x.Hermitian, x.TraceNeutral, x.GaugeCompatible, x.KMSCompatible, x.MassLiftCompatible, x.ImportsEmpiricalData))
	}
	b.WriteString("\n")
	b.WriteString("Because these ledgers are mutually distinct and all pass the same native tests, the boundary intersection does not produce a unique coefficient assignment. This is the decisive obstruction.\n\n")

	b.WriteString("## Coefficient ledger closure\n\n")
	b.WriteString(FormatCoefficientLedger(a.Ledger) + "\n\n")
	b.WriteString(FormatClosure(a.Closure) + "\n\n")
	b.WriteString("The correct registry update is not to promote coefficients, but to close the amplitude lane: structural support is native where proven, numerical flavor amplitudes remain quarantined.\n\n")

	b.WriteString("## Phenomenology/firewall audit\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No empirical flavor datum was imported. No muon/charm mass value, Yukawa coefficient, CKM angle, CKM phase, or PMNS value is predicted.\n\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range []string{StatusGate446BoundaryInherited, StatusKXYCoefficientArenaFormalized, StatusNativeBoundaryStackApplied, StatusFunctionalSelectorSieveCompleted, StatusCounterLedgerWitnessesConstructed, StatusCoefficientFirewallClosed, StatusEmpiricalFirewallPreserved, StatusFailedNoNativeSectorCoefficientRule, StatusFailedMultipleLedgersSurvive, StatusFailedTraceKMSGaugeDoNotSelectValues, StatusFailedNoMuonCharmMassPrediction, StatusFailedNoCKMPMNSPrediction, StatusNineCoefficientsRemainQuarantined} {
		b.WriteString("- `" + s + "`\n")
	}

	b.WriteString("\n## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
