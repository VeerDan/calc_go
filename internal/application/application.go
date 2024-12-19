package application

import (
	"encoding/json"
	"net/http"
	"os"
	"fmt"
	calculation "github.com/VeerDan/calc_go/pkg/calculation"
	"log/slog"
	"time"
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

type Response struct {
	Result float64 `json:"result"`
}

type Request struct {
	Expression string `json:"expression"`
}

func TimeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		t := time.Now()
		elapsed := t.Sub(start)
		slog.Info(fmt.Sprintf("Время ответа сервера: %v", elapsed))
	})
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error(fmt.Sprintf("error: Internal server error; status_code: %d", 500))
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
		return
	}
	result, err := calculation.Calc(request.Expression)
	if err != nil {
		slog.Error(fmt.Sprintf("error: %s; status_code: %d", err, 422))
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err), http.StatusUnprocessableEntity )
	} else {
		slog.Info(fmt.Sprintf("result: %f; status_code: %d", result, 200))
		fmt.Fprintf(w, `{"result":"%f"}`, result)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/", TimeMiddleware(CalcHandler))
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
