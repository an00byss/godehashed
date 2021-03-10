# Godehashed
A golang tool that uses the dehashed API to search for compromised assets.

## Screenshot

![](screenshot.jpeg)

## Dehashed API

You must supply the tool an api key. See apikey_template.txt for example.
Use with switch ./godehashed -i apikey.txt

## Usage

```
Usage():
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
Specify what we are searching for: name, email or username. Then Add corrisponding switch.  
-u string  
Username we are searching for
```

## Notice

```
[!] Legal disclaimer: Usage of godehashed for attacking targets without
prior mutual consent is illegal. It is the end user's responsibility
to obey all applicable local, state and federal laws. Developers assume
no liability and are not responsible for any misuse or damage caused.
```


