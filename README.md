# Go-Commons

Este repositorio contiene funciones de utilidad para proyectos de microservicios con Go.



### Base64Encode 

Codifica una cadena como base64

```go
import (
	"github.com/LuisEGR/go-commons"
)

b64 := commons.Base64Encode("Hola Mundo")
print(b64) // SG9sYSBNdW5kbw
```

### Base64Decode

Decodifica una cadena de base64 a utf
```go
import (
	"github.com/LuisEGR/go-commons"
)

txt := commons.Base64Decode("RGVjb2RpZmljYWRv")
print(txt) // Decodificado
```

### GetDataJWT

Obtiene un valor del payload de un JWT, utilizando gjson para obtener el valor.
```go
import (
	"github.com/LuisEGR/go-commons"
)

// Token Payload:
// {
//     idUser: "12312"
// }

token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJpZFVzZXIiOiIxMjMxMiJ9.5cqpNmlEcroWY7w-xj-zkO9p7Vz2-VaOFUOPvOblWzc"
value, _ := commons.GetDataJWT(token, "idUser")
print(value) // 12312
```
más información en https://github.com/tidwall/gjson


### GetLogger
Obtiene una instancia de logrus-loger, preconfigurada para trabajar con parámetros específicos dependiendo de la variable de entorno `ENVIRONMENT`.

```go
import (
	"github.com/LuisEGR/go-commons"
)

log := commons.GetLogger()

log.Trace("Something very low level.")
log.Debug("Useful debugging information.")
log.Info("Something noteworthy happened!")
log.Warn("You should probably take a look at this.")
log.Error("Something failed but I'm not quitting.")
// Calls os.Exit(1) after logging
log.Fatal("Bye.")
// Calls panic() after logging
log.Panic("I'm bailing.")
```
más información en https://github.com/sirupsen/logrus
