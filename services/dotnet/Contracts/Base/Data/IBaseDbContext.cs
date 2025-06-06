using Contracts.Base.Models;
using Contracts.Extensions;
using Microsoft.EntityFrameworkCore;

namespace Contracts.Base.Data;

public abstract class IBaseDbContext : DbContext
{
    public IBaseDbContext(DbContextOptions options) : base(options)
    {
    }

    public override Task<int> SaveChangesAsync(CancellationToken cancellationToken = default)
    {
        // Update timestamps for added or modified entities
        UpdateTimestamps();
        // Create slugs for slugable entities
        CreateSlugs();

        // Call the base SaveChangesAsync method
        return base.SaveChangesAsync(cancellationToken);
    }

    public override int SaveChanges()
    {
        // Update timestamps for added or modified entities
        UpdateTimestamps();
        // Create slugs for slugable entities
        CreateSlugs();

        // Call the base SaveChanges method
        return base.SaveChanges();
    }

    protected virtual void UpdateTimestamps()
    {
        var entries = ChangeTracker.Entries()
            .Where(e => e.State == EntityState.Added || e.State == EntityState.Modified);

        var now = DateTime.UtcNow;

        foreach (var entry in entries)
        {
            if (entry.Entity is IBaseEntity entity)
            {
                if (entry.State == EntityState.Added)
                {
                    entity.CreatedAt = now;
                }
                else if (entry.State == EntityState.Modified)
                {
                    entity.ModifiedAt = now;
                }
            }
        }
    }

    protected virtual void CreateSlugs()
    {
        var entries = ChangeTracker.Entries()
            .Where(e => e.State == EntityState.Added || e.State == EntityState.Modified);

        foreach (var entry in entries)
        {
            if (entry.Entity is ISlugable slugable && string.IsNullOrWhiteSpace(slugable.Slug))
            {
                slugable.Slug = slugable.Title.ToSlug();
            }
        }
    }

    protected virtual void SoftDeleteEntities()
    {
        var entries = ChangeTracker.Entries()
            .Where(e => e.State == EntityState.Deleted && e.Entity is ISoftDelete softDelete);

        foreach (var entry in entries)
        {
            if (entry.Entity is ISoftDelete softDelete && softDelete.IsDeleted == false)
            {
                softDelete.IsDeleted = true;
                softDelete.DeletedAt = DateTime.UtcNow;
                entry.State = EntityState.Modified;
            }
        }
    }
}
