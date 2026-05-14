package carrieractivationseal

import "testing"

func TestGate205Inherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.PreviousGate205
	if !p.Gate205Inherited || !p.CarrierActivationObstructed || !p.GaugeChargeObstructed || !p.SpinStatisticsObstructed || !p.MassActivationObstructed || p.ContactModesAudited != 7 || p.ContactModesPromotedToBetaRows || !p.Gate201ShapesRemainConditional || p.PhysicalUnificationClaimed {
		t.Fatalf("bad Gate 205 inheritance: %s", FormatGate205(p))
	}
}

func TestNativeSearchStillObstructed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	n := a.NativeSearch
	if n.ContactModesAudited != 7 || !n.BRSTCohomologyRouteAudited || n.BRSTNonzeroCanonicalDifferential || n.BRSTZeroBetaLedger || !n.CliffordOctonionGradingRouteAudited || n.CanonicalNontrivialParityGrading || n.GaugeChargeFunctorDerived || n.SpinStatisticsFunctorDerived || n.MassActivationPredicateDerived || n.NativeCarrierActivationDerived {
		t.Fatalf("native semantic firewall leak: %s", FormatNativeSearch(n))
	}
}

func TestEmpiricalCarrierSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.Seal
	if s.Name != "EmpiricalCarrierSeal" || !s.ExplicitAxiom || !s.Quarantined || !s.RequiredByGate205 || !s.BypassesChargeSemantics || !s.BypassesSpinStatisticsSemantics || !s.BypassesMassActivationSemantics || s.UsesObservedInputForFiniteCore || s.CarriesFiniteDerivationClaim || !s.AllowsConditionalThresholdCarriers || len(s.AllowedRepresentations) != 2 || s.ConditionalStatus != StatusConditionalOnCarrierSeal {
		t.Fatalf("bad carrier seal: %s", FormatSeal(s))
	}
}

func TestAnomalyCompatibility(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Anomaly.ChecksAudited != 2 || !a.Anomaly.AllPerturbativeAnomaliesZero || !a.Anomaly.AllGlobalSU2WittenSafe || !a.Anomaly.AllMixedGravitationalSafe || !a.Anomaly.AllCarriersCompatible || !a.Anomaly.CombinedVector.Zero() {
		t.Fatalf("anomaly audit failed: %s", FormatAnomalyAudit(a.Anomaly))
	}
	for _, c := range a.AnomalyChecks {
		if !c.PerturbativeGaugeAnomalyFree || !c.GlobalSU2WittenSafe || !c.MixedGravitationalSafe || !c.Vector.Zero() {
			t.Fatalf("carrier anomaly failed: %s", FormatAnomalyCheck(c))
		}
	}
}

func TestConditionalPredictions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Predictions) != 2 || a.PredictionAudit.PredictionsEmitted != 2 || !a.PredictionAudit.AllAnomalyCompatible || !a.PredictionAudit.AllCloseUOneBoundary || !a.PredictionAudit.AllOrderedPositiveScales || !a.PredictionAudit.AllConditionalOnCarrierSeal || !a.PredictionAudit.UniversalCompletionStillExternal || !a.PredictionAudit.AlphaGUTFixedByUOneSeal || a.PredictionAudit.AbsoluteMassPredictionClaimed || a.PredictionAudit.PhysicalUnificationClaimed {
		t.Fatalf("bad prediction audit: %s :: %s", FormatPredictionAudit(a.PredictionAudit), FormatPredictions(a.Predictions))
	}
	for _, p := range a.Predictions {
		if p.ThresholdScaleMBGeV <= 0 || p.BoundaryScaleMStarGeV <= p.ThresholdScaleMBGeV || p.MaxClosureResidual >= 1e-7 || p.TriangleAreaAfterActivation != 0 || !p.ConditionalOnCarrierSeal || !p.ConditionalOnUniversalCompletion || !p.AnomalyCompatible || p.FiniteDerived || p.AbsolutePredictionClaimed {
			t.Fatalf("bad conditional prediction: %s", FormatPrediction(p))
		}
	}
}

func TestFirewallSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if !f.Gate205Inherited || !f.NativeSearchObstructed || !f.CarrierSealExplicit || !f.CarrierSealQuarantined || f.ObservedInputUsedForFiniteCore || f.ContactModesPromotedWithoutSeal || f.ContactModesClaimedFiniteParticles || f.UniversalBetaSourceDerived || f.FiniteMatchingCorrectionsDerived || f.AbsoluteMassPredicted || f.PhysicalUnificationClaimed || f.ThresholdCorrectedPhysicalFitClaimed || !f.NumericalPredictionsConditional {
		t.Fatalf("firewall leak: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != f.StrictNullityAfter || f.CarrierSealNullityBefore != 1 || f.CarrierSealNullityAfter != 0 || f.PhysicalPredictionNullityBefore != f.PhysicalPredictionNullityAfter {
		t.Fatalf("unexpected nullity accounting: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := CarrierActivationSealLocalFieldSemanticBifurcationAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
