# go-trans

go-trans is a package that can translate your static texts easily

yaml file exist in /app/descriptions/trans.yml path:
```yaml
poem:
  fa: 'اومدم خونتون خونه نبودی'
  en: 'I came to your home but you were not there'
payment:
  fa: 'مبلغ قابل پرداخت :price'
  en: 'Payable amount :amount'
```

Usage:

```go
type unmarshaler struct{}

func (unmarshaler) Unmarshal(d []byte, v *map[string]map[string]string) error {
	err := yaml.Unmarshal(d, v)
	return err
}

func main() {
	err := go_trans.Initialize("/app/descriptions", unmarshaler{})
	if err != nil {
	    log.Fatal(err)
	}

	//Without replace will return اومدم خونتون خونه نبودی
	go_trans.Trans("trans.poem", "fa")

	//With replace will return Payable amount 2000
	go_trans.Trans("trans.payment", "en", go_trans.R{"price":"2000"})
}
```