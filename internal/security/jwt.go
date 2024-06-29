package security

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Secret key used to sign the JWTs
var jwtKey = []byte("your_secret_key")

// Claims represents the JWT claims
type Claims struct {
	jwt.StandardClaims
}

func Init() {
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// ValidateToken parses and validates a JWT token
func validateToken(tokenString string) (bool, error) {
	jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY")) // Использование значения из переменной окружения

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}

// UnaryInterceptor intercepts and validates JWT token
func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "retrieving metadata failed")
		}

		authHeader, ok := md["authorization"]
		if !ok || len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "authorization token is required")
		}

		token := strings.TrimPrefix(authHeader[0], "Bearer ")
		if valid, err := validateToken(token); !valid || err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token: "+err.Error())
		}

		return handler(ctx, req)
	}
}
