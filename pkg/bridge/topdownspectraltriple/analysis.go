// Package topdownspectraltriple implements Gate 166: a top-down finite
// spectral-triple ansatz on the 16-dimensional Fock-spinor carrier.
//
// The previous gates proved that no bottom-up contact-mode classification, total
// contact/Fock/scalar representation, or nontrivial contact Dirac operator is
// currently available. Gate 166 deliberately tries the opposite direction: use
// the one-generation 16-state Fock-spinor bookkeeping as the total Hilbert
// space, place the eight Gate-25 Yukawa selection channels into an off-diagonal
// finite Dirac support matrix, and test the resulting fourth-trace gauge-sector
// functional.
//
// The gate records two facts at once:
//
//  1. The unit-incidence version of this top-down ansatz reproduces the
//     representation-trace normalization diag(1,1,1,5/3), hence sin^2=3/8.
//  2. This is not yet a physical Dirac/Yukawa theorem, because Gate 25 derives
//     channel support but not channel amplitudes. Once arbitrary channel
//     amplitudes are allowed, Tr(D_F^4 G^2) is not invariant and the 5/3 ratio
//     is lost. The reproduction is therefore a representation-trace certificate,
//     not a threshold/RG/mass/coupling derivation.
package topdownspectraltriple

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

type HilbertAudit struct {
	Source                 string
	Dimension              int
	FockStateCount         int
	LeftDimension          int
	RightDimension         int
	IdentifiedWithFock16   bool
	UsesObservedInput      bool
	RequiresBranchChoice   bool
	CanonicalBottomUp      bool
	TopDownAnsatzAvailable bool
	Verdict                string
}

type FiniteTripleAudit struct {
	HilbertDimension               int
	DiracMatrixRows                int
	DiracMatrixCols                int
	DiracSymmetric                 bool
	DiracOffDiagonal               bool
	YukawaChannelCount             int
	YukawaChannelSupportComplete   bool
	YukawaAmplitudesDerived        bool
	RealStructureAvailable         bool
	RealStructureInvolutive        bool
	RealStructureCommutesWithD     bool
	RealStructureAnticommutesGamma bool
	GammaAvailable                 bool
	GammaInvolutive                bool
	GammaTraceZero                 bool
	GammaAnticommutesWithD         bool
	OrderOneAxiomTested            bool
	OrderOneAxiomVerified          bool
	PromotableSpectralTriple       bool
	Verdict                        string
}

type ChannelTrace struct {
	Kind         yukawaintertwiner.FermionKind
	Pairs        int
	D4Trace      float64
	LeftY2Trace  float64
	RightY2Trace float64
}

type GaugeTraceAudit struct {
	KSU2T1                  float64
	KSU2T2                  float64
	KSU2T3                  float64
	KU1Y                    float64
	NormalizedT1            float64
	NormalizedT2            float64
	NormalizedT3            float64
	NormalizedY             float64
	BoundaryDiagMatched     bool
	WeakAngleSeed           float64
	WeakAngleSeedMatched    bool
	TraceD4                 float64
	SectorTraces            []ChannelTrace
	FunctionalFormula       string
	UnitIncidenceRequired   bool
	RepresentationTraceOnly bool
	Verdict                 string
}

type AmplitudeSensitivityAudit struct {
	ArbitraryAmplitudesAllowedByPriorGates bool
	UnitAmplitudesDerivedByGate25          bool
	UnitRatioYOverSU2                      float64
	DeformedRatioYOverSU2                  float64
	UnitWeakAngle                          float64
	DeformedWeakAngle                      float64
	BoundaryRatioStable                    bool
	WeakAngleStable                        bool
	ExampleDeformation                     string
	Verdict                                string
}

type FirewallAudit struct {
	ContactModeClassificationBypassedForBoundaryTrace bool
	ContactModeClassificationSolved                   bool
	ThresholdCorrectionsDerived                       bool
	RGRunningDerived                                  bool
	PhysicalCouplingsDerived                          bool
	MassSpectrumDerived                               bool
	GaugeKineticRowsDerived                           int
	BoundaryRowsReproduced                            int
	ResidualNullityBefore                             int
	ResidualNullityAfter                              int
	Verdict                                           string
}

type Analysis struct {
	Fock   spinor.FockSpace
	Yukawa yukawaintertwiner.Analysis

	Hilbert              HilbertAudit
	Triple               FiniteTripleAudit
	Gauge                GaugeTraceAudit
	AmplitudeSensitivity AmplitudeSensitivityAudit
	Firewall             FirewallAudit

	BasisLeft  []string
	BasisRight []string
	Dirac      linear.Matrix
	Gamma      linear.Matrix
	J          linear.Matrix
	D4         linear.Matrix
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
	f, err := spinor.NewCovariantPhaseFockSpace(4)
	if err != nil {
		return Analysis{}, err
	}
	y, err := yukawaintertwiner.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(f, y, 1e-10)
}

func Build(f spinor.FockSpace, y yukawaintertwiner.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if f.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("Gate 166 requires the Gate-14 16-dimensional Fock-spinor carrier, got %d", f.StateCount())
	}
	if !y.ChargeCompatibleYukawaChannelsDerived || len(y.Channels) != 8 || y.LeftDimension != 8 || y.RightDimension != 8 {
		return Analysis{}, fmt.Errorf("Gate 166 requires the Gate-25 eight-channel Yukawa support, got channels=%d left=%d right=%d complete=%t", len(y.Channels), y.LeftDimension, y.RightDimension, y.ChargeCompatibleYukawaChannelsDerived)
	}

	basisLeft := make([]string, y.LeftDimension)
	for _, s := range y.SU2L.States {
		basisLeft[s.Index] = s.Name
	}
	basisRight := make([]string, y.RightDimension)
	rightIndex := map[string]int{}
	for i, s := range y.RightStates {
		basisRight[i] = s.Name
		rightIndex[s.Name] = i
	}

	d, err := buildUnitIncidenceDirac(y, rightIndex)
	if err != nil {
		return Analysis{}, err
	}
	gamma := buildLRGamma(16)
	j, err := buildChannelConjugation(y, rightIndex)
	if err != nil {
		return Analysis{}, err
	}
	d2, err := d.Mul(d)
	if err != nil {
		return Analysis{}, err
	}
	d4, err := d2.Mul(d2)
	if err != nil {
		return Analysis{}, err
	}

	hilbert := HilbertAudit{
		Source:                 "Gate14 H_Fock with top-down identification H_Fock ≅ H_L ⊕ H_R from Gate25",
		Dimension:              16,
		FockStateCount:         f.StateCount(),
		LeftDimension:          y.LeftDimension,
		RightDimension:         y.RightDimension,
		IdentifiedWithFock16:   f.StateCount() == y.LeftDimension+y.RightDimension,
		UsesObservedInput:      false,
		RequiresBranchChoice:   true,
		CanonicalBottomUp:      false,
		TopDownAnsatzAvailable: true,
		Verdict:                "dimensionally exact and useful as a top-down ansatz; not a bottom-up canonical identification of the Fock occupation basis with the left/right Weyl table",
	}

	triple, err := auditTriple(d, gamma, j, eps)
	if err != nil {
		return Analysis{}, err
	}
	triple.YukawaChannelCount = len(y.Channels)
	triple.YukawaChannelSupportComplete = y.ChargeCompatibleYukawaChannelsDerived
	triple.YukawaAmplitudesDerived = false
	triple.PromotableSpectralTriple = false
	triple.OrderOneAxiomTested = false
	triple.OrderOneAxiomVerified = false
	triple.Verdict = "the candidate has D, J, and gamma with the expected matrix identities, but it is still an ansatz because the total algebra representation and Yukawa amplitudes are not derived"

	gauge, err := auditGaugeTrace(y, d4, eps)
	if err != nil {
		return Analysis{}, err
	}
	amp, err := auditAmplitudeSensitivity(y, rightIndex, eps)
	if err != nil {
		return Analysis{}, err
	}

	firewall := FirewallAudit{
		ContactModeClassificationBypassedForBoundaryTrace: gauge.BoundaryDiagMatched && gauge.WeakAngleSeedMatched,
		ContactModeClassificationSolved:                   false,
		ThresholdCorrectionsDerived:                       false,
		RGRunningDerived:                                  false,
		PhysicalCouplingsDerived:                          false,
		MassSpectrumDerived:                               false,
		GaugeKineticRowsDerived:                           0,
		BoundaryRowsReproduced:                            4,
		ResidualNullityBefore:                             3,
		ResidualNullityAfter:                              3,
		Verdict:                                           "the top-down unit-incidence representation trace reproduces the embedded boundary normalization, but it does not derive physical running, thresholds, masses, or couplings",
	}

	return Analysis{
		Fock:                 f,
		Yukawa:               y,
		Hilbert:              hilbert,
		Triple:               triple,
		Gauge:                gauge,
		AmplitudeSensitivity: amp,
		Firewall:             firewall,
		BasisLeft:            basisLeft,
		BasisRight:           basisRight,
		Dirac:                d,
		Gamma:                gamma,
		J:                    j,
		D4:                   d4,
	}, nil
}

func buildUnitIncidenceDirac(y yukawaintertwiner.Analysis, rightIndex map[string]int) (linear.Matrix, error) {
	dim := y.LeftDimension + y.RightDimension
	d := linear.NewMatrix(dim, dim)
	seenLeft := map[int]bool{}
	seenRight := map[int]bool{}
	for _, ch := range y.Channels {
		li := ch.Left.Index
		ri, ok := rightIndex[ch.Right.Name]
		if !ok {
			return linear.Matrix{}, fmt.Errorf("right state %s not in right basis", ch.Right.Name)
		}
		if seenLeft[li] {
			return linear.Matrix{}, fmt.Errorf("left state %s appears in more than one minimal channel", ch.Left.Name)
		}
		if seenRight[ri] {
			return linear.Matrix{}, fmt.Errorf("right state %s appears in more than one minimal channel", ch.Right.Name)
		}
		seenLeft[li] = true
		seenRight[ri] = true
		d.Set(li, y.LeftDimension+ri, 1)
		d.Set(y.LeftDimension+ri, li, 1)
	}
	if len(seenLeft) != y.LeftDimension || len(seenRight) != y.RightDimension {
		return linear.Matrix{}, fmt.Errorf("Yukawa channel support is not a perfect L/R matching: left=%d/%d right=%d/%d", len(seenLeft), y.LeftDimension, len(seenRight), y.RightDimension)
	}
	return d, nil
}

func buildLRGamma(dim int) linear.Matrix {
	g := linear.NewMatrix(dim, dim)
	for i := 0; i < dim; i++ {
		v := -1.0
		if i >= dim/2 {
			v = 1
		}
		g.Set(i, i, v)
	}
	return g
}

func buildChannelConjugation(y yukawaintertwiner.Analysis, rightIndex map[string]int) (linear.Matrix, error) {
	dim := y.LeftDimension + y.RightDimension
	j := linear.NewMatrix(dim, dim)
	for _, ch := range y.Channels {
		li := ch.Left.Index
		ri, ok := rightIndex[ch.Right.Name]
		if !ok {
			return linear.Matrix{}, fmt.Errorf("right state %s not in right basis", ch.Right.Name)
		}
		j.Set(li, y.LeftDimension+ri, 1)
		j.Set(y.LeftDimension+ri, li, 1)
	}
	return j, nil
}

func auditTriple(d, gamma, j linear.Matrix, eps float64) (FiniteTripleAudit, error) {
	dim := d.Rows()
	if d.Rows() != d.Cols() || gamma.Rows() != dim || gamma.Cols() != dim || j.Rows() != dim || j.Cols() != dim {
		return FiniteTripleAudit{}, fmt.Errorf("invalid D/J/gamma dimensions")
	}
	d2, err := d.Mul(d)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	_ = d2
	g2, err := gamma.Mul(gamma)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	j2, err := j.Mul(j)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	jd, err := j.Mul(d)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	dj, err := d.Mul(j)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	jdMinus, err := jd.Sub(dj)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	gd, err := gamma.Mul(d)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	dg, err := d.Mul(gamma)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	gdPlus, err := gd.Add(dg)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	jg, err := j.Mul(gamma)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	gj, err := gamma.Mul(j)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	jgPlus, err := jg.Add(gj)
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	trGamma, err := gamma.Trace()
	if err != nil {
		return FiniteTripleAudit{}, err
	}
	return FiniteTripleAudit{
		HilbertDimension:               dim,
		DiracMatrixRows:                d.Rows(),
		DiracMatrixCols:                d.Cols(),
		DiracSymmetric:                 d.IsSymmetric(eps),
		DiracOffDiagonal:               diagonalMaxAbs(d) < eps,
		RealStructureAvailable:         true,
		RealStructureInvolutive:        j2.AlmostEqual(linear.Identity(dim), eps),
		RealStructureCommutesWithD:     jdMinus.FrobeniusNorm() < eps,
		RealStructureAnticommutesGamma: jgPlus.FrobeniusNorm() < eps,
		GammaAvailable:                 true,
		GammaInvolutive:                g2.AlmostEqual(linear.Identity(dim), eps),
		GammaTraceZero:                 math.Abs(trGamma) < eps,
		GammaAnticommutesWithD:         gdPlus.FrobeniusNorm() < eps,
	}, nil
}

func auditGaugeTrace(y yukawaintertwiner.Analysis, d4 linear.Matrix, eps float64) (GaugeTraceAudit, error) {
	t1Left, err := y.SU2L.TPlus.Add(y.SU2L.TMinus)
	if err != nil {
		return GaugeTraceAudit{}, err
	}
	t1Left = t1Left.Scale(0.5)
	t3Left := y.SU2L.T3
	t1 := blockLeft(t1Left, y.LeftDimension, y.RightDimension)
	t3 := blockLeft(t3Left, y.LeftDimension, y.RightDimension)
	ygen := buildHyperchargeGenerator(y)

	kT1, err := weightedSquareTrace(d4, t1)
	if err != nil {
		return GaugeTraceAudit{}, err
	}
	kT3, err := weightedSquareTrace(d4, t3)
	if err != nil {
		return GaugeTraceAudit{}, err
	}
	kY, err := weightedSquareTrace(d4, ygen)
	if err != nil {
		return GaugeTraceAudit{}, err
	}
	trD4, err := d4.Trace()
	if err != nil {
		return GaugeTraceAudit{}, err
	}
	kT2 := kT1 // The complex T2 generator has the same squared trace as T1 on each SU(2) doublet.
	base := kT3
	weak := base / (base + kY)
	sectors := sectorTraces(y, d4)
	return GaugeTraceAudit{
		KSU2T1:                  kT1,
		KSU2T2:                  kT2,
		KSU2T3:                  kT3,
		KU1Y:                    kY,
		NormalizedT1:            kT1 / base,
		NormalizedT2:            kT2 / base,
		NormalizedT3:            kT3 / base,
		NormalizedY:             kY / base,
		BoundaryDiagMatched:     close(kT1/base, 1, eps) && close(kT2/base, 1, eps) && close(kT3/base, 1, eps) && close(kY/base, 5.0/3.0, eps),
		WeakAngleSeed:           weak,
		WeakAngleSeedMatched:    close(weak, 3.0/8.0, eps),
		TraceD4:                 trD4,
		SectorTraces:            sectors,
		FunctionalFormula:       "K_a = Tr(D_F^4 T_a^2) for unit-incidence D_F; T2 row uses SU(2) isotropy of the complex generator",
		UnitIncidenceRequired:   true,
		RepresentationTraceOnly: true,
		Verdict:                 "unit-incidence D_F^4 is the identity, so the functional reduces to the one-generation representation trace and reproduces diag(1,1,1,5/3) and sin^2=3/8",
	}, nil
}

func auditAmplitudeSensitivity(y yukawaintertwiner.Analysis, rightIndex map[string]int, eps float64) (AmplitudeSensitivityAudit, error) {
	unitD, err := buildWeightedDirac(y, rightIndex, map[yukawaintertwiner.FermionKind]float64{
		yukawaintertwiner.UpType:       1,
		yukawaintertwiner.DownType:     1,
		yukawaintertwiner.NeutrinoType: 1,
		yukawaintertwiner.ElectronType: 1,
	})
	if err != nil {
		return AmplitudeSensitivityAudit{}, err
	}
	deformedD, err := buildWeightedDirac(y, rightIndex, map[yukawaintertwiner.FermionKind]float64{
		yukawaintertwiner.UpType:       2,
		yukawaintertwiner.DownType:     1,
		yukawaintertwiner.NeutrinoType: 1,
		yukawaintertwiner.ElectronType: 1,
	})
	if err != nil {
		return AmplitudeSensitivityAudit{}, err
	}
	unitD4, err := fourthPower(unitD)
	if err != nil {
		return AmplitudeSensitivityAudit{}, err
	}
	defD4, err := fourthPower(deformedD)
	if err != nil {
		return AmplitudeSensitivityAudit{}, err
	}
	unitGauge, err := auditGaugeTrace(y, unitD4, eps)
	if err != nil {
		return AmplitudeSensitivityAudit{}, err
	}
	defGauge, err := auditGaugeTrace(y, defD4, eps)
	if err != nil {
		return AmplitudeSensitivityAudit{}, err
	}
	return AmplitudeSensitivityAudit{
		ArbitraryAmplitudesAllowedByPriorGates: true,
		UnitAmplitudesDerivedByGate25:          false,
		UnitRatioYOverSU2:                      unitGauge.NormalizedY,
		DeformedRatioYOverSU2:                  defGauge.NormalizedY,
		UnitWeakAngle:                          unitGauge.WeakAngleSeed,
		DeformedWeakAngle:                      defGauge.WeakAngleSeed,
		BoundaryRatioStable:                    close(unitGauge.NormalizedY, defGauge.NormalizedY, eps),
		WeakAngleStable:                        close(unitGauge.WeakAngleSeed, defGauge.WeakAngleSeed, eps),
		ExampleDeformation:                     "set the three up-type channel amplitudes to 2 while all other channel amplitudes remain 1",
		Verdict:                                "the 5/3 boundary ratio is not invariant under allowed Yukawa-amplitude deformations; Gate 25 supplies support, not the unit-amplitude Dirac spectrum",
	}, nil
}

func buildWeightedDirac(y yukawaintertwiner.Analysis, rightIndex map[string]int, weights map[yukawaintertwiner.FermionKind]float64) (linear.Matrix, error) {
	dim := y.LeftDimension + y.RightDimension
	d := linear.NewMatrix(dim, dim)
	for _, ch := range y.Channels {
		w, ok := weights[ch.Right.Kind]
		if !ok {
			return linear.Matrix{}, fmt.Errorf("missing weight for %s", ch.Right.Kind)
		}
		ri, ok := rightIndex[ch.Right.Name]
		if !ok {
			return linear.Matrix{}, fmt.Errorf("right state %s not in right basis", ch.Right.Name)
		}
		d.Set(ch.Left.Index, y.LeftDimension+ri, w)
		d.Set(y.LeftDimension+ri, ch.Left.Index, w)
	}
	return d, nil
}

func fourthPower(d linear.Matrix) (linear.Matrix, error) {
	d2, err := d.Mul(d)
	if err != nil {
		return linear.Matrix{}, err
	}
	return d2.Mul(d2)
}

func blockLeft(left linear.Matrix, leftDim, rightDim int) linear.Matrix {
	out := linear.NewMatrix(leftDim+rightDim, leftDim+rightDim)
	for r := 0; r < leftDim; r++ {
		for c := 0; c < leftDim; c++ {
			out.Set(r, c, left.At(r, c))
		}
	}
	return out
}

func buildHyperchargeGenerator(y yukawaintertwiner.Analysis) linear.Matrix {
	dim := y.LeftDimension + y.RightDimension
	out := linear.NewMatrix(dim, dim)
	for _, s := range y.SU2L.States {
		out.Set(s.Index, s.Index, s.Hypercharge)
	}
	for i, s := range y.RightStates {
		out.Set(y.LeftDimension+i, y.LeftDimension+i, s.Hypercharge)
	}
	return out
}

func weightedSquareTrace(weight linear.Matrix, g linear.Matrix) (float64, error) {
	g2, err := g.Mul(g)
	if err != nil {
		return 0, err
	}
	wg2, err := weight.Mul(g2)
	if err != nil {
		return 0, err
	}
	return wg2.Trace()
}

func sectorTraces(y yukawaintertwiner.Analysis, d4 linear.Matrix) []ChannelTrace {
	byKind := map[yukawaintertwiner.FermionKind]*ChannelTrace{}
	for _, ch := range y.Channels {
		ct := byKind[ch.Right.Kind]
		if ct == nil {
			ct = &ChannelTrace{Kind: ch.Right.Kind}
			byKind[ch.Right.Kind] = ct
		}
		li := ch.Left.Index
		ri := y.LeftDimension + rightStateIndex(y.RightStates, ch.Right.Name)
		ct.Pairs++
		ct.D4Trace += d4.At(li, li) + d4.At(ri, ri)
		ct.LeftY2Trace += d4.At(li, li) * ch.Left.Hypercharge * ch.Left.Hypercharge
		ct.RightY2Trace += d4.At(ri, ri) * ch.Right.Hypercharge * ch.Right.Hypercharge
	}
	out := make([]ChannelTrace, 0, len(byKind))
	for _, v := range byKind {
		out = append(out, *v)
	}
	sort.Slice(out, func(i, j int) bool { return string(out[i].Kind) < string(out[j].Kind) })
	return out
}

func rightStateIndex(states []yukawaintertwiner.RightSinglet, name string) int {
	for i, s := range states {
		if s.Name == name {
			return i
		}
	}
	return -1
}

func diagonalMaxAbs(m linear.Matrix) float64 {
	max := 0.0
	for i := 0; i < m.Rows() && i < m.Cols(); i++ {
		if v := math.Abs(m.At(i, i)); v > max {
			max = v
		}
	}
	return max
}

func close(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatBasis(left, right []string) string {
	return fmt.Sprintf("left=[%s]; right=[%s]", strings.Join(left, ", "), strings.Join(right, ", "))
}

func FormatHilbertAudit(a HilbertAudit) string {
	return fmt.Sprintf("dim=%d fock=%d left=%d right=%d identified=%t branchChoice=%t canonicalBottomUp=%t", a.Dimension, a.FockStateCount, a.LeftDimension, a.RightDimension, a.IdentifiedWithFock16, a.RequiresBranchChoice, a.CanonicalBottomUp)
}

func FormatTripleAudit(a FiniteTripleAudit) string {
	return fmt.Sprintf("D=%dx%d symmetric=%t offdiag=%t channels=%d support=%t J^2=%t JD=DJ=%t Jgamma=-gammaJ=%t gamma^2=%t Tr(gamma)=0:%t gammaD=-Dgamma:%t amplitudesDerived=%t triple=%t", a.DiracMatrixRows, a.DiracMatrixCols, a.DiracSymmetric, a.DiracOffDiagonal, a.YukawaChannelCount, a.YukawaChannelSupportComplete, a.RealStructureInvolutive, a.RealStructureCommutesWithD, a.RealStructureAnticommutesGamma, a.GammaInvolutive, a.GammaTraceZero, a.GammaAnticommutesWithD, a.YukawaAmplitudesDerived, a.PromotableSpectralTriple)
}

func FormatGaugeAudit(a GaugeTraceAudit) string {
	return fmt.Sprintf("K=(%.12g, %.12g, %.12g, %.12g), normalized=(%.12g, %.12g, %.12g, %.12g), sin2=%.12g, TrD4=%.12g", a.KSU2T1, a.KSU2T2, a.KSU2T3, a.KU1Y, a.NormalizedT1, a.NormalizedT2, a.NormalizedT3, a.NormalizedY, a.WeakAngleSeed, a.TraceD4)
}

func FormatSectorTraces(xs []ChannelTrace) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s:pairs=%d D4=%.12g Y2=(L %.12g,R %.12g)", x.Kind, x.Pairs, x.D4Trace, x.LeftY2Trace, x.RightY2Trace)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAmplitudeAudit(a AmplitudeSensitivityAudit) string {
	return fmt.Sprintf("unitY/SU2=%.12g deformedY/SU2=%.12g unitSin2=%.12g deformedSin2=%.12g stableRatio=%t stableSin2=%t unitDerived=%t", a.UnitRatioYOverSU2, a.DeformedRatioYOverSU2, a.UnitWeakAngle, a.DeformedWeakAngle, a.BoundaryRatioStable, a.WeakAngleStable, a.UnitAmplitudesDerivedByGate25)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("boundaryBypassed=%t contactSolved=%t thresholds=%t RG=%t couplings=%t masses=%t gaugeRows=%d boundaryRows=%d nullity=%d->%d", a.ContactModeClassificationBypassedForBoundaryTrace, a.ContactModeClassificationSolved, a.ThresholdCorrectionsDerived, a.RGRunningDerived, a.PhysicalCouplingsDerived, a.MassSpectrumDerived, a.GaugeKineticRowsDerived, a.BoundaryRowsReproduced, a.ResidualNullityBefore, a.ResidualNullityAfter)
}
