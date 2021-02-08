package main

import(
	"fmt"
	"io/ioutil"
	"log"
	//"time"

	"microauth/token"
)

func main() {
	prvKey, err := ioutil.ReadFile("key/oauth-private.key")
	if err != nil {
		log.Fatalln(err)
	}
	pubKey, err := ioutil.ReadFile("key/oauth-public.key")
	if err != nil {
		log.Fatalln(err)
	}

	jwtToken := token.NewJWT(prvKey, pubKey)

	// 1. Create a new JWT token.
	// tok, err := jwtToken.Create(time.Hour, "Can be anything")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("TOKEN:", tok)

	// 2. Validate an existing JWT token.s
	man := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImUwNzMyNmJlYzg4Mzk4MGEzY2JlY2QyNTg4ZDE3MThjOGE2MzdhNTQyMTViZjJlOGVjOGQxYTRiZTUzNzA1NjM2NTA3OTE4MDBhZmE3NGQwIn0.eyJhdWQiOiI1MiIsImp0aSI6ImUwNzMyNmJlYzg4Mzk4MGEzY2JlY2QyNTg4ZDE3MThjOGE2MzdhNTQyMTViZjJlOGVjOGQxYTRiZTUzNzA1NjM2NTA3OTE4MDBhZmE3NGQwIiwiaWF0IjoxNjEyNjkzODA5LCJuYmYiOjE2MTI2OTM4MDksImV4cCI6MTY0NDIyOTgwOSwic3ViIjoiMSIsInNjb3BlcyI6W119.lNfPg_MLLEaCyx3PIm_AS1gqB2uqwKle8G0mnQESU4kuJH6lp5VOVkBtuA2v1WojwR6JAZbrbbirgw4xjdm8l1xgEV8lK9h12tY4NMtRXsxIojW-NMaM46c__aex7SxXkTkcRH2JTSBJ2ZDwlPuktHA6hBtHVlkw07NWStgYFDQa9Asyn39_w39GJnYgh2J_fDcjSLphLA1EJm0-CeX4KEfRhoVxWxks5eKfsEuYLvU1Ss705F1lcXRGel8ZtdCvo3O6ExqtWDTo6s9si4_Qnws6BRsp_HM-ZwoRyBVP14wMaUx6fimDkgJQgfsoKaDx3xStqkQr2Z-43xWnNBm7zMw0h6XNBGuH4rBmze7LXQ1Fg-ybRy-hUpwIx-TX0-UJ39cYHLLT_TZSoiNR3Xxapt8ICmQVkEJ7zlJzUWXAMvfY0_h8609-8rCSk1Mf72mvURQd3pn96Qa3fW2eE9ziRkE1yBAwieogZ270m2OGs-GCy7GQqnvg_T9AWIPdTNfQ9cXqC1jKSuxR3PjsnWvADAdAU9sEsHapGO8u_Heb8nB5vz3bsDN7Vbqq2CdNhp1Sz_n_hPJx4qtArqI2HzEcLISMGxQ2bPP-9NZ5irMM9nvs0boblhKTNimSB-nIV-TT7uaMxToQkvcjxI2ngE0E4HaixHJrfZ2n9BfhW_hCn_0"
	//man := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjNhMjkyOTY0ZGY2ZmU3OTk1NzYzNmEyNDg4NzIwYzI1OWRmOTFjYWE2OTljZTE4MjMwMTIwMDQxZDAxNzdkMzI0ZjM4NDcxNDM2ODQ3NmFmIn0.eyJhdWQiOiI2NCIsImp0aSI6IjNhMjkyOTY0ZGY2ZmU3OTk1NzYzNmEyNDg4NzIwYzI1OWRmOTFjYWE2OTljZTE4MjMwMTIwMDQxZDAxNzdkMzI0ZjM4NDcxNDM2ODQ3NmFmIiwiaWF0IjoxNjEyNzY3NzAzLCJuYmYiOjE2MTI3Njc3MDMsImV4cCI6MTY0NDMwMzcwMywic3ViIjoiMyIsInNjb3BlcyI6W119.lVBgFc-zTDSshYv1enoo54Kj9ke8swKMwR9-dDRMZ4-MPHnXYqcIXATL2l9_NTMRlstoqOgyTFiSGIB1wxIAQ2iJICNT8rPSsMsLfQ90ZD-o1xzn4YxAuB-A-MHuPBzTruOvv6nJmiZk19atAuV0KHUGDAR6KW4wr6sEsQH5oLf-Bt_HhMFTgLKm71-mIGQjK4HLodxLdjR_BP8SvqOzkD0oqJ7vNNVadjBpQXGf_BSphKfkUXmVWKxbwFQg1bA8d_-bISoW5tYwdTYNKahPTZTG2MGH7cfqbuVYda7-BdOQp_cUfzfbLeFjPQT3ledCojyq4LyefGshsj0I-es4SE27HxKR5_ZoHvNU1vOp16ASqnMuoIYgpuv-_X19Y4w4lB_HLusZfwKcYhLu627ZI_QqTacl8chGqf3XtAwOCJ6aBxocIrGpHqlYYSEkIYOCEmBdok8FbVxY4BKiHWgUhZ_E_dTqakfkAxyBM2NzpetypJ4XRNJKP77GTSNQbsMAwZpJydCN2iBsFQ9EtFUTGbSTGmb2PaVmHCfXL0QbQemHAHNk4Pe6G85g0IKq8w8yX4AYpk5lC4tRGwt6cfa9x8W9warkScEkMw2JQhzgE1Bq-63FBSw6NeagzngO7E1K9rJ6ZnSoAz70qTeXciws8ij5kyT9av8gSiRuupAXU34"
	content, err := jwtToken.Validate(man)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("CONTENT:", content)
}