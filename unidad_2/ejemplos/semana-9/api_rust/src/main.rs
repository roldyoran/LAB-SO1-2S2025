use actix_web::{web, App, HttpServer, HttpResponse, Result, middleware::Logger};
use serde::{Deserialize, Serialize};
use std::env;

// Estructura para el JSON que recibiremos en /clima
#[derive(Debug, Deserialize, Serialize)]
struct DatosClima {
    ciudad: String,
    temperatura: f32,
    humedad: f32,
    condicion: String,
}

// Handler para el endpoint raíz
async fn hola() -> Result<HttpResponse> {
    log::info!("Endpoint raíz accedido");
    Ok(HttpResponse::Ok().body("¡Hola! Bienvenido a la API del Clima"))
}

// Handler para el endpoint /clima
async fn post_clima(datos: web::Json<DatosClima>) -> Result<HttpResponse> {
    log::info!("Endpoint /clima accedido con datos: {:?}", datos);
    
    // Crear la oración con los datos recibidos
    let oracion = format!(
        "En {} hace {}°C con {}% de humedad y condiciones {}.",
        datos.ciudad, datos.temperatura, datos.humedad, datos.condicion
    );
    
    log::info!("Oración generada: {}", oracion);
    
    Ok(HttpResponse::Ok().json(serde_json::json!({
        "mensaje": oracion,
        "datos_recibidos": datos.into_inner()
    })))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // Configurar logger
    env_logger::init_from_env(env_logger::Env::new().default_filter_or("info"));
    
    let port = 8080;
    
    log::info!("Iniciando servidor en http://localhost:{}", port);
    
    HttpServer::new(|| {
        App::new()
            // Middleware de logging para todos los endpoints
            .wrap(Logger::default())
            .wrap(Logger::new("%a %{User-Agent}i"))
            // Configurar rutas
            .route("/", web::get().to(hola))
            .route("/clima", web::post().to(post_clima))
    })
    .bind(("127.0.0.1", port))?
    .run()
    .await
}