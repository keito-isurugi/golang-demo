package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang_demo/middleware"
)

var secret = []byte("secret")

func Auth() {
	http.HandleFunc("/api/login", middleware.CORS(loginHndler))
	http.HandleFunc("/restricted", restricted)

	// log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

type LoginData struct {
	// jsonで型定義
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Jwt struct {
	Token string `json:"token"`
}

func loginHndler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(login(w, r))
	// login(w, r)
}

func login(w http.ResponseWriter, r *http.Request) Jwt{
	// テストデータ
	var testData LoginData
	testData.Id = 123
	testData.Email = "test1@email.com"
	testData.Password = "test1"

	// 入力された値を取得
	var loginData LoginData
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		log.Fatal(err)
	}
	fmt.Println(loginData.Email, loginData.Password)
	
	// ここで、ユーザー名とパスワードを検証し、正当な場合にのみJWTトークンを生成する
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	var tokenString Jwt
	var err error
	if testData.Email == loginData.Email && testData.Password == loginData.Password {
		fmt.Println("正解！")
		// JWTトークンを生成し、クライアントに返す
		claims["user_id"] = testData.Id 
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		tokenString.Token, err = token.SignedString(secret)
		// w.Write([]byte(tokenString))
	fmt.Println(tokenString)
	} else {
		// return err.Error()
		// tokenString, err = token.SignedString(secret)
		// w.Write([]byte(tokenString))
		// fmt.Println("残念、、")
		fmt.Println(err)
	}
	return tokenString

	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "Error creating token", http.StatusInternalServerError)
	// 	return err.Error()
	// }

}

func restricted(w http.ResponseWriter, r *http.Request) {
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
			return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// トークンが有効な場合の処理
			// fmt.Println(claims)
			fmt.Printf("user_id: %v\n", int64(claims["user_id"].(float64)))
			fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
			// 省略
    } else {
			http.Error(w, "Invalid token 1", http.StatusUnauthorized)
    }
}
