package store
 
import (
  "crypto/rand"
  "time"
  "errors"

  // "github.com/go-pg/pg/v10"
	"github.com/rs/zerolog/log"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  ID             int
  Username       string `binding:"required,min=5,max=30"`
  Password       string `pg:"-" binding:"required,min=7,max=32"`
  HashedPassword []byte `json:"-"`
  Salt           []byte `json:"-"`
  CreatedAt      time.Time
  ModifiedAt     time.Time
  Images         []*Image `json:"-" pg:"fk:user_id,rel:has-many,on_delete:CASCADE"`
}
 
func AddUser(user *User) error {
  salt, err := GenerateSalt()
  if err != nil {
    return err
  }
  toHash := append([]byte(user.Password), salt...)
  hashedPassword, err := bcrypt.GenerateFromPassword(toHash, bcrypt.DefaultCost)
  if err != nil {
    return err
  }
 
  user.Salt = salt
  user.HashedPassword = hashedPassword
 
  _, err = db.Model(user).Returning("*").Insert()
  if err != nil {
    return err
  }
  return err
}
 
func Authenticate(username, password string) (*User, error) {
  user := new(User)
  if err := db.Model(user).Where(
    "username = ?", username).Select(); err != nil {
    return nil, err
  }
  salted := append([]byte(password), user.Salt...)
  if err := bcrypt.CompareHashAndPassword (user.HashedPassword, salted); err != nil {
    return nil, errors.New("Password not valid.")
  }
  return user, nil
}

func GenerateSalt() ([]byte, error) {
  salt := make([]byte, 16)
  if _, err := rand.Read(salt); err != nil {
    log.Error().Err(err).Msg("Unable to create salt")
    return nil, err
  }
  return salt, nil
}