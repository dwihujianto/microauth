package main

import(
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	//fmt.Println(parts, parts[0:2], parts[2])

	pubKey, _ := ioutil.ReadFile("key/oauth-public.key")
	key, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)

	if err != nil {
		fmt.Printf("Haduh", err)
	}

	tokenString := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImUwNzMyNmJlYzg4Mzk4MGEzY2JlY2QyNTg4ZDE3MThjOGE2MzdhNTQyMTViZjJlOGVjOGQxYTRiZTUzNzA1NjM2NTA3OTE4MDBhZmE3NGQwIn0.eyJhdWQiOiI1MiIsImp0aSI6ImUwNzMyNmJlYzg4Mzk4MGEzY2JlY2QyNTg4ZDE3MThjOGE2MzdhNTQyMTViZjJlOGVjOGQxYTRiZTUzNzA1NjM2NTA3OTE4MDBhZmE3NGQwIiwiaWF0IjoxNjEyNjkzODA5LCJuYmYiOjE2MTI2OTM4MDksImV4cCI6MTY0NDIyOTgwOSwic3ViIjoiMSIsInNjb3BlcyI6W119.lNfPg_MLLEaCyx3PIm_AS1gqB2uqwKle8G0mnQESU4kuJH6lp5VOVkBtuA2v1WojwR6JAZbrbbirgw4xjdm8l1xgEV8lK9h12tY4NMtRXsxIojW-NMaM46c__aex7SxXkTkcRH2JTSBJ2ZDwlPuktHA6hBtHVlkw07NWStgYFDQa9Asyn39_w39GJnYgh2J_fDcjSLphLA1EJm0-CeX4KEfRhoVxWxks5eKfsEuYLvU1Ss705F1lcXRGel8ZtdCvo3O6ExqtWDTo6s9si4_Qnws6BRsp_HM-ZwoRyBVP14wMaUx6fimDkgJQgfsoKaDx3xStqkQr2Z-43xWnNBm7zMw0h6XNBGuH4rBmze7LXQ1Fg-ybRy-hUpwIx-TX0-UJ39cYHLLT_TZSoiNR3Xxapt8ICmQVkEJ7zlJzUWXAMvfY0_h8609-8rCSk1Mf72mvURQd3pn96Qa3fW2eE9ziRkE1yBAwieogZ270m2OGs-GCy7GQqnvg_T9AWIPdTNfQ9cXqC1jKSuxR3PjsnWvADAdAU9sEsHapGO8u_Heb8nB5vz3bsDN7Vbqq2CdNhp1Sz_n_hPJx4qtArqI2HzEcLISMGxQ2bPP-9NZ5irMM9nvs0boblhKTNimSB-nIV-TT7uaMxToQkvcjxI2ngE0E4HaixHJrfZ2n9BfhW_hCn_0"
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTI2ODAzNDgsInVzZXJfaWQiOjF9.KOl_c5PMhptXo6KMkwN46YeCWR4oqbR1m2XrwEFH9Es"

	parts := strings.Split(tokenString, ".")
	method := jwt.GetSigningMethod("RS256")

	err = method.Verify(strings.Join(parts[0:2], "."), parts[2], key)

	if err != nil {
		fmt.Printf("[%v] Error while verifying key: %v", "RS256", err)
	}

	ok, duh := method.Sign(strings.Join(parts[0:2], "."), key)

	if duh != nil {
		fmt.Printf("[%v] Duh: %v", "RS256", duh)
	} else {
		fmt.Println(ok)
	}

}