## path traversal?

Path traversal is also known as directory traversal. These vulnerabilities can 
enable an attacker to read arbitrary files on the server that is running an application. 
This might include:

1. Application code and data.
2. Credentials for back-end systems.
3. Sensitive operating system files.

** 

In some cases, an attacker might be able to write to arbitrary files on the server, allowing 
them to modify application data or behavior, and ultimately take full control of the server.

### usage

Here is how to use this program.

```bash
    go run main.go payload.json
```

Here is the syntax of the `payload.json` file.

```json
{
  "host": "https://0000000000000.web-security-academy.net/image",
  "payload": "../../../etc/passwd",
  "session": "sdsedfdfdsfdsfsd",
  "method": "GET"
}
```

Get session from the browser once the lab is launched.
Get host from the lab link once you launch it and remember to append `/image` since thats the endpoint you will be exploiting
Payload is the file path
All lab endpoints use `GET` requests.
