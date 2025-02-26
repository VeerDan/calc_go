package application

import (
	"encoding/json"
	"net/http"
	"os"
	"fmt"
	calculation "github.com/VeerDan/calc_go/pkg/calculation"
	"log/slog"
	"time"
	"sync"
)


var Expressions []*Expression
var Mu sync.Mutex
var Id int = 0

type Expression struct {
	Id int
	Status string
	Result int
}

func AssignId() int {
	Mu.Lock()
	defer Mu.Unlock()
	Id += 1
	return Id
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error(fmt.Sprintf("error: %v; status_code: %d", err, 500))
		http.Error(w, fmt.Sprintf(`{"error":"%v"}`, err), http.StatusInternalServerError)
		return
	}
	if calculation.IsValid(request.Expression) == nil {
		Mu.Lock()
		defer Mu.Unlock()
		expression := new(Expression)
		expression.Id = AssignId()
		expression.Status = "Calculating"
		expression.Result = 0
		Expressions = append(Expressions, expression)
		slog.Info(fmt.Sprintf("working on expression; id: %v", expression.Id))
		fmt.Fprintf(w, `{"id":"%v"}`, expression.Id)
	} else {
		slog.Error(fmt.Sprintf("error: Unvalid expression; status_code: %d", 422))
		http.Error(w, `{"error":"Unvalid expression"}`, http.StatusUnprocessableEntity)
	}
}


func ExpressionHandler(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(Expressions)
	if err != nil {
		slog.Error("error: Internal server error; status code: 500")
		http.Error(w, `{"error":"Internal server error"}`, 500)
	}
	fmt.Fprint(w, string(res))
}


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
		slog.Info(fmt.Sprintf("result: %v; status_code: %d", result, 200))
		fmt.Fprintf(w, `{"result":"%v"}`, result)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/", TimeMiddleware(CalcHandler))
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
