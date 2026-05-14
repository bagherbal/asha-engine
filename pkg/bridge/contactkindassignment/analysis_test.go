package contactkindassignment

import (
	"math"
	"testing"

	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

func TestContactWeightTargetIsTwoHighTwoLow(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Target.ExactShape != "1197/4624" {
		t.Fatalf("unexpected scalar target: %+v", a.Target)
	}
	if a.Target.HighMultiplicity != 2 || a.Target.LowMultiplicity != 2 {
		t.Fatalf("expected two high and two low weights: %+v", a.Target)
	}
	if !(a.Target.HighWeightApprox > a.Target.LowWeightApprox) || math.Abs(a.Target.SquaredAmplitudeRatio-1.46404703870) > 1e-6 {
		t.Fatalf("bad high/low ratio: %+v", a.Target)
	}
	if !a.Target.RequiresKindAssignment || a.Target.UsesObservedInput {
		t.Fatalf("target should require kind assignment and no observed input: %+v", a.Target)
	}
}

func TestKindSignaturesMatchGate25Support(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.KindSignatures) != 4 {
		t.Fatalf("expected four kind signatures: %s", FormatKindSignatures(a.KindSignatures))
	}
	byKind := map[yukawaintertwiner.FermionKind]KindSignature{}
	for _, s := range a.KindSignatures {
		byKind[s.Kind] = s
	}
	if byKind[yukawaintertwiner.UpType].ChannelCount != 3 || byKind[yukawaintertwiner.DownType].ChannelCount != 3 {
		t.Fatalf("expected three color channels for up/down: %s", FormatKindSignatures(a.KindSignatures))
	}
	if byKind[yukawaintertwiner.NeutrinoType].ChannelCount != 1 || byKind[yukawaintertwiner.ElectronType].ChannelCount != 1 {
		t.Fatalf("expected one lepton channel for neutrino/electron: %s", FormatKindSignatures(a.KindSignatures))
	}
	if byKind[yukawaintertwiner.UpType].ScalarBranch != "Φ_+" || byKind[yukawaintertwiner.NeutrinoType].ScalarBranch != "Φ_+" {
		t.Fatalf("expected up/neutrino on Phi+ branch: %s", FormatKindSignatures(a.KindSignatures))
	}
	if byKind[yukawaintertwiner.DownType].ScalarBranch != "Φ_-" || byKind[yukawaintertwiner.ElectronType].ScalarBranch != "Φ_-" {
		t.Fatalf("expected down/electron on Phi- branch: %s", FormatKindSignatures(a.KindSignatures))
	}
}

func TestMultipleCanonicalPartitionsButNoContactAssignment(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.AssignmentAudit.OrientedHighLowAssignments != 6 || a.AssignmentAudit.ComplementUnorientedPartitions != 3 {
		t.Fatalf("bad assignment count: %+v", a.AssignmentAudit)
	}
	if a.AssignmentAudit.CanonicalPartitionsFound != 2 || !a.AssignmentAudit.MultipleIncompatiblePartitions {
		t.Fatalf("expected two incompatible canonical partitions: %+v partitions=%s", a.AssignmentAudit, FormatPartitions(a.Partitions))
	}
	if a.AssignmentAudit.CanonicalOrientedAssignmentsFound != 0 || a.AssignmentAudit.ContactTiedAssignmentsFound != 0 {
		t.Fatalf("no oriented/contact-tied assignment should be derived: %+v", a.AssignmentAudit)
	}
	if a.AssignmentAudit.UniqueContactKindAssignment || a.AssignmentAudit.SurvivingBranchChoices != 6 {
		t.Fatalf("all six assignments should remain branch choices: %+v", a.AssignmentAudit)
	}
}

func TestScalarShapeAndMassFirewallRemainOpen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Consequence.ConditionalShapeStillValid || a.Consequence.ContactKindAssignmentDerived || a.Consequence.ScalarShapeClosed {
		t.Fatalf("scalar shape should remain conditional: %+v", a.Consequence)
	}
	if !a.Firewall.GaugeRatioClosed || !a.Firewall.ScalarShapeTargetAvailable || !a.Firewall.FourKindSupportQuotientVisible {
		t.Fatalf("expected gauge/scalar/support prerequisites present: %+v", a.Firewall)
	}
	if a.Firewall.YukawaAmplitudesDerived || a.Firewall.GenerationTextureDerived || a.Firewall.FermionMassesDerived || a.Firewall.PhysicalConstantsDerived {
		t.Fatalf("mass/constant firewall should remain closed: %+v", a.Firewall)
	}
}
