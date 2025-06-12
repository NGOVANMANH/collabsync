using System.Net.Mail;

namespace file_service.Repositories;

public interface IVideoRepository
{
    Task<Guid> Create();
    Task<bool> Update(Guid id, string fileName, long size, string contentType, bool isLocalFile, string fileUrl);
    Task<bool> Delete(Guid id);
    Task<Attachment?> GetById(Guid id);
    Task<IEnumerable<Attachment>> GetAll();
}