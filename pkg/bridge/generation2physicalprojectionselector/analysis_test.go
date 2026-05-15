package generation2physicalprojectionselector

import (
	"strings"
	"testing"
)

func TestGate528Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate527Inherited || !a.Inheritance.Gate527ProjectionAirlockDefined || !a.Inheritance.Gate527Physical3Plus1Blocked || !a.Inheritance.Gate527NoObservedDataImported || !a.Inheritance.Gate527NativeWriteBlocked || a.Inheritance.Gate527ReopenedSealedFirewalls {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Idempotents.VolumeElementAvailable || !a.Idempotents.ChiralityProjectorsAvailable || !a.Idempotents.ChiralityProjectorsIdempotent || !a.Idempotents.ChiralityActsOnSpinorParity || a.Idempotents.ChiralityProjectsVectorSpace44 || !a.Idempotents.PrimitiveIdempotentsAbundant || a.Idempotents.PrimitiveIdempotentsCanonical {
		t.Fatalf("bad idempotent sieve: %+v", a.Idempotents)
	}
	if a.Rank44.VectorDimension != 8 || a.Rank44.CandidateExternalRank != 4 || a.Rank44.CandidateInternalRank != 4 || !a.Rank44.RankArithmeticValid || !a.Rank44.ChosenFourPlaneProjectorIdempotent || !a.Rank44.ProjectorComplementary || !a.Rank44.ProjectorRequiresBasisChoice || a.Rank44.Spin17InvariantRank4ProjectorFound || a.Rank44.MutuallyCommutingSubalgebrasNative || !a.Rank44.GradedTensorFactorizationBridgeOnly || a.Rank44.InternalComplementUniqueNative {
		t.Fatalf("bad 4+4 audit: %+v", a.Rank44)
	}
	if a.Selector.ExternalLorentzSignatureCandidate != "1+3" || !a.Selector.TimeLikeDirectionAvailable || !a.Selector.TimeIncludedByChosenBridgePlane || a.Selector.TimeAssignmentNativeSelected || a.Selector.OrientationAndArrowSelected || a.Selector.Physical3Plus1ProjectorIdentified || !a.Selector.Physical3Plus1BridgeSocketReady || a.Selector.InternalGaugeSpaceIdentified {
		t.Fatalf("bad selector: %+v", a.Selector)
	}
	if a.Firewall.ObservedDimensionImported || a.Firewall.ObservedConstantsImported || a.Firewall.ObservedMassesImported || a.Firewall.ObservedTopologyImported || a.Firewall.NativeChiralityVectorWrite || a.Firewall.NativeFourPlaneWrite || a.Firewall.NativeInternalComplementWrite || a.Firewall.NativeTimeAssignmentWrite || a.Firewall.Native3Plus1ProjectionWrite || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
}

func TestGate528Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 528 Registry Audit", StatusChiralityVolumeProjectorSocketFound, StatusFailedNoSpin17InvariantRank4VectorProjector, StatusFailedPhysical3Plus1ProjectorNotIdentified, StatusFirewallProjectionNativeWriteBlocked, "Gate 529"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate528Theorem(t *testing.T) {
	result := Generation2Physical3Plus1ProjectionAndInternalComplementSelectorAuditTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
