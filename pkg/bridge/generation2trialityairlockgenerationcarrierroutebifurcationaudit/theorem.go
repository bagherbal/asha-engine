package generation2trialityairlockgenerationcarrierroutebifurcationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE957-GENERATION2TRIALITYAIRLOCKGENERATIONCARRIERROUTEBIFURCATIONAUDIT"
	theoremName = "Gate 957: Triality Airlock and GenerationCarrier Route Bifurcation Audit"
)

func Generation2TrialityAirlockGenerationCarrierRouteBifurcationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 957 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, ItemSupports(a.Items))
		allFailures := appendAll(a.Failures, ItemFailures(a.Items))
		checks := []theorem.Check{
			{Name: "inherits Gate 956 K7-minus/triality blocker", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "preserves R3 dual seal and avoids downstream overclaim", Passed: a.R3DualSealRequired && !a.NativeR3 && !a.GenerationCarrierCertified && !a.Decision.GenerationCarrierCertified && !a.Decision.FlavorOrientationCertified && !a.Decision.IndividualYukawaCertified && !a.Decision.PhysicalAssignmentCertified && !a.Decision.OfficialLedgerUpdate, Detail: "dual seal and flavor/Yukawa/official firewalls preserved"},
			{Name: "parent triality board not certified in active board", Passed: !a.Parent.Identified && !a.Parent.NativeD4Spin8SourceCertified && !a.Parent.ReplacesGate955AbstractC3, Detail: TrialityParentBoardName},
			{Name: "triality airlock to Lambda4/K7 is absent", Passed: !a.Airlock.ToLambda4Certified && !a.Airlock.ToK7Certified && !a.Airlock.PreservesK7, Detail: TrialityAirlockName},
			{Name: "K7-minus route is not reopened", Passed: !a.Airlock.PreservesK7Minus && !a.Decision.K7MinusRouteReopened, Detail: "no native restriction to K7^-"},
			{Name: "no alternative three-carrier selected by triality airlock", Passed: !a.Airlock.SelectsAlternativeThreeCarrier && a.Airlock.SelectsNoCanonicalThreeCarrier && !a.Decision.AlternativeCarrierSelected, Detail: "no canonical alternate carrier from triality airlock"},
			{Name: "route bifurcates to alternative carrier search", Passed: a.Decision.AlternativeSearchRequired && a.Decision.Decision == Verdict, Detail: a.Decision.Decision},
			{Name: "K7 polarity dimensions retained as candidate context", Passed: a.K7MinusDimension == 3 && a.K7PlusDimension == 4 && a.K7MinusDimension+a.K7PlusDimension == K7Dim, Detail: "K7^-=3, K7^+=4"},
			{Name: "R3 trace rows remain aggregate only", Passed: a.R3TraceRows == R3TraceRows, Detail: a.R3TraceRows},
			{Name: "records required supports", Passed: containsAll(allSupports, RequiredSupports()), Detail: stringsJoin(RequiredSupports())},
			{Name: "preserves required firewalls", Passed: containsAll(allFailures, RequiredFailures()), Detail: stringsJoin(RequiredFailures())},
			{Name: "matches verdict and classification", Passed: a.Verdict == Verdict && a.Classification == Classification && a.ShortStatus == ShortStatus, Detail: a.Verdict},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, a.Final, NextGate}
		notes = append(notes, ItemNotes(a.Items)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
