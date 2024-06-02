# hu-design-project

Request kısmını ve devices kısmına bakmıyorum şu anlık sadece alerts kısmına bakacağım. Ek olarak dashboard kısmına bakacağım. Anladığım şeyleri yazıyorum sana yanlışım var ise düzelt:

Alerts:
- Alert ekleyebileceğimiz POST /alerts endpointi lazım.
- Alertleri alabileceğin GET /alerts endpointi lazım.
- Tek bir tane alert alabileceğin GET /alerts/{id} endpointi lazım.
- Whitelist veya blackliste eklendiğinde alertü silen DELETE /alerts/{id} endpointi lazım.
  Whitelist ve Blacklist:
- POST /whitelist ve POST /blacklist lazım
- GET /whitelist ve GET /blacklist lazım


{
"userId": "baverkacar@gmail.com",
"oldPassword": "",
"newPassword": ""
}

/users/change-password


/forget-password