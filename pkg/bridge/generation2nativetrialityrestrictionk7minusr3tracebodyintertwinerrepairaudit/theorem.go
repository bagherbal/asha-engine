package generation2nativetrialityrestrictionk7minusr3tracebodyintertwinerrepairaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE956-GENERATION2NATIVETRIALITYRESTRICTIONK7MINUSR3TRACEBODYINTERTWINERREPAIRAUDIT"
	theoremName = "Gate 956: Native Triality Restriction to K7Minus and R3 Tracebody Intertwiner Repair Audit"
)

func Generation2NativeTrialityRestrictionK7MinusR3TracebodyIntertwinerRepairAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 956 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, ItemSupports(a.Items))
		allFailures := appendAll(a.Failures, ItemFailures(a.Items))
		checks := []theorem.Check{
			{Name: "inherits Gate 955 abstract C3/no-generation-map wound", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "preserves dual-sealed R3 and blocks downstream overclaim", Passed: a.R3DualSealRequired && !a.NativeR3 && !a.OfficialLedgerUpdate && !a.GenerationCarrierCertified && !a.FlavorOrientationCertified && !a.IndividualYukawaCertified && !a.PhysicalAssignmentCertified, Detail: "R3 dual seal preserved; no generation/flavor/Yukawa overclaim"},
			{Name: "native triality operator remains absent", Passed: !a.Transport.NativeTrialityOperatorConstructed && !a.Transport.OrderThree && !a.Transport.Nontrivial, Detail: "no T_tri^native supplied"},
			{Name: "triality transport chain remains uncertified", Passed: !a.Transport.TransportToLambda4Certified && !a.Transport.PreservesLambda4 && !a.Transport.PreservesK7ContactCarrier && !a.Transport.PreservesK7Minus && a.Transport.LeakageToK7PlusUnknown, Detail: "no Lambda4/K7/K7^- restriction certificate"},
			{Name: "Gate 955 abstract C3 model remains noncanonical", Passed: !a.Transport.AbstractC3FromGate955Realized, Detail: "abstract action not native-realized"},
			{Name: "R3 tracebody intertwiner remains uncertified", Passed: !a.Intertwiner.Certified && a.Intertwiner.RequiresNativeK7MinusAction && a.Intertwiner.UsesArbitraryBasisFit && a.Intertwiner.UsesR3RowsAsGenerationLabels && a.Intertwiner.PreservesR3DualSeal, Detail: a.Intertwiner.CandidateName},
			{Name: "noncircularity and flavor firewall preserved", Passed: !a.Intertwiner.UsesFlavorBacksolve && !a.Intertwiner.UsesObservedMassesOrYukawas && !a.Intertwiner.UsesCKMPMNSInput, Detail: "no CKM/PMNS/mass/Yukawa/flavor backsolve input"},
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
