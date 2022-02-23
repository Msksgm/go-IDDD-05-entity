package identity

// func TestNewUser(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()

// 		tenantId, err := NewTenantId(uu)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		userName := "userName"
// 		password := "qwerty!ASDFG#"
// 		bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		got, err := NewUser(*tenantId, userName, password)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		want := &User{tenantId: TenantId{id: uu}, userName: userName, password: string(bcryptedPassword)}

// 		opts := cmp.Options{
// 			cmp.AllowUnexported(User{}, TenantId{}),
// 			cmpopts.IgnoreFields(User{}, "password"),
// 		}
// 		if diff := cmp.Diff(want, got, opts); diff != "" {
// 			t.Errorf("mismatch (-want, +got):\n%s", diff)
// 		}
// 		if err := bcrypt.CompareHashAndPassword([]byte(want.password), []byte(password)); err != nil {
// 			t.Error(err)
// 		}
// 	})
// 	t.Run("fail username is required.", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()

// 		tenantId, err := NewTenantId(uu)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		userName := ""
// 		password := "qwerty!ASDFG#"

// 		_, err = NewUser(*tenantId, userName, password)
// 		want := fmt.Sprintf("user.setUserName(%s): The username is required.", userName)
// 		if got := err.Error(); got != want {
// 			t.Errorf("got %s, want %s", got, want)
// 		}
// 	})
// 	t.Run("fail username is lower than 3 characters.", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()

// 		tenantId, err := NewTenantId(uu)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		userName := "na"
// 		password := "qwerty!ASDFG#"

// 		_, err = NewUser(*tenantId, userName, password)
// 		want := fmt.Sprintf("user.setUserName(%s): The username must be 3 to 250 characters.", userName)
// 		if got := err.Error(); got != want {
// 			t.Errorf("got %s, want %s", got, want)
// 		}
// 	})
// 	t.Run("fail username is over than 250 characters.", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()

// 		tenantId, err := NewTenantId(uu)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		userName := RandString(251)
// 		password := "qwerty!ASDFG#"

// 		_, err = NewUser(*tenantId, userName, password)
// 		want := fmt.Sprintf("user.setUserName(%s): The username must be 3 to 250 characters.", userName)
// 		if got := err.Error(); got != want {
// 			t.Errorf("got %s, want %s", got, want)
// 		}
// 	})
// }

// func TestAssertPasswordNotSame(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()
// 		tenantId := TenantId{id: uu}

// 		userName := "user"
// 		password := "qwerty!ASDFG#"
// 		user := &User{tenantId: tenantId, userName: userName, password: password}
// 		changedPassword := "ASDFG#qwerty!"

// 		if err = user.assertPasswordNotSame(password, changedPassword); err != nil {
// 			t.Error(err)
// 		}
// 	})
// 	t.Run("fail", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()
// 		tenantId := TenantId{id: uu}

// 		userName := "user"
// 		password := "qwerty!ASDFG#"
// 		user := &User{tenantId: tenantId, userName: userName, password: password}
// 		changedPassword := "qwerty!ASDFG#"

// 		err = user.assertPasswordNotSame(password, changedPassword)
// 		want := fmt.Sprintf("user.assertPasswordNotSame(%s, %s): The password is unchanged", password, changedPassword)
// 		if got := err.Error(); got != want {
// 			t.Errorf("got %s, want %s", got, want)
// 		}
// 	})
// }

// func TestAssertUsernamePasswordNotSame(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()
// 		tenantId := TenantId{id: uu}

// 		userName := "user"
// 		password := "qwerty!ASDFG#"
// 		user := &User{tenantId: tenantId, userName: userName, password: password}
// 		changedPassword := "qwerty!ASDFG#"

// 		if err := user.assertUsernamePasswordNotSame(changedPassword); err != nil {
// 			t.Error(err)
// 		}
// 	})
// 	t.Run("fail", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()
// 		tenantId := TenantId{id: uu}

// 		userName := "user"
// 		password := "qwerty!ASDFG#"
// 		user := &User{tenantId: tenantId, userName: userName, password: password}
// 		changedPassword := "user"

// 		err = user.assertUsernamePasswordNotSame(changedPassword)
// 		want := fmt.Sprintf("user.assertUsernamePasswordNotSame(%s): The username and password must not be the same.", changedPassword)
// 		if got := err.Error(); got != want {
// 			t.Errorf("got %s, want %s", got, want)
// 		}
// 	})
// }

// func TestAssertPasswordNotWeak(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()
// 		tenantId := TenantId{id: uu}

// 		userName := "user"
// 		password := "qwerty!ASDFG#"
// 		user := &User{tenantId: tenantId, userName: userName, password: password}
// 		changedPassword := "qwerty!ASDFG"
// 		if err := user.assertPasswordNotWeak(changedPassword); err != nil {
// 			t.Error(err)
// 		}
// 	})
// 	t.Run("fail password empty", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()
// 		tenantId := TenantId{id: uu}

// 		userName := "user"
// 		password := "qwerty!ASDFG#"
// 		user := &User{tenantId: tenantId, userName: userName, password: password}
// 		changedPassword := ""
// 		err = user.assertPasswordNotWeak(changedPassword)
// 		want := fmt.Sprintf("user.assertPasswordNotWeak(%s): The password must not be empty", changedPassword)
// 		if got := err.Error(); got != want {
// 			t.Errorf("got %s, want %s", got, want)
// 		}
// 	})
// 	t.Run("fail password is weak", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()
// 		tenantId := TenantId{id: uu}

// 		userName := "user"
// 		password := "qwerty!ASDFG#"
// 		user := &User{tenantId: tenantId, userName: userName, password: password}
// 		changedPassword := "123456"
// 		err = user.assertPasswordNotWeak(changedPassword)
// 		want := fmt.Sprintf("user.assertPasswordNotWeak(%s): The password must be stronger.", changedPassword)
// 		if got := err.Error(); got != want {
// 			t.Errorf("got %s, want %s", got, want)
// 		}
// 	})
// }

// func TestUserEquals(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		u, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu := u.String()

// 		tenantId, err := NewTenantId(uu)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		userName := "userName"
// 		password := "qwerty!ASDFG#"
// 		bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		user, err := NewUser(*tenantId, userName, password)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		other := &User{tenantId: TenantId{id: uu}, userName: userName, password: string(bcryptedPassword)}

// 		if !user.Equals(*other) {
// 			t.Errorf("user: %v must be equal to other :%v", user, other)
// 		}
// 	})
// 	t.Run("fail tenantId is not equal", func(t *testing.T) {
// 		u1, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu1 := u1.String()
// 		tenantId1 := &TenantId{id: uu1}

// 		userName := "userName"
// 		password := "qwerty!ASDFG#"
// 		bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		user, err := NewUser(*tenantId1, userName, password)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		u2, err := uuid.NewRandom()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		uu2 := u2.String()
// 		tenantId2 := &TenantId{id: uu2}

// 		other := &User{tenantId: *tenantId2, userName: userName, password: string(bcryptedPassword)}

// 		if user.Equals(*other) {
// 			t.Errorf("user: %v must be equal to other :%v", user, other)
// 		}
// 	})
// }
