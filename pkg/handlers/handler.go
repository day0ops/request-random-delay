package handlers

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	c "github.com/day0ops/request-random-delay/pkg/config"
)

type Resp struct {
	Message string
}

type Option func(*Handler)

func LogWith(logger *zap.Logger) Option {
	return func(h *Handler) {
		h.logger = logger
	}
}

type Handler struct {
	logger *zap.Logger
	mux    *http.ServeMux
}

func NewHandler(options ...Option) *Handler {
	h := &Handler{}

	for _, o := range options {
		o(h)
	}

	h.mux = http.NewServeMux()
	h.mux.HandleFunc("/", h.index)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Sugar().Infof("%s %s", r.Method, r.URL.Path)

	h.mux.ServeHTTP(w, r)
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	baseDelay, err := strconv.Atoi(c.BaseDelay)
	if err != nil {
		panic(err)
	}

	randomDelay := rand.Intn(100)
	totalDelay := time.Duration(baseDelay+randomDelay) * time.Millisecond
	time.Sleep(totalDelay)

	responseMessage := fmt.Sprintf("Response from server after %d ms", totalDelay.Milliseconds())
	m := Resp{Message: responseMessage}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(m)
	if err != nil {
		http.Error(w, fmt.Sprintf("error building the response, %v", err), http.StatusInternalServerError)
		return
	}
}
