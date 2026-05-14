// Package fockrepresentationtrace implements Gate 167: representation-trace
// gauge-ratio rigidity and Yukawa-amplitude separation.
//
// Gate 166 showed that a unit-incidence top-down Fock/Yukawa ansatz reproduces
// the embedded boundary normalization through Tr(D_F^4 T_a^2), but also showed
// that the same expression changes when the Yukawa channel amplitudes are
// deformed. Gate 167 corrects the interpretation: the gauge kinetic ratio in a
// finite spectral-action representation is the representation trace Tr_rep(T_a^2)
// over the one-generation fermion content. It counts charge-squared weights and
// is independent of the finite Dirac/Yukawa amplitudes.
//
// The amplitudes are not discarded. They are moved to their correct location:
// the off-diagonal entries of D_F are the finite Yukawa/mass texture data. Their
// singular values encode fermion masses after a scalar scale is supplied, and in
// the replicated generation problem their left-eigenbasis misalignment encodes
// CKM/PMNS mixing. Thus Gate 167 closes the boundary-ratio question while
// reopening the mass-generation problem as a finite Dirac eigenvalue problem.
package fockrepresentationtrace

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/topdownspectraltriple"
	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type Rational struct {
	v *big.Rat
}

func NewRational(n, d int64) Rational {
	if d == 0 {
		panic("zero denominator")
	}
	return Rational{v: new(big.Rat).SetFrac(big.NewInt(n), big.NewInt(d))}
}

func RationalFromSixth(x float64) Rational {
	return NewRational(int64(math.Round(x*6)), 6)
}

func RationalFromHalf(x float64) Rational {
	return NewRational(int64(math.Round(x*2)), 2)
}

func (r Rational) Add(s Rational) Rational { return Rational{v: new(big.Rat).Add(r.v, s.v)} }
func (r Rational) Sub(s Rational) Rational { return Rational{v: new(big.Rat).Sub(r.v, s.v)} }
func (r Rational) Mul(s Rational) Rational { return Rational{v: new(big.Rat).Mul(r.v, s.v)} }
func (r Rational) Div(s Rational) Rational {
	if s.v.Sign() == 0 {
		panic("division by zero")
	}
	return Rational{v: new(big.Rat).Quo(r.v, s.v)}
}
func (r Rational) Square() Rational      { return r.Mul(r) }
func (r Rational) Equal(s Rational) bool { return r.v.Cmp(s.v) == 0 }
func (r Rational) String() string        { return r.v.RatString() }
func (r Rational) Float64() float64 {
	f, _ := r.v.Float64()
	return f
}

type RepresentationState struct {
	Name           string
	Sector         string
	Chirality      string
	Color          int
	T3             Rational
	Hypercharge    Rational
	ContributesSU2 bool
	Y2             Rational
	T3Squared      Rational
}

type SectorTrace struct {
	Sector      string
	States      int
	SU2Doublets int
	T1Squared   Rational
	T2Squared   Rational
	T3Squared   Rational
	Y2          Rational
	Verdict     string
}

type RepresentationTraceAudit struct {
	FermionStates             int
	LeftStates                int
	RightStates               int
	SU2Doublets               int
	HyperchargeStates         int
	KSU2T1                    Rational
	KSU2T2                    Rational
	KSU2T3                    Rational
	KU1Y                      Rational
	NormalizedT1              Rational
	NormalizedT2              Rational
	NormalizedT3              Rational
	NormalizedY               Rational
	WeakAngleSeed             Rational
	BoundaryDiagMatched       bool
	WeakAngleSeedMatched      bool
	AmplitudeIndependent      bool
	UsesDiracFourthPower      bool
	UsesObservedInput         bool
	UsesContactModeClassifier bool
	FunctionalFormula         string
	Verdict                   string
}

type AmplitudeSeparationAudit struct {
	Gate166UnitDWeightedRatio               Rational
	Gate166DeformedDWeightedRatio           Rational
	Gate166UnitDWeightedSin2                Rational
	Gate166DeformedDWeightedSin2            Rational
	RepresentationUnitRatio                 Rational
	RepresentationDeformedRatio             Rational
	RepresentationUnitSin2                  Rational
	RepresentationDeformedSin2              Rational
	DWeightedFunctionalAmplitudeDependent   bool
	RepresentationTraceAmplitudeIndependent bool
	DWeightedDemotedToDiagnostic            bool
	GaugeKineticFunctionalCorrected         bool
	Verdict                                 string
}

type YukawaAmplitudeAudit struct {
	YukawaChannels                   int
	OneGenerationAmplitudeSlots      int
	FermionKindBlocks                int
	ColorUniversalAmplitudesDerived  bool
	NumericalAmplitudesDerived       bool
	FiniteDiracAmplitudeMeaning      string
	OneGenerationSpectrumRule        string
	TrialityReplicatedProblem        string
	MassEigenvaluesAreSingularValues bool
	CKMFromLeftMisalignment          bool
	PMNSFromLeftMisalignment         bool
	ConnectsGate28TextureSearch      bool
	PhysicalMassesDerived            bool
	MixingMatricesDerived            bool
	Verdict                          string
}

type FirewallAudit struct {
	BoundaryGaugeRatioClosed                   bool
	BoundaryWeakAngleSeedClosed                bool
	ContactModeClassificationNeededForBoundary bool
	ContactModeClassificationSolved            bool
	ThresholdCorrectionsDerived                bool
	RGRunningDerived                           bool
	PhysicalCouplingsDerived                   bool
	PhysicalMassesDerived                      bool
	CKMPMNSDerived                             bool
	YukawaTextureProblemOpened                 bool
	ResidualNullityBefore                      int
	ResidualNullityAfter                       int
	Verdict                                    string
}

type Analysis struct {
	Previous topdownspectraltriple.Analysis
	Yukawa   yukawaintertwiner.Analysis

	States         []RepresentationState
	SectorTraces   []SectorTrace
	TraceAudit     RepresentationTraceAudit
	Separation     AmplitudeSeparationAudit
	YukawaAudit    YukawaAmplitudeAudit
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
		defaultValue, defaultErr = buildDefaultUncached()
	})
	return defaultValue, defaultErr
}

func buildDefaultUncached() (Analysis, error) {
	prev, err := topdownspectraltriple.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(prev, 1e-10)
}

func Build(prev topdownspectraltriple.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	y := prev.Yukawa
	if !y.ChargeCompatibleYukawaChannelsDerived || len(y.Channels) != 8 {
		return Analysis{}, fmt.Errorf("Gate 167 requires Gate-25 eight-channel Yukawa support, got %d", len(y.Channels))
	}
	states := buildRepresentationStates(y)
	sectors := buildSectorTraces(states)
	trace := auditRepresentationTrace(states, sectors, eps)
	sep := auditAmplitudeSeparation(prev, trace)
	yuk := auditYukawaAmplitudes(y)
	fw := FirewallAudit{
		BoundaryGaugeRatioClosed:                   trace.BoundaryDiagMatched && trace.AmplitudeIndependent,
		BoundaryWeakAngleSeedClosed:                trace.WeakAngleSeedMatched && trace.AmplitudeIndependent,
		ContactModeClassificationNeededForBoundary: false,
		ContactModeClassificationSolved:            false,
		ThresholdCorrectionsDerived:                false,
		RGRunningDerived:                           false,
		PhysicalCouplingsDerived:                   false,
		PhysicalMassesDerived:                      false,
		CKMPMNSDerived:                             false,
		YukawaTextureProblemOpened:                 true,
		ResidualNullityBefore:                      3,
		ResidualNullityAfter:                       3,
		Verdict:                                    "the embedded boundary gauge ratio is closed as an amplitude-independent representation trace; physical running, thresholds, couplings, masses, and mixing remain open",
	}
	return Analysis{
		Previous:       prev,
		Yukawa:         y,
		States:         states,
		SectorTraces:   sectors,
		TraceAudit:     trace,
		Separation:     sep,
		YukawaAudit:    yuk,
		Firewall:       fw,
		TruthStatement: "diag(1,1,1,5/3) and sin^2_*=3/8 are representation-trace invariants of the one-generation Fock charge table; D_F amplitudes are Yukawa/mass texture data, not gauge-normalization data.",
	}, nil
}

func buildRepresentationStates(y yukawaintertwiner.Analysis) []RepresentationState {
	out := make([]RepresentationState, 0, y.LeftDimension+y.RightDimension)
	for _, s := range y.SU2L.States {
		sector := "L_L"
		if s.Kind == su2lgauge.QuarkDoublet {
			sector = "Q_L"
		}
		t3 := RationalFromHalf(s.T3)
		hy := RationalFromSixth(s.Hypercharge)
		out = append(out, RepresentationState{
			Name:           s.Name,
			Sector:         sector,
			Chirality:      "L",
			Color:          s.Color,
			T3:             t3,
			Hypercharge:    hy,
			ContributesSU2: true,
			Y2:             hy.Square(),
			T3Squared:      t3.Square(),
		})
	}
	for _, s := range y.RightStates {
		sector := string(s.Kind) + "_R"
		switch s.Kind {
		case yukawaintertwiner.UpType:
			sector = "u_R"
		case yukawaintertwiner.DownType:
			sector = "d_R"
		case yukawaintertwiner.NeutrinoType:
			sector = "nu_R"
		case yukawaintertwiner.ElectronType:
			sector = "e_R"
		}
		hy := RationalFromSixth(s.Hypercharge)
		out = append(out, RepresentationState{
			Name:           s.Name,
			Sector:         sector,
			Chirality:      "R",
			Color:          s.Color,
			T3:             NewRational(0, 1),
			Hypercharge:    hy,
			ContributesSU2: false,
			Y2:             hy.Square(),
			T3Squared:      NewRational(0, 1),
		})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Chirality != out[j].Chirality {
			return out[i].Chirality < out[j].Chirality
		}
		return out[i].Name < out[j].Name
	})
	return out
}

func buildSectorTraces(states []RepresentationState) []SectorTrace {
	order := map[string]int{"Q_L": 0, "L_L": 1, "u_R": 2, "d_R": 3, "nu_R": 4, "e_R": 5}
	by := map[string]*SectorTrace{}
	for _, s := range states {
		st := by[s.Sector]
		if st == nil {
			st = &SectorTrace{Sector: s.Sector, T1Squared: NewRational(0, 1), T2Squared: NewRational(0, 1), T3Squared: NewRational(0, 1), Y2: NewRational(0, 1)}
			by[s.Sector] = st
		}
		st.States++
		st.T3Squared = st.T3Squared.Add(s.T3Squared)
		st.Y2 = st.Y2.Add(s.Y2)
	}
	for _, st := range by {
		if st.Sector == "Q_L" || st.Sector == "L_L" {
			st.SU2Doublets = st.States / 2
			st.T1Squared = NewRational(int64(st.SU2Doublets), 2)
			st.T2Squared = NewRational(int64(st.SU2Doublets), 2)
			st.Verdict = "complete SU(2)_L doublet sector; Tr(T1^2)=Tr(T2^2)=Tr(T3^2)=doublets/2"
		} else {
			st.Verdict = "right-singlet sector; contributes to hypercharge trace only"
		}
	}
	out := make([]SectorTrace, 0, len(by))
	for _, st := range by {
		out = append(out, *st)
	}
	sort.Slice(out, func(i, j int) bool { return order[out[i].Sector] < order[out[j].Sector] })
	return out
}

func auditRepresentationTrace(states []RepresentationState, sectors []SectorTrace, eps float64) RepresentationTraceAudit {
	k1, k2, k3, ky := NewRational(0, 1), NewRational(0, 1), NewRational(0, 1), NewRational(0, 1)
	left, right, doublets := 0, 0, 0
	for _, st := range sectors {
		k1 = k1.Add(st.T1Squared)
		k2 = k2.Add(st.T2Squared)
		k3 = k3.Add(st.T3Squared)
		ky = ky.Add(st.Y2)
		doublets += st.SU2Doublets
	}
	for _, s := range states {
		if s.Chirality == "L" {
			left++
		} else {
			right++
		}
	}
	n1 := k1.Div(k3)
	n2 := k2.Div(k3)
	n3 := k3.Div(k3)
	ny := ky.Div(k3)
	sin2 := k3.Div(k3.Add(ky))
	return RepresentationTraceAudit{
		FermionStates:             len(states),
		LeftStates:                left,
		RightStates:               right,
		SU2Doublets:               doublets,
		HyperchargeStates:         len(states),
		KSU2T1:                    k1,
		KSU2T2:                    k2,
		KSU2T3:                    k3,
		KU1Y:                      ky,
		NormalizedT1:              n1,
		NormalizedT2:              n2,
		NormalizedT3:              n3,
		NormalizedY:               ny,
		WeakAngleSeed:             sin2,
		BoundaryDiagMatched:       n1.Equal(NewRational(1, 1)) && n2.Equal(NewRational(1, 1)) && n3.Equal(NewRational(1, 1)) && ny.Equal(NewRational(5, 3)),
		WeakAngleSeedMatched:      sin2.Equal(NewRational(3, 8)),
		AmplitudeIndependent:      true,
		UsesDiracFourthPower:      false,
		UsesObservedInput:         false,
		UsesContactModeClassifier: false,
		FunctionalFormula:         "K_a = Tr_rep(T_a^2) over H_Fock one-generation fermion content; no D_F amplitudes appear",
		Verdict:                   "the embedded boundary normalization is an exact representation-trace invariant, independent of the Yukawa amplitudes",
	}
}

func auditAmplitudeSeparation(prev topdownspectraltriple.Analysis, trace RepresentationTraceAudit) AmplitudeSeparationAudit {
	unitDWeightedRatio := rationalApprox(prev.AmplitudeSensitivity.UnitRatioYOverSU2)
	defDWeightedRatio := rationalApprox(prev.AmplitudeSensitivity.DeformedRatioYOverSU2)
	unitDWeightedSin2 := rationalApprox(prev.AmplitudeSensitivity.UnitWeakAngle)
	defDWeightedSin2 := rationalApprox(prev.AmplitudeSensitivity.DeformedWeakAngle)
	return AmplitudeSeparationAudit{
		Gate166UnitDWeightedRatio:               unitDWeightedRatio,
		Gate166DeformedDWeightedRatio:           defDWeightedRatio,
		Gate166UnitDWeightedSin2:                unitDWeightedSin2,
		Gate166DeformedDWeightedSin2:            defDWeightedSin2,
		RepresentationUnitRatio:                 trace.NormalizedY,
		RepresentationDeformedRatio:             trace.NormalizedY,
		RepresentationUnitSin2:                  trace.WeakAngleSeed,
		RepresentationDeformedSin2:              trace.WeakAngleSeed,
		DWeightedFunctionalAmplitudeDependent:   !unitDWeightedRatio.Equal(defDWeightedRatio) && !unitDWeightedSin2.Equal(defDWeightedSin2),
		RepresentationTraceAmplitudeIndependent: true,
		DWeightedDemotedToDiagnostic:            true,
		GaugeKineticFunctionalCorrected:         true,
		Verdict:                                 "Tr(D_F^4 T_a^2) is a useful unit-incidence diagnostic but not the gauge kinetic functional; Tr_rep(T_a^2) is amplitude-independent",
	}
}

func rationalApprox(x float64) Rational {
	r := new(big.Rat).SetFloat64(x)
	if r == nil {
		return NewRational(0, 1)
	}
	return Rational{v: new(big.Rat).Set(r)}.limitDenominator(1000000)
}

func (r Rational) limitDenominator(max int64) Rational {
	// Continued-fraction approximation. This keeps inherited Gate-166 float
	// diagnostics readable as exact small fractions such as 295/159.
	f, _ := r.v.Float64()
	best := new(big.Rat).SetFloat64(f)
	if best == nil {
		return r
	}
	bestErr := math.Inf(1)
	for d := int64(1); d <= max; d++ {
		n := int64(math.Round(f * float64(d)))
		cand := NewRational(n, d)
		err := math.Abs(cand.Float64() - f)
		if err < bestErr {
			bestErr = err
			best = cand.v
			if err < 1e-12 {
				break
			}
		}
	}
	return Rational{v: new(big.Rat).Set(best)}
}

func auditYukawaAmplitudes(y yukawaintertwiner.Analysis) YukawaAmplitudeAudit {
	blocks := map[yukawaintertwiner.FermionKind]bool{}
	for _, ch := range y.Channels {
		blocks[ch.Right.Kind] = true
	}
	return YukawaAmplitudeAudit{
		YukawaChannels:                   len(y.Channels),
		OneGenerationAmplitudeSlots:      len(y.Channels),
		FermionKindBlocks:                len(blocks),
		ColorUniversalAmplitudesDerived:  false,
		NumericalAmplitudesDerived:       false,
		FiniteDiracAmplitudeMeaning:      "off-diagonal D_F entries are finite Yukawa coupling amplitudes on the Gate-25 channel support",
		OneGenerationSpectrumRule:        "for a perfect one-generation matching, each 2x2 block [[0,y_i],[y_i,0]] has eigenvalues ±|y_i|; D_F^2 carries y_i^2",
		TrialityReplicatedProblem:        "after Gate-26 triality replication, each fermion kind requires a 3x3 Yukawa matrix; quark color is an additional identity factor unless a later theorem breaks it",
		MassEigenvaluesAreSingularValues: true,
		CKMFromLeftMisalignment:          true,
		PMNSFromLeftMisalignment:         true,
		ConnectsGate28TextureSearch:      true,
		PhysicalMassesDerived:            false,
		MixingMatricesDerived:            false,
		Verdict:                          "amplitudes are not a gauge-normalization problem; they are exactly the open mass-generation texture problem in finite-Dirac form",
	}
}

func FormatR(r Rational) string { return r.String() }

func FormatTraceAudit(a RepresentationTraceAudit) string {
	return fmt.Sprintf("K=(%s,%s,%s,%s) normalized=(%s,%s,%s,%s) sin2=%s states=%d left=%d right=%d doublets=%d amplitudeIndependent=%t usesD4=%t",
		a.KSU2T1, a.KSU2T2, a.KSU2T3, a.KU1Y,
		a.NormalizedT1, a.NormalizedT2, a.NormalizedT3, a.NormalizedY,
		a.WeakAngleSeed, a.FermionStates, a.LeftStates, a.RightStates, a.SU2Doublets, a.AmplitudeIndependent, a.UsesDiracFourthPower)
}

func FormatSectorTraces(xs []SectorTrace) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s:states=%d doublets=%d T=(%s,%s,%s) Y2=%s", x.Sector, x.States, x.SU2Doublets, x.T1Squared, x.T2Squared, x.T3Squared, x.Y2)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSeparation(a AmplitudeSeparationAudit) string {
	return fmt.Sprintf("D4-weighted unit/deformed ratio=%s/%s sin2=%s/%s; rep-trace unit/deformed ratio=%s/%s sin2=%s/%s; D4dependent=%t repIndependent=%t",
		a.Gate166UnitDWeightedRatio, a.Gate166DeformedDWeightedRatio,
		a.Gate166UnitDWeightedSin2, a.Gate166DeformedDWeightedSin2,
		a.RepresentationUnitRatio, a.RepresentationDeformedRatio,
		a.RepresentationUnitSin2, a.RepresentationDeformedSin2,
		a.DWeightedFunctionalAmplitudeDependent, a.RepresentationTraceAmplitudeIndependent)
}

func FormatYukawaAudit(a YukawaAmplitudeAudit) string {
	return fmt.Sprintf("channels=%d amplitudeSlots=%d kindBlocks=%d numericDerived=%t colorUniversalDerived=%t masses=%t CKM=%t PMNS=%t gate28=%t",
		a.YukawaChannels, a.OneGenerationAmplitudeSlots, a.FermionKindBlocks, a.NumericalAmplitudesDerived, a.ColorUniversalAmplitudesDerived, a.PhysicalMassesDerived, a.CKMFromLeftMisalignment, a.PMNSFromLeftMisalignment, a.ConnectsGate28TextureSearch)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("boundaryRatioClosed=%t weakSeedClosed=%t contactNeededForBoundary=%t contactSolved=%t thresholds=%t RG=%t couplings=%t masses=%t CKM/PMNS=%t textureOpened=%t nullity=%d->%d",
		a.BoundaryGaugeRatioClosed, a.BoundaryWeakAngleSeedClosed, a.ContactModeClassificationNeededForBoundary, a.ContactModeClassificationSolved, a.ThresholdCorrectionsDerived, a.RGRunningDerived, a.PhysicalCouplingsDerived, a.PhysicalMassesDerived, a.CKMPMNSDerived, a.YukawaTextureProblemOpened, a.ResidualNullityBefore, a.ResidualNullityAfter)
}
