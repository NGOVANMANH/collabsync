using BlogService.Models;
using Contracts.Base.Data;
using Microsoft.EntityFrameworkCore;
using MongoDB.EntityFrameworkCore.Extensions;

namespace BlogService.Data;

public class BlogServiceDbContext : IBaseDbContext
{
    public BlogServiceDbContext(DbContextOptions options) : base(options)
    {
        Database.AutoTransactionBehavior = AutoTransactionBehavior.Never;
    }

    public DbSet<Blog> Blogs { get; init; } = null!;
    public DbSet<User> Users { get; init; } = null!;

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        base.OnModelCreating(modelBuilder);
        modelBuilder.Entity<Blog>().ToCollection("blogs");
        modelBuilder.Entity<User>().ToCollection("users");
    }
}
