fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .type_attribute(
            ".weathertweet.WeatherTweetResponse",
            "#[derive(serde::Serialize, serde::Deserialize)]",
        )
        .compile(&["proto/weathertweet.proto"], &["proto"])?;
    Ok(())
}