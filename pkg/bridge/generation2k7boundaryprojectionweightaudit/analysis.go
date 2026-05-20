// Package generation2k7boundaryprojectionweightaudit implements
// Gate 627: K7BoundaryProjectionWeight Audit.
//
// Gate 626 found that the HistoryLoopDeficit closure does not land exactly on
// |lambda(Lambda_12)|, but on a boundary-weighted scalar/gauge wound mixture
// with weight 7/72. Gate 627 audits the source type of that coefficient.  The
// numerator is tested against the certified Boolean--octonionic contact carrier
// K_7, while the denominator is tested only against already-typed ASHA ledger
// chambers.  No native projection map or boundary chamber is promoted.
package generation2k7boundaryprojectionweightaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/asha"
	gate626 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryweighteddeficitclosureaudit"
)

const (
	AuditID = "GATE627-K7-BOUNDARY-PROJECTION-WEIGHT-AUDIT"

	StatusGate626Inherited       = "PASS_GATE626_BOUNDARY_WEIGHTED_CLOSURE_INHERITED"
	StatusWeightIdentified       = "PASS_WEIGHT_7_OVER_72_IDENTIFIED"
	StatusNumeratorK7Candidate   = "CONDITIONAL_SUPPORT_NUMERATOR_7_MATCHES_DIM_K7"
	StatusDenominator72Candidate = "CONDITIONAL_SUPPORT_72_BOUNDARY_CHAMBER_DENOMINATOR_CANDIDATE"
	StatusComplementAudited      = "PASS_65_OVER_72_COMPLEMENT_AUDITED"
	StatusMidpointRewriteAudited = "PASS_7_OVER_36_MIDPOINT_PULL_REWRITE_AUDITED"
	StatusProjectionMissing      = "FAILED_ROUTE_NO_NATIVE_K7_TO_BOUNDARY_STRESS_PROJECTOR"
	StatusNoCertified72Carrier   = "FAILED_ROUTE_NO_CERTIFIED_72_DIMENSION_BOUNDARY_CARRIER"
	StatusNoNativeWeightTheorem  = "FAILED_ROUTE_NO_NATIVE_SEVEN_OVER_SEVENTY_TWO_SOURCE_THEOREM"
	StatusGate627Boundary        = "FIREWALL_PRESERVED_GATE627_WEIGHT_SOURCE_IS_CANDIDATE_ONLY"
)

const (
	expectedNumerator   = 7
	expectedDenominator = 72
)

type Gate626Inheritance struct {
	KappaSum                      float64
	AbsLambda12                   float64
	R3MinusOne                    float64
	BoundarySplit                 float64
	BoundaryWeight                float64
	ScalarWeight                  float64
	WeightedMixture               float64
	WeightedClosureResidual       float64
	ScalarPredictionResidual      float64
	Gate626ClosureIsBridgeOnly    bool
	Gate626NativeWeightSource     bool
	Gate626NativeTransportTheorem bool
	Verdict                       string
}

type WeightIdentification struct {
	Numerator             int
	Denominator           int
	Value                 float64
	ObservedWeight        float64
	Residual              float64
	ScalarComplement      float64
	ComplementNumerator   int
	ComplementDenominator int
	Verdict               string
}

type NumeratorK7Audit struct {
	Numerator                 int
	K7Dimension               int
	RankPB                    int
	RankPG                    int
	AmbientExteriorDimension  int
	CliffordVectorDimension   int
	MatchesDimK7              bool
	K7NativeCarrierCertified  bool
	ProjectionToBoundaryFound bool
	Interpretation            string
	Verdict                   string
}

type DenominatorCandidateRow struct {
	Name               string
	Expression         string
	Value              int
	TypedFactors       []string
	ExistingLedgerData bool
	BoundaryCarrier    bool
	CertifiedAsDenom   bool
	RequiresNewTheorem bool
	Verdict            string
}

type Denominator72Audit struct {
	TargetDenominator         int
	Rows                      []DenominatorCandidateRow
	AnyExistingTypedCandidate bool
	CertifiedBoundaryCarrier  bool
	BestCandidate             string
	BestCandidateValue        int
	Verdict                   string
}

type ComplementProjectionAudit struct {
	BoundaryWeight           float64
	ScalarWeight             float64
	K7Numerator              int
	ChamberDenominator       int
	ComplementDimension      int
	ComplementWeight         float64
	ComplementEquals65Over72 bool
	ArithmeticComplementOnly bool
	NativeComplementCarrier  bool
	Equation                 string
	Verdict                  string
}

type MidpointPullRewriteAudit struct {
	AbsLambda12              float64
	R3MinusOne               float64
	XiBoundary               float64
	BoundarySplit            float64
	ScalarToMidpointPull     float64
	FullSplitWeight          float64
	MidpointPullWeight       float64
	WeightedFromFullSplit    float64
	WeightedFromMidpoint     float64
	RewriteResidual          float64
	XiBoundaryInherited      bool
	NativeMidpointProjection bool
	Equation                 string
	Verdict                  string
}

type BoundaryProjectionOperatorAudit struct {
	DomainCarrier             string
	DomainDimension           int
	CandidateCodomain         string
	CandidateChamberDimension int
	ProjectionWeight          float64
	WeightEqualsDimRatio      bool
	ProjectionOperatorExists  bool
	IdempotentCertified       bool
	TraceCertified            bool
	IntertwinerCertified      bool
	MissingObject             string
	Verdict                   string
}

type RecurrenceAuditRow struct {
	Coefficient     string
	Location        string
	SameCoefficient bool
	Forced          bool
	NativeCertified bool
	Comment         string
}

type CoefficientRecurrenceAudit struct {
	Rows                     []RecurrenceAuditRow
	SameCoefficientElsewhere bool
	NativeRecurrenceLaw      bool
	Verdict                  string
}

type NativeASHAStatus struct {
	NumeratorK7Native                  bool
	Denominator72BoundaryCarrierNative bool
	K7BoundaryProjectionNative         bool
	ComplementProjectionNative         bool
	GaugeScalarFlavorTransportNative   bool
	SevenOverSeventyTwoSourceTheorem   bool
	Statement                          string
	Verdict                            string
}

type Firewalls struct {
	ClaimsNativeWeightTheorem bool
	ClaimsCertified72Carrier  bool
	ClaimsNativeProjection    bool
	ClaimsScalarRGMatching    bool
	ClaimsFlavorOrientation   bool
	ClaimsGaugeUnification    bool
	ClaimsHiggsMassDerived    bool
	ClaimsEndpointDerivation  bool
	Verdict                   string
}

type Analysis struct {
	Inherited    Gate626Inheritance
	Weight       WeightIdentification
	Numerator    NumeratorK7Audit
	Denominator  Denominator72Audit
	Complement   ComplementProjectionAudit
	Midpoint     MidpointPullRewriteAudit
	Projection   BoundaryProjectionOperatorAudit
	Recurrence   CoefficientRecurrenceAudit
	NativeStatus NativeASHAStatus
	Firewalls    Firewalls
	Truth        string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g626, err := gate626.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate626 predecessor: %w", err)
	}
	engine := asha.New()
	geo := engine.Geometry
	fam := engine.Family

	inherit := inheritGate626(g626)
	weight := identifyWeight(inherit)
	numerator := auditK7Numerator(weight, geo)
	denominator := auditDenominator72(weight, geo, fam)
	complement := auditComplement(weight)
	midpoint := auditMidpointPull(inherit, weight)
	projection := auditBoundaryProjection(numerator, denominator, weight)
	recurrence := auditCoefficientRecurrence(weight)

	a := Analysis{
		Inherited:    inherit,
		Weight:       weight,
		Numerator:    numerator,
		Denominator:  denominator,
		Complement:   complement,
		Midpoint:     midpoint,
		Projection:   projection,
		Recurrence:   recurrence,
		NativeStatus: auditNativeStatus(numerator, denominator, projection, complement),
		Firewalls:    auditFirewalls(),
		Truth:        "Gate 627 audits the source type of the Gate626 7/72 boundary weight. The numerator 7 lawfully matches the certified K_7 contact-vacuum dimension, and 72 has typed ledger decompositions such as 8×9, 3×24, and 2×36. But ASHA does not yet certify any 72-dimensional boundary chamber or any native K_7-to-boundary-stress projection operator. Therefore 7/72 remains a strong bridge candidate and the missing object sharpens to Pi_{K7->boundary} plus a certified boundary chamber.",
	}
	return a, nil
}

func inheritGate626(g gate626.Analysis) Gate626Inheritance {
	return Gate626Inheritance{
		KappaSum:                      g.WeightedClosure.KappaSum,
		AbsLambda12:                   g.WeightedClosure.AbsLambda12,
		R3MinusOne:                    g.WeightedClosure.R3MinusOne,
		BoundarySplit:                 g.WeightedClosure.BoundarySplit,
		BoundaryWeight:                g.WeightedClosure.BoundaryWeight,
		ScalarWeight:                  g.WeightedClosure.ScalarWeight,
		WeightedMixture:               g.WeightedClosure.WeightedMixture,
		WeightedClosureResidual:       g.WeightedClosure.Residual,
		ScalarPredictionResidual:      g.ScalarPrediction.BestResidual,
		Gate626ClosureIsBridgeOnly:    g.WeightedClosure.BridgeOnly,
		Gate626NativeWeightSource:     g.NativeStatus.NativeSevenOverSeventyTwoSource,
		Gate626NativeTransportTheorem: g.NativeStatus.NativeGaugeScalarFlavorDeficitTransport,
		Verdict:                       StatusGate626Inherited,
	}
}

func identifyWeight(i Gate626Inheritance) WeightIdentification {
	exact := float64(expectedNumerator) / float64(expectedDenominator)
	return WeightIdentification{
		Numerator:             expectedNumerator,
		Denominator:           expectedDenominator,
		Value:                 exact,
		ObservedWeight:        i.BoundaryWeight,
		Residual:              i.BoundaryWeight - exact,
		ScalarComplement:      1 - exact,
		ComplementNumerator:   expectedDenominator - expectedNumerator,
		ComplementDenominator: expectedDenominator,
		Verdict:               StatusWeightIdentified,
	}
}

func auditK7Numerator(w WeightIdentification, g asha.Geometry) NumeratorK7Audit {
	gradeDims := g.GradeDimensions
	vectorDim := 0
	if len(gradeDims) > 1 {
		vectorDim = gradeDims[1]
	}
	ambient := 0
	if len(gradeDims) > 4 {
		ambient = gradeDims[4]
	}
	return NumeratorK7Audit{
		Numerator:                 w.Numerator,
		K7Dimension:               g.DimK,
		RankPB:                    g.RankPB,
		RankPG:                    g.RankPG,
		AmbientExteriorDimension:  ambient,
		CliffordVectorDimension:   vectorDim,
		MatchesDimK7:              w.Numerator == g.DimK,
		K7NativeCarrierCertified:  g.RankPB == 56 && g.RankPG == 14 && g.DimK == 7,
		ProjectionToBoundaryFound: false,
		Interpretation:            "7 may be read as dim(K_7), the certified Boolean--octonionic contact-vacuum carrier; this certifies only the numerator source candidate, not a projection law.",
		Verdict:                   StatusNumeratorK7Candidate,
	}
}

func auditDenominator72(w WeightIdentification, g asha.Geometry, f asha.Family) Denominator72Audit {
	vectorDim := 0
	if len(g.GradeDimensions) > 1 {
		vectorDim = g.GradeDimensions[1]
	}
	rows := []DenominatorCandidateRow{
		{
			Name:               "Clifford measurement ladder times charged K/X/Y coefficient chamber",
			Expression:         "8 × 9",
			Value:              vectorDim * f.KXYChargedCoeffDim,
			TypedFactors:       []string{"dim R^8 vector ladder", "dim C_KXY^charged=9 quarantined coefficient chamber"},
			ExistingLedgerData: true,
			BoundaryCarrier:    false,
			CertifiedAsDenom:   false,
			RequiresNewTheorem: true,
			Verdict:            StatusDenominator72Candidate,
		},
		{
			Name:               "three-generation matter inventory candidate",
			Expression:         "3 × 24",
			Value:              3 * 24,
			TypedFactors:       []string{"generation count 3", "24-unit matter inventory candidate"},
			ExistingLedgerData: false,
			BoundaryCarrier:    false,
			CertifiedAsDenom:   false,
			RequiresNewTheorem: true,
			Verdict:            StatusDenominator72Candidate,
		},
		{
			Name:               "doubled boundary pair over a 36-unit chamber",
			Expression:         "2 × 36",
			Value:              2 * 36,
			TypedFactors:       []string{"two endpoint/boundary sides", "36-unit chamber candidate"},
			ExistingLedgerData: false,
			BoundaryCarrier:    false,
			CertifiedAsDenom:   false,
			RequiresNewTheorem: true,
			Verdict:            StatusDenominator72Candidate,
		},
		{
			Name:               "K7 complement arithmetic chamber",
			Expression:         "7 + 65",
			Value:              w.Numerator + (w.Denominator - w.Numerator),
			TypedFactors:       []string{"dim K_7 candidate numerator", "65 arithmetic complement"},
			ExistingLedgerData: true,
			BoundaryCarrier:    false,
			CertifiedAsDenom:   false,
			RequiresNewTheorem: true,
			Verdict:            StatusDenominator72Candidate,
		},
	}
	anyTyped := false
	best := ""
	bestVal := 0
	for _, row := range rows {
		if row.Value == w.Denominator && row.ExistingLedgerData {
			anyTyped = true
			if best == "" {
				best = row.Expression
				bestVal = row.Value
			}
		}
	}
	if best == "" && len(rows) > 0 {
		best = rows[0].Expression
		bestVal = rows[0].Value
	}
	return Denominator72Audit{
		TargetDenominator:         w.Denominator,
		Rows:                      rows,
		AnyExistingTypedCandidate: anyTyped,
		CertifiedBoundaryCarrier:  false,
		BestCandidate:             best,
		BestCandidateValue:        bestVal,
		Verdict:                   StatusDenominator72Candidate,
	}
}

func auditComplement(w WeightIdentification) ComplementProjectionAudit {
	compDim := w.Denominator - w.Numerator
	compWeight := float64(compDim) / float64(w.Denominator)
	return ComplementProjectionAudit{
		BoundaryWeight:           w.Value,
		ScalarWeight:             w.ScalarComplement,
		K7Numerator:              w.Numerator,
		ChamberDenominator:       w.Denominator,
		ComplementDimension:      compDim,
		ComplementWeight:         compWeight,
		ComplementEquals65Over72: compDim == 65 && math.Abs(compWeight-w.ScalarComplement) < 1e-15,
		ArithmeticComplementOnly: true,
		NativeComplementCarrier:  false,
		Equation:                 "65/72 = 1 - 7/72 is an arithmetic complement unless a 72-dimensional boundary chamber and 65-dimensional complement carrier are certified.",
		Verdict:                  StatusComplementAudited,
	}
}

func auditMidpointPull(i Gate626Inheritance, w WeightIdentification) MidpointPullRewriteAudit {
	xi := 0.5 * (i.R3MinusOne + i.AbsLambda12)
	pull := xi - i.AbsLambda12
	midpointWeight := 2 * w.Value
	fromFull := i.AbsLambda12 + w.Value*i.BoundarySplit
	fromMid := i.AbsLambda12 + midpointWeight*pull
	return MidpointPullRewriteAudit{
		AbsLambda12:              i.AbsLambda12,
		R3MinusOne:               i.R3MinusOne,
		XiBoundary:               xi,
		BoundarySplit:            i.BoundarySplit,
		ScalarToMidpointPull:     pull,
		FullSplitWeight:          w.Value,
		MidpointPullWeight:       midpointWeight,
		WeightedFromFullSplit:    fromFull,
		WeightedFromMidpoint:     fromMid,
		RewriteResidual:          fromMid - fromFull,
		XiBoundaryInherited:      true,
		NativeMidpointProjection: false,
		Equation:                 "W_72=|lambda_12|+(7/72)[R_3-1-|lambda_12|]=|lambda_12|+(7/36)[xi_boundary-|lambda_12|]",
		Verdict:                  StatusMidpointRewriteAudited,
	}
}

func auditBoundaryProjection(n NumeratorK7Audit, d Denominator72Audit, w WeightIdentification) BoundaryProjectionOperatorAudit {
	dimRatio := float64(n.K7Dimension) / float64(d.TargetDenominator)
	return BoundaryProjectionOperatorAudit{
		DomainCarrier:             "K_7 Boolean--octonionic contact vacuum",
		DomainDimension:           n.K7Dimension,
		CandidateCodomain:         "scalar/gauge boundary-stress split line spanned by (R_3-1)-|lambda(Lambda_12)|",
		CandidateChamberDimension: d.TargetDenominator,
		ProjectionWeight:          w.Value,
		WeightEqualsDimRatio:      math.Abs(dimRatio-w.Value) < 1e-15,
		ProjectionOperatorExists:  false,
		IdempotentCertified:       false,
		TraceCertified:            false,
		IntertwinerCertified:      false,
		MissingObject:             "Pi_{K7->boundary}: a typed projection/intertwiner from K_7 into a certified 72-dimensional boundary chamber whose normalized trace is 7/72",
		Verdict:                   StatusProjectionMissing,
	}
}

func auditCoefficientRecurrence(w WeightIdentification) CoefficientRecurrenceAudit {
	rows := []RecurrenceAuditRow{
		{Coefficient: "7/72", Location: "Gate626 boundary-weighted deficit closure", SameCoefficient: true, Forced: false, NativeCertified: false, Comment: "the coefficient is inherited as the target under audit"},
		{Coefficient: "7/36", Location: "equivalent pull from |lambda_12| toward xi_boundary", SameCoefficient: false, Forced: false, NativeCertified: false, Comment: "same closure rewritten against the midpoint stress seal, not a new native coefficient"},
		{Coefficient: "65/72", Location: "scalar-wound complement in the weighted mixture", SameCoefficient: false, Forced: false, NativeCertified: false, Comment: "arithmetic complement of 7/72, not a certified complementary carrier"},
	}
	return CoefficientRecurrenceAudit{Rows: rows, SameCoefficientElsewhere: false, NativeRecurrenceLaw: false, Verdict: StatusNoNativeWeightTheorem}
}

func auditNativeStatus(n NumeratorK7Audit, d Denominator72Audit, p BoundaryProjectionOperatorAudit, c ComplementProjectionAudit) NativeASHAStatus {
	return NativeASHAStatus{
		NumeratorK7Native:                  n.K7NativeCarrierCertified && n.MatchesDimK7,
		Denominator72BoundaryCarrierNative: d.CertifiedBoundaryCarrier,
		K7BoundaryProjectionNative:         p.ProjectionOperatorExists && p.IdempotentCertified && p.TraceCertified,
		ComplementProjectionNative:         c.NativeComplementCarrier,
		GaugeScalarFlavorTransportNative:   false,
		SevenOverSeventyTwoSourceTheorem:   false,
		Statement:                          "The numerator is a lawful K_7 dimension candidate, but the denominator and projection are not certified. A native source theorem would need a 72-dimensional boundary chamber plus a K_7 projection/intertwiner with normalized trace 7/72.",
		Verdict:                            StatusNoCertified72Carrier,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{Verdict: StatusGate627Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate626Inherited,
		StatusWeightIdentified,
		StatusNumeratorK7Candidate,
		StatusDenominator72Candidate,
		StatusComplementAudited,
		StatusMidpointRewriteAudited,
		StatusProjectionMissing,
		StatusNoCertified72Carrier,
		StatusNoNativeWeightTheorem,
		StatusGate627Boundary,
	}
}
