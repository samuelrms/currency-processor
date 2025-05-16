# 💱 Pacote de Mapeamento de Moedas

Pacote Go para normalização de nomes de moedas em códigos ISO, ideal para processamento de dados financeiros.

![Conversão de Moedas](https://via.placeholder.com/800x200.png?text=Suporte+Multilíngue→Padronização+ISO→Consistência+em+Dados+Financeiros)

## ✨ Funcionalidades

- **Conformidade ISO 4217**: Códigos de 3 letras padrão
- **Suporte Multilíngue**: Termos em português, inglês e variações comuns
- **Aliases Comuns**: Abreviações, erros comuns e termos regionais
- **Estrutura Extensível**: Fácil adição de novas moedas
- **Case Insensitive**: Aceita qualquer formatação de texto (implementação no consumo)

## 🚀 Instalação

```bash
go get github.com/samuelrms/translate-currency/currency_map
```

## 📦 Uso Básico

Importe o pacote e acesse o mapa diretamente:

```go
import "github.com/samuelrms/translate-currency/currency_map"

func main() {
    // Normaliza entrada de moeda
    entrada := "dólar australiano"
    codigoISO := currency_map.CurrencyMap[strings.ToUpper(entrada)]
    
    // Trata valores não mapeados
    if codigoISO == "" {
        codigoISO = "BRL" // Fallback padrão
    }
}
```

## 🌍 Moedas Suportadas (Lista Parcial)

| Código ISO | Termos Suportados/Traduções               |
|------------|--------------------------------------------|
| BRL        | REAL, REAIS, BRAZILIAN REAL                |
| USD        | DOLAR, DOLLAR, DÓLAR, DÓLAR AMERICANO      |
| EUR        | EURO, EUROS                                |
| GBP        | LIBRA ESTERLINA, LIBRA BRITÂNICA           |
| JPY        | IENE, YEN, IENE JAPONÊS                   |
| AUD        | DÓLAR AUSTRALIANO, DOLAR AUSTRALIANO      |

## 🛠 Personalização

**1. Adicionar Nova Moeda**  
```go
// currency_map.go
var CurrencyMap = map[string]string{
    ...
    "PESO ARGENTINO": "ARS",
    "PESO MEXICANO":  "MXN",
}
```

**2. Modificar Entradas Existentes**  
```go
// Adicionar variação comum
CurrencyMap["DOLAR CANADENSE"] = "CAD" 
CurrencyMap["DINHEIRO"] = "BRL" // Gíria
```

**3. Exemplo de Normalização**  
```go
// Implementação de consumo
func NormalizarMoeda(entrada string) string {
    chave := strings.ToUpper(strings.TrimSpace(entrada))
    if codigo, existe := currency_map.CurrencyMap[chave]; existe {
        return codigo
    }
    return "BRL" // Seu fallback
}
```

## 🪪 Licença
MIT License - Veja [LICENÇA](LICENSE) para detalhes.

---

**Casos de Uso Comuns**:
- Sistemas de pagamento internacional
- Processamento de transações em e-commerce
- Conversão de dados contábeis
- Integração com gateways de pagamento

**Nota**: Implemente normalização de case no seu código. Contribuições com novos termos são bem-vindas!