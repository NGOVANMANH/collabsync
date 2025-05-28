using System.ComponentModel.DataAnnotations;
using Contracts.Base.Models;
using Microsoft.EntityFrameworkCore;
using MongoDB.Bson.Serialization.Attributes;

namespace BlogService.Models;

[Index(nameof(Email), IsUnique = true)]
[Index(nameof(UserName), IsUnique = true)]
public class User : IBaseEntity
{
    [BsonId]
    public Guid Id { get; set; }
    public DateTime CreatedAt { get; set; }
    public DateTime? ModifiedAt { get; set; }
    [Required]
    [MaxLength(50)]
    [MinLength(3)]
    [RegularExpression(@"^[a-zA-Z0-9_]+$", ErrorMessage = "Username can only contain letters, numbers, and underscores.")]
    public string UserName { get; set; } = null!;

    [Required]
    [MaxLength(100)]
    [MinLength(5)]
    [EmailAddress(ErrorMessage = "Invalid email address.")]
    public string Email { get; set; } = null!;
    public DateTime DateOfBirth { get; set; }
    public bool IsFemale { get; set; }
    public string Avatar { get; set; } = null!;

    public IEnumerable<Blog> Blogs { get; set; } = new List<Blog>();
}
