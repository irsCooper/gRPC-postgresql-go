package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	// "github.com/go-playground/locales/sl"
	"github.com/irsCooper/gRPC-postgresql-go/sso/internal/domain/models"
	"github.com/irsCooper/gRPC-postgresql-go/sso/internal/storage"
	"github.com/irsCooper/gRPC-postgresql-go/sso/lib/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	appProvider AppProvider
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentails")
	ErrInvalidAppID = errors.New("invalid app id")
	ErrUserExists = errors.New("user already exists")
)

// New returns a new instance of the Auth service.
func New(log *slog.Logger, userSaver UserSaver, userProvider UserProvider, appProvider AppProvider, tokenTTL time.Duration) *Auth {
	return &Auth{
		log:         log,
		usrSaver:    userSaver,
		usrProvider: userProvider,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
	}
}


// Login chech if user given credentials exist in the system.
// 
// If user exist, but password is invalid, returns error.
// If user doesn't exist, returns error.
func (a *Auth) Login(ctx context.Context, email string, password string, appID int) (string, error) {
	const op = "auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("login user")

	user, err := a.usrProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			// a.log.Warn("user not found", sl.Err(err))
			a.log.Warn("user not found", err)
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		// a.log.Error("failed to get user", sl.Err(err))
		a.log.Error("failed to get user", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}


	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		// a.log.Info("invalid credentials", sl.Err(err))
		a.log.Info("invalid credentials", err)
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	app, err := a.appProvider.App(ctx, appID)
	if err != nil {
		return "", fmt.Errorf("%s %w", op, err)
	}

	log.Info("user logged in successfully")

	token, err := jwt.NewToken(user, app, a.tokenTTL)
	if err != nil {
		// a.log.Error("failed to generate token", sl.Err(err))
		a.log.Error("failed to generate token", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}


// RegisterNewUser registers new user in the system and returns user ID.
// 
// If user with given username already exist, returns error.
func (a Auth) RegisterNewUser(ctx context.Context, email string, pass string) (int64, error) {
	const op = "auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		// log.Error("failed to generate password hash", sl.Err(err))
		log.Error("failed to generate password hash", err)

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.usrSaver.SaveUser(ctx, email, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserExist) {
			// log.Warn("user already exists", sl.Err(err))
			log.Warn("user already exists", err)
			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}

		// log.Error("failed to save user", sl.Err(err))
		log.Error("failed to save user", err)

		return 0, fmt.Errorf("%s: %w", op, err)
	}	

	log.Info("user is registrated")

	return id, nil
}

// IsAdmiv check if user is admin.
func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "auth.IsAdmin"

	log := a.log.With(
		slog.String("op", op),
		slog.Int64("userID", userID),
	)

	log.Info("checking role admin for user")

	isAdmin, err := a.usrProvider.IsAdmin(ctx, userID)
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			// log.Warn("user not found", sl.Err(err))
			log.Warn("user not found", err)
			return false, fmt.Errorf("%s: %w", op, ErrInvalidAppID)
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("checked role admin for user", slog.Bool("isAdmin", isAdmin))

	return isAdmin, nil
}



