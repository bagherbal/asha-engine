package generation2thetaextboundaryresponsepackageconstructionandreadoutobstructionaudit

import (
	"strings"
	"testing"
)

func TestGate785LiftReadoutAndExteriorAlgebra(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate784.Inherited || !strings.Contains(a.Gate784.MissingObject, "Theta_ext") {
		t.Fatalf("bad Gate784 inheritance: %+v", a.Gate784)
	}
	if !closeRel(a.Ledger.M1, 0.0001256543573849177, 1e-14) || !closeRel(a.Ledger.M2, 1.624013231638281e-07, 1e-14) || !closeRel(a.Ledger.M3, 2.0989474869200057e-10, 1e-14) || !a.Ledger.Matches {
		t.Fatalf("bad ledger: %+v", a.Ledger)
	}
	if !a.Separation.Separated || a.Separation.ThetaExtAloneSufficient || !strings.Contains(a.Separation.PolynomialRepresentation, "chi_ext") {
		t.Fatalf("bad separation: %+v", a.Separation)
	}
	if !a.Algebra.Typed || !strings.Contains(a.Algebra.Algebra, "Lambda^2") || !a.Algebra.RequiresDegreeOneAxis || a.Algebra.HasNativeDegreeOneAxis || !a.Algebra.ConditionalLabelledBasis || !strings.Contains(a.Algebra.VolumeForm, "omega_B") {
		t.Fatalf("bad exterior response algebra: %+v", a.Algebra)
	}
}

func TestGate785ConditionalPackageNaturalityMagnitudeAndStop(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Package.Constructed || !strings.Contains(a.Package.ThetaM2, "beta_B") || !strings.Contains(a.Package.ThetaM3, "omega_B") || !a.Package.MatchesFWall3 || a.Package.ReadoutNative {
		t.Fatalf("bad conditional package: %+v", a.Package)
	}
	if !closeRel(a.Package.Chi0, 1, 1e-15) || !closeRel(a.Package.Chi1, a.Ledger.KappaE, 1e-15) || !closeRel(a.Package.Chi2, -2*a.Ledger.P, 1e-15) {
		t.Fatalf("bad package readout coefficients: %+v", a.Package)
	}
	if !a.Naturality.Completed || !a.Naturality.OnlyDimensionTwoNative || a.Naturality.CanonicalNonzeroVectorFromDimension || a.Naturality.CanonicalNonzeroCovectorFromDimension || !a.Naturality.LabelledBasisConditional {
		t.Fatalf("bad naturality audit: %+v", a.Naturality)
	}
	if !a.MagnitudeSign.Separated || !closeRel(a.MagnitudeSign.TwoP, 7.0/36.0, 1e-15) || a.MagnitudeSign.OrientationSignNative || !strings.Contains(a.MagnitudeSign.Sign, "negative") {
		t.Fatalf("bad magnitude/sign audit: %+v", a.MagnitudeSign)
	}
	if !a.CubicStop.Audited || !a.CubicStop.BlockedIfDegreeRule || a.CubicStop.DegreeRuleNative || a.CubicStop.CubicStopNative || !strings.Contains(a.CubicStop.M4Degree, "Lambda^3") || a.CubicStop.M4 <= 0 {
		t.Fatalf("bad cubic stop audit: %+v", a.CubicStop)
	}
}

func TestGate785ExponentialPredictionFirewallsAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Exponential.Audited || a.Exponential.SingleVectorProducesDegreeTwo || !strings.Contains(a.Exponential.SingleVectorExp, "1+beta") || !strings.Contains(a.Exponential.DegreeTwoRequirement, "volume") {
		t.Fatalf("bad exterior exponential audit: %+v", a.Exponential)
	}
	if !a.Prediction.Reclassified || !strings.Contains(a.Prediction.FWall3Status, "Level B+") || a.Prediction.FWallLevelC || !strings.Contains(a.Prediction.CHiggs, "not Level C") {
		t.Fatalf("bad prediction status: %+v", a.Prediction)
	}
	if !a.Firewalls.Enforced || a.Firewalls.ThetaExtPackageNative || a.Firewalls.ConditionalBetaNative || a.Firewalls.ChiExtNative || a.Firewalls.DimensionTwoProof || a.Firewalls.OmegaBSignDerived || a.Firewalls.FWallNative || a.Firewalls.KappaLambdaNative || a.Firewalls.CHistoryIndependent || a.Firewalls.TreeProxyPoleMass || a.Firewalls.YukawaNative {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	if !strings.Contains(a.FinalStatement, "does not construct Theta_ext natively") || !strings.Contains(a.FinalStatement, "chi_ext") || !strings.Contains(a.FinalStatement, "next bottleneck") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
}

func TestGate785TheoremStatuses(t *testing.T) {
	res := Generation2ThetaExtBoundaryResponsePackageConstructionAndReadoutObstructionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
