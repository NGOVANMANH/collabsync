namespace file_service.Models;

public class MetaVideo
{
    public string Codec { get; set; } = null!;
    public int Width { get; set; }
    public int Height { get; set; }
    public double FrameRate { get; set; }
    public TimeSpan DurationSeconds { get; set; }
    public long Bitrate { get; set; }
}