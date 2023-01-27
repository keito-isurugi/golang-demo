package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang_demo/middleware"
	"golang_demo/db"
	"golang_demo/models"
)

var secret = []byte("secret")

func Auth() {
	http.HandleFunc("/api/login", middleware.CORS(loginHndler))
	http.HandleFunc("/api/auth", middleware.CORS(authHndler))

	// log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

type LoginData struct {
	// jsonで型定義
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func loginHndler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(login(w, r))
}
func login(w http.ResponseWriter, r *http.Request) string{

	// 入力された値を取得
	var loginData LoginData
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		log.Fatal(err)
	}

	// ユーザーの存在確認	
	user := db.GetLoginUser(loginData.Email, loginData.Password)
	fmt.Println(user)

	// ここで、ユーザー名とパスワードを検証し、正当な場合にのみJWTトークンを生成する
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	var tokenString string
	var err error
	if user.ID != 0 {
		// JWTトークンを生成し
		fmt.Println("存在します！")
		claims["user_id"] = user.ID 
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		tokenString, err = token.SignedString(secret)
		fmt.Println(tokenString)
	} else {
		fmt.Println("存在しません、、")
		fmt.Println(err)
		return err.Error()
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return err.Error()
	}
	
	return tokenString

}


func authHndler(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(auth(w, r))
}
func auth(w http.ResponseWriter, r *http.Request) models.User{
    // クライアントから送られてきたJWTトークンを検証する
    tokenString := r.Header.Get("Authorization")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// トークンの検証
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return secret, nil
    })

    if err != nil {
			log.Println(err)
			http.Error(w, "Invalid token 1", http.StatusUnauthorized)
			// return
    }

		var user models.User
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// トークンが有効な場合の処理
		fmt.Printf("user_id: %v\n", int64(claims["user_id"].(float64)))
		fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
		
		userId := int(int64(claims["user_id"].(float64)))
		user = db.GetUser(userId)
		} else {
			http.Error(w, "Invalid token 1", http.StatusUnauthorized)
		}
		return user
}
