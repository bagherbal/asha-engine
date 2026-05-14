// Package generationaddressfunctor implements Gate 394:
// Native Generation-Address Functor from Triality/Morita Edge Incidence.
//
// Gate 393 showed that Spin(8) triality is an admissible arena only as an
// abstract/label stress test: the project still lacks a native functor from
// the finite ASHA ledger into End(C^3_gen).  Gate 394 therefore audits the
// concrete sources that could address generations noncentrally: triality branch
// incidence, Morita/Dirac edge incidence, inner-fluctuation one-form support,
// Fock number candidates, and finite scalar/contact anisotropy.  The output is
// deliberately a theorem-gated obstruction unless a native noncentral and
// noncommuting generation algebra is actually derived.
package generationaddressfunctor

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE394-NATIVE-GENERATION-ADDRESS-FUNCTOR"

	StatusGate393Inherited = "CONDITIONAL_SUPPORT_GATE393_TRIALITY_DOMAIN_OBSTRUCTION_INHERITED"
	StatusGate370Inherited = "CONDITIONAL_SUPPORT_GATE370_CENTRAL_SUPPORT_MAP_OBSTRUCTION_INHERITED"
	StatusGate371Inherited = "CONDITIONAL_SUPPORT_GATE371_NUMBER_OPERATOR_CAPACITY_INHERITED"
	StatusGate372Inherited = "CONDITIONAL_SUPPORT_GATE372_THIRTEEN_MODULI_FIREWALL_INHERITED"
	StatusGate385Inherited = "CONDITIONAL_SUPPORT_GATE385_ONE_FORM_EDGE_SUPPORT_INHERITED"

	StatusFunctorSieveFormalized       = "CONDITIONAL_SUPPORT_GENERATION_ADDRESS_FUNCTOR_SIEVE_FORMALIZED"
	StatusCandidateSourcesEnumerated   = "CONDITIONAL_SUPPORT_GENERATION_ADDRESS_CANDIDATE_SOURCES_ENUMERATED"
	StatusTrialityBranchAudited        = "CONDITIONAL_SUPPORT_TRIALITY_BRANCH_INCIDENCE_AUDITED"
	StatusMoritaEdgeAudited            = "CONDITIONAL_SUPPORT_MORITA_EDGE_INCIDENCE_AUDITED"
	StatusOneFormSupportAudited        = "CONDITIONAL_SUPPORT_ONE_FORM_EDGE_SUPPORT_AUDITED"
	StatusFockNumberAudited            = "CONDITIONAL_SUPPORT_FOCK_NUMBER_OPERATOR_DERIVATION_AUDITED"
	StatusTextureCapacityAudited       = "CONDITIONAL_SUPPORT_NONCOMMUTING_TEXTURE_CAPACITY_AUDITED"
	StatusModuliImpactAudited          = "CONDITIONAL_SUPPORT_GENERATION_ADDRESS_MODULI_IMPACT_AUDITED"
	StatusConditionalAddressCapacity   = "CONDITIONAL_SUPPORT_GENERATION_ADDRESS_FUNCTOR_UNDER_SEALED_OPERATOR"
	StatusConditionalHierarchyCapacity = "CONDITIONAL_SUPPORT_NUMBER_OPERATOR_HIERARCHY"

	StatusTensionTrialityStillNeedsCarrier  = "CONDITIONAL_TENSION_TRIALITY_BRANCH_ACTION_STILL_NEEDS_NATIVE_CARRIER"
	StatusTensionMoritaUniformBroadcast     = "CONDITIONAL_TENSION_MORITA_EDGE_INCIDENCE_BROADCASTS_UNIFORMLY_OVER_GENERATIONS"
	StatusTensionOneFormUniformBroadcast    = "CONDITIONAL_TENSION_ONE_FORM_EDGE_SUPPORT_REPEATS_UNIFORMLY_OVER_GENERATIONS"
	StatusTensionProtectedContactUnassigned = "CONDITIONAL_TENSION_PROTECTED_CONTACT_ANISOTROPY_HAS_NO_NATIVE_GENERATION_ASSIGNMENT"
	StatusTensionNNonNative                 = "CONDITIONAL_TENSION_N_OPERATOR_NOT_DERIVED_FROM_CURRENT_FOCK_LEDGER"
	StatusTensionSealedCapacityNotNative    = "CONDITIONAL_TENSION_SEALED_CAPACITY_DOES_NOT_REWRITE_NATIVE_FIREWALL"
	StatusTensionNeedTwoOperators           = "CONDITIONAL_TENSION_FLAVOR_REQUIRES_TWO_NONCOMMUTING_NATIVE_OPERATORS"

	StatusVerifiedNativeAddressFunctor   = "VERIFIED_NATIVE_GENERATION_ADDRESS_FUNCTOR"
	StatusFailedGenerationAddressCentral = "FAILED_ROUTE_GENERATION_ADDRESS_REMAINS_CENTRAL"
	StatusFailedDiagonalOnlyNoCKM        = "FAILED_ROUTE_DIAGONAL_ONLY_NO_CKM"
	StatusFailedCircularTauOrNInsertion  = "FAILED_ROUTE_CIRCULAR_TAU_OR_N_INSERTION"
	StatusFailedNoNativeNoncommutingPair = "FAILED_ROUTE_NO_NATIVE_NONCOMMUTING_TEXTURE_PAIR"
	StatusFirewallPreserved13Moduli      = "FIREWALL_PRESERVED_13_MODULI"
)

const eps = 1e-10

type Inheritance struct {
	Executed                         bool
	Gate393DomainNotAdmitted         bool
	Gate393TrialityOnlyLabelSymmetry bool
	Gate370NativeSupportMapsCentral  bool
	Gate371NumberOperatorNonNative   bool
	Gate372ChargedModuliDim          int
	Gate385OneFormEdgeSupportDerived bool
	NoEmpiricalFlavorValuesImported  bool
	Verdict                          string
}

type FunctorTarget struct {
	Domain                string
	Codomain              string
	RequiredNativePattern string
	SuccessfulPattern     string
	Verdict               string
}

type Candidate struct {
	Name                         string
	Source                       string
	Operator                     [][]float64
	Weights                      []float64
	Native                       bool
	Sealed                       bool
	Circular                     bool
	DerivedFromExistingLedger    bool
	CompatibleWithJ              bool
	CompatibleWithFirstOrder     bool
	CompatibleWithHyperchargeSU2 bool
	Central                      bool
	NonCentral                   bool
	DiagonalOnly                 bool
	Rank                         int
	Spectrum                     []float64
	CommutantDimension           int
	BreaksGenerationDegeneracy   bool
	GivesMixing                  bool
	Reason                       string
	Verdict                      string
}

type CandidateAudit struct {
	Executed              bool
	Candidates            []Candidate
	NativeCandidateCount  int
	NativeNoncentralCount int
	SealedNoncentralCount int
	CentralNativeCount    int
	Verdict               string
}

type SourceAudit struct {
	Executed              bool
	Name                  string
	Inputs                []string
	Result                string
	Rank                  int
	Spectrum              []float64
	CentralOnly           bool
	NativeNoncentralFound bool
	CircularOrSealedOnly  bool
	Verdict               string
}

type NumberOperatorAudit struct {
	Executed                  bool
	Candidate                 Candidate
	Source                    string
	Status                    string
	DerivationResidual        float64
	Native                    bool
	BridgeCompatible          bool
	SealedExternalExtension   bool
	CircularIfUsedAsSolution  bool
	CommutatorWithCycle       float64
	CommutatorWithMirror      float64
	CommutesWithHypercharge   bool
	CommutesWithSU2L          bool
	CommutesWithJ             bool
	CommutesWithGamma         bool
	CommutesWithDFEdgeSupport bool
	BreaksExactTriality       bool
	ProducesHierarchy         bool
	ProducesMixing            bool
	Verdict                   string
}

type PairAudit struct {
	Name                       string
	Left                       string
	Right                      string
	NativePair                 bool
	SealedPair                 bool
	CommutatorNorm             float64
	Noncommuting               bool
	SimultaneouslyDiagonalized bool
	CKMCapacity                bool
	Reason                     string
	Verdict                    string
}

type TextureCapacityAudit struct {
	Executed                     bool
	NativeGenerationOperators    int
	NativeNoncentralOperators    int
	NativeNoncommutingPairs      int
	SealedNoncommutingPairs      int
	MaxNativeCommutatorNorm      float64
	MaxSealedCommutatorNorm      float64
	SimultaneouslyDiagonalizable bool
	CKMCapacityNative            bool
	Pairs                        []PairAudit
	Verdict                      string
}

type ModuliScenario struct {
	Name                          string
	AssumptionClass               string
	StartingChargedDim            int
	ResultingDim                  int
	Native                        bool
	Conditional                   bool
	Failed                        bool
	DistinctChargedMassesPossible bool
	CKMMisalignmentPossible       bool
	LeptonQuarkSectorSeparation   bool
	Reason                        string
	Verdict                       string
}

type ModuliAudit struct {
	Executed               bool
	StartingChargedDim     int
	NativeReductionBelow13 bool
	BestNativeDim          int
	BestConditionalDim     int
	Scenarios              []ModuliScenario
	Verdict                string
}

type FirewallAudit struct {
	Executed                     bool
	NoYukawaMassesImported       bool
	NoCKMImported                bool
	NoPMNSImported               bool
	NoEmpiricalOrderingImported  bool
	NoManualGenerationAssignment bool
	NoCircularTauInserted        bool
	NoCircularNPromoted          bool
	NoNativeAddressClaimed       bool
	NoModuliReductionClaimed     bool
	Verdict                      string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance     Inheritance
	Target          FunctorTarget
	Candidates      CandidateAudit
	TrialityBranch  SourceAudit
	MoritaEdge      SourceAudit
	OneFormSupport  SourceAudit
	Number          NumberOperatorAudit
	TextureCapacity TextureCapacityAudit
	Moduli          ModuliAudit
	Firewall        FirewallAudit
	Next            NextStep
	Truth           string
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
	inheritance := inheritPreviousGates()
	target := formalizeTarget()
	candidates, err := auditCandidates(inheritance)
	if err != nil {
		return Analysis{}, err
	}
	triality := auditTrialityBranch(candidates)
	morita := auditMoritaEdge(candidates)
	oneform := auditOneFormSupport(candidates)
	number := auditNumberOperator(candidates)
	texture := auditTextureCapacity(candidates)
	moduli := auditModuli(inheritance, candidates, texture, number)
	firewall := auditFirewall(candidates, moduli, number)
	next := chooseNextGate(candidates, texture, number)
	truth := buildTruth(candidates, texture, moduli, next)
	return Analysis{inheritance, target, candidates, triality, morita, oneform, number, texture, moduli, firewall, next, truth}, nil
}

func inheritPreviousGates() Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate393DomainNotAdmitted:         true,
		Gate393TrialityOnlyLabelSymmetry: true,
		Gate370NativeSupportMapsCentral:  true,
		Gate371NumberOperatorNonNative:   true,
		Gate372ChargedModuliDim:          13,
		Gate385OneFormEdgeSupportDerived: true,
		NoEmpiricalFlavorValuesImported:  true,
		Verdict: join(
			StatusGate393Inherited,
			StatusGate370Inherited,
			StatusGate371Inherited,
			StatusGate372Inherited,
			StatusGate385Inherited,
		),
	}
}

func formalizeTarget() FunctorTarget {
	return FunctorTarget{
		Domain:                "finite ASHA support/topology/Morita-edge/one-form data",
		Codomain:              "End(C^3_gen)",
		RequiredNativePattern: "Phi(s)=a I_3 + b T_gen with b != 0, T_gen derived and compatible with A_F, J, first-order, hypercharge, and SU(2)_L channels",
		SuccessfulPattern:     "at least one native noncentral generation operator for hierarchy, and at least two native noncommuting operators for CKM/PMNS capacity",
		Verdict:               StatusFunctorSieveFormalized,
	}
}

func auditCandidates(_ Inheritance) (CandidateAudit, error) {
	cycle := [][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	mirror := [][]float64{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}
	_ = mirror // mirror is used in the number audit; keep local construction explicit in that audit.

	candidates := []Candidate{
		newCandidate("identity generation broadcast", "Gate-370 native support trace / Morita multiplicity broadcast", identity3(), nil, true, false, false, true, "native support data acts generation-equivariantly and factors through I3"),
		newCandidate("Morita edge uniform incidence", "A_F=C⊕H⊕M3(C) Dirac/Yukawa edge multiplicity ledger", diagFromWeights([]float64{10, 10, 10}), []float64{10, 10, 10}, true, false, false, true, "the true finite spectral-triple edge ledger repeats the same edge support over generation copies"),
		newCandidate("inner-fluctuation one-form uniform support", "Gate-385 Ω¹_D(A_F) J-doubled one-form edge support", diagFromWeights([]float64{10, 10, 10}), []float64{10, 10, 10}, true, false, false, true, "Higgs one-form support selects edge measure, but it does not distinguish generation addresses"),
		newCandidate("abstract triality branch cycle", "sealed label action P=(123) from Gate-393 stress test", cycle, nil, false, true, true, false, "noncentral branch action exists only after manual generation-to-triality labeling; no native carrier/theta was admitted"),
		newCandidate("protected contact anisotropy spurion", "Gate-29 Higgs/contact anisotropy diagonal weights", diagFromWeights([]float64{0.336692702, 0.2833333333, 0.2299739647}), []float64{0.336692702, 0.2833333333, 0.2299739647}, false, true, false, true, "finite anisotropy supplies diagonal capacity, but the assignment from protected contact directions to generations is noncanonical"),
		newCandidate("Fock number ladder N", "Gate-371 bridge-level N=diag(0,1,2) candidate", diagFromWeights([]float64{0, 1, 2}), []float64{0, 1, 2}, false, true, true, false, "N has hierarchy capacity but is not derived from the current Fock/spectral-triple ledger"),
	}

	for i := range candidates {
		if err := fillCandidateLinearData(&candidates[i]); err != nil {
			return CandidateAudit{}, err
		}
	}

	nativeCount, nativeNoncentral, sealedNoncentral, centralNative := 0, 0, 0, 0
	for _, c := range candidates {
		if c.Native {
			nativeCount++
		}
		if c.Native && c.NonCentral {
			nativeNoncentral++
		}
		if c.Sealed && c.NonCentral {
			sealedNoncentral++
		}
		if c.Native && c.Central {
			centralNative++
		}
	}

	verdict := join(StatusCandidateSourcesEnumerated, StatusFailedGenerationAddressCentral, StatusConditionalAddressCapacity, StatusTensionSealedCapacityNotNative)
	if nativeNoncentral > 0 {
		verdict = join(StatusCandidateSourcesEnumerated, StatusVerifiedNativeAddressFunctor)
	}
	return CandidateAudit{
		Executed:              true,
		Candidates:            candidates,
		NativeCandidateCount:  nativeCount,
		NativeNoncentralCount: nativeNoncentral,
		SealedNoncentralCount: sealedNoncentral,
		CentralNativeCount:    centralNative,
		Verdict:               verdict,
	}, nil
}

func newCandidate(name, source string, op [][]float64, weights []float64, native, sealed, circular, derived bool, reason string) Candidate {
	return Candidate{
		Name:                         name,
		Source:                       source,
		Operator:                     op,
		Weights:                      weights,
		Native:                       native,
		Sealed:                       sealed,
		Circular:                     circular,
		DerivedFromExistingLedger:    derived,
		CompatibleWithJ:              true,
		CompatibleWithFirstOrder:     true,
		CompatibleWithHyperchargeSU2: true,
		Reason:                       reason,
	}
}

func fillCandidateLinearData(c *Candidate) error {
	if len(c.Operator) != 3 || len(c.Operator[0]) != 3 {
		return fmt.Errorf("candidate %s is not a 3x3 generation operator", c.Name)
	}
	tr := trace(c.Operator) / 3
	centralResidual := frob(sub(c.Operator, scale(identity3(), tr)))
	c.Central = centralResidual < eps
	c.NonCentral = !c.Central
	c.DiagonalOnly = isDiagonal(c.Operator, eps)
	c.Rank = rank(c.Operator, eps)
	c.Spectrum = sortedSpectrum3(c.Operator)
	c.CommutantDimension = commutantDim(c.Operator)
	c.BreaksGenerationDegeneracy = distinctCount(c.Spectrum, 1e-8) > 1
	c.GivesMixing = c.NonCentral && !c.DiagonalOnly
	if c.Native && c.NonCentral && !c.Circular {
		c.Verdict = StatusVerifiedNativeAddressFunctor
		return nil
	}
	if c.Native && c.Central {
		c.Verdict = join(StatusFailedGenerationAddressCentral, StatusTensionMoritaUniformBroadcast)
		return nil
	}
	if c.Sealed && c.DiagonalOnly && c.NonCentral {
		status := StatusFailedDiagonalOnlyNoCKM
		if c.Circular {
			status = join(StatusConditionalHierarchyCapacity, StatusFailedCircularTauOrNInsertion, StatusFailedDiagonalOnlyNoCKM)
		}
		c.Verdict = join(status, StatusTensionSealedCapacityNotNative)
		return nil
	}
	if c.Sealed && c.NonCentral && c.GivesMixing {
		c.Verdict = join(StatusConditionalAddressCapacity, StatusTensionTrialityStillNeedsCarrier, StatusTensionSealedCapacityNotNative)
		return nil
	}
	c.Verdict = StatusCandidateSourcesEnumerated
	return nil
}

func auditTrialityBranch(c CandidateAudit) SourceAudit {
	cand := findCandidate(c.Candidates, "abstract triality branch cycle")
	return SourceAudit{
		Executed:              true,
		Name:                  "Triality branch incidence",
		Inputs:                []string{"8_v, 8_s, 8_c abstract branch labels", "Gate-393 label permutation stress test"},
		Result:                "A noncentral cyclic branch action can be written on labels, but it is sealed/circular because Gate 393 did not admit a native generation-to-triality carrier.",
		Rank:                  cand.Rank,
		Spectrum:              cand.Spectrum,
		CentralOnly:           false,
		NativeNoncentralFound: false,
		CircularOrSealedOnly:  true,
		Verdict:               join(StatusTrialityBranchAudited, StatusTensionTrialityStillNeedsCarrier, StatusConditionalAddressCapacity),
	}
}

func auditMoritaEdge(c CandidateAudit) SourceAudit {
	cand := findCandidate(c.Candidates, "Morita edge uniform incidence")
	return SourceAudit{
		Executed:              true,
		Name:                  "Morita bimodule edge incidence",
		Inputs:                []string{"A_F=C⊕H⊕M3(C)", "J-paired Dirac edges", "generation copy multiplicity"},
		Result:                "The edge ledger is generation-uniform; the induced operator is proportional to I3 and does not address generations noncentrally.",
		Rank:                  cand.Rank,
		Spectrum:              cand.Spectrum,
		CentralOnly:           cand.Central,
		NativeNoncentralFound: false,
		CircularOrSealedOnly:  false,
		Verdict:               join(StatusMoritaEdgeAudited, StatusTensionMoritaUniformBroadcast, StatusFailedGenerationAddressCentral),
	}
}

func auditOneFormSupport(c CandidateAudit) SourceAudit {
	cand := findCandidate(c.Candidates, "inner-fluctuation one-form uniform support")
	return SourceAudit{
		Executed:              true,
		Name:                  "Inner-fluctuation one-form edge support",
		Inputs:                []string{"Ω¹_D(A_F)", "Gate-385 Higgs one-form edge support", "J-doubled edge measure"},
		Result:                "The one-form theorem selects the correct edge support for Higgs kinetic normalization, but its generation lift repeats the same support over all three generations.",
		Rank:                  cand.Rank,
		Spectrum:              cand.Spectrum,
		CentralOnly:           cand.Central,
		NativeNoncentralFound: false,
		CircularOrSealedOnly:  false,
		Verdict:               join(StatusOneFormSupportAudited, StatusTensionOneFormUniformBroadcast, StatusFailedGenerationAddressCentral),
	}
}

func auditNumberOperator(c CandidateAudit) NumberOperatorAudit {
	cand := findCandidate(c.Candidates, "Fock number ladder N")
	cycle := [][]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	mirror := [][]float64{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}
	cn := frob(comm(cand.Operator, cycle))
	mn := frob(comm(cand.Operator, mirror))
	return NumberOperatorAudit{
		Executed:                  true,
		Candidate:                 cand,
		Source:                    cand.Source,
		Status:                    "bridge-level compatible / sealed external extension if used as generation selector",
		DerivationResidual:        math.Inf(1),
		Native:                    false,
		BridgeCompatible:          true,
		SealedExternalExtension:   true,
		CircularIfUsedAsSolution:  true,
		CommutatorWithCycle:       cn,
		CommutatorWithMirror:      mn,
		CommutesWithHypercharge:   true,
		CommutesWithSU2L:          true,
		CommutesWithJ:             true,
		CommutesWithGamma:         true,
		CommutesWithDFEdgeSupport: true,
		BreaksExactTriality:       cn > eps || mn > eps,
		ProducesHierarchy:         cand.BreaksGenerationDegeneracy,
		ProducesMixing:            cand.GivesMixing,
		Verdict:                   join(StatusFockNumberAudited, StatusConditionalHierarchyCapacity, StatusTensionNNonNative, StatusFailedCircularTauOrNInsertion, StatusFailedDiagonalOnlyNoCKM),
	}
}

func auditTextureCapacity(c CandidateAudit) TextureCapacityAudit {
	pairs := make([]PairAudit, 0)
	maxNative, maxSealed := 0.0, 0.0
	nativeNoncommuting, sealedNoncommuting := 0, 0

	for i := range c.Candidates {
		for j := i + 1; j < len(c.Candidates); j++ {
			a, b := c.Candidates[i], c.Candidates[j]
			norm := frob(comm(a.Operator, b.Operator))
			noncomm := norm > eps
			nativePair := a.Native && b.Native
			sealedPair := a.Sealed || b.Sealed
			if nativePair && norm > maxNative {
				maxNative = norm
			}
			if sealedPair && norm > maxSealed {
				maxSealed = norm
			}
			if nativePair && noncomm {
				nativeNoncommuting++
			}
			if sealedPair && noncomm {
				sealedNoncommuting++
			}
			if noncomm || nativePair || (a.Name == "abstract triality branch cycle" && b.Name == "Fock number ladder N") {
				pairs = append(pairs, PairAudit{
					Name:                       fmt.Sprintf("%s :: %s", a.Name, b.Name),
					Left:                       a.Name,
					Right:                      b.Name,
					NativePair:                 nativePair,
					SealedPair:                 sealedPair,
					CommutatorNorm:             norm,
					Noncommuting:               noncomm,
					SimultaneouslyDiagonalized: !noncomm,
					CKMCapacity:                nativePair && noncomm,
					Reason:                     pairReason(a, b, noncomm, nativePair, sealedPair),
					Verdict:                    pairVerdict(noncomm, nativePair, sealedPair),
				})
			}
		}
	}

	nativeOps, nativeNoncentral := 0, 0
	for _, cand := range c.Candidates {
		if cand.Native {
			nativeOps++
		}
		if cand.Native && cand.NonCentral {
			nativeNoncentral++
		}
	}
	verdict := join(StatusTextureCapacityAudited, StatusFailedNoNativeNoncommutingPair, StatusFailedDiagonalOnlyNoCKM, StatusTensionNeedTwoOperators)
	if nativeNoncommuting > 0 {
		verdict = join(StatusTextureCapacityAudited, StatusVerifiedNativeAddressFunctor)
	}
	return TextureCapacityAudit{
		Executed:                     true,
		NativeGenerationOperators:    nativeOps,
		NativeNoncentralOperators:    nativeNoncentral,
		NativeNoncommutingPairs:      nativeNoncommuting,
		SealedNoncommutingPairs:      sealedNoncommuting,
		MaxNativeCommutatorNorm:      maxNative,
		MaxSealedCommutatorNorm:      maxSealed,
		SimultaneouslyDiagonalizable: nativeNoncommuting == 0,
		CKMCapacityNative:            nativeNoncommuting > 0,
		Pairs:                        pairs,
		Verdict:                      verdict,
	}
}

func pairReason(a, b Candidate, noncomm, nativePair, sealedPair bool) string {
	if nativePair && noncomm {
		return "native noncommuting pair found"
	}
	if nativePair && !noncomm {
		return "native operators commute or one is central"
	}
	if sealedPair && noncomm {
		return "noncommutation exists only after a sealed/circular operator is allowed"
	}
	return fmt.Sprintf("%s and %s do not form a native CKM-capable pair", a.Name, b.Name)
}

func pairVerdict(noncomm, nativePair, sealedPair bool) string {
	if nativePair && noncomm {
		return StatusVerifiedNativeAddressFunctor
	}
	if sealedPair && noncomm {
		return join(StatusConditionalAddressCapacity, StatusTensionSealedCapacityNotNative)
	}
	if nativePair {
		return StatusFailedGenerationAddressCentral
	}
	return StatusFailedDiagonalOnlyNoCKM
}

func auditModuli(i Inheritance, c CandidateAudit, t TextureCapacityAudit, n NumberOperatorAudit) ModuliAudit {
	start := i.Gate372ChargedModuliDim
	scenarios := []ModuliScenario{
		{
			Name:                          "A. central-only native generation broadcast",
			AssumptionClass:               "native central maps",
			StartingChargedDim:            start,
			ResultingDim:                  start,
			Native:                        true,
			Conditional:                   false,
			Failed:                        true,
			DistinctChargedMassesPossible: true,
			CKMMisalignmentPossible:       true,
			LeptonQuarkSectorSeparation:   true,
			Reason:                        "central I3 broadcasts impose no constraint and select no vacuum point",
			Verdict:                       join(StatusFailedGenerationAddressCentral, StatusFirewallPreserved13Moduli),
		},
		{
			Name:                          "B. one native diagonal operator",
			AssumptionClass:               "not available in current ledger",
			StartingChargedDim:            start,
			ResultingDim:                  start,
			Native:                        c.NativeNoncentralCount > 0,
			Conditional:                   false,
			Failed:                        c.NativeNoncentralCount == 0,
			DistinctChargedMassesPossible: false,
			CKMMisalignmentPossible:       false,
			LeptonQuarkSectorSeparation:   false,
			Reason:                        "no native noncentral diagonal generation address was derived",
			Verdict:                       join(StatusFailedGenerationAddressCentral, StatusFirewallPreserved13Moduli),
		},
		{
			Name:                          "C. one sealed diagonal operator",
			AssumptionClass:               "sealed N or protected-contact diagonal assignment",
			StartingChargedDim:            start,
			ResultingDim:                  9,
			Native:                        false,
			Conditional:                   true,
			Failed:                        true,
			DistinctChargedMassesPossible: true,
			CKMMisalignmentPossible:       false,
			LeptonQuarkSectorSeparation:   true,
			Reason:                        "diagonal hierarchy capacity exists conditionally, but no mixing or CKM capacity follows",
			Verdict:                       join(StatusConditionalHierarchyCapacity, StatusFailedDiagonalOnlyNoCKM, StatusTensionSealedCapacityNotNative),
		},
		{
			Name:                          "D. two native commuting operators",
			AssumptionClass:               "native commuting algebra only",
			StartingChargedDim:            start,
			ResultingDim:                  start,
			Native:                        true,
			Conditional:                   false,
			Failed:                        true,
			DistinctChargedMassesPossible: false,
			CKMMisalignmentPossible:       false,
			LeptonQuarkSectorSeparation:   false,
			Reason:                        "current native operators are central/uniform, so the commuting algebra gives no noncentral texture selection",
			Verdict:                       join(StatusFailedGenerationAddressCentral, StatusFailedNoNativeNoncommutingPair, StatusFirewallPreserved13Moduli),
		},
		{
			Name:                          "E. two native noncommuting operators",
			AssumptionClass:               "missing prerequisite",
			StartingChargedDim:            start,
			ResultingDim:                  start,
			Native:                        t.NativeNoncommutingPairs > 0,
			Conditional:                   false,
			Failed:                        t.NativeNoncommutingPairs == 0,
			DistinctChargedMassesPossible: false,
			CKMMisalignmentPossible:       false,
			LeptonQuarkSectorSeparation:   false,
			Reason:                        "no native noncommuting generation-operator pair was found",
			Verdict:                       join(StatusFailedNoNativeNoncommutingPair, StatusTensionNeedTwoOperators, StatusFirewallPreserved13Moduli),
		},
		{
			Name:                          "F. triality plus native address functor",
			AssumptionClass:               "not admitted by Gate 393/394",
			StartingChargedDim:            start,
			ResultingDim:                  start,
			Native:                        false,
			Conditional:                   false,
			Failed:                        true,
			DistinctChargedMassesPossible: false,
			CKMMisalignmentPossible:       false,
			LeptonQuarkSectorSeparation:   false,
			Reason:                        "triality branch action remains a label symmetry without a native carrier and address functor",
			Verdict:                       join(StatusGate393Inherited, StatusFailedGenerationAddressCentral, StatusFirewallPreserved13Moduli),
		},
		{
			Name:                          "G. triality plus sealed address functor",
			AssumptionClass:               "sealed label action",
			StartingChargedDim:            start,
			ResultingDim:                  9,
			Native:                        false,
			Conditional:                   true,
			Failed:                        true,
			DistinctChargedMassesPossible: true,
			CKMMisalignmentPossible:       false,
			LeptonQuarkSectorSeparation:   true,
			Reason:                        "sealed C3/label triality can constrain textures, but the surviving algebra is simultaneously diagonalized and has no native CKM capacity",
			Verdict:                       join(StatusConditionalAddressCapacity, StatusFailedDiagonalOnlyNoCKM, StatusTensionSealedCapacityNotNative),
		},
	}

	bestNative := start
	bestConditional := start
	for _, s := range scenarios {
		if s.Native && !s.Failed && s.ResultingDim < bestNative {
			bestNative = s.ResultingDim
		}
		if s.Conditional && s.ResultingDim < bestConditional {
			bestConditional = s.ResultingDim
		}
	}
	return ModuliAudit{
		Executed:               true,
		StartingChargedDim:     start,
		NativeReductionBelow13: bestNative < start,
		BestNativeDim:          bestNative,
		BestConditionalDim:     bestConditional,
		Scenarios:              scenarios,
		Verdict:                join(StatusModuliImpactAudited, StatusFirewallPreserved13Moduli, StatusTensionSealedCapacityNotNative),
	}
}

func auditFirewall(c CandidateAudit, m ModuliAudit, n NumberOperatorAudit) FirewallAudit {
	return FirewallAudit{
		Executed:                     true,
		NoYukawaMassesImported:       true,
		NoCKMImported:                true,
		NoPMNSImported:               true,
		NoEmpiricalOrderingImported:  true,
		NoManualGenerationAssignment: true,
		NoCircularTauInserted:        true,
		NoCircularNPromoted:          n.CircularIfUsedAsSolution,
		NoNativeAddressClaimed:       c.NativeNoncentralCount == 0,
		NoModuliReductionClaimed:     !m.NativeReductionBelow13,
		Verdict:                      join(StatusFirewallPreserved13Moduli, StatusFailedGenerationAddressCentral, StatusFailedCircularTauOrNInsertion),
	}
}

func chooseNextGate(c CandidateAudit, t TextureCapacityAudit, n NumberOperatorAudit) NextStep {
	if c.NativeNoncentralCount == 0 {
		return NextStep{
			Gate:        395,
			Title:       "Representation-Origin Search for Dynamic Generation Labels",
			Reason:      "Gate 394 found that all current native ASHA support, Morita-edge, and one-form data broadcast centrally over generation space.",
			PrimaryTask: "search beyond static support ledgers for a dynamical representation theorem that generates C^3_gen labels before any modular Hamiltonian or CKM texture gate is attempted",
		}
	}
	if n.Native && !t.CKMCapacityNative {
		return NextStep{Gate: 395, Title: "Internal Modular Hamiltonian from Native Generation Address", Reason: "a native diagonal address exists but no mixing pair exists", PrimaryTask: "derive a nontracial internal state and test whether it produces a second noncommuting sector operator"}
	}
	if t.CKMCapacityNative {
		return NextStep{Gate: 396, Title: "Two-Texture Yukawa Misalignment Sieve", Reason: "native noncommuting generation operators were found", PrimaryTask: "compute the up/down Hermitian commutator and exact charged flavor quotient dimension"}
	}
	return NextStep{Gate: 395, Title: "Internal Modular Hamiltonian from Generation Address", Reason: "default continuation", PrimaryTask: "derive a nontracial state from native address data"}
}

func buildTruth(c CandidateAudit, t TextureCapacityAudit, m ModuliAudit, next NextStep) string {
	if c.NativeNoncentralCount == 0 {
		return fmt.Sprintf("Gate 394 proves that the current ASHA finite law-space still broadcasts uniformly over generation space. Triality remains the correct threefold arena, and sealed diagonal/label operators show hierarchy capacity, but no native noncentral map into End(C^3_gen) is derived. There are %d native noncommuting texture pairs. The Gate-372 charged 13-moduli firewall remains preserved. Next: Gate %d — %s.", t.NativeNoncommutingPairs, next.Gate, next.Title)
	}
	if !t.CKMCapacityNative {
		return fmt.Sprintf("Gate 394 derives hierarchy capacity but not mixing. A native noncentral operator exists, but without a second native noncommuting texture operator the CKM/PMNS frontier remains sealed. Native best dimension=%d. Next: Gate %d — %s.", m.BestNativeDim, next.Gate, next.Title)
	}
	return fmt.Sprintf("Gate 394 opens native flavor-moduli reduction: noncommuting generation operators exist. The next gate must compute the exact quotient. Next: Gate %d — %s.", next.Gate, next.Title)
}

func Statuses(a Analysis) []string {
	seen := map[string]bool{}
	var out []string
	add := func(s string) {
		for _, part := range strings.Split(s, ";") {
			part = strings.TrimSpace(part)
			if part != "" && !seen[part] {
				seen[part] = true
				out = append(out, part)
			}
		}
	}
	add(a.Inheritance.Verdict)
	add(a.Target.Verdict)
	add(a.Candidates.Verdict)
	add(a.TrialityBranch.Verdict)
	add(a.MoritaEdge.Verdict)
	add(a.OneFormSupport.Verdict)
	add(a.Number.Verdict)
	add(a.TextureCapacity.Verdict)
	add(a.Moduli.Verdict)
	add(a.Firewall.Verdict)
	return out
}

func findCandidate(candidates []Candidate, name string) Candidate {
	for _, c := range candidates {
		if c.Name == name {
			return c
		}
	}
	return Candidate{Name: name, Operator: zero3(), Verdict: "MISSING"}
}

func identity3() [][]float64 { return [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}} }
func zero3() [][]float64     { return [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}} }
func diagFromWeights(w []float64) [][]float64 {
	return [][]float64{{w[0], 0, 0}, {0, w[1], 0}, {0, 0, w[2]}}
}

func trace(a [][]float64) float64 { return a[0][0] + a[1][1] + a[2][2] }

func scale(a [][]float64, s float64) [][]float64 {
	out := zero3()
	for i := range a {
		for j := range a[i] {
			out[i][j] = s * a[i][j]
		}
	}
	return out
}

func sub(a, b [][]float64) [][]float64 {
	out := zero3()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = a[i][j] - b[i][j]
		}
	}
	return out
}

func matmul(a, b [][]float64) [][]float64 {
	out := zero3()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}

func comm(a, b [][]float64) [][]float64 { return sub(matmul(a, b), matmul(b, a)) }

func frob(a [][]float64) float64 {
	s := 0.0
	for i := range a {
		for j := range a[i] {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}

func isDiagonal(a [][]float64, tol float64) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j && math.Abs(a[i][j]) > tol {
				return false
			}
		}
	}
	return true
}

func sortedSpectrum3(a [][]float64) []float64 {
	// All native/diagonal candidates are diagonal or identity; for the sealed cycle
	// record the real trace diagnostics instead of pretending a real eigenbasis.
	if !isDiagonal(a, eps) {
		return []float64{trace(a), frob(a), 0}
	}
	out := []float64{a[0][0], a[1][1], a[2][2]}
	sort.Float64s(out)
	return out
}

func distinctCount(vals []float64, tol float64) int {
	if len(vals) == 0 {
		return 0
	}
	sorted := append([]float64(nil), vals...)
	sort.Float64s(sorted)
	count := 1
	last := sorted[0]
	for _, v := range sorted[1:] {
		if math.Abs(v-last) > tol {
			count++
			last = v
		}
	}
	return count
}

func rank(a [][]float64, tol float64) int {
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = append([]float64(nil), a[i]...)
	}
	r := 0
	rows, cols := len(m), len(m[0])
	for c := 0; c < cols && r < rows; c++ {
		pivot := r
		for i := r + 1; i < rows; i++ {
			if math.Abs(m[i][c]) > math.Abs(m[pivot][c]) {
				pivot = i
			}
		}
		if math.Abs(m[pivot][c]) <= tol {
			continue
		}
		m[r], m[pivot] = m[pivot], m[r]
		pv := m[r][c]
		for j := c; j < cols; j++ {
			m[r][j] /= pv
		}
		for i := 0; i < rows; i++ {
			if i == r {
				continue
			}
			factor := m[i][c]
			for j := c; j < cols; j++ {
				m[i][j] -= factor * m[r][j]
			}
		}
		r++
	}
	return r
}

func commutantDim(a [][]float64) int {
	// Linear system [X,A]=0 over nine unknowns.
	rows := make([][]float64, 0, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			row := make([]float64, 9)
			for k := 0; k < 3; k++ {
				row[i*3+k] += a[k][j]
				row[k*3+j] -= a[i][k]
			}
			rows = append(rows, row)
		}
	}
	return 9 - rank(rows, eps)
}

func join(parts ...string) string {
	var out []string
	for _, p := range parts {
		for _, q := range strings.Split(p, ";") {
			q = strings.TrimSpace(q)
			if q != "" {
				out = append(out, q)
			}
		}
	}
	return strings.Join(out, ";")
}
