package handler

// func (h *Manager) CreateBook(c echo.Context) error {
// 	input := struct {
// 		Title     string `json:"title"     validate:"required,min=5"`
// 		Login    string `json:"login"    validate:"required,min=5"`
// 		Password string `json:"password" validate:"required,min=5"`
// 	}{}

// 	if err := c.Bind(&input); err != nil {
// 		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(CreateUser): bad request"})
// 	}

// 	if err := c.Validate(input); err != nil {
// 		return c.JSON(http.StatusBadRequest, envelope{
// 			"error":         "handler(CreateUser): validation failed",
// 			"errorValidate": err.Error(),
// 		})
// 	}

// 	user := model.User{
// 		Name:     input.Name,
// 		Login:    input.Login,
// 		Password: input.Password,
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Я вот не знаю где и как лучше обрабатывать ошибки и как их лучше стандартизировать, не довело видеть хороших примеров, у меня не правильно мне кажеться сделано, хотелось бы знать как лучше.

// 	err := h.s.User.Create(ctx, user)
// 	switch {
// 	case errors.Is(err, gorm.ErrDuplicatedKey):
// 		return c.JSON(http.StatusUnprocessableEntity, envelope{"error": "handler(CreateUser): user already exist || " + err.Error()})
// 	case err != nil:
// 		return c.JSON(http.StatusInternalServerError, envelope{"error": "handler(CreateUser): " + err.Error()})
// 	default:
// 		return c.JSON(http.StatusCreated, envelope{"info": "user created"})
// 	}
// }

// func (h *Manager) GetUser(c echo.Context) error {
// 	input := struct {
// 		ID int `param:"id" validate:"required,min=1"`
// 	}{}

// 	if err := c.Bind(&input); err != nil {
// 		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(GetUser): bad request"})
// 	}

// 	if err := c.Validate(input); err != nil {
// 		return c.JSON(http.StatusBadRequest, envelope{
// 			"error":         "handler(GetUser): validation failed",
// 			"errorValidate": err.Error(),
// 		})
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	user, err := h.s.User.Get(ctx, input.ID)
// 	switch {
// 	case errors.Is(err, gorm.ErrRecordNotFound):
// 		return c.JSON(http.StatusNotFound, envelope{"info": "handler(GetUser): user is not exist || " + err.Error()})
// 	case err != nil:
// 		return c.JSON(http.StatusInternalServerError, envelope{"error": "handler(GetUser): " + err.Error()})
// 	default:
// 		return c.JSON(http.StatusFound, user)
// 	}
// }
