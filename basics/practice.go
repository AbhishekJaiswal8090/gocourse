package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	fmt.Println("==== Welcome To The Axis Bank ====")
	fmt.Println("Our Services are\n")
	fmt.Println("1.Creating Account\n")
	fmt.Println("2.Deposit Money\n")
	fmt.Println("3.Withdraw Money\n")
	var choice int
	fmt.Println("Enter Your choice Plaese\n 1 For Creating Saving Account\n 2.For Depositing MOney \n 3.For Withdrawl of Money\n")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		CreateAccount()
	default:
		fmt.Println("You are sorry to bother you")
	}

}

type Bank struct {
	domesticUser PersonDomestic
	merchantUser PersonMerchant
}

type PersonDomestic struct {
	id            int
	name          string
	address       Address
	total_balance int
}

type Address struct {
	city  string
	phone string
}

type PersonMerchant struct {
	id              int
	name            string
	buisnessname    string
	buisnessrevenue int
	address         AddressMerchant
	total_balance   int
}

type AddressMerchant struct {
	city  string
	phone string
}

func CreateAccount() {
	var typeOfAccount string
	fmt.Println(" => What type of account do you want to create \n")
	fmt.Println("1.Domestic \n2.Merchant \n=> enter 1 for domestic or 2 for merchant")
	fmt.Scan(&typeOfAccount)
	switch typeOfAccount {
	case "1":
		id, detail := createSavingAccount()
		fmt.Println(detail)
		fmt.Printf("\nYour Saving Account has been created your account no. is: %d ,heres the detail %s \n", id, detail)
	case "2":
		id, detail := CreateMerchantAccount()
		fmt.Printf("\n Your Merchant Account has been created Your account no. is: %d,heres the detail %d \n", id, detail)
	default:
		fmt.Println("Sorry to bother You ,My mistake")
	}

}

func createSavingAccount() (interface{}, string) {
	var name string
	var address string
	var phone string
	fmt.Print(" => Enter Your Full name\n")
	fmt.Scan(&name)

	fmt.Println(" => Enter you Address\n")
	fmt.Scan(&address)

	fmt.Println(" => Enter your phone num\n")
	fmt.Scan(&phone)

	rand.Seed(time.Now().UnixNano()) // Seed with current time
	user_id := rand.Intn(1000000)

	total_id := make(map[int]string)
	total_id[user_id] = name
	fmt.Println("User account has been created\n")
	userAddress := Address{
		city:  address,
		phone: phone,
	}
	userDetail := PersonDomestic{
		id:            user_id,
		name:          name,
		address:       userAddress,
		total_balance: 0,
	}
	
	detailBytes, err := json.Marshal(userDetail)
	if err != nil {
		return userDetail.id, "Error serializing user details"
	}
	return userDetail.id, string(detailBytes)
}

func (b *PersonDomestic) Deposit(amount int) (interface{}, interface{}) {
	b.total_balance += amount
	file, err := os.Create("bank.txt")
	if err != nil {
		fmt.Println("Error creating file\n", err)
		return err, false
	}
	data := fmt.Sprintf("%.2f Amount has been deposited\n", amount)
	file.WriteString(data)
	return amount, "deposited"
}

func (b *PersonDomestic) Withdrawl(amount int) (interface{}, interface{}) {
	if amount > b.total_balance {
		fmt.Println("Not Enough Balance")
		return amount, false
	}
	b.total_balance -= amount
	file, err := os.Create("bank.txt")
	if err != nil {
		fmt.Println("error creating file ")
		return err, false
	}
	data := fmt.Sprintf("%.2f Amount has been withdrawl\n", amount)
	file.WriteString(data)
	return amount, "withdrawl"
}

func CreateMerchantAccount() (interface{}, interface{}) {
	var name string
	var buisness_name string
	var phone string
	var address string
	var revenue int
	fmt.Print("\nEnter Your Full name\n")
	fmt.Scan(&name)
	fmt.Println("\nEnter your city name\n")
	fmt.Scan(&address)
	fmt.Println("\nEnter your phone num\n")
	fmt.Scan(&phone)
	fmt.Println("\nEnter your buisness name\n")
	fmt.Scan(&buisness_name)
	fmt.Println("\nEnter your buisness annual revenue\n")
	fmt.Scan(&revenue)

	if revenue < 500000 {
		fmt.Println("\nSince your revenue is short we can't create your merchant account\n")
		return errors.New("Error :Your account can't be created "), false
	}
	rand.Seed(time.Now().UnixNano()) // Seed with current time
	user_id := rand.Intn(1000000)

	mer_address := AddressMerchant{
		city:  address,
		phone: phone,
	}

	userDetail := PersonMerchant{
		id:              user_id,
		name:            name,
		buisnessname:    buisness_name,
		buisnessrevenue: revenue,
		address:         mer_address,
		total_balance:   0,
	}

	fmt.Println("\nYour Merchant Account has been created\n")
	fmt.Println("Your Account id is :", userDetail.id)
	return userDetail.id, userDetail
}

func (b *PersonMerchant) Deposit(amount int) (interface{}, interface{}) {
	b.total_balance += amount
	file, err := os.Create("bank.txt")
	if err != nil {
		fmt.Println("Error creating file", err)
		return err, false
	}
	data := fmt.Sprintf("%.2f Amount has been deposited\n", amount)
	file.WriteString(data)
	return amount, "deposited"
}

func (b *PersonMerchant) Withdrawl(amount int) (interface{}, interface{}) {
	if amount > b.total_balance {
		fmt.Println("\nNot Enough Balance\n")
		return amount, false
	}
	b.total_balance -= amount
	file, err := os.Create("bank.txt")
	if err != nil {
		fmt.Println("error creating file ")
		return err, false
	}
	data := fmt.Sprintf("%.2f Amount has been withdrawl\n", amount)
	file.WriteString(data)
	return amount, "withdrawl"
}
