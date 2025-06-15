using file_service.Models;
using Microsoft.EntityFrameworkCore;

namespace file_service.Data;

public class FileDbContext : DbContext
{
    public FileDbContext(DbContextOptions<FileDbContext> options) : base(options)
    {
    }

    public DbSet<Attachment> Attachments { get; set; } = null!;
}