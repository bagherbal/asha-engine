package scalarorientationseal

import (
	"testing"

	"github.com/bagherbal/asha-engine/pkg/bridge/betamatching"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcomplex"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcovariant"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarorientationsource"
	"github.com/bagherbal/asha-engine/pkg/bridge/topologicalnormalization"
)

func TestBuildDefaultSpontaneousScalarOrientationSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Seal.ExplicitAxiom || !a.Seal.Quarantined || a.Seal.DerivedFromFiniteSelector || a.Seal.UsesObservedInput {
		t.Fatalf("seal is not correctly quarantined: %s", FormatSeal(a.Seal))
	}
	if !a.Trivialization.ProjectorIntertwiningVerified || !a.Trivialization.PhysicalScalarBundleTrivialized {
		t.Fatalf("sealed trivialization failed: %s", FormatTrivialization(a.Trivialization))
	}
	if !a.GaugePullback.T3LPreservesFibers || !a.GaugePullback.YPhiPreservesFibers || !a.GaugePullback.T1MixesFibers || !a.GaugePullback.T2MixesFibers {
		t.Fatalf("gauge pullback did not have expected block/off-block structure: %s", FormatGaugePullback(a.GaugePullback))
	}
	if a.Firewall.ChernWeilCarrierDerived || a.Firewall.HeatKernelMatchingDerived || a.Firewall.ThresholdCorrectedBetaDerived || a.Firewall.AbsoluteCouplingPromoted || a.Firewall.PhysicalConstantsDerived {
		t.Fatalf("firewall leaked: %s", FormatFirewall(a.Firewall))
	}
}

func TestEtaToLowAlternativeSealAlsoTrivializes(t *testing.T) {
	prev, err := scalarorientationsource.BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	sc, err := scalarcovariant.BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	cx, err := scalarcomplex.BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	top, err := topologicalnormalization.BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	beta, err := betamatching.BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	a, err := Build(prev, sc, cx, top, beta, EtaToLow, 1e-9)
	if err != nil {
		t.Fatalf("Build eta_to_low failed: %v", err)
	}
	if !a.SealedFrame.SealMapsAToLow || !a.SealedFrame.SealMapsBToHigh {
		t.Fatalf("eta_to_low seal not recorded: %s", FormatSealedFrame(a.SealedFrame))
	}
	if !a.Trivialization.ProjectorIntertwiningVerified || !a.Firewall.ConditionalPhysicalBundleDerived {
		t.Fatalf("eta_to_low trivialization failed: %s", FormatTrivialization(a.Trivialization))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := SpontaneousScalarOrientationSealGaugeFixedHphiTrivializationTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
