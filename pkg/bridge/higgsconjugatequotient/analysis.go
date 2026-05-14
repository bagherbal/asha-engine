// Package higgsconjugatequotient implements Gate 170: a strict audit of the
// proposed Higgs-conjugate reduction of the eight Gate-25 Yukawa channels.
//
// Gate 169 found a conditional scalar-shape match if the eight finite Yukawa
// support slots could be quotiented into four amplitude classes and those four
// squared amplitudes were identified with the four active contact/Higgs weights.
// Gate 170 checks the premise against the actual Gate-25 channel table.
//
// The result is a refinement rather than a scalar-sector closure.  The eight
// minimal Gate-25 channels are not two Higgs-conjugate copies of four fermion
// kinds.  Hypercharge selects a unique scalar branch for each kind: up and
// neutrino use Φ_+, down and electron use Φ_-.  The eight support slots are
// instead 3 up-color channels + 3 down-color channels + one neutrino + one
// electron.  Thus a four-kind support quotient is visible, but it is a
// color/kind quotient, not a Higgs-conjugate pair collapse, and it still does
// not select the two-high/two-low contact assignment or the Yukawa amplitudes.
package higgsconjugatequotient

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/yukawashapeconstraint"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type KindChannelGroup struct {
	Kind                          yukawaintertwiner.FermionKind
	ChannelCount                  int
	ColorChannels                 int
	LeptonChannels                int
	ScalarBranches                []string
	UniqueScalarBranch            bool
	ContainsBothConjugateBranches bool
	HyperchargeSelectedBranch     bool
	SupportQuotientClass          bool
	AmplitudeEqualityDerived      bool
	Detail                        string
}

type HiggsConjugateAudit struct {
	Gate25MinimalChannels          int
	ScalarBranches                 int
	FermionKindBlocks              int
	KindsWithBothBranches          int
	KindsWithUniqueBranch          int
	HyperchargeSelectsUniqueBranch bool
	HiggsConjugatePairsAvailable   bool
	HiggsConjugatePairCollapse     bool
	RejectedBecause                string
	Verdict                        string
}

type ColorKindQuotientAudit struct {
	KindBlocks                        int
	SupportPattern                    string
	UpColorChannels                   int
	DownColorChannels                 int
	NeutrinoChannels                  int
	ElectronChannels                  int
	FourKindSupportQuotientVisible    bool
	FourKindSupportQuotientCanonical  bool
	ColorAmplitudeUniversalityDerived bool
	FourAmplitudeClassQuotientDerived bool
	Verdict                           string
}

type ScalarShapeConsequenceAudit struct {
	Gate169TargetExact             string
	Gate169ConditionalMatchFound   bool
	HiggsConjugatePremiseRejected  bool
	FourKindQuotientStillAvailable bool
	ContactWeightAssignments       int
	CanonicalContactKindAssignment bool
	ScalarShapeClosed              bool
	AmplitudeTextureSelected       bool
	Verdict                        string
}

type FirewallAudit struct {
	GaugeRatioClosed                  bool
	ScalarShapeTargetAvailable        bool
	HiggsConjugateQuotientDerived     bool
	FourKindSupportQuotientVisible    bool
	FourAmplitudeClassQuotientDerived bool
	ContactKindAssignmentDerived      bool
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
	Previous yukawashapeconstraint.Analysis

	Groups         []KindChannelGroup
	HiggsAudit     HiggsConjugateAudit
	KindQuotient   ColorKindQuotientAudit
	Consequence    ScalarShapeConsequenceAudit
	Firewall       FirewallAudit
	TruthStatement string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := yukawashapeconstraint.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev yukawashapeconstraint.Analysis) (Analysis, error) {
	y := prev.Previous.Previous.Yukawa
	if !y.ChargeCompatibleYukawaChannelsDerived || len(y.Channels) != 8 {
		return Analysis{}, fmt.Errorf("Gate 170 requires the Gate-25 eight-channel Yukawa support, got %d", len(y.Channels))
	}
	groups := buildKindGroups(y.Channels)
	higgs := auditHiggsConjugation(y, groups)
	kind := auditColorKindQuotient(groups)
	cons := ScalarShapeConsequenceAudit{
		Gate169TargetExact:             prev.Target.ExactValue,
		Gate169ConditionalMatchFound:   prev.Best.MatchesTarget,
		HiggsConjugatePremiseRejected:  !higgs.HiggsConjugatePairCollapse,
		FourKindQuotientStillAvailable: kind.FourKindSupportQuotientVisible,
		ContactWeightAssignments:       prev.PairCollapse.KindAssignmentAmbiguity,
		CanonicalContactKindAssignment: false,
		ScalarShapeClosed:              false,
		AmplitudeTextureSelected:       false,
		Verdict:                        "Gate 169's scalar target survives as a finite moment constraint, but the proposed Higgs-conjugate 8→4 collapse is not the mechanism; the remaining obstruction is the canonical assignment of the two high/two low contact weights to fermion kinds plus actual amplitude selection",
	}
	fw := FirewallAudit{
		GaugeRatioClosed:                  prev.Firewall.GaugeRatioClosed,
		ScalarShapeTargetAvailable:        prev.Firewall.ScalarShapeTargetAvailable,
		HiggsConjugateQuotientDerived:     higgs.HiggsConjugatePairCollapse,
		FourKindSupportQuotientVisible:    kind.FourKindSupportQuotientVisible,
		FourAmplitudeClassQuotientDerived: kind.FourAmplitudeClassQuotientDerived,
		ContactKindAssignmentDerived:      cons.CanonicalContactKindAssignment,
		YukawaAmplitudesDerived:           false,
		GenerationTextureDerived:          false,
		FermionMassesDerived:              false,
		CKMPMNSDerived:                    false,
		PhysicalConstantsDerived:          false,
		ResidualNullityBefore:             3,
		ResidualNullityAfter:              3,
		Verdict:                           "the Higgs-conjugate quotient is rejected; a four-kind support quotient is visible but does not derive amplitude equality, contact-weight assignment, masses, mixing, or physical constants",
	}
	return Analysis{
		Previous:       prev,
		Groups:         groups,
		HiggsAudit:     higgs,
		KindQuotient:   kind,
		Consequence:    cons,
		Firewall:       fw,
		TruthStatement: "Gate 170 corrects the Gate-169 quotient premise: the eight Yukawa support channels are not scalar-conjugate pairs. Hypercharge gives each fermion kind exactly one Higgs branch. The natural 8→4 compression is a fermion-kind/color support quotient, and it leaves the contact-weight assignment and Yukawa amplitude problem open.",
	}, nil
}

func buildKindGroups(channels []yukawaintertwiner.Channel) []KindChannelGroup {
	byKind := map[yukawaintertwiner.FermionKind][]yukawaintertwiner.Channel{}
	for _, ch := range channels {
		byKind[ch.Right.Kind] = append(byKind[ch.Right.Kind], ch)
	}
	kinds := []yukawaintertwiner.FermionKind{
		yukawaintertwiner.UpType,
		yukawaintertwiner.DownType,
		yukawaintertwiner.NeutrinoType,
		yukawaintertwiner.ElectronType,
	}
	out := make([]KindChannelGroup, 0, len(kinds))
	for _, kind := range kinds {
		chs := byKind[kind]
		branches := distinctBranches(chs)
		color, lepton := 0, 0
		for _, ch := range chs {
			if ch.Right.Color > 0 {
				color++
			} else {
				lepton++
			}
		}
		unique := len(branches) == 1
		both := contains(branches, "Φ_+") && contains(branches, "Φ_-")
		out = append(out, KindChannelGroup{
			Kind:                          kind,
			ChannelCount:                  len(chs),
			ColorChannels:                 color,
			LeptonChannels:                lepton,
			ScalarBranches:                branches,
			UniqueScalarBranch:            unique,
			ContainsBothConjugateBranches: both,
			HyperchargeSelectedBranch:     unique && !both,
			SupportQuotientClass:          len(chs) > 0,
			AmplitudeEqualityDerived:      false,
			Detail:                        fmt.Sprintf("%s channels=%d colors=%d leptons=%d branches=%s", kind, len(chs), color, lepton, strings.Join(branches, ",")),
		})
	}
	return out
}

func distinctBranches(chs []yukawaintertwiner.Channel) []string {
	m := map[string]bool{}
	for _, ch := range chs {
		m[ch.Scalar.Name] = true
	}
	out := make([]string, 0, len(m))
	for b := range m {
		out = append(out, b)
	}
	sort.Strings(out)
	return out
}

func auditHiggsConjugation(y yukawaintertwiner.Analysis, groups []KindChannelGroup) HiggsConjugateAudit {
	both, unique := 0, 0
	for _, g := range groups {
		if g.ContainsBothConjugateBranches {
			both++
		}
		if g.UniqueScalarBranch {
			unique++
		}
	}
	return HiggsConjugateAudit{
		Gate25MinimalChannels:          len(y.Channels),
		ScalarBranches:                 len(y.ScalarBranches),
		FermionKindBlocks:              len(groups),
		KindsWithBothBranches:          both,
		KindsWithUniqueBranch:          unique,
		HyperchargeSelectsUniqueBranch: unique == len(groups),
		HiggsConjugatePairsAvailable:   both == len(groups),
		HiggsConjugatePairCollapse:     false,
		RejectedBecause:                "Gate-25 hypercharge balance Y_R=Y_L+Y_Φ selects Φ_+ for up/neutrino and Φ_- for down/electron; no fermion kind carries both conjugate scalar branches",
		Verdict:                        "no canonical Higgs-conjugate 8→4 channel quotient exists on the actual Gate-25 support",
	}
}

func auditColorKindQuotient(groups []KindChannelGroup) ColorKindQuotientAudit {
	up, down, nu, e := 0, 0, 0, 0
	for _, g := range groups {
		switch g.Kind {
		case yukawaintertwiner.UpType:
			up = g.ChannelCount
		case yukawaintertwiner.DownType:
			down = g.ChannelCount
		case yukawaintertwiner.NeutrinoType:
			nu = g.ChannelCount
		case yukawaintertwiner.ElectronType:
			e = g.ChannelCount
		}
	}
	visible := len(groups) == 4 && up == 3 && down == 3 && nu == 1 && e == 1
	return ColorKindQuotientAudit{
		KindBlocks:                        len(groups),
		SupportPattern:                    "3_u + 3_d + 1_ν + 1_e → {u,d,ν,e}",
		UpColorChannels:                   up,
		DownColorChannels:                 down,
		NeutrinoChannels:                  nu,
		ElectronChannels:                  e,
		FourKindSupportQuotientVisible:    visible,
		FourKindSupportQuotientCanonical:  visible,
		ColorAmplitudeUniversalityDerived: false,
		FourAmplitudeClassQuotientDerived: false,
		Verdict:                           "the support table canonically groups into four fermion kinds, but equality of color amplitudes and the contact-spectrum weights are not derived by the Higgs-conjugate quotient audit",
	}
}

func FormatGroups(gs []KindChannelGroup) string {
	parts := make([]string, 0, len(gs))
	for _, g := range gs {
		parts = append(parts, fmt.Sprintf("%s:n=%d color=%d lepton=%d branches=%s unique=%t both=%t ampEq=%t", g.Kind, g.ChannelCount, g.ColorChannels, g.LeptonChannels, strings.Join(g.ScalarBranches, "/"), g.UniqueScalarBranch, g.ContainsBothConjugateBranches, g.AmplitudeEqualityDerived))
	}
	sort.Strings(parts)
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatHiggsAudit(a HiggsConjugateAudit) string {
	return fmt.Sprintf("channels=%d scalarBranches=%d kinds=%d uniqueBranchKinds=%d bothBranchKinds=%d hyperchargeUnique=%t conjugatePairs=%t collapse=%t reason=%s", a.Gate25MinimalChannels, a.ScalarBranches, a.FermionKindBlocks, a.KindsWithUniqueBranch, a.KindsWithBothBranches, a.HyperchargeSelectsUniqueBranch, a.HiggsConjugatePairsAvailable, a.HiggsConjugatePairCollapse, a.RejectedBecause)
}

func FormatKindQuotient(a ColorKindQuotientAudit) string {
	return fmt.Sprintf("kindBlocks=%d pattern=%s up=%d down=%d nu=%d e=%d supportVisible=%t supportCanonical=%t colorAmp=%t amplitudeQuotient=%t", a.KindBlocks, a.SupportPattern, a.UpColorChannels, a.DownColorChannels, a.NeutrinoChannels, a.ElectronChannels, a.FourKindSupportQuotientVisible, a.FourKindSupportQuotientCanonical, a.ColorAmplitudeUniversalityDerived, a.FourAmplitudeClassQuotientDerived)
}

func FormatConsequence(a ScalarShapeConsequenceAudit) string {
	return fmt.Sprintf("target=%s conditional169=%t higgsPremiseRejected=%t fourKindVisible=%t assignments=%d canonicalAssignment=%t scalarClosed=%t amplitudes=%t", a.Gate169TargetExact, a.Gate169ConditionalMatchFound, a.HiggsConjugatePremiseRejected, a.FourKindQuotientStillAvailable, a.ContactWeightAssignments, a.CanonicalContactKindAssignment, a.ScalarShapeClosed, a.AmplitudeTextureSelected)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gaugeClosed=%t scalarTarget=%t higgsQuotient=%t fourKindSupport=%t fourAmplitude=%t contactKindAssign=%t amplitudes=%t generation=%t masses=%t CKM/PMNS=%t constants=%t nullity=%d->%d", a.GaugeRatioClosed, a.ScalarShapeTargetAvailable, a.HiggsConjugateQuotientDerived, a.FourKindSupportQuotientVisible, a.FourAmplitudeClassQuotientDerived, a.ContactKindAssignmentDerived, a.YukawaAmplitudesDerived, a.GenerationTextureDerived, a.FermionMassesDerived, a.CKMPMNSDerived, a.PhysicalConstantsDerived, a.ResidualNullityBefore, a.ResidualNullityAfter)
}

func contains(xs []string, v string) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}
