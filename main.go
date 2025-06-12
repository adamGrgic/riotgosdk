package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adamGrgic/riotgosdk/publicconstants/gamemodes"
	"github.com/adamGrgic/riotgosdk/publicconstants/leagues"
	"github.com/adamGrgic/riotgosdk/riotclient"
	"github.com/joho/godotenv"
)

func GetLeagueEntry(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	leagueId := query.Get("leagueId")

	res, err := riotclient.GetLeagueEntry(apiKey, leagueId)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetLeagueEntries(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()

	gameMode := gamemodes.GameMode(query.Get("gameMode"))
	league := leagues.League(query.Get("league"))
	tier := query.Get("tier")

	res, err := riotclient.GetLeagueEntries(apiKey, gameMode, league, tier)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetLeagueEntryFromPuuid(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	leagueId := query.Get("leagueId")

	res, err := riotclient.GetLeagueEntryFromPuuid(apiKey, leagueId)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetLeagueEntryFromSummonerId(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	leagueId := query.Get("leagueId")

	res, err := riotclient.GetLeagueEntryFromSummonerId(apiKey, leagueId)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetChallengerLeagueEntriesFromQueueId(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	queueId := query.Get("queueId")

	res, err := riotclient.GetChallengerLeagueEntriesFromQueueId(apiKey, queueId)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetGrandmasterLeagueEntriesFromQueueId(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	queueId := query.Get("queueId")

	res, err := riotclient.GetGrandmasterLeagueEntriesFromQueueId(apiKey, queueId)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetMasterLeagueEntriesFromQueueId(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	queueId := query.Get("queueId")

	res, err := riotclient.GetMasterLeagueEntriesFromQueueId(apiKey, queueId)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetRecentMatches(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	puuid := query.Get("puuid")
	start := query.Get("start")
	count := query.Get("count")

	res, err := riotclient.GetRecentMatches(w, apiKey, puuid, start, count)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetMatchData(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	matchId := query.Get("matchId")

	res, err := riotclient.GetMatchData(w, apiKey, matchId)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetMatchTimelineData(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	matchId := query.Get("matchId")

	res, err := riotclient.GetMatchTimelineData(w, apiKey, matchId)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetAccountFromGameName(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	tagLine := query.Get("tagLine")
	gameName := query.Get("gameName")

	// fmt.Fprintf(w, "tagLine: %s\ngameName: %s", tagLine, gameName)

	res, err := riotclient.GetAccountFromGameName(w, apiKey, gameName, tagLine)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetAccountFromPuuid(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	puuid := query.Get("puuid")

	// fmt.Fprintf(w, "tagLine: %s\ngameName: %s", tagLine, gameName)

	res, err := riotclient.GetAccountFromPuuid(w, apiKey, puuid)
	if err != nil {
		errMsg := fmt.Sprintf("Something went wrong getting league entries: %s", err)
		w.Write([]byte(errMsg))
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the Content-Type header and write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Riot API Service Response OK")
}

func main() {
	fmt.Println("Starting up Riot API Service...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", Ping)
	http.HandleFunc("/league-entry", GetLeagueEntry)
	http.HandleFunc("/league-entries", GetLeagueEntries)
	http.HandleFunc("/league-entry-from-puuid", GetLeagueEntryFromPuuid)
	http.HandleFunc("/league-entry-from-summoner-id", GetLeagueEntryFromSummonerId)
	http.HandleFunc("/master-leagues-from-queue-id", GetMasterLeagueEntriesFromQueueId)
	http.HandleFunc("/grandmaster-leagues-from-queue-id", GetGrandmasterLeagueEntriesFromQueueId)
	http.HandleFunc("/challenger-leagues-from-queue-id", GetChallengerLeagueEntriesFromQueueId)
	http.HandleFunc("/recent-matches", GetRecentMatches)
	http.HandleFunc("/match-data", GetMatchData)
	http.HandleFunc("/match-timeline", GetMatchTimelineData)
	http.HandleFunc("/account-from-puuid", GetAccountFromPuuid)
	http.HandleFunc("/account-from-game-name", GetAccountFromGameName)

	port := ":8081"
	fmt.Println("Riot API Service is running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
