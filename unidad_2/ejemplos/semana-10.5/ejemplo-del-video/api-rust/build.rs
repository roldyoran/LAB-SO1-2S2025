fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .type_attribute(
            ".weathertweet.WeatherTweetResponse",
            "#[derive(serde::Serialize, serde::Deserialize)]",
        )
        .compile_protos(&["proto/weathertweet.proto"], &["proto"])?;
    Ok(())
}