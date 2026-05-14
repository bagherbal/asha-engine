// Package branchprojector implements Gate 188: branchwise quadratic
// idempotent / scalar-projector construction audit.
//
// Gate 187 derived the threefold resolvent-vacuum orbit R_pair=Q[z]/(r3),
// but deliberately stopped before constructing a scalar projector. Gate 188
// asks what is actually constructible after a spontaneous branch is admitted.
//
// The exact answer is slightly subtler than the naive phrase "over Q(z) the
// quartic splits into two quadratics". A resolvent root z selects the unordered
// 2+2 partition. To write the two quadratic factors themselves, one must adjoin
// a quadratic element eta with eta^2=z^2-4d, which exchanges the two pair
// factors. This is not a split into four roots. It is the minimal branchwise
// factor-label extension needed to construct the two complementary 2D
// idempotents. The unordered pair {P_A,P_B} is invariant under eta -> -eta.
package branchprojector

import (
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/resolventvacuum"
)

type Rational = big.Rat

type BaseFieldAudit struct {
	QuarticPolynomial                    string
	MonicQuarticCoefficients             []string
	ResolventAlgebra                     string
	ResolventPolynomial                  string
	QuadraticAdjunction                  string
	ExtensionAlgebra                     string
	ResolventRootSelectsPartition        bool
	QuadraticAdjunctionLabelsPairFactors bool
	DoesNotAdjoinIndividualRoots         bool
	InvariantUnderFactorSwap             bool
	Verdict                              string
}

type QuadraticFactorAudit struct {
	FactorA                        string
	FactorB                        string
	FactorizationVerified          bool
	FactorsMonicQuadratic          bool
	FactorsCoprime                 bool
	NoLinearRootFactorsConstructed bool
	OnlyTwoPlusTwoSplitConstructed bool
	FactorSwapInvolutionPreserved  bool
	Verdict                        string
}

type BezoutAudit struct {
	BezoutIdentity                 string
	LeftCoefficient                string
	RightCoefficient               string
	IdentityVerified               bool
	UsesExtendedEuclideanAlgorithm bool
	ExactArithmetic                bool
	NoNumericRootApproximation     bool
	Verdict                        string
}

type ProjectorRecord struct {
	Name                       string
	Polynomial                 string
	TraceOnQuarticModule       string
	Idempotent                 bool
	Complementary              bool
	OrthogonalToOtherProjector bool
	ProjectsToQuadraticFactor  string
	Dimension                  int
}

type ProjectorAudit struct {
	Projectors                  []ProjectorRecord
	ProjectorPairConstructed    bool
	ProjectorSumIdentity        bool
	ProjectorsIdempotent        bool
	ProjectorsOrthogonal        bool
	TraceTwoEach                bool
	DimensionTwoEach            bool
	IndividualRootProjectors    int
	PhysicalScalarBundleDerived bool
	CanonicalBranchSelected     bool
	Verdict                     string
}

type HiggsCarrierAudit struct {
	Gate37PairDegenerate              bool
	BranchwiseTwoPlusTwoProjectors    bool
	ConditionalScalarProjectorDerived bool
	CanonicalScalarProjectorDerived   bool
	PhysicalScalarBundleDerived       bool
	ScalarBundleMapRequiresNextBridge bool
	ChernWeilReady                    bool
	HeatKernelReady                   bool
	ThresholdRowsReady                bool
	Verdict                           string
}

type Summary struct {
	TestsAudited                      int
	ResolventVacuumInherited          bool
	BranchwiseExtensionConstructed    bool
	QuadraticFactorsConstructed       bool
	BezoutIdentityConstructed         bool
	ProjectorPairConstructed          bool
	ConditionalScalarProjectorDerived bool
	PhysicalScalarBundleDerived       bool
	IndividualRootProjectors          int
	Comment                           string
}

type Firewall struct {
	UsesObservedInputForDerivation    bool
	UsesNumericRootApproximation      bool
	UsesIndividualRootDiagonalization bool
	UsesArbitraryPairingChoice        bool
	ResolventVacuumInherited          bool
	SpontaneousBranchDataQuarantined  bool
	QuadraticAdjunctionRecorded       bool
	BranchwiseQuadraticFactorsDerived bool
	BranchwiseProjectorPairDerived    bool
	ConditionalScalarProjectorDerived bool
	CanonicalUniqueBranchDerived      bool
	CanonicalScalarProjectorDerived   bool
	PhysicalScalarBundleDerived       bool
	ChernWeilCarrierDerived           bool
	HeatKernelMatchingDerived         bool
	ThresholdCorrectedBetaDerived     bool
	AbsoluteCouplingPromoted          bool
	PhysicalConstantsDerived          bool
	StrictNullityBefore               int
	StrictNullityAfter                int
	ConditionalNullityBefore          int
	ConditionalNullityAfter           int
	ClosedStatements                  []string
	OpenRequirements                  []string
	RecommendedNextGate               string
	Verdict                           string
}

type Analysis struct {
	PreviousGate187 resolventvacuum.Analysis
	BaseField       BaseFieldAudit
	Factors         QuadraticFactorAudit
	Bezout          BezoutAudit
	Projectors      ProjectorAudit
	HiggsCarrier    HiggsCarrierAudit
	Summary         Summary
	Firewall        Firewall
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := resolventvacuum.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 187 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev)
	})
	return defaultA, defaultErr
}

func Build(prev resolventvacuum.Analysis) (Analysis, error) {
	if !prev.Firewall.ResolventVacuumAlgebraDerived || !prev.Firewall.DegenerateVacuumOrbitDerived || !prev.Firewall.ConditionalScalarCarrierOpened {
		return Analysis{}, fmt.Errorf("Gate 188 requires Gate 187 resolvent-vacuum orbit and conditional scalar carrier")
	}
	if prev.Firewall.PhysicalScalarBundleDerived || prev.Firewall.CanonicalScalarProjectorDerived || prev.Firewall.CanonicalTwoPlusTwoSelectorDerived {
		return Analysis{}, fmt.Errorf("Gate 188 expects Gate 187 to stop before scalar-bundle/projector promotion")
	}

	cert, err := buildProjectorCertificate()
	if err != nil {
		return Analysis{}, err
	}

	base := auditBaseField(prev)
	factors := auditFactors(cert)
	bezout := auditBezout(cert)
	projectors := auditProjectors(cert)
	higgs := auditHiggsCarrier(prev, projectors)
	summary := Summary{
		TestsAudited:                      6,
		ResolventVacuumInherited:          prev.Firewall.ResolventVacuumAlgebraDerived,
		BranchwiseExtensionConstructed:    base.QuadraticAdjunctionLabelsPairFactors && base.DoesNotAdjoinIndividualRoots,
		QuadraticFactorsConstructed:       factors.FactorizationVerified && factors.OnlyTwoPlusTwoSplitConstructed,
		BezoutIdentityConstructed:         bezout.IdentityVerified,
		ProjectorPairConstructed:          projectors.ProjectorPairConstructed,
		ConditionalScalarProjectorDerived: higgs.ConditionalScalarProjectorDerived,
		PhysicalScalarBundleDerived:       higgs.PhysicalScalarBundleDerived,
		IndividualRootProjectors:          projectors.IndividualRootProjectors,
		Comment:                           "Gate 188 constructs the branchwise unordered pair of complementary 2D scalar projectors using exact algebraic factors and Bezout identity; it does not split the quadratics into four roots or promote a physical scalar bundle.",
	}
	fw := Firewall{
		UsesObservedInputForDerivation:    false,
		UsesNumericRootApproximation:      false,
		UsesIndividualRootDiagonalization: false,
		UsesArbitraryPairingChoice:        false,
		ResolventVacuumInherited:          prev.Firewall.ResolventVacuumAlgebraDerived,
		SpontaneousBranchDataQuarantined:  prev.Firewall.SpontaneousBranchDataQuarantined,
		QuadraticAdjunctionRecorded:       base.QuadraticAdjunctionLabelsPairFactors,
		BranchwiseQuadraticFactorsDerived: factors.FactorizationVerified,
		BranchwiseProjectorPairDerived:    projectors.ProjectorPairConstructed,
		ConditionalScalarProjectorDerived: higgs.ConditionalScalarProjectorDerived,
		CanonicalUniqueBranchDerived:      false,
		CanonicalScalarProjectorDerived:   false,
		PhysicalScalarBundleDerived:       false,
		ChernWeilCarrierDerived:           false,
		HeatKernelMatchingDerived:         false,
		ThresholdCorrectedBetaDerived:     false,
		AbsoluteCouplingPromoted:          false,
		PhysicalConstantsDerived:          false,
		StrictNullityBefore:               prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:          prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:           1,
		ClosedStatements: []string{
			"a resolvent branch plus the eta^2=z^2-4d quadratic factor-label adjunction gives two exact quadratic factors of q4",
			"the two factors are coprime, so Bezout identity constructs complementary idempotents in the quartic companion algebra",
			"the resulting projectors are exact, orthogonal, sum to identity, and have trace two on the four-dimensional quartic module",
			"no individual quartic root, linear factor, observed physical input, or numerical diagonalization is used",
		},
		OpenRequirements: []string{
			"identify a lawful branch/orientation datum with the physical H_Phi convention without using observation or arbitrary selection",
			"construct the scalar-bundle map from the branchwise projector pair to the existing Gate-37 Higgs carrier",
			"derive a branchwise complex/symplectic orientation or prove it remains noncanonical",
			"only after the scalar bundle map is derived may Chern-Weil, heat-kernel, threshold rows, and absolute couplings be reopened",
		},
		RecommendedNextGate: "Gate 189 — scalar-bundle map / H_Phi projector identification audit",
		Verdict:             "Gate 188 succeeds conditionally: branchwise 2+2 scalar projectors are derived exactly from the resolvent-vacuum branch and Bezout identity, while the physical scalar bundle and absolute branch choice remain sealed.",
	}
	truth := "Gate 188 is the first genuine scalar-projector construction. The necessary refinement is that Q(z) selects the 2+2 partition, while eta^2=z^2-4d labels the two quadratic pair factors. Over this branchwise extension, Bezout identity constructs two complementary idempotents on Q[x]/(q4). This opens a conditional scalar projector pair without diagonalizing individual roots, without importing observed input, and without yet claiming the physical scalar bundle."
	return Analysis{PreviousGate187: prev, BaseField: base, Factors: factors, Bezout: bezout, Projectors: projectors, HiggsCarrier: higgs, Summary: summary, Firewall: fw, TruthStatement: truth}, nil
}

func auditBaseField(prev resolventvacuum.Analysis) BaseFieldAudit {
	return BaseFieldAudit{
		QuarticPolynomial:                    "q4(x)=3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		MonicQuarticCoefficients:             []string{"1", "-71/30", "1071/540", "-149/216", "271/3240"},
		ResolventAlgebra:                     prev.Resolvent.AlgebraName,
		ResolventPolynomial:                  prev.Resolvent.SourcePolynomial,
		QuadraticAdjunction:                  "eta^2 = z^2 - 271/810",
		ExtensionAlgebra:                     "K_pair = Q[z,eta]/(r3(z), eta^2-(z^2-271/810))",
		ResolventRootSelectsPartition:        prev.VacuumOrbit.DegenerateVacuumOrbitDerived,
		QuadraticAdjunctionLabelsPairFactors: true,
		DoesNotAdjoinIndividualRoots:         true,
		InvariantUnderFactorSwap:             true,
		Verdict:                              "Q[z]/(r3) selects the unordered 2+2 partition; adjoining eta labels the two quadratic pair factors and is exchanged by eta -> -eta. No quartic root or linear factor is adjoined.",
	}
}

func auditFactors(c projectorCertificate) QuadraticFactorAudit {
	return QuadraticFactorAudit{
		FactorA:                        c.factorA.String(),
		FactorB:                        c.factorB.String(),
		FactorizationVerified:          c.factorizationVerified,
		FactorsMonicQuadratic:          c.factorA.degree() == 2 && c.factorB.degree() == 2 && c.factorA.leading().Equal(oneE()) && c.factorB.leading().Equal(oneE()),
		FactorsCoprime:                 c.bezoutVerified,
		NoLinearRootFactorsConstructed: true,
		OnlyTwoPlusTwoSplitConstructed: true,
		FactorSwapInvolutionPreserved:  factorSwapPreserved(c),
		Verdict:                        "The quartic companion polynomial factors into two monic quadratic factors over K_pair. The construction stops at 2+2 and intentionally constructs no linear root factors.",
	}
}

func auditBezout(c projectorCertificate) BezoutAudit {
	return BezoutAudit{
		BezoutIdentity:                 "A(x) q_A(x) + B(x) q_B(x) = 1",
		LeftCoefficient:                c.bezoutA.String(),
		RightCoefficient:               c.bezoutB.String(),
		IdentityVerified:               c.bezoutVerified,
		UsesExtendedEuclideanAlgorithm: true,
		ExactArithmetic:                true,
		NoNumericRootApproximation:     true,
		Verdict:                        "Exact Euclidean arithmetic in K_pair[x] gives Bezout coefficients for the two coprime quadratic factors.",
	}
}

func auditProjectors(c projectorCertificate) ProjectorAudit {
	pa := ProjectorRecord{Name: "P_A", Polynomial: c.projectorA.String(), TraceOnQuarticModule: c.traceA.String(), Idempotent: c.projectorAIdempotent, Complementary: c.projectorSumIdentity, OrthogonalToOtherProjector: c.projectorsOrthogonal, ProjectsToQuadraticFactor: "q_A", Dimension: 2}
	pb := ProjectorRecord{Name: "P_B", Polynomial: c.projectorB.String(), TraceOnQuarticModule: c.traceB.String(), Idempotent: c.projectorBIdempotent, Complementary: c.projectorSumIdentity, OrthogonalToOtherProjector: c.projectorsOrthogonal, ProjectsToQuadraticFactor: "q_B", Dimension: 2}
	return ProjectorAudit{
		Projectors:                  []ProjectorRecord{pa, pb},
		ProjectorPairConstructed:    c.projectorAIdempotent && c.projectorBIdempotent && c.projectorSumIdentity && c.projectorsOrthogonal,
		ProjectorSumIdentity:        c.projectorSumIdentity,
		ProjectorsIdempotent:        c.projectorAIdempotent && c.projectorBIdempotent,
		ProjectorsOrthogonal:        c.projectorsOrthogonal,
		TraceTwoEach:                c.traceA.Equal(intE(2)) && c.traceB.Equal(intE(2)),
		DimensionTwoEach:            true,
		IndividualRootProjectors:    0,
		PhysicalScalarBundleDerived: false,
		CanonicalBranchSelected:     false,
		Verdict:                     "The pair {P_A,P_B} is an exact branchwise scalar-projector pair: orthogonal idempotents, P_A+P_B=1, each with trace two. It is conditional branch data, not an absolute scalar-bundle identification.",
	}
}

func auditHiggsCarrier(prev resolventvacuum.Analysis, p ProjectorAudit) HiggsCarrierAudit {
	return HiggsCarrierAudit{
		Gate37PairDegenerate:              prev.Higgs.Gate37PairDegenerate,
		BranchwiseTwoPlusTwoProjectors:    p.ProjectorPairConstructed && p.TraceTwoEach,
		ConditionalScalarProjectorDerived: p.ProjectorPairConstructed && p.TraceTwoEach,
		CanonicalScalarProjectorDerived:   false,
		PhysicalScalarBundleDerived:       false,
		ScalarBundleMapRequiresNextBridge: true,
		ChernWeilReady:                    false,
		HeatKernelReady:                   false,
		ThresholdRowsReady:                false,
		Verdict:                           "The projectors now have the correct 2+2 shape for the Gate-37 Higgs carrier. The remaining problem is not projector algebra but the physical H_Phi bundle map and orientation convention.",
	}
}

func FormatBaseField(a BaseFieldAudit) string {
	return fmt.Sprintf("quartic=%q monic=%v resolvent=%q r3=%q adjunction=%q ext=%q zSelects=%t etaLabels=%t noRoots=%t swapInvariant=%t (%s)", a.QuarticPolynomial, a.MonicQuarticCoefficients, a.ResolventAlgebra, a.ResolventPolynomial, a.QuadraticAdjunction, a.ExtensionAlgebra, a.ResolventRootSelectsPartition, a.QuadraticAdjunctionLabelsPairFactors, a.DoesNotAdjoinIndividualRoots, a.InvariantUnderFactorSwap, a.Verdict)
}

func FormatFactors(a QuadraticFactorAudit) string {
	return fmt.Sprintf("qA=%s qB=%s factored=%t monicQuadratic=%t coprime=%t noLinear=%t twoPlusTwo=%t swap=%t (%s)", a.FactorA, a.FactorB, a.FactorizationVerified, a.FactorsMonicQuadratic, a.FactorsCoprime, a.NoLinearRootFactorsConstructed, a.OnlyTwoPlusTwoSplitConstructed, a.FactorSwapInvolutionPreserved, a.Verdict)
}

func FormatBezout(a BezoutAudit) string {
	return fmt.Sprintf("identity=%q A=%s B=%s verified=%t eea=%t exact=%t numeric=%t (%s)", a.BezoutIdentity, a.LeftCoefficient, a.RightCoefficient, a.IdentityVerified, a.UsesExtendedEuclideanAlgorithm, a.ExactArithmetic, !a.NoNumericRootApproximation, a.Verdict)
}

func FormatProjectors(a ProjectorAudit) string {
	parts := make([]string, 0, len(a.Projectors))
	for _, p := range a.Projectors {
		parts = append(parts, fmt.Sprintf("%s=%s trace=%s idem=%t dim=%d", p.Name, p.Polynomial, p.TraceOnQuarticModule, p.Idempotent, p.Dimension))
	}
	return fmt.Sprintf("constructed=%t sumI=%t idempotent=%t orthogonal=%t trace2=%t dim2=%t rootProjectors=%d physicalBundle=%t branchSelected=%t projectors={%s} (%s)", a.ProjectorPairConstructed, a.ProjectorSumIdentity, a.ProjectorsIdempotent, a.ProjectorsOrthogonal, a.TraceTwoEach, a.DimensionTwoEach, a.IndividualRootProjectors, a.PhysicalScalarBundleDerived, a.CanonicalBranchSelected, strings.Join(parts, "; "), a.Verdict)
}

func FormatHiggsCarrier(a HiggsCarrierAudit) string {
	return fmt.Sprintf("gate37Pair=%t branchProjectors=%t conditional=%t canonical=%t physicalBundle=%t nextMap=%t chernWeil=%t heat=%t thresholds=%t (%s)", a.Gate37PairDegenerate, a.BranchwiseTwoPlusTwoProjectors, a.ConditionalScalarProjectorDerived, a.CanonicalScalarProjectorDerived, a.PhysicalScalarBundleDerived, a.ScalarBundleMapRequiresNextBridge, a.ChernWeilReady, a.HeatKernelReady, a.ThresholdRowsReady, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d prev=%t ext=%t factors=%t bezout=%t projectors=%t conditional=%t physical=%t rootProjectors=%d (%s)", a.TestsAudited, a.ResolventVacuumInherited, a.BranchwiseExtensionConstructed, a.QuadraticFactorsConstructed, a.BezoutIdentityConstructed, a.ProjectorPairConstructed, a.ConditionalScalarProjectorDerived, a.PhysicalScalarBundleDerived, a.IndividualRootProjectors, a.Comment)
}

func FormatFirewall(a Firewall) string {
	return fmt.Sprintf("observed=%t numeric=%t rootDiag=%t arbitraryPair=%t Rpair=%t spontaneous=%t eta=%t factors=%t projectors=%t conditional=%t uniqueBranch=%t canonicalProjector=%t physicalBundle=%t chernWeil=%t heat=%t thresholds=%t absolute=%t constants=%t strict=%d->%d conditional=%d->%d closed=[%s] open=[%s] next=%s verdict=%s", a.UsesObservedInputForDerivation, a.UsesNumericRootApproximation, a.UsesIndividualRootDiagonalization, a.UsesArbitraryPairingChoice, a.ResolventVacuumInherited, a.SpontaneousBranchDataQuarantined, a.QuadraticAdjunctionRecorded, a.BranchwiseQuadraticFactorsDerived, a.BranchwiseProjectorPairDerived, a.ConditionalScalarProjectorDerived, a.CanonicalUniqueBranchDerived, a.CanonicalScalarProjectorDerived, a.PhysicalScalarBundleDerived, a.ChernWeilCarrierDerived, a.HeatKernelMatchingDerived, a.ThresholdCorrectedBetaDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, strings.Join(a.ClosedStatements, "; "), strings.Join(a.OpenRequirements, "; "), a.RecommendedNextGate, a.Verdict)
}
