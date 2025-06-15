using file_service.Models;
using file_service.Repositories;
namespace file_service.Services;

public class FileService : IFileService
{
    private readonly IFileRepository _fileRepository;

    public FileService(IFileRepository fileRepository)
    {
        _fileRepository = fileRepository;
    }

    public Task<Attachment?> GetFileAsync(Guid id)
    {
        if (id == Guid.Empty)
        {
            throw new ArgumentException("Invalid file ID.", nameof(id));
        }

        return _fileRepository.Get(id);
    }

    public async Task<IEnumerable<Attachment>> ProcessFilesAsync(List<IFormFile> files)
    {
        var tasks = files.Select(ProcessFile);
        var results = await Task.WhenAll(tasks);

        await _fileRepository.Create(results.ToList());

        return results;
    }

    private async Task<Attachment> ProcessFile(IFormFile file)
    {
        if (file == null || file.Length == 0)
        {
            throw new ArgumentException("File is null or empty.");
        }

        var contentType = file.ContentType.ToLowerInvariant();

        var storagePath = contentType switch
        {
            var ct when ct!.StartsWith("image/") => $"Images",
            var ct when ct!.StartsWith("video/") => $"Videos",
            var ct when ct!.StartsWith("audio/") => $"Audios",
            var ct when ct!.StartsWith("text/") || ct == "application/pdf" || ct == "application/msword" => $"Documents",
            _ => $"Other"
        };

        var fileId = Guid.NewGuid();
        var fileUrl = $"{storagePath}/{fileId}{Path.GetExtension(file.FileName)}";

        // Define two tasks
        await using var stream = new FileStream(Path.Combine(Directory.GetCurrentDirectory(), Constants.FILE_STORAGE_ROOT, fileUrl), FileMode.Create);
        await file.CopyToAsync(stream);

        var attach = new Attachment
        {
            Id = fileId,
            FileName = file.FileName,
            Size = file.Length,
            ContentType = contentType!,
            IsLocalFile = true,
            FileUrl = fileUrl,
            CreatedAt = DateTime.UtcNow,
        };

        // Check type and set metadata if necessary use storagePath -> add real metadata from file
        if (storagePath.Contains("Images"))
        {
            attach.SetMetadata(new { Width = 0, Height = 0 }); // Placeholder for image metadata
        }
        else if (storagePath.Contains("Videos"))
        {
            attach.SetMetadata(new { Duration = 0 }); // Placeholder for video metadata
        }
        else if (storagePath.Contains("Audios"))
        {
            attach.SetMetadata(new { Duration = 0 }); // Placeholder for audio metadata
        }
        else if (storagePath.Contains("Documents"))
        {
            attach.SetMetadata(new { PageCount = 0 }); // Placeholder for document metadata
        }

        return attach;
    }
}
