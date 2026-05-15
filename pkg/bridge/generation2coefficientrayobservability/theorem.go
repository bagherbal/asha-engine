package generation2coefficientrayobservability

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2CoefficientRayObservabilityRankAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 coefficient-ray observability rank audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate454 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate 453 empirical interface and post-444 structural laws", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate445TriangleForced && a.Inheritance.Gate447CoefficientsSealed && a.Inheritance.Gate450RatiosRequireAmplitudes && a.Inheritance.Gate452NearestNeighborNotGauge && a.Inheritance.Gate453EmpiricalInterfaceDefined && a.Inheritance.Gate453PromotionRejected && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "derives two-dimensional projective coefficient ray", Passed: a.Ray.Executed && a.Ray.ProjectiveDimension == ProjectiveRayDOF && a.Ray.NativeSelectorCount == 0 && len(a.Ray.RayParameters) == 2, Detail: FormatRayModel(a.Ray)},
			{Name: "spectrum-only comparator has rank one and remains underdetermined", Passed: a.Rank.Executed && a.Rank.SpectrumOnlyRank == 1 && a.Rank.SpectrumOnlyResidualDOF == 1 && a.Rank.SpectrumOnlyRejected, Detail: FormatRankAudit(a.Rank)},
			{Name: "two labelled scalar comparators identify the ray locally", Passed: a.Rank.MinimumLocalScalars == 2 && a.Rank.TwoScalarLocalWorks && a.Rank.GenericJacobianNonzero, Detail: FormatRankAudit(a.Rank)},
			{Name: "CP-oriented global branch remains explicitly tagged", Passed: a.Rank.MinimumOrientedScalars == 3 && a.Rank.CPBranchTagRequired, Detail: FormatRankAudit(a.Rank)},
			{Name: "comparator protocol rejects native coefficient promotion", Passed: a.Protocol.Executed && a.Protocol.AllowsNativeLedger && a.Protocol.AllowsLocalRayFit && a.Protocol.AllowsCPOrientedRayFit && a.Protocol.RequiresExplicitEmpiricalLabel && a.Protocol.RequiresSectorTag && a.Protocol.RequiresRenormalizationTag && a.Protocol.RequiresBranchTagForCPOrientation && !a.Protocol.AllowsNativeCoefficientClaim && !a.Protocol.AllowsSpectrumOnlyRayClaim, Detail: FormatProtocol(a.Protocol)},
			{Name: "13-moduli and coefficient firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoCurveFitPromoted && a.Firewall.NoGSTPromotion && a.Firewall.NoNativeCoefficientRayValue && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate is a dry-run adapter firewall test", Passed: a.Next.Gate == 455, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusGate453Inherited, StatusRayDimensionDerived, StatusSpectrumOnlyRankOne, StatusTwoScalarLocalRankTwo, StatusCPBranchTagRequired, StatusComparatorProtocolDefined, StatusNoNativeCoefficientValues, StatusEmpiricalFirewallPreserved, StatusFailedSpectrumOnlyUnderdetermined, StatusFailedNativeCoefficientSelectorAbsent, StatusFailedCPOrientationNotNative, a.Truth}}
	}}
}
