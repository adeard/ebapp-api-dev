package job

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

func MainJob() {
	c := cron.New()

	schedule := os.Getenv("EMAIL_CRON_SCHEDULE")
	if schedule == "" {
		schedule = "0 6 * * *" // default setiap hari jam 06:00
	}
	_, err := c.AddFunc(schedule, func() {
		log.Println("=== Cron ProsesSendEmail Dimulai ===")

		if err := ProsesSendEmail(); err != nil {
			log.Printf("ERROR ProsesSendEmail: %v", err)
			return
		}

		log.Println("=== Cron ProsesSendEmail Selesai ===")
	})

	if err != nil {
		log.Printf("Gagal menambahkan cron: %v", err)
		return
	}

	c.Start()

	log.Println("Cron Job berhasil dijalankan")
}

func ProsesSendEmail() error {

	authToken, err := GetAuthToken()
	if err != nil {
		return err
	}

	client := &http.Client{}

	listStatus := []string{"Waiting", "OnProgress"}
	bodyListStatus, _ := json.Marshal(listStatus)

	urlDoc := fmt.Sprintf(
		"%swf_helper/get_doc_by_status?doc_alias=eBAPP",
		os.Getenv("SERVER_URL_WR"),
	)

	req, err := http.NewRequest(http.MethodPost, urlDoc, bytes.NewBuffer(bodyListStatus))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authenticationToken", authToken)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var docs GetDocResponse

	if err := json.NewDecoder(resp.Body).Decode(&docs); err != nil {
		return err
	}

	if !docs.Result {
		return fmt.Errorf("gagal mengambil document")
	}

	dummyEmail := strings.TrimSpace(os.Getenv("EMAIL_DUMMY"))
	for _, doc := range docs.Objek {

		split := strings.Split(doc.Description, "|")
		if len(split) == 0 {
			continue
		}

		listMasterValueID := []string{split[0]}

		body, _ := json.Marshal(listMasterValueID)

		emails := []string{}

		if strings.TrimSpace(doc.PotensialActivityOwner) == "bapp.kontraktorvendor" {
			continue
		}

		if strings.TrimSpace(doc.ActivityOwner) != "" {

			urlUser := fmt.Sprintf(
				"%suser/email_by_username?username=%s",
				os.Getenv("SERVER_URL_UM"),
				url.QueryEscape(doc.ActivityOwner),
			)

			reqUser, err := http.NewRequest(http.MethodGet, urlUser, nil)
			if err != nil {
				log.Println(err)
				continue
			}

			reqUser.Header.Set("Accept", "application/json")
			reqUser.Header.Set("authenticationToken", authToken)

			respUser, err := client.Do(reqUser)
			if err != nil {
				log.Println(err)
				continue
			}
			defer respUser.Body.Close()

			var userResp UserEmailByUsernameResponse
			if err := json.NewDecoder(respUser.Body).Decode(&userResp); err != nil {
				log.Println(err)
				continue
			}

			if !userResp.Result || strings.TrimSpace(userResp.Objek) == "" {
				continue
			}

			emails = append(emails, userResp.Objek)

		} else {

			urlUser := fmt.Sprintf(
				"%suser/list_email_by_role_master_value_id?application_role=%s",
				os.Getenv("SERVER_URL_UM"),
				url.QueryEscape(doc.PotensialActivityOwner),
			)

			reqUser, err := http.NewRequest(
				http.MethodPost,
				urlUser,
				bytes.NewBuffer(body),
			)
			if err != nil {
				log.Println(err)
				continue
			}

			reqUser.Header.Set("Content-Type", "application/json")
			reqUser.Header.Set("authenticationToken", authToken)

			respUser, err := client.Do(reqUser)
			if err != nil {
				log.Println(err)
				continue
			}
			defer respUser.Body.Close()

			var users UserEmailResponse
			if err := json.NewDecoder(respUser.Body).Decode(&users); err != nil {
				log.Println(err)
				continue
			}

			if !users.Result {
				continue
			}

			for _, user := range users.Objek {
				if strings.TrimSpace(user.Email) != "" {
					emails = append(emails, user.Email)
				}
			}
		}

		if len(emails) == 0 {
			continue
		}

		for _, user := range emails {

			if user == "" {
				continue
			}

			// Default kirim ke email user
			sendTo := user
			projectID := ""
			estate := ""
			projectName := ""
			cc := []string{}

			if len(split) > 0 {
				projectID = split[0]
			}
			if len(split) > 1 {
				estate = split[1]
			}
			if len(split) > 2 {
				projectName = split[2]
			}
			if len(split) > 3 && strings.TrimSpace(split[3]) != "" {
				for _, email := range strings.Split(split[3], ",") {
					email = strings.TrimSpace(email)
					if email != "" {
						cc = append(cc, email)
					}
				}
			}

			// Jika EMAIL_DUMMY diisi, override tujuan email
			if dummyEmail != "" {
				sendTo = dummyEmail
				cc = []string{dummyEmail}
			}

			approvalLink := fmt.Sprintf(
				"https://appcore.indoagri.co.id/ebapp/#/view-progress/%s/%s",
				estate,
				doc.DocumentID,
			)

			if strings.Contains(strings.ToUpper(doc.DocumentID), "/CPP") {
				approvalLink = fmt.Sprintf(
					"https://appcore.indoagri.co.id/ebapp/#/view-cpp/%s/%s",
					estate,
					doc.DocumentID,
				)
			}
			body := fmt.Sprintf(`
			<html>
			<body style="font-family:Arial,Helvetica,sans-serif;font-size:14px;color:#000">

			<p>Dear Bp/Ibu,</p>

			<p>
			Berikut informasi document project <b>eBAPP</b> yang masih menunggu persetujuan :
			</p>

			<table cellpadding="6" cellspacing="0">
				<tr>
					<td width="120"><b>Document Id</b></td>
					<td>: %s</td>
				</tr>
				<tr>
					<td><b>Project ID</b></td>
					<td>: %s</td>
				</tr>
				<tr>
					<td><b>Nama Project</b></td>
					<td>: %s</td>
				</tr>
				<tr>
					<td><b>Link Approval</b></td>
					<td>: <a href="%s">%s</a></td>
				</tr>
			</table>

			<br><br>

			<p>
			Untuk informasi lebih lanjut dapat menghubungi
			<a href="mailto:eBAPP.support@simp.co.id">
			eBAPP.support@simp.co.id
			</a>
			</p>

			<p>
			<i>
			Email ini merupakan email otomatis yang dikirim dari sistem.
			Harap tidak membalas email ini.
			</i>
			</p>

			<br>

			<p>
			Terima kasih.<br>
			Regards,
			</p>

			<p>
			Admin Project
			</p>

			</body>
			</html>
			`,
				doc.DocumentID,
				projectID,
				projectName,
				approvalLink,
				approvalLink,
			)

			err := SendEmail(
				sendTo,
				cc,
				fmt.Sprintf(`eBAPP- Reminder Approval %s`, doc.DocumentID),
				body,
			)

			if err != nil {
				log.Printf("Gagal kirim email ke %s : %v",
					sendTo,
					err,
				)
				continue
			}

			log.Printf("Berhasil kirim email ke %s",
				sendTo,
			)
		}
	}

	return nil
}

func GetAuthToken() (string, error) {

	url := fmt.Sprintf("%suser/login", os.Getenv("SERVER_URL_UM"))

	reqBody := map[string]string{
		"user_name": "user_bapp",
		"password":  "P@ssw0rd",
		"apps":      "BAPP",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result LoginResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if !result.Result {
		return "", fmt.Errorf("login gagal: %s", result.Message)
	}

	return result.Datas, nil
}

func SendEmail(to string, cc []string, subject string, body string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("SMTP_FROM")

	addr := host + ":" + port

	client, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer client.Quit()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	if ok, _ := client.Extension("STARTTLS"); ok {
		if err := client.StartTLS(tlsConfig); err != nil {
			return err
		}
	}

	if strings.TrimSpace(pass) != "" {
		auth := smtp.PlainAuth("", user, pass, host)

		if ok, _ := client.Extension("AUTH"); ok {
			if err := client.Auth(auth); err != nil {
				return err
			}
		}
	}
	if err := client.Mail(from); err != nil {
		return err
	}

	if err := client.Rcpt(to); err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	message := []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Cc: " + strings.Join(cc, ",") + "\r\n" +
			"Subject: " + subject + " - " + time.Now().Format("02 Jan 2006 15:04:05") + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
			body,
	)

	_, err = w.Write(message)
	if err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return client.Quit()
}
