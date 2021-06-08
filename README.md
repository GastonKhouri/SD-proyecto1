# Sistemas Distribuidos - Proyecto #1

## Características del proyecto

El proyecto cuenta con las siguientes características:

- Servicio de Cuentas
- Servidor UDP
- Servidores TCP (procesos e hilos)
- Cliente UDP – Con interfaz a usuario
- Cliente TCP – Con Interfaz a usuario
- Consola local – Con interfaz a usuario 

Donde el servicio de Cuentas, mantiene las cuenta y los diferentes servidores se comunican con él vía una cola de mensajesy el accede, vía un semáforo, la cuenta en memoria compartida.

La consola local provee una interfaz de usuario en una terminal local que permite manipular la cuenta. Para lo cual accede, vía un semáforo, la cuenta en memoria compartida, de la misma forma provee información del servicio, la cola de mensaje y servidores (UDP y TCP). El servicio, así como los servidores deben provee la capacidad de monitorear suactividad.

La consola remota ofrece las mismas funcionalidades de la consola local, y para su implementaciónse debe hacer uso de RPC (llamado a procedimientos remotos) y la función remota en el servidor debe accede, vía un semáforo, la cuenta en memoria compartida.

## Integrantes
- Laura Andara, 27.186.209
- Gastón Lozano, 27.955.802
- Samuel Mesa, 26.116.630
- Florentino Muñoz, 27.077.374
