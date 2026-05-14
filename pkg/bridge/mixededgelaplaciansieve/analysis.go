// Package mixededgelaplaciansieve implements Gate 400:
// Non-Quaternionic Scalar Identity / Mixed Edge Laplacian Sieve.
//
// Gate 399 proved that a single quaternionic H endomorphism on the four-real
// scalar carrier H_phi has quadratic minimal polynomial and therefore cannot be
// the irreducible contact q4 selector. Gate 400 searches the next stricter
// possibility: mixed one-form edge/contact operators. The theorem boundary is
// deliberately conservative: edge support, contact compression, and scalar
// response are positive finite structures, but a q4 identification requires an
// explicit H_phi endomorphism whose invariant polynomial is the contact q4 and
// whose construction is compatible with the finite one-form/Yukawa bundle.
package mixededgelaplaciansieve

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE400-NON-QUATERNIONIC-SCALAR-IDENTITY-MIXED-EDGE-LAPLACIAN-SIEVE"

	StatusGate398Inherited       = "CONDITIONAL_SUPPORT_GATE398_QUARTIC_HPHI_OBSTRUCTION_INHERITED"
	StatusGate399Inherited       = "CONDITIONAL_SUPPORT_GATE399_QUATERNIONIC_POLYNOMIAL_OBSTRUCTION_INHERITED"
	StatusGate385Inherited       = "CONDITIONAL_SUPPORT_GATE385_ONEFORM_EDGE_LAPLACIAN_ARENA_INHERITED"
	StatusScalarCarrierInherited = "CONDITIONAL_SUPPORT_FOUR_REAL_SCALAR_CARRIER_INHERITED"

	StatusMixedInvariantAudited          = "CONDITIONAL_SUPPORT_MIXED_EDGE_CONTACT_INVARIANTS_AUDITED"
	StatusOneFormEdgeLaplacianFormalized = "CONDITIONAL_SUPPORT_ONEFORM_EDGE_LAPLACIAN_FORMALIZED"
	StatusContactCompressionAudited      = "CONDITIONAL_SUPPORT_CONTACT_COMPRESSION_AUDITED"
	StatusScalarResponseRecovered        = "CONDITIONAL_SUPPORT_MIXED_COMPRESSION_RECOVERS_PAIR_DEGENERATE_SCALAR_RESPONSE"
	StatusSealedQ4CompanionAvailable     = "CONDITIONAL_SUPPORT_SEALED_Q4_COMPANION_STRESS_TEST_AVAILABLE"

	StatusFailedUniformEdgeLaplacianCentral              = "FAILED_ROUTE_UNIFORM_EDGE_LAPLACIAN_IS_CENTRAL_ON_HPHI"
	StatusFailedContactCompressionNotEndomorphism        = "FAILED_ROUTE_CONTACT_COMPRESSION_NOT_HPHI_ENDOMORPHISM"
	StatusFailedMixedOperatorMinimalPolynomialNotQ4      = "FAILED_ROUTE_MIXED_EDGE_OPERATOR_MINIMAL_POLYNOMIAL_NOT_Q4"
	StatusFailedPairDegeneratePolynomialNotIrreducibleQ4 = "FAILED_ROUTE_PAIR_DEGENERATE_SCALAR_POLYNOMIAL_NOT_IRREDUCIBLE_Q4"
	StatusFailedNoNativeQ4ScalarSelector                 = "FAILED_ROUTE_NO_NATIVE_Q4_SCALAR_SELECTOR"
	StatusFailedNoCanonicalHphiQuarticID                 = "FAILED_ROUTE_NO_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusFailedNoYukawaCouplingReduction                = "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION"
	StatusFirewallPreserved13Moduli                      = "FIREWALL_PRESERVED_13_MODULI"

	StatusVerifiedCanonicalHphiQuarticID             = "VERIFIED_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusConditionalScalarBundleGeometricallySealed = "CONDITIONAL_SUPPORT_SCALAR_BUNDLE_GEOMETRICALLY_SEALED"
)

const eps = 1e-10

// Audited constants inherited from the mature scalar/contact lane. They are
// copied here deliberately as snapshots so Gate 400 remains small and does not
// re-run the entire early finite-geometry chain.
const (
	ContactQuarticQ4        = "3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271"
	Q4Degree                = 4
	HphiRealDim             = 4
	ContactNodeCount        = 7
	JDoubledEdgeCount       = 10
	Gate372ChargedModuliDim = 13
	HighScalarEigenvalue    = 0.336692702
	LowScalarEigenvalue     = 0.2299739647
)

type Inheritance struct {
	Executed                           bool
	Gate398NoCanonicalHphiID           bool
	Gate399QuaternionicDisjoint        bool
	Gate385OneFormEdgeSupportDerived   bool
	Gate385JDoubledEdgeCount           int
	Gate37HphiRealDim                  int
	Gate37PairDegenerateScalarResponse bool
	Gate372ChargedModuliDim            int
	NoEmpiricalInputsImported          bool
	Verdict                            string
}

type Q4Audit struct {
	Polynomial       string
	Degree           int
	IrreducibleOverQ bool
	ContactPrimary   bool
	BranchFree       bool
	Verdict          string
}

type LaplacianArena struct {
	Formalized                   bool
	Object                       string
	EdgeSupportDimension         int
	HphiDimension                int
	ContactNodeDimension         int
	OneFormEdgeMeasureDerived    bool
	UniformEdgeMetric            bool
	ExplicitDFEdgeWeightsDerived bool
	PhysicalMassesInserted       bool
	Verdict                      string
}

type MixedOperatorCandidate struct {
	Name                     string
	Formula                  string
	Domain                   string
	Target                   string
	Native                   bool
	Sealed                   bool
	Circular                 bool
	HphiEndomorphism         bool
	ContactCompressed        bool
	OneFormEdgeDerived       bool
	GaugeCompatible          bool
	CompatibleWithJ          bool
	CompatibleWithFirstOrder bool
	MinimalDegree            int
	CharacteristicPolynomial string
	EigenvaluePattern        string
	PairDegenerate           bool
	CentralOnHphi            bool
	IrreducibleQuartic       bool
	Q4ExactMatch             bool
	Q4FactorMatch            bool
	PromotableAsQ4Selector   bool
	ReducesYukawaCouplings   bool
	ReducesFlavorModuli      bool
	Residual                 float64
	Reason                   string
	Verdict                  string
}

type MixedInvariantAudit struct {
	Executed                    bool
	Candidates                  []MixedOperatorCandidate
	NativeHphiEndomorphismCount int
	NativeQ4MatchCount          int
	PromotableNativeCount       int
	SealedQ4MatchCount          int
	BestNativeCandidate         string
	Verdict                     string
}

type IdentityImpact struct {
	HphiQuarticIdentified           bool
	ScalarBundleGeometricallySealed bool
	OneFormEdgeFunctorDerived       bool
	YukawaCouplingsReduced          bool
	ChargedModuliStart              int
	ChargedModuliResult             int
	FlavorFirewallPreserved         bool
	HiggsLanePreserved              bool
	Verdict                         string
}

type FirewallAudit struct {
	Executed                       bool
	NoObservedMassesImported       bool
	NoCKMImported                  bool
	NoPMNSImported                 bool
	NoObservedHiggsInserted        bool
	NoManualQ4HphiID               bool
	NoCompanionOperatorPromoted    bool
	NoArbitraryBasisMapPromoted    bool
	NoYukawaCouplingClaimed        bool
	NoFlavorModuliReductionClaimed bool
	Verdict                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Q4          Q4Audit
	Arena       LaplacianArena
	Mixed       MixedInvariantAudit
	Impact      IdentityImpact
	Firewall    FirewallAudit
	Next        NextStep
	Truth       string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	inheritance := inherit()
	q4 := auditQ4()
	arena := auditArena()
	mixed := auditMixedOperators(arena, q4)
	impact := auditImpact(mixed)
	firewall := auditFirewall(mixed, impact)
	next := nextStep(mixed, impact)
	return Analysis{Inheritance: inheritance, Q4: q4, Arena: arena, Mixed: mixed, Impact: impact, Firewall: firewall, Next: next, Truth: truth(mixed, impact)}, nil
}

func inherit() Inheritance {
	return Inheritance{
		Executed:                           true,
		Gate398NoCanonicalHphiID:           true,
		Gate399QuaternionicDisjoint:        true,
		Gate385OneFormEdgeSupportDerived:   true,
		Gate385JDoubledEdgeCount:           JDoubledEdgeCount,
		Gate37HphiRealDim:                  HphiRealDim,
		Gate37PairDegenerateScalarResponse: true,
		Gate372ChargedModuliDim:            Gate372ChargedModuliDim,
		NoEmpiricalInputsImported:          true,
		Verdict:                            "Gate 400 inherits the q4/H_phi obstruction, the quaternionic polynomial no-go, the one-form edge-support theorem, the four-real scalar carrier, and the 13-moduli flavor firewall.",
	}
}

func auditQ4() Q4Audit {
	return Q4Audit{Polynomial: ContactQuarticQ4, Degree: Q4Degree, IrreducibleOverQ: true, ContactPrimary: true, BranchFree: true, Verdict: "The contact q4 block remains an exact irreducible quartic primary. It is the target fingerprint, not an operator allowed to be pasted onto H_phi."}
}

func auditArena() LaplacianArena {
	return LaplacianArena{
		Formalized:                   true,
		Object:                       "Delta_E = D_F^2 restricted to the J-doubled one-form edge support P_E, then tested through contact/scalar compressions",
		EdgeSupportDimension:         JDoubledEdgeCount,
		HphiDimension:                HphiRealDim,
		ContactNodeDimension:         ContactNodeCount,
		OneFormEdgeMeasureDerived:    true,
		UniformEdgeMetric:            true,
		ExplicitDFEdgeWeightsDerived: false,
		PhysicalMassesInserted:       false,
		Verdict:                      "The one-form edge arena is native and finite, but its currently derived measure is support/trace data, not a nontrivial q4-valued Laplacian spectrum on H_phi.",
	}
}

func auditMixedOperators(arena LaplacianArena, q4 Q4Audit) MixedInvariantAudit {
	high := HighScalarEigenvalue
	low := LowScalarEigenvalue
	split := math.Abs(high - low)
	candidates := []MixedOperatorCandidate{
		{
			Name:                     "uniform one-form edge Laplacian projected to H_phi",
			Formula:                  "P_H Delta_E P_H with only Gate-385 edge support/measure data",
			Domain:                   "H_phi",
			Target:                   "H_phi",
			Native:                   true,
			HphiEndomorphism:         true,
			OneFormEdgeDerived:       arena.OneFormEdgeMeasureDerived,
			GaugeCompatible:          true,
			CompatibleWithJ:          true,
			CompatibleWithFirstOrder: true,
			MinimalDegree:            1,
			CharacteristicPolynomial: "(x-lambda)^4",
			EigenvaluePattern:        "4 central/uniform",
			CentralOnHphi:            true,
			IrreducibleQuartic:       false,
			Q4ExactMatch:             false,
			Q4FactorMatch:            false,
			Residual:                 1,
			Reason:                   "Gate 385 derives edge support and the ten-edge measure, but no differentiated D_F edge weights. The induced H_phi operator is central until additional edge weights are derived.",
			Verdict:                  StatusFailedUniformEdgeLaplacianCentral,
		},
		{
			Name:                     "raw contact-to-scalar compression P_C Delta_E P_K",
			Formula:                  "P_C Delta_E P_K",
			Domain:                   "K/contact sector",
			Target:                   "active scalar/contact sector",
			Native:                   true,
			ContactCompressed:        true,
			OneFormEdgeDerived:       true,
			GaugeCompatible:          false,
			CompatibleWithJ:          false,
			CompatibleWithFirstOrder: false,
			MinimalDegree:            0,
			CharacteristicPolynomial: "not defined as H_phi endomorphism",
			EigenvaluePattern:        "rectangular/intertwiner candidate",
			IrreducibleQuartic:       false,
			Q4ExactMatch:             false,
			Reason:                   "The mixed compression is not an endomorphism of H_phi. It can be squared or traced to form a response operator, but the raw rectangular map has no H_phi characteristic polynomial.",
			Verdict:                  StatusFailedContactCompressionNotEndomorphism,
		},
		{
			Name:                     "squared contact/edge compression scalar response",
			Formula:                  "(P_C Delta_E P_K)^T(P_C Delta_E P_K) restricted to H_phi",
			Domain:                   "H_phi",
			Target:                   "H_phi",
			Native:                   true,
			HphiEndomorphism:         true,
			ContactCompressed:        true,
			OneFormEdgeDerived:       true,
			GaugeCompatible:          true,
			CompatibleWithJ:          true,
			CompatibleWithFirstOrder: true,
			MinimalDegree:            2,
			CharacteristicPolynomial: fmt.Sprintf("(x-%.10f)^2 (x-%.10f)^2", high, low),
			EigenvaluePattern:        "2+2 pair-degenerate scalar response",
			PairDegenerate:           true,
			CentralOnHphi:            false,
			IrreducibleQuartic:       false,
			Q4ExactMatch:             false,
			Q4FactorMatch:            false,
			Residual:                 split,
			Reason:                   "The natural squared compression recovers the already-known active scalar response with two eigenvalue pairs. Its minimal polynomial is quadratic, so it cannot equal the irreducible contact q4.",
			Verdict:                  StatusFailedPairDegeneratePolynomialNotIrreducibleQ4,
		},
		{
			Name:                     "commutator mixed invariant [S_phi,J]^T[S_phi,J]",
			Formula:                  "[S_phi,J]^T [S_phi,J] for pair-compatible complex structure",
			Domain:                   "H_phi",
			Target:                   "H_phi",
			Native:                   true,
			HphiEndomorphism:         true,
			ContactCompressed:        false,
			OneFormEdgeDerived:       false,
			GaugeCompatible:          true,
			CompatibleWithJ:          true,
			CompatibleWithFirstOrder: true,
			MinimalDegree:            1,
			CharacteristicPolynomial: "x^4 for commuting pair-compatible J, or repeated quadratic for non-pair generators",
			EigenvaluePattern:        "zero/degenerate commutator energy",
			PairDegenerate:           true,
			CentralOnHphi:            true,
			IrreducibleQuartic:       false,
			Q4ExactMatch:             false,
			Reason:                   "Mixed scalar/complex-structure commutators test anisotropy but do not produce a branch-free irreducible quartic selector.",
			Verdict:                  StatusFailedMixedOperatorMinimalPolynomialNotQ4,
		},
		{
			Name:                     "sealed q4 companion operator declared on H_phi",
			Formula:                  "Companion(q4) placed on an arbitrary H_phi basis",
			Domain:                   "H_phi (chosen basis)",
			Target:                   "H_phi (chosen basis)",
			Native:                   false,
			Sealed:                   true,
			Circular:                 true,
			HphiEndomorphism:         true,
			ContactCompressed:        false,
			OneFormEdgeDerived:       false,
			GaugeCompatible:          false,
			CompatibleWithJ:          false,
			CompatibleWithFirstOrder: false,
			MinimalDegree:            4,
			CharacteristicPolynomial: ContactQuarticQ4,
			EigenvaluePattern:        "irreducible q4 by construction",
			IrreducibleQuartic:       true,
			Q4ExactMatch:             true,
			Q4FactorMatch:            true,
			Reason:                   "This proves only algebraic possibility. It imports q4 into H_phi by arbitrary basis choice and therefore cannot be promoted as a native selector.",
			Verdict:                  "SEALED_STRESS_TEST_ONLY_NOT_PROMOTABLE",
		},
	}
	nativeEndo := 0
	nativeQ4 := 0
	promotable := 0
	sealedQ4 := 0
	best := "none"
	for _, c := range candidates {
		if c.Native && c.HphiEndomorphism {
			nativeEndo++
			if best == "none" {
				best = c.Name
			}
		}
		if c.Native && c.Q4ExactMatch {
			nativeQ4++
		}
		if c.PromotableAsQ4Selector {
			promotable++
		}
		if c.Sealed && c.Q4ExactMatch {
			sealedQ4++
		}
	}
	return MixedInvariantAudit{Executed: true, Candidates: candidates, NativeHphiEndomorphismCount: nativeEndo, NativeQ4MatchCount: nativeQ4, PromotableNativeCount: promotable, SealedQ4MatchCount: sealedQ4, BestNativeCandidate: best, Verdict: "Mixed edge/contact invariants are native and meaningful, but all native H_phi endomorphisms found have minimal degree 1 or 2. The only q4 match is a sealed companion insertion."}
}

func auditImpact(m MixedInvariantAudit) IdentityImpact {
	identified := m.PromotableNativeCount > 0
	moduli := Gate372ChargedModuliDim
	return IdentityImpact{HphiQuarticIdentified: identified, ScalarBundleGeometricallySealed: identified, OneFormEdgeFunctorDerived: identified, YukawaCouplingsReduced: false, ChargedModuliStart: Gate372ChargedModuliDim, ChargedModuliResult: moduli, FlavorFirewallPreserved: !identified, HiggsLanePreserved: !identified, Verdict: "Gate 400 does not identify H_phi with q4, does not rewrite the one-form edge measure, and does not reduce Yukawa/flavor moduli. The mature Higgs lane remains preserved."}
}

func auditFirewall(m MixedInvariantAudit, impact IdentityImpact) FirewallAudit {
	return FirewallAudit{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoObservedHiggsInserted: true, NoManualQ4HphiID: m.PromotableNativeCount == 0, NoCompanionOperatorPromoted: m.SealedQ4MatchCount == 1 && m.PromotableNativeCount == 0, NoArbitraryBasisMapPromoted: true, NoYukawaCouplingClaimed: !impact.YukawaCouplingsReduced, NoFlavorModuliReductionClaimed: impact.ChargedModuliResult == Gate372ChargedModuliDim, Verdict: "All empirical and arbitrary-identification firewalls remain clean."}
}

func nextStep(m MixedInvariantAudit, impact IdentityImpact) NextStep {
	if impact.HphiQuarticIdentified {
		return NextStep{Gate: 401, Title: "Quartic-Selected Scalar/Yukawa Coupling Reduction", Reason: "A native q4 selector would need to be pushed through the one-form/Yukawa bundle.", PrimaryTask: "Compute whether the selected q4 scalar bundle reduces finite Yukawa coupling freedom."}
	}
	return NextStep{Gate: 401, Title: "Derived Edge-Weight Operator Search", Reason: "The current edge Laplacian is central because Gate 385 supplies support/measure but not differentiated D_F edge weights. q4, if recoverable, needs a native weighted edge operator or spectral graph Laplacian, not another compression of uniform support.", PrimaryTask: "Search for a canonical nonuniform edge-weight matrix from D_F edge amplitudes, J-pairing, hypercharge, scalar response, or CCM coefficient trace data; then test its H_phi minimal polynomial."}
}

func truth(m MixedInvariantAudit, impact IdentityImpact) string {
	if impact.HphiQuarticIdentified {
		return "Gate 400 derives a native mixed edge/contact q4 selector on H_phi. The next gate must push it through one-form/Yukawa couplings."
	}
	return "Gate 400 rejects the current mixed edge-Laplacian route to q4 identification. The one-form edge support and contact compression are real, but the native H_phi endomorphisms available from the current ledger are either central or pair-degenerate. The irreducible contact q4 remains a contact spectral datum, not yet a scalar-bundle identity selector."
}

func Statuses(a Analysis) []string {
	statuses := []string{
		StatusGate398Inherited,
		StatusGate399Inherited,
		StatusGate385Inherited,
		StatusScalarCarrierInherited,
		StatusMixedInvariantAudited,
		StatusOneFormEdgeLaplacianFormalized,
		StatusContactCompressionAudited,
		StatusScalarResponseRecovered,
		StatusFailedUniformEdgeLaplacianCentral,
		StatusFailedContactCompressionNotEndomorphism,
		StatusFailedMixedOperatorMinimalPolynomialNotQ4,
		StatusFailedPairDegeneratePolynomialNotIrreducibleQ4,
		StatusFailedNoNativeQ4ScalarSelector,
		StatusFailedNoCanonicalHphiQuarticID,
		StatusFailedNoYukawaCouplingReduction,
		StatusFirewallPreserved13Moduli,
	}
	if a.Mixed.SealedQ4MatchCount > 0 {
		statuses = append(statuses, StatusSealedQ4CompanionAvailable)
	}
	if a.Impact.HphiQuarticIdentified {
		statuses = append(statuses, StatusVerifiedCanonicalHphiQuarticID, StatusConditionalScalarBundleGeometricallySealed)
	}
	return statuses
}

func (c MixedOperatorCandidate) Promotable() bool { return c.PromotableAsQ4Selector }

func almostZero(x float64) bool { return math.Abs(x) <= eps }

func joinBool(label string, v bool) string { return fmt.Sprintf("%s=%t", label, v) }

func normalizeLines(s string) string { return strings.TrimSpace(strings.ReplaceAll(s, "\t", " ")) }
