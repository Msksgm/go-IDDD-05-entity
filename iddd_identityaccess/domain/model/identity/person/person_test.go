package person

import (
	"log"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/contactinformation"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/emailaddress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/fullname"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/postaladdress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/telephone"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/tenantid"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

const (
	address                                                     = "sample@mail.com"
	streetAddress, city, stateProvince, postalCode, countryCode = "streetAddress", "city", "stateProvince", "postalCode", "00"
	primaryNumber                                               = "090-1234-5678"
	secondaryNumber                                             = "090-5678-1234"
	firstName                                                   = "FirstName"
	lastName                                                    = "lastName"
)

var (
	contactInformation *contactinformation.ContactInformation
	fullName           *fullname.FullName
	tenantId           *tenantid.TenantId
)

func init() {
	emailAddress, err := emailaddress.NewEmailAddress(address)
	if err != nil {
		log.Fatal(err)
	}
	postalAddress, err := postaladdress.NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	primaryTelephone, err := telephone.NewTelephone(primaryNumber)
	if err != nil {
		log.Fatal(err)
	}
	secondaryTelephone, err := telephone.NewTelephone(secondaryNumber)
	if err != nil {
		log.Fatal(err)
	}
	contactInformation, err = contactinformation.NewContactInformation(*emailAddress, *postalAddress, *primaryTelephone, *secondaryTelephone)
	if err != nil {
		log.Fatal(err)
	}

	fullName, err = fullname.NewFullName(firstName, lastName)
	if err != nil {
		log.Fatal(err)
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	uuString := uuid.String()
	tenantId, err = tenantid.NewTenantId(uuString)
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewPerson(t *testing.T) {
	got, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	want := &Person{tenantId: *tenantId, name: *fullName, contactInformation: *contactInformation}

	allowUnexported := cmp.AllowUnexported(Person{}, tenantid.TenantId{}, fullname.FullName{}, contactinformation.ContactInformation{}, emailaddress.EmailAddress{}, postaladdress.PostalAddress{}, telephone.Telephone{})
	if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestChangeContactInformation(t *testing.T) {
	person, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	changedAddressString := "changed@email.com"
	changedAddress, err := emailaddress.NewEmailAddress(changedAddressString)
	if err != nil {
		t.Fatal(err)
	}
	changedContactInfomation, err := contactInformation.ChangeEmailAddress(*changedAddress)
	if err != nil {
		t.Fatal(err)
	}
	want, err := NewPerson(*tenantId, *fullName, *changedContactInfomation)
	if err != nil {
		t.Fatal(err)
	}
	if err := person.ChangeContactInformation(*changedContactInfomation); err != nil {
		t.Fatal(err)
	}

	allowUnexported := cmp.AllowUnexported(Person{}, tenantid.TenantId{}, fullname.FullName{}, contactinformation.ContactInformation{}, emailaddress.EmailAddress{}, postaladdress.PostalAddress{}, telephone.Telephone{})
	if diff := cmp.Diff(want, person, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestChangeName(t *testing.T) {
	person, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	changedFirstName := "ChangedFirstName"
	changedFullName, err := fullName.WithChangedFirstName(changedFirstName)
	if err != nil {
		t.Fatal(err)
	}

	if err := person.ChangeName(*changedFullName); err != nil {
		t.Fatal(err)
	}

	want, err := NewPerson(*tenantId, *changedFullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	allowUnexported := cmp.AllowUnexported(Person{}, tenantid.TenantId{}, fullname.FullName{}, contactinformation.ContactInformation{}, emailaddress.EmailAddress{}, postaladdress.PostalAddress{}, telephone.Telephone{})
	if diff := cmp.Diff(want, person, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestContactInfomation(t *testing.T) {
	person, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}
	want := *contactInformation
	got := person.ContactInformation()
	allowUnexported := cmp.AllowUnexported(Person{}, tenantid.TenantId{}, fullname.FullName{}, contactinformation.ContactInformation{}, emailaddress.EmailAddress{}, postaladdress.PostalAddress{}, telephone.Telephone{})
	if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestEmailAddress(t *testing.T) {
	person, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}
	want := *contactInformation.EmailAddress()
	got := person.EmailAddress()
	allowUnexported := cmp.AllowUnexported(Person{}, tenantid.TenantId{}, fullname.FullName{}, contactinformation.ContactInformation{}, emailaddress.EmailAddress{}, postaladdress.PostalAddress{}, telephone.Telephone{})
	if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
