package application

import (
	"encoding/json"
	"net/http"
	"os"
	"fmt"
	calculation "github.com/VeerDan/calc_go/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}
type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}


type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, `"error": "Internal server error"`, http.StatusInternalServerError)
		return
	}
	result, err := calculation.Calc(request.Expression)
	if err != nil {
		http.Error(w, fmt.Sprintf(`"error": "%s"`, err), http.StatusUnprocessableEntity )
	} else {
		fmt.Fprintf(w, `"result": "%f"`, result)
	}

}

func (a *Application) RunServer() error {
	http.HandleFunc("/", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
