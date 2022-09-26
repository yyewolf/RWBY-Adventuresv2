package export

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"rwby-adventures/config"

	"github.com/yyewolf/goth"
)

func Code(code string) (*goth.User, error) {
	fmt.Println(os.Getenv("AUTH_KEY"))
	values := map[string]string{
		"code":     code,
		"password": os.Getenv("AUTH_KEY"),
	}
	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%sapi/login/code", config.AuthHost), "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
	m := res["user"].(map[string]interface{})
	d, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	var u goth.User
	err = json.Unmarshal(d, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
