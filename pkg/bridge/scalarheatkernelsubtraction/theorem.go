package scalarheatkernelsubtraction

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ScalarHeatKernelSubtractionHiggsPotentialChannelSeparationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-SCALAR-HEAT-KERNEL-SUBTRACTION-HIGGS-POTENTIAL-CHANNEL-SEPARATION-AUDIT"
	const name = "Scalar Heat-Kernel Subtraction / Higgs Potential Channel Separation Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 305 scalar heat-kernel subtraction audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 304 f0 promotion is inherited only for a4 normalization while higher moments remain open", Passed: a.Input.ContactF0Promoted && a.Input.PromotedF0Value == 7 && a.Input.F0Positive && a.Input.KineticNormalizationAnchored && !a.Input.HigherMomentsLocked && !a.Input.UniqueProfileShapeDerived && !a.Input.HeatKernelSubtractionClaimed && !a.Input.NumericalZHComputed && !a.Input.HiggsMassPredictionClaimed && !a.Input.NumericalYukawasInserted, Detail: FormatGate304Inheritance(a.Input)},
			{Name: "raw a2 coefficient is decomposed into vacuum/reference carrier and scalar-power-2 dynamical candidate", Passed: a.A2.DecompositionFormalized && a.A2.FieldIndependentVacuumSeen && a.A2.ScalarPower2ChannelSeen && a.A2.MixedTermsFirewalled && !a.A2.NumericalCoefficientUsed && len(a.A2.Components) >= 3, Detail: FormatRawA2(a.A2)},
			{Name: "vacuum-referenced scalar a2 subtraction scheme is formalized without unique counterterm claim", Passed: a.Subtraction.Formalized && a.Subtraction.LinearityRequired && a.Subtraction.GaugeCovariant && !a.Subtraction.BackgroundIndependent && !a.Subtraction.SchemeUnique && !a.Subtraction.CountertermPhysicallyFixed && !a.Subtraction.NumericalCountertermUsed && len(a.Subtraction.SubtractedVacuumPieces) > 0 && len(a.Subtraction.RetainedDynamicalPieces) > 0, Detail: FormatSubtraction(a.Subtraction)},
			{Name: "Higgs mass parameter extraction map uses subtracted a2 and ZH normalization but computes no number", Passed: a.MassMap.MapFormalized && a.MassMap.UsesGate300Normalization && a.MassMap.UsesSubtractedA2 && a.MassMap.RequiresPositiveZH && a.MassMap.RequiresF2 && a.MassMap.RequiresCutoffScale && a.MassMap.RequiresYukawaAmplitudes && !a.MassMap.NumericalMassComputed, Detail: FormatMassMap(a.MassMap)},
			{Name: "f2 dependency sieve records that f0=7 does not lock the a2 moment", Passed: a.F2.DependencyFormalized && !a.F2.F2LockedByGate304 && a.F2.SameProfileCouldVaryF2 && a.F2.RequiresProfileShapeRule && a.F2.RequiresCutoffScaleLambda && !a.F2.CanClaimMassWithoutF2 && a.F2.PredictivePowerLostIfFreeF2, Detail: FormatF2(a.F2)},
			{Name: "channel separation isolates the quadratic Higgs candidate without disturbing a4 kinetic/gauge or quartic channels", Passed: a.Channels.QuadraticChannelIsolated && a.Channels.VacuumChannelSubtracted && a.Channels.A4F0SealPreserved && a.Channels.A2F2StillOpen && a.Channels.NoDynamicsOverclaimed && !a.Channels.QuarticChannelTouched && !a.Channels.GaugeKineticDisturbed, Detail: FormatChannels(a.Channels)},
			{Name: "mass, f2, Lambda, ZH, Yukawa, quartic, subtraction uniqueness, and B-gap firewalls are preserved", Passed: a.Firewalls.NoNumericalF2Inserted && a.Firewalls.NoCutoffScaleInserted && a.Firewalls.NoYukawaNumbersInserted && a.Firewalls.NoNumericalZHComputed && a.Firewalls.NoHiggsMassPredictionClaimed && a.Firewalls.NoHiggsQuarticPredictionClaimed && a.Firewalls.NoBGapInstantonClaimed && a.Firewalls.NoUniqueSubtractionSchemeClaimed && a.Firewalls.NoProfileHigherMomentLockClaimed && a.Firewalls.F0SealPreservedOnlyForA4 && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary authorizes the subtraction algorithm but refuses numerical Higgs dynamics", Passed: a.Summary.Gate304Inherited && a.Summary.RawA2Decomposed && a.Summary.SubtractionSchemeFormalized && a.Summary.QuadraticChannelSeparated && a.Summary.MassMapFormalized && a.Summary.F2DependencyFormalized && a.Summary.F0SealPreserved && !a.Summary.NumericalHiggsMassDerived && !a.Summary.PhysicalDynamicsDerived && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 305 is a subtraction/channel-separation theorem, not a Higgs-mass theorem.", "The legal mass map remains conditional on f2, Lambda, numerical amplitude data, absolute ZH, and a declared renormalization/subtraction convention."}}
	}}
}
