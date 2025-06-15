using System.ComponentModel.DataAnnotations;
using System.Text.Json;
using file_service.Config;

namespace file_service.Models;

public class Attachment
{
    [Key]
    public Guid Id { get; set; }
    public string FileName { get; set; } = null!;
    public long Size { get; set; }
    public string ContentType { get; set; } = null!;
    public bool IsLocalFile { get; set; }
    public string FileUrl { get; set; } = null!;
    public string? Metadata { get; set; } = null!;
    public DateTime CreatedAt { get; set; }

    public TMetadata? GetMetadata<TMetadata>()
          where TMetadata : class, new()
    {
        return string.IsNullOrWhiteSpace(Metadata)
            ? new TMetadata()
            : JsonSerializer.Deserialize<TMetadata>(Metadata);
    }

    public void SetMetadata<TMetadata>(TMetadata metadata)
    {
        Metadata = JsonSerializer.Serialize(metadata, JsonOptions.CamelCase);
    }
}
