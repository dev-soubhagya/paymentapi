package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//user can send amount to merchant
func UserToMerchant(httpres http.ResponseWriter, req *http.Request) {
	var userdata User
	var merchantdata Merchant
	var request Request
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in parsing", err)
		httpres.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("Got request ", request)
	err = db.QueryRow("SELECT * FROM user WHERE userid = ?", request.Userid).Scan(&userdata.Userid, &userdata.Balance, &userdata.Status)
	ErrorCheck(err)
	fmt.Println("User details ", userdata)
	if userdata.Balance > 0 && request.Amount <= userdata.Balance {
		balance := userdata.Balance - request.Amount
		//user debit
		stmt, e := db.Prepare("update user SET balance=? where userid=?")
		ErrorCheck(e)
		fmt.Println("Balance ", balance, "userid ", userdata.Userid)
		res, e := stmt.Exec(balance, userdata.Userid)
		ErrorCheck(e)
		a, e := res.RowsAffected()
		ErrorCheck(e)
		fmt.Println(a)
		//merchant credit
		err := db.QueryRow("SELECT * FROM merchant WHERE merchantid = ?", request.Merchantid).Scan(&merchantdata.Merchantid, &merchantdata.Balance, &merchantdata.Status)
		ErrorCheck(err)
		fmt.Println("Merchant details ", merchantdata)
		merchantamount := merchantdata.Balance + request.Amount
		stmt1, e := db.Prepare("UPDATE merchant SET balance=? WHERE merchantid=?")
		ErrorCheck(e)
		fmt.Println("merchant AMOUNT AFTER CREDIT ADD ", merchantamount, "merchantid ", merchantdata.Merchantid)
		res1, e := stmt1.Exec(merchantamount, merchantdata.Merchantid)
		ErrorCheck(e)
		b, e := res1.RowsAffected()
		ErrorCheck(e)
		fmt.Println(b)
		fmt.Println("USER TO MERCHANT TRANSACTION SUCCESS ")
		//insert in user transaction status
		transdata := Response{TxnId: GenerateId(), Status: "SUCCESS", Status_desc: "Successfully Credited in Merchant"}
		statement, e := db.Prepare("INSERT INTO user_transaction_status(txnid,userid,amount,balance,status,type) VALUES (?,?,?,?,?,?)")
		ErrorCheck(e)
		res2, e := statement.Exec(transdata.TxnId, userdata.Userid, request.Amount, balance, transdata.Status, "DEBIT")
		ErrorCheck(e)
		c, e := res2.RowsAffected()
		ErrorCheck(e)
		fmt.Println(c)
		statement1, e := db.Prepare("INSERT INTO merchant_transaction_status(txnid,merchantid,amount,balance,status,type) VALUES (?,?,?,?,?,?)")
		ErrorCheck(e)
		res3, e := statement1.Exec(transdata.TxnId, merchantdata.Merchantid, request.Amount, merchantamount, transdata.Status, "CREDIT")
		ErrorCheck(e)
		d, e := res3.RowsAffected()
		ErrorCheck(e)
		fmt.Println(d)
		bytedata, _ := json.Marshal(transdata)
		httpres.Header().Set("Content-Type", "application/json")
		httpres.Write(bytedata)
	} else {
		fmt.Println("Balance Not exist ")
	}
}

//merchant can send amount to user
func MerchantToUser(httpres http.ResponseWriter, req *http.Request) {
	var userdata User
	var merchantdata Merchant
	var request Request
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in parsing", err)
		httpres.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("Got request ", request)
	err = db.QueryRow("SELECT * FROM merchant WHERE merchantid = ?", request.Merchantid).Scan(&merchantdata.Merchantid, &merchantdata.Balance, &merchantdata.Status)
	ErrorCheck(err)
	fmt.Println("Merchant details ", merchantdata)
	if merchantdata.Balance > 0 && request.Amount <= merchantdata.Balance {
		merchantamount := merchantdata.Balance - request.Amount
		//merchant debit
		stmt1, e := db.Prepare("UPDATE merchant SET balance=? WHERE merchantid=?")
		ErrorCheck(e)
		fmt.Println("Debit balance ", merchantamount, "merchantid ", merchantdata.Merchantid)
		res1, e := stmt1.Exec(merchantamount, merchantdata.Merchantid)
		ErrorCheck(e)
		b, e := res1.RowsAffected()
		ErrorCheck(e)
		fmt.Println(b)
		//user credit
		err = db.QueryRow("SELECT * FROM user WHERE userid = ?", request.Userid).Scan(&userdata.Userid, &userdata.Balance, &userdata.Status)
		ErrorCheck(err)
		fmt.Println("User details ", userdata)
		useramount := userdata.Balance + request.Amount
		stmt, e := db.Prepare("UPDATE user SET balance=? WHERE userid=?")
		ErrorCheck(e)
		fmt.Println("Balance ", useramount, "userid ", userdata.Userid)
		res, e := stmt.Exec(useramount, userdata.Userid)
		ErrorCheck(e)
		a, e := res.RowsAffected()
		ErrorCheck(e)
		fmt.Println(a)
		fmt.Println("USER TO MERCHANT TRANSACTION SUCCESS ")
		//insert in user transaction status
		transdata := Response{TxnId: GenerateId(), Status: "SUCCESS", Status_desc: "Successfully Refunded"}
		statement, e := db.Prepare("INSERT INTO user_transaction_status(txnid,userid,amount,balance,status,type) VALUES (?,?,?,?,?,?)")
		ErrorCheck(e)
		res2, e := statement.Exec(transdata.TxnId, userdata.Userid, request.Amount, useramount, transdata.Status, "CREDIT")
		ErrorCheck(e)
		c, e := res2.RowsAffected()
		ErrorCheck(e)
		fmt.Println(c)
		statement1, e := db.Prepare("INSERT INTO merchant_transaction_status(txnid,merchantid,amount,balance,status,type) VALUES (?,?,?,?,?,?)")
		ErrorCheck(e)
		res3, e := statement1.Exec(transdata.TxnId, merchantdata.Merchantid, request.Amount, merchantamount, transdata.Status, "DEBIT")
		ErrorCheck(e)
		d, e := res3.RowsAffected()
		ErrorCheck(e)
		fmt.Println(d)
		bytedata, _ := json.Marshal(transdata)
		httpres.Header().Set("Content-Type", "application/json")
		httpres.Write(bytedata)
	} else {
		fmt.Println("Balance Not exist ")
	}
}

//merchant can withdraw amount from his account(merchantid)
func MerchantWithraw(httpres http.ResponseWriter, req *http.Request) {
	var merchantdata Merchant
	var request Request
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in parsing", err)
		httpres.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("Got request ", request)
	err = db.QueryRow("SELECT * FROM merchant WHERE merchantid = ?", request.Merchantid).Scan(&merchantdata.Merchantid, &merchantdata.Balance, &merchantdata.Status)
	ErrorCheck(err)
	fmt.Println("Merchant details ", merchantdata)
	if merchantdata.Balance > 0 && request.Amount <= merchantdata.Balance {
		merchantamount := merchantdata.Balance - request.Amount
		//merchant debit
		stmt1, e := db.Prepare("UPDATE merchant SET balance=? WHERE merchantid=?")
		ErrorCheck(e)
		fmt.Println("Debit balance ", merchantamount, "merchantid ", merchantdata.Merchantid)
		res1, e := stmt1.Exec(merchantamount, merchantdata.Merchantid)
		ErrorCheck(e)
		b, e := res1.RowsAffected()
		ErrorCheck(e)
		fmt.Println(b)
		transdata := Response{TxnId: GenerateId(), Status: "SUCCESS", Status_desc: "Withraw SuccessfullY"}
		statement1, e := db.Prepare("INSERT INTO merchant_transaction_status(txnid,merchantid,amount,balance,status,type) VALUES (?,?,?,?,?,?)")
		ErrorCheck(e)
		res3, e := statement1.Exec(transdata.TxnId, merchantdata.Merchantid, request.Amount, merchantamount, transdata.Status, "MERCHANT_WITHRAW")
		ErrorCheck(e)
		d, e := res3.RowsAffected()
		ErrorCheck(e)
		fmt.Println(d)
		bytedata, _ := json.Marshal(transdata)
		httpres.Header().Set("Content-Type", "application/json")
		httpres.Write(bytedata)
	} else {
		fmt.Println("Amount cant be withraw ")
	}
}

//merchant can all transaction related to his account(merchantid)
func MerchantTransactionCheck(httpres http.ResponseWriter, req *http.Request) {
	var request Request
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in parsing", err)
		httpres.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("Got request ", request)
	stmt, err := db.Prepare("select * from merchant_transaction_status WHERE merchantid=?")
	ErrorCheck(err)
	defer stmt.Close()
	rows, err := stmt.Query(request.Merchantid)
	ErrorCheck(err)
	defer rows.Close()
	allmerchantTxn := []Merchant_transaction_status{}
	for rows.Next() {
		var merchantTxn Merchant_transaction_status
		if err := rows.Scan(&merchantTxn.Txnid, &merchantTxn.Merchantid, &merchantTxn.Amount, &merchantTxn.Balance, &merchantTxn.Status, &merchantTxn.Type); err != nil {
			ErrorCheck(err)
		}
		allmerchantTxn = append(allmerchantTxn, merchantTxn)
	}
	if err := rows.Err(); err != nil {
		ErrorCheck(err)
	}
	fmt.Println("All Merchant Txn Data====>>>>>> ", allmerchantTxn)
	var data1 []byte
	if data1, err = json.Marshal(allmerchantTxn); err != nil {
		ErrorCheck(err)
	}
	httpres.Header().Set("Content-Type", "application/json")
	httpres.Write(data1)
}
