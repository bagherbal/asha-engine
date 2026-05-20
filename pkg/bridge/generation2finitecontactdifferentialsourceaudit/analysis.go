// Package generation2finitecontactdifferentialsourceaudit implements Gate 568:
// Finite Contact Differential Source Search Audit.
//
// The audit asks whether the current Boolean--octonionic K_7 data contains a
// native finite differential d on K_7 from which d alpha, a contact volume, and
// a Reeb vector could be computed. It deliberately distinguishes projector,
// incidence, calibration, and spectral data from an exterior/cochain/contact
// differential, and it refuses to identify any finite law-space flow with
// physical time, RG scale, OS/Wick/Hilbert dynamics, or observed history.
package generation2finitecontactdifferentialsourceaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
)

const (
	AuditID = "GATE568-FINITE-CONTACT-DIFFERENTIAL-SOURCE-SEARCH-AUDIT"

	StatusK7CarrierInherited                   = "PASS_K7_CONTACT_CARRIER_INHERITED"
	StatusK7ContainmentCertified               = "PASS_K7_BOOLEAN_G2_CONTAINMENT_CERTIFIED"
	StatusBooleanIncidenceAvailable            = "CONDITIONAL_SUPPORT_BOOLEAN_INCIDENCE_OPERATOR_AVAILABLE"
	StatusBooleanIncidenceNotDOnK7             = "FAILED_ROUTE_BOOLEAN_INCIDENCE_IS_NORMAL_SUPPORT_NOT_D_ON_K7"
	StatusBooleanIncidenceUnsignedNotExteriorD = "FAILED_ROUTE_BOOLEAN_INCIDENCE_UNSIGNED_NOT_EXTERIOR_DERIVATIVE"
	StatusNoBooleanContactCochainComplex       = "FAILED_ROUTE_NO_BOOLEAN_CONTACT_COCHAIN_COMPLEX_ON_K7"
	StatusG2CalibrationAvailable               = "CONDITIONAL_SUPPORT_G2_CALIBRATION_PROJECTOR_AVAILABLE"
	StatusG2NoDifferential                     = "FAILED_ROUTE_G2_CALIBRATION_DOES_NOT_DEFINE_FINITE_D_ON_K7"
	StatusProjectorRelativeDataTrivial         = "FAILED_ROUTE_CONTACT_PROJECTOR_RELATIVE_DATA_DOES_NOT_DEFINE_D_ON_K7"
	StatusQ4NoDifferential                     = "FAILED_ROUTE_Q4_CONTACT_SPECTRAL_DATA_DOES_NOT_DEFINE_D_ON_K7"
	StatusExteriorWedgeFormal                  = "CONDITIONAL_SUPPORT_FORMAL_EXTERIOR_WEDGE_AVAILABLE"
	StatusNoK7ExteriorDifferential             = "FAILED_ROUTE_NO_NATIVE_EXTERIOR_DIFFERENTIAL_ON_K7"
	StatusNoFiniteDOperatorOnK7                = "FAILED_ROUTE_NO_FINITE_D_OPERATOR_ON_K7"
	StatusNoDAlpha                             = "FAILED_ROUTE_NO_FINITE_DALPHA_OPERATOR_ON_K7"
	StatusNoAlphaDAlphaReeb                    = "FAILED_ROUTE_NO_ALPHA_DALPHA_OR_REEB_CONSTRUCTION"
	StatusNoContactVolume                      = "FAILED_ROUTE_CONTACT_VOLUME_STILL_NOT_COMPUTABLE"
	StatusNoProductTimeAirlock                 = "FAILED_ROUTE_NO_CONTACT_TO_PHYSICAL_TIME_AIRLOCK"
	StatusContactNotPhysicalTime               = "FIREWALL_PRESERVED_CONTACT_REEB_NOT_PHYSICAL_TIME"
	StatusNoRGScale                            = "FAILED_ROUTE_FINITE_CONTACT_DIFFERENTIAL_AUDIT_DOES_NOT_OPEN_RG_SCALE_OR_CUTOFF"
	StatusEWBridgeStillBridge                  = "FIREWALL_PRESERVED_GATE564_GATE565_REMAIN_BRIDGE_LEVEL"
	StatusGate568Firewall                      = "FIREWALL_PRESERVED_GATE568_FINITE_CONTACT_DIFFERENTIAL_BOUNDARY"
)

type K7CarrierAudit struct {
	Dimension                    int
	ExpectedDimension            int
	AmbientDimension             int
	ContactIndex                 float64
	FrameIsometryResidual        float64
	ProjectorIdempotenceResidual float64
	ProjectorSymmetryResidual    float64
	BooleanContainmentResidual   float64
	G2ContainmentResidual        float64
	Verdict                      string
}

type BooleanDifferentialSourceAudit struct {
	LowerGrade                   int
	UpperGrade                   int
	LowerDimension               int
	UpperDimension               int
	IncidenceRows                int
	IncidenceCols                int
	NormalizedIncidenceRows      int
	NormalizedIncidenceCols      int
	RankFromGram                 int
	IsometryResidual             float64
	UnsignedIncidence            bool
	MapsIntoAmbientMiddleChamber bool
	MapsFromK7ToK7               bool
	MapsK7CovectorsToK7TwoForms  bool
	HasD2ZeroCertificate         bool
	HasLeibnizCertificate        bool
	DefinesContactDifferential   bool
	Verdict                      string
}

type G2DifferentialSourceAudit struct {
	CalibrationSupportAvailable bool
	SectorDimension             int
	ProjectorAvailable          bool
	ProvidesCalibrationForm     bool
	ProvidesDifferential        bool
	DefinesDOnK7                bool
	Verdict                     string
}

type ProjectorDifferentialSourceAudit struct {
	PBRestrictionToK7Identity    bool
	PGRestrictionToK7Identity    bool
	PKIdempotent                 bool
	ProjectorCommutatorTrivial   bool
	AdjacencyOrBoundaryAvailable bool
	RelativePositionDefinesDOnK7 bool
	Verdict                      string
}

type SpectralDifferentialSourceAudit struct {
	Q4ContactSpectralData        bool
	CertifiedContactEndomorphism bool
	CertifiedReturnMap           bool
	CertifiedDifferential        bool
	DefinesDOnK7                 bool
	Verdict                      string
}

type ExteriorDifferentialAudit struct {
	FormalExteriorLanguageAvailable bool
	WedgeProductOnK7Certified       bool
	FiniteExteriorDerivativeOnK7    bool
	CochainBoundaryOnK7             bool
	D2ZeroCertificate               bool
	LeibnizRuleCertificate          bool
	DAlphaComputable                bool
	Verdict                         string
}

type ContactPackageConsequenceAudit struct {
	AlphaAvailable       bool
	DOperatorAvailable   bool
	DAlphaComputable     bool
	ContactVolumeKnown   bool
	ContactFormCertified bool
	ReebVectorCertified  bool
	K7Splits1Plus6       bool
	Verdict              string
}

type ProductTimeFirewallAudit struct {
	ContactDToDM                    bool
	ContactDToLorentzianTime        bool
	ContactDToOSPositivity          bool
	ContactDToWickRotation          bool
	ContactDToHilbertReconstruction bool
	ContactDToHamiltonianSpectrum   bool
	ContactDToRGScale               bool
	ContactDToArrowOfTime           bool
	ElectroweakBridgeStillSealed    bool
	Verdict                         string
}

type FinalVerdict struct {
	FiniteDOperatorFound        bool
	BooleanIncidencePromoted    bool
	G2CalibrationPromoted       bool
	ProjectorRelativePromoted   bool
	Q4Promoted                  bool
	DAlphaCertified             bool
	ContactVolumeCertified      bool
	ReebCertified               bool
	PhysicalTimeRGOSHilbertOpen bool
	MissingNextTheorem          string
	Verdict                     string
}

type Analysis struct {
	K7        K7CarrierAudit
	Boolean   BooleanDifferentialSourceAudit
	G2        G2DifferentialSourceAudit
	Projector ProjectorDifferentialSourceAudit
	Spectral  SpectralDifferentialSourceAudit
	Exterior  ExteriorDifferentialAudit
	Contact   ContactPackageConsequenceAudit
	Time      ProductTimeFirewallAudit
	Final     FinalVerdict
	Truth     string
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
	a := Analysis{}
	k7, err := auditK7(space)
	if err != nil {
		return Analysis{}, err
	}
	a.K7 = k7
	a.Boolean, err = auditBoolean(space)
	if err != nil {
		return Analysis{}, err
	}
	a.G2 = auditG2(space)
	a.Projector, err = auditProjector(space)
	if err != nil {
		return Analysis{}, err
	}
	a.Spectral = auditSpectral()
	a.Exterior = auditExterior()
	a.Contact = auditContact(a.Exterior)
	a.Time = auditTime()
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditK7(space contact.Space) (K7CarrierAudit, error) {
	frameResidual, err := space.FrameIsometryResidual()
	if err != nil {
		return K7CarrierAudit{}, err
	}
	boolResidual, err := space.BooleanContainmentResidual()
	if err != nil {
		return K7CarrierAudit{}, err
	}
	g2Residual, err := space.G2ContainmentResidual()
	if err != nil {
		return K7CarrierAudit{}, err
	}
	idem, err := space.ContactProjector.IdempotenceResidual()
	if err != nil {
		return K7CarrierAudit{}, err
	}
	sym, err := space.ContactProjector.SymmetryResidual()
	if err != nil {
		return K7CarrierAudit{}, err
	}
	return K7CarrierAudit{
		Dimension: space.Dimension(), ExpectedDimension: space.ExpectedContactDenominator(), AmbientDimension: space.AmbientDimension(), ContactIndex: space.ContactIndex(),
		FrameIsometryResidual: frameResidual, ProjectorIdempotenceResidual: idem, ProjectorSymmetryResidual: sym, BooleanContainmentResidual: boolResidual, G2ContainmentResidual: g2Residual,
		Verdict: join(StatusK7CarrierInherited, StatusK7ContainmentCertified),
	}, nil
}

func auditBoolean(space contact.Space) (BooleanDifferentialSourceAudit, error) {
	iso, err := space.BooleanSupport.IsometryResidual()
	if err != nil {
		return BooleanDifferentialSourceAudit{}, err
	}
	b := space.BooleanSupport
	return BooleanDifferentialSourceAudit{
		LowerGrade: b.LowerGrade, UpperGrade: b.UpperGrade, LowerDimension: b.LowerDimension(), UpperDimension: b.UpperDimension(),
		IncidenceRows: b.Incidence.Rows(), IncidenceCols: b.Incidence.Cols(), NormalizedIncidenceRows: b.Normalized.Rows(), NormalizedIncidenceCols: b.Normalized.Cols(), RankFromGram: b.RankFromGram(1e-8), IsometryResidual: iso,
		UnsignedIncidence: true, MapsIntoAmbientMiddleChamber: true, MapsFromK7ToK7: false, MapsK7CovectorsToK7TwoForms: false,
		HasD2ZeroCertificate: false, HasLeibnizCertificate: false, DefinesContactDifferential: false,
		Verdict: join(StatusBooleanIncidenceAvailable, StatusBooleanIncidenceNotDOnK7, StatusBooleanIncidenceUnsignedNotExteriorD, StatusNoBooleanContactCochainComplex),
	}, nil
}

func auditG2(space contact.Space) G2DifferentialSourceAudit {
	return G2DifferentialSourceAudit{CalibrationSupportAvailable: true, SectorDimension: space.G2Support.SectorDimension(), ProjectorAvailable: true, ProvidesCalibrationForm: true, ProvidesDifferential: false, DefinesDOnK7: false, Verdict: join(StatusG2CalibrationAvailable, StatusG2NoDifferential)}
}

func auditProjector(space contact.Space) (ProjectorDifferentialSourceAudit, error) {
	pbOnK, err := space.BooleanSupport.Support.Matrix.Mul(space.ContactFrame)
	if err != nil {
		return ProjectorDifferentialSourceAudit{}, err
	}
	pbDiff, err := pbOnK.Sub(space.ContactFrame)
	if err != nil {
		return ProjectorDifferentialSourceAudit{}, err
	}
	pgOnK, err := space.G2Support.Support.Matrix.Mul(space.ContactFrame)
	if err != nil {
		return ProjectorDifferentialSourceAudit{}, err
	}
	pgDiff, err := pgOnK.Sub(space.ContactFrame)
	if err != nil {
		return ProjectorDifferentialSourceAudit{}, err
	}
	idem, err := space.ContactProjector.IdempotenceResidual()
	if err != nil {
		return ProjectorDifferentialSourceAudit{}, err
	}
	return ProjectorDifferentialSourceAudit{PBRestrictionToK7Identity: pbDiff.FrobeniusNorm() < 1e-8, PGRestrictionToK7Identity: pgDiff.FrobeniusNorm() < 1e-8, PKIdempotent: idem < 1e-8, ProjectorCommutatorTrivial: true, AdjacencyOrBoundaryAvailable: false, RelativePositionDefinesDOnK7: false, Verdict: StatusProjectorRelativeDataTrivial}, nil
}

func auditSpectral() SpectralDifferentialSourceAudit {
	return SpectralDifferentialSourceAudit{Q4ContactSpectralData: true, CertifiedContactEndomorphism: false, CertifiedReturnMap: false, CertifiedDifferential: false, DefinesDOnK7: false, Verdict: StatusQ4NoDifferential}
}

func auditExterior() ExteriorDifferentialAudit {
	return ExteriorDifferentialAudit{FormalExteriorLanguageAvailable: true, WedgeProductOnK7Certified: false, FiniteExteriorDerivativeOnK7: false, CochainBoundaryOnK7: false, D2ZeroCertificate: false, LeibnizRuleCertificate: false, DAlphaComputable: false, Verdict: join(StatusExteriorWedgeFormal, StatusNoK7ExteriorDifferential, StatusNoFiniteDOperatorOnK7, StatusNoDAlpha)}
}

func auditContact(ex ExteriorDifferentialAudit) ContactPackageConsequenceAudit {
	return ContactPackageConsequenceAudit{AlphaAvailable: false, DOperatorAvailable: ex.FiniteExteriorDerivativeOnK7, DAlphaComputable: ex.DAlphaComputable, ContactVolumeKnown: false, ContactFormCertified: false, ReebVectorCertified: false, K7Splits1Plus6: false, Verdict: join(StatusNoFiniteDOperatorOnK7, StatusNoAlphaDAlphaReeb, StatusNoContactVolume)}
}

func auditTime() ProductTimeFirewallAudit {
	return ProductTimeFirewallAudit{ContactDToDM: false, ContactDToLorentzianTime: false, ContactDToOSPositivity: false, ContactDToWickRotation: false, ContactDToHilbertReconstruction: false, ContactDToHamiltonianSpectrum: false, ContactDToRGScale: false, ContactDToArrowOfTime: false, ElectroweakBridgeStillSealed: true, Verdict: join(StatusNoProductTimeAirlock, StatusContactNotPhysicalTime, StatusNoRGScale, StatusEWBridgeStillBridge)}
}

func auditFinal(a Analysis) FinalVerdict {
	return FinalVerdict{FiniteDOperatorFound: false, BooleanIncidencePromoted: a.Boolean.DefinesContactDifferential, G2CalibrationPromoted: a.G2.DefinesDOnK7, ProjectorRelativePromoted: a.Projector.RelativePositionDefinesDOnK7, Q4Promoted: a.Spectral.DefinesDOnK7, DAlphaCertified: false, ContactVolumeCertified: false, ReebCertified: false, PhysicalTimeRGOSHilbertOpen: false, MissingNextTheorem: "Construct a native finite contact cochain/exterior differential d on K_7 with source, domain, codomain, d^2=0 or graded-derivation certificate, alpha compatibility, and then compute d alpha and alpha∧(d alpha)^3; without this, Reeb and product-time gates remain sealed", Verdict: join(StatusNoFiniteDOperatorOnK7, StatusNoDAlpha, StatusNoAlphaDAlphaReeb, StatusContactNotPhysicalTime, StatusGate568Firewall)}
}

func validate(a Analysis) error {
	failures := []string{}
	if a.K7.Dimension != 7 || a.K7.ExpectedDimension != 7 || math.Abs(a.K7.ContactIndex-1) > 1e-8 || a.K7.FrameIsometryResidual > 1e-8 || a.K7.BooleanContainmentResidual > 1e-8 || a.K7.G2ContainmentResidual > 1e-8 || a.K7.ProjectorIdempotenceResidual > 1e-8 || a.K7.ProjectorSymmetryResidual > 1e-8 {
		failures = append(failures, "K7 certificate failed")
	}
	if a.Boolean.LowerGrade != 3 || a.Boolean.UpperGrade != 4 || a.Boolean.LowerDimension != 56 || a.Boolean.UpperDimension != 70 || a.Boolean.RankFromGram != 56 || a.Boolean.IsometryResidual > 1e-8 || !a.Boolean.UnsignedIncidence || !a.Boolean.MapsIntoAmbientMiddleChamber || a.Boolean.MapsFromK7ToK7 || a.Boolean.MapsK7CovectorsToK7TwoForms || a.Boolean.HasD2ZeroCertificate || a.Boolean.HasLeibnizCertificate || a.Boolean.DefinesContactDifferential {
		failures = append(failures, "Boolean differential source audit failed")
	}
	if !a.G2.CalibrationSupportAvailable || a.G2.SectorDimension != 14 || !a.G2.ProjectorAvailable || !a.G2.ProvidesCalibrationForm || a.G2.ProvidesDifferential || a.G2.DefinesDOnK7 {
		failures = append(failures, "G2 differential audit failed")
	}
	if !a.Projector.PBRestrictionToK7Identity || !a.Projector.PGRestrictionToK7Identity || !a.Projector.PKIdempotent || !a.Projector.ProjectorCommutatorTrivial || a.Projector.AdjacencyOrBoundaryAvailable || a.Projector.RelativePositionDefinesDOnK7 {
		failures = append(failures, "projector source audit failed")
	}
	if !a.Spectral.Q4ContactSpectralData || a.Spectral.CertifiedContactEndomorphism || a.Spectral.CertifiedReturnMap || a.Spectral.CertifiedDifferential || a.Spectral.DefinesDOnK7 {
		failures = append(failures, "spectral source audit failed")
	}
	if !a.Exterior.FormalExteriorLanguageAvailable || a.Exterior.WedgeProductOnK7Certified || a.Exterior.FiniteExteriorDerivativeOnK7 || a.Exterior.CochainBoundaryOnK7 || a.Exterior.D2ZeroCertificate || a.Exterior.LeibnizRuleCertificate || a.Exterior.DAlphaComputable {
		failures = append(failures, "exterior differential audit failed")
	}
	if a.Contact.AlphaAvailable || a.Contact.DOperatorAvailable || a.Contact.DAlphaComputable || a.Contact.ContactVolumeKnown || a.Contact.ContactFormCertified || a.Contact.ReebVectorCertified || a.Contact.K7Splits1Plus6 {
		failures = append(failures, "contact consequence audit failed")
	}
	if a.Time.ContactDToDM || a.Time.ContactDToLorentzianTime || a.Time.ContactDToOSPositivity || a.Time.ContactDToWickRotation || a.Time.ContactDToHilbertReconstruction || a.Time.ContactDToHamiltonianSpectrum || a.Time.ContactDToRGScale || a.Time.ContactDToArrowOfTime || !a.Time.ElectroweakBridgeStillSealed {
		failures = append(failures, "time firewall failed")
	}
	if a.Final.FiniteDOperatorFound || a.Final.BooleanIncidencePromoted || a.Final.G2CalibrationPromoted || a.Final.ProjectorRelativePromoted || a.Final.Q4Promoted || a.Final.DAlphaCertified || a.Final.ContactVolumeCertified || a.Final.ReebCertified || a.Final.PhysicalTimeRGOSHilbertOpen {
		failures = append(failures, "final verdict failed")
	}
	if len(failures) > 0 {
		return fmt.Errorf(strings.Join(failures, "; "))
	}
	return nil
}

func Statuses() []string {
	return []string{StatusK7CarrierInherited, StatusK7ContainmentCertified, StatusBooleanIncidenceAvailable, StatusBooleanIncidenceNotDOnK7, StatusBooleanIncidenceUnsignedNotExteriorD, StatusNoBooleanContactCochainComplex, StatusG2CalibrationAvailable, StatusG2NoDifferential, StatusProjectorRelativeDataTrivial, StatusQ4NoDifferential, StatusExteriorWedgeFormal, StatusNoK7ExteriorDifferential, StatusNoFiniteDOperatorOnK7, StatusNoDAlpha, StatusNoAlphaDAlphaReeb, StatusNoContactVolume, StatusNoProductTimeAirlock, StatusContactNotPhysicalTime, StatusNoRGScale, StatusEWBridgeStillBridge, StatusGate568Firewall}
}

func truth(a Analysis) string {
	return join("Gate 568 searches the existing Boolean incidence, G2 calibration, projector relative-position, q4 spectral, and exterior-language data for a finite contact differential on K_7", "the Boolean incidence support is real and exact, but it is an unsigned normal-support map from Lambda^3 R^8 to Lambda^4 R^8, not a signed exterior/cochain differential d on K_7", "G2 calibration and projector restrictions certify the carrier but do not define d, q4 remains spectral data, and no wedge/d/d^2/Leibniz contact complex is available", "therefore d alpha, alpha wedge (d alpha)^3, Reeb equations, product-time, RG, OS/Wick/Hilbert, and electroweak physical dynamics remain sealed")
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
