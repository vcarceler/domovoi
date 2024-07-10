/* This program is free software: you can redistribute it and/or modify it under the
terms of the GNU General Public License as published by the Free Software
Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or
FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
more details.

You should have received a copy of the GNU General Public License along with this
program. If not, see <https://www.gnu.org/licenses/>.  */

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var address string
var port int
var secret string
var command string

func execCommand(w http.ResponseWriter, r *http.Request) {
	sec := r.PathValue("secret")
	p1 := r.PathValue("p1")
	msg := fmt.Sprintf("/command/%s/%s p1=%s remoteaddress=%s", sec, p1, p1, r.RemoteAddr)

	if sec != secret {
		msg = fmt.Sprintf("%s error='bad secret'", msg)
		fmt.Fprintf(w, msg)
		log.Printf(msg)
		return
	}

	// Ejecutamos la tarea
	fmt.Fprintf(w, msg)
	log.Printf(msg)

	// Hacemos flush
	// Ver: https://simon-frey.com/blog/manual-flush-golang-http-responsewriter/
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	cmd := exec.Command(command, p1)

	// Captura la salida estándar del comando
	output, err := cmd.Output()
	if err != nil {
		msg = fmt.Sprintf("Error ejecutando el comando: %v", err)
		fmt.Fprintf(w, msg)
		log.Printf(msg)
		return
	}

	// Imprime la salida del comando
	msg = fmt.Sprintf("\n\nSalida del comando:\n===================\n%s", output)
	fmt.Fprintf(w, msg)
	log.Printf(msg)
}

func main() {
	flag.StringVar(&address, "address", "0.0.0.0", "Dirección para recibir peticiones")
	flag.IntVar(&port, "port", 1234, "Puerto")
	flag.StringVar(&secret, "secret", "DOMOVOI", "Token secreto")
	flag.StringVar(&command, "command", "./mi_script.sh", "Comando a ejecutar")

	flag.Parse()

	log.Printf("domovoi -address %s -port %d -secret XXXX -command %s", address, port, command)

	http.HandleFunc("/command/{secret}/{p1}", execCommand)

	addr := fmt.Sprintf("%s:%d", address, port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
