# go-jplaw-xml

Go struct for [Japanese Standard Law XML Schema (法令標準XMLスキーマ)][xmldoc]

## Usage

```sh
go get github.com/ngs/go-jplaw-xml
```

```go
file, err := os.Open("path/to/law.xml")
if err != nil {
	log.Fatalf("Failed to open file: %v", err)
}
defer file.Close()

data, err := io.ReadAll(file)
if err != nil {
	log.Fatalf("Failed to read file: %v", err)
}

var schema jplaw.Law
err = xml.Unmarshal(data, &schema)

if err != nil {
	log.Fatalf("Failed to parse schema: %v", err)
}

log.Println("LawNum:", schema.LawNum)
log.Println("LawTitle:", schema.LawBody.LawTitle.CharData)
```

## Author

[Atushi Nagase]

## License

Copyright &copy; 2025 [Atushi Nagase]. All rights reserved.

[Atushi Nagase]: https://ngs.io/
[xmldoc]: https://laws.e-gov.go.jp/docs/law-data-basic/419a603-xml-schema-for-japanese-law/

