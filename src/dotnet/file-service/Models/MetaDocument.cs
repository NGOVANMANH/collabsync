namespace file_service.Models;

public class MetaDocument
{
    public int? PageCount { get; set; }
    public string? Language { get; set; }
    public bool IsEncrypted { get; set; } = false;
}