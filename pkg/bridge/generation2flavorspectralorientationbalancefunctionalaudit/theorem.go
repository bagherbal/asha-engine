package generation2flavorspectralorientationbalancefunctionalaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2FlavorSpectralOrientationBalanceFunctionalAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 flavor spectral orientation balance functional audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate594 flavor spectral balance audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate593 orientation-balance invariant values", Passed: a.Inherited.ResidualInsideSigma && a.Inherited.PMNSTrace == 4*a.Inherited.PMNSQuarter && almostEqual(a.Inherited.BFlavValue, -a.Inherited.Delta590, 1e-18), Detail: FormatInherited(a.Inherited)},
			{Name: "define common flavor spectral algebra", Passed: a.Algebra.Name == "A_flav" && len(a.Algebra.Generators) == 4 && a.Algebra.ObservedOnly && !a.Algebra.NativeAlgebraPresent, Detail: FormatAlgebra(a.Algebra)},
			{Name: "define charged-lepton root-spectrum epsilon functional", Passed: a.ChargedLepton.RequiresSpectralCalculus && a.ChargedLepton.RequiresFourthRootOfHE && a.ChargedLepton.RequiresChamberOrderSeal && !a.ChargedLepton.NativeFunctionalPresent && a.ChargedLepton.Gate352ObstructionPreserved, Detail: FormatChargedLepton(a.ChargedLepton)},
			{Name: "write PMNS reactor term as spectral projector overlap", Passed: a.PMNS.TraceValue == a.Inherited.PMNSTrace && a.PMNS.QuarterValue == a.Inherited.PMNSQuarter && !a.PMNS.NativeDerivation && a.PMNS.ObservedLedger && len(a.PMNS.RequiredLabels) == 4, Detail: FormatPMNS(a.PMNS)},
			{Name: "write CKM term as normalized spectral commutator area", Passed: a.CKM.JValue == a.Inherited.JCKM && a.CKM.RequiresNondegenerateSpectra && !a.CKM.NativeDerivation && a.CKM.ObservedLedger && len(a.CKM.RequiredLabels) == 4, Detail: FormatCKM(a.CKM)},
			{Name: "construct B_flav and reproduce Gate590/593 residual with sign", Passed: a.Balance.LeftKappa == a.Inherited.KappaObs && a.Balance.RightHandCandidate == a.Inherited.RightHandCandidate && almostEqual(a.Balance.BFlav, -a.Balance.Delta590, 1e-18) && a.Balance.ResidualInsideSigma, Detail: FormatBalance(a.Balance)},
			{Name: "audit invariance and label seals", Passed: a.Invariance.AllRequiredSealsNamed && a.Invariance.BasisInvariantWithSeals && len(a.Invariance.Items) == 5, Detail: FormatInvariance(a.Invariance)},
			{Name: "audit ASHA availability and reject native B_flav theorem", Passed: len(a.Availability.Items) >= 5 && !a.Availability.NativeFlavorSpectralAlgebra && !a.Availability.NativeEpsilonFunctional && !a.Availability.NativePMNSProjector && !a.Availability.NativeCKMCommutator && !a.Availability.NativeBFlavZeroTheorem, Detail: FormatAvailability(a.Availability)},
			{Name: "define missing FlavorSpectralOrientationBalanceFunctional theorem target", Passed: a.Target.Name == "FlavorSpectralOrientationBalanceFunctional" && a.Target.MustUseSpectralCalculus && a.Target.MustBeBasisInvariantWithSeals && a.Target.MustPreserveSectorTyping && !a.Target.NativePresent, Detail: FormatTarget(a.Target)},
			{Name: "preserve flavor/root-trace firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesYukawas && !a.Firewalls.DerivesNeutrinoData && !a.Firewalls.DerivesFlavorTexture && !a.Firewalls.PromotesObservedData && !a.Firewalls.AddsNewCarrier && !a.Firewalls.AddsNewSelector && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return B_flav=0 as environmental only", Passed: a.Final.AllTermsInOneSpectralAlgebra && !a.Final.NativeBFlavOperatorPresent && a.Final.BFlavEnvironmental, Detail: FormatFinal(a.Final)},
		}
		notes := append(Statuses(), a.Truth, a.Final.Decision)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
