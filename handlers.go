package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/jacobmichaellemon/language-learner/internal/data"
	"github.com/jacobmichaellemon/language-learner/internal/download"
)

const sessionCookieName = "vocab_session_id"

type QuizApp struct {
	Questions []data.Translation
	Tmpl      *template.Template
	// Thread-safe map to track current question index per user/session
	// Key: session ID or simple ID -> Value: current question index
	mu       sync.RWMutex
	Sessions map[string]UserSession
	DB       *sql.DB
}

type UserSession struct {
	QuestionIndex   int
	CurrentIndex    int
	Total           int
	Score           int
	Started         bool
	ToLang          string
	FromLang        string
	NumberQuestions int
	Languages       map[string]string
}

func (app *QuizApp) handleCreateQuiz(w http.ResponseWriter, r *http.Request) {
	sessionID := getOrCreateSessionID(w, r)
	app.mu.RLock()
	session := app.Sessions[sessionID]
	app.mu.RUnlock()

	// Check if the quiz is started
	if session.Started {
		http.Redirect(w, r, "/quiz", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		// Parse submitted form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		nativeLang := r.FormValue("native_lang")
		targetLang := r.FormValue("target_lang")
		numQuestions, _ := strconv.Atoi(r.FormValue("num_questions"))

		db, err := download.GetDictionary(nativeLang, targetLang)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		app.DB = db

		questions, err := GetQuizWords(db, numQuestions, 3.0)
		if err != nil {
			log.Fatal(err)
		}
		app.Questions = questions

		// Construct state for the active quiz
		data := UserSession{
			ToLang:          nativeLang,
			FromLang:        targetLang,
			NumberQuestions: len(app.Questions),
			Started:         true,
		}
		app.Sessions[sessionID] = data
		// Redirect back to GET "/" to show the next question
		http.Redirect(w, r, "/quiz", http.StatusSeeOther)
	}
	data := UserSession{Started: false, Languages: languages}
	tmpl.ExecuteTemplate(w, "setup.html", data)
}

func (app *QuizApp) handleQuiz(w http.ResponseWriter, r *http.Request) {
	sessionID := getOrCreateSessionID(w, r)
	app.mu.RLock()
	session := app.Sessions[sessionID]
	app.mu.RUnlock()

	// Check if the quiz is finished
	if session.CurrentIndex >= len(app.Questions) {
		http.Redirect(w, r, "/results", http.StatusSeeOther)
		return
	}

	currentQ := app.Questions[session.CurrentIndex]

	data := struct {
		QuestionIndex int
		Total         int
		Question      data.Translation
	}{
		QuestionIndex: session.CurrentIndex + 1,
		Total:         len(app.Questions),
		Question:      currentQ,
	}

	tmpl.ExecuteTemplate(w, "quiz.html", data)
}

func (app *QuizApp) handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/quiz", http.StatusSeeOther)
		return
	}

	sessionID := getOrCreateSessionID(w, r)
	userAnswer := r.FormValue("answer")

	app.mu.Lock()
	session := app.Sessions[sessionID]
	currentQ := app.Questions[session.CurrentIndex]
	isCorrect := false

	// 1. Evaluate answer logic here (e.g., update user score)
	for word := range strings.SplitSeq(currentQ.TransList, "|") {
		if strings.EqualFold(userAnswer, strings.TrimSpace(word)) {
			fmt.Println("Correct!")
			isCorrect = true
			session.Score++
			break
		}
	}
	if !isCorrect {
		fmt.Println("Incorrect")
	}
	// 2. Advance to the next question for this session
	session.CurrentIndex++
	app.Sessions[sessionID] = session
	app.mu.Unlock()

	// Redirect back to GET "/" to show the next question
	http.Redirect(w, r, "/quiz", http.StatusSeeOther)
}

func (app *QuizApp) handleResults(w http.ResponseWriter, r *http.Request) {

	sessionID := getOrCreateSessionID(w, r)
	session := app.Sessions[sessionID]

	// Pass the score data to results.html
	data := struct {
		Score int
		Total int
	}{
		Score: session.Score,
		Total: len(app.Questions),
	}

	// Ensure template execution errors are logged
	err := app.Tmpl.ExecuteTemplate(w, "results.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *QuizApp) handleReset(w http.ResponseWriter, r *http.Request) {
	// Clear session/cookies explicitly on restart
	cookie := &http.Cookie{
		Name:   "vocab_session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Deletes cookie in browser
	}
	http.SetCookie(w, cookie)

	// Redirect back to root setup form
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// getOrCreateSessionID checks if the user already has a session cookie.
// If not, it generates a new unique ID and sets a cookie in their browser.
func getOrCreateSessionID(w http.ResponseWriter, r *http.Request) string {
	// 1. Try to read the existing session cookie from the request
	cookie, err := r.Cookie(sessionCookieName)
	if err == nil && cookie.Value != "" {
		return cookie.Value
	}

	// 2. If no cookie exists, generate a random unique ID
	sessionID := generateRandomID()

	// 3. Set the cookie on the user's browser response
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true, // Prevents JavaScript from reading the cookie (security best practice)
		SameSite: http.SameSiteLaxMode,
	})

	return sessionID
}

// generateRandomID creates a secure 16-byte hex string (e.g., "a3f1c20e...")
func generateRandomID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
