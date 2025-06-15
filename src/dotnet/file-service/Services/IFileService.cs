using file_service.Models;

namespace file_service.Services;

public interface IFileService
{
    Task<IEnumerable<Attachment>> ProcessFilesAsync(List<IFormFile> files);
    Task<Attachment?> GetFileAsync(Guid id);
}