package identity

import (
	"github.com/google/uuid"
)

type TenantId struct {
	id string
}

func NewTenantId(uu string) (_ *TenantId, err TenantIdError) {
	// defer ierrors.Wrap(&err, "tenantid.NewTenantId(%s)", uu)
	tenantId := new(TenantId)

	// setId
	if _, err := uuid.Parse(uu); err != nil {
		return nil, &TenantIdParseError{TenantIdArgments{uuid: uu}, err}
	}
	tenantId.id = uu

	return tenantId, nil
}

type TenantIdArgments struct {
	uuid string
}

type TenantIdError interface {
	getArguments() TenantIdArgments
	getError() error
}

type TenantIdParseError struct {
	argments TenantIdArgments
	err      error
}

func (tenantIdParseError *TenantIdParseError) getArguments() TenantIdArgments {
	return tenantIdParseError.argments
}

func (tenantIdParseError *TenantIdParseError) getError() error {
	return tenantIdParseError.err
}

func (tenantIdParseError *TenantIdParseError) Error() string {
	return tenantIdParseError.err.Error()
}
