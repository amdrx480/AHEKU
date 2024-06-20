package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// var rolePermissions = map[string][]string{
// 	"superadmin": {"read", "write", "delete"},
// 	"admin":      {"read", "write"},
// }

// var rolePermissions = map[string][]string{
// 	"superadmin": {"read", "write", "delete"}, // ID 1 untuk Super Admin
// 	"admin":      {"read", "write"},           // ID 2 untuk Admin
// }

// JwtCustomClaims defines custom claims for JWT
type JwtCustomClaims struct {
	ID int `json:"id"`
	// Role     admin.Role
	// RoleName string `json:"role_name"`
	jwt.RegisteredClaims
}

// JWTConfig holds the JWT configuration
type JWTConfig struct {
	SecretKey       string
	ExpiresDuration int
}

// Init initializes the JWT configuration for Echo
func (jwtConfig *JWTConfig) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(jwtConfig.SecretKey),
		// sus
		// TokenLookup: "header:Authorization",
		// ContextKey:  "user",
		// AuthScheme:  "Bearer",
	}
}

// GenerateToken generates a new JWT token for a user
// func (jwtConfig *JWTConfig) GenerateToken(userID int, role string) (string, error) {
func (jwtConfig *JWTConfig) GenerateToken(userID int, RoleName string) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(jwtConfig.ExpiresDuration)))

	claims := &JwtCustomClaims{
		ID: userID,
		// RoleName: RoleName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(jwtConfig.SecretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUser extracts the user data from context
func GetUser(c echo.Context) (*JwtCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	claims := user.Claims.(*JwtCustomClaims)
	if claims == nil {
		return nil, errors.New("invalid token claims")
	}

	// Accessing RoleID
	// roleID := claims.RoleID
	// fmt.Println("RoleID:", roleID)

	return claims, nil
}

// VerifyToken middleware to verify the token and set user data in context
func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userData, err := GetUser(c)

		isInvalid := userData == nil || err != nil

		if isInvalid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		return next(c)
	}
}

// RBAC middleware to check user's role and permissions
// func RBAC(requiredPermission string) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			userData, err := GetUser(c)
// 			if err != nil {
// 				return c.JSON(http.StatusUnauthorized, map[string]string{
// 					"message": "invalid token",
// 				})
// 			}

// 			permissions, roleExists := rolePermissions[userData.RoleName]
// 			if !roleExists {
// 				return c.JSON(http.StatusForbidden, map[string]string{
// 					"message": "role not found",
// 				})
// 			}

// 			hasPermission := false
// 			for _, permission := range permissions {
// 				if permission == requiredPermission {
// 					hasPermission = true
// 					break
// 				}
// 			}

// 			if !hasPermission {
// 				return c.JSON(http.StatusForbidden, map[string]interface{}{
// 					"message": "forbidden",
// 					"details": map[string]interface{}{
// 						"user_id":             userData.ID,
// 						"role_name":           userData.RoleName,
// 						"required_permission": requiredPermission,
// 					},
// 				})
// 			}

// 			return next(c)
// 		}
// 	}
// }

// RBAC middleware to check user's role and permissions
// func RBAC(requiredPermission string) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			userData, err := GetUser(c)
// 			if err != nil {
// 				return c.JSON(http.StatusUnauthorized, map[string]string{
// 					"message": "invalid token",
// 				})
// 			}

// 			permissions, roleExists := rolePermissions[userData.RoleName]
// 			if !roleExists {
// 				return c.JSON(http.StatusForbidden, map[string]string{
// 					"message": "role not found",
// 				})
// 			}

// 			requiredPermissionStr := strconv.Itoa(requiredPermission)
// 			hasPermission := false
// 			for _, permission := range permissions {
// 				if permission == requiredPermissionStr {
// 					hasPermission = true
// 					break
// 				}
// 			}
// 			if !hasPermission {
// 				return c.JSON(http.StatusForbidden, map[string]interface{}{
// 					"message": "forbidden",
// 					"details": map[string]interface{}{
// 						"user_id":             userData.ID,
// 						"role_id":             userData.RoleName,
// 						"required_permission": requiredPermission,
// 					},
// 				})
// 			}

// 			return next(c)
// 		}
// 	}
// }

// func RBAC(requiredPermission string) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			userData, err := GetUser(c)
// 			if err != nil {
// 				return c.JSON(http.StatusUnauthorized, map[string]string{
// 					"message": "invalid token",
// 				})
// 			}

// 			permissions, roleExists := rolePermissions[userData.RoleID]
// 			if !roleExists {
// 				return c.JSON(http.StatusForbidden, map[string]string{
// 					"message": "role not found",
// 				})
// 			}

// 			hasPermission := false
// 			for _, permission := range permissions {
// 				if permission == requiredPermission {
// 					hasPermission = true
// 					break
// 				}
// 			}

// 			if !hasPermission {
// 				return c.JSON(http.StatusForbidden, map[string]string{
// 					"message": "forbidden",
// 				})
// 			}

// 			return next(c)
// 		}
// 	}
// }
