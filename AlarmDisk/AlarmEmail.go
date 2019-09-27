package AlarmDisk

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
	"strconv"
)

func AlarmDiskSizeEmail(kv map[string]float64) bool {

	if len(kv) == 0 {
		return true
	}

	var body string
	for server,size := range kv {
		body += server+strconv.FormatFloat(size, 'E', -1, 64)+"\n"
	}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", "XXX@XXX.com", "dev")
	m.SetAddressHeader("To", "XXX@XXX.com", "XXX")
	m.SetAddressHeader("To", "XXX@XXX.com", "XXX")
	m.SetHeader("Subject", "磁盘报警")
	m.SetBody("text/html", "<h1>"+body+"</h1>")

	d := gomail.NewPlainDialer("smtp.exmail.qq.com", 465, "XXX@XXX.com", "XXXXX")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return false
	}
	return true
}
