
using file_service.Models;

namespace file_service.Repositories;

public interface IFileRepository
{
    Task<Guid> Create(Attachment attachment);
    Task<List<Guid>> Create(List<Attachment> attachments);
    Task Update(Guid id, Attachment attachment);
    Task Delete(Guid id);
    Task<Attachment?> Get(Guid id);
    Task<List<Attachment>> Get();
    Task<int> Save();
}