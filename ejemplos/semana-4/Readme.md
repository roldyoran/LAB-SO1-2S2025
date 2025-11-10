### 1. **Prepara tu entorno de desarrollo:**
   - Asegúrate de tener instalado un compilador de kernel, como `gcc`.
   - Ten acceso a los encabezados del kernel (`kernel headers`). En Ubuntu y sistemas basados en el mismo, puedes instalarlos con:

     ```bash
     sudo apt-get install gcc linux-headers-$(uname -r)
     ```

### 2. **Crea el archivo del módulo:**
   - Abre un editor de texto y copia el código proporcionado en un archivo con extensión `.c`, por ejemplo, `sysinfo.c`.

### 3. **Crea un Makefile para compilar el módulo:**
   - En el mismo directorio donde guardaste `sysinfo.c`, crea un archivo llamado `Makefile` con el siguiente contenido:

     ```makefile
     obj-m += sysinfo.o

     all:
        make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules

     clean:
        make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean
     ```

   - Este `Makefile` compila el módulo para tu versión actual del kernel.

### 4. **Compila el módulo:**
   - En la terminal, navega al directorio donde guardaste `sysinfo.c` y `Makefile`. Luego ejecuta:

     ```bash
     make
     ```

   - Si todo está correcto, se generará un archivo con la extensión `.ko` (por ejemplo, `sysinfo.ko`), que es tu módulo de kernel.

### 5. **Carga el módulo en el kernel:**
   - Usa `insmod` para cargar el módulo:

     ```bash
     sudo insmod sysinfo.ko
     ```

   - Para verificar que el módulo se ha cargado, usa:

     ```bash
     lsmod | grep sysinfo
     ```

   - También puedes verificar si el archivo `/proc/sysinfo` ha sido creado:

     ```bash
     cat /proc/sysinfo
     ```

   - Deberías ver el JSON con la información del sistema.

### 6. **Desinstala el módulo cuando termines:**
   - Para remover el módulo del kernel, usa:

     ```bash
     sudo rmmod sysinfo
     ```

   - Verifica que el módulo fue removido correctamente:

     ```bash
     lsmod | grep sysinfo
     ```



### 7. **Depura y ajusta según sea necesario:**
   - Si encuentras errores, puedes revisar los mensajes del kernel usando:

     ```bash
     dmesg | tail
     ```

   - Ajusta el código del módulo para satisfacer todos los requisitos y, si es necesario, recompila el módulo.
