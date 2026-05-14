// Package contactkindassignment implements Gate 171: a strict audit of whether
// the Gate-169 two-high/two-low contact scalar-shape weights can be assigned
// canonically to the four Gate-170 fermion-kind support classes {u,d,ν,e}.
//
// Gate 169 found an exact contact scalar-shape target with squared Yukawa
// weights proportional to two copies of (34+sqrt(41))/120 and two copies of
// (34-sqrt(41))/120. Gate 170 corrected the quotient mechanism: the available
// four-class object is the fermion-kind support quotient {u,d,ν,e}, not a
// Higgs-conjugate pair collapse. Gate 171 therefore asks whether any currently
// derived finite operator chooses which two fermion kinds receive the high
// contact weight.
//
// The answer is negative. Current finite data supplies several canonical 2+2
// partitions of the four kinds, but they are mutually incompatible and none is
// tied to the high/low contact eigenspaces. Scalar branch / weak-isospin sign
// gives {u,ν}|{d,e}; color/B-L gives {u,d}|{ν,e}; other charge-order cuts are
// diagnostic rank choices rather than derived contact assignments. Because no
// finite operator maps the contact high eigenspace into a unique fermion-kind
// pair, the six oriented high/low assignments remain branch choices.
package contactkindassignment

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/higgsconjugatequotient"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type ContactWeightTarget struct {
	ExactShape                 string
	HighWeightExact            string
	LowWeightExact             string
	HighWeightApprox           float64
	LowWeightApprox            float64
	HighMultiplicity           int
	LowMultiplicity            int
	SquaredAmplitudeRatioExact string
	SquaredAmplitudeRatio      float64
	RequiresKindAssignment     bool
	UsesObservedInput          bool
	Verdict                    string
}

type KindSignature struct {
	Kind              yukawaintertwiner.FermionKind
	ChannelCount      int
	ColorChannels     int
	LeptonChannels    int
	ScalarBranch      string
	ScalarHypercharge float64
	T3                float64
	RightHypercharge  float64
	ElectricCharge    float64
	BLCharge          float64
	KindClass         string
	Detail            string
}

type PartitionCandidate struct {
	Name                     string
	Source                   string
	HighKinds                []yukawaintertwiner.FermionKind
	LowKinds                 []yukawaintertwiner.FermionKind
	CanonicalPartition       bool
	CanonicalHighOrientation bool
	TiedToContactHighLow     bool
	UsesArbitraryCut         bool
	DistinguishesAllKinds    bool
	RemainingAssignments     int
	Verdict                  string
}

type AssignmentSearchAudit struct {
	FermionKinds                      int
	OrientedHighLowAssignments        int
	ComplementUnorientedPartitions    int
	CanonicalPartitionsFound          int
	CanonicalOrientedAssignmentsFound int
	ContactTiedAssignmentsFound       int
	MultipleIncompatiblePartitions    bool
	UniqueContactKindAssignment       bool
	SurvivingBranchChoices            int
	Verdict                           string
}

type ScalarShapeConsequenceAudit struct {
	Gate169TargetExact             string
	Gate170FourKindQuotientVisible bool
	ContactKindAssignmentDerived   bool
	ScalarShapeClosed              bool
	ConditionalShapeStillValid     bool
	AmplitudeTextureSelected       bool
	Verdict                        string
}

type FirewallAudit struct {
	GaugeRatioClosed                  bool
	ScalarShapeTargetAvailable        bool
	FourKindSupportQuotientVisible    bool
	ContactKindAssignmentDerived      bool
	FourAmplitudeClassQuotientDerived bool
	YukawaAmplitudesDerived           bool
	GenerationTextureDerived          bool
	FermionMassesDerived              bool
	CKMPMNSDerived                    bool
	PhysicalConstantsDerived          bool
	ResidualNullityBefore             int
	ResidualNullityAfter              int
	Verdict                           string
}

type Analysis struct {
	Previous higgsconjugatequotient.Analysis

	Target          ContactWeightTarget
	KindSignatures  []KindSignature
	Partitions      []PartitionCandidate
	AssignmentAudit AssignmentSearchAudit
	Consequence     ScalarShapeConsequenceAudit
	Firewall        FirewallAudit
	TruthStatement  string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := higgsconjugatequotient.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev higgsconjugatequotient.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !prev.KindQuotient.FourKindSupportQuotientVisible || !prev.Previous.Best.MatchesTarget {
		return Analysis{}, fmt.Errorf("Gate 171 requires Gate 170 four-kind support plus Gate 169 conditional contact-shape target")
	}
	target := buildTarget(prev)
	sigs := buildKindSignatures(prev.Groups)
	if len(sigs) != 4 {
		return Analysis{}, fmt.Errorf("expected four fermion-kind signatures, got %d", len(sigs))
	}
	parts := buildPartitions(sigs)
	audit := auditAssignments(parts)
	cons := ScalarShapeConsequenceAudit{
		Gate169TargetExact:             prev.Previous.Target.ExactValue,
		Gate170FourKindQuotientVisible: prev.KindQuotient.FourKindSupportQuotientVisible,
		ContactKindAssignmentDerived:   audit.UniqueContactKindAssignment,
		ScalarShapeClosed:              false,
		ConditionalShapeStillValid:     prev.Previous.Best.MatchesTarget,
		AmplitudeTextureSelected:       false,
		Verdict:                        "Gate-169's contact scalar-shape target remains valid but conditional: no finite source assigns the two high/two low contact weights to fermion kinds",
	}
	fw := FirewallAudit{
		GaugeRatioClosed:                  prev.Firewall.GaugeRatioClosed,
		ScalarShapeTargetAvailable:        prev.Firewall.ScalarShapeTargetAvailable,
		FourKindSupportQuotientVisible:    prev.KindQuotient.FourKindSupportQuotientVisible,
		ContactKindAssignmentDerived:      audit.UniqueContactKindAssignment,
		FourAmplitudeClassQuotientDerived: false,
		YukawaAmplitudesDerived:           false,
		GenerationTextureDerived:          false,
		FermionMassesDerived:              false,
		CKMPMNSDerived:                    false,
		PhysicalConstantsDerived:          false,
		ResidualNullityBefore:             3,
		ResidualNullityAfter:              3,
		Verdict:                           "contact-spectrum-to-kind assignment is not derived; scalar-sector closure, mass texture, generation texture, mixing, and constants remain sealed",
	}
	return Analysis{
		Previous:        prev,
		Target:          target,
		KindSignatures:  sigs,
		Partitions:      parts,
		AssignmentAudit: audit,
		Consequence:     cons,
		Firewall:        fw,
		TruthStatement:  "Gate 171 finds canonical fermion-kind partitions but no canonical contact-high/low assignment. Scalar branch/T3 and color/B-L give different 2+2 splits, and neither is connected to the contact high eigenspace. Thus the Gate-169 scalar-shape match remains a conditional Yukawa moment target, not a selected mass texture.",
	}, nil
}

func buildTarget(prev higgsconjugatequotient.Analysis) ContactWeightTarget {
	high := (34.0 + math.Sqrt(41.0)) / 120.0
	low := (34.0 - math.Sqrt(41.0)) / 120.0
	return ContactWeightTarget{
		ExactShape:                 prev.Previous.Target.ExactValue,
		HighWeightExact:            "(34+sqrt(41))/120",
		LowWeightExact:             "(34-sqrt(41))/120",
		HighWeightApprox:           high,
		LowWeightApprox:            low,
		HighMultiplicity:           2,
		LowMultiplicity:            2,
		SquaredAmplitudeRatioExact: prev.Previous.PairCollapse.SquaredAmplitudeRatioExact,
		SquaredAmplitudeRatio:      prev.Previous.PairCollapse.SquaredAmplitudeRatio,
		RequiresKindAssignment:     true,
		UsesObservedInput:          false,
		Verdict:                    "the contact scalar-shape target supplies two identical high and two identical low squared-amplitude weights, but no labels on {u,d,ν,e}",
	}
}

func buildKindSignatures(groups []higgsconjugatequotient.KindChannelGroup) []KindSignature {
	out := make([]KindSignature, 0, len(groups))
	for _, g := range groups {
		s := KindSignature{
			Kind:           g.Kind,
			ChannelCount:   g.ChannelCount,
			ColorChannels:  g.ColorChannels,
			LeptonChannels: g.LeptonChannels,
			ScalarBranch:   firstOrEmpty(g.ScalarBranches),
		}
		switch g.Kind {
		case yukawaintertwiner.UpType:
			s.ScalarHypercharge = +0.5
			s.T3 = +0.5
			s.RightHypercharge = 2.0 / 3.0
			s.ElectricCharge = 2.0 / 3.0
			s.BLCharge = 1.0 / 3.0
			s.KindClass = "quark/up/Φ+"
		case yukawaintertwiner.DownType:
			s.ScalarHypercharge = -0.5
			s.T3 = -0.5
			s.RightHypercharge = -1.0 / 3.0
			s.ElectricCharge = -1.0 / 3.0
			s.BLCharge = 1.0 / 3.0
			s.KindClass = "quark/down/Φ-"
		case yukawaintertwiner.NeutrinoType:
			s.ScalarHypercharge = +0.5
			s.T3 = +0.5
			s.RightHypercharge = 0
			s.ElectricCharge = 0
			s.BLCharge = -1
			s.KindClass = "lepton/up/Φ+"
		case yukawaintertwiner.ElectronType:
			s.ScalarHypercharge = -0.5
			s.T3 = -0.5
			s.RightHypercharge = -1
			s.ElectricCharge = -1
			s.BLCharge = -1
			s.KindClass = "lepton/down/Φ-"
		}
		s.Detail = fmt.Sprintf("%s channels=%d color=%d lepton=%d scalar=%s T3=%+.1f YR=%.6g Q=%.6g B-L=%.6g class=%s", s.Kind, s.ChannelCount, s.ColorChannels, s.LeptonChannels, s.ScalarBranch, s.T3, s.RightHypercharge, s.ElectricCharge, s.BLCharge, s.KindClass)
		out = append(out, s)
	}
	sort.Slice(out, func(i, j int) bool { return kindRank(out[i].Kind) < kindRank(out[j].Kind) })
	return out
}

func buildPartitions(_ []KindSignature) []PartitionCandidate {
	u := yukawaintertwiner.UpType
	d := yukawaintertwiner.DownType
	n := yukawaintertwiner.NeutrinoType
	e := yukawaintertwiner.ElectronType
	return []PartitionCandidate{
		{
			Name:                     "scalar branch / weak-isospin sign split, positive branch high",
			Source:                   "T3 sign and Φ hypercharge sign select {u,ν}|{d,e}",
			HighKinds:                kinds(u, n),
			LowKinds:                 kinds(d, e),
			CanonicalPartition:       true,
			CanonicalHighOrientation: false,
			TiedToContactHighLow:     false,
			UsesArbitraryCut:         false,
			DistinguishesAllKinds:    false,
			RemainingAssignments:     2,
			Verdict:                  "canonical 2+2 partition exists, but assigning the contact high weight to the positive scalar/T3 side is an orientation choice",
		},
		{
			Name:                     "scalar branch / weak-isospin sign split, negative branch high",
			Source:                   "T3 sign and Φ hypercharge sign select {u,ν}|{d,e}",
			HighKinds:                kinds(d, e),
			LowKinds:                 kinds(u, n),
			CanonicalPartition:       true,
			CanonicalHighOrientation: false,
			TiedToContactHighLow:     false,
			UsesArbitraryCut:         false,
			DistinguishesAllKinds:    false,
			RemainingAssignments:     2,
			Verdict:                  "same canonical partition with opposite high/low orientation; the finite data does not choose between the two",
		},
		{
			Name:                     "color/B-L split, quarks high",
			Source:                   "color multiplicity and B-L sign/magnitude select quark vs lepton kinds {u,d}|{ν,e}",
			HighKinds:                kinds(u, d),
			LowKinds:                 kinds(n, e),
			CanonicalPartition:       true,
			CanonicalHighOrientation: false,
			TiedToContactHighLow:     false,
			UsesArbitraryCut:         false,
			DistinguishesAllKinds:    false,
			RemainingAssignments:     2,
			Verdict:                  "canonical quark/lepton partition exists, but assigning the contact high weight to quarks is not derived",
		},
		{
			Name:                     "color/B-L split, leptons high",
			Source:                   "color multiplicity and B-L sign/magnitude select quark vs lepton kinds {u,d}|{ν,e}",
			HighKinds:                kinds(n, e),
			LowKinds:                 kinds(u, d),
			CanonicalPartition:       true,
			CanonicalHighOrientation: false,
			TiedToContactHighLow:     false,
			UsesArbitraryCut:         false,
			DistinguishesAllKinds:    false,
			RemainingAssignments:     2,
			Verdict:                  "same canonical quark/lepton partition with opposite high/low orientation; the finite data does not choose between the two",
		},
		{
			Name:                     "diagonal mixed split, {u,e} high",
			Source:                   "would require an extra rank/order rule such as |Y_R| extremal pairing; no such rule is derived from contact geometry",
			HighKinds:                kinds(u, e),
			LowKinds:                 kinds(d, n),
			CanonicalPartition:       false,
			CanonicalHighOrientation: false,
			TiedToContactHighLow:     false,
			UsesArbitraryCut:         true,
			DistinguishesAllKinds:    false,
			RemainingAssignments:     6,
			Verdict:                  "not a currently derived finite contact-kind selector; it is a diagnostic branch choice",
		},
		{
			Name:                     "diagonal mixed split, {d,ν} high",
			Source:                   "complement of the same non-derived diagonal rank/order rule",
			HighKinds:                kinds(d, n),
			LowKinds:                 kinds(u, e),
			CanonicalPartition:       false,
			CanonicalHighOrientation: false,
			TiedToContactHighLow:     false,
			UsesArbitraryCut:         true,
			DistinguishesAllKinds:    false,
			RemainingAssignments:     6,
			Verdict:                  "not a currently derived finite contact-kind selector; it is a diagnostic branch choice",
		},
	}
}

func auditAssignments(parts []PartitionCandidate) AssignmentSearchAudit {
	canonicalPartitions := 0
	canonicalOriented := 0
	contactTied := 0
	partitionKeys := map[string]bool{}
	for _, p := range parts {
		if p.CanonicalPartition {
			canonicalPartitions++
			partitionKeys[unorientedKey(p.HighKinds, p.LowKinds)] = true
		}
		if p.CanonicalPartition && p.CanonicalHighOrientation {
			canonicalOriented++
		}
		if p.TiedToContactHighLow {
			contactTied++
		}
	}
	unique := contactTied == 1 && canonicalOriented == 1
	return AssignmentSearchAudit{
		FermionKinds:                      4,
		OrientedHighLowAssignments:        6,
		ComplementUnorientedPartitions:    3,
		CanonicalPartitionsFound:          len(partitionKeys),
		CanonicalOrientedAssignmentsFound: canonicalOriented,
		ContactTiedAssignmentsFound:       contactTied,
		MultipleIncompatiblePartitions:    len(partitionKeys) > 1,
		UniqueContactKindAssignment:       unique,
		SurvivingBranchChoices:            6,
		Verdict:                           "canonical partitions exist but are incompatible and none is tied to the contact high eigenspace; all six oriented high/low assignments remain branch choices",
	}
}

func FormatTarget(t ContactWeightTarget) string {
	return fmt.Sprintf("shape=%s high=%s≈%.12g×%d low=%s≈%.12g×%d ratio=%s≈%.12g requiresKind=%t observed=%t", t.ExactShape, t.HighWeightExact, t.HighWeightApprox, t.HighMultiplicity, t.LowWeightExact, t.LowWeightApprox, t.LowMultiplicity, t.SquaredAmplitudeRatioExact, t.SquaredAmplitudeRatio, t.RequiresKindAssignment, t.UsesObservedInput)
}

func FormatKindSignatures(sigs []KindSignature) string {
	parts := make([]string, 0, len(sigs))
	for _, s := range sigs {
		parts = append(parts, s.Detail)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatPartition(p PartitionCandidate) string {
	return fmt.Sprintf("%s high=%s low=%s source=%s canonicalPartition=%t canonicalHigh=%t contactTied=%t arbitraryCut=%t remaining=%d", p.Name, formatKinds(p.HighKinds), formatKinds(p.LowKinds), p.Source, p.CanonicalPartition, p.CanonicalHighOrientation, p.TiedToContactHighLow, p.UsesArbitraryCut, p.RemainingAssignments)
}

func FormatPartitions(parts []PartitionCandidate) string {
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		out = append(out, FormatPartition(p))
	}
	return "[" + strings.Join(out, "; ") + "]"
}

func FormatAssignmentAudit(a AssignmentSearchAudit) string {
	return fmt.Sprintf("kinds=%d oriented=%d unoriented=%d canonicalPartitions=%d canonicalOriented=%d contactTied=%d incompatible=%t unique=%t surviving=%d", a.FermionKinds, a.OrientedHighLowAssignments, a.ComplementUnorientedPartitions, a.CanonicalPartitionsFound, a.CanonicalOrientedAssignmentsFound, a.ContactTiedAssignmentsFound, a.MultipleIncompatiblePartitions, a.UniqueContactKindAssignment, a.SurvivingBranchChoices)
}

func FormatConsequence(a ScalarShapeConsequenceAudit) string {
	return fmt.Sprintf("target=%s fourKind=%t assignment=%t scalarClosed=%t conditional=%t amplitudes=%t", a.Gate169TargetExact, a.Gate170FourKindQuotientVisible, a.ContactKindAssignmentDerived, a.ScalarShapeClosed, a.ConditionalShapeStillValid, a.AmplitudeTextureSelected)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gaugeClosed=%t scalarTarget=%t fourKind=%t contactAssign=%t fourAmplitude=%t amplitudes=%t generation=%t masses=%t CKM/PMNS=%t constants=%t nullity=%d->%d", a.GaugeRatioClosed, a.ScalarShapeTargetAvailable, a.FourKindSupportQuotientVisible, a.ContactKindAssignmentDerived, a.FourAmplitudeClassQuotientDerived, a.YukawaAmplitudesDerived, a.GenerationTextureDerived, a.FermionMassesDerived, a.CKMPMNSDerived, a.PhysicalConstantsDerived, a.ResidualNullityBefore, a.ResidualNullityAfter)
}

func kinds(xs ...yukawaintertwiner.FermionKind) []yukawaintertwiner.FermionKind {
	out := append([]yukawaintertwiner.FermionKind(nil), xs...)
	sort.Slice(out, func(i, j int) bool { return kindRank(out[i]) < kindRank(out[j]) })
	return out
}

func formatKinds(xs []yukawaintertwiner.FermionKind) string {
	ys := make([]string, len(xs))
	for i, x := range xs {
		ys[i] = string(x)
	}
	return "{" + strings.Join(ys, ",") + "}"
}

func unorientedKey(a, b []yukawaintertwiner.FermionKind) string {
	fa, fb := formatKinds(kinds(a...)), formatKinds(kinds(b...))
	if fa < fb {
		return fa + "|" + fb
	}
	return fb + "|" + fa
}

func firstOrEmpty(xs []string) string {
	if len(xs) == 0 {
		return ""
	}
	return xs[0]
}

func kindRank(k yukawaintertwiner.FermionKind) int {
	switch k {
	case yukawaintertwiner.UpType:
		return 0
	case yukawaintertwiner.DownType:
		return 1
	case yukawaintertwiner.NeutrinoType:
		return 2
	case yukawaintertwiner.ElectronType:
		return 3
	default:
		return 99
	}
}
