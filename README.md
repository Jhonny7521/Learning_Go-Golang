# :abc: :full_moon_with_face: APRENDIENDO GO/GOLANG :full_moon_with_face: :abc:

## INTRODUCCION

### ¿Qué es Go (Golang)?

- Go es un lenguaje de programación compilado.
- Es un proyecto OpenSource.
- Fue creado por Google en el 2007.
- Sus autores son Robert Griesemen (Suiza), Rob Pike (Canada) y Ken Thompson (USA).
- Fue lanzado al mercado el 10 de Noviembre del 2009.
- Fue creado con la intención de combinar el dinamismo de Python y el performance de C o C++.
- Es la solución interna de Google al aprovechamiento de los procesadores multicore.

### ¿Cuáles son las diferencias entre Go y lenguajes como PHP, Python y Ruby?

- El compilador convierte el código de Go en código binario.
- Al realizar la compilación revisa si existen errores.
- El archivo compilado no requiere un interprete para ejecutarse.
- Se puede compilar directamente desde y a cualquier sistema operativo (Windows, Linux, Mac, etc) y/o arquitectura (32 o 64 bits).

### ¿Qué empresas utilizan Go?

- Uber
- Google
- Docker
- Kubernetes
- Twitch
- MongoDB
- Dropbox

### ¿Cuáles son algunas de las ventajas de usar Go?

- Es sumamente rápido.
- No requiere de un interprete o una maquina virtual.
- Muy sencillo de usar debido a su sintaxis.
- Multiplataforma, se puede ejecutar en Windows, Mac, Linux, etc.
- Se puede compilar directamente desde y a cualquier SO y/o arquitectura.
- Programación asíncrona sencilla de utilizar (rutinas).
- El tiempo de compilación es prácticamente inmediato.
- Obliga al programador a tomar conciencia real sobre el manejo de errores.
- Viene precargado (no requiere de frameworks).

### ¿Cuáles son algunas de las las desventajas de usar Go?

- Sigue siendo un lenguaje joven por lo que carece de la madurez de otros lenguajes como Java, Python, PHP, etc.
- Comparado con otros lenguajes como Python, PHP o Ruby, resulta ser algo improductivo, pues carece de muchos atajos que tienen estos otros lenguajes.
- La documentación no es la mas sencilla de interpretar para quienes comienzan a utilizar el lenguaje.
- Aun carece de paquetes oficiales para manejo de elementos importantes como algunas bases de datos.
- No existe un mercado laboral tan grande como en otros lenguajes.

### ¿Qué tipo de aplicaciones se pueden crear con Go?

- Páginas web dinámicas.
- REST API’s.
- Servicios de red.
- Aplicaciones de línea de comandos.

### ¿Qué tipo de aplicaciones no se pueden crear con Go?

- Aplicaciones de escritorio.
- Aplicaciones para dispositivos móviles (Android, iOS, etc).

### Recursos

- [Documentación de Go](https://go.dev/doc/)
- [Lista de paquetes oficiales](https://pkg.go.dev/std)
- [Buscador de paquetes](https://pkg.go.dev/)


## INSTALACIÓN DE GO/GOLANG (Linux)

### Descargar archivos de instalación de Go en Linux

* Descargar el instalador de https://golang.org/doc/install.

* Descomprimir el archivo dentro de /usr/local.

  $ tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
  Por lo general estos comandos requieren el uso de sudo.

### Verificar la instalación de Go en Linux

  $ go version  

### Definir las variables de entorno de Go en Linux

* Crear el directorio del workspace de Go.

* Agregar la variable GOPATH al archivo .profile

  $ GOPATH="$HOME/go"

* Agregar Go al PATH en el archivo .profile.

  $ PATH="/usr/local/go/bin:$PATH"

### Verificar la variables de entorno de Go en Linux

* Verificar que la variable de entorno GOPATH este definida.

  $ echo $GOPATH

* Verificar que la variable de entorno PATH contenga la ruta a los binarios de Go.

  $ echo $PATH

  [Link - video instructivo](https://www.youtube.com/watch?v=E1yiCT2Rdj8&list=PLt1J5u9LpM5-L-Ps8jjr91pKhFxAnxKJp&index=4)


##  INSTALACIÓN DE GO/GOLANG (Windows)

### Descargar archivos de instalación de Go en Windows

* Descargar el instalador de [Golang.org](https://golang.org/doc/install).

* Ejecutar el instalador, una ves completado el proceso de instalación verificar que Go esta presente en el sistema.

### Verificar la instalación de Go en Windows

  $ go version

### Definir las variables de entorno de Go en Windows

* El instalador va a agregar el path c:\Go\bin a tus variables de entorno, pero quizás sea necesario reiniciar el command prompt.

* Crea un directorio para tu workspace en %USERPROFILE%\go.

### Verificar la variables de entorno de Go en Windows

  $ echo %GOPATH%

  [Link - video instructivo](https://www.youtube.com/watch?v=9LWX7FKqbUc&list=PLt1J5u9LpM5-L-Ps8jjr91pKhFxAnxKJp&index=30)


## INSTALACIÓN DE GO/GOLANG (macOs)

### Descargar archivos de instalación de Go en macOS

* Descargar el instalador de [Golang.org](https://golang.org/doc/install).

* Ejecutar el instalador.

### Verificar la instalación de Go en macOS

  $ go version

### Definir las variables de entorno de Go en macOS

* Crear el directorio del workspace de Go.

* Agregar la variable GOPATH al archivo .profile

  $ GOPATH="$HOME/go"

* Agregar Go al PATH en el archivo .profile.

  $ PATH="/usr/local/go/bin:$PATH"

* Reinicia tu terminal.

### Verificar la variables de entorno de Go en macOS

* Verificar que la variable de entorno GOPATH este definida.

  $ echo $GOPATH

* Verificar que la variable de entorno PATH contenga la ruta a los binarios de Go.

  $ echo $PATH

  [Link - video instructivo](https://www.youtube.com/watch?v=MbS1wn0B-fk)

## EDITOR DE CÓDIGO

* El editor a utilizar es a preferencia de cada uno. En mi caso será Visual Studio Code

- [Link de descarga de Visual Studio Code](https://code.visualstudio.com/download)
- [Extension para Golang en VSC](https://marketplace.visualstudio.com/items?itemName=golang.Go)

* Algunos de los editores de código para Golang son:
  - Sublime
  - Atom
  - Textmate
  - Vim
  - Emacs


### FUENTES DE INFORMACIÓN

- Página Web: [Apuntes de Programador](https://apuntes.de/golang/introduccion-a-golang/#gsc.tab=0)
- Canal de YouTube: [Roelcode](https://www.youtube.com/c/Roelcode)
- Canal de YouTube: [Bitfumes](https://www.youtube.com/c/Bitfumes)
