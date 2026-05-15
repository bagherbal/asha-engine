package generation2phaseorientation

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Conclusion.SignedCycleForced || a.Conclusion.ComplexPhaseForced || a.Conclusion.YGenPromotedToNative {
		t.Fatalf("Gate 446 must not promote a phase/orientation: %+v", a.Conclusion)
	}
	if !a.Firewall.YGenRemainsQuarantined || !a.Firewall.SignedCycleOrientationSealed || !a.Firewall.ComplexPhaseSealed {
		t.Fatalf("phase firewall not preserved: %+v", a.Firewall)
	}
}

func TestRealSignSieveTwoGaugeClasses(t *testing.T) {
	s := buildRealSignSieve()
	if s.UniqueSignedCycle {
		t.Fatalf("signed cycle should not be unique: %s", FormatRealSignSieve(s))
	}
	if len(s.Candidates) != 8 || s.PositiveCycleCount != 4 || s.NegativeCycleCount != 4 || s.Z2GaugeClasses != 2 {
		t.Fatalf("unexpected real sign counts: %s", FormatRealSignSieve(s))
	}
	for _, c := range s.Candidates {
		if !c.MassLiftCompatible || !c.JGammaCompatible || !c.EtaTraceNeutral {
			t.Fatalf("candidate should pass structural boundaries: %s", FormatRealSignCandidate(c))
		}
	}
}

func TestComplexPhaseSieveContinuumAndCPPairs(t *testing.T) {
	s := buildComplexPhaseSieve()
	if !s.ContinuumSurvives || !s.CPConjugatePairsSurvive || s.UniqueComplexPhase || s.CPPhaseValuePredicted {
		t.Fatalf("unexpected complex phase sieve: %s", FormatComplexPhaseSieve(s))
	}
	var plus, minus, pureImag bool
	for _, sample := range s.Samples {
		switch sample.Label {
		case "pi/4":
			plus = sample.MassLiftCompatible && sample.CPCapable
		case "-pi/4":
			minus = sample.MassLiftCompatible && sample.CPCapable
		case "pi/2":
			pureImag = !sample.MassLiftCompatible && sample.CPCapable
		}
	}
	if !plus || !minus || !pureImag {
		t.Fatalf("phase samples do not show the intended survivor/failure split: %s", FormatComplexPhaseSieve(s))
	}
}

func TestTheoremPassesAsFailedRouteAudit(t *testing.T) {
	res := Generation2SignedCycleComplexPhaseOrientationSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit checks did not pass:\n%s", res.Details())
	}
	if string(res.Status) == "" {
		t.Fatalf("empty theorem status")
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusFailedSignedOrientationNotUnique, StatusFailedComplexPhaseContinuum, StatusFailedYGenNotNative, "Phi = arg(z12 z23 conjugate(z13))"} {
		if !stringsContains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func stringsContains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
