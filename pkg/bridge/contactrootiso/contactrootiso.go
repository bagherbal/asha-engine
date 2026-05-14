// Package contactrootiso implements Gate 150: exact contact
// root-isolation / row-wise eigenprojector assignment theorem.
//
// Gate 149 certified the exact rational contact-overlap matrix and its exact
// characteristic polynomial. Gate 150 upgrades the spectral side by isolating
// the seven non-unit contact roots with exact rational certificates. Three
// roots are rational factors and four roots are isolated by exact sign-change
// intervals for the certified quartic factor.
//
// The gate deliberately does not claim physical row semantics. Exact root
// isolation is not the same as an exact eigenprojector construction in the
// quartic number field, and neither gives T3R, B-L, hypercharge, local fields,
// mass activation, decoupling, or threshold beta permission.
package contactrootiso

import (
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactmatrixcert"
)

type IsolatedRoot struct {
	Label           string
	Factor          string
	Kind            string
	Point           string
	IntervalLeft    string
	IntervalRight   string
	LeftSign        int
	RightSign       int
	SignChange      bool
	Multiplicity    int
	Certified       bool
	Approx          string
	PhysicalMeaning string
}

type QuarticIsolation struct {
	Polynomial             string
	Intervals              []IsolatedRoot
	IntervalsDisjoint      bool
	SignChanges            int
	Degree                 int
	AllQuarticRootsCovered bool
	Verdict                string
}

type RootIsolationCertificate struct {
	ExactCharpolyInherited bool
	RationalRootsCertified int
	QuarticRootsCertified  int
	TotalPartialRoots      int
	UnitRootMultiplicity   int
	RootIsolationComplete  bool
	OrderedRoots           []IsolatedRoot
	Verdict                string
}

type EigenprojectorAudit struct {
	ExactMatrixCertificateInherited bool
	ExactRootIsolationAvailable     bool
	ExactNumberFieldProjectors      int
	RationalProjectors              int
	QuarticProjectors               int
	RowwiseProjectorAssignments     int
	RootToContactRowSemantics       int
	ChargeSemanticRows              int
	RepresentationRows              int
	BetaRowsAllowed                 int
	Verdict                         string
}

type ConstructionRequirements struct {
	ExactRationalMatrix           bool
	ExactCharacteristicPolynomial bool
	ExactRootIsolation            bool
	ExactNumberFieldArithmetic    bool
	EigenprojectorFormulaInField  bool
	RowwiseProjectorAssignment    bool
	ChargeOperatorSelected        bool
	RepresentationRowsSelected    bool
	ObservedInputFree             bool
	AllSatisfiedForPhysics        bool
	Verdict                       string
}

type Summary struct {
	ContactRows                int
	ExactMatrixCertificates    int
	ExactCharpolyCertificates  int
	RootIsolationCertificates  int
	RationalRootCertificates   int
	QuarticRootCertificates    int
	UnitEigenMultiplicity      int
	RowAssignmentProofs        int
	ChargeSemanticRows         int
	RepresentationCompleteRows int
	RepresentationOpenRows     int
	ContactBetaRowsAllowed     int
	ContactZeroRowsProved      int
	ResidualS6Choices          int
	ResidualNullityBefore      int
	ResidualNullityAfter       int
}

type Analysis struct {
	Previous contactmatrixcert.Analysis

	Quartic      QuarticIsolation
	Certificate  RootIsolationCertificate
	Projectors   EigenprojectorAudit
	Requirements ConstructionRequirements
	Summary      Summary

	ContactRows                  int
	ExactRationalOverlapMatrix   bool
	ExactCharacteristicCertified bool
	ExactRootIsolationCertified  bool
	ExactNumberFieldProjectors   int
	RootIsolationCertificates    int
	RowwiseRootAssignmentProofs  int
	ChargeSemanticRows           int
	T3RRowsDerived               int
	ChiralityRowsDerived         int
	BMinusLRowsDerived           int
	SU2LRowsDerived              int
	HyperchargeRowsDerived       int
	RepresentationCompleteRows   int
	RepresentationOpenRows       int
	ContactBetaRowsAllowed       int
	ContactZeroRowsProved        int
	BetaPermissionFirewallClosed bool
	ThresholdCorrectedBeta       bool
	FullBetaMatchingTensor       bool
	ResidualS6Choices            int
	ResidualNullityBefore        int
	ResidualNullityAfter         int
	HiddenObservedInputUsed      bool
	PhysicalWeakAngleDerived     bool
	FineStructureDerived         bool
	PhysicalMassesDerived        bool
	PhysicalScaleDerived         bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var defaultOnce sync.Once
var defaultValue Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := contactmatrixcert.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactmatrixcert.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || !prev.ExactRationalOverlapMatrix || !prev.ExactCharacteristicCertified || !prev.ExactAnnihilationCertified || prev.RootIsolationCertificates != 0 || prev.RowwiseRootAssignmentProofs != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 150 requires Gate 149 exact matrix/charpoly certificate with root isolation still absent")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 150 refuses hidden observed physical input")
	}

	rational := []IsolatedRoot{
		{Label: "partial rational root r1", Factor: "3x-1", Kind: "rational", Point: "1/3", IntervalLeft: "1/3", IntervalRight: "1/3", LeftSign: 0, RightSign: 0, SignChange: false, Multiplicity: 1, Certified: true, Approx: "0.3333333333", PhysicalMeaning: "exact rational contact spectral root, diagnostic only"},
		{Label: "partial rational root r2", Factor: "2x-1", Kind: "rational", Point: "1/2", IntervalLeft: "1/2", IntervalRight: "1/2", LeftSign: 0, RightSign: 0, SignChange: false, Multiplicity: 1, Certified: true, Approx: "0.5000000000", PhysicalMeaning: "exact rational contact spectral root, diagnostic only"},
		{Label: "partial rational root r3", Factor: "3x-2", Kind: "rational", Point: "2/3", IntervalLeft: "2/3", IntervalRight: "2/3", LeftSign: 0, RightSign: 0, SignChange: false, Multiplicity: 1, Certified: true, Approx: "0.6666666667", PhysicalMeaning: "exact rational contact spectral root, diagnostic only"},
	}

	quarticIntervals := []struct{ label, left, right, approx string }{
		{"quartic low root q1", "2839/10000", "2840/10000", "0.2839121926"},
		{"quartic mid-low root q2", "4411/10000", "4412/10000", "0.4411227573"},
		{"quartic mid-high root q3", "7440/10000", "7441/10000", "0.7440966380"},
		{"quartic high root q4", "8975/10000", "8976/10000", "0.8975350788"},
	}
	quarticRoots := make([]IsolatedRoot, 0, len(quarticIntervals))
	signChanges := 0
	for _, in := range quarticIntervals {
		l := mustRat(in.left)
		r := mustRat(in.right)
		ls := quarticEval(l).Sign()
		rs := quarticEval(r).Sign()
		change := ls*rs < 0
		if change {
			signChanges++
		}
		quarticRoots = append(quarticRoots, IsolatedRoot{
			Label:           in.label,
			Factor:          "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
			Kind:            "quartic-isolated",
			IntervalLeft:    in.left,
			IntervalRight:   in.right,
			LeftSign:        ls,
			RightSign:       rs,
			SignChange:      change,
			Multiplicity:    1,
			Certified:       change,
			Approx:          in.approx,
			PhysicalMeaning: "isolated exact quartic spectral root, diagnostic only",
		})
	}

	quartic := QuarticIsolation{
		Polynomial:             "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		Intervals:              quarticRoots,
		IntervalsDisjoint:      intervalsDisjoint(quarticRoots),
		SignChanges:            signChanges,
		Degree:                 4,
		AllQuarticRootsCovered: intervalsDisjoint(quarticRoots) && signChanges == 4,
		Verdict:                "four disjoint rational intervals each contain a sign change of the exact quartic factor; since the factor has degree four, the four quartic roots are isolated one per interval",
	}

	ordered := []IsolatedRoot{quarticRoots[0], rational[0], quarticRoots[1], rational[1], rational[2], quarticRoots[2], quarticRoots[3]}
	rootIsolationComplete := quartic.AllQuarticRootsCovered && allCertified(rational) && allCertified(quarticRoots)
	cert := RootIsolationCertificate{
		ExactCharpolyInherited: true,
		RationalRootsCertified: len(rational),
		QuarticRootsCertified:  signChanges,
		TotalPartialRoots:      7,
		UnitRootMultiplicity:   prev.Summary.UnitEigenMultiplicity,
		RootIsolationComplete:  rootIsolationComplete,
		OrderedRoots:           ordered,
		Verdict:                "the seven non-unit contact spectral roots are isolated exactly: three rational roots plus four quartic roots in disjoint rational sign-change intervals",
	}

	projectors := EigenprojectorAudit{
		ExactMatrixCertificateInherited: true,
		ExactRootIsolationAvailable:     rootIsolationComplete,
		ExactNumberFieldProjectors:      0,
		RationalProjectors:              0,
		QuarticProjectors:               0,
		RowwiseProjectorAssignments:     0,
		RootToContactRowSemantics:       0,
		ChargeSemanticRows:              0,
		RepresentationRows:              0,
		BetaRowsAllowed:                 0,
		Verdict:                         "root isolation orders the spectral roots but does not construct exact number-field eigenprojectors or assign roots to physical contact rows",
	}

	req := ConstructionRequirements{
		ExactRationalMatrix:           prev.ExactRationalOverlapMatrix,
		ExactCharacteristicPolynomial: prev.ExactCharacteristicCertified,
		ExactRootIsolation:            rootIsolationComplete,
		ExactNumberFieldArithmetic:    false,
		EigenprojectorFormulaInField:  false,
		RowwiseProjectorAssignment:    false,
		ChargeOperatorSelected:        false,
		RepresentationRowsSelected:    false,
		ObservedInputFree:             true,
		AllSatisfiedForPhysics:        false,
		Verdict:                       "root isolation is certified, but exact number-field projectors, row-wise semantic assignments, charge operators, representation rows, local fields, mass activation, and decoupling are not derived",
	}

	summary := Summary{
		ContactRows:                prev.ContactRows,
		ExactMatrixCertificates:    1,
		ExactCharpolyCertificates:  1,
		RootIsolationCertificates:  boolInt(rootIsolationComplete) * 7,
		RationalRootCertificates:   len(rational),
		QuarticRootCertificates:    signChanges,
		UnitEigenMultiplicity:      prev.Summary.UnitEigenMultiplicity,
		RowAssignmentProofs:        0,
		ChargeSemanticRows:         0,
		RepresentationCompleteRows: 0,
		RepresentationOpenRows:     prev.RepresentationOpenRows,
		ContactBetaRowsAllowed:     0,
		ContactZeroRowsProved:      0,
		ResidualS6Choices:          prev.ResidualS6Choices,
		ResidualNullityBefore:      prev.ResidualNullityAfter,
		ResidualNullityAfter:       prev.ResidualNullityAfter,
	}

	truth := "Gate 150 upgrades the exact contact spectral certificate by isolating all seven non-unit roots. Three roots are exact rational factors 1/3, 1/2, and 2/3; four roots are certified by disjoint rational sign-change intervals of the exact quartic factor. This gives exact root isolation and spectral ordering, but it still does not construct number-field eigenprojectors, assign roots to physical contact rows, or derive charge, representation, local-field, mass-activation, decoupling, threshold-beta, or physical-constant semantics."

	return Analysis{
		Previous:                     prev,
		Quartic:                      quartic,
		Certificate:                  cert,
		Projectors:                   projectors,
		Requirements:                 req,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		ExactRationalOverlapMatrix:   prev.ExactRationalOverlapMatrix,
		ExactCharacteristicCertified: prev.ExactCharacteristicCertified,
		ExactRootIsolationCertified:  rootIsolationComplete,
		ExactNumberFieldProjectors:   0,
		RootIsolationCertificates:    summary.RootIsolationCertificates,
		RowwiseRootAssignmentProofs:  0,
		ChargeSemanticRows:           0,
		T3RRowsDerived:               0,
		ChiralityRowsDerived:         0,
		BMinusLRowsDerived:           0,
		SU2LRowsDerived:              0,
		HyperchargeRowsDerived:       0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       prev.RepresentationOpenRows,
		ContactBetaRowsAllowed:       0,
		ContactZeroRowsProved:        0,
		BetaPermissionFirewallClosed: true,
		ThresholdCorrectedBeta:       false,
		FullBetaMatchingTensor:       false,
		ResidualS6Choices:            prev.ResidualS6Choices,
		ResidualNullityBefore:        prev.ResidualNullityAfter,
		ResidualNullityAfter:         prev.ResidualNullityAfter,
		HiddenObservedInputUsed:      false,
		PhysicalWeakAngleDerived:     false,
		FineStructureDerived:         false,
		PhysicalMassesDerived:        false,
		PhysicalScaleDerived:         false,
		TruthStatement:               truth,
		RejectedClaims: []string{
			"root isolation implies contact charge semantics",
			"quartic roots may be treated as hypercharge rows",
			"spectral ordering is a physical row assignment",
			"exact roots open threshold beta rows",
			"observed constants may be used to orient or normalize the contact roots",
		},
		RemainingUnknowns: []string{
			"exact number-field arithmetic for quartic-root eigenprojectors",
			"row-wise eigenprojector assignment to contact modes",
			"semantic map from isolated roots to T3R, B-L, hypercharge, or representation rows",
			"local field variables, mass activation, and decoupling",
			"threshold-corrected beta tensor and physical-flow selector",
		},
		RecommendedNextGate: "Gate 151 — exact contact eigenprojector number-field / spectral idempotent construction attempt",
	}, nil
}

func quarticEval(x *big.Rat) *big.Rat {
	// 3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271
	x2 := new(big.Rat).Mul(x, x)
	x3 := new(big.Rat).Mul(x2, x)
	x4 := new(big.Rat).Mul(x3, x)
	acc := new(big.Rat).Mul(rat(3240, 1), x4)
	acc.Sub(acc, new(big.Rat).Mul(rat(7668, 1), x3))
	acc.Add(acc, new(big.Rat).Mul(rat(6426, 1), x2))
	acc.Sub(acc, new(big.Rat).Mul(rat(2235, 1), x))
	acc.Add(acc, rat(271, 1))
	return acc
}

func intervalsDisjoint(rs []IsolatedRoot) bool {
	for i := 0; i < len(rs)-1; i++ {
		if mustRat(rs[i].IntervalRight).Cmp(mustRat(rs[i+1].IntervalLeft)) >= 0 {
			return false
		}
	}
	return true
}

func allCertified(rs []IsolatedRoot) bool {
	for _, r := range rs {
		if !r.Certified {
			return false
		}
	}
	return true
}

func mustRat(s string) *big.Rat {
	r := new(big.Rat)
	if _, ok := r.SetString(s); !ok {
		panic("invalid rational literal: " + s)
	}
	return r
}

func rat(n, d int64) *big.Rat { return new(big.Rat).SetFrac(big.NewInt(n), big.NewInt(d)) }

func boolInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func signString(s int) string {
	if s > 0 {
		return "+"
	}
	if s < 0 {
		return "-"
	}
	return "0"
}

func FormatRoot(r IsolatedRoot) string {
	if r.Kind == "rational" {
		return fmt.Sprintf("%s %s=%s mult=%d certified=%t (%s)", r.Label, r.Factor, r.Point, r.Multiplicity, r.Certified, r.PhysicalMeaning)
	}
	return fmt.Sprintf("%s interval=[%s,%s] signs=%s/%s change=%t mult=%d approx=%s certified=%t (%s)", r.Label, r.IntervalLeft, r.IntervalRight, signString(r.LeftSign), signString(r.RightSign), r.SignChange, r.Multiplicity, r.Approx, r.Certified, r.PhysicalMeaning)
}

func FormatRoots(rs []IsolatedRoot) string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		parts = append(parts, FormatRoot(r))
	}
	return strings.Join(parts, "; ")
}

func FormatQuartic(q QuarticIsolation) string {
	return fmt.Sprintf("poly=%s degree=%d disjoint=%t signChanges=%d covered=%t roots=[%s] (%s)", q.Polynomial, q.Degree, q.IntervalsDisjoint, q.SignChanges, q.AllQuarticRootsCovered, FormatRoots(q.Intervals), q.Verdict)
}

func FormatCertificate(c RootIsolationCertificate) string {
	return fmt.Sprintf("exactCharInherited=%t rational=%d quartic=%d totalPartial=%d unitMult=%d complete=%t ordered=[%s] (%s)", c.ExactCharpolyInherited, c.RationalRootsCertified, c.QuarticRootsCertified, c.TotalPartialRoots, c.UnitRootMultiplicity, c.RootIsolationComplete, FormatRoots(c.OrderedRoots), c.Verdict)
}

func FormatProjectors(p EigenprojectorAudit) string {
	return fmt.Sprintf("matrixInherited=%t rootIso=%t numberFieldProjectors=%d rationalProjectors=%d quarticProjectors=%d rowAssignments=%d rootSemantics=%d charge=%d reps=%d beta=%d (%s)", p.ExactMatrixCertificateInherited, p.ExactRootIsolationAvailable, p.ExactNumberFieldProjectors, p.RationalProjectors, p.QuarticProjectors, p.RowwiseProjectorAssignments, p.RootToContactRowSemantics, p.ChargeSemanticRows, p.RepresentationRows, p.BetaRowsAllowed, p.Verdict)
}

func FormatRequirements(r ConstructionRequirements) string {
	return fmt.Sprintf("matrix=%t char=%t rootIso=%t numberField=%t projectors=%t rowAssign=%t charge=%t reps=%t observedFree=%t physicsAll=%t (%s)", r.ExactRationalMatrix, r.ExactCharacteristicPolynomial, r.ExactRootIsolation, r.ExactNumberFieldArithmetic, r.EigenprojectorFormulaInField, r.RowwiseProjectorAssignment, r.ChargeOperatorSelected, r.RepresentationRowsSelected, r.ObservedInputFree, r.AllSatisfiedForPhysics, r.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d matrixCert=%d charCert=%d rootIso=%d rational=%d quartic=%d unitMult=%d rowProof=%d semantic=%d reps=%d/%d beta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.ExactMatrixCertificates, s.ExactCharpolyCertificates, s.RootIsolationCertificates, s.RationalRootCertificates, s.QuarticRootCertificates, s.UnitEigenMultiplicity, s.RowAssignmentProofs, s.ChargeSemanticRows, s.RepresentationCompleteRows, s.RepresentationOpenRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
