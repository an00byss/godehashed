
# Godehashed
A golang tool that uses the dehashed.com API to search for compromised assets. Results can then be compiled into a CSV for further analysis.

## Screenshot

![](screenshot.jpg)

## Dehashed API

You must supply the tool an api key. See apikey_template.txt for example.

## Installation

To install the tool in CLI run the following command. Your $GOPATH must already be set.

```go get https://github.com/an00byss/godehashed```

## Usage
./godehashed -s email -i apikey.txt -e SOMEDOMAIN -o leaks.csv

```
Godehashed Usage():
-e string  
Email we are searching for  
-i string  
Name of apikey to import.  
-n string  
Name we are searching for.  
-o string  
Outfile file name, will output in CSV Format.  
-p int  
Phone number we are searching for  
-s string  
Specify what we are searching for: "name", "email", "phone" or "username". Then add corresponding switch.  
-u string  
Username we are searching for
-l list  
Search a list of emails.
```

|       Search Term         |Examples                         |
|----------------|-------------------------------|
|Email|`godehashed -s email -i apikey.txt -e SOMEDOMAIN.com -o leaks.csv`            |
|Username          |`godehashed -s username -i apikey.txt -u USERNAME -o leaks.csv`            |
|Name          |`godehashed -s name -i apikey.txt -n "Name" -o leaks.csv`|
|Phone          |`godehashed -s phone -i apikey.txt -p "phonenumber" -o leaks.csv`|
|list          |`godehashed -e email -i apikey.txt -l list.txt -o leaks.csv`|

## Notice

```
[!] Legal disclaimer: Usage of godehashed for attacking targets without
prior mutual consent is illegal. It is the end user's responsibility
to obey all applicable local, state and federal laws. Developers assume
no liability and are not responsible for any misuse or damage caused.
```


