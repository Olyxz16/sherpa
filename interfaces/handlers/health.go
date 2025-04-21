package handlers

import (
	"encoding/json"
	"net/http"

    "go.uber.org/zap"
	db "github.com/Olyxz16/sherpa/infrastructure/persistence"
)

/*********************/
/* Healthcheck utils */
/*********************/

func Health(w http.ResponseWriter, r *http.Request) {
	service, err := db.Instance()
    if err != nil || !service.Health() {
        zap.L().Error("HEALTHCHECK NOT PASSING !")
        resp, _ := json.Marshal(map[string]string {"message": "KO"})
        w.WriteHeader(500)
        w.Write(resp)
		return
    }
    resp, _ := json.Marshal(map[string]string {"message": "OK"})
    w.Write(resp)
}
