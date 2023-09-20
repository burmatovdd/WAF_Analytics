package resource

import (
	"database/sql"
	"time"
)

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserAuth struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Access   string `json:"access"`
}

type User struct {
	ID    sql.NullInt32  `json:"id"`
	Email sql.NullString `json:"email"`
	FIO   sql.NullString `json:"FIO"`
}

type RequestStatistic struct {
	AllServers      int `json:"allServers"`
	ErrorServers    int `json:"errorServers"`
	WorkingServers  int `json:"workingServers"`
	WithWaf         int `json:"withWaf"`
	AllCertificates int `json:"allCertificates"`
	OkCertificates  int `json:"okCertificates"`
}

type ResponseStatistic struct {
	ID              int
	Date            time.Time
	AllServers      int
	ErrorServers    int
	WorkServers     int
	WithWaf         int
	Possible        float64
	WafProcPossible float64
	WafProc         float64
	WithKas         int
	WafAndKas       int
	WafAndKasProc   float64
	AllCertificate  int
	OkCertificate   int
}

type Weeks struct {
	Last    Week
	Current Week
}

type Week struct {
	Monday time.Time
	Friday time.Time
}

type WeeksStatistic struct {
	Last    WeekStatistic `json:"last"`
	Current WeekStatistic `json:"current"`
}

type WeekStatistic struct {
	NoResolve      int                     `json:"no_resolve"`
	NewWaf         int                     `json:"new_waf"`
	NoResResource  []WeekStatisticResource `json:"no_res_resource"`
	NewWafResource []WeekStatisticResource `json:"new_waf_resource"`
}

type WeekStatisticResource struct {
	Resource string `json:"resource"`
	Date     string `json:"date"`
}

type URL struct {
	Url   string `json:"url"`
	Email string `json:"email"`
	Owner string `json:"owner"`
}

type ResourceTable struct {
	ID        sql.NullInt32  `json:"ID"`
	NameURL   sql.NullString `json:"NameURL"`
	IpFirst   sql.NullString `json:"IpFirst"`
	IpNow     sql.NullString `json:"IpNow"`
	DateFirst sql.NullTime   `json:"DateFirst"`
	Status    sql.NullString `json:"Status"`
	DateNoRes sql.NullTime   `json:"DateNoRes"`
	WafDate   sql.NullTime   `json:"WafDate"`
	WafIp     sql.NullString `json:"WafIp"`
}

type UrlTable struct {
	ID         sql.NullInt32
	URL        sql.NullString
	IP         sql.NullString
	Err        sql.NullString
	Waf        sql.NullString
	IDUser     sql.NullInt64
	IDOwner    sql.NullInt64
	CommonName sql.NullString
	Issuer     sql.NullString
	EndDate    sql.NullString
}

type Owner struct {
	ID        sql.NullInt32  `json:"id"`
	FullName  sql.NullString `json:"nameOwner"`
	ShortName sql.NullString `json:"shortName"`
}

type CheckDataResult struct {
	UserID  sql.NullInt32
	OwnerId sql.NullInt32
	Result  bool
}

type CheckResource struct {
	URL     string         `json:"URL"`
	IP      string         `json:"IP"`
	Status  bool           `json:"Status"`
	WAF     bool           `json:"WAF"`
	SSL     UrlCertificate `json:"SSL"`
	DateEnd string         `json:"DateEnd"`
	Email   string         `json:"Email"`
	FIO     string         `json:"FIO"`
	Owner   string         `json:"Owner"`
}

type UpdateData struct {
	Url   string `json:"url"`
	Email string `json:"email"`
}

type GeneralStat struct {
	Resources          int `json:"resources"`
	DeactivateResource int `json:"deactivateResource"`
	Owners             int `json:"owners"`
	Waf                int `json:"waf"`
}

type CertificateInfo struct {
	Current []Certificate `json:"current"`
	Next    []Certificate `json:"next"`
}

type Certificate struct {
	Resource string `json:"resource"`
	Date     string `json:"date"`
}

type Months struct {
	Current string
	Next    string
}

type ResolveInfo struct {
	Ip        string     `json:"ip"`
	Status    string     `json:"status"`
	DateNoRes *time.Time `json:"dateNoRes"`
	WafDate   *time.Time `json:"wafDate"`
	Waf       string     `json:"waf"`
	WafIp     *string    `json:"wafIp"`
	NameUrl   string     `json:"nameurl"`
}

type UrlCertificate struct {
	CommonName string `json:"common_name"`
	Issuer     string `json:"issuer"`
	DateCert   string `json:"date_cert"`
}

type AddResourceCollection struct {
	Resolve     ResolveInfo    `json:"resolve"`
	Certificate UrlCertificate `json:"certificate"`
}
