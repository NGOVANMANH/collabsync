namespace file_service.Models;

public class Attachment
{
    public Guid Id { get; set; }
    public string FileName { get; set; } = null!;
    public long Size { get; set; }
    public string ContentType { get; set; } = null!;
    public bool IsLocalFile { get; set; } = true;
    public string FileUrl { get; set; } = null!;
}
