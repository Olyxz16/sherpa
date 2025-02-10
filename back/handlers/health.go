package handlers

import (
	"encoding/json"
	"net/http"

    "go.uber.org/zap"

	"github.com/Olyxz16/sherpa/model"
)

/*********************/
/* Healthcheck utils */
/*********************/

func Health(w http.ResponseWriter, r *http.Request) {
    if !model.New().Health() {
        zap.L().Error("HEALTHCHECK NOT PASSING !")
        resp, _ := json.Marshal(map[string]string {"message": "KO"})
        w.WriteHeader(500)
        w.Write(resp)
    }
    resp, _ := json.Marshal(map[string]string {"message": "OK"})
    w.Write(resp)
}
