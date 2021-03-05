package zabbix

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/cavaliercoder/go-zabbix"
	cnf "github.com/koscom-cloud/zabbix-exporter-3000/config"
)

var Session, err = Connect()
var Query *zabbix.Request

func Connect() (*zabbix.Session, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: cnf.Cnf.SslSkip}}}

	cache := zabbix.NewSessionFileCache().SetFilePath("./zabbix_session")
	session, err := zabbix.CreateClient(cnf.Cnf.Server).
		WithCache(cache).
		WithHTTPClient(client).
		WithCredentials(cnf.Cnf.User, cnf.Cnf.Password).
		Connect()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	version, err := session.GetVersion()

	if err != nil {
		panic(err)
	}

	authToken := session.AuthToken()
	sToken := strings.Split(authToken, "")
	log.Print("Auth: ", sToken[1], sToken[2], sToken[3], sToken[4], sToken[5], sToken[6])
	strRequestWithAuth := strings.Replace(cnf.Cnf.Query, "%auth-token%", authToken, -1)

	// fmt.Print(cnf.Cnf.Query)
	err = json.Unmarshal([]byte(strRequestWithAuth), &Query)
	if err != nil {
		log.Print("ERROR While convert request to JSON: ", err)
	}

	log.Print("Connected to Zabbix API v", version)
	return session, err
}
