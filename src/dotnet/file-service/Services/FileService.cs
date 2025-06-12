using file_service.Extensions;
using file_service.Models;

namespace file_service.Services;

public class FileService : IFileService
{
    public Task<IEnumerable<Attachment>> ProcessFilesAsync(List<IFormFile> files)
    {
        var filesProcessing = files.Select(file =>
        {
            return ProcessFile(file);
        });

        return Task.FromResult(
            files.Select(file => new Attachment
            {
                FileName = file.FileName,
                ContentType = file.ContentType,
                Size = file.Length,
            })
        );
    }

    private static Attachment ProcessFile(IFormFile file)
    {
        if (file == null || file.Length == 0)
        {
            throw new ArgumentException("File is null or empty.");
        }

        if (file.Length > 10 * Constants.MB)
        {
            throw new ArgumentException($"File {file.FileName} exceeds the maximum allowed size of 10 MB.");
        }

        var fileId = Guid.NewGuid();

        switch (file.ContentType.ToLowerInvariant().Split('/')[0])
        {
            case "video":
            case "image":
            case "audio":
            default: break;
        }

        return new Attachment
        {
            Id = fileId,
            FileName = file.FileName,
            Size = file.Length,
            ContentType = file.ContentType,
            IsLocalFile = true,
            FileUrl = $"/images/{fileId}",
            CreatedAt = DateTime.UtcNow
        };
    }
}
