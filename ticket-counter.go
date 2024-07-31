//package main
//
//import (
//	"bytes"
//	"fmt"
//	"net/smtp"
//)
//
//func main()  {
//	arr := []int{1,2,3,4,5,6,7,8,9}
//	fmt.Println(distributeTicket(arr, 3))
//}
//
//func distributeTicket(seats []int, k int) int {
//	var frontRo = true
//	var last = false
//
//
//
//	return 0
//}


package main

import (
"bytes"
"fmt"
"html/template"
"net/smtp"
)

var auth smtp.Auth

func main() {
	auth = smtp.PlainAuth(
		"",
		"ayushdadhich900@gmail.com",
		"sqij xqxq vwgx zegj",
		"smtp.gmail.com",
	)
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Dhanush",
		URL:  "http://geektrust.in",
	}
	r := NewRequest([]string{"shekharpandey93@gmail.com"}, "Hello Junk!", "Hello, World!")
	//err := r.ParseTemplate("template.html", templateData)
	if err := r.ParseTemplate("template.html", templateData); err == nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}

}

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "ayushdadhich900@gmail.com", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
