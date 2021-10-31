package main
// eu ainda tô estudando Go XD
import (
  "fmt"
  gomail "gopkg.in/mail.v2"
)
func main() {
  var gmail,senha,body,title,gmail_vitima string
  var valor int
  fmt.Printf("Ative a opção de Apps menos seguros antes de utilizar o script.\nDigite seu endereço de Email : ")
  fmt.Scan(&gmail)
  fmt.Printf("Digite sua senha : ")
  fmt.Scan(&senha)
  fmt.Printf("Digite o titulo do seu Email : ")
  fmt.Scan(&title)
  fmt.Printf("Digite a mensagem que quer enviar : ")
  fmt.Scan(&body)
  fmt.Printf("Digite o email da vítima : ")
  fmt.Scan(&gmail_vitima)
  fmt.Printf("Digite a quantidade de Emails que você quer enviar : ")
  fmt.Scan(&valor)
  
  for number:=valor;number>=0;number--{
    Gmail:=gomail.NewMessage()
    Gmail.SetHeader("From", gmail)
    Gmail.SetHeader("To", gmail_vitima)
    Gmail.SetHeader("Subject",title)
    Gmail.SetHeader("text/plain", body)
    Gmail_ver=gomail.NewDailer("smtp.gmail.com",587,gmail,senha)
    if erro:=Gmail_ver.DialAndSender(Gmail); erro!=nil{
      fmt.Printf(erro);panic(erro)
    }
  }
}
