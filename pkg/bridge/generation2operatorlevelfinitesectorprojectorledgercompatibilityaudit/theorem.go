package generation2operatorlevelfinitesectorprojectorledgercompatibilityaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_OPERATOR_LEVEL_FINITE_SECTOR_PROJECTOR_LEDGER_COMPATIBILITY_AUDIT"
	theoremName = "Gate 888 — Operator-Level FiniteSector ProjectorLedger Compatibility Audit Under Dual Seal"
)

func Generation2OperatorLevelFiniteSectorProjectorLedgerCompatibilityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "define oriented finite-sector projector ledger", Passed: a.Ledger.Name == FiniteSectorLedger && len(a.Ledger.Atoms) == 3 && containsAll(a.Ledger.Supports, []string{SupportOrientedLedgerExists, SupportR3CandidateUnderDualSeal}), Detail: FormatLedger(a.Ledger)},
			{Name: "projectors are idempotent orthogonal complete on H_R^min", Passed: a.Ledger.Idempotent && a.Ledger.Orthogonal && a.Ledger.CompleteOnHRMin && a.Ledger.Rank == RankHRMin && allAtomsProjectorOK(a.Ledger.Atoms), Detail: FormatLedger(a.Ledger)},
			{Name: "ledger stable under A_F^orient but not full A_F", Passed: a.Ledger.StableUnderAFOrient && !a.Ledger.StableUnderFullAF && containsAll(a.Ledger.Failures, []string{FailurePostOrientNotFullAF, FailureNoNativeFiniteSectorTheorem}), Detail: FormatLedger(a.Ledger)},
			{Name: "edge and trace readout compatibility", Passed: a.Ledger.EdgeCompatible && a.Ledger.TraceReadoutCompatible && near(a.Ledger.TraceTotal, 3+3*AlphaB) && near(a.Ledger.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) && near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) && near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic), Detail: FormatLedger(a.Ledger)},
			{Name: "dual seal dependency explicit", Passed: a.Seals.BoundaryAlphaSealSuppliesWeights && a.Seals.PostOrientationSealSuppliesProjectors && !a.Seals.NativeAlphaFunctorCertified && !a.Seals.NativeFullAFProjectorsCertified && !a.Seals.OfficialR3Eligible && containsAll(a.Seals.Failures, []string{FailureAlphaStillSealed, FailureNoNativeFiniteSectorTheorem, FailureNoNativeR3SectorLedger}), Detail: FormatSeals(a.Seals)},
			{Name: "preserve physical generation flavor and individual-yukawa firewalls", Passed: !hasPhysicalLeak(a) && containsAll(a.Ledger.Failures, []string{FailureSocketAtomsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}), Detail: FormatLedger(a.Ledger)},
			{Name: "preserve official ledger freeze", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) && !near(a.Freeze.OperatorCYukawa, a.Freeze.OfficialCYukawa), Detail: FormatFreeze(a.Freeze)},
			{Name: "preserve Gate 888 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatSeals(a.Seals), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func allAtomsProjectorOK(atoms []ProjectorAtom) bool {
	for _, atom := range atoms {
		if !atom.Idempotent || !atom.OrthogonalToOthers || !atom.CompletePartOfHRMin || !atom.StableUnderAFOrient || atom.StableUnderFullAF || !atom.EdgeCompatible {
			return false
		}
	}
	return true
}

func hasPhysicalLeak(a Audit) bool {
	for _, atom := range a.Ledger.Atoms {
		if atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			return true
		}
	}
	return false
}
