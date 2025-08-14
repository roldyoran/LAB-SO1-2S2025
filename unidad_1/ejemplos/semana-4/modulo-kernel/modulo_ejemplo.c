// Incluye las cabeceras necesarias para crear un módulo del kernel de Linux
#include <linux/module.h>   // Funciones y macros para módulos
#include <linux/kernel.h>   // Funciones para imprimir mensajes en el kernel
#include <linux/init.h>     // Macros para marcar funciones de inicio y fin


// Información del módulo
MODULE_LICENSE("GPL"); // Tipo de licencia (obligatorio)
MODULE_AUTHOR("autor"); // Autor del módulo
MODULE_DESCRIPTION("Un módulo de kernel simple que muestra mensajes al cargar y descargar"); // Descripción


// Esta función se ejecuta cuando el módulo es cargado en el kernel
static int __init modulo_inicio(void)
{
    printk(KERN_INFO "Módulo cargado\n"); // Imprime un mensaje en el log del sistema
    printk(KERN_INFO "¡Bienvenido! El módulo se ha insertado correctamente.\n"); // Segundo mensaje
    return 0; // Indica que la carga fue exitosa
}


// Esta función se ejecuta cuando el módulo es eliminado del kernel
static void __exit modulo_fin(void)
{
    printk(KERN_INFO "Módulo eliminado\n"); // Imprime un mensaje al eliminar el módulo
    printk(KERN_INFO "¡Adiós! El módulo ha sido removido del kernel.\n"); // Segundo mensaje
}


// Macros que indican cuál es la función de inicio y de salida del módulo
module_init(modulo_inicio); // Al cargar el módulo, se llama a modulo_inicio
module_exit(modulo_fin);    // Al eliminar el módulo, se llama a modulo_fin