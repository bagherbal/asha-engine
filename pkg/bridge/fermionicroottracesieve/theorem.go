package fermionicroottracesieve

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FermionicEffectiveActionRootTracePfaffianSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-FERMIONIC-EFFECTIVE-ACTION-ROOT-TRACE-PFAFFIAN-SIEVE"
	const name = "Fermionic Effective Action / Root-Trace (Pfaffian) Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 352 fermionic root-trace audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "span inherits Gate 351 without adding a fit", Passed: a.Span.InheritedGate == 351 && !a.Span.AddsFit, Detail: FormatSpan(a.Span)},
			{Name: "fermionic Pfaffian effective action is formalized", Passed: a.FermionicAction.Formalized, Detail: FormatFermionicAction(a.FermionicAction)},
			{Name: "Pfaffian gives root determinant not root trace", Passed: a.Pfaffian.Executed && a.Pfaffian.RootDeterminant > 0 && a.Pfaffian.RootTrace > 0 && !a.Pfaffian.PfaffianCanGenerateKoide, Detail: FormatPfaffian(a.Pfaffian)},
			{Name: "root-trace observable is audited and remains non-native", Passed: a.RootTrace.Audited && a.RootTrace.BosonicEvenTraceBarrier && !a.RootTrace.RootTraceNative, Detail: FormatRootTrace(a.RootTrace)},
			{Name: "Dixmier/contact trace does not lock finite Yukawa root trace", Passed: a.Dixmier.Audited && a.Dixmier.FiniteRankYukawaDixmierZero && !a.Dixmier.LocksYukawaRootTrace, Detail: FormatDixmier(a.Dixmier)},
			{Name: "Koide promotion remains empirical", Passed: a.KoidePromotion.Executed && a.KoidePromotion.EmpiricalAlignment && !a.KoidePromotion.NativePromotion, Detail: FormatKoidePromotion(a.KoidePromotion)},
			{Name: "parameter census remains fifteen", Passed: a.Census.StartingVacuumInputs == 15 && a.Census.AdditionalReduction == 0 && a.Census.RemainingVacuumInputs == 15 && !a.Census.SevenSealTargetReached, Detail: FormatCensus(a.Census)},
			{Name: "summary preserves root-trace firewall", Passed: a.Summary.Executed && !a.Summary.RootTracePromoted && a.Summary.RemainingVacuumInputs == 15, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 352 proves that Pfaffian square-root mechanics are real but insufficient for Koide: a root determinant is not a root trace."}}
	}}
}
