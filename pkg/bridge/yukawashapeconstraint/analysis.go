// Package yukawashapeconstraint implements Gate 169: a finite Yukawa amplitude
// texture search constrained by the Gate-37 contact/Higgs scalar shape.
//
// Gate 168 showed that the scalar spectral-action shape is not a gauge-like
// representation trace. For finite Yukawa amplitudes y_i it is
//
//	B/A^2 = (Σ_i |y_i|^4)/(Σ_i |y_i|^2)^2,
//
// while Gate 37 independently supplies the contact/Higgs target
// λ_contact = Tr(M_K²)/Tr(M_K)². Gate 169 therefore asks a more precise
// question: does the existing finite data select a Yukawa amplitude texture
// whose moment shape equals the contact target?
//
// The gate finds a conditional shape-level match if the eight Gate-25 scalar
// channels are first quotiented into four Higgs-conjugate amplitude classes and
// those squared amplitudes are identified with the four active contact/Higgs
// eigenvalues. This is not yet a mass theorem: the conjugate-channel quotient,
// the assignment of the two high and two low contact weights to fermion kinds,
// the generation lift, phases, scales, and mixing matrices are still missing.
package yukawashapeconstraint

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarfockspectralpotential"
)

type ShapeTargetAudit struct {
	SourceGate             string
	Formula                string
	ExactValue             string
	FloatValue             float64
	RequiredEffectiveSlots float64
	EightSlotRangeMin      float64
	EightSlotRangeMax      float64
	FourClassRangeMin      float64
	FourClassRangeMax      float64
	InEightSlotRange       bool
	InFourClassRange       bool
	IntegerEqualSlotMatch  bool
	UsesObservedMassInput  bool
	Verdict                string
}

type TextureCandidate struct {
	Name                    string
	SlotCount               int
	IndependentClasses      int
	SquaredAmplitudeWeights []float64
	Shape                   float64
	EffectiveSlots          float64
	MatchesTarget           bool
	UsesContactEigenvalues  bool
	RequiresPairCollapse    bool
	RequiresKindAssignment  bool
	CanonicalSelected       bool
	DerivesMasses           bool
	DerivesMixing           bool
	Detail                  string
}

type PairCollapseAudit struct {
	Gate25Channels                     int
	FermionKindBlocks                  int
	ScalarConjugateMultiplicity        int
	FourClassQuotientAvailable         bool
	FourClassQuotientDerived           bool
	ContactSpectrumClasses             int
	ContactSpectrumPairDegenerate      bool
	ConditionalShapeMatch              bool
	PairCollapseRequiredForMatch       bool
	DirectEightChannelDuplicationFails bool
	KindAssignmentAmbiguity            int
	SquaredAmplitudeRatio              float64
	AmplitudeRatio                     float64
	SquaredAmplitudeRatioExact         string
	Verdict                            string
}

type GenerationTextureAudit struct {
	GenerationCount            int
	FermionKindBlocks          int
	KindTextureMatrices        int
	EntriesPerGeneralTexture   int
	TotalGeneralTextureEntries int
	ScalarShapeConstraints     int
	ShapeConstraintOnlyMoment  bool
	TextureUnderdetermined     bool
	PhasesDerived              bool
	CKMPMNSDerived             bool
	FermionMassesDerived       bool
	Verdict                    string
}

type FirewallAudit struct {
	GaugeRatioClosed                     bool
	ScalarShapeTargetAvailable           bool
	ConditionalFourClassMatchFound       bool
	EightChannelAmplitudeTextureSelected bool
	PairCollapseDerived                  bool
	KindAssignmentDerived                bool
	GenerationTextureDerived             bool
	YukawaAmplitudesDerived              bool
	FermionMassesDerived                 bool
	CKMPMNSDerived                       bool
	ElectroweakScaleDerived              bool
	HiggsMassDerived                     bool
	ThresholdCorrectionsDerived          bool
	RGRunningDerived                     bool
	PhysicalConstantsDerived             bool
	ResidualNullityBefore                int
	ResidualNullityAfter                 int
	Verdict                              string
}

type Analysis struct {
	Previous scalarfockspectralpotential.Analysis

	Target         ShapeTargetAudit
	Candidates     []TextureCandidate
	Best           TextureCandidate
	PairCollapse   PairCollapseAudit
	Generation     GenerationTextureAudit
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
		prev, err := scalarfockspectralpotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev scalarfockspectralpotential.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !prev.Comparison.ConstraintOpened || prev.ContactShape.LambdaShape <= eps {
		return Analysis{}, fmt.Errorf("Gate 169 requires the Gate-168 scalar-shape constraint")
	}
	active := append([]float64(nil), prev.Contact.ActiveSpectrum...)
	if len(active) != 4 {
		return Analysis{}, fmt.Errorf("Gate 169 requires four active contact/Higgs eigenvalues, got %d", len(active))
	}
	sort.Slice(active, func(i, j int) bool { return active[i] > active[j] })

	target := auditTarget(prev, eps)
	unit8 := candidate("unit eight-channel incidence", ones(8), prev.ContactShape.LambdaShape, eps, false, false, false,
		"The Gate-25 support with all eight channel amplitudes equal. This is the Gate-168 baseline and does not match the contact scalar shape.")
	equal4 := candidate("equal four-class Higgs quotient", ones(4), prev.ContactShape.LambdaShape, eps, false, true, false,
		"A hypothetical quotient to four fermion-kind amplitude classes, but with no contact anisotropy. It gives 1/4 and is close but not equal to the Gate-37 target.")
	duplicated := candidate("contact spectrum duplicated across Phi± channels", duplicate(active), prev.ContactShape.LambdaShape, eps, true, false, true,
		"Every active contact eigenvalue is copied into both scalar-conjugate channels. A and B both double, so the shape is exactly half of the contact target.")
	collapsed := candidate("four-class contact-spectrum amplitude target", active, prev.ContactShape.LambdaShape, eps, true, true, true,
		"The squared amplitudes of four Higgs-conjugate Yukawa classes are proportional to the four active contact/Higgs eigenvalues. This matches the scalar shape exactly, conditionally on a future pair-collapse and kind-assignment theorem.")

	candidates := []TextureCandidate{unit8, equal4, duplicated, collapsed}
	best := collapsed
	pc := auditPairCollapse(prev, active, duplicated, collapsed, eps)
	gen := auditGeneration(prev)
	fw := FirewallAudit{
		GaugeRatioClosed:                     prev.Firewall.GaugeRatioAlreadyClosed,
		ScalarShapeTargetAvailable:           target.InEightSlotRange && target.InFourClassRange,
		ConditionalFourClassMatchFound:       collapsed.MatchesTarget,
		EightChannelAmplitudeTextureSelected: false,
		PairCollapseDerived:                  pc.FourClassQuotientDerived,
		KindAssignmentDerived:                false,
		GenerationTextureDerived:             false,
		YukawaAmplitudesDerived:              false,
		FermionMassesDerived:                 false,
		CKMPMNSDerived:                       false,
		ElectroweakScaleDerived:              false,
		HiggsMassDerived:                     false,
		ThresholdCorrectionsDerived:          false,
		RGRunningDerived:                     false,
		PhysicalConstantsDerived:             false,
		ResidualNullityBefore:                3,
		ResidualNullityAfter:                 3,
		Verdict:                              "Gate37 supplies a finite scalar-shape target and a conditional four-class amplitude pattern, but the actual Yukawa amplitudes, generation textures, masses, mixing, thresholds, and running remain underived",
	}

	return Analysis{
		Previous:       prev,
		Target:         target,
		Candidates:     candidates,
		Best:           best,
		PairCollapse:   pc,
		Generation:     gen,
		Firewall:       fw,
		TruthStatement: "Gate 169 turns the scalar mismatch into a finite Yukawa moment target: λ_contact=1197/4624 is matched by a four-class contact-spectrum amplitude pattern, but only conditionally. The eight-channel Dirac amplitude texture is not yet selected because scalar-conjugate pair collapse, fermion-kind assignment, generation lift, phases, and non-commuting texture data remain open.",
	}, nil
}

func auditTarget(prev scalarfockspectralpotential.Analysis, eps float64) ShapeTargetAudit {
	target := prev.ContactShape.LambdaShape
	min8 := 1.0 / float64(prev.UnitYukawa.ChannelCount)
	max8 := 1.0
	min4 := 1.0 / 4.0
	max4 := 1.0
	integerMatch := false
	for n := 1; n <= prev.UnitYukawa.ChannelCount; n++ {
		if math.Abs(target-1.0/float64(n)) <= eps {
			integerMatch = true
			break
		}
	}
	return ShapeTargetAudit{
		SourceGate:             "Gate37 via Gate168",
		Formula:                "λ_contact = Tr(M_K²)/Tr(M_K)²; finite certificate λ_contact = 1197/4624",
		ExactValue:             "1197/4624",
		FloatValue:             target,
		RequiredEffectiveSlots: 1 / target,
		EightSlotRangeMin:      min8,
		EightSlotRangeMax:      max8,
		FourClassRangeMin:      min4,
		FourClassRangeMax:      max4,
		InEightSlotRange:       target >= min8-eps && target <= max8+eps,
		InFourClassRange:       target >= min4-eps && target <= max4+eps,
		IntegerEqualSlotMatch:  integerMatch,
		UsesObservedMassInput:  false,
		Verdict:                "the contact/Higgs scalar shape is an exact finite target inside the allowed Yukawa moment range, but it is not equal to any equal-amplitude integer-slot shape",
	}
}

func candidate(name string, squaredWeights []float64, target, eps float64, usesContact, requiresKindAssignment, requiresPairCollapse bool, detail string) TextureCandidate {
	sh := shape(squaredWeights)
	eff := math.Inf(1)
	if sh > 0 {
		eff = 1 / sh
	}
	return TextureCandidate{
		Name:                    name,
		SlotCount:               len(squaredWeights),
		IndependentClasses:      len(distinct(squaredWeights, eps)),
		SquaredAmplitudeWeights: append([]float64(nil), squaredWeights...),
		Shape:                   sh,
		EffectiveSlots:          eff,
		MatchesTarget:           math.Abs(sh-target) <= eps,
		UsesContactEigenvalues:  usesContact,
		RequiresPairCollapse:    requiresPairCollapse,
		RequiresKindAssignment:  requiresKindAssignment,
		CanonicalSelected:       false,
		DerivesMasses:           false,
		DerivesMixing:           false,
		Detail:                  detail,
	}
}

func auditPairCollapse(prev scalarfockspectralpotential.Analysis, active []float64, duplicated, collapsed TextureCandidate, eps float64) PairCollapseAudit {
	high := active[0]
	low := active[len(active)-1]
	sqRatio := math.Inf(1)
	ampRatio := math.Inf(1)
	if low > eps {
		sqRatio = high / low
		ampRatio = math.Sqrt(sqRatio)
	}
	kindBlocks := prev.Previous.YukawaAudit.FermionKindBlocks
	if kindBlocks == 0 {
		kindBlocks = 4
	}
	ambiguity := 0
	if kindBlocks == 4 {
		ambiguity = 6 // choose which two of four fermion kinds receive the high contact pair.
	}
	return PairCollapseAudit{
		Gate25Channels:                     prev.UnitYukawa.ChannelCount,
		FermionKindBlocks:                  kindBlocks,
		ScalarConjugateMultiplicity:        2,
		FourClassQuotientAvailable:         prev.UnitYukawa.ChannelCount == 2*kindBlocks,
		FourClassQuotientDerived:           false,
		ContactSpectrumClasses:             len(active),
		ContactSpectrumPairDegenerate:      prev.ContactShape.PairDegenerate,
		ConditionalShapeMatch:              collapsed.MatchesTarget,
		PairCollapseRequiredForMatch:       collapsed.RequiresPairCollapse,
		DirectEightChannelDuplicationFails: !duplicated.MatchesTarget,
		KindAssignmentAmbiguity:            ambiguity,
		SquaredAmplitudeRatio:              sqRatio,
		AmplitudeRatio:                     ampRatio,
		SquaredAmplitudeRatioExact:         "(34+sqrt(41))/(34-sqrt(41))",
		Verdict:                            "a four-class Higgs-conjugate quotient would let the contact active spectrum match the scalar shape, but the quotient and the assignment of the two high/two low weights to fermion kinds are not derived by current gates",
	}
}

func auditGeneration(prev scalarfockspectralpotential.Analysis) GenerationTextureAudit {
	kinds := prev.Previous.YukawaAudit.FermionKindBlocks
	if kinds == 0 {
		kinds = 4
	}
	generationCount := 3
	entriesPer := generationCount * generationCount
	return GenerationTextureAudit{
		GenerationCount:            generationCount,
		FermionKindBlocks:          kinds,
		KindTextureMatrices:        kinds,
		EntriesPerGeneralTexture:   entriesPer,
		TotalGeneralTextureEntries: kinds * entriesPer,
		ScalarShapeConstraints:     1,
		ShapeConstraintOnlyMoment:  true,
		TextureUnderdetermined:     true,
		PhasesDerived:              false,
		CKMPMNSDerived:             false,
		FermionMassesDerived:       false,
		Verdict:                    "the scalar shape gives one global scale-free moment constraint; after triality, the mass problem is still a set of four 3x3 Yukawa matrices with undetermined entries, phases, eigenvalues, and relative eigenbases",
	}
}

func shape(squaredWeights []float64) float64 {
	sum, sq := 0.0, 0.0
	for _, w := range squaredWeights {
		sum += w
		sq += w * w
	}
	if sum == 0 {
		return 0
	}
	return sq / (sum * sum)
}

func ones(n int) []float64 {
	out := make([]float64, n)
	for i := range out {
		out[i] = 1
	}
	return out
}

func duplicate(xs []float64) []float64 {
	out := make([]float64, 0, 2*len(xs))
	for _, x := range xs {
		out = append(out, x, x)
	}
	return out
}

func distinct(xs []float64, eps float64) []float64 {
	if len(xs) == 0 {
		return nil
	}
	ys := append([]float64(nil), xs...)
	sort.Float64s(ys)
	out := []float64{ys[0]}
	for _, y := range ys[1:] {
		if math.Abs(y-out[len(out)-1]) > eps {
			out = append(out, y)
		}
	}
	return out
}

func FormatTarget(a ShapeTargetAudit) string {
	return fmt.Sprintf("%s=%s≈%.12g Neff=%.12g range8=[%.12g,%.12g] range4=[%.12g,%.12g] in8=%t in4=%t integerEqual=%t observed=%t", a.Formula, a.ExactValue, a.FloatValue, a.RequiredEffectiveSlots, a.EightSlotRangeMin, a.EightSlotRangeMax, a.FourClassRangeMin, a.FourClassRangeMax, a.InEightSlotRange, a.InFourClassRange, a.IntegerEqualSlotMatch, a.UsesObservedMassInput)
}

func FormatCandidate(c TextureCandidate) string {
	return fmt.Sprintf("%s slots=%d classes=%d shape=%.12g Neff=%.12g match=%t contact=%t pairCollapse=%t kindAssign=%t canonical=%t masses=%t mixing=%t weights=%s", c.Name, c.SlotCount, c.IndependentClasses, c.Shape, c.EffectiveSlots, c.MatchesTarget, c.UsesContactEigenvalues, c.RequiresPairCollapse, c.RequiresKindAssignment, c.CanonicalSelected, c.DerivesMasses, c.DerivesMixing, formatFloats(c.SquaredAmplitudeWeights))
}

func FormatCandidates(cs []TextureCandidate) string {
	parts := make([]string, 0, len(cs))
	for _, c := range cs {
		parts = append(parts, FormatCandidate(c))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatPairCollapse(a PairCollapseAudit) string {
	return fmt.Sprintf("gate25=%d kindBlocks=%d scalarConj=%d quotientAvailable=%t quotientDerived=%t contactClasses=%d pairDegenerate=%t conditionalMatch=%t direct8Fails=%t ambiguity=%d sqRatio=%.12g ampRatio=%.12g exactSqRatio=%s", a.Gate25Channels, a.FermionKindBlocks, a.ScalarConjugateMultiplicity, a.FourClassQuotientAvailable, a.FourClassQuotientDerived, a.ContactSpectrumClasses, a.ContactSpectrumPairDegenerate, a.ConditionalShapeMatch, a.DirectEightChannelDuplicationFails, a.KindAssignmentAmbiguity, a.SquaredAmplitudeRatio, a.AmplitudeRatio, a.SquaredAmplitudeRatioExact)
}

func FormatGeneration(a GenerationTextureAudit) string {
	return fmt.Sprintf("generations=%d kindMatrices=%d entriesPer=%d totalEntries=%d scalarConstraints=%d momentOnly=%t underdetermined=%t phases=%t masses=%t CKM/PMNS=%t", a.GenerationCount, a.KindTextureMatrices, a.EntriesPerGeneralTexture, a.TotalGeneralTextureEntries, a.ScalarShapeConstraints, a.ShapeConstraintOnlyMoment, a.TextureUnderdetermined, a.PhasesDerived, a.FermionMassesDerived, a.CKMPMNSDerived)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gaugeClosed=%t scalarTarget=%t conditional4Class=%t selected8=%t pairCollapse=%t kindAssign=%t generationTexture=%t amplitudes=%t masses=%t CKM/PMNS=%t EWscale=%t HiggsMass=%t thresholds=%t RG=%t constants=%t nullity=%d->%d", a.GaugeRatioClosed, a.ScalarShapeTargetAvailable, a.ConditionalFourClassMatchFound, a.EightChannelAmplitudeTextureSelected, a.PairCollapseDerived, a.KindAssignmentDerived, a.GenerationTextureDerived, a.YukawaAmplitudesDerived, a.FermionMassesDerived, a.CKMPMNSDerived, a.ElectroweakScaleDerived, a.HiggsMassDerived, a.ThresholdCorrectionsDerived, a.RGRunningDerived, a.PhysicalConstantsDerived, a.ResidualNullityBefore, a.ResidualNullityAfter)
}

func formatFloats(xs []float64) string {
	if len(xs) == 0 {
		return "[]"
	}
	out := make([]string, len(xs))
	for i, x := range xs {
		out[i] = fmt.Sprintf("%.10g", x)
	}
	return "[" + strings.Join(out, ", ") + "]"
}
