namespace file_service.Models;

public class MetaAudio
{
    public string Codec { get; set; } = null!;
    public double DurationSeconds { get; set; }
    public int SampleRate { get; set; }
    public int Channels { get; set; }
    public long Bitrate { get; set; }
}