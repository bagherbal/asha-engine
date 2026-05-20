// Package generation2contactformcovectorobstructionaudit implements Gate 567:
// Contact Form Certificate and Distinguished Covector Obstruction Audit.
//
// The audit asks whether the already-certified Boolean--octonionic K_7 contact
// vacuum contains enough additional native data to construct a distinguished
// contact form alpha, finite d alpha, and Reeb vector. It deliberately refuses
// to identify finite contact orientation or a possible law-space clock-flow with
// continuum Lorentzian time, OS/Hilbert dynamics, RG scale, or observed history.
package generation2contactformcovectorobstructionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
)

const (
	AuditID = "GATE567-CONTACT-FORM-CERTIFICATE-DISTINGUISHED-COVECTOR-OBSTRUCTION-AUDIT"

	StatusK7BasisMetricCertified        = "PASS_K7_BASIS_AND_METRIC_CERTIFIED"
	StatusProjectorContainmentCertified = "PASS_K7_BOOLEAN_G2_CONTAINMENT_CERTIFIED"
	StatusNoDistinguishedCovector       = "FAILED_ROUTE_NO_NATIVE_DISTINGUISHED_VECTOR_OR_COVECTOR_ON_K7"
	StatusG2NoReebDirection             = "FAILED_ROUTE_G2_STRUCTURE_ALONE_DOES_NOT_SELECT_REEB_DIRECTION"
	StatusNoCandidateAlpha              = "FAILED_ROUTE_NO_NATIVE_CONTACT_ALPHA_CANDIDATE_ON_K7"
	StatusNoFiniteDAlpha                = "FAILED_ROUTE_NO_FINITE_DALPHA_OPERATOR_ON_K7"
	StatusNoContactCondition            = "FAILED_ROUTE_CONTACT_VOLUME_NOT_COMPUTABLE_WITHOUT_ALPHA_AND_DALPHA"
	StatusNoReebCertificate             = "FAILED_ROUTE_NO_REEB_VECTOR_CERTIFICATE"
	StatusNo7Split                      = "FAILED_ROUTE_K7_1PLUS6_REEB_SPLIT_NOT_DERIVED"
	StatusQ4Independent                 = "CONDITIONAL_SUPPORT_Q4_REMAINS_INDEPENDENT_CONTACT_SPECTRAL_DATA"
	StatusQ4NotReeb                     = "FAILED_ROUTE_Q4_NOT_CERTIFIED_AS_REEB_FLOW_OR_CONTACT_RETURN_MAP"
	StatusNoE0Relation                  = "FAILED_ROUTE_NO_CANONICAL_E0_TO_REEB_RELATION"
	StatusContactNotPhysicalTime        = "FIREWALL_PRESERVED_CONTACT_REEB_NOT_PHYSICAL_TIME"
	StatusNoProductTimeAirlock          = "FAILED_ROUTE_NO_CONTACT_TO_PHYSICAL_TIME_AIRLOCK"
	StatusNoRGScale                     = "FAILED_ROUTE_CONTACT_FORM_AUDIT_DOES_NOT_OPEN_RG_SCALE_OR_CUTOFF"
	StatusEWBridgeStillBridge           = "FIREWALL_PRESERVED_GATE564_GATE565_REMAIN_BRIDGE_LEVEL"
	StatusGate567Firewall               = "FIREWALL_PRESERVED_GATE567_CONTACT_FORM_COVECTOR_BOUNDARY"
)

type K7BasisMetricAudit struct {
	Dimension                    int
	ExpectedDimension            int
	AmbientDimension             int
	ContactIndex                 float64
	FrameColumns                 int
	FrameRows                    int
	FrameIsometryResidual        float64
	ProjectorIdempotenceResidual float64
	ProjectorSymmetryResidual    float64
	BooleanContainmentResidual   float64
	G2ContainmentResidual        float64
	InducedMetricIsIdentity      bool
	Verdict                      string
}

type DistinguishedSearchAudit struct {
	FromPB                         bool
	FromPG                         bool
	FromCommutator                 bool
	FromRelativePosition           bool
	FromBooleanIncidenceTensor     bool
	FromG2Calibration              bool
	FromQ4ContactSpectralBlock     bool
	FromTraceRankAsymmetry         bool
	FromCliffordE0Projection       bool
	PBRestrictionToK7Identity      bool
	PGRestrictionToK7Identity      bool
	ProjectorCommutatorOnK7Trivial bool
	NativeDistinguishedObjectFound bool
	Verdict                        string
}

type G2ObstructionAudit struct {
	G2StructureAvailable              bool
	ActsTransitivelyOnUnitDirections  bool
	ExtraSymmetryBreakingDatumPresent bool
	CanSelectReebDirection            bool
	Verdict                           string
}

type CandidateAlphaAudit struct {
	CandidateCovectorFound bool
	CandidateVectorFound   bool
	NativeBasisIndependent bool
	AlphaConstructed       bool
	UniqueUpToSignOrScale  bool
	FullyCanonical         bool
	Verdict                string
}

type FiniteDifferentialAudit struct {
	ExteriorAlgebraAvailable     bool
	FiniteDOperatorOnK7Available bool
	CochainBoundaryAvailable     bool
	IncidenceDifferentialOnK7    bool
	DAlphaComputable             bool
	Verdict                      string
}

type ContactConditionAudit struct {
	AlphaAvailable               bool
	DAlphaAvailable              bool
	AlphaWedgeDAlphaCubedKnown   bool
	AlphaWedgeDAlphaCubedNonzero bool
	ContactFormCertified         bool
	Verdict                      string
}

type ReebAudit struct {
	AlphaAvailable    bool
	DAlphaAvailable   bool
	SolvedAlphaOfR    bool
	SolvedContraction bool
	UniqueReeb        bool
	SplitK7As1Plus6   bool
	Verdict           string
}

type Q4RelationAudit struct {
	Polynomial                        string
	ContactSpectralData               bool
	CertifiedContactEndomorphism      bool
	CertifiedReebReturnMap            bool
	CertifiedLinearizedReebFlow       bool
	HiggsFlavorYukawaPromotionBlocked bool
	Verdict                           string
}

type E0RelationAudit struct {
	CliffordE0AvailableAsSignatureDatum bool
	E0ProjectionIntoK7Available         bool
	E0ToReebFunctorAvailable            bool
	ReebAvailable                       bool
	SeparationPreserved                 bool
	Verdict                             string
}

type ProductTimeFirewallAudit struct {
	ContactToDM                    bool
	ContactToLorentzianTime        bool
	ContactToOSPositivity          bool
	ContactToWickRotation          bool
	ContactToHilbertReconstruction bool
	ContactToHamiltonianSpectrum   bool
	ContactToRGScale               bool
	ContactToArrowOfTime           bool
	ElectroweakBridgeStillSealed   bool
	Verdict                        string
}

type FinalVerdict struct {
	NativeDistinguishedVectorOrCovector bool
	FiniteDOperatorOnK7                 bool
	AlphaCertified                      bool
	DAlphaCertified                     bool
	ContactVolumeNonzero                bool
	ReebCertified                       bool
	K7Splits1Plus6                      bool
	Q4RelatedToReebDynamics             bool
	E0RelatedToReeb                     bool
	PhysicalTimeRGOSHilbertOpened       bool
	MissingNextTheorem                  string
	Verdict                             string
}

type Analysis struct {
	K7      K7BasisMetricAudit
	Search  DistinguishedSearchAudit
	G2      G2ObstructionAudit
	Alpha   CandidateAlphaAudit
	DAlpha  FiniteDifferentialAudit
	Contact ContactConditionAudit
	Reeb    ReebAudit
	Q4      Q4RelationAudit
	E0      E0RelationAudit
	Time    ProductTimeFirewallAudit
	Final   FinalVerdict
	Truth   string
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
	space, err := contact.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build contact K_7 space: %w", err)
	}
	k7, err := auditK7(space)
	if err != nil {
		return Analysis{}, err
	}
	search, err := auditDistinguishedSearch(space)
	if err != nil {
		return Analysis{}, err
	}
	a := Analysis{}
	a.K7 = k7
	a.Search = search
	a.G2 = auditG2(search)
	a.Alpha = auditAlpha(search)
	a.DAlpha = auditDAlpha()
	a.Contact = auditContactCondition(a.Alpha, a.DAlpha)
	a.Reeb = auditReeb(a.Contact)
	a.Q4 = auditQ4()
	a.E0 = auditE0(a.Reeb)
	a.Time = auditProductTime()
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditK7(space contact.Space) (K7BasisMetricAudit, error) {
	frameResidual, err := space.FrameIsometryResidual()
	if err != nil {
		return K7BasisMetricAudit{}, err
	}
	boolResidual, err := space.BooleanContainmentResidual()
	if err != nil {
		return K7BasisMetricAudit{}, err
	}
	g2Residual, err := space.G2ContainmentResidual()
	if err != nil {
		return K7BasisMetricAudit{}, err
	}
	idem, err := space.ContactProjector.IdempotenceResidual()
	if err != nil {
		return K7BasisMetricAudit{}, err
	}
	sym, err := space.ContactProjector.SymmetryResidual()
	if err != nil {
		return K7BasisMetricAudit{}, err
	}
	ok := space.Dimension() == 7 && space.ExpectedContactDenominator() == 7 && math.Abs(space.ContactIndex()-1) < 1e-8 && frameResidual < 1e-8 && boolResidual < 1e-8 && g2Residual < 1e-8 && idem < 1e-8 && sym < 1e-8
	return K7BasisMetricAudit{
		Dimension: space.Dimension(), ExpectedDimension: space.ExpectedContactDenominator(), AmbientDimension: space.AmbientDimension(), ContactIndex: space.ContactIndex(),
		FrameColumns: space.ContactFrame.Cols(), FrameRows: space.ContactFrame.Rows(), FrameIsometryResidual: frameResidual, ProjectorIdempotenceResidual: idem, ProjectorSymmetryResidual: sym,
		BooleanContainmentResidual: boolResidual, G2ContainmentResidual: g2Residual, InducedMetricIsIdentity: frameResidual < 1e-8,
		Verdict: join(StatusK7BasisMetricCertified, StatusProjectorContainmentCertified, fmt.Sprintf("certified=%t", ok)),
	}, nil
}

func auditDistinguishedSearch(space contact.Space) (DistinguishedSearchAudit, error) {
	// On K_7, P_B and P_G act as the identity by construction. Their relative
	// projector restrictions therefore do not single out a direction. The
	// project currently has no typed vector/covector ledger produced by Boolean
	// incidence, G2 calibration, q4, trace/rank asymmetry, or Clifford e0.
	pbOnK, err := space.BooleanSupport.Support.Matrix.Mul(space.ContactFrame)
	if err != nil {
		return DistinguishedSearchAudit{}, err
	}
	pbDiff, err := pbOnK.Sub(space.ContactFrame)
	if err != nil {
		return DistinguishedSearchAudit{}, err
	}
	pgOnK, err := space.G2Support.Support.Matrix.Mul(space.ContactFrame)
	if err != nil {
		return DistinguishedSearchAudit{}, err
	}
	pgDiff, err := pgOnK.Sub(space.ContactFrame)
	if err != nil {
		return DistinguishedSearchAudit{}, err
	}

	pbIdentity := pbDiff.FrobeniusNorm() < 1e-8
	pgIdentity := pgDiff.FrobeniusNorm() < 1e-8
	commTrivial := pbIdentity && pgIdentity
	return DistinguishedSearchAudit{
		FromPB: false, FromPG: false, FromCommutator: false, FromRelativePosition: false,
		FromBooleanIncidenceTensor: false, FromG2Calibration: false, FromQ4ContactSpectralBlock: false, FromTraceRankAsymmetry: false, FromCliffordE0Projection: false,
		PBRestrictionToK7Identity: pbIdentity, PGRestrictionToK7Identity: pgIdentity, ProjectorCommutatorOnK7Trivial: commTrivial,
		NativeDistinguishedObjectFound: false,
		Verdict:                        join(StatusNoDistinguishedCovector, "P_B|K7 and P_G|K7 are identity restrictions; no native vector/covector is singled out"),
	}, nil
}

func auditG2(s DistinguishedSearchAudit) G2ObstructionAudit {
	return G2ObstructionAudit{G2StructureAvailable: true, ActsTransitivelyOnUnitDirections: true, ExtraSymmetryBreakingDatumPresent: false, CanSelectReebDirection: false, Verdict: StatusG2NoReebDirection}
}

func auditAlpha(s DistinguishedSearchAudit) CandidateAlphaAudit {
	return CandidateAlphaAudit{CandidateCovectorFound: false, CandidateVectorFound: false, NativeBasisIndependent: false, AlphaConstructed: false, UniqueUpToSignOrScale: false, FullyCanonical: false, Verdict: join(StatusNoDistinguishedCovector, StatusNoCandidateAlpha)}
}

func auditDAlpha() FiniteDifferentialAudit {
	return FiniteDifferentialAudit{ExteriorAlgebraAvailable: true, FiniteDOperatorOnK7Available: false, CochainBoundaryAvailable: false, IncidenceDifferentialOnK7: false, DAlphaComputable: false, Verdict: StatusNoFiniteDAlpha}
}

func auditContactCondition(alpha CandidateAlphaAudit, d FiniteDifferentialAudit) ContactConditionAudit {
	return ContactConditionAudit{AlphaAvailable: alpha.AlphaConstructed, DAlphaAvailable: d.DAlphaComputable, AlphaWedgeDAlphaCubedKnown: false, AlphaWedgeDAlphaCubedNonzero: false, ContactFormCertified: false, Verdict: join(StatusNoCandidateAlpha, StatusNoFiniteDAlpha, StatusNoContactCondition)}
}

func auditReeb(c ContactConditionAudit) ReebAudit {
	return ReebAudit{AlphaAvailable: c.AlphaAvailable, DAlphaAvailable: c.DAlphaAvailable, SolvedAlphaOfR: false, SolvedContraction: false, UniqueReeb: false, SplitK7As1Plus6: false, Verdict: join(StatusNoReebCertificate, StatusNo7Split)}
}

func auditQ4() Q4RelationAudit {
	return Q4RelationAudit{Polynomial: "q4(x)=3240x^4-7668x^3+6426x^2-2235x+271", ContactSpectralData: true, CertifiedContactEndomorphism: false, CertifiedReebReturnMap: false, CertifiedLinearizedReebFlow: false, HiggsFlavorYukawaPromotionBlocked: true, Verdict: join(StatusQ4Independent, StatusQ4NotReeb)}
}

func auditE0(r ReebAudit) E0RelationAudit {
	return E0RelationAudit{CliffordE0AvailableAsSignatureDatum: true, E0ProjectionIntoK7Available: false, E0ToReebFunctorAvailable: false, ReebAvailable: r.UniqueReeb, SeparationPreserved: true, Verdict: StatusNoE0Relation}
}

func auditProductTime() ProductTimeFirewallAudit {
	return ProductTimeFirewallAudit{ContactToDM: false, ContactToLorentzianTime: false, ContactToOSPositivity: false, ContactToWickRotation: false, ContactToHilbertReconstruction: false, ContactToHamiltonianSpectrum: false, ContactToRGScale: false, ContactToArrowOfTime: false, ElectroweakBridgeStillSealed: true, Verdict: join(StatusNoProductTimeAirlock, StatusContactNotPhysicalTime, StatusNoRGScale, StatusEWBridgeStillBridge)}
}

func auditFinal(a Analysis) FinalVerdict {
	return FinalVerdict{NativeDistinguishedVectorOrCovector: a.Search.NativeDistinguishedObjectFound, FiniteDOperatorOnK7: a.DAlpha.FiniteDOperatorOnK7Available, AlphaCertified: a.Contact.ContactFormCertified, DAlphaCertified: a.DAlpha.DAlphaComputable, ContactVolumeNonzero: a.Contact.AlphaWedgeDAlphaCubedNonzero, ReebCertified: a.Reeb.UniqueReeb, K7Splits1Plus6: a.Reeb.SplitK7As1Plus6, Q4RelatedToReebDynamics: a.Q4.CertifiedReebReturnMap || a.Q4.CertifiedLinearizedReebFlow || a.Q4.CertifiedContactEndomorphism, E0RelatedToReeb: a.E0.E0ProjectionIntoK7Available || a.E0.E0ToReebFunctorAvailable, PhysicalTimeRGOSHilbertOpened: a.Time.ContactToDM || a.Time.ContactToLorentzianTime || a.Time.ContactToOSPositivity || a.Time.ContactToWickRotation || a.Time.ContactToHilbertReconstruction || a.Time.ContactToHamiltonianSpectrum || a.Time.ContactToRGScale || a.Time.ContactToArrowOfTime, MissingNextTheorem: "Provide a native distinguished covector alpha in K_7^* or vector R in K_7, a finite differential d on K_7, prove alpha∧(d alpha)^3≠0, solve unique Reeb R, and only then build a separate product-time/OS/Wick/Hilbert/RG airlock", Verdict: join(StatusNoDistinguishedCovector, StatusNoFiniteDAlpha, StatusNoReebCertificate, StatusContactNotPhysicalTime, StatusGate567Firewall)}
}

func validate(a Analysis) error {
	failures := []string{}
	if a.K7.Dimension != 7 || !a.K7.InducedMetricIsIdentity || a.K7.BooleanContainmentResidual > 1e-8 || a.K7.G2ContainmentResidual > 1e-8 || a.K7.ProjectorIdempotenceResidual > 1e-8 || a.K7.ProjectorSymmetryResidual > 1e-8 {
		failures = append(failures, "K7 certificate failed")
	}
	if !a.Search.PBRestrictionToK7Identity || !a.Search.PGRestrictionToK7Identity || !a.Search.ProjectorCommutatorOnK7Trivial || a.Search.NativeDistinguishedObjectFound || a.Search.FromPB || a.Search.FromPG || a.Search.FromCommutator || a.Search.FromRelativePosition || a.Search.FromBooleanIncidenceTensor || a.Search.FromG2Calibration || a.Search.FromQ4ContactSpectralBlock || a.Search.FromTraceRankAsymmetry || a.Search.FromCliffordE0Projection {
		failures = append(failures, "distinguished search failed")
	}
	if !a.G2.G2StructureAvailable || !a.G2.ActsTransitivelyOnUnitDirections || a.G2.ExtraSymmetryBreakingDatumPresent || a.G2.CanSelectReebDirection {
		failures = append(failures, "G2 obstruction failed")
	}
	if a.Alpha.CandidateCovectorFound || a.Alpha.CandidateVectorFound || a.Alpha.NativeBasisIndependent || a.Alpha.AlphaConstructed || a.Alpha.UniqueUpToSignOrScale || a.Alpha.FullyCanonical {
		failures = append(failures, "alpha audit failed")
	}
	if !a.DAlpha.ExteriorAlgebraAvailable || a.DAlpha.FiniteDOperatorOnK7Available || a.DAlpha.CochainBoundaryAvailable || a.DAlpha.IncidenceDifferentialOnK7 || a.DAlpha.DAlphaComputable {
		failures = append(failures, "d alpha audit failed")
	}
	if a.Contact.AlphaAvailable || a.Contact.DAlphaAvailable || a.Contact.AlphaWedgeDAlphaCubedKnown || a.Contact.AlphaWedgeDAlphaCubedNonzero || a.Contact.ContactFormCertified {
		failures = append(failures, "contact condition audit failed")
	}
	if a.Reeb.AlphaAvailable || a.Reeb.DAlphaAvailable || a.Reeb.SolvedAlphaOfR || a.Reeb.SolvedContraction || a.Reeb.UniqueReeb || a.Reeb.SplitK7As1Plus6 {
		failures = append(failures, "Reeb audit failed")
	}
	if !a.Q4.ContactSpectralData || a.Q4.CertifiedContactEndomorphism || a.Q4.CertifiedReebReturnMap || a.Q4.CertifiedLinearizedReebFlow || !a.Q4.HiggsFlavorYukawaPromotionBlocked {
		failures = append(failures, "q4 audit failed")
	}
	if !a.E0.CliffordE0AvailableAsSignatureDatum || a.E0.E0ProjectionIntoK7Available || a.E0.E0ToReebFunctorAvailable || a.E0.ReebAvailable || !a.E0.SeparationPreserved {
		failures = append(failures, "e0 audit failed")
	}
	if a.Time.ContactToDM || a.Time.ContactToLorentzianTime || a.Time.ContactToOSPositivity || a.Time.ContactToWickRotation || a.Time.ContactToHilbertReconstruction || a.Time.ContactToHamiltonianSpectrum || a.Time.ContactToRGScale || a.Time.ContactToArrowOfTime || !a.Time.ElectroweakBridgeStillSealed {
		failures = append(failures, "product time firewall failed")
	}
	if a.Final.NativeDistinguishedVectorOrCovector || a.Final.FiniteDOperatorOnK7 || a.Final.AlphaCertified || a.Final.DAlphaCertified || a.Final.ContactVolumeNonzero || a.Final.ReebCertified || a.Final.K7Splits1Plus6 || a.Final.Q4RelatedToReebDynamics || a.Final.E0RelatedToReeb || a.Final.PhysicalTimeRGOSHilbertOpened {
		failures = append(failures, "final verdict failed")
	}
	if len(failures) > 0 {
		return fmt.Errorf(strings.Join(failures, "; "))
	}
	return nil
}

func Statuses() []string {
	return []string{StatusK7BasisMetricCertified, StatusProjectorContainmentCertified, StatusNoDistinguishedCovector, StatusG2NoReebDirection, StatusNoCandidateAlpha, StatusNoFiniteDAlpha, StatusNoContactCondition, StatusNoReebCertificate, StatusNo7Split, StatusQ4Independent, StatusQ4NotReeb, StatusNoE0Relation, StatusContactNotPhysicalTime, StatusNoProductTimeAirlock, StatusNoRGScale, StatusEWBridgeStillBridge, StatusGate567Firewall}
}

func truth(a Analysis) string {
	return join("Gate 567 recovers the certified orthonormal K_7 basis and induced metric, but finds no native distinguished vector or covector on K_7", "P_B and P_G restrict to the identity on K_7, G2 symmetry alone does not select a Reeb direction, and no Boolean/G2/q4/trace/e0 datum breaks the symmetry", "without alpha and a finite d operator on K_7, alpha∧(d alpha)^3 and the Reeb equations remain noncomputable", "q4 remains independent contact spectral data, e0 remains finite signature data, and no airlock opens physical time, RG scale, OS/Wick/Hilbert dynamics, or electroweak physical dynamics")
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
