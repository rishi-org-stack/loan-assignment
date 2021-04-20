package model

import (
	"context"
	// "fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type Loan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerName string             `bson:"customer_name" json:"customer_name"`
	PhoneNo      string             `bson:"phone_no" json:"phone_no"`
	Email        string             `bson:"email" json:"email"`
	LoanAmount   float64            `bson:"loan_amount" json:"loan_amount"`
	Status       string             `bson:"status" json:"status"`
	CreditScore  int                `bson:"credit_score" json:"credit_score"`
}

func (l *Loan) SomethingIsNotPresent() (ans bool) {
	ans = false
	if l.CustomerName==""||l.Email==""||l.PhoneNo==""||l.Status==""||l.LoanAmount==0||l.CreditScore==0{
		ans=true
	}
	return
}

func (db *DB) GetAllLoan() []Loan {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	cursor, err := db.Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var objects []Loan
	if err = cursor.All(ctx, &objects); err != nil {
		log.Fatal(err)

	}
	defer cancel()
	return objects
}

func (db *DB) GetAllLoansWithLoanFilter(val float64) []Loan {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	cursor, err := db.Collection.Find(ctx, bson.D{{"loan_amount", bson.D{{"$gt", val}}}})
	if err != nil {
		log.Fatal(err)
	}
	var objects []Loan
	if err = cursor.All(ctx, &objects); err != nil {
		log.Fatal(err)

	}
	defer cancel()
	return objects
}
func (db *DB) GetAllLoansWithStatusFilter(what string, status string) []Loan {
	coll := db.Collection
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	cursor, err := coll.Find(ctx, bson.M{"status": status})
	var objects []Loan
	if err = cursor.All(ctx, &objects); err != nil {
		log.Fatal(err)

	}

	return objects
}
