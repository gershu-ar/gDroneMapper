public class Config
{
    public string? inputDirectory { get; set; }
    public string? outputDirectory { get; set; }
    public string? log_output_folder { get; set; }
    public bool recurse { get; set; }
    public double distance { get; set; }
    public string? unit { get; set; }
    public bool drawHeading { get; set; }
    public string? language { get; set; }
    public bool scan_images { get; set; }
    public bool scan_videos { get; set; }
    public bool enable_logging { get; set; }
    public bool show_cones { get; set; }
}
