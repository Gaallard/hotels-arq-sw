package middleware

import (
	"fmt"
	"net/http"
	"strings"
	client "users-api/users-api/client"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// Extraer el token eliminando el prefijo "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Imprimir el token recibido para depuración
		fmt.Printf("Received Token: %s\n", tokenString)

		// Decodificar el token JWT sin verificar la firma
		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
		if err == nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				// Imprimir todos los reclamos para depuración
				fmt.Printf("Unverified Claims: %v\n", claims)
			}
		} else {
			fmt.Printf("Error parsing unverified token: %v\n", err)
		}

		// Decodificar el token JWT verificando la firma
		token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			fmt.Printf("Error parsing token: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Imprimir todos los reclamos para depuración
			fmt.Printf("Claims: %v\n", claims)

			idUser, _ := claims["id_user"].(float64)
			user, err := client.GetUserById(int(idUser))

			// Verificar si se extrajeron correctamente los reclamos
			if user.Id_user == 0 || err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}
		} else {
			fmt.Printf("Token is invalid or claims are not of type jwt.MapClaims\n")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}

func GetUserIdByToken(context *gin.Context) (int, error) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("no token provided")
	}

	// Extraer el token eliminando el prefijo "Bearer "
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return 0, fmt.Errorf("invalid token format")
	}

	// Decodificar el token JWT verificando la firma
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idUser, _ := claims["id_user"].(float64)
		return int(idUser), nil
	}

	return 0, fmt.Errorf("invalid token claims")
}

func AuthMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// Extraer el token eliminando el prefijo "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Imprimir el token recibido para depuración
		fmt.Printf("Received Token: %s\n", tokenString)

		// Decodificar el token JWT sin verificar la firma
		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
		if err == nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				// Imprimir todos los reclamos para depuración
				fmt.Printf("Unverified Claims: %v\n", claims)
			}
		} else {
			fmt.Printf("Error parsing unverified token: %v\n", err)
		}

		// Decodificar el token JWT verificando la firma
		token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			fmt.Printf("Error parsing token: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Imprimir todos los reclamos para depuración
			fmt.Printf("Claims: %v\n", claims)

			idUser, _ := claims["id_user"].(float64)
			user, err := client.GetUserById(int(idUser))

			// Verificar si se extrajeron correctamente los reclamos
			if user.Id_user == 0 || err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}

			if user.Role != 1 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not an admin"})
				c.Abort()
				return
			}
		} else {
			fmt.Printf("Token is invalid or claims are not of type jwt.MapClaims\n")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
