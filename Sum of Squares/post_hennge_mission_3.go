package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func HOTP(secret string, message int, digits int) string {
	// HTOP:
	// K is the shared key
	// C is the counter value(message)
	// digits control the response length

	secret_byte := []byte(secret)

	message_byte := make([]byte, 8)
	binary.BigEndian.PutUint64(message_byte, uint64(message))

	hmac512 := hmac.New(sha512.New, secret_byte)
	hmac512.Write((message_byte))
	hash := hmac512.Sum(nil)

	offset := hash[63] & 15

	var password uint32

	r := bytes.NewReader(hash[offset : offset+4])
	_ = binary.Read(r, binary.BigEndian, &password)

	return fmt.Sprintf("%010d", password)
}

func TOTP(secret string, digits int, timeref int64, timestep int64) string {
	// TOTP, time-based variant of HOTP
	// digits control the response length
	// the C (message) in HOTP is replaced by ( (currentTime - timeref) / timestep )
	message := (time.Now().Unix() - timeref) / timestep

	return HOTP(secret, int(message), digits)
}

func main() {
	url := "https://api.challenge.hennge.com/challenges/003"
	content_type := "application/json"
	userid := "khondokar.hash@gmail.com"
	secret_suffix := "HENNGECHALLENGE003"
	git_gist := "https://gist.github.com/peeyalk/fb7820c38a6077a79877d129e1ae33e9"
	shared_secret := userid + secret_suffix
	password_digit := 10
	timeref := 0
	timestep := 30

	password := TOTP(shared_secret, password_digit, int64(timeref), int64(timestep))

	fmt.Println(password)

	fmt.Println("URL:>", url)

	var data = []byte(`
		{
			"github_url": "` + git_gist + `",
			"contact_email": "` + userid + `"
		}
	`)

	fmt.Println(string(data))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", content_type)
	req.Header.Add("Authorization", "Basic "+basicAuth(userid, password))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
