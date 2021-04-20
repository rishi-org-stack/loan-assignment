package view

import (
	"encoding/json"
	"fmt"
	"strings"

	// "fmt"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/gorilla/mux"
	"github.com/rishi-org-stack/loan/model"
	"github.com/rishi-org-stack/loan/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Query struct {
	Where interface{}
	Val   interface{}
}

const (
	database = "Loan"
)

var db = model.Instantiate().Connect().CreateDb(database).LinkToCollection("user")
var res = &response.Response{}

func Shout(w http.ResponseWriter, r *http.Request) {
	val := make(map[string]string)
	val["weather"] = "is great"
	json.NewEncoder(w).Encode(val)
}

func Register(w http.ResponseWriter, r *http.Request) {
	res.Method = r.Method
	w.Header().Add("content-type", "application/json")
	var p model.Loan
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		res.ServerError(err.Error(), r.Body)
	}
	p.Status = "New"
	if !p.SomethingIsNotPresent() {
		db.Insert(p)
		res.Success("Loan is Successfully Added", p)
	}else{
		res.InsufficientCredtials("some fields are missing",nil)
	}

	json.NewEncoder(w).Encode(res)
}

func GetAllLoan(w http.ResponseWriter, r *http.Request) {
	res.Method = r.Method
	w.Header().Add("content-type", "json")
	q := r.URL.Query()

	if len(q) == 0 {
		loans := db.GetAllLoan()
		res.Success("ALL LOAN is extracted", loans)
	} else {
		if len(q) == 2 {
			string_loan := q["loanAmountGreater"][0]
			s, _ := strconv.ParseFloat(string_loan, 64)

			loans := (db.GetAllLoansWithLoanFilter(s))
			final_loan := make([]model.Loan, 0)
			fmt.Println(len(q["status"]))
			ans := strings.Split(q.Get("status"), ",")
			fmt.Println(ans)
			if len(ans) == 2 {
				for _, val := range loans {
					if val.Status == ans[0] || val.Status == ans[1] {

						final_loan = append(final_loan, val)
					}

				}
			} else {
				for _, val := range loans {
					if val.Status == q.Get("status") {
						final_loan = append(final_loan, val)
					}

				}
			}

			res.Success("a few loans is selected", final_loan)
		}
		if len(q) == 1 {
			loans := db.GetAllLoan()
			final_loan := make([]model.Loan, 0)
			ans := strings.Split(q.Get("status"), ",")
			fmt.Println(ans)
			if len(ans) == 2 {
				for _, val := range loans {
					if val.Status == ans[0] || val.Status == ans[1] {

						final_loan = append(final_loan, val)
					}

				}
			} else {
				for _, val := range loans {
					if val.Status == q.Get("status") {
						final_loan = append(final_loan, val)
					}

				}
			}

			res.Success("a few loans is selected", final_loan)
		}

	}

	json.NewEncoder(w).Encode(res)
}

func GetaLoan(w http.ResponseWriter, r *http.Request) {
	res.Method = r.Method
	w.Header().Add("content-type", "json")
	var loan model.Loan
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		res.ServerError(err.Error(), nil)
	}
	db.Get("_id", id, &loan)
	if loan.CustomerName == "" {
		res.NosuchDoc("no such Document", nil)
	} else {
		res.Success("process succesfull", loan)

	}
	json.NewEncoder(w).Encode(res)
}

func UpdateLoan(w http.ResponseWriter, r *http.Request) {
	res.Method = r.Method
	w.Header().Add("content-type", "json")
	var q model.Loan
	json.NewDecoder(r.Body).Decode(&q)
	vars := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(vars["id"])
	find := make(map[string]interface{})
	find["identity"] = "_id"
	find["val"] = id
	val := make(map[string]interface{})
	val["status"] = q.Status
	c, err := db.UpdateaDocument(find, val)
	if err != nil {
		res.ErrorUpdateDoc(err.Error(), c)
	}
	res.Success("Update succesfull", q)

	json.NewEncoder(w).Encode(res)
}

func DeleteLoan(w http.ResponseWriter, r *http.Request) {
	res.Method = r.Method
	w.Header().Add("content-type", "json")
	var q = make(map[string]interface{})
	vars := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(vars["id"])
	// find:=make(map[string]interface{})
	q["identity"] = "_id"
	q["val"] = id
	count, err := db.DeleteADocument(q)
	if err != nil {
		res.ErrorDeleteDoc(err.Error(), count)
	}
	res.Success("Delete succesfull", q)
	json.NewEncoder(w).Encode(res)
}
