package generation2flavorspectralbalancefunctionaltypeadmissibilityaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2FlavorSpectralBalanceFunctionalTypeAdmissibilityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 flavor spectral balance functional type-admissibility audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate595 flavor spectral type-admissibility audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate594 B_flav functional and residual", Passed: a.Inherited.ResidualInsideSigma && almostEqual(a.Inherited.BFlav, -a.Inherited.Delta590, 1e-18), Detail: FormatInherited(a.Inherited)},
			{Name: "type all three B_flav terms", Passed: a.TermTyping.Complete && a.TermTyping.Epsilon.RequiresFractional && a.TermTyping.PMNS.RequiresSpectralCalc && a.TermTyping.CKM.RequiresSpectralCalc, Detail: FormatTermTyping(a.TermTyping)},
			{Name: "certify B_flav as environmental spectral functional", Passed: a.Final.BFlavEnvironmentalWellDefined && a.Admissibility.SpectralProjectorsAdmitted && a.Admissibility.NormalizedCKMAdmitted, Detail: FormatAdmissibility(a.Admissibility)},
			{Name: "locate primary native obstruction at epsilon(H_e)", Passed: a.Obstruction.EpsilonHEBlocked && a.Obstruction.PMNSMoreAdmissible && a.Obstruction.CKMMoreAdmissible && a.Final.PrimaryNativeObstruction != "", Detail: FormatObstruction(a.Obstruction)},
			{Name: "reject fourth-root/root-trace/chamber wall native promotion", Passed: !a.Admissibility.FractionalFourthRootAdmitted && !a.Admissibility.RootTraceAdmitted && !a.Admissibility.ChamberWallFunctionalAdmitted, Detail: FormatAdmissibility(a.Admissibility)},
			{Name: "reject native PMNS/CKM and cross-sector balance theorem", Passed: !a.TermTyping.PMNS.NativePresent && !a.TermTyping.CKM.NativePresent && !a.Admissibility.CrossSectorEquationAdmitted && !a.Final.NativeBFlavZeroTheoremPresent, Detail: FormatFinal(a.Final)},
			{Name: "define exact promotion requirements", Passed: !a.Requirements.AllPresent && len(a.Requirements.Items) == 6 && a.Requirements.ExactMissingTheorem != "", Detail: FormatRequirements(a.Requirements)},
			{Name: "preserve flavor/root-trace firewalls and avoid residual fitting", Passed: !a.Firewalls.FitsNewResiduals && !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesYukawas && !a.Firewalls.DerivesNeutrinoPhysics && !a.Firewalls.DerivesFlavorTexture && !a.Firewalls.PromotesLedgers && !a.Firewalls.AddsCarrier && !a.Firewalls.AddsSelector && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth, a.Final.Decision)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
