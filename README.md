# Examen Mercadolibre

### Pasos para iniciar
1. Clonar el proyecto
2. Importar las dependencias
3. ```go run main.go```
4. Queda corriendo en el puerto 8080

### Pasos para probar servicios en LOCAL
```
TipoMensaje: POST 
Url: http://localhost:8080/mutant/
Body: {
      	"dna":[
      		"AAAAGA",
      		"CAGTGC",
      		"TTATGT",
      		"AGTAGG",
      		"CTCCTA",
      		"TTACTG"
      		]
      }
Respuesta: 200 o 403
```
```
TipoMensaje: GET 
Url: http://localhost:8080/stats
Respuesta: {
               "count_human_dna": 7,
               "count_mutant_dna": 4,
               "ratio": 0.5714286
           }
```

### Pasos para probar servicios en AWS
#### Son los mismos que para local, pero cambiando el url.
- Cambiar "localhost" por "ec2-3-17-193-20.us-east-2.compute.amazonaws.com"
```
    TipoMensaje: POST 
    Url: http://ec2-3-17-193-20.us-east-2.compute.amazonaws.com:8080/mutant/

    TipoMensaje: GET 
    Url: http://ec2-3-17-193-20.us-east-2.compute.amazonaws.com:8080/stats
```

### Notas
1. Se implementó por base de datos [bitcask](https://github.com/prologic/bitcask) porque es una base de datos hecha en Go,
de tipo clave-valor, muy rapida en tiempo de respuesta y de implementar. Una base de datos relacional o documental, tendria muchas 
caracteristicas que no se utilizarian (en este caso solo necesitamos Obtener y Guardar).

2. El archivo `mutantes.conf` y databases se generan dentro de la carpeta del proyecto para facilitar la ejecucion del sistema,
si fuera un sistema productivo hubiera puesto el archivo `mutantes.conf` en alguna ruta absoluta e igual la base de datos. 
Aunque esto trajo como problema al testear que no se encontrara el archivo `mutantes.conf`. 
