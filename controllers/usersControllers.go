package controllers

import (
	"accountanApp/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUsersData(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT userId, name, username, type FROM users")

	if err != nil {
		sendErrorResponse(w, string(err.Error()), http.StatusBadRequest)
		return
	} else {
		var user models.User
		var users []models.User

		for rows.Next() {
			if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Type); err != nil {
				fmt.Print(err)
				sendErrorResponse(w, "Error Field Undefined", http.StatusBadRequest)
				return
			} else {
				users = append(users, user)
			}
		}

		if len(users) != 0 {
			sendSuccessResponse(w, users)
		} else {
			sendErrorResponse(w, "No data in table users", http.StatusBadRequest)
		}
	}
}

func InsertNewSaving(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var saving models.Saving
	var user models.User
	err := decoder.Decode(&saving)
	if err != nil {
		fmt.Print(err)
		return
	}

	queryA := "SELECT userId, name, username, saving_balance FROM users WHERE userId=?"
	row := db.QueryRow(queryA, &saving.UserID)

	if err := row.Scan(&user.ID, &user.Name, &user.Username, &user.SavingBalance); err != nil {
		sendErrorResponse(w, "No user found", http.StatusBadRequest)
		return
	}

	fmt.Println("Saving balance: ", user.SavingBalance)

	saving.Jsp = user.SavingBalance / 5000
	currBalance := user.SavingBalance + saving.Ss

	_, errQ := db.Exec("INSERT INTO savings(user_id, date, ss, jsp) VALUES (?,?,?,?)",
		saving.UserID,
		saving.Date,
		saving.Ss,
		saving.Jsp,
	)

	if errQ != nil {
		fmt.Println(errQ)
		sendErrorResponse(w, "Error saving new data", http.StatusBadRequest)
		return
	} else {
		fmt.Println("user id: ", user.ID)
		fmt.Println("current balance: ", currBalance)

		_, errQ2 := db.Exec("UPDATE users SET saving_balance=? WHERE userId=?",
			currBalance,
			user.ID,
		)

		if errQ2 != nil {
			fmt.Println(errQ)
			sendErrorResponse(w, "Error update new saving balance", http.StatusBadRequest)
			return
		} else {
			var model models.UserSaving
			model.User = user
			model.Saving = saving
			sendSuccessResponse(w, model)
		}
	}
}
