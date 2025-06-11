using System.ComponentModel.DataAnnotations;
using Contracts.Base.Models;

namespace UserService.Models;

public class User : IBaseEntity, ISoftDelete
{
    [Key]
    public Guid Id { get; set; }
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
    public DateTime CreatedAt { get; set; }
    public DateTime? ModifiedAt { get; set; }
    public bool IsDeleted { get; set; }
    public DateTime? DeletedAt { get; set; }
}
