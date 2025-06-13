using file_service.Models;

namespace file_service.Services;

public class FileService : IFileService
{
    public async Task<IEnumerable<Attachment>> ProcessFilesAsync(List<IFormFile> files)
    {
        var tasks = files.Select(ProcessFile);
        var results = await Task.WhenAll(tasks);
        return results;
    }

    private static async Task<Attachment> ProcessFile(IFormFile file)
    {
        if (file == null || file.Length == 0)
        {
            throw new ArgumentException("File is null or empty.");
        }

        var contentType = file.ContentType.ToLowerInvariant();

        var storagePath = contentType switch
        {
            var ct when ct!.StartsWith("image/") => $"Uploads/Images",
            var ct when ct!.StartsWith("video/") => $"Uploads/Videos",
            var ct when ct!.StartsWith("audio/") => $"Uploads/Audios",
            var ct when ct!.StartsWith("text/") || ct == "application/pdf" || ct == "application/msword" => $"Uploads/Documents",
            _ => $"Uploads/Other"
        };

        var fileId = Guid.NewGuid();
        var fileUrl = $"{storagePath}/{fileId}{Path.GetExtension(file.FileName)}";

        // Save the file to the appropriate location
        using (var stream = new FileStream(Path.Combine(Directory.GetCurrentDirectory(), fileUrl), FileMode.Create))
        {
            await file.CopyToAsync(stream);
        }

        return new Attachment
        {
            Id = fileId,
            FileName = file.FileName,
            Size = file.Length,
            ContentType = contentType!,
            IsLocalFile = true,
            FileUrl = fileUrl,
            CreatedAt = DateTime.UtcNow
        };
    }
}
