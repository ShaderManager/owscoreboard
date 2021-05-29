package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"

	hook "github.com/robotn/gohook"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port int `yaml:"port"`

	Keymaps struct {
		IncTankWins   string `yaml:"increment-tank-wins"`
		IncTankTies   string `yaml:"increment-tank-ties"`
		IncTankLosses string `yaml:"increment-tank-losses"`
		IncDpsWins    string `yaml:"increment-dps-wins"`
		IncDpsTies    string `yaml:"increment-dps-ties"`
		IncDpsLosses  string `yaml:"increment-dps-losses"`
		IncSupWins    string `yaml:"increment-sup-wins"`
		IncSupTies    string `yaml:"increment-sup-ties"`
		IncSupLosses  string `yaml:"increment-sup-losses"`
		ResetStats    string `yaml:"reset-stats"`
	} `yaml:"keymaps"`
}

func NewConfig() Config {
	res := Config{}
	res.Port = 8080

	res.Keymaps.IncTankWins = "numpad 7"
	res.Keymaps.IncTankTies = "numpad 8"
	res.Keymaps.IncTankLosses = "numpad 9"
	res.Keymaps.IncDpsWins = "numpad 4"
	res.Keymaps.IncDpsTies = "numpad 5"
	res.Keymaps.IncDpsLosses = "numpad 6"
	res.Keymaps.IncSupWins = "numpad 1"
	res.Keymaps.IncSupTies = "numpad 2"
	res.Keymaps.IncSupLosses = "numpad 3"
	res.Keymaps.ResetStats = "numpad 0"

	return res
}

func loadConfig(path string) (Config, error) {
	f, err := os.Open(path)
	cfg := NewConfig()

	if err != nil {
		return cfg, err
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Printf("Failed to parse %s: %+v", path, err)
	}

	return cfg, err
}

var cfg Config

func init() {
	cfg = NewConfig()
}

func main() {
	cfg, _ = loadConfig("config.yaml")

	go setupHook()

	// Workaround for Windows
	mime.AddExtensionType(".js", "text/javascript")

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/getResults", getResultsHandler)

	log.Printf("Starting web server at http://localhost:%d/", cfg.Port)
	http.ListenAndServe(":"+fmt.Sprintf("%d", cfg.Port), mux)
}

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

type Results struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Ties   int `json:"ties"`
}

var tankResults = &Results{0, 0, 0}
var dpsResults = &Results{0, 0, 0}
var supResults = &Results{0, 0, 0}

func getResultsHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	class := params.Get("class")

	var res *Results

	switch {
	case class == "tank":
		res = tankResults
	case class == "dps":
		res = dpsResults
	case class == "sup":
		res = supResults
	default:
		http.Error(w, "Invalid class", http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(res)
	w.Write(data)
}

func setupHook() {
	defer hook.End()

	hook.Register(hook.KeyDown, []string{""}, func(e hook.Event) {
		increment := func() int {
			if (e.Mask & 2) != 0 {
				return -1
			} else {
				return 1
			}
		}()

		switch e.Rawcode {
		case hook.KeychartoRawcode(cfg.Keymaps.IncSupWins):
			supResults.Wins += increment
		case hook.KeychartoRawcode(cfg.Keymaps.IncSupTies):
			supResults.Ties += increment
		case hook.KeychartoRawcode(cfg.Keymaps.IncSupLosses):
			supResults.Losses += increment
		case hook.KeychartoRawcode(cfg.Keymaps.IncDpsWins):
			dpsResults.Wins += increment
		case hook.KeychartoRawcode(cfg.Keymaps.IncDpsTies):
			dpsResults.Ties += increment
		case hook.KeychartoRawcode(cfg.Keymaps.IncDpsLosses):
			dpsResults.Losses += increment
		case hook.KeychartoRawcode(cfg.Keymaps.IncTankWins):
			tankResults.Wins += increment
		case hook.KeychartoRawcode(cfg.Keymaps.IncTankTies):
			tankResults.Ties += increment
		case hook.KeychartoRawcode(cfg.Keymaps.IncTankLosses):
			tankResults.Losses += increment
		case hook.KeychartoRawcode(cfg.Keymaps.ResetStats):
			tankResults = &Results{0, 0, 0}
			dpsResults = &Results{0, 0, 0}
			supResults = &Results{0, 0, 0}
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
