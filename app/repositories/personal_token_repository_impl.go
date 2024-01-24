package repositories

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/Zainal21/go-bone/app/dtos"
	"github.com/Zainal21/go-bone/app/entity"
	"github.com/Zainal21/go-bone/app/utils/query"
	"github.com/Zainal21/go-bone/app/utils/sanctum"
	"github.com/Zainal21/go-bone/pkg/database/mysql"
)

type personalTokenImpl struct {
	db             mysql.Adapter
	Token          sanctum.TokenI
	userRepository UserRepository
}

// DeleteByUserId implements PersonalTokenRepository.
func (p *personalTokenImpl) DeleteByUserId(ctx context.Context, user_id string) error {
	_query := `DELETE FROM personal_access_tokens WHERE tokenable_id = ?`
	_, err := p.db.Exec(ctx, _query, user_id)

	if err != nil {
		return err
	}

	return nil
}

// Delete implements PersonalTokenRepository.
func (p *personalTokenImpl) Delete(ctx context.Context, token string) error {
	idStr, hash, err := p.Token.Split(token)

	if err != nil {
		return err
	}

	_query := `DELETE FROM personal_access_tokens WHERE token = ? AND id = ?`
	_, err = p.db.Exec(ctx, _query, hash, idStr)

	if err != nil {
		return err
	}

	return nil

}

// Verify implements PersonalTokenRepository.
func (p *personalTokenImpl) Verify(ctx context.Context, token string) (*entity.User, error) {
	idStr, hash, err := p.Token.Split(token)
	if err != nil {
		return nil, err
	}

	_query := query.SelectQuery(
		"personal_access_tokens",
		[]string{
			"id",
			"name",
			"token",
			"expires_at",
			"tokenable_id",
		},
		"id = ? AND AND id = ? ",
		1,
		0,
	)

	var personalAccessToken entity.PersonalAccessToken

	row := p.db.QueryRowX(ctx, _query, hash, idStr)

	var expirationTimeStr string

	if err = row.Scan(
		&personalAccessToken.Id,
		&personalAccessToken.Name,
		&personalAccessToken.Token,
		&expirationTimeStr,
		&personalAccessToken.UserId,
	); err != nil {
		return nil, err
	}

	expirationTime, err := time.Parse("2006-01-02T15:04:05Z", expirationTimeStr)
	if err != nil {
		return nil, err
	}

	personalAccessToken.ExpirationAt = &expirationTime

	if personalAccessToken.ExpirationAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	user, err := p.userRepository.FindById(ctx, personalAccessToken.UserId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetPermissionByUserId implements PersonalTokenRepository.
func (p *personalTokenImpl) Create(ctx context.Context, personalTokenDto *dtos.PersonalAccessTokenDto) (string, error) {
	expirationAt := time.Now().Add(time.Hour * 24 * 7)

	tokenItem, err := p.Token.Create()
	if err != nil {
		return "", err
	}

	_query := `INSERT INTO personal_access_tokens
			(
				tokenable_type, 
				tokenable_id, 
				name,
				token, 
				abilities, 
				last_used_at, 
				expires_at, 
				created_at, 
				updated_at
			)
			VALUES(
				'App\\Models\\User', 
				?, 
				'kejaksaan-agung', 
				? , 
				'["*"]', 
				NULL, 
				?, 
				?, 
				?
			);
		`
	_, err = p.db.Exec(ctx, _query, personalTokenDto.TokenableId, tokenItem.Hash, expirationAt, time.Now(), time.Now())
	if err != nil {
		return "", err
	}

	var generatedID int

	result := p.db.QueryRowX(ctx, "SELECT LAST_INSERT_ID()")

	err = result.Scan(&generatedID)
	if err != nil {
		return "", err
	}

	personalAccessToken := &entity.PersonalAccessToken{
		Id: generatedID,
	}

	return tokenItem.GetPlainText(strconv.Itoa(personalAccessToken.Id)), nil
}

func NewPersonalToken(db mysql.Adapter, Token sanctum.TokenI, userRepo UserRepository) PersonalTokenRepository {
	return &personalTokenImpl{
		db:             db,
		Token:          Token,
		userRepository: userRepo,
	}
}
