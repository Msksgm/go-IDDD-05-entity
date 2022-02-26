package user

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/utils"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/tenantid"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	userName = "userName"
	password = "qwerty!ASDFG#"
)

var tenantId *tenantid.TenantId

func init() {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	uu := u.String()
	tenantId, err = tenantid.NewTenantId(uu)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	argumentLengthError   *ierrors.ArgumentLengthError
	argumentNotEmptyError *ierrors.ArgumentNotEmptyError
)

func TestNewUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			t.Fatal(err)
		}

		got, err := NewUser(*tenantId, userName, password)
		if err != nil {
			t.Fatal(err)
		}

		want := &User{tenantId: *tenantId, userName: userName, password: string(bcryptedPassword)}

		opts := cmp.Options{
			cmp.AllowUnexported(User{}, tenantid.TenantId{}),
			cmpopts.IgnoreFields(User{}, "password"),
		}
		if diff := cmp.Diff(want, got, opts); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
		if err := bcrypt.CompareHashAndPassword([]byte(want.password), []byte(password)); err != nil {
			t.Error(err)
		}
	})
	t.Run("fail username is required.", func(t *testing.T) {
		_, err := NewUser(*tenantId, "", password)
		if !errors.As(err, &argumentNotEmptyError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentNotEmptyError))
		}
	})
	t.Run("fail username is lower than 3 characters.", func(t *testing.T) {
		_, err := NewUser(*tenantId, "na", password)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail username is over than 250 characters.", func(t *testing.T) {
		_, err := NewUser(*tenantId, utils.RandString(251), password)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
}

func TestAssertPasswordNotSame(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user := &User{tenantId: *tenantId, userName: userName, password: password}
		changedPassword := "ASDFG#qwerty!"

		if err := user.assertPasswordNotSame(password, changedPassword); err != nil {
			t.Error(err)
		}
	})
	t.Run("fail", func(t *testing.T) {
		user := &User{tenantId: *tenantId, userName: userName, password: password}
		changedPassword := "qwerty!ASDFG#"

		err := user.assertPasswordNotSame(password, changedPassword)
		want := fmt.Sprintf("user.assertPasswordNotSame(%s, %s): The password is unchanged", password, changedPassword)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestAssertUsernamePasswordNotSame(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user := &User{tenantId: *tenantId, userName: userName, password: password}
		changedPassword := "qwerty!ASDFG#"

		if err := user.assertUsernamePasswordNotSame(changedPassword); err != nil {
			t.Error(err)
		}
	})
	t.Run("fail", func(t *testing.T) {
		user := &User{tenantId: *tenantId, userName: userName, password: password}
		changedPassword := "userName"

		err := user.assertUsernamePasswordNotSame(changedPassword)
		want := fmt.Sprintf("user.assertUsernamePasswordNotSame(%s): The username and password must not be the same.", changedPassword)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestAssertPasswordNotWeak(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user := &User{tenantId: *tenantId, userName: userName, password: password}
		changedPassword := "qwerty!ASDFG"
		if err := user.assertPasswordNotWeak(changedPassword); err != nil {
			t.Error(err)
		}
	})
	t.Run("fail password empty", func(t *testing.T) {
		user := &User{tenantId: *tenantId, userName: userName, password: password}
		changedPassword := ""
		err := user.assertPasswordNotWeak(changedPassword)
		want := fmt.Sprintf("user.assertPasswordNotWeak(%s): The password must not be empty", changedPassword)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail password is weak", func(t *testing.T) {
		user := &User{tenantId: *tenantId, userName: userName, password: password}
		changedPassword := "123456"
		err := user.assertPasswordNotWeak(changedPassword)
		want := fmt.Sprintf("user.assertPasswordNotWeak(%s): The password must be stronger.", changedPassword)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestUserEquals(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			t.Fatal(err)
		}

		user, err := NewUser(*tenantId, userName, password)
		if err != nil {
			t.Fatal(err)
		}

		other := &User{tenantId: *tenantId, userName: userName, password: string(bcryptedPassword)}

		if !user.Equals(*other) {
			t.Errorf("user: %v must be equal to other :%v", user, other)
		}
	})
	t.Run("fail tenantId is not equal", func(t *testing.T) {
		bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			t.Fatal(err)
		}

		user, err := NewUser(*tenantId, userName, password)
		if err != nil {
			t.Fatal(err)
		}

		u2, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu2 := u2.String()
		tenantId2, err := tenantid.NewTenantId(uu2)
		if err != nil {
			t.Fatal(err)
		}

		other := &User{tenantId: *tenantId2, userName: userName, password: string(bcryptedPassword)}

		if user.Equals(*other) {
			t.Errorf("user: %v must be equal to other :%v", user, other)
		}
	})
}
