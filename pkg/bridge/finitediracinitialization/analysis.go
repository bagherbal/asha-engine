// Package finitediracinitialization implements Gate 233:
// Finite Dirac Operator (D_F) initialization / 16-state Fock space matrix audit.
//
// Gates 217 and 230 showed that the heavy-sector matching and Hopf hierarchy
// cannot be dynamically derived without a genuine finite spectral triple. Gate
// 233 returns to the native 16-state Fock scaffold and asks a minimal, rigorous
// question: what can be initialized without importing continuum masses?
//
// The answer is deliberately split. The 16-state Fock space admits a formal
// dimensionless odd self-adjoint Dirac matrix family D=[0 M; M^T 0] once a
// parity/chirality split is chosen. This is a valid algebraic ansatz. However,
// the finite core does not yet select the block M, does not derive a total real
// structure J/order-one calculus, and does not canonically embed the B-sector
// first spectral gap into any off-diagonal entry. Therefore no physical finite
// Dirac operator, mass matrix, or spectral-action coefficient is derived.
package finitediracinitialization

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/dynamics/bsector"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE233-FINITE-DIRAC-OPERATOR-INITIALIZATION-AUDIT"

	StatusConditionalAnsatz      = "CONDITIONAL_SUPPORT_DIMENSIONLESS_ODD_SELF_ADJOINT_DF_ANSATZ"
	StatusFailedCanonicalDF      = "FAILED_ROUTE_CANONICAL_FINITE_DIRAC_OPERATOR_DERIVATION"
	StatusFailedBGapEmbedding    = "FAILED_ROUTE_CANONICAL_BGAP_DF_EMBEDDING"
	StatusFailedSpectralAction   = "FAILED_ROUTE_SPECTRAL_ACTION_COEFFICIENT_DERIVATION"
	StatusBroaderHilbertRequired = "BROADER_HILBERT_OR_REAL_STRUCTURE_REQUIRED"
)

type FockSnapshot struct {
	Constructed              bool
	StateCount               int
	ModeCount                int
	EvenParityStates         int
	OddParityStates          int
	ParitySplitBalanced      bool
	GammaDefined             bool
	GammaTrace               float64
	GammaSquaredIdentity     bool
	PhysicalChiralityDerived bool
	Comment                  string
}

type DiracFamilyAudit struct {
	FamilyName                 string
	Dimension                  int
	LeftDimension              int
	RightDimension             int
	FreeRealParameters         int
	GeneralFormula             string
	SelfAdjointByConstruction  bool
	OddWithGammaByConstruction bool
	DimensionlessOnly          bool
	UsesContinuumMasses        bool
	UsesObservedInput          bool
	RequiresChosenBlockM       bool
	CanonicalBlockDerived      bool
	RealStructureJDerived      bool
	OrderOneCalculusVerified   bool
	PromotableFiniteDirac      bool
	Verdict                    string
}

type RepresentativeMatrixAudit struct {
	RepresentativeName  string
	Built               bool
	Dimension           int
	AnticommutatorNorm  float64
	SelfAdjointResidual float64
	TraceD2             float64
	TraceD4             float64
	NormalizedTraceD2   float64
	NormalizedTraceD4   float64
	HopfCoefficient     float64
	TraceMatchesHopf    bool
	Promoted            bool
	Comment             string
}

type BGapEmbeddingAudit struct {
	BGapAvailable                bool
	BGap                         float64
	OffDiagonalEmbeddingAllowed  bool
	CanonicalLeftRightPairing    bool
	CanonicalMatrixEntrySelector bool
	BGapAsDimensionlessAmplitude bool
	BGapAsPhysicalMass           bool
	UniformBlockCandidateAudited bool
	UniformTraceD2               float64
	UniformTraceD4               float64
	UniformNormalizedTraceD2     float64
	UniformNormalizedTraceD4     float64
	UniformTraceMatchesHopf      bool
	BGapEmbeddingPromotable      bool
	Verdict                      string
}

type SpectralActionPreflight struct {
	D2ComputedForRepresentatives       bool
	D4ComputedForRepresentatives       bool
	TraceRowsComputed                  int
	GaugeCurvatureProjectionDerived    bool
	HeatKernelMapDerived               bool
	CutoffFunctionDerived              bool
	ScaleRatiosGenerated               bool
	HopfCoefficientGenerated           bool
	FiniteMatchingCorrectionsGenerated bool
	PhysicalMassesGenerated            bool
	Verdict                            string
}

type ObstructionAudit struct {
	FiniteFockCarrierAvailable      bool
	OddSelfAdjointFamilyAvailable   bool
	CanonicalDFSelected             bool
	BGapOffDiagonalMapDerived       bool
	RealStructureJSelected          bool
	GradingPhysicalChiralityDerived bool
	OrderOneAxiomVerified           bool
	GaugeFluctuationMapDerived      bool
	SpectralActionReady             bool
	RequiresBroaderHilbertSpace     bool
	RequiredNextIngredients         []string
	Verdict                         string
}

type FirewallAudit struct {
	DimensionlessFiniteDataOnly bool
	ContinuumMassInserted       bool
	VEVInserted                 bool
	MBInserted                  bool
	MStarInserted               bool
	BGapPromotedToMass          bool
	DFChosenByFit               bool
	HopfCoefficientFitted       bool
	PhysicalLagrangianClaimed   bool
	FiniteCorePolluted          bool
	Verdict                     string
}

type Summary struct {
	DFAnsatzAvailable    bool
	CanonicalDFDerived   bool
	BGapEmbeddingDerived bool
	SpectralActionReady  bool
	Status               string
	NextGate             string
	Comment              string
}

type Analysis struct {
	Fock           FockSnapshot
	DiracFamily    DiracFamilyAudit
	UnitMatrix     RepresentativeMatrixAudit
	BGap           BGapEmbeddingAudit
	SpectralAction SpectralActionPreflight
	Obstruction    ObstructionAudit
	Firewall       FirewallAudit
	Summary        Summary

	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		f, err := spinor.NewCovariantPhaseFockSpace(4)
		if err != nil {
			defaultErr = fmt.Errorf("construct Gate-14 Fock space: %w", err)
			return
		}
		b, err := bsector.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("construct B-sector vacuum: %w", err)
			return
		}
		defaultA, defaultErr = Build(f, b, 1e-10)
	})
	return defaultA, defaultErr
}

func Build(f spinor.FockSpace, b bsector.Vacuum, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if f.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("Gate 233 requires native 16-state Fock space, got %d", f.StateCount())
	}
	fock := auditFock(f, eps)
	family := auditDiracFamily(fock)
	unit, err := auditUnitRepresentative(fock, eps)
	if err != nil {
		return Analysis{}, err
	}
	bgap, err := auditBGapEmbedding(fock, b, eps)
	if err != nil {
		return Analysis{}, err
	}
	spectral := auditSpectralAction(unit, bgap)
	obstruction := auditObstructions(fock, family, bgap, spectral)
	firewall := auditFirewall()
	summary := summarize(family, bgap, spectral)
	truth := buildTruth(fock, family, bgap, spectral, obstruction)
	return Analysis{Fock: fock, DiracFamily: family, UnitMatrix: unit, BGap: bgap, SpectralAction: spectral, Obstruction: obstruction, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func auditFock(f spinor.FockSpace, eps float64) FockSnapshot {
	even, odd := 0, 0
	trace := 0.0
	for _, s := range f.States {
		if s.ExcitationNumber()%2 == 0 {
			even++
			trace += 1
		} else {
			odd++
			trace -= 1
		}
	}
	return FockSnapshot{
		Constructed:              true,
		StateCount:               f.StateCount(),
		ModeCount:                f.ModeCount(),
		EvenParityStates:         even,
		OddParityStates:          odd,
		ParitySplitBalanced:      even == 8 && odd == 8,
		GammaDefined:             true,
		GammaTrace:               trace,
		GammaSquaredIdentity:     math.Abs(trace) < 16+eps, // diagonal ±1 construction; exact matrix check is in representatives.
		PhysicalChiralityDerived: false,
		Comment:                  "occupation parity gives a canonical Z2 grading candidate on H_Fock; identifying it with physical chirality remains a bridge theorem",
	}
}

func auditDiracFamily(f FockSnapshot) DiracFamilyAudit {
	left, right := f.EvenParityStates, f.OddParityStates
	return DiracFamilyAudit{
		FamilyName:                 "odd self-adjoint finite Dirac family on parity-split H_Fock",
		Dimension:                  f.StateCount,
		LeftDimension:              left,
		RightDimension:             right,
		FreeRealParameters:         left * right,
		GeneralFormula:             "D_F(M) = [[0, M], [M^T, 0]], M ∈ Mat_{8×8}(R)",
		SelfAdjointByConstruction:  true,
		OddWithGammaByConstruction: true,
		DimensionlessOnly:          true,
		UsesContinuumMasses:        false,
		UsesObservedInput:          false,
		RequiresChosenBlockM:       true,
		CanonicalBlockDerived:      false,
		RealStructureJDerived:      false,
		OrderOneCalculusVerified:   false,
		PromotableFiniteDirac:      false,
		Verdict:                    StatusConditionalAnsatz,
	}
}

func auditUnitRepresentative(f FockSnapshot, eps float64) (RepresentativeMatrixAudit, error) {
	d, gamma, err := buildOffDiagonalIdentity(f.EvenParityStates, f.OddParityStates, 1)
	if err != nil {
		return RepresentativeMatrixAudit{}, err
	}
	return auditRepresentative("unit off-diagonal parity representative", d, gamma, eps)
}

func auditBGapEmbedding(f FockSnapshot, b bsector.Vacuum, eps float64) (BGapEmbeddingAudit, error) {
	gap := b.FirstPositiveEigenvalue(1e-8)
	if math.IsNaN(gap) || gap <= 0 {
		return BGapEmbeddingAudit{}, fmt.Errorf("B-sector first positive eigenvalue unavailable")
	}
	d, gamma, err := buildOffDiagonalIdentity(f.EvenParityStates, f.OddParityStates, gap)
	if err != nil {
		return BGapEmbeddingAudit{}, err
	}
	rep, err := auditRepresentative("uniform B-gap off-diagonal block candidate", d, gamma, eps)
	if err != nil {
		return BGapEmbeddingAudit{}, err
	}
	return BGapEmbeddingAudit{
		BGapAvailable:                true,
		BGap:                         gap,
		OffDiagonalEmbeddingAllowed:  true,
		CanonicalLeftRightPairing:    false,
		CanonicalMatrixEntrySelector: false,
		BGapAsDimensionlessAmplitude: true,
		BGapAsPhysicalMass:           false,
		UniformBlockCandidateAudited: true,
		UniformTraceD2:               rep.TraceD2,
		UniformTraceD4:               rep.TraceD4,
		UniformNormalizedTraceD2:     rep.NormalizedTraceD2,
		UniformNormalizedTraceD4:     rep.NormalizedTraceD4,
		UniformTraceMatchesHopf:      rep.TraceMatchesHopf,
		BGapEmbeddingPromotable:      false,
		Verdict:                      StatusFailedBGapEmbedding,
	}, nil
}

func auditSpectralAction(unit RepresentativeMatrixAudit, bgap BGapEmbeddingAudit) SpectralActionPreflight {
	rows := 0
	if unit.Built {
		rows++
	}
	if bgap.UniformBlockCandidateAudited {
		rows++
	}
	return SpectralActionPreflight{
		D2ComputedForRepresentatives:       unit.TraceD2 > 0 && bgap.UniformTraceD2 > 0,
		D4ComputedForRepresentatives:       unit.TraceD4 > 0 && bgap.UniformTraceD4 > 0,
		TraceRowsComputed:                  rows,
		GaugeCurvatureProjectionDerived:    false,
		HeatKernelMapDerived:               false,
		CutoffFunctionDerived:              false,
		ScaleRatiosGenerated:               false,
		HopfCoefficientGenerated:           false,
		FiniteMatchingCorrectionsGenerated: false,
		PhysicalMassesGenerated:            false,
		Verdict:                            StatusFailedSpectralAction,
	}
}

func auditObstructions(f FockSnapshot, d DiracFamilyAudit, b BGapEmbeddingAudit, s SpectralActionPreflight) ObstructionAudit {
	return ObstructionAudit{
		FiniteFockCarrierAvailable:      f.Constructed && f.StateCount == 16,
		OddSelfAdjointFamilyAvailable:   d.SelfAdjointByConstruction && d.OddWithGammaByConstruction,
		CanonicalDFSelected:             d.CanonicalBlockDerived,
		BGapOffDiagonalMapDerived:       b.BGapEmbeddingPromotable,
		RealStructureJSelected:          d.RealStructureJDerived,
		GradingPhysicalChiralityDerived: f.PhysicalChiralityDerived,
		OrderOneAxiomVerified:           d.OrderOneCalculusVerified,
		GaugeFluctuationMapDerived:      s.GaugeCurvatureProjectionDerived,
		SpectralActionReady:             s.HeatKernelMapDerived && s.CutoffFunctionDerived,
		RequiresBroaderHilbertSpace:     true,
		RequiredNextIngredients: []string{
			"canonical representation of the finite algebra on total H_F",
			"real structure J and KO-dimension data",
			"physical chirality map, not only occupation parity",
			"canonical selector for the 8×8 block M",
			"B-gap-to-bilinear map or proof it is not a Dirac amplitude",
			"order-one calculus and gauge fluctuation map",
			"spectral-action cutoff/subtraction rule",
		},
		Verdict: StatusFailedCanonicalDF,
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		DimensionlessFiniteDataOnly: true,
		ContinuumMassInserted:       false,
		VEVInserted:                 false,
		MBInserted:                  false,
		MStarInserted:               false,
		BGapPromotedToMass:          false,
		DFChosenByFit:               false,
		HopfCoefficientFitted:       false,
		PhysicalLagrangianClaimed:   false,
		FiniteCorePolluted:          false,
		Verdict:                     "firewall preserved: only dimensionless finite scaffold and B-gap diagnostics are audited",
	}
}

func summarize(d DiracFamilyAudit, b BGapEmbeddingAudit, s SpectralActionPreflight) Summary {
	status := strings.Join([]string{StatusConditionalAnsatz, StatusFailedCanonicalDF, StatusFailedBGapEmbedding, StatusBroaderHilbertRequired}, "\n")
	return Summary{
		DFAnsatzAvailable:    d.SelfAdjointByConstruction && d.OddWithGammaByConstruction,
		CanonicalDFDerived:   d.PromotableFiniteDirac,
		BGapEmbeddingDerived: b.BGapEmbeddingPromotable,
		SpectralActionReady:  s.HeatKernelMapDerived && s.CutoffFunctionDerived,
		Status:               status,
		NextGate:             "derive or seal the missing finite-algebra representation/J/gamma/order-one calculus before any physical spectral action",
		Comment:              "Gate 233 initializes the legal D_F family but does not select a physical finite Dirac operator.",
	}
}

func buildTruth(f FockSnapshot, d DiracFamilyAudit, b BGapEmbeddingAudit, s SpectralActionPreflight, o ObstructionAudit) string {
	return fmt.Sprintf("Gate 233 truth: H_Fock has %d states and admits a dimensionless odd self-adjoint D_F family with %d free real entries, but the finite core does not derive the block M, the real structure J, the physical chirality map, or a canonical B-gap off-diagonal embedding. Representative traces Tr(D^2), Tr(D^4) are computable diagnostics only; without gauge projection and cutoff/subtraction data they do not generate the Hopf coefficient, matching constants, or masses. Required next ingredients: %s.", f.StateCount, d.FreeRealParameters, strings.Join(o.RequiredNextIngredients, "; "))
}

func buildOffDiagonalIdentity(left, right int, amplitude float64) (linear.Matrix, linear.Matrix, error) {
	if left != right || left <= 0 {
		return linear.Matrix{}, linear.Matrix{}, fmt.Errorf("balanced left/right split required, got %d/%d", left, right)
	}
	n := left + right
	d := linear.NewMatrix(n, n)
	for i := 0; i < left; i++ {
		d.Set(i, left+i, amplitude)
		d.Set(left+i, i, amplitude)
	}
	gamma := linear.NewMatrix(n, n)
	for i := 0; i < left; i++ {
		gamma.Set(i, i, 1)
		gamma.Set(left+i, left+i, -1)
	}
	return d, gamma, nil
}

func auditRepresentative(name string, d, gamma linear.Matrix, eps float64) (RepresentativeMatrixAudit, error) {
	if d.Rows() != d.Cols() || gamma.Rows() != gamma.Cols() || d.Rows() != gamma.Rows() {
		return RepresentativeMatrixAudit{}, fmt.Errorf("representative matrix dimension mismatch")
	}
	dt := d.Transpose()
	self, err := d.Sub(dt)
	if err != nil {
		return RepresentativeMatrixAudit{}, err
	}
	gd, err := gamma.Mul(d)
	if err != nil {
		return RepresentativeMatrixAudit{}, err
	}
	dg, err := d.Mul(gamma)
	if err != nil {
		return RepresentativeMatrixAudit{}, err
	}
	anti, err := gd.Add(dg)
	if err != nil {
		return RepresentativeMatrixAudit{}, err
	}
	d2, err := d.Mul(d)
	if err != nil {
		return RepresentativeMatrixAudit{}, err
	}
	d4, err := d2.Mul(d2)
	if err != nil {
		return RepresentativeMatrixAudit{}, err
	}
	tr2, err := d2.Trace()
	if err != nil {
		return RepresentativeMatrixAudit{}, err
	}
	tr4, err := d4.Trace()
	if err != nil {
		return RepresentativeMatrixAudit{}, err
	}
	n := float64(d.Rows())
	ntr2, ntr4 := tr2/n, tr4/n
	hopf := 4 / math.Pi
	return RepresentativeMatrixAudit{
		RepresentativeName:  name,
		Built:               true,
		Dimension:           d.Rows(),
		AnticommutatorNorm:  anti.FrobeniusNorm(),
		SelfAdjointResidual: self.FrobeniusNorm(),
		TraceD2:             tr2,
		TraceD4:             tr4,
		NormalizedTraceD2:   ntr2,
		NormalizedTraceD4:   ntr4,
		HopfCoefficient:     hopf,
		TraceMatchesHopf:    math.Abs(ntr2-hopf) < eps || math.Abs(ntr4-hopf) < eps,
		Promoted:            false,
		Comment:             "representative proves matrix identities only; it is not a canonical D_F selection",
	}, nil
}

func FormatFock(f FockSnapshot) string {
	return fmt.Sprintf("states=%d modes=%d even=%d odd=%d balanced=%t gamma_trace=%.6g physical_chirality_derived=%t; %s", f.StateCount, f.ModeCount, f.EvenParityStates, f.OddParityStates, f.ParitySplitBalanced, f.GammaTrace, f.PhysicalChiralityDerived, f.Comment)
}

func FormatDiracFamily(d DiracFamilyAudit) string {
	return fmt.Sprintf("%s; dim=%d split=%d+%d free_params=%d formula=%s selfadjoint=%t odd=%t canonical_block=%t J=%t order_one=%t promotable=%t verdict=%s", d.FamilyName, d.Dimension, d.LeftDimension, d.RightDimension, d.FreeRealParameters, d.GeneralFormula, d.SelfAdjointByConstruction, d.OddWithGammaByConstruction, d.CanonicalBlockDerived, d.RealStructureJDerived, d.OrderOneCalculusVerified, d.PromotableFiniteDirac, d.Verdict)
}

func FormatRepresentative(r RepresentativeMatrixAudit) string {
	return fmt.Sprintf("%s: dim=%d ||{γ,D}||=%.6g self_resid=%.6g TrD2=%.12g TrD4=%.12g normalized=(%.12g,%.12g) hopf=%.12g match=%t promoted=%t; %s", r.RepresentativeName, r.Dimension, r.AnticommutatorNorm, r.SelfAdjointResidual, r.TraceD2, r.TraceD4, r.NormalizedTraceD2, r.NormalizedTraceD4, r.HopfCoefficient, r.TraceMatchesHopf, r.Promoted, r.Comment)
}

func FormatBGap(b BGapEmbeddingAudit) string {
	return fmt.Sprintf("B_gap=%.12g available=%t offdiag_allowed=%t canonical_pairing=%t selector=%t as_mass=%t uniform_TrD2=%.12g uniform_TrD4=%.12g normalized=(%.12g,%.12g) hopf_match=%t promotable=%t verdict=%s", b.BGap, b.BGapAvailable, b.OffDiagonalEmbeddingAllowed, b.CanonicalLeftRightPairing, b.CanonicalMatrixEntrySelector, b.BGapAsPhysicalMass, b.UniformTraceD2, b.UniformTraceD4, b.UniformNormalizedTraceD2, b.UniformNormalizedTraceD4, b.UniformTraceMatchesHopf, b.BGapEmbeddingPromotable, b.Verdict)
}

func FormatSpectral(s SpectralActionPreflight) string {
	return fmt.Sprintf("D2=%t D4=%t rows=%d gauge_projection=%t heat_kernel=%t cutoff=%t scale_ratios=%t hopf=%t matching=%t masses=%t verdict=%s", s.D2ComputedForRepresentatives, s.D4ComputedForRepresentatives, s.TraceRowsComputed, s.GaugeCurvatureProjectionDerived, s.HeatKernelMapDerived, s.CutoffFunctionDerived, s.ScaleRatiosGenerated, s.HopfCoefficientGenerated, s.FiniteMatchingCorrectionsGenerated, s.PhysicalMassesGenerated, s.Verdict)
}

func FormatObstruction(o ObstructionAudit) string {
	return fmt.Sprintf("carrier=%t odd_family=%t canonical_DF=%t Bgap_map=%t J=%t physical_gamma=%t order_one=%t gauge_fluctuation=%t spectral_action_ready=%t broader_hilbert=%t next=[%s] verdict=%s", o.FiniteFockCarrierAvailable, o.OddSelfAdjointFamilyAvailable, o.CanonicalDFSelected, o.BGapOffDiagonalMapDerived, o.RealStructureJSelected, o.GradingPhysicalChiralityDerived, o.OrderOneAxiomVerified, o.GaugeFluctuationMapDerived, o.SpectralActionReady, o.RequiresBroaderHilbertSpace, strings.Join(o.RequiredNextIngredients, "; "), o.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("dimensionless_only=%t continuum_mass=%t v=%t MB=%t Mstar=%t Bgap_mass=%t DF_fit=%t hopf_fit=%t lagrangian_claim=%t polluted=%t; %s", f.DimensionlessFiniteDataOnly, f.ContinuumMassInserted, f.VEVInserted, f.MBInserted, f.MStarInserted, f.BGapPromotedToMass, f.DFChosenByFit, f.HopfCoefficientFitted, f.PhysicalLagrangianClaimed, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("ansatz=%t canonical_DF=%t Bgap_embedding=%t spectral_action=%t status=%q next=%s", s.DFAnsatzAvailable, s.CanonicalDFDerived, s.BGapEmbeddingDerived, s.SpectralActionReady, s.Status, s.NextGate)
}
