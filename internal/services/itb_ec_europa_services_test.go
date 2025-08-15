package services

import (
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients"
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients/enums"
	"testing"
)

const validDcatAp = `@prefix dcat: <http://www.w3.org/ns/dcat#> .
@prefix dct: <http://purl.org/dc/terms/> .
@prefix ex: <http://example.com/ns#> .

ex:ValidExample a dcat:Dataset ;
dct:title "Example Valid Dataset" ;
dct:description "This is an example of a dataset that should pass the ITB validation." .
`

const invalidDcatAp = `@prefix dcat: <http://www.w3.org/ns/dcat#> .
@prefix dct: <http://purl.org/dc/terms/> .
@prefix ex: <http://example.com/ns#> .

ex:ValidExample a dcat:Dataset ;
dct:title "Example Valid Dataset" ;
dct:description "This is an example of a dataset that should fail the ITB validation." ;
dcat:theme <http://www.wikidata.org/entity/Q14944328> .
`

const validHealthri = `@prefix dcat: <http://www.w3.org/ns/dcat#> .
@prefix dcterms: <http://purl.org/dc/terms/> .
@prefix foaf: <http://xmlns.com/foaf/0.1/> .
@prefix ns1: <http://data.europa.eu/r5r/> .
@prefix v: <http://www.w3.org/2006/vcard/ns#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .

<http://www.example.com/dataset/ZLOYOJ> a dcat:Dataset ;
ns1:applicableLegislation <http://data.europa.eu/eli/reg/2025/327/oj> ;
dcterms:accessRights <http://publications.europa.eu/resource/authority/access-right/RESTRICTED> ;
dcterms:creator [ a foaf:Agent ;
    dcterms:identifier "https://ror.org/05wg1m734" ;
    foaf:homepage <https://www.xumc.nl/> ;
    foaf:mbox <mailto:data-access-committee@xumc.nl> ;
    foaf:name "Academic Medical Center" ] ;
dcterms:description "The primary aim of the PRISMA study was to investigate the potential value of risk-tailored versus traditional breast cancer screening protocols in the Netherlands. Data collection took place between 2014-2019, resulting in ∼67,000 mammograms, ∼38,000 surveys, ∼10,000 blood samples and ∼600 saliva samples." ;
dcterms:identifier "http://www.example.com/dataset/ZLOYOJ" ;
dcterms:issued "2024-07-01T11:11:11+00:00"^^xsd:dateTime ;
dcterms:license <https://creativecommons.org/licenses/by-sa/4.0/> ;
dcterms:modified "2024-06-04T13:36:10+00:00"^^xsd:dateTime ;
dcterms:publisher [ a foaf:Agent ;
    dcterms:identifier "https://ror.org/05wg1m734" ;
    foaf:homepage <https://www.xumc.nl/> ;
    foaf:mbox <mailto:data-access-committee@xumc.nl> ;
    foaf:name "Academic Medical Center" ] ;
dcterms:title "Questionnaire data of the Personalised RISk-based MAmmascreening Study (PRISMA)" ;
dcat:contactPoint [ a v:Kind ;
    v:fn "Data Access Committee of the x UMC" ;
    v:hasEmail <mailto:data-access-committee@xumc.nl> ] ;
dcat:keyword "example" ;
dcat:theme <http://publications.europa.eu/resource/authority/data-theme/HEAL> .
`

const invalidHealthri = `@prefix dcat: <http://www.w3.org/ns/dcat#> .
@prefix dcterms: <http://purl.org/dc/terms/> .
@prefix foaf: <http://xmlns.com/foaf/0.1/> .
@prefix ns1: <http://data.europa.eu/r5r/> .
@prefix v: <http://www.w3.org/2006/vcard/ns#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .

<http://www.example.com/dataset/ZLOYOJ> a dcat:Dataset ;
ns1:applicableLegislation <http://data.europa.eu/eli/reg/2025/327/oj> ;
dcterms:creator [ a foaf:Agent ;
    dcterms:identifier "https://ror.org/05wg1m734" ;
    foaf:homepage <https://www.xumc.nl/> ;
    foaf:mbox <mailto:data-access-committee@xumc.nl> ;
    foaf:name "Academic Medical Center" ] ;
dcterms:description "The primary aim of the PRISMA study was to investigate the potential value of risk-tailored versus traditional breast cancer screening protocols in the Netherlands. Data collection took place between 2014-2019, resulting in ∼67,000 mammograms, ∼38,000 surveys, ∼10,000 blood samples and ∼600 saliva samples." ;
dcterms:identifier "http://www.example.com/dataset/ZLOYOJ" ;
dcterms:issued "2024-07-01T11:11:11+00:00"^^xsd:dateTime ;
dcterms:license <https://creativecommons.org/licenses/by-sa/4.0/> ;
dcterms:modified "2024-06-04T13:36:10+00:00"^^xsd:dateTime ;
dcterms:publisher [ a foaf:Agent ;
    dcterms:identifier "https://ror.org/05wg1m734" ;
    foaf:homepage <https://www.xumc.nl/> ;
    foaf:mbox <mailto:data-access-committee@xumc.nl> ;
    foaf:name "Academic Medical Center" ] ;
dcterms:title "Questionnaire data of the Personalised RISk-based MAmmascreening Study (PRISMA)" ;
dcat:contactPoint [ a v:Kind ;
    v:fn "Data Access Committee of the x UMC" ;
    v:hasEmail <mailto:data-access-committee@xumc.nl> ] ;
dcat:theme <http://publications.europa.eu/resource/authority/data-theme/HEAL> .
`

func TestITBValidation(t *testing.T) {
	client := clients.NewItbEuropaClient()

	service, err := NewItbEuropaServices(client)
	if err != nil {
		t.Fatalf("Failed to create service: %v", err)
	}

	tests := []struct {
		title          string
		domain         string
		content        string
		contextSyntax  enums.ContextSyntax
		validationType enums.ValidationType
		expectPass     bool
	}{
		{"Valid DCAT-AP", "dcat-ap", validDcatAp, enums.Turtle, enums.V3Full1, true},
		{"Valid HealthRI", "healthri", validHealthri, enums.Turtle, enums.V200, true},
		{"Invalid DCAT-AP", "dcat-ap", invalidDcatAp, enums.Turtle, enums.V3Full1, false},
		{"Invalid HealthRI", "healthri", invalidHealthri, enums.Turtle, enums.V200, false},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {

			valid, err := service.ValidateContent(tt.domain, tt.content, tt.contextSyntax, tt.validationType)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if valid != tt.expectPass {
				t.Errorf("Expected %v, got %v", tt.expectPass, valid)
			}
		})
	}
}
