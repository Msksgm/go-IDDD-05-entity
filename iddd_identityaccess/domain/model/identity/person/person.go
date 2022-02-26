package person

import (
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/contactinformation"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/emailaddress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/fullname"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/tenantid"
)

type Person struct {
	tenantId           tenantid.TenantId
	name               fullname.FullName
	contactInformation contactinformation.ContactInformation
}

func NewPerson(aTenantId tenantid.TenantId, aName fullname.FullName, aContactInformation contactinformation.ContactInformation) (*Person, error) {
	person := new(Person)

	person.tenantId = aTenantId
	person.name = aName
	person.contactInformation = aContactInformation
	return person, nil
}

func (person *Person) ChangeContactInformation(aContactInformation contactinformation.ContactInformation) error {
	person.contactInformation = aContactInformation
	return nil
}

func (person *Person) ChangeName(aName fullname.FullName) error {
	person.name = aName
	return nil
}

func (person *Person) ContactInformation() contactinformation.ContactInformation {
	return person.contactInformation
}

func (person *Person) EmailAddress() emailaddress.EmailAddress {
	return *person.contactInformation.EmailAddress()
}
