using Contracts.Base.Data;
using Microsoft.EntityFrameworkCore;
using UserService.Models;

namespace UserService.Data;

public class UserServiceDbContext : IBaseDbContext
{
    public UserServiceDbContext(DbContextOptions options) : base(options)
    {
    }

    public DbSet<User> Users { get; set; } = null!;
}
