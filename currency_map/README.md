# 💱 Currency Normalization Map Package

Go package providing comprehensive currency name-to-ISO code mapping for financial data processing.

![Currency Conversion](https://via.placeholder.com/800x200.png?text=Multi-language+Support→ISO+Standardization→Financial+Data+Consistency)

## ✨ Features

- **ISO 4217 Compliance**: Standard 3-letter currency codes
- **Multi-language Support**: English, Portuguese, and common variations
- **Common Aliases**: Handles abbreviations, typos, and regional terms
- **Extensible Structure**: Easy to add new currency mappings
- **Case Insensitive**: Use any text case for lookups (handled by consumer implementation)

## 🚀 Installation

```bash
go get github.com/samuelrms/translate-currency/currency_map
```

## 📦 Usage

Import the package and access the map directly:

```go
import "github.com/samuelrms/translate-currency/currency_map"

func main() {
    // Normalize currency input
    rawInput := "dólar australiano"
    isoCode := currency_map.CurrencyMap[strings.ToUpper(rawInput)]
    
    // Handle missing entries
    if isoCode == "" {
        isoCode = "BRL" // Default fallback
    }
}
```

## 🌍 Supported Currencies (Partial List)

| ISO Code | Supported Terms/Translations              |
|----------|--------------------------------------------|
| BRL      | REAL, REAIS, BRAZILIAN REAL                |
| USD      | DOLAR, DOLLAR, DÓLAR, US DOLLAR           |
| EUR      | EURO, EUROS                                |
| GBP      | LIBRA ESTERLINA, POUND, BRITISH POUND     |
| JPY      | IENE, YEN, JAPANESE YEN                   |
| AUD      | DOLAR AUSTRALIANO, AUSTRALIAN DOLLAR      |

## 🛠 Customization

**1. Add New Currency**  
```go
// currency_map.go
var CurrencyMap = map[string]string{
    ...
    "PESO ARGENTINO": "ARS",
    "ARGENTINE PESO": "ARS",
}
```

**2. Modify Existing Entries**  
```go
// Add common bitcoin alias
CurrencyMap["XBT"] = "BTC" 
CurrencyMap["BITCOIN"] = "BTC"
```

**3. Case Handling Example**  
```go
// Consumer implementation example
func NormalizeCurrency(input string) string {
    key := strings.ToUpper(strings.TrimSpace(input))
    if code, exists := currency_map.CurrencyMap[key]; exists {
        return code
    }
    return "USD" // Your default
}
```

## 🪪 License
MIT License - See [LICENSE](../LICENSE) for details.

---

**Key Use Cases**:
- Financial data normalization pipelines
- E-commerce transaction processing
- Multi-currency accounting systems
- International payment gateways

**Note**: Always pair with case normalization in your implementation. Contributions for additional currency variants welcome!