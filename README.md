[![Domovoi por Iván Bilibin](https://upload.wikimedia.org/wikipedia/commons/8/84/Domovoi_Bilibin.jpg)](https://es.wikipedia.org/wiki/Iv%C3%A1n_Bilibin)
# Домовой / Domovoi

Domovoi es un servicio web permite lanzar un comando a partir de la solicitud de un cliente.

Está pensado para, por ejemplo, lanzar un `playbook` de Ansible que prepara portátiles para los alumnos. Pero puede lanzar cualquier comando pasando los parámetros indicados en la petición web.

Cuando se lanza `domovoi` (normalmente utilizando `domovoi.service`) se especifica el comando a ejecutar. Cada vez que `domovoi` reciba una petición se ejecutará el comando indicado.

# Instalación de Domovoi

Será suficiente con copiar el ejecutable en algún directorio.

Es posible compilar Domovoi clonando el repositorio y ejecutando:

~~~
cd domovoi
go build
~~~

Lo que producirá el ejecutable `domovoi`.

# Ejecución de Domovoi

Será posible ejecutar `domovoi` utilizando los parámetros por defecto:

~~~
vcarceler@baba-yaga-2404:~$ ./domovoi 
2024/07/06 15:56:00 domovoi -address 0.0.0.0 -port 1234 -secret XXXX -command ./mi_script.sh
~~~

En este caso `domovoi`:

 * Se conectará a todas las interfaces de red (`0.0.0.0`).
 * Atenderá en el puerto `1234`.
 * Se utilizará como `secret` la cadena `DOMOVOI`.

Pero se podrá indicar un valor adecuado para cualquiera de estos parámetros:

 ~~~
vcarceler@baba-yaga-2404:~$ ./domovoi --help
Usage of ./domovoi:
  -address string
    	Dirección para recibir peticiones (default "0.0.0.0")
  -command string
    	Comando a ejecutar (default "./mi_script.sh")
  -port int
    	Puerto (default 1234)
  -secret string
    	Token secreto (default "DOMOVOI")
vcarceler@baba-yaga-2404:~$
 ~~~

El parámetro `secret` permite especificar la cadena que se utilizará en las peticiones. Las peticiones que no incluyan el valor correcto de `secret` no se atenderán.

Es posible utilizar una unidad de `systemd` para lanzar `domovoi`.

Por ejemplo:

~~~
[Unit]
Description=Domovoi recibe peticiones de los clientes web y lanza comandos en el sistema

[Service]
User=vcarceler
Group=vcarceler
Restart=on-failure
WorkingDirectory=/home/vcarceler
ExecStart=/opt/domovoi -command /home/vcarceler/playbooks-elpuig/cron/domovoi-init-alumnes-u2404

[Install]
WantedBy=default.target
~~~

Cada vez que `domovoi` reciba una petición `<ip>:<port>/command/<secret>/<parameters>` ejecutará el script `/home/vcarceler/playbooks-elpuig/cron/domovoi-init-alumnes-u2404` pasándole `<parameters>`.

Durante el funcionamiento se irán registrando las solicitudes recibidas y la salida de la ejecución tanto en el navegador como en los ficheros de registro.

~~~
jul 06 13:05:29 baba-yaga-2404 domovoi[575770]: 2024/07/06 13:05:29 domovoi -address 0.0.0.0 -port 1234 -secret XXXX -command />
jul 06 13:05:43 baba-yaga-2404 domovoi[575770]: 2024/07/06 13:05:43 /command/DOMOVOI/10.0.3.120 p1=10.0.3.120 remoteaddress=10.>
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]: 2024/07/06 13:12:52
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]: Salida del comando:
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]: ===================
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]: Operations to perform:
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]:   Apply all migrations: admin, api, auth, contenttypes, db, sessions
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]: Running migrations:
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]:   No migrations to apply.
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]: PLAY [portatils_alumnes_u2404] *************************************************
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]: TASK [Gathering Facts] *********************************************************
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]: ok: [10.0.3.120]
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]: TASK [delete-users : Borrado del usuario usuariinstall] ************************
jul 06 13:12:52 baba-yaga-2404 domovoi[575770]: ok: [10.0.3.120]
~~~



## Built with ❤️

* [Go](https://go.dev/) - Build simple, secure, scalable systems with Go.
* [GNU/Linux](https://es.wikipedia.org/wiki/GNU/Linux) - Un sistema operativo libre.

## Authors

* Victor Carceler

## License

This project is licensed under the GNU General Public License v3.0 - see the [COPYING](COPYING) file for details.