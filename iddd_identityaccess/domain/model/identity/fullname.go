package identity

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(aFirstName string, aLastName string) (_ *FullName, err error) {
	// defer ierrors.Wrap(&err, "fullname.NewFullName(%s, %s)", aFirstName, aLastName)
	fullName := new(FullName)
	errors := []FullNameError{}

	// set firstName
	if aFirstName == "" {
		errors = append(errors, &NameIsRequiredError{FullNameArgments{firstName: aFirstName, lastName: aLastName}})
		// return nil, &NameIsRequiredError{FullNameArgments{firstName: aFirstName, lastName: aLastName}}
	}
	if len(aFirstName) < 1 || 50 < len(aFirstName) {
		return nil, &Max50CharactersError{FullNameArgments{firstName: aFirstName, lastName: aLastName}}
	}
	if !regexp.MustCompile(`^[A-Z][a-z]*`).MatchString(aFirstName) {
		return nil, &RequiredAlphabetOnlyError{FullNameArgments{firstName: aFirstName, lastName: aLastName}}
	}
	fullName.firstName = aFirstName

	// set lastName
	if aLastName == "" {
		return nil, fmt.Errorf("Last name is required.")
	}
	if len(aLastName) < 1 || 50 < len(aLastName) {
		return nil, fmt.Errorf("Last name must be 50 characters or less.")
	}
	if !regexp.MustCompile(`^[a-zA-Z'][ a-zA-Z'-]*[a-zA-Z']?`).MatchString(aLastName) {
		return nil, fmt.Errorf("Last name must be at least one character in length.")
	}
	fullName.lastName = aLastName

	return fullName, nil
}

// TODO CopyFullName as shallow copy

func (fullName *FullName) AsFormattedName() string {
	return fmt.Sprintf("%s %s", fullName.firstName, fullName.lastName)
}

func (fullName *FullName) FirstName() string {
	return fullName.firstName
}

func (fullName *FullName) LastName() string {
	return fullName.lastName
}

func (fullName *FullName) WithChangedFirstName(aFirstName string) (_ *FullName, err error) {
	defer ierrors.Wrap(&err, "fullname.WithChangedFirstName(%s)", aFirstName)
	fullName, err = NewFullName(aFirstName, fullName.lastName)
	if err != nil {
		return nil, err
	}
	return fullName, nil
}

func (fullName *FullName) WithChangedLastName(aLastName string) (_ *FullName, err error) {
	defer ierrors.Wrap(&err, "fullname.WithChangedLastName(%s)", aLastName)
	fullName, err = NewFullName(fullName.firstName, aLastName)
	if err != nil {
		return nil, err
	}
	return fullName, nil
}

func (fullName *FullName) Equal(otherFullName *FullName) bool {
	isFirstNameEqual := reflect.DeepEqual(fullName.firstName, otherFullName.firstName)
	isLastNameEqual := reflect.DeepEqual(fullName.lastName, otherFullName.lastName)

	return isFirstNameEqual && isLastNameEqual
}

func (fullName *FullName) String() string {
	return fmt.Sprintf("FullName [firstName=" + fullName.firstName + ", lastName=" + fullName.lastName + "]")
}

type FullNameArgments struct {
	firstName string
	lastName  string
}

type FullNameError interface {
	getArguments() FullNameArgments
	getError() error
}

// type FullNameParseError struct {
// }

type NameIsRequiredError struct {
	argments FullNameArgments
	// err      error
}

type Max50CharactersError struct {
	argments FullNameArgments
	// err      error
}

type RequiredAlphabetOnlyError struct {
	argments FullNameArgments
	// err      error
}

// func (fullNameParseError *FullNameParseError) getArguments() FullNameArgments {
// 	return fullNameParseError.argments
// }

// func (fullNameParseError *FullNameParseError) getError() error {
// 	return fullNameParseError.err
// }

// func (fullNameParseError *FullNameParseError) Error() string {
// 	return fullNameParseError.err.Error()
// }

func (nameIsRequiredError *NameIsRequiredError) getArguments() FullNameArgments {
	return nameIsRequiredError.argments
}

func (nameIsRequiredError *NameIsRequiredError) getError() error {
	return nameIsRequiredError
}

func (nameIsRequiredError *NameIsRequiredError) Error() string {
	return "name is requred"
}

func (max50CharactersError *Max50CharactersError) getArguments() FullNameArgments {
	return max50CharactersError.argments
}

func (max50CharactersError *Max50CharactersError) getError() error {
	return max50CharactersError
}

func (max50CharactersError *Max50CharactersError) Error() string {
	return "max 50 character"
}

func (requiredAlphabetOnlyError *RequiredAlphabetOnlyError) getArguments() FullNameArgments {
	return requiredAlphabetOnlyError.argments
}

func (requiredAlphabetOnlyError *RequiredAlphabetOnlyError) getError() error {
	return requiredAlphabetOnlyError
}

func (requiredAlphabetOnlyError *RequiredAlphabetOnlyError) Error() string {
	return "alphabet only"
}
