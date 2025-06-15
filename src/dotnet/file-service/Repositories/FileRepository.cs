using file_service.Data;
using file_service.Models;
using Microsoft.EntityFrameworkCore;

namespace file_service.Repositories;

public class FileRepository : IFileRepository
{
    private readonly FileDbContext _context;

    public FileRepository(FileDbContext context)
    {
        _context = context;
    }
    public async Task<Guid> Create(Attachment attachment)
    {
        await _context.Attachments.AddAsync(attachment);
        await _context.SaveChangesAsync();
        return attachment.Id;
    }

    public async Task<List<Guid>> Create(List<Attachment> attachments)
    {
        if (attachments == null || !attachments.Any())
        {
            throw new ArgumentException("Attachments list cannot be null or empty.", nameof(attachments));
        }

        _context.Attachments.AddRange(attachments);
        await _context.SaveChangesAsync();
        return attachments.Select(a => a.Id).ToList();
    }

    public async Task Delete(Guid id)
    {
        var attachment = await _context.Attachments.FindAsync(id);
        if (attachment != null)
        {
            _context.Attachments.Remove(attachment);
            await _context.SaveChangesAsync();
        }
        else
        {
            throw new KeyNotFoundException($"Attachment with ID {id} not found.");
        }
    }

    public Task<Attachment?> Get(Guid id)
    {
        return _context.Attachments.FindAsync(id).AsTask();
    }

    public Task<List<Attachment>> Get()
    {
        return _context.Attachments.ToListAsync();
    }

    public async Task<int> Save()
    {
        return await _context.SaveChangesAsync();
    }

    public Task Update(Guid id, Attachment attachment)
    {
        throw new NotImplementedException();
    }
}
