package generation2sealedyukawamatrixoperatorconstructionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE973-GENERATION2SEALEDYUKAWAMATRIXOPERATORCONSTRUCTIONAUDIT"
	theoremName = "Gate 973: Sealed Yukawa Matrix Operator Construction Audit"
)

func Generation2SealedYukawaMatrixOperatorConstructionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 973 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "inherits sealed rail", Passed: a.Decision.InheritedSealedRail && a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "preserves R3 dual seal and external orientation/generation seals", Passed: a.Decision.R3DualSealPreserved && a.Decision.ScalarSourceSealPreserved && a.Decision.PostOrientationSealPreserved && a.Decision.ExternalGenerationCarrierSealPreserved && a.Decision.ExternalFlavorOrientationSealPreserved, Detail: a.SealLane},
			{Name: "allows only sealed operation under audit", Passed: a.Decision.AllowsSealedOperation, Detail: stringsJoin(a.Allowed)},
			{Name: "does not derive native flavor, Yukawa matrix, individual values, CKM/PMNS, particles, or official ledger", Passed: !a.Decision.DerivesNativeFlavor && !a.Decision.DerivesNativeYukawaMatrix && !a.Decision.DerivesIndividualYukawas && !a.Decision.DerivesCKMPMNS && !a.Decision.AssignsPhysicalParticles && !a.Decision.UpdatesOfficialLedger, Detail: stringsJoin(a.Forbidden)},
			{Name: "records required supports", Passed: containsAll(a.Supports, RequiredSupports()), Detail: stringsJoin(RequiredSupports())},
			{Name: "preserves required firewalls", Passed: containsAll(a.Failures, RequiredFailures()), Detail: stringsJoin(RequiredFailures())},
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
		notes = append(notes, a.Allowed...)
		notes = append(notes, a.Forbidden...)
		notes = append(notes, a.MatrixNormalForm...)
		notes = append(notes, a.Supports...)
		notes = append(notes, a.Failures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
