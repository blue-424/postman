package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/myhandler", myHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// محتوای درخواست را بخوانید
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "خطا در پردازش فرم", http.StatusBadRequest)
			return
		}
		{
			name := r.FormValue("name")
			lastname := r.FormValue("last name")
			address := r.FormValue("address")
			phoneNumber := r.FormValue("phone number")

			fmt.Fprintf(w, "Name = %s\n", name)
			fmt.Fprintf(w, "Last name = %s\n", lastname)
			fmt.Fprintf(w, "Address = %s\n", address)
			fmt.Fprintf(w, "Phone number = %s\n", phoneNumber)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")


		fmt.Fprintf(w, "Name : %s\nLast name: %s\nAddress : %s\nPhone number: %s\n", name, lastname, address, phoneNumber)
	} else {
		http.Error(w, "تنها درخواست‌های POST پشتیبانی می‌شوند", http.StatusMethodNotAllowed)
	}
}
