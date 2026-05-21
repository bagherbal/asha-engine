package generation2externalflavororientationsealinstallationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE964-GENERATION2EXTERNALFLAVORORIENTATIONSEALINSTALLATIONAUDIT"
	theoremName = "Gate 964: ExternalFlavorOrientationSeal Installation Audit"
)

func Generation2ExternalFlavorOrientationSealInstallationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 964 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, ComponentSupports(a.Components))
		allFailures := appendAll(a.Failures, ComponentFailures(a.Components))
		checks := []theorem.Check{
			{Name: "inherits Gate 963 seal requirement", Passed: a.Inherited == InheritedStatus && a.Decision.InheritedRequiresExternalFlavorOrientationSeal, Detail: a.Inherited},
			{Name: "ExternalFlavorOrientationSeal installed as non-native quarantine", Passed: a.Decision.ExternalFlavorOrientationSealInstalled && !a.Decision.ExternalFlavorOrientationSealNative, Detail: a.InstalledSeal},
			{Name: "diagnostic representative selected without canonical selector", Passed: a.Decision.RepresentativeChosenForDiagnostics && a.Decision.U3OrbitAcknowledged && !a.Decision.CanonicalFlavorSelectorCertified, Detail: "Phi_flav^seal in [Phi_flav]_{U(3)}"},
			{Name: "triple-sealed lane active", Passed: a.Decision.TripleSealLaneActive && a.Decision.R3DualSealPreserved && a.Decision.ExternalGenerationCarrierSealPreserved, Detail: a.TripleSealLane},
			{Name: "downstream tests allowed only as sealed diagnostics", Passed: a.Decision.DownstreamFlavorLedgerTestsAllowed, Detail: stringsJoin(a.Allowed)},
			{Name: "does not derive Yukawas, CKM/PMNS, particles, flavor theorem, or ledger updates", Passed: !a.Decision.NativeFlavorTheoremCertified && !a.Decision.YukawaEigenvaluesDerived && !a.Decision.CKMPMNSDerived && !a.Decision.PMNSDerived && !a.Decision.PhysicalParticlesAssigned && !a.Decision.OfficialLedgerUpdateAllowed, Detail: stringsJoin(a.Forbidden)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, a.InstalledSeal, a.TripleSealLane, a.Final, a.NextGate}
		notes = append(notes, ComponentNotes(a.Components)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
