package coloredoctetportal

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ColoredOctetPureSMPortalSearchSpectrumFalsificationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-COLORED-OCTET-PURE-SM-PORTAL-SEARCH"
	const name = "Colored-octet pure-SM portal search / Spectrum falsification audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 223 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 222 partial relic result is inherited", Passed: a.Gate222.Gate222Inherited && a.Gate222.TripletPortalSupported && a.Gate222.ColoredOctetUnresolved && !a.Gate222.FullRelicSealGranted, Detail: a.Gate222.TruthStatement},
			{Name: "pure-SM tensor search is explicit and dimension bounded", Passed: a.TensorSearch.CombinationsScanned > 0 && a.TensorSearch.DimensionEligibleScanned == a.TensorSearch.CombinationsScanned && a.TensorSearch.SearchLimit != "", Detail: FormatTensorSearch(a.TensorSearch)},
			{Name: "false octet-Q mass mixing remains rejected", Passed: a.TensorSearch.FalseQMixingStillRejected && !a.Firewall.FalseOctetQMixingClaimed, Detail: "(8,2,Y=1/2) is not (3,2,Y=1/6)"},
			{Name: "dimension-6 pure-SM colored-octet portal exists", Passed: a.TensorSearch.OctetPortalFound && len(a.TensorSearch.ValidPortals) > 0 && a.TensorSearch.BestPortal.ValidPortal && a.TensorSearch.BestPortal.TotalDimHalf == 12, Detail: FormatCombination(a.TensorSearch.BestPortal)},
			{Name: "BBN safety bound is computable for the sealed portal", Passed: a.Kinematics.BBNSafeForPerturbativeWilson && a.Kinematics.ConservativeLambdaMaxGeV > a.Kinematics.MBGeV && a.Kinematics.RequiredWidthGeV > 0, Detail: FormatKinematics(a.Kinematics)},
			{Name: "RelicDecaySeal is granted only conditionally on EFT portals", Passed: a.RelicSeal.SealPreviouslyDenied && a.RelicSeal.SealGranted && a.RelicSeal.StillNotFiniteDerived && a.Summary.FullRelicDecaySeal && !a.Summary.Rank1SpectrumFalsified, Detail: FormatRelicSeal(a.RelicSeal)},
			{Name: "fallback spectrum replacement is not triggered", Passed: !a.Fallback.Triggered && !a.Summary.Rank1SpectrumFalsified, Detail: FormatFallback(a.Fallback)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate222Inherited && a.Firewall.ThresholdSpectrumSealActive && a.Firewall.MatchingCorrectionSealActive && a.Firewall.EmpiricalCarrierSealActive && a.Firewall.LeptoquarkDynamicsSealActive && !a.Firewall.NewMediatorInvented && !a.Firewall.LeptoquarkSealViolated && !a.Firewall.FalseOctetQMixingClaimed && !a.Firewall.FiniteOperatorClaimed && !a.Firewall.WilsonCoefficientFixed && !a.Firewall.RelicAbundanceComputed && a.Firewall.BBNUsedAsFilterOnly && !a.Firewall.FallbackTuningPerformed && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "Gate 223 saves the Rank-1 spectrum from immediate colored-relic falsification, but only by granting a phenomenological RelicDecaySeal. The finite core still does not derive the Wilson coefficient, flavor choice, or relic abundance."}}
	}}
}
