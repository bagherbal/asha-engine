// Package scalarfockspectralpotential implements Gate 168: a scalar-sector
// convergence test between the top-down Fock finite-Dirac spectral action and
// the independently derived contact/Higgs scalar potential of Gate 37.
//
// Gate 167 closed the gauge-kinetic boundary ratio by proving that it is a
// representation trace, not a Yukawa-amplitude trace. Gate 168 deliberately
// tests whether an analogous amplitude-independent closure exists for the
// scalar potential. It does not: the spectral-action scalar coefficients depend
// on the finite Dirac/Yukawa amplitudes through
//
//	A = Tr(Y†Y)       = Σ_i |y_i|²
//	B = Tr((Y†Y)²)   = Σ_i |y_i|⁴
//
// while the full doubled finite Dirac has Tr(D_F²)=2A and Tr(D_F⁴)=2B. The
// scale-free scalar shape comparable to Gate 37 is therefore B/A², not the
// gauge-sector representation trace. Unit incidence gives B/A²=1/8, whereas
// Gate 37 gives λ_shape≈0.258866782 from the contact/Higgs active sector.
//
// The result is a disciplined negative convergence theorem: the two towers do
// not yet agree on the scalar shape under unit incidence, but Gate 37's value
// lies inside the mathematically allowed Yukawa-shape range [1/8,1]. Thus the
// scalar comparison becomes a finite amplitude-texture constraint rather than a
// new gauge-like rigidity theorem.
package scalarfockspectralpotential

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fockrepresentationtrace"
	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type YukawaMomentAudit struct {
	AmplitudeName        string
	ChannelCount         int
	QuadraticMomentA     float64 // Σ y_i² over one chiral Yukawa support.
	QuarticMomentB       float64 // Σ y_i⁴ over one chiral Yukawa support.
	FullDiracTraceD2     float64 // 2A on H_L ⊕ H_R.
	FullDiracTraceD4     float64 // 2B on H_L ⊕ H_R.
	RawQuarticQuadratic  float64 // B/A; cutoff-dependent in the spectral potential.
	ChiralShape          float64 // B/A²; scale-free, comparable to Gate 37 λ_shape.
	FullDiracShape       float64 // Tr(D⁴)/Tr(D²)² = B/(2A²), included as doubled-carrier diagnostic.
	EffectiveSlots       float64 // 1/(B/A²), participation count of amplitudes.
	WeightsByKind        map[string]float64
	Formula              string
	UsesObservedInput    bool
	AmplitudeIndependent bool
	Verdict              string
}

type ContactShapeAudit struct {
	ActiveRealDimension     int
	ProtectedDirectionCount int
	VacuumRadiusSquared     float64
	QuarticTrace            float64
	LambdaShape             float64
	EffectiveParticipation  float64
	PairDegenerate          bool
	ShiftedNormalForm       bool
	ElectroweakScaleDerived bool
	HiggsMassDerived        bool
	Formula                 string
	Verdict                 string
}

type ShapeComparisonAudit struct {
	ComparedShapeFormula          string
	FockUnitShape                 float64
	ContactLambdaShape            float64
	AbsoluteDifference            float64
	RelativeDifference            float64
	UnitIncidenceMatchesContact   bool
	ContactWithinYukawaShapeRange bool
	ShapeRangeMin                 float64
	ShapeRangeMax                 float64
	UnitEffectiveSlots            float64
	ContactEffectiveSlots         float64
	ConvergenceClosed             bool
	ConstraintOpened              bool
	Verdict                       string
}

type SpectralActionAudit struct {
	PotentialTemplate               string
	QuadraticCoefficientSource      string
	QuarticCoefficientSource        string
	CutoffMomentsRequired           bool
	CutoffMomentsDerived            bool
	DimensionalMuDerived            bool
	PhysicalHiggsQuarticDerived     bool
	PhysicalHiggsMassDerived        bool
	ScalarShapeComparable           bool
	ScalarShapeAmplitudeDependent   bool
	GaugeLikeRepresentationRigidity bool
	Verdict                         string
}

type FirewallAudit struct {
	GaugeRatioAlreadyClosed         bool
	ScalarConvergenceClosed         bool
	ScalarAmplitudeConstraintOpened bool
	YukawaAmplitudesDerived         bool
	FermionMassesDerived            bool
	CKMPMNSDerived                  bool
	ElectroweakScaleDerived         bool
	HiggsMassDerived                bool
	ThresholdCorrectionsDerived     bool
	RGRunningDerived                bool
	PhysicalConstantsDerived        bool
	ResidualNullityBefore           int
	ResidualNullityAfter            int
	Verdict                         string
}

type Analysis struct {
	Previous fockrepresentationtrace.Analysis
	Contact  scalarpotential.Analysis

	UnitYukawa     YukawaMomentAudit
	DeformedYukawa YukawaMomentAudit
	ContactShape   ContactShapeAudit
	Comparison     ShapeComparisonAudit
	SpectralAction SpectralActionAudit
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
		prev, err := fockrepresentationtrace.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		contact, err := scalarpotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, contact, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev fockrepresentationtrace.Analysis, contact scalarpotential.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	y := prev.Yukawa
	if !y.ChargeCompatibleYukawaChannelsDerived || len(y.Channels) != 8 {
		return Analysis{}, fmt.Errorf("Gate 168 requires the Gate-25 eight-channel Yukawa support, got %d", len(y.Channels))
	}
	unit := auditYukawaMoments("unit incidence", y, map[yukawaintertwiner.FermionKind]float64{
		yukawaintertwiner.UpType:       1,
		yukawaintertwiner.DownType:     1,
		yukawaintertwiner.NeutrinoType: 1,
		yukawaintertwiner.ElectronType: 1,
	})
	deformed := auditYukawaMoments("up-type amplitude deformation", y, map[yukawaintertwiner.FermionKind]float64{
		yukawaintertwiner.UpType:       2,
		yukawaintertwiner.DownType:     1,
		yukawaintertwiner.NeutrinoType: 1,
		yukawaintertwiner.ElectronType: 1,
	})
	cs := auditContactShape(contact)
	comp := compareShapes(unit, cs, eps)
	spectral := SpectralActionAudit{
		PotentialTemplate:               "V(H) = -c2*f2*Λ²*A*|H|² + c4*f0*B*|H|⁴, with A=Σ|y_i|² and B=Σ|y_i|⁴",
		QuadraticCoefficientSource:      "Tr(D_F²)=2A; equivalently the Yukawa quadratic moment A=Tr(Y†Y)",
		QuarticCoefficientSource:        "Tr(D_F⁴)=2B; equivalently the Yukawa quartic moment B=Tr((Y†Y)²)",
		CutoffMomentsRequired:           true,
		CutoffMomentsDerived:            false,
		DimensionalMuDerived:            false,
		PhysicalHiggsQuarticDerived:     false,
		PhysicalHiggsMassDerived:        false,
		ScalarShapeComparable:           true,
		ScalarShapeAmplitudeDependent:   true,
		GaugeLikeRepresentationRigidity: false,
		Verdict:                         "the finite spectral-action scalar potential is well-formed as a moment ledger, but its shape is a Yukawa-amplitude invariant, not an amplitude-independent representation trace",
	}
	fw := FirewallAudit{
		GaugeRatioAlreadyClosed:         prev.Firewall.BoundaryGaugeRatioClosed && prev.Firewall.BoundaryWeakAngleSeedClosed,
		ScalarConvergenceClosed:         comp.ConvergenceClosed,
		ScalarAmplitudeConstraintOpened: comp.ConstraintOpened,
		YukawaAmplitudesDerived:         false,
		FermionMassesDerived:            false,
		CKMPMNSDerived:                  false,
		ElectroweakScaleDerived:         false,
		HiggsMassDerived:                false,
		ThresholdCorrectionsDerived:     false,
		RGRunningDerived:                false,
		PhysicalConstantsDerived:        false,
		ResidualNullityBefore:           3,
		ResidualNullityAfter:            3,
		Verdict:                         "scalar-sector convergence is not closed; the comparison opens a finite Yukawa-amplitude shape constraint without deriving masses or physical constants",
	}
	return Analysis{
		Previous:       prev,
		Contact:        contact,
		UnitYukawa:     unit,
		DeformedYukawa: deformed,
		ContactShape:   cs,
		Comparison:     comp,
		SpectralAction: spectral,
		Firewall:       fw,
		TruthStatement: "Unlike the gauge ratio, the scalar potential shape is amplitude-sensitive: unit Fock incidence gives B/A^2=1/8, while the contact/Higgs active sector gives λ_shape≈0.258866782. The value is allowed by finite Yukawa moments but is not selected by current gates.",
	}, nil
}

func auditYukawaMoments(name string, y yukawaintertwiner.Analysis, weights map[yukawaintertwiner.FermionKind]float64) YukawaMomentAudit {
	a, b := 0.0, 0.0
	byKind := map[string]float64{}
	for _, ch := range y.Channels {
		w := weights[ch.Right.Kind]
		a += w * w
		b += w * w * w * w
		byKind[string(ch.Right.Kind)] = w
	}
	chiralShape := 0.0
	fullShape := 0.0
	raw := 0.0
	eff := math.Inf(1)
	if a > 0 {
		raw = b / a
		chiralShape = b / (a * a)
	}
	if 2*a > 0 {
		fullShape = (2 * b) / ((2 * a) * (2 * a))
	}
	if chiralShape > 0 {
		eff = 1 / chiralShape
	}
	return YukawaMomentAudit{
		AmplitudeName:        name,
		ChannelCount:         len(y.Channels),
		QuadraticMomentA:     a,
		QuarticMomentB:       b,
		FullDiracTraceD2:     2 * a,
		FullDiracTraceD4:     2 * b,
		RawQuarticQuadratic:  raw,
		ChiralShape:          chiralShape,
		FullDiracShape:       fullShape,
		EffectiveSlots:       eff,
		WeightsByKind:        byKind,
		Formula:              "A=Σ|y_i|², B=Σ|y_i|⁴, Tr(D_F²)=2A, Tr(D_F⁴)=2B, scalar shape=B/A²",
		UsesObservedInput:    false,
		AmplitudeIndependent: false,
		Verdict:              "valid finite Yukawa moment ledger; numerical shape changes when amplitudes change",
	}
}

func auditContactShape(c scalarpotential.Analysis) ContactShapeAudit {
	eff := math.Inf(1)
	if c.LambdaShape > 0 {
		eff = 1 / c.LambdaShape
	}
	return ContactShapeAudit{
		ActiveRealDimension:     c.ActiveRealDimension,
		ProtectedDirectionCount: c.ProtectedDirectionCount,
		VacuumRadiusSquared:     c.VacuumRadiusSquared,
		QuarticTrace:            c.QuarticTrace,
		LambdaShape:             c.LambdaShape,
		EffectiveParticipation:  eff,
		PairDegenerate:          c.PairDegenerate,
		ShiftedNormalForm:       c.ShiftedNormalFormAvailable,
		ElectroweakScaleDerived: c.ElectroweakScaleDerived,
		HiggsMassDerived:        c.HiggsMassDerived,
		Formula:                 "λ_shape = Tr(M_K²)/Tr(M_K)² from the active contact/Higgs sector",
		Verdict:                 "independent contact/Higgs scalar normal form; dimensionless only, with no electroweak vev or Higgs mass derived",
	}
}

func compareShapes(unit YukawaMomentAudit, contact ContactShapeAudit, eps float64) ShapeComparisonAudit {
	diff := math.Abs(unit.ChiralShape - contact.LambdaShape)
	rel := math.Inf(1)
	if contact.LambdaShape > eps {
		rel = diff / contact.LambdaShape
	}
	minShape := 1.0 / float64(unit.ChannelCount)
	maxShape := 1.0
	inRange := contact.LambdaShape >= minShape-eps && contact.LambdaShape <= maxShape+eps
	match := diff <= eps
	return ShapeComparisonAudit{
		ComparedShapeFormula:          "Fock/Yukawa B/A² versus Gate37 contact λ_shape=Tr(M_K²)/Tr(M_K)²",
		FockUnitShape:                 unit.ChiralShape,
		ContactLambdaShape:            contact.LambdaShape,
		AbsoluteDifference:            diff,
		RelativeDifference:            rel,
		UnitIncidenceMatchesContact:   match,
		ContactWithinYukawaShapeRange: inRange,
		ShapeRangeMin:                 minShape,
		ShapeRangeMax:                 maxShape,
		UnitEffectiveSlots:            unit.EffectiveSlots,
		ContactEffectiveSlots:         contact.EffectiveParticipation,
		ConvergenceClosed:             match,
		ConstraintOpened:              !match && inRange,
		Verdict:                       "unit-incidence Fock scalar shape does not match Gate37, but Gate37 lies in the allowed Yukawa moment range; scalar agreement requires a nontrivial amplitude-selection theorem",
	}
}

func FormatYukawaMoment(a YukawaMomentAudit) string {
	return fmt.Sprintf("%s: channels=%d A=%.12g B=%.12g TrD2=%.12g TrD4=%.12g B/A=%.12g B/A^2=%.12g fullDshape=%.12g Neff=%.12g weights={%s}", a.AmplitudeName, a.ChannelCount, a.QuadraticMomentA, a.QuarticMomentB, a.FullDiracTraceD2, a.FullDiracTraceD4, a.RawQuarticQuadratic, a.ChiralShape, a.FullDiracShape, a.EffectiveSlots, formatWeights(a.WeightsByKind))
}

func FormatContactShape(a ContactShapeAudit) string {
	return fmt.Sprintf("active=%d protected=%d r0^2=%.12g TrMK2=%.12g lambda_shape=%.12g Neff=%.12g pair=%t shifted=%t EWscale=%t HiggsMass=%t", a.ActiveRealDimension, a.ProtectedDirectionCount, a.VacuumRadiusSquared, a.QuarticTrace, a.LambdaShape, a.EffectiveParticipation, a.PairDegenerate, a.ShiftedNormalForm, a.ElectroweakScaleDerived, a.HiggsMassDerived)
}

func FormatComparison(a ShapeComparisonAudit) string {
	return fmt.Sprintf("unitShape=%.12g contactShape=%.12g diff=%.12g rel=%.12g range=[%.12g,%.12g] inRange=%t match=%t unitNeff=%.12g contactNeff=%.12g", a.FockUnitShape, a.ContactLambdaShape, a.AbsoluteDifference, a.RelativeDifference, a.ShapeRangeMin, a.ShapeRangeMax, a.ContactWithinYukawaShapeRange, a.UnitIncidenceMatchesContact, a.UnitEffectiveSlots, a.ContactEffectiveSlots)
}

func FormatSpectralAction(a SpectralActionAudit) string {
	return fmt.Sprintf("template=%q quadratic=%q quartic=%q cutoffRequired=%t cutoffDerived=%t scalarShapeComparable=%t amplitudeDependent=%t gaugeLikeRigidity=%t", a.PotentialTemplate, a.QuadraticCoefficientSource, a.QuarticCoefficientSource, a.CutoffMomentsRequired, a.CutoffMomentsDerived, a.ScalarShapeComparable, a.ScalarShapeAmplitudeDependent, a.GaugeLikeRepresentationRigidity)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gaugeClosed=%t scalarClosed=%t scalarConstraint=%t amplitudes=%t masses=%t CKM/PMNS=%t EWscale=%t HiggsMass=%t thresholds=%t RG=%t constants=%t nullity=%d->%d", a.GaugeRatioAlreadyClosed, a.ScalarConvergenceClosed, a.ScalarAmplitudeConstraintOpened, a.YukawaAmplitudesDerived, a.FermionMassesDerived, a.CKMPMNSDerived, a.ElectroweakScaleDerived, a.HiggsMassDerived, a.ThresholdCorrectionsDerived, a.RGRunningDerived, a.PhysicalConstantsDerived, a.ResidualNullityBefore, a.ResidualNullityAfter)
}

func formatWeights(m map[string]float64) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s:%.6g", k, m[k]))
	}
	return strings.Join(parts, ",")
}
