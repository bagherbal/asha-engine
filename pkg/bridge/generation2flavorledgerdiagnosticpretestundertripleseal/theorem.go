package generation2flavorledgerdiagnosticpretestundertripleseal

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE965-GENERATION2FLAVORLEDGERDIAGNOSTICPRETESTUNDERTRIPLESEAL"
	theoremName = "Gate 965: FlavorLedger Diagnostic Pretest Under Triple Seal"
)

func Generation2FlavorLedgerDiagnosticPretestUnderTripleSealTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 965 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, DiagnosticSupports(a.Diagnostics))
		allFailures := appendAll(a.Failures, DiagnosticFailures(a.Diagnostics))
		checks := []theorem.Check{
			{Name: "inherits Gate 964 sealed flavor-orientation lane", Passed: a.Inherited == InheritedStatus && a.Decision.InheritedTripleSealActive, Detail: a.Inherited},
			{Name: "preserves all active seals", Passed: a.Decision.R3DualSealPreserved && a.Decision.ScalarSourceSealPreserved && a.Decision.PostOrientationSealPreserved && a.Decision.ExternalGenerationCarrierSealPreserved && a.Decision.ExternalFlavorOrientationSealPreserved, Detail: a.SealLane},
			{Name: "epsilon_e diagnostic allowed only under triple seal", Passed: a.Decision.EpsilonLedgerDiagnosticAllowed, Detail: "epsilon_e as sealed downstream target only"},
			{Name: "kappa diagnostics allowed only under triple seal", Passed: a.Decision.KappaLedgerDiagnosticAllowed, Detail: "kappa_e/kappa_lambda as sealed downstream targets only"},
			{Name: "Koide-shadow diagnostic allowed only under triple seal", Passed: a.Decision.KoideShadowDiagnosticAllowed, Detail: "Koide-shadow compatibility only"},
			{Name: "CKM/PMNS compatibility diagnostic allowed only under triple seal", Passed: a.Decision.CKMPMNSLedgerDiagnosticAllowed, Detail: "ledger compatibility only"},
			{Name: "no native flavor theorem, Yukawa spectrum, CKM/PMNS, particles, or ledger update", Passed: !a.Decision.NativeFlavorTheoremDerived && !a.Decision.YukawaSpectrumDerived && !a.Decision.CKMPNMSTheoremDerived && !a.Decision.PMNSTheoremDerived && !a.Decision.PhysicalParticlesAssigned && !a.Decision.OfficialLedgerUpdateAllowed, Detail: stringsJoin(a.Forbidden)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, a.SealLane, a.Final, a.NextGate}
		notes = append(notes, DiagnosticNotes(a.Diagnostics)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
