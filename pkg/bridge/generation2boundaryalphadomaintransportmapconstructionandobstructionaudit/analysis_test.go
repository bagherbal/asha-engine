package generation2boundaryalphadomaintransportmapconstructionandobstructionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate828SupportTraceWeightsAndAlpha(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Weights.LinearWeight-3.0/10.0) > 1e-15 {
		t.Fatalf("bad linear support weight: %s", FormatWeights(a.Weights))
	}
	if math.Abs(a.Weights.QuadraticWeight-7.0/72.0) > 1e-15 {
		t.Fatalf("bad quadratic support weight: %s", FormatWeights(a.Weights))
	}
	if math.Abs(a.Ledger.AlphaB-0.0003878958469680527) > 1e-18 {
		t.Fatalf("bad alpha reconstruction: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.LinearAlpha+a.Ledger.QuadraticAlpha-a.Ledger.AlphaB) > 1e-21 {
		t.Fatalf("alpha lanes do not add: %s", FormatLedger(a.Ledger))
	}
}

func TestGate828TransportObstructionCriteria(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Criteria.HasSourceScalar || !a.Criteria.HasTypedTargetCarriers || !a.Criteria.HasSupportTraceWeights {
		t.Fatalf("candidate criteria missing typed inputs: %s", FormatCriteria(a.Criteria))
	}
	if a.Criteria.HasConcreteLinearMap || a.Criteria.HasConcreteQuadraticMap || a.Criteria.HasSharedFunctor || a.Criteria.HasPowerLawDerivation || a.Criteria.HasVariationalPrinciple || a.Criteria.CertifiesNativeAlphaTheorem {
		t.Fatalf("transport over-promoted: %s", FormatCriteria(a.Criteria))
	}
	for _, want := range []string{FailureNoBoundaryAlphaMap, FailureNoLinearTransport, FailureNoQuadraticTransport, FailureNoSharedFunctor, FailureNoNativeAlphaTheorem} {
		if !containsAll(a.Criteria.Failures, []string{want}) {
			t.Fatalf("missing failure %s in criteria: %+v", want, a.Criteria.Failures)
		}
	}
}

func TestGate828NonCircularityFreezeAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.NonCircular.UsesNEffToDefineAlpha || a.NonCircular.UsesObservedYukawas || a.NonCircular.UsesHiggsMass {
		t.Fatalf("noncircularity failed: %s", FormatNonCircularity(a.NonCircular))
	}
	if a.Impact.CanPromoteTotalOperator || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("impact firewall failed: %s", FormatImpact(a.Impact))
	}
	res := Generation2BoundaryAlphaDomainTransportMapConstructionAndObstructionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
