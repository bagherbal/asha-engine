package generation2leptocolorfockcarrierrepresentationsealaudit

import (
	"strings"
	"testing"
)

func TestGate837SharedLeptoColorCarrierAndBMinusL(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Carrier.Dim != 4 || a.Carrier.P1Rank != 1 || a.Carrier.P3Rank != 3 || !a.Carrier.P1P3Orthogonal || !a.Carrier.P1PlusP3CompletesW || !a.Carrier.BMinusLTraceZero {
		t.Fatalf("bad lepto-color carrier: %s", FormatCarrier(a.Carrier))
	}
	if !containsAll(a.Carrier.Supports, []string{SupportSharedWUnifiesFockOnePlusThreeAndM3ColorModule, SupportP1P3SourceBMinusLInternally}) {
		t.Fatalf("missing carrier supports: %s", strings.Join(a.Carrier.Supports, ","))
	}
	if !containsAll(a.Carrier.Failures, []string{FailureSealNotNativeDerivation}) {
		t.Fatalf("missing seal/firewall failure: %s", strings.Join(a.Carrier.Failures, ","))
	}
}

func TestGate837M3ActsOnP3WButColorAtomsRemainFrameDependent(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.M3.ActsOnP3WBySealDefinition || !a.M3.P3WIsM3Fundamental || a.M3.M3IdentityTrace != 3 || !a.M3.MatrixUnitsExist || !a.M3.MatrixUnitsActWithinP3W || !a.M3.P1InvariantUnderM3 || !a.M3.BlockLevelCanonical || a.M3.IndividualColorAtomsCanonical || a.M3.CanonicalColorFrameCertified || !a.M3.NoSeparateTripletBridgeNeeded || a.M3.ContradictsGate833 {
		t.Fatalf("bad M3/P3 seal classification: %s", FormatM3(a.M3))
	}
	if !containsAll(a.M3.Failures, []string{FailureNoCanonicalColorAtomFrame, FailureM3MatrixUnitsBasisDependent}) {
		t.Fatalf("missing color-frame failures: %s", strings.Join(a.M3.Failures, ","))
	}
}

func TestGate837FiniteBodySkeletonButNotCompleteTriple(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Body.HPartDim != 16 || a.Body.HFDim != 32 || a.Body.RightSlotDim != 2 || a.Body.LeftSlotDim != 2 || a.Body.WDim != 4 || !a.Body.RhoFRoleDeclared || !a.Body.GammaFRoleDeclared || !a.Body.JFRoleDeclared || !a.Body.DFRoleDeclared {
		t.Fatalf("bad finite body skeleton: %s", FormatBody(a.Body))
	}
	if a.Body.CompleteRhoFActionLedger || a.Body.ExplicitGammaFOperator || a.Body.ExplicitJFOperator || a.Body.SymbolicDFEdgeMatrix || a.Body.ObservedDataUsed {
		t.Fatalf("finite body over-certified: %s", FormatBody(a.Body))
	}
	if !containsAll(a.Body.Failures, []string{FailureNoCompleteRhoFActionLedger, FailureNoCompleteFiniteTripleData, FailureNoGammaFOperatorMatrices, FailureNoJFOperatorMatrices, FailureNoDFSymbolicEdgeMatrix, FailureDFSocketsNotYukawaMagnitudes}) {
		t.Fatalf("missing body failures: %s", strings.Join(a.Body.Failures, ","))
	}
}

func TestGate837ProjectorsCompressionImpactAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Projectors.CanConstructBlockProjectors || !a.Projectors.CanConstructChiralityTimesLeptoColorSupports || !a.Projectors.CanConstructOppositeCopySupports || a.Projectors.PiSectorFCertified || a.Projectors.TraceMagnitudeReadoutCertified || a.Projectors.SectorProjectorsAreMagnitudes {
		t.Fatalf("projector prospect over/under-certified: %s", FormatProjectors(a.Projectors))
	}
	if !a.Compression.DirectionCorrected || !a.Compression.FiniteBodyToAggregateShadow || a.Compression.AggregateToFiniteBody || a.Compression.R2OperatorSectorLedger || a.Compression.CompressionMapCertified || a.Compression.SigmaPullbackCertified {
		t.Fatalf("compression classification failed: %s", FormatCompression(a.Compression))
	}
	if !a.Impact.CarrierSealConstructed || !a.Impact.CarrierProblemSolvedAtBlockLevel || a.Impact.CompleteFiniteTripleData || a.Impact.PiSectorFCertified || a.Impact.CompressionMapCertified || a.Impact.TraceMagnitudeReadoutCertified || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("impact over-promoted: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.CarrierSealNotNativeDerivation || !a.Firewalls.NoCanonicalColorAtoms || !a.Firewalls.NoPiSectorF || !a.Firewalls.NoCompressionMap || !a.Firewalls.NoMagnitudeReadout || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoCYukawaUpdate || a.Firewalls.Verdict != StatusFirewallGate837 {
		t.Fatalf("firewall failed: %+v", a.Firewalls)
	}
	res := Generation2LeptoColorFockCarrierRepresentationSealAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem construction failure: %+v", res)
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
