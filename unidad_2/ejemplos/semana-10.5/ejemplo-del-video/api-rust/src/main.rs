use actix_web::{post, get, web, App, HttpResponse, HttpServer, Responder};
use serde::{Deserialize, Serialize};
use tonic::transport::Channel;

pub mod weathertweet {
    tonic::include_proto!("weathertweet");
}

use weathertweet::{
    weather_tweet_service_client::WeatherTweetServiceClient,
    WeatherTweetRequest, WeatherTweetResponse,
};

#[derive(Debug, Serialize, Deserialize)]
struct WeatherInput {
    municipality: i32,
    temperature: i32,
    humidity: i32,
    weather: i32,
}

#[get("/")]
async fn root() -> impl Responder {
    HttpResponse::Ok().body("hola")
}

#[post("/clima")]
async fn clima(data: web::Json<WeatherInput>) -> impl Responder {
    println!("Rust API recibió JSON: {:?}", data);

    // Usar variable de entorno o fallback a localhost para desarrollo local
    let grpc_url = std::env::var("GRPC_SERVER_URL").unwrap_or_else(|_| "http://0.0.0.0:50051".to_string());
    
    let mut client = match WeatherTweetServiceClient::connect(grpc_url).await {
        Ok(c) => c,
        Err(e) => return HttpResponse::InternalServerError().body(format!("Error conectando: {}", e)),
    };

    let request = tonic::Request::new(WeatherTweetRequest {
        municipality: data.municipality,
        temperature: data.temperature,
        humidity: data.humidity,
        weather: data.weather,
    });

    match client.send_tweet(request).await {
        Ok(response) => {
            let resp: WeatherTweetResponse = response.into_inner();
            HttpResponse::Ok().json(resp)
        }
        Err(e) => HttpResponse::InternalServerError().body(format!("Error gRPC: {}", e)),
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("🚀 Rust REST API corriendo en http://0.0.0.0:8080");
    HttpServer::new(|| {
        App::new()
            .service(root)
            .service(clima)
    })
    .bind(("0.0.0.0", 8080))?
    .run()
    .await
}